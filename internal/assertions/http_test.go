// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
	"slices"
	"testing"
)

// TestHTTPStatus runs all HTTP status assertions against a unified set of handlers.
//
// The expected pass/fail is determined by [expectedHTTPStatus].
func TestHTTPStatus(t *testing.T) {
	t.Parallel()

	for tc := range httpStatusCases() {
		t.Run(tc.name, testAllHTTPStatusAssertions(tc))
	}
}

// TestHTTPBody runs all HTTP body assertions against a unified set of cases.
//
// The expected pass/fail is determined by [expectedHTTPBody].
func TestHTTPBody(t *testing.T) {
	t.Parallel()

	for tc := range httpBodyCases() {
		t.Run(tc.name, testAllHTTPBodyAssertions(tc))
	}
}

func TestHTTPErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, httpFailCases())
}

// ============================================================================
// Unified HTTP status assertion tests (truth matrix pattern)
// ============================================================================

// handlerKind represents the HTTP response category of a handler.
type handlerKind int

const (
	successHandler    handlerKind = iota // 2xx
	redirectHandler                      // 3xx
	errorHandler                         // 4xx/5xx
	otherHandler                         // 1xx or other non-standard
	badRequestHandler                    // request construction fails (e.g. invalid method)
)

type httpAssertionKind int

const (
	httpSuccessAssertion httpAssertionKind = iota
	httpRedirectAssertion
	httpErrorAssertion
)

func testAllHTTPStatusAssertions(tc httpStatusTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("HTTPSuccess", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPSuccess(mock, tc.handler, tc.method, "/", tc.params)
			shouldPassOrFail(t, mock, res, expectedHTTPStatus(httpSuccessAssertion, tc.kind))
		})

		t.Run("HTTPRedirect", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPRedirect(mock, tc.handler, tc.method, "/", tc.params)
			shouldPassOrFail(t, mock, res, expectedHTTPStatus(httpRedirectAssertion, tc.kind))
		})

		t.Run("HTTPError", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPError(mock, tc.handler, tc.method, "/", tc.params)
			shouldPassOrFail(t, mock, res, expectedHTTPStatus(httpErrorAssertion, tc.kind))
		})

		t.Run("HTTPStatusCode/match", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPStatusCode(mock, tc.handler, tc.method, "/", tc.params, tc.code)
			shouldPassOrFail(t, mock, res, tc.kind != badRequestHandler)
		})

		t.Run("HTTPStatusCode/mismatch", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPStatusCode(mock, tc.handler, tc.method, "/", tc.params, 0)
			shouldPassOrFail(t, mock, res, false)
		})
	}
}

// ============================================================================
// Unified HTTP body assertion tests (truth matrix pattern)
// ============================================================================

// bodyMatchKind represents whether a string is found in the response body.
type bodyMatchKind int

const (
	bodyMatches bodyMatchKind = iota // string is found in response body
	bodyMisses                       // string is not found in response body
)

type httpBodyAssertionKind int

const (
	httpBodyContainsAssertion httpBodyAssertionKind = iota
	httpBodyNotContainsAssertion
)

func testAllHTTPBodyAssertions(tc httpBodyTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("HTTPBodyContains", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPBodyContains(mock, tc.handler, tc.method, "/", tc.params, tc.str)
			shouldPassOrFail(t, mock, res, expectedHTTPBody(httpBodyContainsAssertion, tc.kind))
		})

		t.Run("HTTPBodyNotContains", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := HTTPBodyNotContains(mock, tc.handler, tc.method, "/", tc.params, tc.str)
			shouldPassOrFail(t, mock, res, expectedHTTPBody(httpBodyNotContainsAssertion, tc.kind))
		})
	}
}

// ============================================================================
// HTTP status test cases
// ============================================================================

type httpStatusTestCase struct {
	name    string
	handler http.HandlerFunc
	kind    handlerKind
	code    int        // actual HTTP status code returned by the handler
	method  string     // HTTP method for the request
	params  url.Values // optional query params (nil for none)
}

