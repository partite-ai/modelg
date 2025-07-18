package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"go/types"
	"io"
	"maps"
	"os"
	"reflect"
	"slices"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/partite-ai/modelg/cmd/modelg/parser"
	"golang.org/x/tools/go/packages"
)

type queryType string

const (
	queryTypeGet        queryType = "get"
	queryTypeList       queryType = "list"
	queryTypeExec       queryType = "exec"
	queryTypeExecResult queryType = "execresult"
)

type scanMode string

const (
	scanModeNames     scanMode = "names"
	scanModePositions scanMode = "positions"
)

//go:embed queries_gen.tmpl
var queriesTemplateText string
var queriesTemplate = template.Must(template.New("queries").Parse(queriesTemplateText))

type QuerySet struct {
	queries    map[string]parser.QueryNode
	methods    map[string]*types.Func
	interfaces []types.Type
}

func LoadQuerySet(queryFile string) (*QuerySet, error) {
	f, err := os.Open(queryFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &QuerySet{}, nil
		}
		return nil, err
	}
	defer f.Close()

	qs, err := LoadQueries(f)
	if err != nil {
		return nil, fmt.Errorf("failed to load queries from %q: %w", queryFile, err)
	}
	return qs, nil
}

func LoadQueries(f io.Reader) (*QuerySet, error) {
	queries, err := parser.ParseQueries(f)
	if err != nil {
		return nil, err
	}
	qs := &QuerySet{
		queries: queries,
		methods: make(map[string]*types.Func),
	}
	return qs, nil
}

func (qs *QuerySet) HasQuery(name string) bool {
	_, ok := qs.queries[name]
	return ok
}

func (qs *QuerySet) AddMethod(fn *types.Func) {
	qs.methods[fn.Name()] = fn
}

func (qs *QuerySet) AddInterface(iface types.Type) error {
	ifaceType, ok := iface.Underlying().(*types.Interface)
	if !ok {
		return fmt.Errorf("type %q is not an interface", iface)
	}

	for i := 0; i < ifaceType.NumMethods(); i++ {
		fn := ifaceType.Method(i)
		qs.AddMethod(fn)
	}
	qs.interfaces = append(qs.interfaces, iface)
	return nil
}

func (qs *QuerySet) HasMethod(name string) bool {
	_, ok := qs.methods[name]
	return ok
}

