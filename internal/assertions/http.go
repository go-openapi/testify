// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// HTTPSuccess asserts that a specified handler returns a success status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// # Examples
//
//	success: httpOK, "GET", "/", nil
//	failure: httpError, "GET", "/", nil
func HTTPSuccess(t T, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}

	err := httpCodeInRange(handler, http.StatusOK, http.StatusPartialContent, "HTTP success", method, url, values)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return true
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// # Examples
//
//	success: httpRedirect, "GET", "/", nil
//	failure: httpError, "GET", "/", nil
func HTTPRedirect(t T, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}

	err := httpCodeInRange(handler, http.StatusMultipleChoices, http.StatusTemporaryRedirect, "HTTP redirect", method, url, values)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return true
}

// HTTPError asserts that a specified handler returns an error status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// # Examples
//
//	success: httpError, "GET", "/", nil
//	failure: httpOK, "GET", "/", nil
func HTTPError(t T, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}

	err := httpCodeInRange(handler, http.StatusBadRequest, -1, "HTTP error", method, url, values)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return true
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// # Examples
//
//	success: httpOK, "GET", "/", nil, http.StatusOK
//	failure: httpError, "GET", "/", nil, http.StatusOK
func HTTPStatusCode(t T, handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}

	err := httpCodeInRange(handler, statuscode, statuscode, fmt.Sprintf("HTTP status code %d", statuscode), method, url, values)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return true
}

// HTTPBody is a helper that returns the HTTP body of the response.
// It returns the empty string if building a new request fails.
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string {
	// Domain: http
	w := httptest.NewRecorder()
	if len(values) > 0 {
		url += "?" + values.Encode()
	}
	req, err := http.NewRequestWithContext(context.Background(), method, url, http.NoBody)
	if err != nil {
		return ""
	}
	handler(w, req)
	return w.Body.String()
}

// HTTPBodyContains asserts that a specified handler returns a body that contains a string.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// # Examples
//
//	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"
//	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!"
func HTTPBodyContains(t T, handler http.HandlerFunc, method, url string, values url.Values, str any, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if !contains {
		Fail(t, fmt.Sprintf("Expected response body for %q to contain %q but found %q", url+"?"+values.Encode(), str, body), msgAndArgs...)
	}

	return contains
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// # Examples
//
//	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!"
//	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!"
func HTTPBodyNotContains(t T, handler http.HandlerFunc, method, url string, values url.Values, str any, msgAndArgs ...any) bool {
	// Domain: http
	if h, ok := t.(H); ok {
		h.Helper()
	}
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if contains {
		Fail(t, fmt.Sprintf("Expected response body for %q to NOT contain %q but found %q", url+"?"+values.Encode(), str, body), msgAndArgs...)
	}

	return !contains
}

func httpCodeInRange(handler http.HandlerFunc, minStatusCode, maxStatusCode int, expectedStatus, method, url string, values url.Values) error {
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		return fmt.Errorf("failed to build test request, got error: %w", err)
	}

	if (code < minStatusCode && minStatusCode != -1) || (code > maxStatusCode && maxStatusCode != -1) {
		return fmt.Errorf("expected %s status code for %q but received %d", expectedStatus, url+"?"+values.Encode(), code)
	}

	return nil
}

// httpCode is a helper that returns the HTTP code of the response.
//
// It returns -1 and an error if building a new request fails.
func httpCode(handler http.HandlerFunc, method, url string, values url.Values) (int, error) {
	// maintainer: should inject t.Context()
	w := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(context.Background(), method, url, http.NoBody)
	if err != nil {
		return -1, err
	}

	req.URL.RawQuery = values.Encode()
	handler(w, req)

	return w.Code, nil
}
