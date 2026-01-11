---
title: "Ordering"
description: "Asserting How Collections Are Ordered"
modified: 2026-01-11
weight: 11
domains:
  - "ordering"
keywords:
  - "IsDecreasing"
  - "IsDecreasingf"
  - "IsIncreasing"
  - "IsIncreasingf"
  - "IsNonDecreasing"
  - "IsNonDecreasingf"
  - "IsNonIncreasing"
  - "IsNonIncreasingf"
---

Asserting How Collections Are Ordered

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.

### IsDecreasing

IsDecreasing asserts that the collection is decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsDecreasing(t, []int{2, 1, 0})
	assertions.IsDecreasing(t, []float{2, 1})
	assertions.IsDecreasing(t, []string{"b", "a"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{3, 2, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasing) | package-level function |
| [`assert.IsDecreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasingf) | formatted variant |
| [`assert.(*Assertions).IsDecreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsDecreasing) | method variant |
| [`assert.(*Assertions).IsDecreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsDecreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasing) | package-level function |
| [`require.IsDecreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasingf) | formatted variant |
| [`require.(*Assertions).IsDecreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsDecreasing) | method variant |
| [`require.(*Assertions).IsDecreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsDecreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsDecreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L63)
{{% /tab %}}
{{< /tabs >}}

### IsIncreasing

IsIncreasing asserts that the collection is increasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsIncreasing(t, []int{1, 2, 3})
	assertions.IsIncreasing(t, []float{1, 2})
	assertions.IsIncreasing(t, []string{"a", "b"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 2, 3}
	failure: []int{1, 1, 2}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasing) | package-level function |
| [`assert.IsIncreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasingf) | formatted variant |
| [`assert.(*Assertions).IsIncreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsIncreasing) | method variant |
| [`assert.(*Assertions).IsIncreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsIncreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasing) | package-level function |
| [`require.IsIncreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasingf) | formatted variant |
| [`require.(*Assertions).IsIncreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsIncreasing) | method variant |
| [`require.(*Assertions).IsIncreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsIncreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsIncreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L23)
{{% /tab %}}
{{< /tabs >}}

### IsNonDecreasing

IsNonDecreasing asserts that the collection is not decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNonDecreasing(t, []int{1, 1, 2})
	assertions.IsNonDecreasing(t, []float{1, 2})
	assertions.IsNonDecreasing(t, []string{"a", "b"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 1, 2}
	failure: []int{2, 1, 1}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsNonDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasing) | package-level function |
| [`assert.IsNonDecreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasingf) | formatted variant |
| [`assert.(*Assertions).IsNonDecreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonDecreasing) | method variant |
| [`assert.(*Assertions).IsNonDecreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonDecreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasing) | package-level function |
| [`require.IsNonDecreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasingf) | formatted variant |
| [`require.(*Assertions).IsNonDecreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonDecreasing) | method variant |
| [`require.(*Assertions).IsNonDecreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonDecreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonDecreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L83)
{{% /tab %}}
{{< /tabs >}}

### IsNonIncreasing

IsNonIncreasing asserts that the collection is not increasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNonIncreasing(t, []int{2, 1, 1})
	assertions.IsNonIncreasing(t, []float{2, 1})
	assertions.IsNonIncreasing(t, []string{"b", "a"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{2, 1, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsNonIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasing) | package-level function |
| [`assert.IsNonIncreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasingf) | formatted variant |
| [`assert.(*Assertions).IsNonIncreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonIncreasing) | method variant |
| [`assert.(*Assertions).IsNonIncreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonIncreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasing) | package-level function |
| [`require.IsNonIncreasingf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasingf) | formatted variant |
| [`require.(*Assertions).IsNonIncreasing(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonIncreasing) | method variant |
| [`require.(*Assertions).IsNonIncreasingf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonIncreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonIncreasing(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L43)
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
