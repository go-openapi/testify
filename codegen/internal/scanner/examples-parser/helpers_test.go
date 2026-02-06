// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"iter"
	"slices"
	"strings"
	"testing"
)

func TestExampleFuncName(t *testing.T) {
	t.Parallel()

	for c := range exampleFuncNameCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := exampleFuncName(c.input)
			if got != c.expected {
				t.Errorf("exampleFuncName(%q) = %q, expected %q", c.input, got, c.expected)
			}
		})
	}
}

type exampleFuncNameCase struct {
	name     string
	input    string
	expected string
}

func exampleFuncNameCases() iter.Seq[exampleFuncNameCase] {
	return slices.Values([]exampleFuncNameCase{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "simple function name",
			input:    "Equal",
			expected: "Equal",
		},
		{
			name:     "function with suffix",
			input:    "Equal_basic",
			expected: "Equal",
		},
		{
			name:     "underscore followed by uppercase is part of name",
			input:    "T_Method",
			expected: "T_Method",
		},
		{
			name:     "multiple underscores with suffix",
			input:    "Foo_Bar_baz",
			expected: "Foo_Bar",
		},
		{
			name:     "trailing underscore without next char",
			input:    "Foo_",
			expected: "Foo_",
		},
		{
			name:     "single character",
			input:    "X",
			expected: "X",
		},
		{
			name:     "underscore only",
			input:    "_",
			expected: "_",
		},
		{
			name:     "suffix starts immediately after underscore",
			input:    "Contains_slice",
			expected: "Contains",
		},
		{
			name:     "multiple suffixes only first split applies",
			input:    "HTTPStatusCode_redirect_permanent",
			expected: "HTTPStatusCode",
		},
	})
}

func TestExampleSuffix(t *testing.T) {
	t.Parallel()

	for c := range exampleSuffixCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := exampleSuffix(c.exampleName, c.funcName)
			if got != c.expected {
				t.Errorf("exampleSuffix(%q, %q) = %q, expected %q", c.exampleName, c.funcName, got, c.expected)
			}
		})
	}
}

type exampleSuffixCase struct {
	name        string
	exampleName string
	funcName    string
	expected    string
}

func exampleSuffixCases() iter.Seq[exampleSuffixCase] {
	return slices.Values([]exampleSuffixCase{
		{
			name:        "no suffix",
			exampleName: "Equal",
			funcName:    "Equal",
			expected:    "",
		},
		{
			name:        "simple suffix",
			exampleName: "Equal_basic",
			funcName:    "Equal",
			expected:    "basic",
		},
		{
			name:        "suffix with uppercase name",
			exampleName: "Foo_Bar_baz",
			funcName:    "Foo_Bar",
			expected:    "baz",
		},
		{
			name:        "multi-word suffix",
			exampleName: "Contains_with_custom_message",
			funcName:    "Contains",
			expected:    "with_custom_message",
		},
	})
}

func TestStripOutputComments(t *testing.T) {
	t.Parallel()

	for c := range stripOutputCommentsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := stripOutputComments(c.input)
			if got != c.expected {
				t.Errorf("stripOutputComments() =\n%q\nexpected:\n%q", got, c.expected)
			}
		})
	}
}

type stripOutputCommentsCase struct {
	name     string
	input    string
	expected string
}

func stripOutputCommentsCases() iter.Seq[stripOutputCommentsCase] {
	return slices.Values([]stripOutputCommentsCase{
		{
			name:     "no output comment",
			input:    "x := 1\ny := 2",
			expected: "x := 1\ny := 2",
		},
		{
			name:     "trailing output comment",
			input:    "fmt.Println(x)\n// Output: 42",
			expected: "fmt.Println(x)",
		},
		{
			name:     "lowercase output comment",
			input:    "fmt.Println(x)\n// output: 42",
			expected: "fmt.Println(x)",
		},
		{
			name:     "output comment followed by blank lines",
			input:    "fmt.Println(x)\n// Output: 42\n\n",
			expected: "fmt.Println(x)",
		},
		{
			name:     "only trailing blanks stripped",
			input:    "// Output: first\nfmt.Println(x)\n// Output: 42",
			expected: "// Output: first\nfmt.Println(x)",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only output comment",
			input:    "// Output: 42",
			expected: "",
		},
		{
			name:     "multiple trailing output lines",
			input:    "x := 1\n// Output:\n// output: result",
			expected: "x := 1",
		},
	})
}

