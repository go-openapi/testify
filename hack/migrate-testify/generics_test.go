// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"iter"
	"slices"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

type genericsTestCase struct {
	name     string
	input    string
	expected string
}

func genericsTestCases() iter.Seq[genericsTestCase] {
	return slices.Values([]genericsTestCase{
		{
			name: "equal int upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestEqual(t *testing.T) {
	assert.Equal(t, 42, 42)
}`,
			expected: `assert.EqualT(t, 42, 42)`,
		},
		{
			name: "notequal string upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestNotEqual(t *testing.T) {
	assert.NotEqual(t, "a", "b")
}`,
			expected: `assert.NotEqualT(t, "a", "b")`,
		},
		{
			name: "greater upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestGreater(t *testing.T) {
	assert.Greater(t, 2, 1)
}`,
			expected: `assert.GreaterT(t, 2, 1)`,
		},
		{
			name: "positive upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestPositive(t *testing.T) {
	assert.Positive(t, 42)
}`,
			expected: `assert.PositiveT(t, 42)`,
		},
		{
			name: "skip any",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestSkipAny(t *testing.T) {
	var x any = 42
	assert.Equal(t, x, x)
}`,
			expected: `assert.Equal(t, x, x)`,
		},
		{
			name: "skip different types",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestSkipDifferent(t *testing.T) {
	assert.Equal(t, int32(1), int64(1))
}`,
			expected: `assert.Equal(t, int32(1), int64(1))`,
		},
		{
			name: "contains string upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestContains(t *testing.T) {
	assert.Contains(t, "hello world", "world")
}`,
			expected: `assert.StringContainsT(t, "hello world", "world")`,
		},
		{
			name: "contains slice upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestContains(t *testing.T) {
	assert.Contains(t, []int{1, 2, 3}, 2)
}`,
			expected: `assert.SliceContainsT(t, []int{1, 2, 3}, 2)`,
		},
		{
			name: "contains map upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestContains(t *testing.T) {
	assert.Contains(t, map[string]int{"a": 1}, "a")
}`,
			expected: `assert.MapContainsT(t, map[string]int{"a": 1}, "a")`,
		},
		{
			name: "true/false bool upgrade",
			input: `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestBool(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}`,
			expected: "assert.TrueT(t, true)\n\tassert.FalseT(t, false)",
		},
	})
}

func TestGenericUpgradeUnit(t *testing.T) {
	t.Parallel()

	for c := range genericsTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "test.go", c.input, parser.ParseComments)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}

			// Create a mock assert package so type-checking succeeds.
			info := typeCheckWithMockAssert(t, fset, f)
			if info == nil {
				t.Fatal("type-check failed")
			}

			// Build a fake *packages.Package to pass to upgradeFile.
			pkg := &packages.Package{
				TypesInfo: info,
				Syntax:    []*ast.File{f},
				GoFiles:   []string{"test.go"},
			}

			rpt := &report{}
			upgradeFile(f, pkg, fset, rpt, "test.go", true)

			// Extract the assertion call(s) from the output.
			var buf strings.Builder
			pcfg := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
			if err := pcfg.Fprint(&buf, fset, f); err != nil {
				t.Fatalf("print: %v", err)
			}

			got := buf.String()
			if !strings.Contains(got, c.expected) {
				t.Errorf("expected output to contain:\n  %s\ngot:\n%s", c.expected, got)
			}
		})
	}
}

func TestGenericUpgradeTracksUpgrades(t *testing.T) {
	t.Parallel()

	input := `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestEqual(t *testing.T) {
	assert.Equal(t, 42, 42)
	assert.True(t, true)
}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", input, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	info := typeCheckWithMockAssert(t, fset, f)
	if info == nil {
		t.Fatal("type-check failed")
	}

	pkg := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{f},
		GoFiles:   []string{"test.go"},
	}

	rpt := &report{}
	changes := upgradeFile(f, pkg, fset, rpt, "test.go", false)

	if changes != 2 {
		t.Errorf("expected 2 changes, got %d", changes)
	}
	if len(rpt.upgraded) != 2 {
		t.Errorf("expected 2 upgrade entries, got %d", len(rpt.upgraded))
	}
	if rpt.upgraded["Equal → EqualT"] != 1 {
		t.Errorf("expected Equal → EqualT upgrade, got %v", rpt.upgraded)
	}
	if rpt.upgraded["True → TrueT"] != 1 {
		t.Errorf("expected True → TrueT upgrade, got %v", rpt.upgraded)
	}
}

