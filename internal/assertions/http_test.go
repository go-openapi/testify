// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"testing"
)

func TestHTTPSuccess(t *testing.T) {
	t.Parallel()

	mock1 := new(mockT)
	if result := HTTPSuccess(mock1, httpOK, "GET", "/", nil); !result {
		t.Error("expected HTTPSuccess to return true for httpOK")
	}
	if mock1.Failed() {
		t.Error("expected mock not to have failed")
	}

	mock2 := new(mockT)
	if result := HTTPSuccess(mock2, httpRedirect, "GET", "/", nil); result {
		t.Error("expected HTTPSuccess to return false for httpRedirect")
	}
	if !mock2.Failed() {
		t.Error("expected mock to have failed")
	}

	mock3 := new(mockT)
	if result := HTTPSuccess(
		mock3, httpError, "GET", "/", nil,
		"was not expecting a failure here",
	); result {
		t.Error("expected HTTPSuccess to return false for httpError")
	}
	if !mock3.Failed() {
		t.Error("expected mock to have failed")
	}

	mock4 := new(mockT)
	if result := HTTPSuccess(mock4, httpStatusCode, "GET", "/", nil); result {
		t.Error("expected HTTPSuccess to return false for httpStatusCode")
	}
	if !mock4.Failed() {
		t.Error("expected mock to have failed")
	}

	mock5 := new(mockT)
	if result := HTTPSuccess(mock5, httpReadBody, "POST", "/", nil); !result {
		t.Error("expected HTTPSuccess to return true for httpReadBody")
	}
	if mock5.Failed() {
		t.Error("expected mock not to have failed")
	}
}

func TestHTTPRedirect(t *testing.T) {
	t.Parallel()

	mock1 := new(mockT)
	if result := HTTPRedirect(
		mock1, httpOK, "GET", "/", nil,
		"was expecting a 3xx status code. Got 200.",
	); result {
		t.Error("expected HTTPRedirect to return false for httpOK")
	}
	if !mock1.Failed() {
		t.Error("expected mock to have failed")
	}

	mock2 := new(mockT)
	if result := HTTPRedirect(mock2, httpRedirect, "GET", "/", nil); !result {
		t.Error("expected HTTPRedirect to return true for httpRedirect")
	}
	if mock2.Failed() {
		t.Error("expected mock not to have failed")
	}

	mock3 := new(mockT)
	if result := HTTPRedirect(mock3, httpError, "GET", "/", nil); result {
		t.Error("expected HTTPRedirect to return false for httpError")
	}
	if !mock3.Failed() {
		t.Error("expected mock to have failed")
	}

	mock4 := new(mockT)
	if result := HTTPRedirect(mock4, httpStatusCode, "GET", "/", nil); result {
		t.Error("expected HTTPRedirect to return false for httpStatusCode")
	}
	if !mock4.Failed() {
		t.Error("expected mock to have failed")
	}
}

func TestHTTPError(t *testing.T) {
	t.Parallel()

	mock1 := new(mockT)
	if result := HTTPError(mock1, httpOK, "GET", "/", nil); result {
		t.Error("expected HTTPError to return false for httpOK")
	}
	if !mock1.Failed() {
		t.Error("expected mock to have failed")
	}

	mock2 := new(mockT)
	if result := HTTPError(
		mock2, httpRedirect, "GET", "/", nil,
		"Expected this request to error out. But it didn't",
	); result {
		t.Error("expected HTTPError to return false for httpRedirect")
	}
	if !mock2.Failed() {
		t.Error("expected mock to have failed")
	}

	mock3 := new(mockT)
	if result := HTTPError(mock3, httpError, "GET", "/", nil); !result {
		t.Error("expected HTTPError to return true for httpError")
	}
	if mock3.Failed() {
		t.Error("expected mock not to have failed")
	}

	mock4 := new(mockT)
	if result := HTTPError(mock4, httpStatusCode, "GET", "/", nil); result {
		t.Error("expected HTTPError to return false for httpStatusCode")
	}
	if !mock4.Failed() {
		t.Error("expected mock to have failed")
	}
}

func TestHTTPStatusCode(t *testing.T) {
	t.Parallel()

	mock1 := new(mockT)
	if result := HTTPStatusCode(mock1, httpOK, "GET", "/", nil, http.StatusSwitchingProtocols); result {
		t.Error("expected HTTPStatusCode to return false for httpOK")
	}
	if !mock1.Failed() {
		t.Error("expected mock to have failed")
	}

	mock2 := new(mockT)
	if result := HTTPStatusCode(mock2, httpRedirect, "GET", "/", nil, http.StatusSwitchingProtocols); result {
		t.Error("expected HTTPStatusCode to return false for httpRedirect")
	}
	if !mock2.Failed() {
		t.Error("expected mock to have failed")
	}

	mock3 := new(mockT)
	if result := HTTPStatusCode(
		mock3, httpError, "GET", "/", nil, http.StatusSwitchingProtocols,
		"Expected the status code to be %d", http.StatusSwitchingProtocols,
	); result {
		t.Error("expected HTTPStatusCode to return false for httpError")
	}
	if !mock3.Failed() {
		t.Error("expected mock to have failed")
	}

	mock4 := new(mockT)
	if result := HTTPStatusCode(mock4, httpStatusCode, "GET", "/", nil, http.StatusSwitchingProtocols); !result {
		t.Error("expected HTTPStatusCode to return true for httpStatusCode")
	}
	if mock4.Failed() {
		t.Error("expected mock not to have failed")
	}
}

