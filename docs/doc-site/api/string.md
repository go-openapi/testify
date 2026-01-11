---
title: "String"
description: "Asserting Strings"
modified: 2026-01-11
weight: 13
domains:
  - "string"
keywords:
  - "NotRegexp"
  - "NotRegexpf"
  - "Regexp"
  - "Regexpf"
---

Asserting Strings

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 2 functionalities.

### NotRegexp

NotRegexp asserts that a specified regexp does not match a string.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
	assertions.NotRegexp(t, "^start", "it's not starting")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: "^start", "not starting"
	failure: "^start", "starting"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotRegexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexp) | package-level function |
| [`assert.NotRegexpf(t T, rx any, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexpf) | formatted variant |
| [`assert.(*Assertions).NotRegexp(rx any, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotRegexp) | method variant |
| [`assert.(*Assertions).NotRegexpf(rx any, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotRegexpf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotRegexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexp) | package-level function |
| [`require.NotRegexpf(t T, rx any, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexpf) | formatted variant |
| [`require.(*Assertions).NotRegexp(rx any, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotRegexp) | method variant |
| [`require.(*Assertions).NotRegexpf(rx any, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotRegexpf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotRegexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotRegexp) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotRegexp](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L53)
{{% /tab %}}
{{< /tabs >}}

### Regexp

Regexp asserts that a specified regexp matches a string.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Regexp(t, regexp.MustCompile("start"), "it's starting")
	assertions.Regexp(t, "start...$", "it's not starting")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: "^start", "starting"
	failure: "^start", "not starting"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Regexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexp) | package-level function |
| [`assert.Regexpf(t T, rx any, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexpf) | formatted variant |
| [`assert.(*Assertions).Regexp(rx any, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Regexp) | method variant |
| [`assert.(*Assertions).Regexpf(rx any, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Regexpf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Regexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Regexp) | package-level function |
| [`require.Regexpf(t T, rx any, str any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Regexpf) | formatted variant |
| [`require.(*Assertions).Regexp(rx any, str any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Regexp) | method variant |
| [`require.(*Assertions).Regexpf(rx any, str any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Regexpf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Regexp(t T, rx any, str any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Regexp) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Regexp](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L22)
{{% /tab %}}
{{< /tabs >}}

---

---

Generated with github.com/go-openapi/testify/codegen/v2

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/codegen/v2 DO NOT EDIT.

Generated on 2026-01-11 (version ca82e58) using codegen version v2.1.9-0.20260111184010-ca82e58db12c+dirty [sha: ca82e58db12cbb61bfcae58c3684b3add9599d10]
-->
