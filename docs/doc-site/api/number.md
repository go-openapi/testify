---
title: "Number"
description: "Asserting Numbers"
modified: 2026-01-11
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
  - "InEpsilon"
  - "InEpsilonf"
  - "InEpsilonSlice"
  - "InEpsilonSlicef"
---

Asserting Numbers

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 5 functionalities.

### InDelta

InDelta asserts that the two numerals are within delta of each other.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
assertions.InDelta(t, math.Pi, 22/7.0, 0.01)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1.0, 1.01, 0.02
	failure: 1.0, 1.1, 0.05
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDelta](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L24)
{{% /tab %}}
{{< /tabs >}}

### InDeltaMapValues

InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02
	failure: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDeltaMapValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L101)
{{% /tab %}}
{{< /tabs >}}

### InDeltaSlice

InDeltaSlice is the same as InDelta, except it compares two slices.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02
	failure: []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InDeltaSlice](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L67)
{{% /tab %}}
{{< /tabs >}}

### InEpsilon

InEpsilon asserts that expected and actual have a relative error less than epsilon.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 100.0, 101.0, 0.02
	failure: 100.0, 110.0, 0.05
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InEpsilon](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L155)
{{% /tab %}}
{{< /tabs >}}

### InEpsilonSlice

InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02
	failure: []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#InEpsilonSlice](https://github.com/go-openapi/testify/blob/master/internal/assertions/number.go#L188)
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