func (qs *QuerySet) CreateQueriesFile(pkg *packages.Package, model string, converters []converterInfo) error {
	for name := range qs.queries {
		_, hasMethod := qs.methods[name]
		if !hasMethod {
			return fmt.Errorf("query %q has no method", name)
		}
	}

	for name := range qs.methods {
		_, hasQuery := qs.queries[name]
		if !hasQuery {
			return fmt.Errorf("method %q has no query", name)
		}
	}

	imports := newImportSet(pkg.PkgPath, map[string]string{
		"fmt":                          "fmt",
		"context":                      "context",
		"database/sql":                 "sql",
		"strings":                      "strings",
		"github.com/partite-ai/modelg": "modelg",
	})

	var methodSignatures []string
	methodNames := slices.Collect(maps.Keys(qs.methods))
	slices.Sort(methodNames)
	for _, name := range methodNames {
		fn := qs.methods[name]
		sig := fn.Signature()
		args := make([]string, sig.Params().Len())
		for i := 0; i < sig.Params().Len(); i++ {
			param := sig.Params().At(i)
			args[i] = param.Name() + " " + types.TypeString(param.Type(), imports.resolvePackageAlias)
		}

		returns := make([]string, sig.Results().Len())
		for i := 0; i < sig.Results().Len(); i++ {
			result := sig.Results().At(i)
			returns[i] = types.TypeString(result.Type(), imports.resolvePackageAlias)
		}

		methodSignatures = append(methodSignatures, fmt.Sprintf("%s(%s) (%s)", name, strings.Join(args, ", "), strings.Join(returns, ", ")))
	}

	var queryInfos []*queryInfo
	queryNames := slices.Collect(maps.Keys(qs.queries))
	slices.Sort(queryNames)
	for _, name := range queryNames {
		query := qs.queries[name]
		fn := qs.methods[name]
		sig := fn.Signature()
		args := make([]string, sig.Params().Len())
		for i := 0; i < sig.Params().Len(); i++ {
			param := sig.Params().At(i)
			args[i] = param.Name() + " " + types.TypeString(param.Type(), imports.resolvePackageAlias)
		}

		var rowType rowTypeInfo
		var queryType queryType
		returns := make([]string, sig.Results().Len())
		switch len(returns) {
		case 1:
			// Single return value, should be an error. Query type will be update.
			queryType = queryTypeExec
			returnType := types.TypeString(sig.Results().At(0).Type(), imports.resolvePackageAlias)
			if returnType != "error" {
				return fmt.Errorf("method %q with a single return value should return an error", name)
			}
		case 2:
			// Two return values, first should be the result, second should be an error.
			secondReturnType := types.TypeString(sig.Results().At(1).Type(), imports.resolvePackageAlias)
			if secondReturnType != "error" {
				return fmt.Errorf("method %q with two return values should return an error as the second value", name)
			}

			firstReturnType := sig.Results().At(0).Type()
			var err error
			rowType, queryType, err = processReturnType(name, firstReturnType, imports.resolvePackageAlias, converters)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("method %q should have one or two return values", name)
		}
		for i := 0; i < sig.Results().Len(); i++ {
			result := sig.Results().At(i)
			returns[i] = types.TypeString(result.Type(), imports.resolvePackageAlias)
		}

		env := &methodArgsQueryEnv{
			pkg:        pkg.Types,
			args:       sig.Params(),
			converters: converters,
			imports:    imports,
		}

		lines, err := query.Apply(env)
		if err != nil {
			return err
		}

		queryInfos = append(queryInfos, &queryInfo{
			Name:      name,
			Args:      strings.Join(args, ", "),
			Returns:   strings.Join(returns, ", "),
			QueryType: queryType,
			QueryInit: strings.Join(lines, "\n"),
			RowType:   rowType,
		})
	}

	var interfaces []string
	for _, iface := range qs.interfaces {
		interfaces = append(interfaces, types.TypeString(iface, imports.resolvePackageAlias))
	}

	var buf bytes.Buffer
	err := queriesTemplate.Execute(&buf, map[string]interface{}{
		"PackageName":      pkg.Name,
		"QueryInterfaces":  interfaces,
		"Imports":          strings.Join(imports.importsList(), "\n"),
		"ModelName":        model,
		"MethodSignatures": methodSignatures,
		"Queries":          queryInfos,
	})
	if err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return fmt.Errorf("cannot format source: %w", err)
	}

	f, err := os.Create(strcase.ToSnake(model) + "_queries_gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(formatted)
	return err
}

type queryInfo struct {
	Name        string
	Args        string
	Returns     string
	QueryType   queryType
	QueryInit   string
	ScanTargets []scanTargetInfo
	RowType     rowTypeInfo
}

type rowTypeInfo struct {
	CollectionType string
	ItemType       string
	InitExpr       string
	ReturnExpr     string
	ScanMode       scanMode
	ScanTargets    []scanTargetInfo
}

type scanTargetInfo struct {
	ColumnName  string
	PrepareExpr string
	TargetExpr  string
}

