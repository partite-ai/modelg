package main

import (
	"bytes"
	"fmt"
	"go/types"
	"strings"

	"github.com/partite-ai/modelg/cmd/modelg/typeutil"
	"golang.org/x/tools/go/packages"
)

type converterInfo struct {
	Name        string
	ScannerFunc types.Object
	ValuerFunc  types.Object
}

func (ci *converterInfo) loadScannerFunction(pkgMap map[string]*packages.Package, name string) error {
	lastDot := strings.LastIndex(name, ".")
	if lastDot == -1 || lastDot == len(name)-1 {
		return fmt.Errorf("invalid scanner function %q", name)
	}
	pkgName := name[:lastDot]
	funcName := name[lastDot+1:]

	pkg := pkgMap[pkgName]
	obj := pkg.Types.Scope().Lookup(funcName)

	scannerFunc, ok := obj.Type().Underlying().(*types.Signature)
	if !ok {
		return fmt.Errorf("scanner function %q is not a function", name)
	}

	if scannerFunc.Params().Len() < 1 {
		return fmt.Errorf("scanner function %q must have at least one parameter", name)
	}

	param1 := scannerFunc.Params().At(0)
	_, ok = param1.Type().Underlying().(*types.Pointer)
	if !ok {
		return fmt.Errorf("scanner function %q first parameter must be a pointer", name)
	}

	for i := 1; i < scannerFunc.Params().Len(); i++ {
		param := scannerFunc.Params().At(i)
		paramSig, ok := param.Type().Underlying().(*types.Signature)
		if !ok {
			return fmt.Errorf("scanner function %q, parameter %s is expected to be of the form func(*T) sql.Scanner", name, param.Name())
		}

		if paramSig.Params().Len() != 1 {
			return fmt.Errorf("scanner function %q, parameter %s is expected to be of the form func(*T) sql.Scanner", name, param.Name())
		}

		_, ok = paramSig.Params().At(0).Type().Underlying().(*types.Pointer)
		if !ok {
			return fmt.Errorf("scanner function %q, parameter %s is expected to be of the form func(*T) sql.Scanner", name, param.Name())
		}
	}

	if scannerFunc.Results().Len() != 1 {
		return fmt.Errorf("scanner function %q must have exactly one result", name)
	}

	ci.ScannerFunc = obj
	return nil
}

func (ci *converterInfo) loadValuerFunction(pkgMap map[string]*packages.Package, name string) error {
	lastDot := strings.LastIndex(name, ".")
	if lastDot == -1 || lastDot == len(name)-1 {
		return fmt.Errorf("invalid valuer function %q", name)
	}
	pkgName := name[:lastDot]
	funcName := name[lastDot+1:]

	pkg := pkgMap[pkgName]
	obj := pkg.Types.Scope().Lookup(funcName)

	valuerFunc, ok := obj.Type().Underlying().(*types.Signature)
	if !ok {
		return fmt.Errorf("valuer function %q is not a function", name)
	}

	if valuerFunc.Params().Len() < 1 {
		return fmt.Errorf("valuer function %q must have at least one parameter", name)
	}

	for i := 1; i < valuerFunc.Params().Len(); i++ {
		param := valuerFunc.Params().At(i)
		paramSig, ok := param.Type().Underlying().(*types.Signature)
		if !ok {
			return fmt.Errorf("valuer function %q, parameter %s is expected to be of the form func(T) driver.Valuer", name, param.Name())
		}

		if paramSig.Params().Len() != 1 {
			return fmt.Errorf("valuer function %q, parameter %s is expected to be of the form func(T) driver.Valuer", name, param.Name())
		}
	}

	if valuerFunc.Results().Len() != 1 {
		return fmt.Errorf("valuer function %q must have exactly one result", name)
	}

	ci.ValuerFunc = obj
	return nil
}

func (ci *converterInfo) CanScan(name string, typ types.Type) bool {
	if ci.ScannerFunc == nil {
		return false
	}

	if ci.Name != name {
		return false
	}

	scannerFunc, ok := ci.ScannerFunc.Type().Underlying().(*types.Signature)
	if !ok {
		return false
	}

	argTyp := types.NewPointer(typ)
	_, ok = typeutil.InferTypeFromFirstParam(scannerFunc, argTyp)
	return ok
}

