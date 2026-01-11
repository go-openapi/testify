---
title: "Comparison"
description: "Comparing Ordered Values"
modified: 2026-01-11
weight: 3
domains:
  - "comparison"
keywords:
  - "Greater"
  - "Greaterf"
  - "GreaterOrEqual"
  - "GreaterOrEqualf"
  - "Less"
  - "Lessf"
  - "LessOrEqual"
  - "LessOrEqualf"
  - "Negative"
  - "Negativef"
  - "Positive"
  - "Positivef"
---

Comparing Ordered Values

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 6 functionalities.

### Greater

Greater asserts that the first element is strictly greater than the second.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Greater(t, 2, 1)
	assertions.Greater(t, float64(2), float64(1))
	assertions.Greater(t, "b", "a")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 2, 1
	failure: 1, 2
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Greater(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Greater) | package-level function |
| [`assert.Greaterf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Greaterf) | formatted variant |
| [`assert.(*Assertions).Greater(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Greater) | method variant |
| [`assert.(*Assertions).Greaterf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Greaterf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Greater(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Greater) | package-level function |
| [`require.Greaterf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Greaterf) | formatted variant |
| [`require.(*Assertions).Greater(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Greater) | method variant |
| [`require.(*Assertions).Greaterf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Greaterf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Greater(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Greater) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Greater](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L56)
{{% /tab %}}
{{< /tabs >}}

### GreaterOrEqual

GreaterOrEqual asserts that the first element is greater than or equal to the second.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.GreaterOrEqual(t, 2, 1)
	assertions.GreaterOrEqual(t, 2, 2)
	assertions.GreaterOrEqual(t, "b", "a")
	assertions.GreaterOrEqual(t, "b", "b")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 2, 1
	failure: 1, 2
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterOrEqual) | package-level function |
| [`assert.GreaterOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterOrEqualf) | formatted variant |
| [`assert.(*Assertions).GreaterOrEqual(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.GreaterOrEqual) | method variant |
| [`assert.(*Assertions).GreaterOrEqualf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.GreaterOrEqualf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterOrEqual) | package-level function |
| [`require.GreaterOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterOrEqualf) | formatted variant |
| [`require.(*Assertions).GreaterOrEqual(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.GreaterOrEqual) | method variant |
| [`require.(*Assertions).GreaterOrEqualf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.GreaterOrEqualf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#GreaterOrEqual) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#GreaterOrEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L78)
{{% /tab %}}
{{< /tabs >}}

### Less

Less asserts that the first element is strictly less than the second.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Less(t, 1, 2)
	assertions.Less(t, float64(1), float64(2))
	assertions.Less(t, "a", "b")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1, 2
	failure: 2, 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Less(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Less) | package-level function |
| [`assert.Lessf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Lessf) | formatted variant |
| [`assert.(*Assertions).Less(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Less) | method variant |
| [`assert.(*Assertions).Lessf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Lessf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Less(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Less) | package-level function |
| [`require.Lessf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Lessf) | formatted variant |
| [`require.(*Assertions).Less(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Less) | method variant |
| [`require.(*Assertions).Lessf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Lessf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Less(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Less) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Less](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L99)
{{% /tab %}}
{{< /tabs >}}

### LessOrEqual

LessOrEqual asserts that the first element is less than or equal to the second.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.LessOrEqual(t, 1, 2)
	assertions.LessOrEqual(t, 2, 2)
	assertions.LessOrEqual(t, "a", "b")
	assertions.LessOrEqual(t, "b", "b")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1, 2
	failure: 2, 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessOrEqual) | package-level function |
| [`assert.LessOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessOrEqualf) | formatted variant |
| [`assert.(*Assertions).LessOrEqual(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.LessOrEqual) | method variant |
| [`assert.(*Assertions).LessOrEqualf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.LessOrEqualf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessOrEqual) | package-level function |
| [`require.LessOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessOrEqualf) | formatted variant |
| [`require.(*Assertions).LessOrEqual(e1 any, e2 any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.LessOrEqual) | method variant |
| [`require.(*Assertions).LessOrEqualf(e1 any, e2 any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.LessOrEqualf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#LessOrEqual) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#LessOrEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L121)
{{% /tab %}}
{{< /tabs >}}

### Negative

Negative asserts that the specified element is strictly negative.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Negative(t, -1)
	assertions.Negative(t, -1.23)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: -1
	failure: 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Negative(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Negative) | package-level function |
| [`assert.Negativef(t T, e any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Negativef) | formatted variant |
| [`assert.(*Assertions).Negative(e any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Negative) | method variant |
| [`assert.(*Assertions).Negativef(e any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Negativef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Negative(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Negative) | package-level function |
| [`require.Negativef(t T, e any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Negativef) | formatted variant |
| [`require.(*Assertions).Negative(e any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Negative) | method variant |
| [`require.(*Assertions).Negativef(e any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Negativef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Negative(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Negative) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Negative](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L162)
{{% /tab %}}
{{< /tabs >}}

### Positive

Positive asserts that the specified element is strictly positive.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Positive(t, 1)
	assertions.Positive(t, 1.23)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1
	failure: -1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Positive(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Positive) | package-level function |
| [`assert.Positivef(t T, e any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Positivef) | formatted variant |
| [`assert.(*Assertions).Positive(e any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Positive) | method variant |
| [`assert.(*Assertions).Positivef(e any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Positivef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Positive(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Positive) | package-level function |
| [`require.Positivef(t T, e any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Positivef) | formatted variant |
| [`require.(*Assertions).Positive(e any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Positive) | method variant |
| [`require.(*Assertions).Positivef(e any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Positivef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Positive(t T, e any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Positive) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Positive](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L141)
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

Generated on 2026-01-11 (version e6b0793) using codegen version v2.1.9-0.20260111152118-e6b0793ba519+dirty [sha: e6b0793ba519fb22dc1887392e1465649a5a95ff]
-->
