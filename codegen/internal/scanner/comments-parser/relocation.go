// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"go/ast"
	"go/format"
	"go/token"
	"strings"
	"unicode"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

const (
	// Exception: PanicTestFunc should always use "assertions" package.
	panicTestFunc = "PanicTestFunc"
	assertions    = "assertions"
)

// RelocateTestValue relocates a test value expression from one package to another.
//
// It walks the AST and:
//  1. Changes qualified identifiers: fromPkg.X → toPkg.X (except exceptions)
//  2. Qualifies unqualified exported identifiers: X → toPkg.X
//
// Exception: PanicTestFunc always stays with assertions package.
func RelocateTestValue(tv model.TestValue, fromPkg, toPkg string) model.TestValue {
	if tv.Expr == nil || fromPkg == "" || toPkg == "" {
		return tv
	}

	// Handle root-level unqualified identifier specially
	expr := tv.Expr
	if ident, ok := expr.(*ast.Ident); ok && shouldQualify(ident) {
		// Special case for PanicTestFunc: always use assertions
		pkg := toPkg
		if ident.Name == panicTestFunc {
			pkg = assertions
		}
		expr = &ast.SelectorExpr{
			X:   &ast.Ident{Name: pkg},
			Sel: ident,
		}
	}

	// Walk the AST and modify selectors in place
	ast.Inspect(expr, func(n ast.Node) bool {
		if node, ok := n.(*ast.SelectorExpr); ok {
			// Handle qualified identifiers: pkg.Identifier
			relocateSelectorExpr(node, fromPkg, toPkg)
		}

		return true
	})

	// Second pass: qualify unqualified exported identifiers in sub-expressions
	qualifyUnqualifiedIdents(expr, toPkg)

	// Re-render the modified AST to string
	fset := token.NewFileSet()
	var buf strings.Builder
	if err := format.Node(&buf, fset, expr); err != nil {
		return model.TestValue{
			Raw:   tv.Raw,
			Expr:  tv.Expr,
			Error: err,
		}
	}

	return model.TestValue{
		Raw:   buf.String(),
		Expr:  expr,
		Error: nil,
	}
}

// relocateSelectorExpr handles selector expressions:
//  1. Package selectors: assertions.CollectT → assert.CollectT
//  2. Method calls on unqualified types: ErrTest.Error() → assert.ErrTest.Error()
func relocateSelectorExpr(sel *ast.SelectorExpr, fromPkg, toPkg string) {
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return
	}

	// Exception: PanicTestFunc always uses assertions package
	if sel.Sel.Name == panicTestFunc {
		ident.Name = assertions
		return
	}

	// Case 1: Package selector (assertions.CollectT → assert.CollectT)
	if ident.Name == fromPkg {
		ident.Name = toPkg
		return
	}

	// Case 2: Unqualified identifier in selector (ErrTest.Error() → assert.ErrTest.Error())
	// Need to wrap the identifier in a SelectorExpr
	if shouldQualify(ident) {
		sel.X = qualifyIdent(ident, toPkg)
	}
}

// qualifyUnqualifiedIdents adds package qualifiers to unqualified exported identifiers.
//
// This recursively processes all type expressions in the AST, replacing unqualified
// identifiers with qualified SelectorExpr nodes.
func qualifyUnqualifiedIdents(expr ast.Expr, pkg string) {
	ast.Inspect(expr, func(n ast.Node) bool {
		switch parent := n.(type) {
		case *ast.CallExpr:
			// Qualify function name
			if ident, ok := parent.Fun.(*ast.Ident); ok && shouldQualify(ident) {
				parent.Fun = qualifyIdent(ident, pkg)
			}
			// Qualify arguments
			for i, arg := range parent.Args {
				if ident, ok := arg.(*ast.Ident); ok && shouldQualify(ident) {
					parent.Args[i] = qualifyIdent(ident, pkg)
				}
			}

		case *ast.CompositeLit:
			// Qualify type in composite literal
			parent.Type = qualifyTypeExpr(parent.Type, pkg)

		case *ast.TypeAssertExpr:
			// Qualify type in type assertion
			parent.Type = qualifyTypeExpr(parent.Type, pkg)

		case *ast.StarExpr:
			// Qualify type inside pointer: *T → *pkg.T
			parent.X = qualifyTypeExpr(parent.X, pkg)

		case *ast.ArrayType:
			// Qualify element type in array/slice: []T → []pkg.T
			parent.Elt = qualifyTypeExpr(parent.Elt, pkg)

		case *ast.MapType:
			// Qualify key and value types in map: map[K]V → map[pkg.K]pkg.V
			parent.Key = qualifyTypeExpr(parent.Key, pkg)
			parent.Value = qualifyTypeExpr(parent.Value, pkg)

		case *ast.ChanType:
			// Qualify type in channel: chan T → chan pkg.T
			parent.Value = qualifyTypeExpr(parent.Value, pkg)

		case *ast.Field:
			// Qualify type in function parameters, struct fields, etc.
			parent.Type = qualifyTypeExpr(parent.Type, pkg)
		}

		return true
	})
}

// qualifyTypeExpr recursively qualifies unqualified identifiers in a type expression.
func qualifyTypeExpr(typ ast.Expr, pkg string) ast.Expr {
	if typ == nil {
		return nil
	}

	switch t := typ.(type) {
	case *ast.Ident:
		if shouldQualify(t) {
			return qualifyIdent(t, pkg)
		}
		return t

	case *ast.StarExpr:
		t.X = qualifyTypeExpr(t.X, pkg)
		return t

	case *ast.ArrayType:
		t.Elt = qualifyTypeExpr(t.Elt, pkg)
		return t

	case *ast.MapType:
		t.Key = qualifyTypeExpr(t.Key, pkg)
		t.Value = qualifyTypeExpr(t.Value, pkg)
		return t

	case *ast.ChanType:
		t.Value = qualifyTypeExpr(t.Value, pkg)
		return t

	default:
		return typ
	}
}

// qualifyIdent wraps an identifier in a SelectorExpr with the appropriate package.
func qualifyIdent(ident *ast.Ident, pkg string) *ast.SelectorExpr {
	// Special case for PanicTestFunc: always use assertions
	targetPkg := pkg
	if ident.Name == panicTestFunc {
		targetPkg = assertions
	}

	return &ast.SelectorExpr{
		X:   &ast.Ident{Name: targetPkg},
		Sel: ident,
	}
}

// shouldQualify returns true if an identifier should be qualified with a package name.
//
// Qualifies exported identifiers (uppercase start) except:
//   - Language keywords and built-ins
func shouldQualify(ident *ast.Ident) bool {
	name := ident.Name
	if name == "" {
		return false
	}

	// Don't qualify if not exported (lowercase start)
	if !unicode.IsUpper(rune(name[0])) {
		return false
	}

	// Don't qualify language built-ins
	switch name {
	case "bool", "byte", "complex64", "complex128",
		"error", "float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"rune", "string",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		return false
	}

	return true
}
