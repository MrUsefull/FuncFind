// Package funcfind provides utilities for discovering Go functions by their return types.
// It uses static analysis to find functions that return specific types without requiring
// code compilation or execution.
//
// The primary use case is for static analysis tools that need to discover functions
// with particular return type signatures across Go packages.
//
// Example usage:
//
//	functions, err := funcfind.Returning("fmt", "error")
//	if err != nil {
//		panic(err)
//	}
//
//	for fn := range functions {
//		fmt.Printf("Function %s returns error\n", fn.Name())
//	}
package funcfind

import (
	"fmt"
	"go/types"
	"iter"

	"golang.org/x/tools/go/packages"
)

// Returning finds all functions in the specified Go package that return the provided
// returnTypes.
//
// The pkgPath parameter should be a valid Go package path (e.g., "fmt", "encoding/json",
// "github.com/user/repo/pkg"). The returnTypes parameter should be the string representation
// of the desired return types (e.g., "error", "string", "github.com/user/repo.CustomType").
//
// Returns an iterator over all matching functions.
// Should package loading fail, returns a non-nil error.
//
// Example:
//
//	// Find all functions in fmt package that return error
//	functions, err := Returning("fmt", "error")
//	if err != nil {
//		return err
//	}
//
//	for fn := range functions {
//		fmt.Printf("Found: %s\n", fn.Name())
//	}
func Returning(pkgPath string, returnTypes ...string) (iter.Seq[*types.Func], error) {
	pkgs, err := packages.Load(
		&packages.Config{
			Mode: packages.NeedTypes | packages.NeedTypesInfo,
		},
		pkgPath,
	)
	if err != nil {
		return nil, fmt.Errorf("load package %s: %w", pkgPath, err)
	}

	return scanPkgForFuncs(pkgs, returnTypes), nil
}

// scanPkgForFuncs discovers all functions in the loaded pkgs package that return exactly returnTypes.
func scanPkgForFuncs(pkgs []*packages.Package, returnTypes []string) iter.Seq[*types.Func] {
	return func(yield func(*types.Func) bool) {
		for _, pkg := range pkgs {
			scope := pkg.Types.Scope()
			for _, name := range scope.Names() {
				obj := scope.Lookup(name)
				if fn, ok := obj.(*types.Func); ok {
					if shouldYield(fn, returnTypes) && !yield(fn) {
						return
					}
				}
			}
		}
	}
}

func shouldYield(fn *types.Func, returnTypes []string) bool {
	if !fn.Exported() {
		return false
	}

	results := fn.Signature().Results()
	if results.Len() != len(returnTypes) {
		return false
	}

	for i := range results.Len() {
		rType := results.At(i).Type().String()
		if rType != returnTypes[i] {
			return false
		}
	}

	return true
}
