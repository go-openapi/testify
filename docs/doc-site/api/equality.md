---
title: "Equality"
description: "Asserting Two Things Are Equal"
weight: 5
domains:
  - "equality"
keywords:
  - "Empty"
  - "Emptyf"
  - "Equal"
  - "Equalf"
  - "EqualExportedValues"
  - "EqualExportedValuesf"
  - "EqualT"
  - "EqualTf"
  - "EqualValues"
  - "EqualValuesf"
  - "Exactly"
  - "Exactlyf"
  - "Nil"
  - "Nilf"
  - "NotEmpty"
  - "NotEmptyf"
  - "NotEqual"
  - "NotEqualf"
  - "NotEqualT"
  - "NotEqualTf"
  - "NotEqualValues"
  - "NotEqualValuesf"
  - "NotNil"
  - "NotNilf"
  - "NotSame"
  - "NotSamef"
  - "NotSameT"
  - "NotSameTf"
  - "Same"
  - "Samef"
  - "SameT"
  - "SameTf"
---

Asserting Two Things Are Equal

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 16 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [Empty](#empty) | angles-right
- [Equal](#equal) | angles-right
- [EqualExportedValues](#equalexportedvalues) | angles-right
- [EqualT[V comparable]](#equaltv-comparable) | star | orange
- [EqualValues](#equalvalues) | angles-right
- [Exactly](#exactly) | angles-right
- [Nil](#nil) | angles-right
- [NotEmpty](#notempty) | angles-right
- [NotEqual](#notequal) | angles-right
- [NotEqualT[V comparable]](#notequaltv-comparable) | star | orange
- [NotEqualValues](#notequalvalues) | angles-right
- [NotNil](#notnil) | angles-right
- [NotSame](#notsame) | angles-right
- [NotSameT[P any]](#notsametp-any) | star | orange
- [Same](#same) | angles-right
- [SameT[P any]](#sametp-any) | star | orange
```

### Empty{#empty}
Empty asserts that the given value is "empty".

Zero values are "empty".

Arrays are "empty" if every element is the zero value of the type (stricter than "empty").

Slices, maps and channels with zero length are "empty".

Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".


[Zero values]: https://go.dev/ref/spec#The_zero_value
{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Empty(t, obj)
	success: ""
	failure: "not empty"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestEmpty(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Empty(t, "")
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
// real-world test would inject *testing.T from TestEmpty(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Empty(t, "")
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
| [`assert.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Empty) | package-level function |
| [`assert.Emptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Emptyf) | formatted variant |
| [`assert.(*Assertions).Empty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Empty) | method variant |
| [`assert.(*Assertions).Emptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Emptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Empty) | package-level function |
| [`require.Emptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Emptyf) | formatted variant |
| [`require.(*Assertions).Empty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Empty) | method variant |
| [`require.(*Assertions).Emptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Emptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Empty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Empty](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_unary.go#L73)
{{% /tab %}}
{{< /tabs >}}

### Equal{#equal}
Equal asserts that two objects are equal.

Pointer variable equality is determined based on the equality of the
referenced values (as opposed to the memory addresses).

Function equality cannot be determined and will always fail.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Equal(t, 123, 123)
	success: 123, 123
	failure: 123, 456
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestEqual(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Equal(t, 123, 123)
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
// real-world test would inject *testing.T from TestEqual(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Equal(t, 123, 123)
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
| [`assert.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) | package-level function |
| [`assert.Equalf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equalf) | formatted variant |
| [`assert.(*Assertions).Equal(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Equal) | method variant |
| [`assert.(*Assertions).Equalf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Equalf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Equal) | package-level function |
| [`require.Equalf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Equalf) | formatted variant |
| [`require.(*Assertions).Equal(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Equal) | method variant |
| [`require.(*Assertions).Equalf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Equalf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Equal) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Equal](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L30)
{{% /tab %}}
{{< /tabs >}}

### EqualExportedValues{#equalexportedvalues}
EqualExportedValues asserts that the types of two objects are equal and their public
fields are also equal.

This is useful for comparing structs that have private fields that could potentially differ.

Function equality cannot be determined and will always fail.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	 type S struct {
		Exported     	int
		notExported   	int
	 }
	assertions.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
	assertions.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
	success: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}
	failure:  &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestEqualExportedValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.EqualExportedValues(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
	fmt.Printf("success: %t\n", success)

}

type dummyStruct struct {
	A string
	b int
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
// real-world test would inject *testing.T from TestEqualExportedValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.EqualExportedValues(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
	fmt.Println("passed")

}

type dummyStruct struct {
	A string
	b int
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
| [`assert.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualExportedValues) | package-level function |
| [`assert.EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualExportedValuesf) | formatted variant |
| [`assert.(*Assertions).EqualExportedValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualExportedValues) | method variant |
| [`assert.(*Assertions).EqualExportedValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualExportedValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualExportedValues) | package-level function |
| [`require.EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualExportedValuesf) | formatted variant |
| [`require.(*Assertions).EqualExportedValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualExportedValues) | method variant |
| [`require.(*Assertions).EqualExportedValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualExportedValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualExportedValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualExportedValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L214)
{{% /tab %}}
{{< /tabs >}}

### EqualT[V comparable] {{% icon icon="star" color=orange %}}{#equaltv-comparable}
EqualT asserts that two objects of the same comparable type are equal.

Pointer variable equality is determined based on the equality of the memory addresses (unlike [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal), but like [Same](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same)).

Functions, slices and maps are not comparable. See also [ComparisonOperators](https://go.dev/ref/spec#Comparison_operators.).

If you need to compare values of non-comparable types, or compare pointers by the value they point to,
use [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) instead.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.EqualT(t, 123, 123)
	success: 123, 123
	failure: 123, 456
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestEqualT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.EqualT(t, 123, 123)
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
// real-world test would inject *testing.T from TestEqualT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.EqualT(t, 123, 123)
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
| [`assert.EqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualT) | package-level function |
| [`assert.EqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualT) | package-level function |
| [`require.EqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.EqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualT](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L67)
{{% /tab %}}
{{< /tabs >}}

### EqualValues{#equalvalues}
EqualValues asserts that two objects are equal or convertible to the larger
type and equal.

Function equality cannot be determined and will always fail.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.EqualValues(t, uint32(123), int32(123))
	success: uint32(123), int32(123)
	failure: uint32(123), int32(456)
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestEqualValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.EqualValues(t, uint32(123), int32(123))
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
// real-world test would inject *testing.T from TestEqualValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.EqualValues(t, uint32(123), int32(123))
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
| [`assert.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualValues) | package-level function |
| [`assert.EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualValuesf) | formatted variant |
| [`assert.(*Assertions).EqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualValues) | method variant |
| [`assert.(*Assertions).EqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualValues) | package-level function |
| [`require.EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualValuesf) | formatted variant |
| [`require.(*Assertions).EqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualValues) | method variant |
| [`require.(*Assertions).EqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L142)
{{% /tab %}}
{{< /tabs >}}

### Exactly{#exactly}
Exactly asserts that two objects are equal in value and type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Exactly(t, int32(123), int64(123))
	success: int32(123), int32(123)
	failure: int32(123), int64(123)
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestExactly(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Exactly(t, int32(123), int32(123))
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
// real-world test would inject *testing.T from TestExactly(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Exactly(t, int32(123), int32(123))
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
| [`assert.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Exactly) | package-level function |
| [`assert.Exactlyf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Exactlyf) | formatted variant |
| [`assert.(*Assertions).Exactly(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Exactly) | method variant |
| [`assert.(*Assertions).Exactlyf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Exactlyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Exactly) | package-level function |
| [`require.Exactlyf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Exactlyf) | formatted variant |
| [`require.(*Assertions).Exactly(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Exactly) | method variant |
| [`require.(*Assertions).Exactlyf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Exactlyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Exactly) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Exactly](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L256)
{{% /tab %}}
{{< /tabs >}}

### Nil{#nil}
Nil asserts that the specified object is nil.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Nil(t, err)
	success: nil
	failure: "not nil"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNil(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Nil(t, nil)
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
// real-world test would inject *testing.T from TestNil(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Nil(t, nil)
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
| [`assert.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Nil) | package-level function |
| [`assert.Nilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Nilf) | formatted variant |
| [`assert.(*Assertions).Nil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Nil) | method variant |
| [`assert.(*Assertions).Nilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Nilf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Nil) | package-level function |
| [`require.Nilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Nilf) | formatted variant |
| [`require.(*Assertions).Nil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Nil) | method variant |
| [`require.(*Assertions).Nilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Nilf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Nil) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Nil](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_unary.go#L21)
{{% /tab %}}
{{< /tabs >}}

### NotEmpty{#notempty}
NotEmpty asserts that the specified object is NOT [Empty](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Empty).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	if assert.NotEmpty(t, obj) {
		assertions.Equal(t, "two", obj[1](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#1))
	}
	success: "not empty"
	failure: ""
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotEmpty(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotEmpty(t, "not empty")
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
// real-world test would inject *testing.T from TestNotEmpty(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotEmpty(t, "not empty")
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
| [`assert.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEmpty) | package-level function |
| [`assert.NotEmptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEmptyf) | formatted variant |
| [`assert.(*Assertions).NotEmpty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEmpty) | method variant |
| [`assert.(*Assertions).NotEmptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEmptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEmpty) | package-level function |
| [`require.NotEmptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEmptyf) | formatted variant |
| [`require.(*Assertions).NotEmpty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEmpty) | method variant |
| [`require.(*Assertions).NotEmptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEmptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEmpty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEmpty](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_unary.go#L98)
{{% /tab %}}
{{< /tabs >}}

### NotEqual{#notequal}
NotEqual asserts that the specified values are NOT equal.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotEqual(t, obj1, obj2)
Pointer variable equality is determined based on the equality of the
referenced values (as opposed to the memory addresses).
Function equality cannot be determined and will always fail.
	success: 123, 456
	failure: 123, 123
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotEqual(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotEqual(t, 123, 456)
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
// real-world test would inject *testing.T from TestNotEqual(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotEqual(t, 123, 456)
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
| [`assert.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqual) | package-level function |
| [`assert.NotEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualf) | formatted variant |
| [`assert.(*Assertions).NotEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqual) | method variant |
| [`assert.(*Assertions).NotEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqual) | package-level function |
| [`require.NotEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualf) | formatted variant |
| [`require.(*Assertions).NotEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqual) | method variant |
| [`require.(*Assertions).NotEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEqual) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L91)
{{% /tab %}}
{{< /tabs >}}

### NotEqualT[V comparable] {{% icon icon="star" color=orange %}}{#notequaltv-comparable}
NotEqualT asserts that the specified values of the same comparable type are NOT equal.

See [EqualT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotEqualT(t, obj1, obj2)
	success: 123, 456
	failure: 123, 123
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotEqualT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotEqualT(t, 123, 456)
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
// real-world test would inject *testing.T from TestNotEqualT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotEqualT(t, 123, 456)
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
| [`assert.NotEqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualT) | package-level function |
| [`assert.NotEqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualT) | package-level function |
| [`require.NotEqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotEqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEqualT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEqualT](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L120)
{{% /tab %}}
{{< /tabs >}}

### NotEqualValues{#notequalvalues}
NotEqualValues asserts that two objects are not equal even when converted to the same type.

Function equality cannot be determined and will always fail.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotEqualValues(t, obj1, obj2)
	success: uint32(123), int32(456)
	failure: uint32(123), int32(123)
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotEqualValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotEqualValues(t, uint32(123), int32(456))
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
// real-world test would inject *testing.T from TestNotEqualValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotEqualValues(t, uint32(123), int32(456))
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
| [`assert.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualValues) | package-level function |
| [`assert.NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualValuesf) | formatted variant |
| [`assert.(*Assertions).NotEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualValues) | method variant |
| [`assert.(*Assertions).NotEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualValues) | package-level function |
| [`require.NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualValuesf) | formatted variant |
| [`require.(*Assertions).NotEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualValues) | method variant |
| [`require.(*Assertions).NotEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L176)
{{% /tab %}}
{{< /tabs >}}

### NotNil{#notnil}
NotNil asserts that the specified object is not nil.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
assertions.NotNil(t, err)
	success: "not nil"
	failure: nil
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotNil(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotNil(t, "not nil")
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
// real-world test would inject *testing.T from TestNotNil(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotNil(t, "not nil")
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
| [`assert.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotNil) | package-level function |
| [`assert.NotNilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotNilf) | formatted variant |
| [`assert.(*Assertions).NotNil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotNil) | method variant |
| [`assert.(*Assertions).NotNilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotNilf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotNil) | package-level function |
| [`require.NotNilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotNilf) | formatted variant |
| [`require.(*Assertions).NotNil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotNil) | method variant |
| [`require.(*Assertions).NotNilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotNilf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotNil) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotNil](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_unary.go#L42)
{{% /tab %}}
{{< /tabs >}}

### NotSame{#notsame}
NotSame asserts that two pointers do not reference the same object.

See [Same](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSame(t, ptr1, ptr2)
	success: &staticVar, ptr("static string")
	failure: &staticVar, staticVarPtr
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotSame(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotSame(t, &staticVar, ptr("static string"))
	fmt.Printf("success: %t\n", success)

}

var staticVar = "static string"

func ptr[T any](value T) *T {
	p := value

	return &p
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
// real-world test would inject *testing.T from TestNotSame(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotSame(t, &staticVar, ptr("static string"))
	fmt.Println("passed")

}

var staticVar = "static string"

func ptr[T any](value T) *T {
	p := value

	return &p
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
| [`assert.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSame) | package-level function |
| [`assert.NotSamef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSamef) | formatted variant |
| [`assert.(*Assertions).NotSame(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSame) | method variant |
| [`assert.(*Assertions).NotSamef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSamef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSame) | package-level function |
| [`require.NotSamef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSamef) | formatted variant |
| [`require.(*Assertions).NotSame(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSame) | method variant |
| [`require.(*Assertions).NotSamef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSamef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSame) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSame](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_pointer.go#L84)
{{% /tab %}}
{{< /tabs >}}

### NotSameT[P any] {{% icon icon="star" color=orange %}}{#notsametp-any}
NotSameT asserts that two pointers do not reference the same object.

See [SameT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SameT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSameT(t, ptr1, ptr2)
	success: &staticVar, ptr("static string")
	failure: &staticVar, staticVarPtr
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotSameT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotSameT(t, &staticVar, ptr("static string"))
	fmt.Printf("success: %t\n", success)

}

var staticVar = "static string"

func ptr[T any](value T) *T {
	p := value

	return &p
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
// real-world test would inject *testing.T from TestNotSameT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotSameT(t, &staticVar, ptr("static string"))
	fmt.Println("passed")

}

var staticVar = "static string"

func ptr[T any](value T) *T {
	p := value

	return &p
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
| [`assert.NotSameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSameT) | package-level function |
| [`assert.NotSameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSameTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotSameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSameT) | package-level function |
| [`require.NotSameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSameTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotSameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSameT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSameT](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_pointer.go#L116)
{{% /tab %}}
{{< /tabs >}}

### Same{#same}
Same asserts that two pointers reference the same object.

Both arguments must be pointer variables.

Pointer variable sameness is determined based on the equality of both type and value.

Unlike [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) pointers, [Same](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same) pointers point to the same memory address.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Same(t, ptr1, ptr2)
	success: &staticVar, staticVarPtr
	failure: &staticVar, ptr("static string")
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSame(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Same(t, &staticVar, staticVarPtr)
	fmt.Printf("success: %t\n", success)

}

var (
	staticVar    = "static string"
	staticVarPtr = &staticVar
)

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSame(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Same(t, &staticVar, staticVarPtr)
	fmt.Println("passed")

}

var (
	staticVar    = "static string"
	staticVarPtr = &staticVar
)

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
| [`assert.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same) | package-level function |
| [`assert.Samef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Samef) | formatted variant |
| [`assert.(*Assertions).Same(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Same) | method variant |
| [`assert.(*Assertions).Samef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Samef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Same) | package-level function |
| [`require.Samef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Samef) | formatted variant |
| [`require.(*Assertions).Same(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Same) | method variant |
| [`require.(*Assertions).Samef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Samef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Same) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Same](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_pointer.go#L24)
{{% /tab %}}
{{< /tabs >}}

### SameT[P any] {{% icon icon="star" color=orange %}}{#sametp-any}
SameT asserts that two pointers of the same type reference the same object.

See [Same](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.SameT(t, ptr1, ptr2)
	success: &staticVar, staticVarPtr
	failure: &staticVar, ptr("static string")
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSameT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.SameT(t, &staticVar, staticVarPtr)
	fmt.Printf("success: %t\n", success)

}

var (
	staticVar    = "static string"
	staticVarPtr = &staticVar
)

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestSameT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.SameT(t, &staticVar, staticVarPtr)
	fmt.Println("passed")

}

var (
	staticVar    = "static string"
	staticVarPtr = &staticVar
)

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
| [`assert.SameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SameT) | package-level function |
| [`assert.SameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#SameTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.SameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SameT) | package-level function |
| [`require.SameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#SameTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.SameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#SameT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#SameT](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal_pointer.go#L57)
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
