// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen; DO NOT EDIT.
// Generated on 2026-01-02 (version v1.2.2-760-g97c29e3) using codegen version master [sha: 97c29e3dbfc40800a080863ceea81db0cfd6e858]

package require

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
