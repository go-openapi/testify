// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

func render(tpl *template.Template, target string, data any, o *imports.Options) error {
	var buffer bytes.Buffer

	if err := tpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("error executing template %q: %w", tpl.Name(), err)
	}

	formatted, err := imports.Process(target, buffer.Bytes(), o)
	if err != nil {
		_ = os.WriteFile(target, buffer.Bytes(), filePermissions)
		return fmt.Errorf("error formatting go code: %w:%w", err, fmt.Errorf("details available at: %v", target))
	}

	// A rendered file with no declaration beyond imports carries no value: this happens
	// for build-variant partitions that have no function in a given category (e.g. a
	// go1.26 variant has no helpers). Skip writing it, and remove any stale leftover so
	// orphaned variant files don't accumulate.
	if !hasDeclarations(formatted) {
		_ = os.Remove(target)

		return nil
	}

	return os.WriteFile(target, formatted, filePermissions)
}

// hasDeclarations reports whether the Go source carries at least one top-level
// declaration other than imports.
func hasDeclarations(src []byte) bool {
	file, err := parser.ParseFile(token.NewFileSet(), "", src, parser.SkipObjectResolution)
	if err != nil {
		return true // be conservative: if we can't parse, let the normal write path keep it
	}

	for _, decl := range file.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.IMPORT {
			continue
		}

		return true
	}

	return false
}

func renderMD(tpl *template.Template, target string, data any) error {
	var buffer bytes.Buffer

	if err := tpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("error executing template %q: %w", tpl.Name(), err)
	}

	return os.WriteFile(target, buffer.Bytes(), filePermissions)
}