func TestGenericUpgradeTracksSkips(t *testing.T) {
	t.Parallel()

	// Test with any type — should be skipped with skipInterfaceType.
	input := `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestSkipAny(t *testing.T) {
	var x any = 42
	assert.Equal(t, x, x)
}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", input, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	info := typeCheckWithMockAssert(t, fset, f)
	if info == nil {
		t.Fatal("type-check failed")
	}

	pkg := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{f},
		GoFiles:   []string{"test.go"},
	}

	rpt := &report{}
	changes := upgradeFile(f, pkg, fset, rpt, "test.go", true)

	if changes != 0 {
		t.Errorf("expected 0 changes for any type, got %d", changes)
	}
	if len(rpt.skipped) == 0 {
		t.Error("expected skip to be tracked for any type")
	}
	if rpt.skipped[string(skipInterfaceType)] != 1 {
		t.Errorf("expected skipInterfaceType, got %v", rpt.skipped)
	}
}

func TestGenericUpgradeTracksDifferentTypeSkip(t *testing.T) {
	t.Parallel()

	input := `package p
import (
	"testing"
	"github.com/go-openapi/testify/v2/assert"
)
func TestSkipDifferent(t *testing.T) {
	assert.Equal(t, int32(1), int64(1))
}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", input, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	info := typeCheckWithMockAssert(t, fset, f)
	if info == nil {
		t.Fatal("type-check failed")
	}

	pkg := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{f},
		GoFiles:   []string{"test.go"},
	}

	rpt := &report{}
	changes := upgradeFile(f, pkg, fset, rpt, "test.go", true)

	if changes != 0 {
		t.Errorf("expected 0 changes for mismatched types, got %d", changes)
	}
	if len(rpt.skipped) == 0 {
		t.Error("expected skip to be tracked for type mismatch")
	}
	if rpt.skipped[string(skipTypeMismatch)] != 1 {
		t.Errorf("expected skipTypeMismatch, got %v", rpt.skipped)
	}
}

// typeCheckWithMockAssert creates a mock "assert" package that has a few functions
// taking (testing.T, any...) and type-checks the file against it.
func typeCheckWithMockAssert(t *testing.T, fset *token.FileSet, f *ast.File) *types.Info {
	t.Helper()

	// We need to provide packages that the source imports.
	// Use the default importer for stdlib.
	stdImporter := importer.ForCompiler(fset, "source", nil)

	// Create a fake "assert" package with the functions we need.
	assertPkg := types.NewPackage("github.com/go-openapi/testify/v2/assert", "assert")

	// T interface (just needs Errorf and Helper).
	tParam := types.NewVar(token.NoPos, assertPkg, "t", types.NewInterfaceType(nil, nil))
	anyType := types.NewInterfaceType(nil, nil)

	// Create variadic msgAndArgs param.
	msgAndArgsVar := types.NewVar(token.NoPos, assertPkg, "msgAndArgs", types.NewSlice(anyType))

	// Helper to create assertion function signatures.
	makeAssertFunc := func(name string, params ...*types.Var) {
		allParams := make([]*types.Var, 0, 1+len(params)+1)
		allParams = append(allParams, tParam)
		allParams = append(allParams, params...)
		allParams = append(allParams, msgAndArgsVar)

		sig := types.NewSignatureType(nil, nil, nil,
			types.NewTuple(allParams...),
			types.NewTuple(types.NewVar(token.NoPos, assertPkg, "", types.Typ[types.Bool])),
			true, // variadic
		)
		assertPkg.Scope().Insert(types.NewFunc(token.NoPos, assertPkg, name, sig))
	}

	// Add known assertion functions to the fake package.
	anyParam := func(name string) *types.Var {
		return types.NewVar(token.NoPos, assertPkg, name, anyType)
	}

	// Equal(t, expected any, actual any, ...)
	makeAssertFunc("Equal", anyParam("expected"), anyParam("actual"))
	makeAssertFunc("EqualT", anyParam("expected"), anyParam("actual"))
	makeAssertFunc("NotEqual", anyParam("expected"), anyParam("actual"))
	makeAssertFunc("NotEqualT", anyParam("expected"), anyParam("actual"))
	makeAssertFunc("Greater", anyParam("e1"), anyParam("e2"))
	makeAssertFunc("GreaterT", anyParam("e1"), anyParam("e2"))
	makeAssertFunc("Less", anyParam("e1"), anyParam("e2"))
	makeAssertFunc("LessT", anyParam("e1"), anyParam("e2"))
	makeAssertFunc("Positive", anyParam("e"))
	makeAssertFunc("PositiveT", anyParam("e"))
	makeAssertFunc("Negative", anyParam("e"))
	makeAssertFunc("NegativeT", anyParam("e"))
	makeAssertFunc("Contains", anyParam("s"), anyParam("contains"))
	makeAssertFunc("StringContainsT", anyParam("s"), anyParam("contains"))
	makeAssertFunc("SliceContainsT", anyParam("s"), anyParam("contains"))
	makeAssertFunc("MapContainsT", anyParam("s"), anyParam("contains"))
	makeAssertFunc("True", anyParam("value"))
	makeAssertFunc("TrueT", anyParam("value"))
	makeAssertFunc("False", anyParam("value"))
	makeAssertFunc("FalseT", anyParam("value"))

	assertPkg.MarkComplete()

	// Custom importer that returns our mock assert package.
	customImporter := &mockImporter{
		base: stdImporter,
		extra: map[string]*types.Package{
			"github.com/go-openapi/testify/v2/assert": assertPkg,
		},
	}

	conf := &types.Config{
		Importer: customImporter,
	}

	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	_, err := conf.Check("test", fset, []*ast.File{f}, info)
	if err != nil {
		t.Logf("type-check error (may be expected): %v", err)
		// Even with errors, the info map may have partial results.
		// Check if we have any type info at all.
		if len(info.Types) == 0 {
			return nil
		}
	}

	return info
}

type mockImporter struct {
	base  types.Importer
	extra map[string]*types.Package
}

func (m *mockImporter) Import(path string) (*types.Package, error) {
	if pkg, ok := m.extra[path]; ok {
		return pkg, nil
	}
	return m.base.Import(path)
}
