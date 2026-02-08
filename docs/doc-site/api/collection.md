---
title: "Collection"
description: "Asserting Slices And Maps"
weight: 2
domains:
  - "collection"
keywords:
  - "Contains"
  - "Containsf"
  - "ElementsMatch"
  - "ElementsMatchf"
  - "ElementsMatchT"
  - "ElementsMatchTf"
  - "Len"
  - "Lenf"
  - "MapContainsT"
  - "MapContainsTf"
  - "MapNotContainsT"
  - "MapNotContainsTf"
  - "NotContains"
  - "NotContainsf"
  - "NotElementsMatch"
  - "NotElementsMatchf"
  - "NotElementsMatchT"
  - "NotElementsMatchTf"
  - "NotSubset"
  - "NotSubsetf"
  - "SeqContainsT"
  - "SeqContainsTf"
  - "SeqNotContainsT"
  - "SeqNotContainsTf"
  - "SliceContainsT"
  - "SliceContainsTf"
  - "SliceNotContainsT"
  - "SliceNotContainsTf"
  - "SliceNotSubsetT"
  - "SliceNotSubsetTf"
  - "SliceSubsetT"
  - "SliceSubsetTf"
  - "StringContainsT"
  - "StringContainsTf"
  - "StringNotContainsT"
  - "StringNotContainsTf"
  - "Subset"
  - "Subsetf"
---

