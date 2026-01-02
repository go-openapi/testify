// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"

	"github.com/go-openapi/testify/v2/internal/assertions/enable/yaml"
)

// YAMLEq asserts that the first documents in the two YAML strings are equivalent.
//
// # Usage
//
//	expected := `---
//	key: value
//	---
//	key: this is a second document, it is not evaluated
//	`
//	actual := `---
//	key: value
//	---
//	key: this is a subsequent document, it is not evaluated
//	`
//	assertions.YAMLEq(t, expected, actual)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
func YAMLEq(t T, expected string, actual string, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}
	var expectedYAMLAsInterface, actualYAMLAsInterface any

	if err := yaml.Unmarshal([]byte(expected), &expectedYAMLAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Expected value ('%s') is not valid yaml.\nYAML parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	// Shortcut if same bytes
	if actual == expected {
		return true
	}

	if err := yaml.Unmarshal([]byte(actual), &actualYAMLAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Input ('%s') needs to be valid yaml.\nYAML error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return Equal(t, expectedYAMLAsInterface, actualYAMLAsInterface, msgAndArgs...)
}