func processReturnType(name string, returnType types.Type, qualifier types.Qualifier, converters []converterInfo) (rowType rowTypeInfo, queryType queryType, err error) {
	for _, converter := range converters {
		if converter.CanScan("", returnType) {
			invExpr, err := converter.ScannerInvocation(qualifier, converters, returnType, "&result__")
			if err != nil {
				return rowTypeInfo{}, "", err
			}
			return rowTypeInfo{
				ItemType:   types.TypeString(returnType, qualifier),
				InitExpr:   fmt.Sprintf("var result__ %s", types.TypeString(returnType, qualifier)),
				ReturnExpr: "result__",
				ScanMode:   scanModePositions,
				ScanTargets: []scanTargetInfo{
					{
						PrepareExpr: fmt.Sprintf("cvt__ := %s", invExpr),
						TargetExpr:  "cvt__",
					},
				},
			}, queryTypeGet, nil

		}
	}

	switch typ := returnType.Underlying().(type) {
	case *types.Named:
		return processReturnType(name, typ.Underlying(), qualifier, converters)
	case *types.Struct:
		typName := types.TypeString(returnType, qualifier)
		rti := rowTypeInfo{
			ItemType:   typName,
			InitExpr:   fmt.Sprintf("var result__ %s", typName),
			ReturnExpr: "result__",
			ScanMode:   scanModeNames,
		}

		err = createStructScanTargets(qualifier, typ, "&result__", &rti, converters, 0)
		return rti, queryTypeGet, err
	case *types.Pointer:
		rowType, queryType, err = processReturnType(name, typ.Elem(), qualifier, converters)
		if err != nil {
			return
		}
		rowType.ItemType = "*" + rowType.ItemType
		rowType.ReturnExpr = "&" + rowType.ReturnExpr
		return
	case *types.Slice:
		// exactly []byte is handled like a value type, not a list.
		// BUT: a typedef that has an underlying type of [] byte is handled normally
		if rts, ok := returnType.(*types.Slice); ok {
			if bt, ok := rts.Elem().(*types.Basic); ok && bt.Kind() == types.Byte {
				return rowTypeInfo{
					ItemType:   "[]byte",
					InitExpr:   "var result__ []byte",
					ReturnExpr: "result__",
					ScanMode:   scanModePositions,
					ScanTargets: []scanTargetInfo{
						{
							TargetExpr: "&result__",
						},
					},
				}, queryTypeGet, nil
			}
		}
		rowType, queryType, err = processReturnType(name, typ.Elem(), qualifier, converters)
		if err != nil {
			return
		}
		rowType.CollectionType = "[]" + rowType.ItemType
		return rowType, queryTypeList, nil
	case *types.Basic:
		return rowTypeInfo{
			ItemType:   types.TypeString(returnType, qualifier),
			InitExpr:   fmt.Sprintf("var result__ %s", types.TypeString(returnType, qualifier)),
			ReturnExpr: "result__",
			ScanMode:   scanModePositions,
			ScanTargets: []scanTargetInfo{
				{
					TargetExpr: "&result__",
				},
			},
		}, queryTypeGet, nil
	case *types.Interface:
		if typ.String() == "database/sql.Result" {
			return rowTypeInfo{}, queryTypeExecResult, nil
		}
		return rowTypeInfo{}, "", fmt.Errorf("method %q with two return values should return a struct, pointer, slice or basic type as the first value", name)
	default:
		return rowTypeInfo{}, "", fmt.Errorf("method %q with two return values should return a struct, pointer, slice or basic type as the first value", name)
	}
}

type methodArgsQueryEnv struct {
	pkg        *types.Package
	args       *types.Tuple
	converters []converterInfo
	imports    *importSet
}

