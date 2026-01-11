---
title: "Collection"
description: "Asserting Slices And Maps"
modified: 2026-01-11
weight: 2
domains:
  - "collection"
keywords:
  - "Contains"
  - "Containsf"
  - "ElementsMatch"
  - "ElementsMatchf"
  - "Len"
  - "Lenf"
  - "NotContains"
  - "NotContainsf"
  - "NotElementsMatch"
  - "NotElementsMatchf"
  - "NotSubset"
  - "NotSubsetf"
  - "Subset"
  - "Subsetf"
---

Asserting Slices And Maps

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 7 functionalities.

### Contains

Contains asserts that the specified string, list(array, slice...) or map contains the
specified substring or element.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Contains(t, "Hello World", "World")
	assertions.Contains(t, []string{"Hello", "World"}, "World")
	assertions.Contains(t, map[string]string{"Hello": "World"}, "Hello")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []string{"A","B"}, "A"
	failure: []string{"A","B"}, "C"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Contains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Contains) | package-level function |
| [`assert.Containsf(t T, s any, contains any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Containsf) | formatted variant |
| [`assert.(*Assertions).Contains(s any, contains any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Contains) | method variant |
| [`assert.(*Assertions).Containsf(s any, contains any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Containsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Contains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Contains) | package-level function |
| [`require.Containsf(t T, s any, contains any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Containsf) | formatted variant |
| [`require.(*Assertions).Contains(s any, contains any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Contains) | method variant |
| [`require.(*Assertions).Containsf(s any, contains any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Containsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Contains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Contains) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Contains](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L64)
{{% /tab %}}
{{< /tabs >}}

### ElementsMatch

ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should match.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
	failure: []int{1, 2, 3}, []int{1, 2, 4}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.ElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ElementsMatch) | package-level function |
| [`assert.ElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ElementsMatchf) | formatted variant |
| [`assert.(*Assertions).ElementsMatch(listA any, listB any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ElementsMatch) | method variant |
| [`assert.(*Assertions).ElementsMatchf(listA any, listB any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ElementsMatchf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ElementsMatch) | package-level function |
| [`require.ElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ElementsMatchf) | formatted variant |
| [`require.(*Assertions).ElementsMatch(listA any, listB any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ElementsMatch) | method variant |
| [`require.(*Assertions).ElementsMatchf(listA any, listB any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ElementsMatchf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.ElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ElementsMatch) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ElementsMatch](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L277)
{{% /tab %}}
{{< /tabs >}}

### Len

Len asserts that the specified object has specific length.

Len also fails if the object has a type that len() does not accept.

The asserted object can be a string, a slice, a map, an array or a channel.

See also [reflect.Len](https://pkg.go.dev/reflect#Len).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Len(t, mySlice, 3)
	assertions.Len(t, myString, 4)
	assertions.Len(t, myMap, 5)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []string{"A","B"}, 2
	failure: []string{"A","B"}, 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Len(t T, object any, length int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Len) | package-level function |
| [`assert.Lenf(t T, object any, length int, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Lenf) | formatted variant |
| [`assert.(*Assertions).Len(object any, length int) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Len) | method variant |
| [`assert.(*Assertions).Lenf(object any, length int, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Lenf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Len(t T, object any, length int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Len) | package-level function |
| [`require.Lenf(t T, object any, length int, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Lenf) | formatted variant |
| [`require.(*Assertions).Len(object any, length int) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Len) | method variant |
| [`require.(*Assertions).Lenf(object any, length int, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Lenf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Len(t T, object any, length int, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Len) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Len](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L31)

> **Maintainer Note**
> The implementation is based on [reflect.Len]. The potential panic is handled with recover.
A better approach could be to check for the [reflect.Type] before calling [reflect.Len].

> **Note**
> (proposals) this does not currently support iterators, or collection objects that have a Len() method.
{{% /tab %}}
{{< /tabs >}}

### NotContains

NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
specified substring or element.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotContains(t, "Hello World", "Earth")
	assertions.NotContains(t, ["Hello", "World"], "Earth")
	assertions.NotContains(t, {"Hello": "World"}, "Earth")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []string{"A","B"}, "C"
	failure: []string{"A","B"}, "B"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotContains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotContains) | package-level function |
| [`assert.NotContainsf(t T, s any, contains any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotContainsf) | formatted variant |
| [`assert.(*Assertions).NotContains(s any, contains any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotContains) | method variant |
| [`assert.(*Assertions).NotContainsf(s any, contains any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotContainsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotContains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotContains) | package-level function |
| [`require.NotContainsf(t T, s any, contains any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotContainsf) | formatted variant |
| [`require.(*Assertions).NotContains(s any, contains any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotContains) | method variant |
| [`require.(*Assertions).NotContainsf(s any, contains any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotContainsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotContains(t T, s any, contains any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotContains) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L94)
{{% /tab %}}
{{< /tabs >}}

### NotElementsMatch

NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should not match.
This is an inverse of ElementsMatch.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false
	assertions.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true
	assertions.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 2, 3}, []int{1, 2, 4}
	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotElementsMatch) | package-level function |
| [`assert.NotElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotElementsMatchf) | formatted variant |
| [`assert.(*Assertions).NotElementsMatch(listA any, listB any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotElementsMatch) | method variant |
| [`assert.(*Assertions).NotElementsMatchf(listA any, listB any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotElementsMatchf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotElementsMatch) | package-level function |
| [`require.NotElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotElementsMatchf) | formatted variant |
| [`require.(*Assertions).NotElementsMatch(listA any, listB any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotElementsMatch) | method variant |
| [`require.(*Assertions).NotElementsMatchf(listA any, listB any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotElementsMatchf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotElementsMatch) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotElementsMatch](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L314)
{{% /tab %}}
{{< /tabs >}}

### NotSubset

NotSubset asserts that the list (array, slice, or map) does NOT contain all
elements given in the subset (array, slice, or map).
Map elements are key-value pairs unless compared with an array or slice where
only the map key is evaluated.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSubset(t, [1, 3, 4], [1, 2])
	assertions.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
	assertions.NotSubset(t, [1, 3, 4], {1: "one", 2: "two"})
	assertions.NotSubset(t, {"x": 1, "y": 2}, ["z"])
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 2, 3}, []int{4, 5}
	failure: []int{1, 2, 3}, []int{1, 2}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotSubset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSubset) | package-level function |
| [`assert.NotSubsetf(t T, list any, subset any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSubsetf) | formatted variant |
| [`assert.(*Assertions).NotSubset(list any, subset any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSubset) | method variant |
| [`assert.(*Assertions).NotSubsetf(list any, subset any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSubsetf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotSubset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSubset) | package-level function |
| [`require.NotSubsetf(t T, list any, subset any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSubsetf) | formatted variant |
| [`require.(*Assertions).NotSubset(list any, subset any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSubset) | method variant |
| [`require.(*Assertions).NotSubsetf(list any, subset any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSubsetf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotSubset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSubset) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSubset](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L205)
{{% /tab %}}
{{< /tabs >}}

### Subset

Subset asserts that the list (array, slice, or map) contains all elements
given in the subset (array, slice, or map).

Map elements are key-value pairs unless compared with an array or slice where
only the map key is evaluated.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Subset(t, [1, 2, 3], [1, 2])
	assertions.Subset(t, {"x": 1, "y": 2}, {"x": 1})
	assertions.Subset(t, [1, 2, 3], {1: "one", 2: "two"})
	assertions.Subset(t, {"x": 1, "y": 2}, ["x"])
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []int{1, 2, 3}, []int{1, 2}
	failure: []int{1, 2, 3}, []int{4, 5}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Subset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Subset) | package-level function |
| [`assert.Subsetf(t T, list any, subset any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Subsetf) | formatted variant |
| [`assert.(*Assertions).Subset(list any, subset any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Subset) | method variant |
| [`assert.(*Assertions).Subsetf(list any, subset any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Subsetf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Subset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Subset) | package-level function |
| [`require.Subsetf(t T, list any, subset any, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Subsetf) | formatted variant |
| [`require.(*Assertions).Subset(list any, subset any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Subset) | method variant |
| [`require.(*Assertions).Subsetf(list any, subset any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Subsetf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Subset(t T, list any, subset any, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Subset) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Subset](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L128)
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