func (ci *converterInfo) ScannerInvocation(qualifier types.Qualifier, converters []converterInfo, typ types.Type, expr string) (string, error) {
	scannerFunc, ok := ci.ScannerFunc.Type().Underlying().(*types.Signature)
	if !ok {
		return "", fmt.Errorf("scanner function %s is not a function", ci.ScannerFunc.Name())
	}

	qName := ci.ScannerFunc.Name()
	pkgName := qualifier(ci.ScannerFunc.Pkg())
	if pkgName != "" {
		qName = pkgName + "." + qName
	}
	instance, ok := typeutil.InferTypeFromFirstParam(scannerFunc, types.NewPointer(typ))
	if !ok {
		return "", fmt.Errorf("could not infer type for scanner function %s", ci.ScannerFunc.Name())
	}

	callExpr := qName + "(" + expr
	for i := 1; i < instance.Params().Len(); i++ {
		param := instance.Params().At(i)
		paramSig, ok := param.Type().Underlying().(*types.Signature)
		if !ok {
			return "", fmt.Errorf("scanner function %s has a non-signature parameter %s", ci.ScannerFunc.Name(), param.Name())
		}

		if paramSig.Params().Len() != 1 {
			return "", fmt.Errorf("scanner function %s parameter %s is expected to be func(*T) sql.Scanner", ci.ScannerFunc.Name(), param.Name())
		}

		paramPtr, ok := paramSig.Params().At(0).Type().Underlying().(*types.Pointer)
		if !ok {
			return "", fmt.Errorf("scanner function %s parameter %s is expected to be func(*T) sql.Scanner", ci.ScannerFunc.Name(), param.Name())
		}

		p0 := paramSig.Params().At(0)

		sigClone := types.NewSignatureType(
			paramSig.Recv(),
			vlistToSlice(paramSig.RecvTypeParams()),
			vlistToSlice(paramSig.TypeParams()),
			types.NewTuple(
				types.NewVar(p0.Pos(), p0.Pkg(), "dst__", p0.Type()),
			),
			paramSig.Results(),
			paramSig.Variadic(),
		)

		convTyp := paramPtr.Elem()
		found := false
		for _, converter := range converters {
			if converter.CanScan("", convTyp) {
				var buf bytes.Buffer
				types.WriteSignature(&buf, sigClone, qualifier)

				invokeConverter, err := converter.ScannerInvocation(qualifier, converters, convTyp, "dst__")
				if err != nil {
					return "", err
				}
				sigExpr := "func" + buf.String() + "{\n return " + invokeConverter + "\n}"
				callExpr += ", " + sigExpr
				found = true
				break
			}
		}

		if !found {
			var buf bytes.Buffer
			types.WriteSignature(&buf, sigClone, qualifier)

			sigExpr := "func" + buf.String() + "{\n return basicScanner{dest: dst__}\n}"
			callExpr += ", " + sigExpr
		}
	}
	callExpr += ")"
	return callExpr, nil
}

func (ci *converterInfo) CanValue(name string, typ types.Type) bool {
	if ci.ValuerFunc == nil {
		return false
	}

	if ci.Name != name {
		return false
	}

	valuerFunc, ok := ci.ValuerFunc.Type().Underlying().(*types.Signature)
	if !ok {
		return false
	}

	_, ok = typeutil.InferTypeFromFirstParam(valuerFunc, typ)
	return ok
}

func (ci *converterInfo) ValuerInvocation(qualifier types.Qualifier, converters []converterInfo, typ types.Type, expr string) (string, error) {
	valuerFunc, ok := ci.ValuerFunc.Type().Underlying().(*types.Signature)
	if !ok {
		return "", fmt.Errorf("valuer function %s is not a function", ci.ValuerFunc.Name())
	}

	qName := ci.ValuerFunc.Name()
	pkgName := qualifier(ci.ValuerFunc.Pkg())
	if pkgName != "" {
		qName = pkgName + "." + qName
	}
	instance, ok := typeutil.InferTypeFromFirstParam(valuerFunc, typ)
	if !ok {
		return "", fmt.Errorf("could not infer type for valuer function %s", ci.ValuerFunc.Name())
	}

	callExpr := qName + "(" + expr
	for i := 1; i < instance.Params().Len(); i++ {
		param := instance.Params().At(i)
		paramSig, ok := param.Type().Underlying().(*types.Signature)
		if !ok {
			return "", fmt.Errorf("valuer function %s has a non-signature parameter %s", ci.ValuerFunc.Name(), param.Name())
		}

		if paramSig.Params().Len() != 1 {
			return "", fmt.Errorf("valuer function %s parameter %s is expected to be func(T) driver.Valuer", ci.ValuerFunc.Name(), param.Name())
		}

		p0 := paramSig.Params().At(0)

		sigClone := types.NewSignatureType(
			paramSig.Recv(),
			vlistToSlice(paramSig.RecvTypeParams()),
			vlistToSlice(paramSig.TypeParams()),
			types.NewTuple(
				types.NewVar(p0.Pos(), p0.Pkg(), "src__", p0.Type()),
			),
			paramSig.Results(),
			paramSig.Variadic(),
		)

		paramTyp := paramSig.Params().At(0).Type()
		found := false
		for _, converter := range converters {
			if converter.CanValue("", paramTyp) {
				var buf bytes.Buffer
				types.WriteSignature(&buf, sigClone, qualifier)

				invokeConverter, err := converter.ValuerInvocation(qualifier, converters, paramTyp, "src__")
				if err != nil {
					return "", err
				}
				sigExpr := "func" + buf.String() + "{\n return " + invokeConverter + "\n}"
				callExpr += ", " + sigExpr
				found = true
				break
			}
		}

		if !found {
			var buf bytes.Buffer
			types.WriteSignature(&buf, sigClone, qualifier)

			sigExpr := "func" + buf.String() + "{\n return basicValuer{src: src__}\n}"
			callExpr += ", " + sigExpr
		}
	}
	callExpr += ")"
	return callExpr, nil
}

type vlist[T any] interface {
	At(i int) T
	Len() int
}

func vlistToSlice[T any](list vlist[T]) []T {
	slice := make([]T, list.Len())
	for i := 0; i < list.Len(); i++ {
		slice[i] = list.At(i)
	}
	return slice
}
