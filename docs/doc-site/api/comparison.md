---
title: "Comparison"
description: "Comparing Ordered Values"
modified: 2026-01-18
weight: 3
domains:
  - "comparison"
keywords:
  - "Greater"
  - "Greaterf"
  - "GreaterOrEqual"
  - "GreaterOrEqualf"
  - "GreaterOrEqualT"
  - "GreaterOrEqualTf"
  - "GreaterT"
  - "GreaterTf"
  - "Less"
  - "Lessf"
  - "LessOrEqual"
  - "LessOrEqualf"
  - "LessOrEqualT"
  - "LessOrEqualTf"
  - "LessT"
  - "LessTf"
  - "Negative"
  - "Negativef"
  - "NegativeT"
  - "NegativeTf"
  - "Positive"
  - "Positivef"
  - "PositiveT"
  - "PositiveTf"
---

Comparing Ordered Values

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 12 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}

```tree
- [Greater](#greater) | angles-right
- [GreaterOrEqual](#greaterorequal) | angles-right
- [GreaterOrEqualT[Orderable Ordered]](#greaterorequaltorderable-ordered) | star | orange
- [GreaterT[Orderable Ordered]](#greatertorderable-ordered) | star | orange
- [Less](#less) | angles-right
- [LessOrEqual](#lessorequal) | angles-right
- [LessOrEqualT[Orderable Ordered]](#lessorequaltorderable-ordered) | star | orange
- [LessT[Orderable Ordered]](#lesstorderable-ordered) | star | orange
- [Negative](#negative) | angles-right
- [NegativeT[SignedNumber SignedNumeric]](#negativetsignednumber-signednumeric) | star | orange
- [Positive](#positive) | angles-right
- [PositiveT[SignedNumber SignedNumeric]](#positivetsignednumber-signednumeric) | star | orange
```

### Greater{#greater}

Greater asserts that the first element is strictly greater than the second.

Both elements must be of the same type in the [reflect.Kind](https://pkg.go.dev/reflect#Kind) sense.
To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Greater](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L30)
{{% /tab %}}
{{< /tabs >}}

### GreaterOrEqual{#greaterorequal}

GreaterOrEqual asserts that the first element is greater than or equal to the second.

See also [Greater].

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#GreaterOrEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L91)
{{% /tab %}}
{{< /tabs >}}

### GreaterOrEqualT[Orderable Ordered] {{% icon icon="star" color=orange %}}{#greaterorequaltorderable-ordered}

GreaterOrEqualT asserts that for two elements of the same type,
the first element is greater than or equal to the second.

