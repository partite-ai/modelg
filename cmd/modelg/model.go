package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/packages"
)

//go:embed model_gen.tmpl
var modelTemplateText string
var modelTemplate = template.Must(template.New("model").Parse(modelTemplateText))

type modelInfo struct {
	Name              string
	PKName            string
	PKType            pkType
	CreateFields      []fieldInfo
	UpdateFields      []fieldInfo
	SelectFields      []fieldInfo
	TableName         string
	DisplayName       string
	PluralDisplayName string
	FieldDisplayNames map[string]string
}

type fieldInfo struct {
	Name          string
	Type          string
	Tag           string
	ConverterName string
	DatabaseName  string
	DisplayName   string
}

type pkType struct {
	PackagePath string
	TypeName    string
}

func LoadModelInfo(model string, typ *types.Struct, imports *importSet) (*modelInfo, error) {
	var createFields []fieldInfo
	var updateFields []fieldInfo
	var selectFields []fieldInfo
	var pkName string
	var pkType pkType
	nameMap := make(map[string]string)

	for i := 0; i < typ.NumFields(); i++ {
		field := typ.Field(i)
		if !field.Exported() {
			continue
		}
		tag := reflect.StructTag(typ.Tag(i))
		tagValue := tag.Get("modelg")

		computed := false
		isPK := false
		immutable := false
		converterName := ""

		if tagValue != "" {
			for part := range strings.SplitSeq(tagValue, ",") {
				switch part {
				case "pk":
					isPK = true
				case "computed":
					computed = true
				case "immutable":
					immutable = true
				}

				if strings.HasPrefix(part, "converter:") {
					converterName = part[len("converter:"):]
				}
			}
		}

		dbName := strcase.ToSnake(field.Name())
		if db, ok := tag.Lookup("db"); ok {
			dbName = db
		}

		nameParts := strings.Split(strcase.ToDelimited(field.Name(), '.'), ".")
		for i, part := range nameParts {
			nameParts[i] = cases.Title(language.English).String(part)
		}
		displayName := strings.Join(nameParts, " ")
		if dn, ok := tag.Lookup("display_name"); ok {
			displayName = dn
		}
		nameMap[dbName] = displayName

		f := fieldInfo{
			Name:          field.Name(),
			ConverterName: converterName,
			DatabaseName:  dbName,
		}

		if tag != "" {
			f.Tag = "`" + string(tag) + "`"
		}

		selectFields = append(selectFields, f)

		if computed && !isPK {
			continue
		}

		if isPK {
			if pkName != "" {
				return nil, fmt.Errorf("model %q has multiple primary keys", model)
			}
			pkName = field.Name()

			switch typ := field.Type().(type) {
			case *types.Basic:
				pkType.TypeName = typ.Name()
			case *types.Named:
				pkType.PackagePath = typ.Obj().Pkg().Path()
				pkType.TypeName = typ.Obj().Name()
			case *types.Alias:
				pkType.PackagePath = typ.Obj().Pkg().Path()
				pkType.TypeName = typ.Obj().Name()
			default:
				pkType.TypeName = types.TypeString(field.Type(), imports.resolvePackageAlias)
			}
		}

		if !computed {
			f.Type = types.TypeString(field.Type(), imports.resolvePackageAlias)
			createFields = append(createFields, f)

			if !immutable {
				updateF := f
				prefix := imports.addImport("github.com/partite-ai/optional", "optional")
				if !strings.HasPrefix(updateF.Type, prefix+".Optional[") {
					updateF.Type = prefix + ".Optional[" + updateF.Type + "]"
				}
				updateFields = append(updateFields, updateF)
			}
		}
	}

	return &modelInfo{
		Name:              model,
		PKName:            pkName,
		PKType:            pkType,
		CreateFields:      createFields,
		UpdateFields:      updateFields,
		SelectFields:      selectFields,
		FieldDisplayNames: nameMap,
	}, nil
}