func httpStatusCases() iter.Seq[httpStatusTestCase] {
	return slices.Values([]httpStatusTestCase{
		{"ok", httpOK, successHandler, http.StatusOK, "GET", nil},
		{"ok/with-params", httpOK, successHandler, http.StatusOK, "GET", url.Values{"id": {"12345"}}},
		{"redirect", httpRedirect, redirectHandler, http.StatusTemporaryRedirect, "GET", nil},
		{"error", httpError, errorHandler, http.StatusInternalServerError, "GET", nil},
		{"status-code-1xx", httpStatusCode, otherHandler, http.StatusSwitchingProtocols, "GET", nil},
		{"read-body", httpReadBody, successHandler, http.StatusOK, "GET", nil},
		{"bad-method", httpOK, badRequestHandler, http.StatusOK, "BAD METHOD", nil},
	})
}

// expectedHTTPStatus determines the expected pass/fail for each assertion based on handler kind.
func expectedHTTPStatus(assertion httpAssertionKind, kind handlerKind) bool {
	switch assertion {
	case httpSuccessAssertion:
		return kind == successHandler
	case httpRedirectAssertion:
		return kind == redirectHandler
	case httpErrorAssertion:
		return kind == errorHandler
	default:
		panic(fmt.Errorf("test case configuration error: invalid httpAssertionKind: %d", assertion))
	}
}

// ============================================================================
// HTTP Body tests
// ============================================================================

type httpBodyTestCase struct {
	name    string
	handler http.HandlerFunc
	method  string
	params  url.Values
	str     string
	kind    bodyMatchKind
}

func httpBodyCases() iter.Seq[httpBodyTestCase] {
	p := url.Values{"name": []string{"World"}}

	return slices.Values([]httpBodyTestCase{
		{"full-match", httpHelloName, "GET", p, "Hello, World!", bodyMatches},
		{"partial-match", httpHelloName, "GET", p, "World", bodyMatches},
		{"case-sensitive-miss", httpHelloName, "GET", p, "world", bodyMisses},
		{"read-body-handler", httpReadBody, "GET", nil, "hello", bodyMatches},
		{"bad-method", httpHelloName, "BAD METHOD", nil, "Hello", bodyMisses},
	})
}

// expectedHTTPBody determines the expected pass/fail for each assertion based on body match kind.
func expectedHTTPBody(assertion httpBodyAssertionKind, kind bodyMatchKind) bool {
	switch assertion {
	case httpBodyContainsAssertion:
		return kind == bodyMatches
	case httpBodyNotContainsAssertion:
		return kind == bodyMisses
	default:
		panic(fmt.Errorf("test case configuration error: invalid httpBodyAssertionKind: %d", assertion))
	}
}

// ============================================================================
// Error message tests
// ============================================================================

func httpFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "HTTPSuccess/error-handler",
			assertion:    func(t T) bool { return HTTPSuccess(t, httpError, "GET", "/", nil) },
			wantContains: []string{"expected HTTP success status code"},
		},
		{
			name:         "HTTPRedirect/ok-handler",
			assertion:    func(t T) bool { return HTTPRedirect(t, httpOK, "GET", "/", nil) },
			wantContains: []string{"expected HTTP redirect status code"},
		},
		{
			name:         "HTTPError/redirect-handler",
			assertion:    func(t T) bool { return HTTPError(t, httpRedirect, "GET", "/", nil) },
			wantContains: []string{"expected HTTP error status code"},
		},
	})
}

// ============================================================================
// Test HTTP handlers
// ============================================================================

func httpHelloName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_, _ = fmt.Fprintf(w, "Hello, %s!", name)
}

func httpOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func httpReadBody(w http.ResponseWriter, r *http.Request) {
	_, _ = io.Copy(io.Discard, r.Body)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
}

func httpRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func httpError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func httpStatusCode(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusSwitchingProtocols)
}