Asserting Slices And Maps

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 19 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [Contains](#contains) | angles-right
- [ElementsMatch](#elementsmatch) | angles-right
- [ElementsMatchT[E comparable]](#elementsmatchte-comparable) | star | orange
- [Len](#len) | angles-right
- [MapContainsT[Map ~map[K]V, K comparable, V any]](#mapcontainstmap-mapkv-k-comparable-v-any) | star | orange
- [MapNotContainsT[Map ~map[K]V, K comparable, V any]](#mapnotcontainstmap-mapkv-k-comparable-v-any) | star | orange
- [NotContains](#notcontains) | angles-right
- [NotElementsMatch](#notelementsmatch) | angles-right
- [NotElementsMatchT[E comparable]](#notelementsmatchte-comparable) | star | orange
- [NotSubset](#notsubset) | angles-right
- [SeqContainsT[E comparable]](#seqcontainste-comparable) | star | orange
- [SeqNotContainsT[E comparable]](#seqnotcontainste-comparable) | star | orange
- [SliceContainsT[Slice ~[]E, E comparable]](#slicecontainstslice-e-e-comparable) | star | orange
- [SliceNotContainsT[Slice ~[]E, E comparable]](#slicenotcontainstslice-e-e-comparable) | star | orange
- [SliceNotSubsetT[Slice ~[]E, E comparable]](#slicenotsubsettslice-e-e-comparable) | star | orange
- [SliceSubsetT[Slice ~[]E, E comparable]](#slicesubsettslice-e-e-comparable) | star | orange
- [StringContainsT[ADoc, EDoc Text]](#stringcontainstadoc-edoc-text) | star | orange
- [StringNotContainsT[ADoc, EDoc Text]](#stringnotcontainstadoc-edoc-text) | star | orange
- [Subset](#subset) | angles-right
```

### Contains{#contains}
Contains asserts that the specified string, list(array, slice...) or map contains the
specified substring or element.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Contains(t, "Hello World", "World")
	assertions.Contains(t, []string{"Hello", "World"}, "World")
	assertions.Contains(t, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"Hello": "World"}, "Hello")
	success: []string{"A","B"}, "A"
	failure: []string{"A","B"}, "C"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestContains(t *testing.T)
	success := assert.Contains(t, []string{"A", "B"}, "A")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestContains(t *testing.T)
	require.Contains(t, []string{"A", "B"}, "A")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Contains](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L65)
{{% /tab %}}
{{< /tabs >}}

### ElementsMatch{#elementsmatch}
ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should match.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
	failure: []int{1, 2, 3}, []int{1, 2, 4}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestElementsMatch(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatch(t *testing.T)
	success := assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestElementsMatch(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatch(t *testing.T)
	require.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ElementsMatch](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L540)
{{% /tab %}}
{{< /tabs >}}

### ElementsMatchT[E comparable] {{% icon icon="star" color=orange %}}{#elementsmatchte-comparable}
ElementsMatchT asserts that the specified listA(array, slice...) is equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should match.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
	failure: []int{1, 2, 3}, []int{1, 2, 4}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestElementsMatchT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatchT(t *testing.T)
	success := assert.ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestElementsMatchT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatchT(t *testing.T)
	require.ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.ElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ElementsMatchT) | package-level function |
| [`assert.ElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ElementsMatchTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ElementsMatchT) | package-level function |
| [`require.ElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ElementsMatchTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ElementsMatchT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ElementsMatchT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L613)
{{% /tab %}}
{{< /tabs >}}

### Len{#len}
Len asserts that the specified object has specific length.

Len also fails if the object has a type that len() does not accept.

The asserted object can be a string, a slice, a map, an array, pointer to array or a channel.

See also [reflect.Len](https://pkg.go.dev/reflect#Len).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Len(t, mySlice, 3)
	assertions.Len(t, myString, 4)
	assertions.Len(t, myMap, 5)
	success: []string{"A","B"}, 2
	failure: []string{"A","B"}, 1
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestLen(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestLen(t *testing.T)
	success := assert.Len(t, []string{"A", "B"}, 2)
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestLen(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestLen(t *testing.T)
	require.Len(t, []string{"A", "B"}, 2)
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Len](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L33)

> **Note**
>
> (proposal for enhancement) this does not currently support iterators, or collection objects that have a Len() method.
>
{{% /tab %}}
{{< /tabs >}}

### MapContainsT[Map ~map[K]V, K comparable, V any] {{% icon icon="star" color=orange %}}{#mapcontainstmap-mapkv-k-comparable-v-any}
MapContainsT asserts that the specified map contains a key.

Go native comparable types are explained there: [comparable-types](https://go.dev/blog/comparable).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.MapContainsT(t, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"Hello": "x","World": "y"}, "World")
	success: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"A": "B"}, "A"
	failure: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"A": "B"}, "C"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestMapContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestMapContainsT(t *testing.T)
	success := assert.MapContainsT(t, map[string]string{"A": "B"}, "A")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestMapContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestMapContainsT(t *testing.T)
	require.MapContainsT(t, map[string]string{"A": "B"}, "A")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.MapContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#MapContainsT) | package-level function |
| [`assert.MapContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#MapContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.MapContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#MapContainsT) | package-level function |
| [`require.MapContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#MapContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.MapContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#MapContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#MapContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L179)
{{% /tab %}}
{{< /tabs >}}

### MapNotContainsT[Map ~map[K]V, K comparable, V any] {{% icon icon="star" color=orange %}}{#mapnotcontainstmap-mapkv-k-comparable-v-any}
MapNotContainsT asserts that the specified map does not contain a key.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.MapNotContainsT(t, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"Hello": "x","World": "y"}, "hi")
	success: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"A": "B"}, "C"
	failure: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)string{"A": "B"}, "A"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestMapNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestMapNotContainsT(t *testing.T)
	success := assert.MapNotContainsT(t, map[string]string{"A": "B"}, "C")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestMapNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestMapNotContainsT(t *testing.T)
	require.MapNotContainsT(t, map[string]string{"A": "B"}, "C")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.MapNotContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#MapNotContainsT) | package-level function |
| [`assert.MapNotContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#MapNotContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.MapNotContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#MapNotContainsT) | package-level function |
| [`require.MapNotContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#MapNotContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.MapNotContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#MapNotContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#MapNotContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L310)
{{% /tab %}}
{{< /tabs >}}

### NotContains{#notcontains}
NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
specified substring or element.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotContains(t, "Hello World", "Earth")
	assertions.NotContains(t, ["Hello", "World"], "Earth")
	assertions.NotContains(t, {"Hello": "World"}, "Earth")
	success: []string{"A","B"}, "C"
	failure: []string{"A","B"}, "B"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotContains(t *testing.T)
	success := assert.NotContains(t, []string{"A", "B"}, "C")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotContains(t *testing.T)
	require.NotContains(t, []string{"A", "B"}, "C")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L206)
{{% /tab %}}
{{< /tabs >}}

### NotElementsMatch{#notelementsmatch}
NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should not match.
This is an inverse of ElementsMatch.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotElementsMatch(t, []int{1, 1, 2, 3}, []int{1, 1, 2, 3}) -> false
	assertions.NotElementsMatch(t, []int{1, 1, 2, 3}, []int{1, 2, 3}) -> true
	assertions.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4}) -> true
	success: []int{1, 2, 3}, []int{1, 2, 4}
	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotElementsMatch(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatch(t *testing.T)
	success := assert.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotElementsMatch(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatch(t *testing.T)
	require.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotElementsMatch](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L577)
{{% /tab %}}
{{< /tabs >}}

### NotElementsMatchT[E comparable] {{% icon icon="star" color=orange %}}{#notelementsmatchte-comparable}
NotElementsMatchT asserts that the specified listA(array, slice...) is NOT equal to specified
listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
the number of appearances of each of them in both lists should not match.
This is an inverse of ElementsMatch.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotElementsMatchT(t, []int{1, 1, 2, 3}, []int{1, 1, 2, 3}) -> false
	assertions.NotElementsMatchT(t, []int{1, 1, 2, 3}, []int{1, 2, 3}) -> true
	assertions.NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4}) -> true
	success: []int{1, 2, 3}, []int{1, 2, 4}
	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotElementsMatchT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatchT(t *testing.T)
	success := assert.NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotElementsMatchT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatchT(t *testing.T)
	require.NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotElementsMatchT) | package-level function |
| [`assert.NotElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotElementsMatchTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotElementsMatchT) | package-level function |
| [`require.NotElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotElementsMatchTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotElementsMatchT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotElementsMatchT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L649)
{{% /tab %}}
{{< /tabs >}}

### NotSubset{#notsubset}
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
	success: []int{1, 2, 3}, []int{4, 5}
	failure: []int{1, 2, 3}, []int{1, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotSubset(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSubset(t *testing.T)
	success := assert.NotSubset(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotSubset(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSubset(t *testing.T)
	require.NotSubset(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSubset](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L446)
{{% /tab %}}
{{< /tabs >}}

### SeqContainsT[E comparable] {{% icon icon="star" color=orange %}}{#seqcontainste-comparable}
SeqContainsT asserts that the specified iterator contains a comparable element.

The sequence may not be consumed entirely: the iteration stops as soon as the specified element is found.

Go native comparable types are explained there: [comparable-types](https://go.dev/blog/comparable).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SeqContainsT(t, slices.Values([]{"Hello","World"}), "World")
	success: slices.Values([]string{"A","B"}), "A"
	failure: slices.Values([]string{"A","B"}), "C"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSeqContainsT(t *testing.T)
package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqContainsT(t *testing.T)
	success := assert.SeqContainsT(t, slices.Values([]string{"A", "B"}), "A")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSeqContainsT(t *testing.T)
package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqContainsT(t *testing.T)
	require.SeqContainsT(t, slices.Values([]string{"A", "B"}), "A")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SeqContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SeqContainsT) | package-level function |
| [`assert.SeqContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SeqContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SeqContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SeqContainsT) | package-level function |
| [`require.SeqContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SeqContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SeqContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SeqContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SeqContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L150)
{{% /tab %}}
{{< /tabs >}}

### SeqNotContainsT[E comparable] {{% icon icon="star" color=orange %}}{#seqnotcontainste-comparable}
SeqNotContainsT asserts that the specified iterator does not contain a comparable element.

See [SeqContainsT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SeqContainsT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SeqContainsT(t, slices.Values([]{"Hello","World"}), "World")
	success: slices.Values([]string{"A","B"}), "C"
	failure: slices.Values([]string{"A","B"}), "A"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSeqNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqNotContainsT(t *testing.T)
	success := assert.SeqNotContainsT(t, slices.Values([]string{"A", "B"}), "C")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSeqNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqNotContainsT(t *testing.T)
	require.SeqNotContainsT(t, slices.Values([]string{"A", "B"}), "C")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SeqNotContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SeqNotContainsT) | package-level function |
| [`assert.SeqNotContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SeqNotContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SeqNotContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SeqNotContainsT) | package-level function |
| [`require.SeqNotContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SeqNotContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SeqNotContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SeqNotContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SeqNotContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L285)
{{% /tab %}}
{{< /tabs >}}

### SliceContainsT[Slice ~[]E, E comparable] {{% icon icon="star" color=orange %}}{#slicecontainstslice-e-e-comparable}
SliceContainsT asserts that the specified slice contains a comparable element.

Go native comparable types are explained there: [comparable-types](https://go.dev/blog/comparable).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SliceContainsT(t, []{"Hello","World"}, "World")
	success: []string{"A","B"}, "A"
	failure: []string{"A","B"}, "C"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceContainsT(t *testing.T)
	success := assert.SliceContainsT(t, []string{"A", "B"}, "A")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceContainsT(t *testing.T)
	require.SliceContainsT(t, []string{"A", "B"}, "A")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SliceContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceContainsT) | package-level function |
| [`assert.SliceContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SliceContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceContainsT) | package-level function |
| [`require.SliceContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SliceContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SliceContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SliceContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L121)
{{% /tab %}}
{{< /tabs >}}

### SliceNotContainsT[Slice ~[]E, E comparable] {{% icon icon="star" color=orange %}}{#slicenotcontainstslice-e-e-comparable}
SliceNotContainsT asserts that the specified slice does not contain a comparable element.

See [SliceContainsT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceContainsT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SliceNotContainsT(t, []{"Hello","World"}, "hi")
	success: []string{"A","B"}, "C"
	failure: []string{"A","B"}, "A"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotContainsT(t *testing.T)
	success := assert.SliceNotContainsT(t, []string{"A", "B"}, "C")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotContainsT(t *testing.T)
	require.SliceNotContainsT(t, []string{"A", "B"}, "C")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SliceNotContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceNotContainsT) | package-level function |
| [`assert.SliceNotContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceNotContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SliceNotContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceNotContainsT) | package-level function |
| [`require.SliceNotContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceNotContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SliceNotContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SliceNotContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SliceNotContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L260)
{{% /tab %}}
{{< /tabs >}}

### SliceNotSubsetT[Slice ~[]E, E comparable] {{% icon icon="star" color=orange %}}{#slicenotsubsettslice-e-e-comparable}
SliceNotSubsetT asserts that a slice of comparable elements does not contain all the elements given in the subset.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SliceNotSubsetT(t, []int{1, 2, 3}, []int{1, 4})
	success: []int{1, 2, 3}, []int{4, 5}
	failure: []int{1, 2, 3}, []int{1, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceNotSubsetT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotSubsetT(t *testing.T)
	success := assert.SliceNotSubsetT(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceNotSubsetT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotSubsetT(t *testing.T)
	require.SliceNotSubsetT(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SliceNotSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceNotSubsetT) | package-level function |
| [`assert.SliceNotSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceNotSubsetTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SliceNotSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceNotSubsetT) | package-level function |
| [`require.SliceNotSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceNotSubsetTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SliceNotSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SliceNotSubsetT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SliceNotSubsetT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L513)
{{% /tab %}}
{{< /tabs >}}

### SliceSubsetT[Slice ~[]E, E comparable] {{% icon icon="star" color=orange %}}{#slicesubsettslice-e-e-comparable}
SliceSubsetT asserts that a slice of comparable elements contains all the elements given in the subset.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
	success: []int{1, 2, 3}, []int{1, 2}
	failure: []int{1, 2, 3}, []int{4, 5}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceSubsetT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceSubsetT(t *testing.T)
	success := assert.SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSliceSubsetT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceSubsetT(t *testing.T)
	require.SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.SliceSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceSubsetT) | package-level function |
| [`assert.SliceSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SliceSubsetTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SliceSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceSubsetT) | package-level function |
| [`require.SliceSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SliceSubsetTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SliceSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) (ok bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SliceSubsetT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SliceSubsetT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L415)
{{% /tab %}}
{{< /tabs >}}

### StringContainsT[ADoc, EDoc Text] {{% icon icon="star" color=orange %}}{#stringcontainstadoc-edoc-text}
StringContainsT asserts that a string contains the specified substring.

Strings may be go strings or []byte according to the type constraint [Text](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Text).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.StringContainsT(t, "Hello World", "World")
	success: "AB", "A"
	failure: "AB", "C"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestStringContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestStringContainsT(t *testing.T)
	success := assert.StringContainsT(t, "AB", "A")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestStringContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestStringContainsT(t *testing.T)
	require.StringContainsT(t, "AB", "A")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.StringContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#StringContainsT) | package-level function |
| [`assert.StringContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#StringContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.StringContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#StringContainsT) | package-level function |
| [`require.StringContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#StringContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.StringContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#StringContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#StringContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L94)
{{% /tab %}}
{{< /tabs >}}

### StringNotContainsT[ADoc, EDoc Text] {{% icon icon="star" color=orange %}}{#stringnotcontainstadoc-edoc-text}
StringNotContainsT asserts that a string does not contain the specified substring.

See [StringContainsT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#StringContainsT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.StringNotContainsT(t, "Hello World", "hi")
	success: "AB", "C"
	failure: "AB", "A"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestStringNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestStringNotContainsT(t *testing.T)
	success := assert.StringNotContainsT(t, "AB", "C")
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestStringNotContainsT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestStringNotContainsT(t *testing.T)
	require.StringNotContainsT(t, "AB", "C")
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.StringNotContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#StringNotContainsT) | package-level function |
| [`assert.StringNotContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#StringNotContainsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.StringNotContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#StringNotContainsT) | package-level function |
| [`require.StringNotContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#StringNotContainsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.StringNotContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#StringNotContainsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#StringNotContainsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L235)
{{% /tab %}}
{{< /tabs >}}

### Subset{#subset}
Subset asserts that the list (array, slice, or map) contains all elements
given in the subset (array, slice, or map).

Map elements are key-value pairs unless compared with an array or slice where
only the map key is evaluated.

nil values are considered as empty sets.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Subset(t, []int{1, 2, 3}, []int{1, 2})
	assertions.Subset(t, []string{"x": 1, "y": 2}, []string{"x": 1})
	assertions.Subset(t, []int{1, 2, 3}, map[int](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#int)string{1: "one", 2: "two"})
	assertions.Subset(t, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)int{"x": 1, "y": 2}, []string{"x"})
	success: []int{1, 2, 3}, []int{1, 2}
	failure: []int{1, 2, 3}, []int{4, 5}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSubset(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSubset(t *testing.T)
	success := assert.Subset(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSubset(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestSubset(t *testing.T)
	require.Subset(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% /cards %}}
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Subset](https://github.com/go-openapi/testify/blob/master/internal/assertions/collection.go#L345)
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
-->
