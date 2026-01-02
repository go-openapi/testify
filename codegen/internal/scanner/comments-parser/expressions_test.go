// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"go/token"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// TestParseTestValues tests parsing of test value expressions.
func TestParseTestValues(t *testing.T) {
	t.Parallel()

	for c := range parseTestValuesTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTestValues(c.input)

			if len(result) != c.expectedCount {
				t.Errorf("Expected %d values, got %d", c.expectedCount, len(result))
				return
			}

			for i, val := range result {
				t.Run("check single parsed test value", testSingleTestValue(c, val, i))
			}
		})
	}
}

func testSingleTestValue(c parseTestCase, val model.TestValue, i int) func(*testing.T) {
	return func(t *testing.T) {
		// Check raw string is preserved
		expectedRaw := c.expectedRaw[i]
		if val.Raw != expectedRaw {
			t.Errorf("Value %d: expected raw %q, got %q", i, expectedRaw, val.Raw)
		}

		// Check parse success/failure
		shouldParse := true // Default: expect successful parse
		if c.shouldParsePerVal != nil && i < len(c.shouldParsePerVal) {
			shouldParse = c.shouldParsePerVal[i]
		}

		if shouldParse {
			if val.Error != nil {
				t.Errorf("Value %d: unexpected parse error: %v", i, val.Error)
			}
			if val.Expr == nil {
				t.Errorf("Value %d: expected parsed expression, got nil", i)
			}
		} else if val.Error == nil {
			t.Errorf("Value %d: expected parse error, got none", i)
		}
	}
}

// TestParseExprWithFileSet tests parsing with FileSet for error reporting.
func TestParseExprWithFileSet(t *testing.T) {
	t.Parallel()

	fset := token.NewFileSet()

	for c := range parseExprWithFileSetTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result, err := ParseExprWithFileSet(fset, "test.go", c.input)

			if c.shouldError {
				if err == nil {
					t.Error("Expected error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if len(result) != c.expectedCount {
				t.Errorf("Expected %d values, got %d", c.expectedCount, len(result))
			}
		})
	}
}

/* Test case iterators */

type parseTestCase struct {
	name              string
	input             string
	expectedCount     int
	expectedRaw       []string
	shouldParsePerVal []bool // per-value parse expectations (nil means all should parse)
}

func parseTestValuesTestCases() iter.Seq[parseTestCase] {
	return slices.Values([]parseTestCase{
		{
			name:          "two integer literals",
			input:         "123, 456",
			expectedCount: 2,
			expectedRaw:   []string{"123", "456"},
		},
		{
			name:          "string literals",
			input:         `"hello", "world"`,
			expectedCount: 2,
			expectedRaw:   []string{`"hello"`, `"world"`},
		},
		{
			name:          "string literals with commas inside",
			input:         `"a,b", "c,d"`,
			expectedCount: 2,
			expectedRaw:   []string{`"a,b"`, `"c,d"`},
		},
		{
			name:          "mixed strings and values with commas",
			input:         `"hello, world", 123, "foo,bar,baz"`,
			expectedCount: 3,
			expectedRaw:   []string{`"hello, world"`, `123`, `"foo,bar,baz"`},
		},
		{
			name:          "nil and boolean",
			input:         "nil, true, false",
			expectedCount: 3,
			expectedRaw:   []string{"nil", "true", "false"},
		},
		{
			name:          "selector expressions",
			input:         "assertions.ErrTest, pkg.Type",
			expectedCount: 2,
			expectedRaw:   []string{"assertions.ErrTest", "pkg.Type"},
		},
		{
			name:          "composite literals",
			input:         "[]int{1,2,3}, []int{4,5,6}",
			expectedCount: 2,
			expectedRaw:   []string{"[]int{1, 2, 3}", "[]int{4, 5, 6}"}, // go/format adds spaces
		},
		{
			name:          "struct literals",
			input:         "pkg.Type{Field: value}, pkg.Type{}",
			expectedCount: 2,
			expectedRaw:   []string{"pkg.Type{Field: value}", "pkg.Type{}"},
		},
		{
			name:          "function calls",
			input:         "fn(a, b), fn(c)",
			expectedCount: 2,
			expectedRaw:   []string{"fn(a, b)", "fn(c)"},
		},
		{
			name:          "unary expressions",
			input:         "&x, *ptr, -123",
			expectedCount: 3,
			expectedRaw:   []string{"&x", "*ptr", "-123"},
		},
		{
			name:          "type assertions",
			input:         "x.(int), y.(string)",
			expectedCount: 2,
			expectedRaw:   []string{"x.(int)", "y.(string)"},
		},
		{
			name:          "empty string",
			input:         "",
			expectedCount: 0,
			expectedRaw:   []string{},
		},
		{
			name:          "whitespace only",
			input:         "   ",
			expectedCount: 0,
			expectedRaw:   []string{},
		},
		{
			name:          "single value",
			input:         "123",
			expectedCount: 1,
			expectedRaw:   []string{"123"},
		},
		{
			name:              "invalid expression",
			input:             "123 +",
			expectedCount:     1,
			expectedRaw:       []string{"123 +"},
			shouldParsePerVal: []bool{false},
		},
		{
			name:              "unclosed paren",
			input:             "(123",
			expectedCount:     1,
			expectedRaw:       []string{"(123"},
			shouldParsePerVal: []bool{false},
		},
		{
			name:              "invalid characters",
			input:             "@#$%",
			expectedCount:     1,
			expectedRaw:       []string{"@#$%"},
			shouldParsePerVal: []bool{false},
		},
		{
			name:              "mixed valid and invalid",
			input:             "123, @invalid",
			expectedCount:     1, // Parse fails for whole input, returns single error
			expectedRaw:       []string{"123, @invalid"},
			shouldParsePerVal: []bool{false}, // Entire input fails to parse
		},
	})
}

type parseWithFileSetTestCase struct {
	name          string
	input         string
	expectedCount int
	shouldError   bool // true if ParseExprWithFileSet should return an error
}

func parseExprWithFileSetTestCases() iter.Seq[parseWithFileSetTestCase] {
	return slices.Values([]parseWithFileSetTestCase{
		{
			name:          "valid expressions",
			input:         "123, 456",
			expectedCount: 2,
			shouldError:   false,
		},
		{
			name:          "invalid expression returns error",
			input:         "123 +",
			expectedCount: 1,
			shouldError:   true,
		},
		{
			name:          "first error is returned",
			input:         "123, @invalid, 789",
			expectedCount: 1, // Parse fails for whole input
			shouldError:   true,
		},
		{
			name:          "empty input",
			input:         "",
			expectedCount: 0,
			shouldError:   false,
		},
	})
}