func (env *methodArgsQueryEnv) ParamCodeExpr(parts []string) (string, error) {
	if len(parts) > 2 {
		return "", fmt.Errorf("invalid name %q", strings.Join(parts, "."))
	}

	paramExpr := make([]string, 0, len(parts))
	var paramType types.Type
	argName := parts[0]
	normalizedArgName := strcase.ToLowerCamel(argName)
	for i := 0; i < env.args.Len(); i++ {
		arg := env.args.At(i)
		normalizedCandidateArgName := strcase.ToLowerCamel(arg.Name())
		if normalizedArgName == normalizedCandidateArgName {
			paramExpr = append(paramExpr, arg.Name())
			paramType = arg.Type()
			break
		}
	}

	converterName := ""
	if len(paramExpr) == 0 {
		for i := 0; i < env.args.Len(); i++ {
			arg := env.args.At(i)
			typ := arg.Type().Underlying()
			if ptr, ok := typ.(*types.Pointer); ok {
				typ = ptr.Elem().Underlying()
			}

			if structType, ok := typ.(*types.Struct); ok {
				for i := 0; i < structType.NumFields(); i++ {
					field := structType.Field(i)
					normalizedFieldName := strcase.ToLowerCamel(field.Name())
					if normalizedArgName == normalizedFieldName {
						paramExpr = append(paramExpr, arg.Name()+"."+field.Name())
						paramType = field.Type()
						tag := reflect.StructTag(structType.Tag(i))
						if modelgTag, ok := tag.Lookup("modelg"); ok {
							parts := strings.Split(modelgTag, ",")
							for _, part := range parts {
								if strings.HasPrefix(part, "converter:") {
									converterName = strings.TrimPrefix(part, "converter:")
								}
							}
						}
						break
					}
				}
			}
		}
	}

	if len(paramExpr) == 0 {
		return "", fmt.Errorf("param %q not found", strings.Join(parts, "."))
	}

	if len(parts) == 2 {
		// resolve the package
		fieldOrMethod, _, _ := types.LookupFieldOrMethod(paramType, true, env.pkg, parts[1])
		if fieldOrMethod == nil {
			return "", fmt.Errorf("field or method %q not found in %q", parts[1], argName)
		}

		switch fm := fieldOrMethod.(type) {
		case *types.Var:
			paramExpr = append(paramExpr, parts[1])
		case *types.Func:
			if fm.Signature().Params().Len() != 0 {
				return "", fmt.Errorf("method %q should not have any parameters", parts[1])
			}
			paramExpr = append(paramExpr, parts[1]+"()")
		}
	} else {
		for _, converter := range env.converters {
			if converter.CanValue(converterName, paramType) {
				return converter.ValuerInvocation(env.imports.resolvePackageAlias, env.converters, paramType, strings.Join(paramExpr, "."))
			}
		}
	}

	return strings.Join(paramExpr, "."), nil
}

func createStructScanTargets(qualifier types.Qualifier, typ *types.Struct, structExpr string, rti *rowTypeInfo, converters []converterInfo, fieldOffset int) error {
	for i := 0; i < typ.NumFields(); i++ {
		field := typ.Field(i)
		if !field.Exported() {
			continue
		}
		tag := reflect.StructTag(typ.Tag(i))
		fieldName := strcase.ToSnake(field.Name())

		if db, ok := tag.Lookup("db"); ok {
			fieldName = db
		}

		var converterName string
		var flatten bool
		if modelgTag, ok := tag.Lookup("modelg"); ok {
			parts := strings.Split(modelgTag, ",")
			for _, part := range parts {
				if strings.HasPrefix(part, "converter:") {
					converterName = strings.TrimPrefix(part, "converter:")
					break
				}
				if part == "flatten" {
					flatten = true
				}
			}
		}

		if flatten && converterName != "" {
			return fmt.Errorf("field %q cannot have both flatten and converter tags", field.Name())
		}

		if flatten {
			if sft, ok := field.Type().Underlying().(*types.Struct); ok {
				err := createStructScanTargets(qualifier, sft, fmt.Sprintf("%s.%s", structExpr, field.Name()), rti, converters, typ.NumFields()+fieldOffset)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("field %q is not a struct", field.Name())
			}
		}

		var sti scanTargetInfo

		hasConverter := false
		for _, converter := range converters {
			if !converter.CanScan(converterName, field.Type()) {
				continue
			}
			hasConverter = true

			invExpr, err := converter.ScannerInvocation(qualifier, converters, field.Type(), fmt.Sprintf("%s.%s", structExpr, field.Name()))
			if err != nil {
				return err
			}

			sti = scanTargetInfo{
				ColumnName:  fieldName,
				PrepareExpr: fmt.Sprintf("cvt__%d := %s", i+fieldOffset, invExpr),
				TargetExpr:  fmt.Sprintf("cvt__%d", i+fieldOffset),
			}
			break

		}

		if !hasConverter {
			sti = scanTargetInfo{
				ColumnName: fieldName,
				TargetExpr: fmt.Sprintf("%s.%s", structExpr, field.Name()),
			}
		}

		rti.ScanTargets = append(rti.ScanTargets, sti)
	}

	return nil
}
