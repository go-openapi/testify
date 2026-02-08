---
title: "Number"
description: "Asserting Numbers"
weight: 10
domains:
  - "number"
keywords:
  - "InDelta"
  - "InDeltaf"
  - "InDeltaMapValues"
  - "InDeltaMapValuesf"
  - "InDeltaSlice"
  - "InDeltaSlicef"
  - "InDeltaT"
  - "InDeltaTf"
  - "InEpsilon"
  - "InEpsilonf"
  - "InEpsilonSlice"
  - "InEpsilonSlicef"
  - "InEpsilonT"
  - "InEpsilonTf"
---

Asserting Numbers

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 7 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [InDelta](#indelta) | angles-right
- [InDeltaMapValues](#indeltamapvalues) | angles-right
- [InDeltaSlice](#indeltaslice) | angles-right
- [InDeltaT[Number Measurable]](#indeltatnumber-measurable) | star | orange
- [InEpsilon](#inepsilon) | angles-right
- [InEpsilonSlice](#inepsilonslice) | angles-right
- [InEpsilonT[Number Measurable]](#inepsilontnumber-measurable) | star | orange
```

### InDelta{#indelta}
InDelta asserts that the two numerals are within delta of each other.

Delta must be greater than or equal to zero.

Expected and actual values should convert to float64.
To compare large integers that can't be represented accurately as float64 (eg. uint64),
prefer [InDeltaT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT) to preserve the original type.

#### Behavior with IEEE floating point arithmetics

  - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
  - expected +Inf is matched only by a +Inf
  - expected -Inf is matched only by a -Inf

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
assertions.InDelta(t, math.Pi, 22/7.0, 0.01)
	success: 1.0, 1.01, 0.02
	failure: 1.0, 1.1, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInDelta(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDelta(t *testing.T)
	success := assert.InDelta(t, 1.0, 1.01, 0.02)
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
// real-world test would inject *testing.T from TestInDelta(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDelta(t *testing.T)
	require.InDelta(t, 1.0, 1.01, 0.02)
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
| [`assert.InDelta(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta) | package-level function |
| [`assert.InDeltaf(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaf) | formatted variant |
| [`assert.(*Assertions).InDelta(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDelta) | method variant |
| [`assert.(*Assertions).InDeltaf(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDeltaf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InDelta(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDelta) | package-level function |
| [`require.InDeltaf(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaf) | formatted variant |
| [`require.(*Assertions).InDelta(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDelta) | method variant |
| [`require.(*Assertions).InDeltaf(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDeltaf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InDelta(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InDelta) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDelta](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L35)
{{% /tab %}}
{{< /tabs >}}

### InDeltaMapValues{#indeltamapvalues}
InDeltaMapValues is the same as [InDelta](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta), but it compares all values between two maps. Both maps must have exactly the same keys.

See [InDelta](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InDeltaMapValues(t, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.0}, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.01}, 0.02)
	success: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.0}, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.01}, 0.02
	failure: map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.0}, map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)float64{"a": 1.1}, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInDeltaMapValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaMapValues(t *testing.T)
	success := assert.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
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
// real-world test would inject *testing.T from TestInDeltaMapValues(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaMapValues(t *testing.T)
	require.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
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
| [`assert.InDeltaMapValues(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaMapValues) | package-level function |
| [`assert.InDeltaMapValuesf(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaMapValuesf) | formatted variant |
| [`assert.(*Assertions).InDeltaMapValues(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDeltaMapValues) | method variant |
| [`assert.(*Assertions).InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDeltaMapValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InDeltaMapValues(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaMapValues) | package-level function |
| [`require.InDeltaMapValuesf(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaMapValuesf) | formatted variant |
| [`require.(*Assertions).InDeltaMapValues(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDeltaMapValues) | method variant |
| [`require.(*Assertions).InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDeltaMapValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InDeltaMapValues(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InDeltaMapValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDeltaMapValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L271)
{{% /tab %}}
{{< /tabs >}}

### InDeltaSlice{#indeltaslice}
InDeltaSlice is the same as [InDelta](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta), except it compares two slices.

See [InDelta](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
	success: []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02
	failure: []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInDeltaSlice(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaSlice(t *testing.T)
	success := assert.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
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
// real-world test would inject *testing.T from TestInDeltaSlice(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaSlice(t *testing.T)
	require.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
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
| [`assert.InDeltaSlice(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaSlice) | package-level function |
| [`assert.InDeltaSlicef(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaSlicef) | formatted variant |
| [`assert.(*Assertions).InDeltaSlice(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDeltaSlice) | method variant |
| [`assert.(*Assertions).InDeltaSlicef(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InDeltaSlicef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InDeltaSlice(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaSlice) | package-level function |
| [`require.InDeltaSlicef(t T, expected any, actual any, delta float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaSlicef) | formatted variant |
| [`require.(*Assertions).InDeltaSlice(expected any, actual any, delta float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDeltaSlice) | method variant |
| [`require.(*Assertions).InDeltaSlicef(expected any, actual any, delta float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InDeltaSlicef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InDeltaSlice(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InDeltaSlice) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDeltaSlice](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L235)
{{% /tab %}}
{{< /tabs >}}

### InDeltaT[Number Measurable] {{% icon icon="star" color=orange %}}{#indeltatnumber-measurable}
InDeltaT asserts that the two numerals of the same type numerical type are within delta of each other.

[InDeltaT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT) accepts any go numeric type, including integer types.

The main difference with [InDelta](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDelta) is that the delta is expressed with the same type as the values, not necessarily a float64.

Delta must be greater than or equal to zero.

#### Behavior with IEEE floating point arithmetics

  - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
  - expected +Inf is matched only by a +Inf
  - expected -Inf is matched only by a -Inf

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
assertions.InDeltaT(t, math.Pi, 22/7.0, 0.01)
	success: 1.0, 1.01, 0.02
	failure: 1.0, 1.1, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInDeltaT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaT(t *testing.T)
	success := assert.InDeltaT(t, 1.0, 1.01, 0.02)
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
// real-world test would inject *testing.T from TestInDeltaT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaT(t *testing.T)
	require.InDeltaT(t, 1.0, 1.01, 0.02)
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
| [`assert.InDeltaT[Number Measurable](t T, expected Number, actual Number, delta Number, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT) | package-level function |
| [`assert.InDeltaTf[Number Measurable](t T, expected Number, actual Number, delta Number, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InDeltaT[Number Measurable](t T, expected Number, actual Number, delta Number, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaT) | package-level function |
| [`require.InDeltaTf[Number Measurable](t T, expected Number, actual Number, delta Number, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InDeltaTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InDeltaT[Number Measurable](t T, expected Number, actual Number, delta Number, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InDeltaT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDeltaT](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L85)
{{% /tab %}}
{{< /tabs >}}

### InEpsilon{#inepsilon}
InEpsilon asserts that expected and actual have a relative error less than epsilon.

#### Behavior with IEEE floating point arithmetics

  - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
  - expected +Inf is matched only by a +Inf
  - expected -Inf is matched only by a -Inf

Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT).

Formula:
  - If expected == 0: fail if |actual - expected| > epsilon
  - If expected != 0: fail if |actual - expected| > epsilon * |expected|

This allows [InEpsilonT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonT) to work naturally across the full numeric range including zero.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
	success: 100.0, 101.0, 0.02
	failure: 100.0, 110.0, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInEpsilon(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilon(t *testing.T)
	success := assert.InEpsilon(t, 100.0, 101.0, 0.02)
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
// real-world test would inject *testing.T from TestInEpsilon(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilon(t *testing.T)
	require.InEpsilon(t, 100.0, 101.0, 0.02)
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
| [`assert.InEpsilon(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilon) | package-level function |
| [`assert.InEpsilonf(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonf) | formatted variant |
| [`assert.(*Assertions).InEpsilon(expected any, actual any, epsilon float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InEpsilon) | method variant |
| [`assert.(*Assertions).InEpsilonf(expected any, actual any, epsilon float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InEpsilonf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InEpsilon(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilon) | package-level function |
| [`require.InEpsilonf(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilonf) | formatted variant |
| [`require.(*Assertions).InEpsilon(expected any, actual any, epsilon float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InEpsilon) | method variant |
| [`require.(*Assertions).InEpsilonf(expected any, actual any, epsilon float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InEpsilonf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InEpsilon(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InEpsilon) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InEpsilon](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L142)
{{% /tab %}}
{{< /tabs >}}

### InEpsilonSlice{#inepsilonslice}
InEpsilonSlice is the same as [InEpsilon](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilon), except it compares each value from two slices.

See [InEpsilon](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilon).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
	success: []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02
	failure: []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInEpsilonSlice(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonSlice(t *testing.T)
	success := assert.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
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
// real-world test would inject *testing.T from TestInEpsilonSlice(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonSlice(t *testing.T)
	require.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
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
| [`assert.InEpsilonSlice(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonSlice) | package-level function |
| [`assert.InEpsilonSlicef(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonSlicef) | formatted variant |
| [`assert.(*Assertions).InEpsilonSlice(expected any, actual any, epsilon float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InEpsilonSlice) | method variant |
| [`assert.(*Assertions).InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.InEpsilonSlicef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InEpsilonSlice(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilonSlice) | package-level function |
| [`require.InEpsilonSlicef(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilonSlicef) | formatted variant |
| [`require.(*Assertions).InEpsilonSlice(expected any, actual any, epsilon float64) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InEpsilonSlice) | method variant |
| [`require.(*Assertions).InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.InEpsilonSlicef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InEpsilonSlice(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InEpsilonSlice) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InEpsilonSlice](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L326)
{{% /tab %}}
{{< /tabs >}}

### InEpsilonT[Number Measurable] {{% icon icon="star" color=orange %}}{#inepsilontnumber-measurable}
InEpsilonT asserts that expected and actual have a relative error less than epsilon.

When expected is zero, epsilon is interpreted as an absolute error threshold,
since relative error is mathematically undefined for zero values.

Unlike [InDeltaT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT), which preserves the original type, [InEpsilonT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonT) converts the expected and actual
numbers to float64, since the relative error doesn't make sense as an integer.

#### Behavior with IEEE floating point arithmetics

  - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
  - expected +Inf is matched only by a +Inf
  - expected -Inf is matched only by a -Inf

Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InDeltaT).

Formula:
  - If expected == 0: fail if |actual - expected| > epsilon
  - If expected != 0: fail if |actual - expected| > epsilon * |expected|

This allows [InEpsilonT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonT) to work naturally across the full numeric range including zero.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
	success: 100.0, 101.0, 0.02
	failure: 100.0, 110.0, 0.05
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestInEpsilonT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonT(t *testing.T)
	success := assert.InEpsilonT(t, 100.0, 101.0, 0.02)
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
// real-world test would inject *testing.T from TestInEpsilonT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonT(t *testing.T)
	require.InEpsilonT(t, 100.0, 101.0, 0.02)
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
| [`assert.InEpsilonT[Number Measurable](t T, expected Number, actual Number, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonT) | package-level function |
| [`assert.InEpsilonTf[Number Measurable](t T, expected Number, actual Number, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#InEpsilonTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.InEpsilonT[Number Measurable](t T, expected Number, actual Number, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilonT) | package-level function |
| [`require.InEpsilonTf[Number Measurable](t T, expected Number, actual Number, epsilon float64, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#InEpsilonTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.InEpsilonT[Number Measurable](t T, expected Number, actual Number, epsilon float64, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#InEpsilonT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InEpsilonT](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L199)
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