func TestExtractFuncBody(t *testing.T) {
	t.Parallel()

	for c := range extractFuncBodyCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := extractFuncBody(c.input)
			if got != c.expected {
				t.Errorf("extractFuncBody() =\n%q\nexpected:\n%q", got, c.expected)
			}
		})
	}
}

type extractFuncBodyCase struct {
	name     string
	input    string
	expected string
}

func extractFuncBodyCases() iter.Seq[extractFuncBodyCase] {
	return slices.Values([]extractFuncBodyCase{
		{
			name:     "standard synthetic file",
			input:    "package p\n\nfunc f() {\n\tx := 1\n\ty := 2\n}\n",
			expected: "x := 1\ny := 2",
		},
		{
			name:     "no function marker",
			input:    "package p\n\nfunc g() {\n\tx := 1\n}\n",
			expected: "package p\n\nfunc g() {\n\tx := 1\n}",
		},
		{
			name:     "empty function body",
			input:    "package p\n\nfunc f() {\n}\n",
			expected: "",
		},
		{
			name:     "single statement",
			input:    "package p\n\nfunc f() {\n\treturn\n}\n",
			expected: "return",
		},
		{
			name:     "no closing brace",
			input:    "package p\n\nfunc f() {\n\tx := 1\n",
			expected: "x := 1",
		},
	})
}

func TestExtractWholeFileBody(t *testing.T) {
	t.Parallel()

	for c := range extractWholeFileBodyCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := extractWholeFileBody(c.input, c.funcName)
			if got != c.expected {
				t.Errorf("extractWholeFileBody() =\n%q\nexpected:\n%q", got, c.expected)
			}
		})
	}
}

type extractWholeFileBodyCase struct {
	name     string
	input    string
	funcName string
	expected string
}

func extractWholeFileBodyCases() iter.Seq[extractWholeFileBodyCase] {
	return slices.Values([]extractWholeFileBodyCase{
		{
			name: "strips package imports and renames main",
			input: `package main

import "fmt"

func main() {
    fmt.Println("hello")
}`,
			funcName: "ExampleFoo",
			expected: `func ExampleFoo() {
    fmt.Println("hello")
}`,
		},
		{
			name: "strips grouped imports",
			input: `package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("hello"))
}`,
			funcName: "ExampleBar",
			expected: `func ExampleBar() {
    fmt.Println(strings.ToUpper("hello"))
}`,
		},
		{
			name: "preserves supporting declarations",
			input: `package main

import "fmt"

type myStruct struct {
    value int
}

func main() {
    s := myStruct{value: 42}
    fmt.Println(s.value)
}

func helper() int {
    return 1
}`,
			funcName: "ExampleBaz",
			expected: `type myStruct struct {
    value int
}

func ExampleBaz() {
    s := myStruct{value: 42}
    fmt.Println(s.value)
}

func helper() int {
    return 1
}`,
		},
		{
			name: "handles single-line import",
			input: `package main

import "fmt"

func main() {
    fmt.Println("ok")
}`,
			funcName: "ExampleSingle",
			expected: `func ExampleSingle() {
    fmt.Println("ok")
}`,
		},
		{
			name:     "empty input",
			input:    "",
			funcName: "ExampleEmpty",
			expected: "",
		},
	})
}

func TestIsWholeFileExample(t *testing.T) {
	t.Parallel()

	for c := range isWholeFileExampleCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := isWholeFileExample(c.file)
			if got != c.expected {
				t.Errorf("isWholeFileExample() = %v, expected %v", got, c.expected)
			}
		})
	}
}

type isWholeFileExampleCase struct {
	name     string
	file     *ast.File
	expected bool
}

func isWholeFileExampleCases() iter.Seq[isWholeFileExampleCase] {
	return slices.Values([]isWholeFileExampleCase{
		{
			name:     "nil file",
			file:     nil,
			expected: false,
		},
		{
			name:     "main only with imports",
			file:     parseFile(`package main; import "fmt"; func main() { fmt.Println() }`),
			expected: false,
		},
		{
			name:     "has helper function",
			file:     parseFile(`package main; func main() {}; func helper() {}`),
			expected: true,
		},
		{
			name:     "has type declaration",
			file:     parseFile(`package main; type Foo struct{}; func main() {}`),
			expected: true,
		},
		{
			name:     "has const declaration",
			file:     parseFile(`package main; const x = 1; func main() {}`),
			expected: true,
		},
		{
			name:     "has var declaration",
			file:     parseFile(`package main; var x int; func main() {}`),
			expected: true,
		},
		{
			name:     "only main and imports",
			file:     parseFile(`package main; import "os"; func main() { _ = os.Args }`),
			expected: false,
		},
	})
}

