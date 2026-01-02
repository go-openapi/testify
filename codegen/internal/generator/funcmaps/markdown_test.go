// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"strings"
	"testing"
)

func TestMarkdownFormatEnhanced(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		contains    []string // strings that should appear in output
		notContains []string // strings that should NOT appear
	}{
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
				`[T]`,        // No package qualifier, left unchanged
				`[TestingT]`, // No package qualifier, left unchanged
			},
			notContains: []string{
				`[T](https://`, // Should NOT be converted to a link
				`[TestingT](https://`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMarkdown(tt.input)

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

func TestMarkdownFormatEnhanced_Output(t *testing.T) {
	input := `Empty asserts that the given value is "empty".

Zero values are "empty".

# Usage

	assertions.Empty(t, obj)

# Examples

	success: ""
	failure: "not empty"
[Zero values]: https://go.dev/ref/spec#The_zero_value`

	result := FormatMarkdown(input)
	t.Logf("Output:\n%s", result)
}
