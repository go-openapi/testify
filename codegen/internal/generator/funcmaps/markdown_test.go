// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

func TestMarkdownFormatEnhanced(t *testing.T) {
	t.Parallel()

	for tt := range markdownTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FormatMarkdown(tt.input, nil)

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("FormatMarkdown() missing expected string:\nwant: %q\ngot: %s", want, result)
				}
			}

			for _, notWant := range tt.notContains {
				if strings.Contains(result, notWant) {
					t.Errorf("FormatMarkdown() contains unexpected string:\ndon't want: %q\ngot: %s", notWant, result)
				}
			}
		})
	}
}

// mockRenderable implements model.Renderable for testing.
type mockRenderable struct {
	content string
}

func (m mockRenderable) Render() string {
	return m.content
}

// TestFormatMarkdownWithFunction verifies FormatMarkdown with a Function object
// that has Context and testable examples - this exercises the stripSections
// branch for Example sections and the addExamples path.
func TestFormatMarkdownWithFunction(t *testing.T) {
	t.Parallel()

	// Build a Function with Context that has ExtraPackages containing examples
	assertPkg := &model.AssertionPackage{
		Package:  "assert",
		Receiver: "Assertions",
		Functions: model.Functions{
			{
				Name: "Equal",
				Examples: []model.Renderable{
					mockRenderable{content: "assert.Equal(t, 1, 1)"},
				},
			},
		},
	}

	doc := &model.Document{
		ExtraPackages: model.ExtraPackages{assertPkg},
	}

	fn := model.Function{
		Name:    "Equal",
		Context: doc,
	}

	input := `Equal asserts that two objects are equal.

# Usage

	assertions.Equal(t, expected, actual)

# Examples

	success: 123, 123
	failure: 123, 456`

	result := FormatMarkdown(input, fn)

	// Should contain Hugo tab structure
	if !strings.Contains(result, `{{% expand title="Examples" %}}`) {
		t.Error("Expected expand shortcode in output")
	}
	if !strings.Contains(result, `{{< tabs >}}`) {
		t.Error("Expected tabs shortcode in output")
	}

	// Should contain testable examples tab
	if !strings.Contains(result, `Testable Examples (assert)`) {
		t.Error("Expected Testable Examples tab in output")
	}

	// Should contain the rendered example
	if !strings.Contains(result, "assert.Equal(t, 1, 1)") {
		t.Error("Expected rendered example code in output")
	}

	// Should contain cards shortcode
	if !strings.Contains(result, `{{% cards %}}`) {
		t.Error("Expected cards shortcode in output")
	}

	// Should contain Go Playground link
	if !strings.Contains(result, "go.dev/play") {
		t.Error("Expected Go Playground link in output")
	}
}

// TestFormatMarkdownWithFunctionNoExamples verifies the Examples section
// is rendered as a regular code tab when no testable examples exist.
func TestFormatMarkdownWithFunctionNoExamples(t *testing.T) {
	t.Parallel()

	fn := model.Function{
		Name:    "Equal",
		Context: nil, // no context
	}

	input := `Equal asserts that two objects are equal.

# Examples

	success: 123, 123`

	result := FormatMarkdown(input, fn)

	// Should still have tab structure
	if !strings.Contains(result, `{{% expand title="Examples" %}}`) {
		t.Error("Expected expand shortcode in output")
	}

	// Should NOT have testable examples (no context with packages)
	if strings.Contains(result, "Testable Examples") {
		t.Error("Should not have Testable Examples when no packages available")
	}
}

// TestFormatMarkdownWithMultipleSections verifies sections beyond Usage/Examples
// are rendered with increased heading depth.
func TestFormatMarkdownWithMultipleSections(t *testing.T) {
	t.Parallel()

	input := `Some description.

# Usage

	assertions.Something(t)

# Details

More info about the function.

# Examples

	success: true`

	result := FormatMarkdown(input, nil)

	// "Details" is not a tab-captured section, it should stay in the result
	// but with increased heading depth (# → ####)
	if !strings.Contains(result, "#### Details") {
		t.Error("Non-tab section should have increased heading depth")
	}

	// The description should be in the main body
	if !strings.Contains(result, "Some description.") {
		t.Error("Description should be in main body")
	}
}

