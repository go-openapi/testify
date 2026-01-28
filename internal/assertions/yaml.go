// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // YAML is actually very similar to JSON but we can't easily factorize this.
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

// YAMLUnmarshalAsT wraps [Equal] after [yaml.Unmarshal].
//
// The input YAML may be a string or []byte.
//
// It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.
//
// Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
// the unmarshaling would take this type to unmarshal as a map[string]any.
//
// # Usage
//
//	expected := struct {
//		A int `yaml:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
func YAMLUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var actual Object
	if err := yaml.Unmarshal([]byte(jazon), &actual); err != nil {
		return Fail(t, fmt.Sprintf("YAML unmarshal failed: %v", err), msgAndArgs...)
	}

	return Equal(t, expected, actual, msgAndArgs...)
}

// YAMLMarshalAsT wraps [YAMLEq] after [yaml.Marshal].
//
// The input YAML may be a string or []byte.
//
// It fails if the marshaling returns an error or if the expected YAML bytes differ semantically
// from the expected ones.
//
// # Usage
//
//	actual := struct {
//		A int `yaml:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
func YAMLMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool {
	// Domain: yaml
	if h, ok := t.(H); ok {
		h.Helper()
	}

	actual, err := yaml.Marshal(object)
	if err != nil {
		return Fail(t, fmt.Sprintf("YAML marshal failed: %v", err), msgAndArgs...)
	}

	return YAMLEqBytes(t, []byte(expected), actual, msgAndArgs...)
}
