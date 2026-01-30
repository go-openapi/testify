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

func TestHTTPSuccess(t *testing.T) {
	t.Parallel()

	mock1 := new(testing.T)
	Equal(t, HTTPSuccess(mock1, httpOK, "GET", "/", nil), true)
	False(t, mock1.Failed())

	mock2 := new(testing.T)
	Equal(t, HTTPSuccess(mock2, httpRedirect, "GET", "/", nil), false)
	True(t, mock2.Failed())

	mock3 := new(mockT)
	Equal(t, HTTPSuccess(
		mock3, httpError, "GET", "/", nil,
		"was not expecting a failure here",
	), false)
	True(t, mock3.Failed())

	mock4 := new(testing.T)
	Equal(t, HTTPSuccess(mock4, httpStatusCode, "GET", "/", nil), false)
	True(t, mock4.Failed())

	mock5 := new(testing.T)
	Equal(t, HTTPSuccess(mock5, httpReadBody, "POST", "/", nil), true)
	False(t, mock5.Failed())
}

func TestHTTPRedirect(t *testing.T) {
	t.Parallel()
	mock1 := new(mockT)

	Equal(t, HTTPRedirect(
		mock1, httpOK, "GET", "/", nil,
		"was expecting a 3xx status code. Got 200.",
	), false)
	True(t, mock1.Failed())

	mock2 := new(testing.T)
	Equal(t, HTTPRedirect(mock2, httpRedirect, "GET", "/", nil), true)
	False(t, mock2.Failed())

	mock3 := new(testing.T)
	Equal(t, HTTPRedirect(mock3, httpError, "GET", "/", nil), false)
	True(t, mock3.Failed())

	mock4 := new(testing.T)
	Equal(t, HTTPRedirect(mock4, httpStatusCode, "GET", "/", nil), false)
	True(t, mock4.Failed())
}

func TestHTTPError(t *testing.T) {
	t.Parallel()

	mock1 := new(testing.T)
	Equal(t, HTTPError(mock1, httpOK, "GET", "/", nil), false)
	True(t, mock1.Failed())

	mock2 := new(mockT)
	Equal(t, HTTPError(
		mock2, httpRedirect, "GET", "/", nil,
		"Expected this request to error out. But it didn't",
	), false)
	True(t, mock2.Failed())

	mock3 := new(testing.T)
	Equal(t, HTTPError(mock3, httpError, "GET", "/", nil), true)
	False(t, mock3.Failed())

	mock4 := new(testing.T)
	Equal(t, HTTPError(mock4, httpStatusCode, "GET", "/", nil), false)
	True(t, mock4.Failed())
}

func TestHTTPStatusCode(t *testing.T) {
	t.Parallel()

	mock1 := new(testing.T)
	Equal(t, HTTPStatusCode(mock1, httpOK, "GET", "/", nil, http.StatusSwitchingProtocols), false)
	True(t, mock1.Failed())

	mock2 := new(testing.T)
	Equal(t, HTTPStatusCode(mock2, httpRedirect, "GET", "/", nil, http.StatusSwitchingProtocols), false)
	True(t, mock2.Failed())

	mock3 := new(mockT)
	Equal(t, HTTPStatusCode(
		mock3, httpError, "GET", "/", nil, http.StatusSwitchingProtocols,
		"Expected the status code to be %d", http.StatusSwitchingProtocols,
	), false)
	True(t, mock3.Failed())

	mock4 := new(testing.T)
	Equal(t, HTTPStatusCode(mock4, httpStatusCode, "GET", "/", nil, http.StatusSwitchingProtocols), true)
	False(t, mock4.Failed())
}

func TestHTTPStatusWrapper(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	Equal(t, HTTPSuccess(mock, httpOK, "GET", "/", nil), true)
	Equal(t, HTTPSuccess(mock, httpRedirect, "GET", "/", nil), false)
	Equal(t, HTTPSuccess(mock, httpError, "GET", "/", nil), false)

	Equal(t, HTTPRedirect(mock, httpOK, "GET", "/", nil), false)
	Equal(t, HTTPRedirect(mock, httpRedirect, "GET", "/", nil), true)
	Equal(t, HTTPRedirect(mock, httpError, "GET", "/", nil), false)

	Equal(t, HTTPError(mock, httpOK, "GET", "/", nil), false)
	Equal(t, HTTPError(mock, httpRedirect, "GET", "/", nil), false)
	Equal(t, HTTPError(mock, httpError, "GET", "/", nil), true)
}

func TestHTTPRequestWithNoParams(t *testing.T) {
	t.Parallel()

	var got *http.Request
	handler := func(w http.ResponseWriter, r *http.Request) {
		got = r
		w.WriteHeader(http.StatusOK)
	}

	True(t, HTTPSuccess(t, handler, "GET", "/url", nil))

	Empty(t, got.URL.Query())
	Equal(t, "/url", got.URL.RequestURI())
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

	True(t, HTTPSuccess(t, handler, "GET", "/url", params))

	Equal(t, url.Values{"id": []string{"12345"}}, got.URL.Query())
	Equal(t, "/url?id=12345", got.URL.String())
	Equal(t, "/url?id=12345", got.URL.RequestURI())
}

func TestHttpBody(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	True(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	True(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	False(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))

	False(t, HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	False(t, HTTPBodyNotContains(
		mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World",
		"Expected the request body to not contain 'World'. But it did.",
	))
	True(t, HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))
	True(t, HTTPBodyContains(mock, httpReadBody, "GET", "/", nil, "hello"))
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

	True(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	True(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	False(t, HTTPBodyContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))

	False(t, HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"))
	False(t, HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "World"))
	True(t, HTTPBodyNotContains(mock, httpHelloName, "GET", "/", url.Values{"name": []string{"World"}}, "world"))
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