func CreateModelFile(currentPackage *packages.Package, modelConf *ModelConfig, typ *types.Struct) error {
	imports := newImportSet(currentPackage.PkgPath, map[string]string{})
	modelInfo, err := LoadModelInfo(modelConf.Name, typ, imports)
	if err != nil {
		return err
	}
	if modelConf.DisplayName != "" {
		modelInfo.DisplayName = modelConf.DisplayName
	} else {
		words := strings.Split(strcase.ToDelimited(modelConf.Name, '.'), ".")
		for i, word := range words {
			words[i] = cases.Title(language.English).String(word)
		}
		modelInfo.DisplayName = strings.Join(words, " ")
	}

	if modelConf.PluralDisplayName != "" {
		modelInfo.PluralDisplayName = modelConf.PluralDisplayName
	} else {
		modelInfo.PluralDisplayName = pluralize.NewClient().Plural(modelConf.DisplayName)
	}

	if modelConf.TableName != "" {
		modelInfo.TableName = modelConf.TableName
	} else {
		modelInfo.TableName = strcase.ToSnake(modelConf.Name)
	}

	generateCreateParams := len(modelInfo.CreateFields) > 0 && !modelConf.SkipCreateParameters
	generateUpdateParams := len(modelInfo.UpdateFields) > 0 && !modelConf.SkipUpdateParameters

	if !generateCreateParams && !generateUpdateParams {
		return nil
	}

	var buf bytes.Buffer
	err = modelTemplate.Execute(&buf, map[string]interface{}{
		"PackageName":          currentPackage.Name,
		"Imports":              strings.Join(imports.importsList(), "\n"),
		"ModelName":            modelConf.Name,
		"GenerateCreateParams": generateCreateParams,
		"GenerateUpdateParams": generateUpdateParams,
		"CreateFields":         modelInfo.CreateFields,
		"UpdateFields":         modelInfo.UpdateFields,
		"TableName":            modelInfo.TableName,
		"DisplayName":          modelInfo.DisplayName,
		"PluralDisplayName":    modelInfo.PluralDisplayName,
		"FieldDisplayNames":    modelInfo.FieldDisplayNames,
	})

	if err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	modelFileName := strcase.ToSnake(modelConf.Name) + "_gen.go"
	f, err := os.Create(modelFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(formatted)
	return err
}

func addCRUDMethods(querySet *QuerySet, currentPackage *packages.Package, packageMap map[string]*packages.Package, mi *modelInfo) error {
	contextPackage := packageMap["context"]
	if querySet.HasQuery("Create"+mi.Name) && !querySet.HasMethod("Create"+mi.Name) {
		params := types.NewTuple(
			types.NewVar(token.NoPos, nil, "ctx", contextPackage.Types.Scope().Lookup("Context").Type()),
			types.NewVar(token.NoPos, nil, "params", types.NewPointer(currentPackage.Types.Scope().Lookup(mi.Name+"CreateParams").Type())),
		)

		results := types.NewTuple(
			types.NewVar(token.NoPos, nil, "", types.NewPointer(currentPackage.Types.Scope().Lookup(mi.Name).Type())),
			types.NewVar(token.NoPos, nil, "", types.Universe.Lookup("error").Type()),
		)
		sig := types.NewSignatureType(nil, nil, nil, params, results, false)
		querySet.AddMethod(types.NewFunc(token.NoPos, currentPackage.Types, "Create"+mi.Name, sig))
	}

	if mi.PKName != "" {
		var pkType types.Type
		if mi.PKType.PackagePath != "" {
			pkPkg := packageMap[mi.PKType.PackagePath]
			pkType = pkPkg.Types.Scope().Lookup(mi.PKType.TypeName).Type()
		} else {
			pkType = types.Universe.Lookup(mi.PKType.TypeName).Type()
		}
		if len(mi.UpdateFields) > 0 && querySet.HasQuery("Update"+mi.Name) && !querySet.HasMethod("Update"+mi.Name) {
			params := types.NewTuple(
				types.NewVar(token.NoPos, nil, "ctx", contextPackage.Types.Scope().Lookup("Context").Type()),
				types.NewVar(token.NoPos, nil, "id", pkType),
				types.NewVar(token.NoPos, nil, "params", types.NewPointer(currentPackage.Types.Scope().Lookup(mi.Name+"UpdateParams").Type())),
			)

			results := types.NewTuple(
				types.NewVar(token.NoPos, nil, "", types.NewPointer(currentPackage.Types.Scope().Lookup(mi.Name).Type())),
				types.NewVar(token.NoPos, nil, "", types.Universe.Lookup("error").Type()),
			)
			sig := types.NewSignatureType(nil, nil, nil, params, results, false)
			querySet.AddMethod(types.NewFunc(token.NoPos, currentPackage.Types, "Update"+mi.Name, sig))
		}

		if querySet.HasQuery("Get"+mi.Name) && !querySet.HasMethod("Get"+mi.Name) {
			params := types.NewTuple(
				types.NewVar(token.NoPos, nil, "ctx", contextPackage.Types.Scope().Lookup("Context").Type()),
				types.NewVar(token.NoPos, nil, "id", pkType),
			)

			results := types.NewTuple(
				types.NewVar(token.NoPos, nil, "", types.NewPointer(currentPackage.Types.Scope().Lookup(mi.Name).Type())),
				types.NewVar(token.NoPos, nil, "", types.Universe.Lookup("error").Type()),
			)
			sig := types.NewSignatureType(nil, nil, nil, params, results, false)
			querySet.AddMethod(types.NewFunc(token.NoPos, currentPackage.Types, "Get"+mi.Name, sig))
		}

		if querySet.HasQuery("Delete"+mi.Name) && !querySet.HasMethod("Delete"+mi.Name) {
			params := types.NewTuple(
				types.NewVar(token.NoPos, nil, "ctx", contextPackage.Types.Scope().Lookup("Context").Type()),
				types.NewVar(token.NoPos, nil, "id", pkType),
			)

			results := types.NewTuple(
				types.NewVar(token.NoPos, nil, "", types.Universe.Lookup("error").Type()),
			)
			sig := types.NewSignatureType(nil, nil, nil, params, results, false)
			querySet.AddMethod(types.NewFunc(token.NoPos, currentPackage.Types, "Delete"+mi.Name, sig))
		}
	}

	return nil
}
