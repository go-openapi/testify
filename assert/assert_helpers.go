// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package assert

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// CallerInfo returns an array of strings containing the file and line number
// of each stack frame leading from the current test to the assert call that
// failed.
func CallerInfo() []string {
	return assertions.CallerInfo()
}

// HTTPBody is a helper that returns the HTTP body of the response.
// It returns the empty string if building a new request fails.
func HTTPBody(handler http.HandlerFunc, method string, url string, values url.Values) string {
	return assertions.HTTPBody(handler, method, url, values)
}

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected any, actual any) bool {
	return assertions.ObjectsAreEqual(expected, actual)
}

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected any, actual any) bool {
	return assertions.ObjectsAreEqualValues(expected, actual)
}

// WithHunkSize sets how many unchanged lines the diff shows around each change.
//
// The diff appears in the failure message of [Equal], [EqualT], [EqualValues],
// [EqualExportedValues] and [Exactly], whenever both values are a struct, map, slice, array
// or string. The default is 1. A value below 1 is clamped to 1, and a value larger than the
// rendered value prints it whole.
//
// Pass it to [New], which is the only place options are read:
//
//	a := assert.New(t, assert.WithHunkSize(4))
//	a.Equal(expected, actual)
func WithHunkSize(n int) Option {
	return assertions.WithHunkSize(n)
}
