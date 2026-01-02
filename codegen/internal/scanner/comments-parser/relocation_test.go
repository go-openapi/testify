// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"iter"
	"slices"
	"testing"
)

// TestRelocateTestValue tests AST-based relocation of test values.
func TestRelocateTestValue(t *testing.T) {
	t.Parallel()

	for c := range relocateTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// Parse the input
			values := ParseTestValues(c.input)
			if len(values) != 1 {
				t.Fatalf("Expected 1 value, got %d", len(values))
			}

			original := values[0]
			if original.Error != nil {
				t.Fatalf("Parse error: %v", original.Error)
			}

			// Relocate
			relocated := RelocateTestValue(original, c.fromPkg, c.toPkg)

			if relocated.Error != nil {
				t.Fatalf("Relocation error: %v", relocated.Error)
			}

			if relocated.Raw != c.expected {
				t.Errorf("Expected %q, got %q", c.expected, relocated.Raw)
			}
		})
	}
}

/* Test case iterators */

type relocateTestCase struct {
	name     string
	input    string
	fromPkg  string
	toPkg    string
	expected string
}

func relocateTestCases() iter.Seq[relocateTestCase] {
	return slices.Values([]relocateTestCase{
		{
			name:     "qualified selector: assertions.ErrTest → assert.ErrTest",
			input:    "assertions.ErrTest",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.ErrTest",
		},
		{
			name:     "qualified selector: assertions.CollectT → assert.CollectT",
			input:    "assertions.CollectT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.CollectT",
		},
		{
			name:     "unqualified type: ErrTest → assert.ErrTest",
			input:    "ErrTest",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.ErrTest", // Should qualify unqualified exported ident
		},
		{
			name:     "unqualified type: CollectT → assert.CollectT",
			input:    "CollectT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.CollectT",
		},
		{
			name:     "exception: assertions.PanicTestFunc stays assertions.PanicTestFunc",
			input:    "assertions.PanicTestFunc",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assertions.PanicTestFunc", // Exception: keep assertions
		},
		{
			name:     "exception: PanicTestFunc → assertions.PanicTestFunc",
			input:    "PanicTestFunc",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assertions.PanicTestFunc", // Exception even when unqualified
		},
		{
			name:     "composite literal: assertions.ErrTest{} → assert.ErrTest{}",
			input:    "assertions.ErrTest{}",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.ErrTest{}",
		},
		{
			name:     "composite literal unqualified: CollectT{} → assert.CollectT{}",
			input:    "CollectT{}",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.CollectT{}",
		},
		{
			name:     "function call: assertions.NewT() → assert.NewT()",
			input:    "assertions.NewT()",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.NewT()",
		},
		{
			name:     "function call unqualified: NewT() → assert.NewT()",
			input:    "NewT()",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.NewT()",
		},
		{
			name:     "lowercase identifier unchanged: true",
			input:    "true",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "true", // Don't qualify lowercase/built-ins
		},
		{
			name:     "nil unchanged",
			input:    "nil",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "nil",
		},
		{
			name:     "string literal unchanged",
			input:    `"hello"`,
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: `"hello"`,
		},
		{
			name:     "integer literal unchanged",
			input:    "123",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "123",
		},
		{
			name:     "complex expression with multiple relocations",
			input:    "func(t assertions.TestingT, fn assertions.PanicTestFunc) bool { return true }",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "func(t assert.TestingT, fn assertions.PanicTestFunc) bool {\n\treturn true\n}", // go/format adds newlines
		},
		{
			name:     "pointer type: *CollectT → *assert.CollectT",
			input:    "*CollectT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "*assert.CollectT",
		},
		{
			name:     "function with pointer parameter: func(c *CollectT) → func(c *assert.CollectT)",
			input:    "func(c *CollectT) { }",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "func(c *assert.CollectT) {\n}", // go/format adds newline
		},
		{
			name:     "slice type: []TestingT → []assert.TestingT",
			input:    "[]TestingT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "[]assert.TestingT",
		},
		{
			name:     "map type: map[string]TestingT → map[string]assert.TestingT",
			input:    "map[string]TestingT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "map[string]assert.TestingT",
		},
		{
			name:     "channel type: chan CollectT → chan assert.CollectT",
			input:    "chan CollectT",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "chan assert.CollectT",
		},
		{
			name:     "method call on unqualified variable: ErrTest.Error() → assert.ErrTest.Error()",
			input:    "ErrTest.Error()",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.ErrTest.Error()",
		},
		{
			name:     "field access on unqualified variable: Config.Value → assert.Config.Value",
			input:    "Config.Value",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "assert.Config.Value",
		},
		{
			name:     "function call with unqualified argument: panic(ErrTest) → panic(assert.ErrTest)",
			input:    "panic(ErrTest)",
			fromPkg:  "assertions",
			toPkg:    "assert",
			expected: "panic(assert.ErrTest)",
		},
	})
}
