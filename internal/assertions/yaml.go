// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"fmt"

	"github.com/go-openapi/testify/v2/internal/assertions/enable/yaml"
)

// YAMLEqBytes asserts that two YAML slices of bytes are equivalent.
//
// Expected and actual must be valid YAML.
//
// # Important
//
// By default, this function is disabled and will panic.
//
// To enable it, you should add a blank import like so:
//
//	import(
//	  "github.com/go-openapi/testify/enable/yaml/v2"
//	)
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
//	panic: []byte("key: value"), []byte("key: value")
//	should panic without the yaml feature enabled.
func YAMLEqBytes(t T, expected, actual []byte, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}
	var expectedYAMLAsInterface, actualYAMLAsInterface any

	if err := yaml.Unmarshal(expected, &expectedYAMLAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Expected value ('%s') is not valid yaml.\nYAML parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	// Shortcut if same bytes
	if bytes.Equal(actual, expected) {
		return true
	}

	if err := yaml.Unmarshal(actual, &actualYAMLAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Input ('%s') needs to be valid yaml.\nYAML error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return Equal(t, expectedYAMLAsInterface, actualYAMLAsInterface, msgAndArgs...)
}

// YAMLEq asserts that two YAML strings are equivalent.
//
// See [YAMLEqBytes].
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
func YAMLEq(t T, expected, actual string, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return YAMLEqBytes(t, []byte(expected), []byte(actual), msgAndArgs)
}

// YAMLEqT asserts that two YAML documents are equivalent.
//
// The expected and actual arguments may be string or []byte. They do not need to be of the same type.
//
// See [YAMLEqBytes].
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
func YAMLEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return YAMLEqBytes(t, []byte(expected), []byte(actual), msgAndArgs)
}
