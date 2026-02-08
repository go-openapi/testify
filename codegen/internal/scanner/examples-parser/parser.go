// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/format"
	"go/printer"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/packages"
)

const tabWidth = 4

// Extractor explores a go package (including test code) and looks for testable examples.
type Extractor struct {
	dir       string
	pkg       string
	buildTags []string
}

// New [Extractor] for a given source package to be scanned.
//
// pkg is an import path or a relative pattern (e.g., "./assert") resolved from the working directory.
func New(pkg string, opts ...Option) *Extractor {
	e := &Extractor{
		pkg: pkg,
		dir: ".",
	}
	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Parse source code including test code.
//
// Builds an index of exported symbols (functions and types).
//
// Attaches all identified testable examples to the index.
//
// It may fail if the code doesn't compile.
func (e *Extractor) Parse() (Examples, error) {
	cfg := &packages.Config{
		Dir:   e.dir,
		Mode:  packages.NeedName | packages.NeedSyntax | packages.NeedCompiledGoFiles,
		Tests: true,
	}
	if len(e.buildTags) > 0 {
		cfg.BuildFlags = []string{"-tags", strings.Join(e.buildTags, ",")}
	}

	pkgs, err := packages.Load(cfg, e.pkg)
	if err != nil {
		return nil, err
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("package not found: %s", e.pkg)
	}

	// packages.Load reports resolution failures as package-level errors.
	for _, pkg := range pkgs {
		for _, pkgErr := range pkg.Errors {
			return nil, fmt.Errorf("loading %s: %s", pkg.ID, pkgErr.Msg)
		}
	}

	// With Tests: true, packages.Load returns:
	//   - the production package (ID = import path)
	//   - internal test variant (ID contains "[", Name = pkg name)
	//   - external test variant (ID contains "[", Name = pkg_test)
	var (
		exported  = make(map[string]bool)
		testFiles []*ast.File
		fset      *token.FileSet
	)

	for _, pkg := range pkgs {
		if fset == nil {
			fset = pkg.Fset
		}

		isTestVariant := strings.Contains(pkg.ID, "[")

		// From the production package, collect exported symbol names
		// (functions and types, so examples for types are linked too).
		if !isTestVariant {
			for _, file := range pkg.Syntax {
				collectExportedSymbols(file, exported)
			}
		}

		// From test variant packages, collect _test.go files.
		if isTestVariant {
			for i, file := range pkg.Syntax {
				if strings.HasSuffix(pkg.CompiledGoFiles[i], "_test.go") {
					testFiles = append(testFiles, file)
				}
			}
		}
	}

	if len(testFiles) == 0 {
		return make(Examples), nil
	}

	// Extract examples using go/doc.
	docExamples := doc.Examples(testFiles...)

	// Build the index: match examples to exported symbols.
	index := make(Examples)
	for _, ex := range docExamples {
		funcName := exampleFuncName(ex.Name)
		if funcName == "" {
			continue // package-level example, skip
		}
		if !exported[funcName] {
			continue // no matching exported symbol
		}

		te := TestableExample{
			Name:      ex.Name,
			Suffix:    exampleSuffix(ex.Name, funcName),
			Doc:       ex.Doc,
			Output:    ex.Output,
			WholeFile: isWholeFileExample(ex.Play),
			code:      ex.Code,
			play:      ex.Play,
			fset:      fset,
		}
		index[funcName] = append(index[funcName], te)
	}

	return index, nil
}

// collectExportedSymbols walks a file's declarations and records exported function
// and type names into the provided set.
func collectExportedSymbols(file *ast.File, exported map[string]bool) {
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv != nil {
				continue // skip methods
			}
			if d.Name.IsExported() {
				exported[d.Name.Name] = true
			}

		case *ast.GenDecl:
			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if ts.Name.IsExported() {
					exported[ts.Name.Name] = true
				}
			}
		}
	}
}