func TestCollectExportedSymbols(t *testing.T) {
	t.Parallel()

	for c := range collectExportedSymbolsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			exported := make(map[string]bool)
			file := parseFile(c.src)
			collectExportedSymbols(file, exported)

			for _, name := range c.want {
				if !exported[name] {
					t.Errorf("expected %q to be exported, but it was not collected", name)
				}
			}
			for _, name := range c.dontWant {
				if exported[name] {
					t.Errorf("expected %q to NOT be collected, but it was", name)
				}
			}

			if len(exported) != len(c.want) {
				t.Errorf("collected %d symbols, expected %d: %v", len(exported), len(c.want), exported)
			}
		})
	}
}

type collectExportedSymbolsCase struct {
	name     string
	src      string
	want     []string
	dontWant []string
}

func collectExportedSymbolsCases() iter.Seq[collectExportedSymbolsCase] {
	return slices.Values([]collectExportedSymbolsCase{
		{
			name:     "exported functions",
			src:      `package p; func Exported() {}; func another() {}`,
			want:     []string{"Exported"},
			dontWant: []string{"another"},
		},
		{
			name:     "methods are skipped",
			src:      `package p; type T struct{}; func (t T) Method() {}; func Standalone() {}`,
			want:     []string{"T", "Standalone"},
			dontWant: []string{"Method"},
		},
		{
			name:     "exported types",
			src:      `package p; type Public struct{}; type private struct{}`,
			want:     []string{"Public"},
			dontWant: []string{"private"},
		},
		{
			name:     "mixed functions and types",
			src:      `package p; func Alpha() {}; type Beta int; func gamma() {}; type delta struct{}`,
			want:     []string{"Alpha", "Beta"},
			dontWant: []string{"gamma", "delta"},
		},
		{
			name:     "no exported symbols",
			src:      `package p; func hidden() {}; type secret struct{}`,
			want:     nil,
			dontWant: []string{"hidden", "secret"},
		},
		{
			name:     "multiple types in one declaration",
			src:      `package p; type ( Foo int; Bar string; baz float64 )`,
			want:     []string{"Foo", "Bar"},
			dontWant: []string{"baz"},
		},
	})
}

// parseFile is a test helper that parses a Go source string into an *ast.File.
func parseFile(src string) *ast.File {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		// Only used in test case setup; a parse error indicates a broken test case.
		panic("parseFile: " + err.Error())
	}

	return file
}

func TestRenderBody_FormatsCode(t *testing.T) {
	t.Parallel()

	// Parse a real example from the assert package and verify Render
	// produces output that:
	//  1. does not contain outer braces
	//  2. does not contain "// Output:" lines
	//  3. is non-empty

	for c := range renderBodyPropertyCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			rendered := c.example.Render()
			if rendered == "" {
				t.Fatal("Render() returned empty string")
			}

			t.Run("should not start with outer braces", func(t *testing.T) {
				trimmed := strings.TrimSpace(rendered)
				if strings.HasPrefix(trimmed, "{") {
					t.Errorf("got:\n%s", rendered)
				}
			})

			t.Run("should not contain output comments", func(t *testing.T) {
				if strings.Contains(rendered, "// Output:") {
					t.Errorf("got:\n%s", rendered)
				}
			})

			t.Run("should not contain package clause", func(t *testing.T) {
				if strings.Contains(rendered, "package ") {
					t.Errorf("got:\n%s", rendered)
				}
			})
		})
	}
}

type renderBodyPropertyCase struct {
	name    string
	example TestableExample
}

func renderBodyPropertyCases() iter.Seq[renderBodyPropertyCase] {
	// Build simple AST examples programmatically to avoid depending on the
	// full package loader in unit tests.
	fset := token.NewFileSet()
	src := `package p

func Example() {
	x := 1
	_ = x
	// Output: 1
}
`
	file, err := parser.ParseFile(fset, "ex_test.go", src, parser.ParseComments)
	if err != nil {
		panic("renderBodyPropertyCases: " + err.Error())
	}

	fn, ok := file.Decls[0].(*ast.FuncDecl)
	if !ok {
		panic("renderBodyPropertyCases: expected *ast.FuncDecl")
	}

	return slices.Values([]renderBodyPropertyCase{
		{
			name: "simple body example",
			example: TestableExample{
				Name: "Example",
				code: fn.Body,
				fset: fset,
			},
		},
	})
}
