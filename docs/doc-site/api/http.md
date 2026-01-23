---
title: "Http"
description: "Asserting HTTP Response And Body"
modified: 2026-01-24
weight: 8
domains:
  - "http"
keywords:
  - "HTTPBody"
  - "HTTPBodyf"
  - "HTTPBodyContains"
  - "HTTPBodyContainsf"
  - "HTTPBodyNotContains"
  - "HTTPBodyNotContainsf"
  - "HTTPError"
  - "HTTPErrorf"
  - "HTTPRedirect"
  - "HTTPRedirectf"
  - "HTTPStatusCode"
  - "HTTPStatusCodef"
  - "HTTPSuccess"
  - "HTTPSuccessf"
---

Asserting HTTP Response And Body

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 7 functionalities.

```tree
- [HTTPBodyContains](#httpbodycontains) | angles-right
- [HTTPBodyNotContains](#httpbodynotcontains) | angles-right
- [HTTPError](#httperror) | angles-right
- [HTTPRedirect](#httpredirect) | angles-right
- [HTTPStatusCode](#httpstatuscode) | angles-right
- [HTTPSuccess](#httpsuccess) | angles-right
```

### HTTPBodyContains{#httpbodycontains}

HTTPBodyContains asserts that a specified handler returns a body that contains a string.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"
	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPBodyContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPBodyContains) | package-level function |
| [`assert.HTTPBodyContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPBodyContainsf) | formatted variant |
| [`assert.(*Assertions).HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPBodyContains) | method variant |
| [`assert.(*Assertions).HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPBodyContainsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPBodyContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPBodyContains) | package-level function |
| [`require.HTTPBodyContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPBodyContainsf) | formatted variant |
| [`require.(*Assertions).HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPBodyContains) | method variant |
| [`require.(*Assertions).HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPBodyContainsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPBodyContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPBodyContains) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPBodyContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L163)
{{% /tab %}}
{{< /tabs >}}

### HTTPBodyNotContains{#httpbodynotcontains}

HTTPBodyNotContains asserts that a specified handler returns a
body that does not contain a string.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!"
	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPBodyNotContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPBodyNotContains) | package-level function |
| [`assert.HTTPBodyNotContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPBodyNotContainsf) | formatted variant |
| [`assert.(*Assertions).HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPBodyNotContains) | method variant |
| [`assert.(*Assertions).HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPBodyNotContainsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPBodyNotContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPBodyNotContains) | package-level function |
| [`require.HTTPBodyNotContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPBodyNotContainsf) | formatted variant |
| [`require.(*Assertions).HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPBodyNotContains) | method variant |
| [`require.(*Assertions).HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPBodyNotContainsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPBodyNotContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPBodyNotContains) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPBodyNotContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L191)
{{% /tab %}}
{{< /tabs >}}

### HTTPError{#httperror}

HTTPError asserts that a specified handler returns an error status code.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpError, "GET", "/", nil
	failure: httpOK, "GET", "/", nil
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPError(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPError) | package-level function |
| [`assert.HTTPErrorf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPErrorf) | formatted variant |
| [`assert.(*Assertions).HTTPError(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPError) | method variant |
| [`assert.(*Assertions).HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPErrorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPError(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPError) | package-level function |
| [`require.HTTPErrorf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPErrorf) | formatted variant |
| [`require.(*Assertions).HTTPError(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPError) | method variant |
| [`require.(*Assertions).HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPErrorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPError(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPError) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPError](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L87)
{{% /tab %}}
{{< /tabs >}}

### HTTPRedirect{#httpredirect}

HTTPRedirect asserts that a specified handler returns a redirect status code.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpRedirect, "GET", "/", nil
	failure: httpError, "GET", "/", nil
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPRedirect(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPRedirect) | package-level function |
| [`assert.HTTPRedirectf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPRedirectf) | formatted variant |
| [`assert.(*Assertions).HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPRedirect) | method variant |
| [`assert.(*Assertions).HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPRedirectf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPRedirect(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPRedirect) | package-level function |
| [`require.HTTPRedirectf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPRedirectf) | formatted variant |
| [`require.(*Assertions).HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPRedirect) | method variant |
| [`require.(*Assertions).HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPRedirectf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPRedirect(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPRedirect) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPRedirect](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L57)
{{% /tab %}}
{{< /tabs >}}

### HTTPStatusCode{#httpstatuscode}

HTTPStatusCode asserts that a specified handler returns a specified status code.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpOK, "GET", "/", nil, http.StatusOK
	failure: httpError, "GET", "/", nil, http.StatusOK
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPStatusCode(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPStatusCode) | package-level function |
| [`assert.HTTPStatusCodef(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPStatusCodef) | formatted variant |
| [`assert.(*Assertions).HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPStatusCode) | method variant |
| [`assert.(*Assertions).HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPStatusCodef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPStatusCode(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPStatusCode) | package-level function |
| [`require.HTTPStatusCodef(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPStatusCodef) | formatted variant |
| [`require.(*Assertions).HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPStatusCode) | method variant |
| [`require.(*Assertions).HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPStatusCodef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPStatusCode(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPStatusCode) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPStatusCode](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L117)
{{% /tab %}}
{{< /tabs >}}

### HTTPSuccess{#httpsuccess}

HTTPSuccess asserts that a specified handler returns a success status code.

Returns whether the assertion was successful (true) or not (false).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: httpOK, "GET", "/", nil
	failure: httpError, "GET", "/", nil
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPSuccess(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPSuccess) | package-level function |
| [`assert.HTTPSuccessf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPSuccessf) | formatted variant |
| [`assert.(*Assertions).HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPSuccess) | method variant |
| [`assert.(*Assertions).HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.HTTPSuccessf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPSuccess(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPSuccess) | package-level function |
| [`require.HTTPSuccessf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPSuccessf) | formatted variant |
| [`require.(*Assertions).HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPSuccess) | method variant |
| [`require.(*Assertions).HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.HTTPSuccessf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPSuccess(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPSuccess) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPSuccess](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L27)
{{% /tab %}}
{{< /tabs >}}

---

## Other helpers

### HTTPBody

HTTPBody is a helper that returns the HTTP body of the response.
It returns the empty string if building a new request fails.


{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.HTTPBody(handler http.HandlerFunc, method string, url string, values url.Values) string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#HTTPBody) | package-level function |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.HTTPBody(handler http.HandlerFunc, method string, url string, values url.Values) string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#HTTPBody) | package-level function |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.HTTPBody(handler http.HandlerFunc, method string, url string, values url.Values) string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#HTTPBody) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#HTTPBody](https://github.com/go-openapi/testify/blob/master/internal/assertions/http.go#L137)
{{% /tab %}}
{{< /tabs >}}

---

Generated with github.com/go-openapi/testify/codegen/v2

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/codegen/v2 DO NOT EDIT.

Generated on 2026-01-24 (version 178304f) using codegen version v2.1.9-0.20260123222731-178304f36678+dirty [sha: 178304f366789315d4db6b11c89786c43d916247]
-->