// exampleFuncName extracts the symbol name from a testable example name.
//
// Go naming convention:
//
//	"Equal"         -> "Equal"         (ExampleEqual)
//	"Equal_basic"   -> "Equal"         (ExampleEqual_basic)
//	"T_Method"      -> "T"             (type method, if T is a type)
//
// We take the leading identifier segment: everything before the first '_'
// that is followed by a lowercase letter (which marks a suffix).
func exampleFuncName(name string) string {
	if name == "" {
		return ""
	}

	// Find the first '_' that separates the identifier from the suffix.
	// The suffix must start with a lowercase letter (Go convention).
	for i, r := range name {
		if r == '_' && i+1 < len(name) {
			next := rune(name[i+1])
			if unicode.IsLower(next) {
				return name[:i]
			}
		}
	}

	return name
}

// exampleSuffix extracts the suffix part from a testable example name
// given the already-identified function name.
func exampleSuffix(name, funcName string) string {
	rest := strings.TrimPrefix(name, funcName)
	rest = strings.TrimPrefix(rest, "_")

	return rest
}

// isWholeFileExample reports whether a Play AST represents a whole-file example.
//
// A whole-file example has top-level declarations beyond the package clause,
// imports, and the main function (i.e., supporting types, helpers, etc.).
func isWholeFileExample(play *ast.File) bool {
	if play == nil {
		return false
	}

	for _, decl := range play.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name != "main" {
				return true
			}
		case *ast.GenDecl:
			if d.Tok != token.IMPORT {
				return true
			}
		}
	}

	return false
}

// Examples is an index of [TestableExample].
//
// Keys are exported symbol names (not fully qualified names).
//
// Each key may have 1 or several examples attached to it.
//
// Example:
//
// for function {package}.MyFunction, the key is "MyFunction".
type Examples map[string][]TestableExample

// TestableExample describes a go testable example and knows how to render as formatted go code.
type TestableExample struct {
	// Name is the full example name after the "Example" prefix.
	//
	// For ExampleEqual, Name is "Equal".
	// For ExampleEqual_basic, Name is "Equal_basic".
	Name string

	// Suffix is the example suffix, without leading '_'.
	//
	// For ExampleEqual, Suffix is "".
	// For ExampleEqual_basic, Suffix is "basic".
	Suffix string

	// Doc is the doc comment on the example function.
	Doc string

	// Output is the expected output string (from "// Output:" comments).
	Output string

	// WholeFile indicates this is a whole-file example with supporting
	// declarations (types, helper functions) outside the example function.
	WholeFile bool

	// unexported fields for rendering
	code ast.Node
	play *ast.File
	fset *token.FileSet
}

// Render the example as a formatted go code snippet.
//
// For whole-file examples, the complete file is rendered (minus package and imports),
// preserving all supporting declarations.
//
// For regular examples, only the function body is rendered.
//
// In both cases, the code is formatted with [format.Source].
func (x TestableExample) Render() string {
	// Render the full Play file as a standalone main program when available.
	if x.play != nil {
		return x.renderPlay()
	}

	return x.renderBody()

	// Previous routing: stripped package/imports/main scaffolding.
	// if x.WholeFile && x.play != nil {
	// 	return x.renderWholeFile()
	// }
	// return x.renderBody()
}

// renderPlay renders the Play AST as-is:
// a complete runnable program with package clause, imports, and func main().
func (x TestableExample) renderPlay() string {
	var buf bytes.Buffer
	p := printer.Config{Mode: printer.UseSpaces, Tabwidth: tabWidth}
	if err := p.Fprint(&buf, x.fset, x.play); err != nil {
		return ""
	}

	raw := buf.String()

	// make sure the output is well formatted
	formatted, err := format.Source([]byte(raw))
	if err != nil {
		return raw
	}

	return string(formatted)
}

// renderBody renders the example function body only.
func (x TestableExample) renderBody() string {
	if x.code == nil {
		return ""
	}

	// Print the raw AST node.
	var buf bytes.Buffer
	p := printer.Config{Mode: printer.UseSpaces, Tabwidth: tabWidth}
	if err := p.Fprint(&buf, x.fset, x.code); err != nil {
		return ""
	}

	body := buf.String()

	// Strip outer braces: the Code node is a *ast.BlockStmt.
	body = strings.TrimSpace(body)
	if strings.HasPrefix(body, "{") && strings.HasSuffix(body, "}") {
		body = body[1 : len(body)-1]
	}

	// Remove trailing "// Output:" comment lines before formatting.
	body = stripOutputComments(body)

	// Wrap in a synthetic file so format.Source can handle formatting.
	synthetic := "package p\n\nfunc f() {\n" + body + "\n}\n"

	formatted, err := format.Source([]byte(synthetic))
	if err != nil {
		return strings.TrimSpace(body)
	}

	return extractFuncBody(string(formatted))
}

