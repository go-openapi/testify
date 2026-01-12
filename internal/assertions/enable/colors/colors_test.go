// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"iter"
	"slices"
	"testing"
)

func TestStringColorizer(t *testing.T) {
	t.Parallel()

	for tc := range colorizerTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := tc.colorizer(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestMakeColorizer(t *testing.T) {
	t.Parallel()

	t.Run("creates colorizer with custom mark", func(t *testing.T) {
		t.Parallel()

		customMark := "\033[1;35m" // bold magenta
		colorizer := makeColorizer(customMark)

		result := colorizer("test")
		expected := "\033[1;35mtest\033[0m"

		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	})

	t.Run("colorizer is reusable", func(t *testing.T) {
		t.Parallel()

		colorizer := makeColorizer(greenMark)

		result1 := colorizer("first")
		result2 := colorizer("second")

		expected1 := "\033[0;32mfirst\033[0m"
		expected2 := "\033[0;32msecond\033[0m"

		if result1 != expected1 {
			t.Errorf("first call: expected %q, got %q", expected1, result1)
		}
		if result2 != expected2 {
			t.Errorf("second call: expected %q, got %q", expected2, result2)
		}
	})
}

func TestNoopColorizer(t *testing.T) {
	t.Parallel()

	inputs := []string{
		"",
		"simple",
		"with\nnewline",
		"\033[0;31malready colored\033[0m",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			t.Parallel()

			result := noopColorizer(input)
			if result != input {
				t.Errorf("noopColorizer should return input unchanged: expected %q, got %q", input, result)
			}
		})
	}
}

type colorizerTestCase struct {
	name      string
	colorizer StringColorizer
	input     string
	expected  string
}

func colorizerTestCases() iter.Seq[colorizerTestCase] {
	return slices.Values([]colorizerTestCase{
		{
			name:      "green colorizer",
			colorizer: greenColorizer,
			input:     "hello",
			expected:  "\033[0;32mhello\033[0m",
		},
		{
			name:      "red colorizer",
			colorizer: redColorizer,
			input:     "world",
			expected:  "\033[0;31mworld\033[0m",
		},
		{
			name:      "bright green colorizer",
			colorizer: brightGreenColorizer,
			input:     "expected",
			expected:  "\033[0;92mexpected\033[0m",
		},
		{
			name:      "bright red colorizer",
			colorizer: brightRedColorizer,
			input:     "actual",
			expected:  "\033[0;91mactual\033[0m",
		},
		{
			name:      "noop colorizer",
			colorizer: noopColorizer,
			input:     "unchanged",
			expected:  "unchanged",
		},
		{
			name:      "empty string",
			colorizer: greenColorizer,
			input:     "",
			expected:  "\033[0;32m\033[0m",
		},
		{
			name:      "string with special characters",
			colorizer: redColorizer,
			input:     "line1\nline2\ttab",
			expected:  "\033[0;31mline1\nline2\ttab\033[0m",
		},
		{
			name:      "string with existing ANSI codes",
			colorizer: greenColorizer,
			input:     "\033[1mbold\033[0m",
			expected:  "\033[0;32m\033[1mbold\033[0m\033[0m",
		},
	})
}
