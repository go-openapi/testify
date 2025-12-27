// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONEqBytes asserts that two JSON byte slices are equivalent.
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
	return JSONEqBytes(t, []byte(expected), []byte(actual), msgAndArgs)
}
