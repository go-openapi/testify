---
title: "Boolean"
description: "Asserting Boolean Values"
modified: 2026-01-11
weight: 1
domains:
  - "boolean"
keywords:
  - "False"
  - "Falsef"
  - "True"
  - "Truef"
---

Asserting Boolean Values

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 2 functionalities.

### False

False asserts that the specified value is false.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.False(t, myBool)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1 == 0
	failure: 1 == 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#False) | package-level function |
| [`assert.Falsef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Falsef) | formatted variant |
| [`assert.(*Assertions).False(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.False) | method variant |
| [`assert.(*Assertions).Falsef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Falsef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#False) | package-level function |
| [`require.Falsef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Falsef) | formatted variant |
| [`require.(*Assertions).False(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.False) | method variant |
| [`require.(*Assertions).Falsef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Falsef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#False) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#False](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L38)
{{% /tab %}}
{{< /tabs >}}

### True

True asserts that the specified value is true.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.True(t, myBool)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1 == 1
	failure: 1 == 0
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#True) | package-level function |
| [`assert.Truef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Truef) | formatted variant |
| [`assert.(*Assertions).True(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.True) | method variant |
| [`assert.(*Assertions).Truef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Truef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#True) | package-level function |
| [`require.Truef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Truef) | formatted variant |
| [`require.(*Assertions).True(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.True) | method variant |
| [`require.(*Assertions).Truef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Truef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#True) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#True](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L16)
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
