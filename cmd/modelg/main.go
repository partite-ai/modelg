package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"go/format"
	"go/types"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

//go:embed package_gen.tmpl
var packageTemplateText string
var packageTemplate = template.Must(template.New("package").Parse(packageTemplateText))

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	if err := run(ctx); err != nil {
		cancel()
		log.Fatalf("Error: %v", err)
	}
}

func run(ctx context.Context) error {
	cf, err := os.Open("modelg.yaml")
	if err != nil {
		return err
	}
	defer cf.Close()

	cnf, err := LoadConfig(cf)
	if err != nil {
		return err
	}

	// Load the Go packages
	needPkgs := []string{
		"",
		"context",
		"github.com/partite-ai/modelg/cmd/modelg",
	}

	for _, converter := range cnf.Converters {
		if converter.Scanner != "" {
			lastDot := strings.LastIndex(converter.Scanner, ".")
			if lastDot == -1 || lastDot == len(converter.Scanner)-1 {
				return fmt.Errorf("invalid scanner function %q", converter.Scanner)
			}

			pkgName := converter.Scanner[:lastDot]
			if !slices.Contains(needPkgs, pkgName) {
				needPkgs = append(needPkgs, pkgName)
			}
		}

		if converter.Valuer != "" {
			lastDot := strings.LastIndex(converter.Valuer, ".")
			if lastDot == -1 || lastDot == len(converter.Valuer)-1 {
				return fmt.Errorf("invalid valuer function %q", converter.Valuer)
			}

			pkgName := converter.Valuer[:lastDot]
			if !slices.Contains(needPkgs, pkgName) {
				needPkgs = append(needPkgs, pkgName)
			}
		}
	}

	currentPackage, _, err := loadPackages(ctx, needPkgs)
	if err != nil {
		return err
	}

	packageImportSet := newImportSet(currentPackage.PkgPath, map[string]string{})
	packageImportSet.addImport("database/sql/driver", "driver")
	packageImportSet.addImport("github.com/partite-ai/modelg", "modelg")

	var buf bytes.Buffer
	err = packageTemplate.Execute(&buf, map[string]interface{}{
		"PackageName": currentPackage.Name,
		"Imports":     strings.Join(packageImportSet.importsList(), "\n"),
	})

	if err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	packageFile, err := os.Create("pkg_gen.go")
	if err != nil {
		return err
	}
	defer packageFile.Close()

	_, err = packageFile.Write(formatted)
	if err != nil {
		return err
	}

	modelInfos := make(map[string]*modelInfo)
	// First generate all param structs - we'll need them for the queries
	for _, model := range cnf.Models {
		obj := currentPackage.Types.Scope().Lookup(model.Name)
		if obj == nil {
			return fmt.Errorf("model %q not found", model.Name)
		}
		modelType, ok := obj.Type().Underlying().(*types.Struct)
		if !ok {
			return fmt.Errorf("model %q is not a struct", model.Name)
		}

		modelInfo, err := LoadModelInfo(model.Name, modelType, newImportSet(currentPackage.PkgPath, map[string]string{}))
		if err != nil {
			return err
		}
		modelInfos[model.Name] = modelInfo

		if modelInfo.PKType.PackagePath != "" {
			if !slices.Contains(needPkgs, modelInfo.PKType.PackagePath) {
				needPkgs = append(needPkgs, modelInfo.PKType.PackagePath)
			}
		}

		if err := CreateModelFile(currentPackage, model, modelType); err != nil {
			return err
		}
	}

	// Load the packages again to get the updated packages
	currentPackage, packageMap, err := loadPackages(ctx, needPkgs)
	if err != nil {
		return err
	}

	// Now that we have the updated packages, load the converters
	var converters []converterInfo

	for _, converter := range cnf.Converters {
		ci := converterInfo{
			Name: converter.Name,
		}

		if converter.Scanner != "" {
			if err := ci.loadScannerFunction(packageMap, converter.Scanner); err != nil {
				return err
			}
		}

		if converter.Valuer != "" {
			if err := ci.loadValuerFunction(packageMap, converter.Valuer); err != nil {
				return err
			}
		}

		converters = append(converters, ci)
	}

	// Now generate the queries
	for _, model := range cnf.Models {
		queryFileName := strcase.ToSnake(model.Name) + ".sql"
		querySet, err := LoadQuerySet(queryFileName)
		if err != nil {
			return err
		}

		var queryIntf *types.Interface
		queriesName := model.Queries
		if queriesName == "" {
			queriesName = strcase.ToLowerCamel(model.Name) + "Queries"
		}
		queriesObj := currentPackage.Types.Scope().Lookup(queriesName)
		if queriesObj == nil {
			if model.Queries != "" {
				return fmt.Errorf("queries %q not found", queriesName)
			}
		} else {
			var ok bool
			queryIntf, ok = queriesObj.Type().Underlying().(*types.Interface)
			if !ok {
				return fmt.Errorf("queries %q is not an interface", model.Queries)
			}
		}

		if queryIntf != nil {
			querySet.AddInterface(queriesObj.Type())
		}

		err = addCRUDMethods(querySet, currentPackage, packageMap, modelInfos[model.Name])
		if err != nil {
			return err
		}

		if err := querySet.CreateQueriesFile(currentPackage, model.Name, converters); err != nil {
			return err
		}
	}

	return nil
}

func loadPackages(ctx context.Context, needPackages []string) (*packages.Package, map[string]*packages.Package, error) {
	pkgs, err := packages.Load(&packages.Config{
		Context: ctx,
		Mode:    packages.NeedName | packages.NeedTypes | packages.NeedImports | packages.NeedTypesInfo | packages.NeedDeps | packages.NeedFiles,
	}, needPackages...)
	if err != nil {
		return nil, nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	cwd, err = filepath.Abs(cwd)
	if err != nil {
		return nil, nil, err
	}

	var currentPkg *packages.Package
	pkgMap := make(map[string]*packages.Package)
	for _, p := range pkgs {
		pkgMap[p.PkgPath] = p
		if len(p.GoFiles) > 0 {
			path, err := filepath.Abs(filepath.Dir(p.GoFiles[0]))
			if err != nil {
				return nil, nil, err
			}

			if cwd == path {
				currentPkg = p
			}
		}
	}

	return currentPkg, pkgMap, nil
}
