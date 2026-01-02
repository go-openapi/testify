// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/generator/funcmaps"
	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

const (
	pkgSource = "github.com/go-openapi/testify/v2/internal/assertions"
)

// TestNew verifies that New creates a valid Generator instance.
func TestNew(t *testing.T) {
	t.Parallel()

	source := model.New()
	source.Package = pkgSource
	source.DocString = "Package assertions provides test assertions"
	source.Copyright = "Copyright 2025"

	gen := New(source)

	if gen == nil {
		t.Fatal("New() returned nil")
	}

	if gen.source != source {
		t.Error("Generator source not set correctly")
	}
}

// TestGeneratorInitContext verifies context initialization.
func TestGeneratorInitContext(t *testing.T) {
	t.Parallel()
	source := model.New()
	source.Package = pkgSource

	for tt := range generatorCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gen := New(source)
			err := gen.initContext(tt.opts)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tt.checkFunc != nil {
					tt.checkFunc(t, gen)
				}
			}
		})
	}
}

// TestTransformArgs verifies argument transformation for generation.
func TestTransformArgs(t *testing.T) {
	t.Parallel()

	source := model.New()
	gen := New(source)

	for tt := range transformCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			params := gen.transformArgs(tt.input)
			result := funcmaps.PrintReturns(params)

			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// TestLoadTemplates verifies template loading for assert and require packages.
func TestLoadTemplates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		targetPkg   string
		expectError bool
		checkFunc   func(*testing.T, *Generator)
	}{
		{
			name:        "load assert templates",
			targetPkg:   pkgAssert,
			expectError: false,
			checkFunc: func(t *testing.T, g *Generator) {
				t.Helper()

				if len(g.ctx.templates) == 0 {
					t.Error("No templates loaded")
				}
				if _, ok := g.ctx.templates["assertion_assertions"]; !ok {
					t.Error("Missing assertion_assertions template")
				}
			},
		},
		{
			name:        "load require templates",
			targetPkg:   pkgRequire,
			expectError: false,
			checkFunc: func(t *testing.T, g *Generator) {
				t.Helper()

				if len(g.ctx.templates) == 0 {
					t.Error("No templates loaded")
				}
				if _, ok := g.ctx.templates["requirement_assertions"]; !ok {
					t.Error("Missing requirement_assertions template")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			source := model.New()
			source.Package = pkgSource

			gen := New(source)
			err := gen.initContext([]GenerateOption{
				WithTargetPackage(tt.targetPkg),
				WithTargetRoot(t.TempDir()),
			})
			if err != nil {
				t.Fatalf("Failed to init context: %v", err)
			}

			err = gen.loadTemplates()

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tt.checkFunc != nil {
					tt.checkFunc(t, gen)
				}
			}
		})
	}
}

type generatorCase struct {
	name        string
	opts        []GenerateOption
	expectError bool
	checkFunc   func(*testing.T, *Generator)
}

func generatorCases(t *testing.T) iter.Seq[generatorCase] {
	t.Helper()

	return slices.Values([]generatorCase{
		{
			name:        "missing target package",
			opts:        []GenerateOption{},
			expectError: true,
		},
		{
			name: "invalid target package",
			opts: []GenerateOption{
				WithTargetPackage("invalid"),
			},
			expectError: true,
		},
		{
			name: "valid assert package",
			opts: []GenerateOption{
				WithTargetPackage(pkgAssert),
				WithTargetRoot(t.TempDir()),
			},
			expectError: false,
			checkFunc: func(t *testing.T, g *Generator) {
				t.Helper()

				if g.ctx.targetBase != pkgAssert {
					t.Errorf("Expected targetBase '%s', got %q", pkgAssert, g.ctx.targetBase)
				}
			},
		},
		{
			name: "valid require package",
			opts: []GenerateOption{
				WithTargetPackage(pkgRequire),
				WithTargetRoot(t.TempDir()),
			},
			expectError: false,
			checkFunc: func(t *testing.T, g *Generator) {
				t.Helper()

				if g.ctx.targetBase != pkgRequire {
					t.Errorf("Expected targetBase '%s', got %q", pkgRequire, g.ctx.targetBase)
				}
			},
		},
	})
}

type transformCase struct {
	name     string
	input    model.Parameters
	expected string
}

func transformCases() iter.Seq[transformCase] {
	return slices.Values([]transformCase{
		{
			name:     "empty parameters",
			input:    model.Parameters{},
			expected: "",
		},
		{
			name: "single parameter without name",
			input: model.Parameters{
				{GoType: "bool"},
			},
			expected: "bool",
		},
		{
			name: "single parameter with name",
			input: model.Parameters{
				{Name: "value", GoType: "bool"},
			},
			expected: "(value bool)",
		},
		{
			name: "multiple parameters",
			input: model.Parameters{
				{Name: "expected", GoType: "any"},
				{Name: "actual", GoType: "any"},
			},
			expected: "(expected any, actual any)",
		},
	})
}
