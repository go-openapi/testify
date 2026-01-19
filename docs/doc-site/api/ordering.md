---
title: "Ordering"
description: "Asserting How Collections Are Ordered"
modified: 2026-01-20
weight: 11
domains:
  - "ordering"
keywords:
  - "IsDecreasing"
  - "IsDecreasingf"
  - "IsDecreasingT"
  - "IsDecreasingTf"
  - "IsIncreasing"
  - "IsIncreasingf"
  - "IsIncreasingT"
  - "IsIncreasingTf"
  - "IsNonDecreasing"
  - "IsNonDecreasingf"
  - "IsNonDecreasingT"
  - "IsNonDecreasingTf"
  - "IsNonIncreasing"
  - "IsNonIncreasingf"
  - "IsNonIncreasingT"
  - "IsNonIncreasingTf"
  - "NotSortedT"
  - "NotSortedTf"
  - "SortedT"
  - "SortedTf"
---

Asserting How Collections Are Ordered

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 10 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}

```tree
- [IsDecreasing](#isdecreasing) | angles-right
- [IsDecreasingT[OrderedSlice ~[]E, E Ordered]](#isdecreasingtorderedslice-~e-e-ordered) | star | orange
- [IsIncreasing](#isincreasing) | angles-right
- [IsIncreasingT[OrderedSlice ~[]E, E Ordered]](#isincreasingtorderedslice-~e-e-ordered) | star | orange
- [IsNonDecreasing](#isnondecreasing) | angles-right
- [IsNonDecreasingT[OrderedSlice ~[]E, E Ordered]](#isnondecreasingtorderedslice-~e-e-ordered) | star | orange
- [IsNonIncreasing](#isnonincreasing) | angles-right
- [IsNonIncreasingT[OrderedSlice ~[]E, E Ordered]](#isnonincreasingtorderedslice-~e-e-ordered) | star | orange
- [NotSortedT[OrderedSlice ~[]E, E Ordered]](#notsortedtorderedslice-~e-e-ordered) | star | orange
- [SortedT[OrderedSlice ~[]E, E Ordered]](#sortedtorderedslice-~e-e-ordered) | star | orange
```

### IsDecreasing{#isdecreasing}

IsDecreasing asserts that the collection is strictly decreasing.

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L170)
{{% /tab %}}
{{< /tabs >}}

### IsDecreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isdecreasingtorderedslice-~e-e-ordered}

IsDecreasingT asserts that a slice of [Ordered] is strictly decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsDecreasingT(t, []int{2, 1, 0})
	assertions.IsDecreasingT(t, []float{2, 1})
	assertions.IsDecreasingT(t, []string{"b", "a"})
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
| [`assert.IsDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasingT) | package-level function |
| [`assert.IsDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasingTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasingT) | package-level function |
| [`require.IsDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasingTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsDecreasingT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsDecreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsDecreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L190)
{{% /tab %}}
{{< /tabs >}}

### IsIncreasing{#isincreasing}

IsIncreasing asserts that the collection is strictly increasing.

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L24)
{{% /tab %}}
{{< /tabs >}}

### IsIncreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isincreasingtorderedslice-~e-e-ordered}

IsIncreasingT asserts that a slice of [Ordered] is strictly increasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsIncreasingT(t, []int{1, 2, 3})
	assertions.IsIncreasingT(t, []float{1, 2})
	assertions.IsIncreasingT(t, []string{"a", "b"})
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
| [`assert.IsIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasingT) | package-level function |
| [`assert.IsIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasingTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasingT) | package-level function |
| [`require.IsIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasingTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsIncreasingT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsIncreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsIncreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L44)
{{% /tab %}}
{{< /tabs >}}

### IsNonDecreasing{#isnondecreasing}

IsNonDecreasing asserts that the collection is not strictly decreasing.

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
	failure: []int{2, 1, 0}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L216)
{{% /tab %}}
{{< /tabs >}}

### IsNonDecreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isnondecreasingtorderedslice-~e-e-ordered}

IsNonDecreasingT asserts that a slice of [Ordered] is not decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNonDecreasingT(t, []int{1, 1, 2})
	assertions.IsNonDecreasingT(t, []float{1, 2})
	assertions.IsNonDecreasingT(t, []string{"a", "b"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 1, 2}
	failure: []int{2, 1, 0}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasingT) | package-level function |
| [`assert.IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasingTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasingT) | package-level function |
| [`require.IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasingTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonDecreasingT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L236)
{{% /tab %}}
{{< /tabs >}}

### IsNonIncreasing{#isnonincreasing}

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L124)
{{% /tab %}}
{{< /tabs >}}

### IsNonIncreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isnonincreasingtorderedslice-~e-e-ordered}

IsNonIncreasingT asserts that a slice of [Ordered] is NOT strictly increasing.

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
| [`assert.IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasingT) | package-level function |
| [`assert.IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasingTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasingT) | package-level function |
| [`require.IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasingTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonIncreasingT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L144)
{{% /tab %}}
{{< /tabs >}}

### NotSortedT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#notsortedtorderedslice-~e-e-ordered}

NotSortedT asserts that the slice of [Ordered] is NOT sorted (i.e. non-strictly increasing).

Unlike [IsDecreasingT], it accepts slices that are neither increasing nor decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSortedT(t, []int{3, 2, 3})
	assertions.NotSortedT(t, []float{2, 1})
	assertions.NotSortedT(t, []string{"b", "a"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{3, 1, 3}
	failure: []int{1, 4, 8}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotSortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSortedT) | package-level function |
| [`assert.NotSortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSortedTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotSortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSortedT) | package-level function |
| [`require.NotSortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSortedTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotSortedT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSortedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSortedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L99)
{{% /tab %}}
{{< /tabs >}}

### SortedT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#sortedtorderedslice-~e-e-ordered}

SortedT asserts that the slice of [Ordered] is sorted (i.e. non-strictly increasing).

Unlike [IsIncreasingT], it accepts elements to be equal.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SortedT(t, []int{1, 2, 3})
	assertions.SortedT(t, []float{1, 2})
	assertions.SortedT(t, []string{"a", "b"})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 1, 3}
	failure: []int{1, 4, 2}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SortedT) | package-level function |
| [`assert.SortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SortedTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SortedT) | package-level function |
| [`require.SortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SortedTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.SortedT(t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SortedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SortedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L72)
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

Generated on 2026-01-20 (version ff38347) using codegen version v2.1.9-0.20260119220113-ff3834752ffb+dirty [sha: ff3834752ffbc6e4e938c8a0293cc8363f861398]
-->
