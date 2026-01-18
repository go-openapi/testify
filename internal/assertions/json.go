// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // we need to duplicate at least some godoc.
package assertions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONEqBytes asserts that two JSON slices of bytes are equivalent.
//
// Expected and actual must be valid JSON.
//
// # Usage
//
//	assertions.JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
//
// # Examples
//
//	success: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`)
//	failure: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`)
func JSONEqBytes(t T, expected, actual []byte, msgAndArgs ...any) bool {
	// Domain: json
	// Maintainer: proposal for enhancement. We could use and indirection for users to inject their favorite JSON
	// library like we do for YAML.
	if h, ok := t.(H); ok {
		h.Helper()
	}
	var expectedJSONAsInterface, actualJSONAsInterface any

	if err := json.Unmarshal(expected, &expectedJSONAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Expected value ('%s') is not valid json.\nJSON parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	// Shortcut if same bytes
	if bytes.Equal(actual, expected) {
		return true
	}

	if err := json.Unmarshal(actual, &actualJSONAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return Equal(t, expectedJSONAsInterface, actualJSONAsInterface, msgAndArgs...)
}

// JSONEq asserts that two JSON strings are equivalent.
//
// Expected and actual must be valid JSON.
//
// # Usage
//
//	assertions.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
//
// # Examples
//
//	success: `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`
//	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
func JSONEq(t T, expected, actual string, msgAndArgs ...any) bool {
	// Domain: json
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return JSONEqBytes(t, []byte(expected), []byte(actual), msgAndArgs)
}

// JSONEqT asserts that two JSON documents are equivalent.
//
// The expected and actual arguments may be string or []byte. They do not need to be of the same type.
//
// Expected and actual must be valid JSON.
//
// # Usage
//
//	assertions.JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
//
// # Examples
//
//	success: `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`)
//	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
func JSONEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool {
	// Domain: json
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return JSONEqBytes(t, []byte(expected), []byte(actual), msgAndArgs)
}
