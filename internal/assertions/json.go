// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// JSONEqBytes asserts that two JSON slices of bytes are semantically equivalent.
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
	// Maintainers: Proposal for enhancement.
	// We could use and indirection for users to inject their favorite JSON
	// library like we do for YAML.
	if h, ok := t.(H); ok {
		h.Helper()
	}
	var expectedJSONAsInterface, actualJSONAsInterface any

	if err := json.Unmarshal(expected, &expectedJSONAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Expected value (%q) is not valid json.\nJSON parsing error: %v", expected, err), msgAndArgs...)
	}

	// Shortcut if same bytes
	if bytes.Equal(actual, expected) {
		return true
	}

	if err := json.Unmarshal(actual, &actualJSONAsInterface); err != nil {
		return Fail(t, fmt.Sprintf("Input (%q) needs to be valid json.\nJSON parsing error: %v", actual, err), msgAndArgs...)
	}

	return Equal(t, expectedJSONAsInterface, actualJSONAsInterface, msgAndArgs...)
}

// JSONEq asserts that two JSON strings are semantically equivalent.
//
// Expected and actual must be valid JSON.
//
// For dynamic redaction of the input text via a callback, use [JSONEqT].
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

	return JSONEqBytes(t, []byte(expected), []byte(actual), msgAndArgs...)
}

// JSONEqT asserts that two JSON documents are semantically equivalent.
//
// The expected and actual arguments may be string or []byte. They do not need to be of the same type.
//
// Expected and actual must be valid JSON.
//
// NOTE: passed values (expected, actual) may be wrapped as functions to redact the input text dynamically.
//
// # Usage
//
//	assertions.JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
//
// # Examples
//
//	success: `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`)
//	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
func JSONEqT[EDoc, ADoc RText](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool {
	// Domain: json
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return JSONEqBytes(t, asBytes(expected), asBytes(actual), msgAndArgs...)
}

// JSONUnmarshalAsT wraps [Equal] after [json.Unmarshal].
//
// The input JSON may be a string or []byte.
//
// It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.
//
// Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
// the unmarshaling would take this type to unmarshal as a map[string]any.
//
// NOTE: passed jazon value may be wrapped as a function to redact the input JSON dynamically.
//
// # Usage
//
//	expected := struct {
//		A int `json:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	success: dummyStruct{A: "a"} , []byte(`{"A": "a"}`)
//	failure: 1, `[{"foo": "bar"}, {"hello": "world"}]`
func JSONUnmarshalAsT[Object any, ADoc RText](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool {
	// Domain: json
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var actual Object
	if err := json.Unmarshal(asBytes(jazon), &actual); err != nil {
		return Fail(t, fmt.Sprintf("JSON unmarshal failed: %v", err), msgAndArgs...)
	}

	return Equal(t, expected, actual, msgAndArgs...)
}

// JSONMarshalAsT wraps [JSONEqT] after [json.Marshal].
//
// The input JSON may be a string or []byte.
//
// It fails if the marshaling returns an error or if the expected JSON bytes differ semantically
// from the expected ones.
//
// NOTE: passed expected value may be wrapped as a function to redact the input text dynamically.
//
// # Usage
//
//	actual := struct {
//		A int `json:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	success: []byte(`{"A": "a"}`), dummyStruct{A: "a"}
//	failure: `[{"foo": "bar"}, {"hello": "world"}]`, 1
func JSONMarshalAsT[EDoc RText](t T, expected EDoc, object any, msgAndArgs ...any) bool {
	// Domain: json
	if h, ok := t.(H); ok {
		h.Helper()
	}

	actual, err := json.Marshal(object)
	if err != nil {
		return Fail(t, fmt.Sprintf("JSON marshal failed: %v", err), msgAndArgs...)
	}

	return JSONEqBytes(t, asBytes(expected), actual, msgAndArgs...)
}

func asBytes[EDoc RText](e EDoc) []byte {
	ie := any(e)

	switch typed := ie.(type) {
	case func() string:
		if typed == nil {
			panic("passed Redactor cannot be nil")
		}
		return []byte(typed())
	case func() []byte:
		if typed == nil {
			panic("passed Redactor cannot be nil")
		}
		return typed()
	case string:
		return []byte(typed)
	case []byte:
		return typed
	default:
		// this edge case (redefined type) requires the input to be converted: the type constraint warrants it to work
		return convertReflectValue[[]byte](e, reflect.ValueOf(e))
	}
}
