package main

import (
	"fmt"
	"go/types"
)

type importSet struct {
	currentPackage string
	imports        map[string]string
	defaultImports map[string]bool
}

func newImportSet(currentPackage string, initImports map[string]string) *importSet {
	is := &importSet{
		currentPackage: currentPackage,
		imports:        make(map[string]string),
		defaultImports: make(map[string]bool),
	}

	for k, v := range initImports {
		is.imports[k] = v
		is.defaultImports[k] = true
	}

	return is
}

func (i *importSet) importsList() []string {
	var result []string
	for pkgPath, alias := range i.imports {
		if i.defaultImports[pkgPath] {
			result = append(result, fmt.Sprintf("%q", pkgPath))
		} else {
			result = append(result, fmt.Sprintf("%s %q", alias, pkgPath))
		}
	}
	return result
}

func (i *importSet) addImport(importPath, defaultAlias string) string {
	if name, ok := i.imports[importPath]; ok {
		return name
	}

	if importPath == i.currentPackage {
		return ""
	}

	name := defaultAlias
	attempt := 0
	for {
		conflict := false
		for _, alias := range i.imports {
			if alias == name {
				conflict = true
				break
			}
		}

		if !conflict {
			break
		}

		name = fmt.Sprintf("%s%d", defaultAlias, attempt)
		attempt++
	}

	if attempt == 0 {
		i.defaultImports[importPath] = true
	}

	i.imports[importPath] = name
	return name
}
func (i *importSet) resolvePackageAlias(importPkg *types.Package) string {
	return i.addImport(importPkg.Path(), importPkg.Name())
}
