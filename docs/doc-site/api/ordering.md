---
title: "Ordering"
description: "Asserting How Collections Are Ordered"
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
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [IsDecreasing](#isdecreasing) | angles-right
- [IsDecreasingT[OrderedSlice ~[]E, E Ordered]](#isdecreasingtorderedslice-e-e-ordered) | star | orange
- [IsIncreasing](#isincreasing) | angles-right
- [IsIncreasingT[OrderedSlice ~[]E, E Ordered]](#isincreasingtorderedslice-e-e-ordered) | star | orange
- [IsNonDecreasing](#isnondecreasing) | angles-right
- [IsNonDecreasingT[OrderedSlice ~[]E, E Ordered]](#isnondecreasingtorderedslice-e-e-ordered) | star | orange
- [IsNonIncreasing](#isnonincreasing) | angles-right
- [IsNonIncreasingT[OrderedSlice ~[]E, E Ordered]](#isnonincreasingtorderedslice-e-e-ordered) | star | orange
- [NotSortedT[OrderedSlice ~[]E, E Ordered]](#notsortedtorderedslice-e-e-ordered) | star | orange
- [SortedT[OrderedSlice ~[]E, E Ordered]](#sortedtorderedslice-e-e-ordered) | star | orange
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
	success: []int{3, 2, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsDecreasing(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsDecreasing(t, []int{3, 2, 1})
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
| [`assert.IsDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasing) | package-level function |
| [`assert.IsDecreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasingf) | formatted variant |
| [`assert.(*Assertions).IsDecreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsDecreasing) | method variant |
| [`assert.(*Assertions).IsDecreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsDecreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasing) | package-level function |
| [`require.IsDecreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsDecreasingf) | formatted variant |
| [`require.(*Assertions).IsDecreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsDecreasing) | method variant |
| [`require.(*Assertions).IsDecreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsDecreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsDecreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L190)
{{% /tab %}}
{{< /tabs >}}

### IsDecreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isdecreasingtorderedslice-e-e-ordered}

IsDecreasingT asserts that a slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is strictly decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsDecreasingT(t, []int{2, 1, 0})
	assertions.IsDecreasingT(t, []float{2, 1})
	assertions.IsDecreasingT(t, []string{"b", "a"})
	success: []int{3, 2, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsDecreasingT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsDecreasingT(t, []int{3, 2, 1})
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
| [`assertions.IsDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsDecreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsDecreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L220)
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
	success: []int{1, 2, 3}
	failure: []int{1, 1, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsIncreasing(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsIncreasing(t, []int{1, 2, 3})
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
| [`assert.IsIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasing) | package-level function |
| [`assert.IsIncreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasingf) | formatted variant |
| [`assert.(*Assertions).IsIncreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsIncreasing) | method variant |
| [`assert.(*Assertions).IsIncreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsIncreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasing) | package-level function |
| [`require.IsIncreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsIncreasingf) | formatted variant |
| [`require.(*Assertions).IsIncreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsIncreasing) | method variant |
| [`require.(*Assertions).IsIncreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsIncreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsIncreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L24)
{{% /tab %}}
{{< /tabs >}}

### IsIncreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isincreasingtorderedslice-e-e-ordered}

IsIncreasingT asserts that a slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is strictly increasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsIncreasingT(t, []int{1, 2, 3})
	assertions.IsIncreasingT(t, []float{1, 2})
	assertions.IsIncreasingT(t, []string{"a", "b"})
	success: []int{1, 2, 3}
	failure: []int{1, 1, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsIncreasingT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsIncreasingT(t, []int{1, 2, 3})
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
| [`assertions.IsIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsIncreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsIncreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L53)
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
	success: []int{1, 1, 2}
	failure: []int{2, 1, 0}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsNonDecreasing(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsNonDecreasing(t, []int{1, 1, 2})
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
| [`assert.IsNonDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasing) | package-level function |
| [`assert.IsNonDecreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonDecreasingf) | formatted variant |
| [`assert.(*Assertions).IsNonDecreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonDecreasing) | method variant |
| [`assert.(*Assertions).IsNonDecreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonDecreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasing) | package-level function |
| [`require.IsNonDecreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonDecreasingf) | formatted variant |
| [`require.(*Assertions).IsNonDecreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonDecreasing) | method variant |
| [`require.(*Assertions).IsNonDecreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonDecreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonDecreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L246)
{{% /tab %}}
{{< /tabs >}}

### IsNonDecreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isnondecreasingtorderedslice-e-e-ordered}

IsNonDecreasingT asserts that a slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is not decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNonDecreasingT(t, []int{1, 1, 2})
	assertions.IsNonDecreasingT(t, []float{1, 2})
	assertions.IsNonDecreasingT(t, []string{"a", "b"})
	success: []int{1, 1, 2}
	failure: []int{2, 1, 0}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsNonDecreasingT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsNonDecreasingT(t, []int{1, 1, 2})
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
| [`assertions.IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonDecreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L275)
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
	success: []int{2, 1, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsNonIncreasing(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsNonIncreasing(t, []int{2, 1, 1})
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
| [`assert.IsNonIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasing) | package-level function |
| [`assert.IsNonIncreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNonIncreasingf) | formatted variant |
| [`assert.(*Assertions).IsNonIncreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonIncreasing) | method variant |
| [`assert.(*Assertions).IsNonIncreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNonIncreasingf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNonIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasing) | package-level function |
| [`require.IsNonIncreasingf(t T, collection any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNonIncreasingf) | formatted variant |
| [`require.(*Assertions).IsNonIncreasing(collection any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonIncreasing) | method variant |
| [`require.(*Assertions).IsNonIncreasingf(collection any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNonIncreasingf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNonIncreasing(t T, collection any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasing) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasing](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L135)
{{% /tab %}}
{{< /tabs >}}

### IsNonIncreasingT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#isnonincreasingtorderedslice-e-e-ordered}

IsNonIncreasingT asserts that a slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is NOT strictly increasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNonIncreasing(t, []int{2, 1, 1})
	assertions.IsNonIncreasing(t, []float{2, 1})
	assertions.IsNonIncreasing(t, []string{"b", "a"})
	success: []int{2, 1, 1}
	failure: []int{1, 2, 3}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestIsNonIncreasingT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.IsNonIncreasingT(t, []int{2, 1, 1})
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
| [`assertions.IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasingT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNonIncreasingT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L164)
{{% /tab %}}
{{< /tabs >}}

### NotSortedT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#notsortedtorderedslice-e-e-ordered}

NotSortedT asserts that the slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is NOT sorted (i.e. non-strictly increasing).

Unlike [IsDecreasingT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsDecreasingT), it accepts slices that are neither increasing nor decreasing.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSortedT(t, []int{3, 2, 3})
	assertions.NotSortedT(t, []float{2, 1})
	assertions.NotSortedT(t, []string{"b", "a"})
	success: []int{3, 1, 3}
	failure: []int{1, 4, 8}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotSortedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotSortedT(t, []int{3, 1, 3})
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
| [`assertions.NotSortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSortedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSortedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L109)
{{% /tab %}}
{{< /tabs >}}

### SortedT[OrderedSlice ~[]E, E Ordered] {{% icon icon="star" color=orange %}}{#sortedtorderedslice-e-e-ordered}

SortedT asserts that the slice of [Ordered](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Ordered) is sorted (i.e. non-strictly increasing).

Unlike [IsIncreasingT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsIncreasingT), it accepts elements to be equal.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SortedT(t, []int{1, 2, 3})
	assertions.SortedT(t, []float{1, 2})
	assertions.SortedT(t, []string{"a", "b"})
	success: []int{1, 1, 3}
	failure: []int{1, 4, 2}
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSortedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.SortedT(t, []int{1, 1, 3})
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
| [`assertions.SortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SortedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SortedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/order.go#L81)
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