func TestHTTPStatusWrapper(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !HTTPSuccess(mock, httpOK, "GET", "/", nil) {
		t.Error("expected HTTPSuccess(httpOK) to return true")
	}
	if HTTPSuccess(mock, httpRedirect, "GET", "/", nil) {
		t.Error("expected HTTPSuccess(httpRedirect) to return false")
	}
	if HTTPSuccess(mock, httpError, "GET", "/", nil) {
		t.Error("expected HTTPSuccess(httpError) to return false")
	}

	if HTTPRedirect(mock, httpOK, "GET", "/", nil) {
		t.Error("expected HTTPRedirect(httpOK) to return false")
	}
	if !HTTPRedirect(mock, httpRedirect, "GET", "/", nil) {
		t.Error("expected HTTPRedirect(httpRedirect) to return true")
	}
	if HTTPRedirect(mock, httpError, "GET", "/", nil) {
		t.Error("expected HTTPRedirect(httpError) to return false")
	}

	if HTTPError(mock, httpOK, "GET", "/", nil) {
		t.Error("expected HTTPError(httpOK) to return false")
	}
	if HTTPError(mock, httpRedirect, "GET", "/", nil) {
		t.Error("expected HTTPError(httpRedirect) to return false")
	}
	if !HTTPError(mock, httpError, "GET", "/", nil) {
		t.Error("expected HTTPError(httpError) to return true")
	}
}

func TestHTTPRequestWithNoParams(t *testing.T) {
	t.Parallel()

	var got *http.Request
	handler := func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.WriteHeader(http.StatusOK)
	}

	if !HTTPSuccess(t, handler, "GET", "/url", nil) {
		t.Error("expected HTTPSuccess to return true")
	}

	if len(got.URL.Query()) != 0 {
		t.Errorf("expected empty query, got %v", got.URL.Query())
	}
	if got.URL.RequestURI() != "/url" {
		t.Errorf("expected RequestURI %q, got %q", "/url", got.URL.RequestURI())
	}
}

func TestHTTPRequestWithParams(t *testing.T) {
	t.Parallel()

	var got *http.Request
	handler := func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.WriteHeader(http.StatusOK)
	}
	params := url.Values{}
	params.Add("id", "12345")

	if !HTTPSuccess(t, handler, "GET", "/url", params) {
		t.Error("expected HTTPSuccess to return true")
	}

	expectedQuery := url.Values{"id": []string{"12345"}}
	if !reflect.DeepEqual(expectedQuery, got.URL.Query()) {
		t.Errorf("expected query %v, got %v", expectedQuery, got.URL.Query())
	}
	if got.URL.String() != "/url?id=12345" {
		t.Errorf("expected URL string %q, got %q", "/url?id=12345", got.URL.String())
	}
	if got.URL.RequestURI() != "/url?id=12345" {
		t.Errorf("expected RequestURI %q, got %q", "/url?id=12345", got.URL.RequestURI())
	}
}

func TestHttpBody(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!") {
		t.Error("expected HTTPBodyContains to return true for 'Hello, World!'")
	}
	if !HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World") {
		t.Error("expected HTTPBodyContains to return true for 'World'")
	}
	if HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world") {
		t.Error("expected HTTPBodyContains to return false for 'world' (case sensitive)")
	}

	if HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!") {
		t.Error("expected HTTPBodyNotContains to return false for 'Hello, World!'")
	}
	if HTTPBodyNotContains(
		mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World",
		"Expected the request body to not contain 'World'. But it did.",
	) {
		t.Error("expected HTTPBodyNotContains to return false for 'World'")
	}
	if !HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world") {
		t.Error("expected HTTPBodyNotContains to return true for 'world' (case sensitive)")
	}
	if !HTTPBodyContains(mock, httpReadBody, "GET", "/", nil, "hello") {
		t.Error("expected HTTPBodyContains to return true for httpReadBody 'hello'")
	}
}

func TestHTTPErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, httpFailCases())
}

// ============================================================================
// TestHTTPErrorMessages
// ============================================================================.
func httpFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "HTTPSuccess/error-handler",
			assertion:    func(t T) bool { return HTTPSuccess(t, httpError, "GET", "/", nil) },
			wantContains: []string{"Expected HTTP success status code"},
		},
		{
			name:         "HTTPRedirect/ok-handler",
			assertion:    func(t T) bool { return HTTPRedirect(t, httpOK, "GET", "/", nil) },
			wantContains: []string{"Expected HTTP redirect status code"},
		},
		{
			name:         "HTTPError/redirect-handler",
			assertion:    func(t T) bool { return HTTPError(t, httpRedirect, "GET", "/", nil) },
			wantContains: []string{"Expected HTTP error status code"},
		},
	})
}

func TestHTTPBodyWrappers(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!") {
		t.Error("expected HTTPBodyContains to return true for 'Hello, World!'")
	}
	if !HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World") {
		t.Error("expected HTTPBodyContains to return true for 'World'")
	}
	if HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world") {
		t.Error("expected HTTPBodyContains to return false for 'world'")
	}

	if HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!") {
		t.Error("expected HTTPBodyNotContains to return false for 'Hello, World!'")
	}
	if HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World") {
		t.Error("expected HTTPBodyNotContains to return false for 'World'")
	}
	if !HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world") {
		t.Error("expected HTTPBodyNotContains to return true for 'world'")
	}
}

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