/*
// renderWholeFile renders a whole-file example, stripping the package clause
// and imports, and renaming "func main()" back to the example function name.
func (x TestableExample) renderWholeFile() string {
	// Print the entire Play file.
	var buf bytes.Buffer
	p := printer.Config{Mode: printer.UseSpaces, Tabwidth: tabWidth}
	if err := p.Fprint(&buf, x.fset, x.play); err != nil {
		return ""
	}

	raw := buf.String()

	// Remove "// Output:" comments.
	raw = stripOutputComments(raw)

	// Format with goimports.
	formatted, err := imports.Process("example.go", []byte(raw), &imports.Options{
		Fragment:   true,
		FormatOnly: true,
	})
	if err != nil {
		formatted = []byte(raw)
	}

	// Strip package clause and imports, rename main -> Example function.
	return extractWholeFileBody(string(formatted), "Example"+x.Name)
}
*/

// extractWholeFileBody strips the package clause and import blocks from a
// formatted Go file, and renames "func main()" to the given example function name.
func extractWholeFileBody(src, exampleFuncName string) string {
	var result []string
	lines := strings.Split(src, "\n")

	inImportBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Skip package declaration.
		if strings.HasPrefix(trimmed, "package ") {
			continue
		}

		// Skip import blocks.
		if trimmed == "import (" {
			inImportBlock = true
			continue
		}
		if inImportBlock {
			if trimmed == ")" {
				inImportBlock = false
			}
			continue
		}

		// Skip single-line imports.
		if strings.HasPrefix(trimmed, "import ") {
			continue
		}

		// Rename "func main()" to the example function name.
		if strings.HasPrefix(trimmed, "func main()") {
			line = strings.Replace(line, "func main()", "func "+exampleFuncName+"()", 1)
		}

		result = append(result, line)
	}

	// Trim leading/trailing blank lines.
	for len(result) > 0 && strings.TrimSpace(result[0]) == "" {
		result = result[1:]
	}
	for len(result) > 0 && strings.TrimSpace(result[len(result)-1]) == "" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, "\n")
}

// stripOutputComments removes trailing "// Output:" lines and any
// blank lines that follow them at the end of the code block.
func stripOutputComments(code string) string {
	lines := strings.Split(code, "\n")

	for len(lines) > 0 {
		trimmed := strings.TrimSpace(lines[len(lines)-1])
		if trimmed == "" || strings.HasPrefix(trimmed, "// Output:") || strings.HasPrefix(trimmed, "// output:") {
			lines = lines[:len(lines)-1]
			continue
		}
		break
	}

	return strings.Join(lines, "\n")
}

// extractFuncBody extracts the body of "func f()" from a formatted synthetic file,
// stripping the wrapper and dedenting.
func extractFuncBody(src string) string {
	// Find "func f() {" and the closing "}".
	const openMarker = "func f() {"
	_, body, found := strings.Cut(src, openMarker)
	if !found {
		return strings.TrimSpace(src)
	}

	// Find the last "}" which closes the function.
	closeIdx := strings.LastIndex(body, "}")
	if closeIdx < 0 {
		return strings.TrimSpace(body)
	}
	body = body[:closeIdx]

	// Dedent by one tab (goimports uses tabs).
	lines := strings.Split(body, "\n")

	// Skip leading/trailing empty lines.
	start := 0
	for start < len(lines) && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	end := len(lines)
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}
	lines = lines[start:end]

	// Remove one leading tab from each line (goimports indents function body by one tab).
	for i, line := range lines {
		if strings.HasPrefix(line, "\t") {
			lines[i] = line[1:]
		}
	}

	return strings.Join(lines, "\n")
}
