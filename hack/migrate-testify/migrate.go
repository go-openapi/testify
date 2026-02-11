// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// runMigration executes pass 1: stretchr/testify → go-openapi/testify/v2.
func runMigration(dir string, opts *options) error {
	fset := token.NewFileSet()
	rpt := &report{}

	// Walk all .go files in the directory tree.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			base := filepath.Base(path)
			if base == "vendor" && opts.skipVendor {
				return filepath.SkipDir
			}
			if base == ".git" || base == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		rpt.filesScanned++
		return migrateFile(fset, path, opts, rpt)
	})
	if err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	rpt.print(opts.verbose)
	rpt.printPass1Summary()

	if !opts.dryRun && !opts.skipGomod {
		if err := updateGoMod(dir, opts.version, false, opts.verbose); err != nil {
			return fmt.Errorf("updating go.mod: %w", err)
		}
	} else if opts.dryRun && !opts.skipGomod {
		if err := updateGoMod(dir, opts.version, true, opts.verbose); err != nil {
			return fmt.Errorf("updating go.mod: %w", err)
		}
	}

	return nil
}

// migrateFile processes a single Go file for pass 1 transformations.
func migrateFile(fset *token.FileSet, filename string, opts *options, rpt *report) error {
	src, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", filename, err)
	}

	if !fileImportsAny(f, "github.com/stretchr/testify") {
		return nil
	}

	changed := false

	// 1. Detect incompatible imports.
	for imp, msg := range incompatibleImports {
		if fileImportsPath(f, imp) {
			rpt.warn(filename, 0, msg)
		}
	}

	// 2. Build alias map before rewriting imports.
	aliases := buildAliasMap(f)

	// 2b. Count API usage for the summary report.
	countAPIUsage(f, aliases, rpt)

	// 3. Rewrite imports.
	for old, replacement := range importRewrites {
		if astutil.RewriteImport(fset, f, old, replacement) {
			changed = true
			rpt.totalChanges++
			if opts.verbose {
				rpt.info(filename, 0, fmt.Sprintf("rewrote import %q → %q", old, replacement))
			}
		}
	}

	// 4. Rename functions and replace PanicTestFunc.
	changes := renameFunctions(f, aliases, fset, rpt, filename, opts.verbose)
	if changes > 0 {
		changed = true
		rpt.totalChanges += changes
	}

	// 5. Replace PanicTestFunc type references.
	ptfChanges := replacePanicTestFunc(f, aliases, fset, rpt, filename, opts.verbose)
	if ptfChanges > 0 {
		changed = true
		rpt.totalChanges += ptfChanges
	}

	// 6. Detect YAML usage and inject enable import.
	if needsYAMLEnable(f, aliases) {
		if !fileImportsPath(f, goopenapiYAMLEnable) {
			astutil.AddNamedImport(fset, f, "_", goopenapiYAMLEnable)
			changed = true
			rpt.totalChanges++
			if opts.verbose {
				rpt.info(filename, 0, "injected enable/yaml import")
			}
		}
	}

	if !changed {
		return nil
	}

	rpt.filesChanged++

	if opts.dryRun {
		return showDiff(fset, f, filename)
	}

	return writeFile(fset, f, filename)
}

// buildAliasMap builds a map from import alias (or default package name) to import path
// for stretchr/testify packages.
func buildAliasMap(f *ast.File) map[string]string {
	aliases := make(map[string]string)
	for _, imp := range f.Imports {
		path := importPath(imp)
		if !strings.HasPrefix(path, "github.com/stretchr/testify") {
			continue
		}

		var localName string
		if imp.Name != nil {
			localName = imp.Name.Name
		} else {
			// Default: last path element.
			parts := strings.Split(path, "/")
			localName = parts[len(parts)-1]
		}
		if localName != "_" && localName != "." {
			aliases[localName] = path
		}
	}
	return aliases
}

// renameFunctions walks the AST and renames function calls that changed names
// between stretchr/testify and go-openapi/testify/v2.
func renameFunctions(f *ast.File, aliases map[string]string, fset *token.FileSet, rpt *report, filename string, verbose bool) int {
	changes := 0

	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		// Handle pkg.Func() calls.
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		funcName := sel.Sel.Name
		newName, exists := migrationRenames[funcName]
		if !exists {
			return true
		}

		// Verify this is a call on a testify package.
		if !isTestifySelector(sel, aliases) {
			return true
		}

		sel.Sel.Name = newName
		changes++
		pos := fset.Position(call.Pos())
		if verbose {
			rpt.info(filename, pos.Line, fmt.Sprintf("renamed %s → %s", funcName, newName))
		}

		return true
	})

	return changes
}

// isTestifySelector checks if a selector expression refers to a stretchr/testify package.
func isTestifySelector(sel *ast.SelectorExpr, aliases map[string]string) bool {
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	_, exists := aliases[ident.Name]
	return exists
}

// replacePanicTestFunc replaces PanicTestFunc type references with func().
func replacePanicTestFunc(f *ast.File, aliases map[string]string, fset *token.FileSet, rpt *report, filename string, verbose bool) int {
	changes := 0

	astutil.Apply(f, func(c *astutil.Cursor) bool {
		sel, ok := c.Node().(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if sel.Sel.Name != "PanicTestFunc" {
			return true
		}

		if !isTestifySelector(sel, aliases) {
			return true
		}

		// Replace with func() — an *ast.FuncType with no params and no results.
		c.Replace(&ast.FuncType{
			Params: &ast.FieldList{},
		})
		changes++
		pos := fset.Position(sel.Pos())
		if verbose {
			rpt.info(filename, pos.Line, "replaced PanicTestFunc with func()")
		}

		return true
	}, nil)

	return changes
}

// countAPIUsage walks the AST and counts every testify assertion call for reporting.
func countAPIUsage(f *ast.File, aliases map[string]string, rpt *report) {
	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if !isTestifySelector(sel, aliases) {
			return true
		}

		ident, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		rpt.trackAPIUsage(ident.Name + "." + sel.Sel.Name)

		return true
	})
}

// needsYAMLEnable checks if any YAML assertion functions are called in the file.
func needsYAMLEnable(f *ast.File, aliases map[string]string) bool {
	found := false

	ast.Inspect(f, func(n ast.Node) bool {
		if found {
			return false
		}

		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if yamlFunctions[sel.Sel.Name] && isTestifySelector(sel, aliases) {
			found = true
			return false
		}

		return true
	})

	return found
}