The [Ordered] type can be any of Go's [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) (strings, numeric types),
[]byte (uses [bytes.Compare](https://pkg.go.dev/bytes#Compare)) and [time.Time](https://pkg.go.dev/time#Time) (uses [time.Time.Compare](https://pkg.go.dev/time#Time.Compare).

Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterOrEqualT] with [*time.Time].

[GreaterOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
use [GreaterOrEqual] instead.

To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.GreaterOrEqualT(t, 2, 1)
	assertions.GreaterOrEqualT(t, 2, 2)
	assertions.GreaterOrEqualT(t, "b", "a")
	assertions.GreaterOrEqualT(t, "b", "b")
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
| [`assert.GreaterOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterOrEqualT) | package-level function |
| [`assert.GreaterOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterOrEqualTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.GreaterOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterOrEqualT) | package-level function |
| [`require.GreaterOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterOrEqualTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.GreaterOrEqualT(t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#GreaterOrEqualT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#GreaterOrEqualT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L124)
{{% /tab %}}
{{< /tabs >}}

### GreaterT[Orderable Ordered] {{% icon icon="star" color=orange %}}{#greatertorderable-ordered}

GreaterT asserts that for two elements of the same type,
the first element is strictly greater than the second.

The [Ordered] type can be any of Go's [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) (strings, numeric types),
[]byte (uses [bytes.Compare](https://pkg.go.dev/bytes#Compare)) and [time.Time](https://pkg.go.dev/time#Time) (uses [time.Time.Compare](https://pkg.go.dev/time#Time.Compare).

Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterT] with [*time.Time].

[GreaterT] ensures type safety at build time. If you need to compare values with a dynamically assigned type, use [Greater] instead.

To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.GreaterT(t, 2, 1)
	assertions.GreaterT(t, float64(2), float64(1))
	assertions.GreaterT(t, "b", "a")
	assertions.GreaterT(t, time.Date(2026,1,1,0,0,0,0,nil), time.Now())
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
| [`assert.GreaterT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterT) | package-level function |
| [`assert.GreaterTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#GreaterTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.GreaterT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterT) | package-level function |
| [`require.GreaterTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#GreaterTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.GreaterT(t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#GreaterT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#GreaterT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L62)
{{% /tab %}}
{{< /tabs >}}

### Less{#less}

Less asserts that the first element is strictly less than the second.

Both elements must be of the same type in the [reflect.Kind](https://pkg.go.dev/reflect#Kind) sense.
To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Less](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L153)
{{% /tab %}}
{{< /tabs >}}

### LessOrEqual{#lessorequal}

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#LessOrEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L211)
{{% /tab %}}
{{< /tabs >}}

### LessOrEqualT[Orderable Ordered] {{% icon icon="star" color=orange %}}{#lessorequaltorderable-ordered}

LessOrEqualT asserts that for two elements of the same type, the first element is less than or equal to the second.

The [Ordered] type can be any of Go's [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) (strings, numeric types),
[]byte (uses [bytes.Compare](https://pkg.go.dev/bytes#Compare)) and [time.Time](https://pkg.go.dev/time#Time) (uses [time.Time.Compare](https://pkg.go.dev/time#Time.Compare).

Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessOrEqualT] with [*time.Time].

[LessOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
use [LessOrEqual] instead.

To compare values that need a type conversion (e.g. float32 against float64), you should use [LessOrEqual] instead.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.LessOrEqualT(t, 1, 2)
	assertions.LessOrEqualT(t, 2, 2)
	assertions.LessOrEqualT(t, "a", "b")
	assertions.LessOrEqualT(t, "b", "b")
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
| [`assert.LessOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessOrEqualT) | package-level function |
| [`assert.LessOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessOrEqualTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.LessOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessOrEqualT) | package-level function |
| [`require.LessOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessOrEqualTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.LessOrEqualT(t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#LessOrEqualT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#LessOrEqualT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L243)
{{% /tab %}}
{{< /tabs >}}

### LessT[Orderable Ordered] {{% icon icon="star" color=orange %}}{#lesstorderable-ordered}

LessT asserts that for two elements of the same type, the first element is strictly less than the second.

The [Ordered] type can be any of Go's [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) (strings, numeric types),
[]byte (uses [bytes.Compare](https://pkg.go.dev/bytes#Compare)) and [time.Time](https://pkg.go.dev/time#Time) (uses [time.Time.Compare](https://pkg.go.dev/time#Time.Compare).

Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessT] with [*time.Time].

[LessT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
use [Less] instead.

To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.LessT(t, 1, 2)
	assertions.LessT(t, float64(1), float64(2))
	assertions.LessT(t, "a", "b")
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
| [`assert.LessT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessT) | package-level function |
| [`assert.LessTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#LessTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.LessT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessT) | package-level function |
| [`require.LessTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#LessTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.LessT(t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#LessT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#LessT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L184)
{{% /tab %}}
{{< /tabs >}}

### Negative{#negative}

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Negative](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L314)
{{% /tab %}}
{{< /tabs >}}

### NegativeT[SignedNumber SignedNumeric] {{% icon icon="star" color=orange %}}{#negativetsignednumber-signednumeric}

NegativeT asserts that the specified element of a signed numeric type is strictly negative.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NegativeT(t, -1)
	assertions.NegativeT(t, -1.23)
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
| [`assert.NegativeT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NegativeT) | package-level function |
| [`assert.NegativeTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NegativeTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NegativeT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NegativeT) | package-level function |
| [`require.NegativeTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NegativeTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NegativeT(t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NegativeT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NegativeT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L335)
{{% /tab %}}
{{< /tabs >}}

### Positive{#positive}

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Positive](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L268)
{{% /tab %}}
{{< /tabs >}}

### PositiveT[SignedNumber SignedNumeric] {{% icon icon="star" color=orange %}}{#positivetsignednumber-signednumeric}

PositiveT asserts that the specified element of a signed numeric type is strictly positive.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.PositiveT(t, 1)
	assertions.PositiveT(t, 1.23)
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
| [`assert.PositiveT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PositiveT) | package-level function |
| [`assert.PositiveTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PositiveTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.PositiveT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PositiveT) | package-level function |
| [`require.PositiveTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PositiveTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.PositiveT(t T, e SignedNumber, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#PositiveT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#PositiveT](https://github.com/go-openapi/testify/blob/master/internal/assertions/compare.go#L289)
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

Generated on 2026-01-18 (version e12affe) using codegen version v2.1.9-0.20260118112101-e12affef2419+dirty [sha: e12affef24198e72ee13eb6d25018d2c3232629f]
-->
