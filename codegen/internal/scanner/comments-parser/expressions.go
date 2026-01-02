// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
)

// ParseTestValues parses a comma-separated list of Go expressions.
//
// Examples:
//   - "123, 456" → two integer literals
//   - "assertions.ErrTest, nil" → selector expression and nil
//   - "[]int{1,2,3}, []int{4,5,6}" → two composite literals
//   - `"a,b", "c,d"` → two string literals with commas inside
//
// Implementation: Wraps input in []any{...} and parses as composite literal,
// then extracts the elements. This leverages Go's parser to handle all edge cases
// (strings, runes, comments, nesting, etc.) correctly.
func ParseTestValues(input string) []model.TestValue {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	// Wrap in composite literal to let Go's parser handle comma splitting
	wrapped := "[]any{" + input + "}"

	// Parse the wrapped expression
	expr, err := parser.ParseExpr(wrapped)
	if err != nil {
		// If parsing fails, return a single TestValue with the error
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: fmt.Errorf("invalid Go expression list %q: %w", input, err),
		}}
	}

	// Extract elements from the composite literal
	compositeLit, ok := expr.(*ast.CompositeLit)
	if !ok {
		// Should never happen if parser succeeded
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: fmt.Errorf("internal error: expected composite literal, got %T", expr),
		}}
	}

	// Convert each element to TestValue
	result := make([]model.TestValue, 0, len(compositeLit.Elts))

	// We need to extract the original source text for each element
	// Since we don't have position info for the original input, we'll format the AST
	fset := token.NewFileSet()
	for _, elt := range compositeLit.Elts {
		// Format the expression back to source code
		var buf strings.Builder
		if err := format.Node(&buf, fset, elt); err != nil {
			result = append(result, model.TestValue{
				Raw:   "<formatting error>",
				Expr:  elt,
				Error: fmt.Errorf("failed to format expression: %w", err),
			})
			continue
		}

		result = append(result, model.TestValue{
			Raw:   buf.String(),
			Expr:  elt,
			Error: nil,
		})
	}

	return result
}

// ParseExprWithFileSet is like ParseTestValues but uses a provided FileSet
// for better error reporting with file/line/column information.
//
// The filename parameter is used for error messages (e.g., "assertion.go:42").
func ParseExprWithFileSet(fset *token.FileSet, filename string, input string) ([]model.TestValue, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}

	// Wrap in composite literal to let Go's parser handle comma splitting
	wrapped := "[]any{" + input + "}"

	// Parse with FileSet for better error reporting
	expr, err := parser.ParseExprFrom(fset, filename, wrapped, 0)
	if err != nil {
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: fmt.Errorf("invalid Go expression list %q: %w", input, err),
		}}, err
	}

	// Extract elements from the composite literal
	compositeLit, ok := expr.(*ast.CompositeLit)
	if !ok {
		err := fmt.Errorf("internal error: expected composite literal, got %T", expr)
		return []model.TestValue{{
			Raw:   input,
			Expr:  nil,
			Error: err,
		}}, err
	}

	// Convert each element to TestValue
	result := make([]model.TestValue, 0, len(compositeLit.Elts))

	for _, elt := range compositeLit.Elts {
		// Format the expression back to source code
		var buf strings.Builder
		if err := format.Node(&buf, fset, elt); err != nil {
			result = append(result, model.TestValue{
				Raw:   "<formatting error>",
				Expr:  elt,
				Error: fmt.Errorf("failed to format expression: %w", err),
			})
			continue
		}

		result = append(result, model.TestValue{
			Raw:   buf.String(),
			Expr:  elt,
			Error: nil,
		})
	}

	return result, nil
}