func TestFormatMarkdownWithMultiplePackages(t *testing.T) {
	t.Parallel()

	assertPkg := &model.AssertionPackage{
		Package: "assert",
		Functions: model.Functions{
			{
				Name: "Equal",
				Examples: []model.Renderable{
					mockRenderable{content: "assert.Equal(t, 1, 1)"},
				},
			},
		},
	}
	requirePkg := &model.AssertionPackage{
		Package: "require",
		Functions: model.Functions{
			{
				Name: "Equal",
				Examples: []model.Renderable{
					mockRenderable{content: "require.Equal(t, 1, 1)"},
				},
			},
		},
	}

	doc := &model.Document{
		ExtraPackages: model.ExtraPackages{assertPkg, requirePkg},
	}

	fn := model.Function{
		Name:    "Equal",
		Context: doc,
	}

	input := `Equal asserts equality.

# Examples

	success: 123, 123`

	result := FormatMarkdown(input, fn)

	// Should have tabs for both packages
	if !strings.Contains(result, `Testable Examples (assert)`) {
		t.Error("Expected assert examples tab")
	}
	if !strings.Contains(result, `Testable Examples (require)`) {
		t.Error("Expected require examples tab")
	}

	// Both rendered examples should appear
	if !strings.Contains(result, "assert.Equal(t, 1, 1)") {
		t.Error("Expected assert rendered example")
	}
	if !strings.Contains(result, "require.Equal(t, 1, 1)") {
		t.Error("Expected require rendered example")
	}
}

func TestFormatMarkdownOutput_Debug(t *testing.T) {
	input := `Empty asserts that the given value is "empty".

Zero values are "empty".

Values can be of type [strings.Builder] or [Boolean].

# Usage

	assertions.Empty(t, obj)

# Examples

	success: ""
	failure: "not empty"
[Zero values]: https://go.dev/ref/spec#The_zero_value`

	result := FormatMarkdown(input, nil)
	t.Logf("Output:\n%s", result)
}

type markdownTestCase struct {
	name        string
	input       string
	contains    []string // strings that should appear in output
	notContains []string // strings that should NOT appear
}

func markdownTestCases() iter.Seq[markdownTestCase] {
	return slices.Values([]markdownTestCase{
		{
			name: "reference-style markdown link (dangling)",
			input: `Empty asserts that the given value is "empty".

# Usage

	assertions.Empty(t, obj)

# Examples

	success: ""
	failure: "not empty"
[Zero values]: https://go.dev/ref/spec#The_zero_value`,
			contains: []string{
				`[Zero values]: https://go.dev/ref/spec#The_zero_value`, // Dangling ref appears after expand
				`{{% expand title="Examples" %}}`,
				`{{< tabs >}}`,
				`{{% /expand %}}`,
			},
			notContains: []string{
				// The reference shouldn't appear inside code blocks
			},
		},
		{
			name: "godoc-style link",
			input: `EqualError asserts that a function returned an error equal to the provided error.

This is a wrapper for [errors.Is].

# Usage

	assertions.EqualError(t, err, expectedErrorString)`,
			contains: []string{
				`[errors.Is](https://pkg.go.dev/errors#Is)`,
			},
			notContains: []string{
				`[errors.Is].`, // Should be converted
			},
		},
		{
			name: "multiple godoc links",
			input: `See [testing.T] and [fmt.Printf] for details.

# Usage

	assertions.Something(t)`,
			contains: []string{
				`[testing.T](https://pkg.go.dev/testing#T)`,
				`[fmt.Printf](https://pkg.go.dev/fmt#Printf)`,
			},
		},
		{
			name: "mixed markdown and godoc links (dangling ref)",
			input: `Uses [errors.Is] to check errors.

# Examples

	success: nil
	failure: ErrTest
[error wrapping]: https://go.dev/blog/go1.13-errors`,
			contains: []string{
				`[errors.Is](https://pkg.go.dev/errors#Is)`,
				`[error wrapping]: https://go.dev/blog/go1.13-errors`, // Dangling ref
			},
			notContains: []string{
				// Nothing
			},
		},
		{
			name: "used reference link",
			input: `Empty asserts values are empty.

See [Zero values] for definition.

# Examples

	success: ""
[Zero values]: https://go.dev/ref/spec#The_zero_value`,
			contains: []string{
				`[Zero values](https://go.dev/ref/spec#The_zero_value)`, // Converted to inline
			},
			notContains: []string{
				`[Zero values]:`, // Definition removed since it was used
			},
		},
		{
			name: "identifier without package qualifier unchanged",
			input: `This function uses [T] as a type parameter.

It also references [TestingT] interface.

# Usage

	assertions.Something[T TestingT](t)`,
			contains: []string{
				`[T](https://pkg.go.dev/github.com/go-openapi`,        // No package qualifier, link to the go doc for our lib
				`[TestingT](https://pkg.go.dev/github.com/go-openapi`, // No package qualifier, link to the go doc for our lib
			},
			notContains: []string{
				`[T](https://pkg.go.dev#T`, // Should not be appended directly
				`[TestingT](https://pkg.go.dev#T`,
			},
		},
	})
}
