---
title: "Panic"
description: "Asserting A Panic Behavior"
modified: 2026-01-11
weight: 12
domains:
  - "panic"
keywords:
  - "NotPanics"
  - "NotPanicsf"
  - "Panics"
  - "Panicsf"
  - "PanicsWithError"
  - "PanicsWithErrorf"
  - "PanicsWithValue"
  - "PanicsWithValuef"
---

Asserting A Panic Behavior

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.

### NotPanics

NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotPanics(t, func(){ RemainCalm() })
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: func() { }
	failure: func() { panic("panicking") }
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotPanics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotPanics) | package-level function |
| [`assert.NotPanicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotPanicsf) | formatted variant |
| [`assert.(*Assertions).NotPanics(f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotPanics) | method variant |
| [`assert.(*Assertions).NotPanicsf(f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotPanicsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotPanics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotPanics) | package-level function |
| [`require.NotPanicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotPanicsf) | formatted variant |
| [`require.(*Assertions).NotPanics(f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotPanics) | method variant |
| [`require.(*Assertions).NotPanicsf(f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotPanicsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotPanics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotPanics) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotPanics](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L116)
{{% /tab %}}
{{< /tabs >}}

### Panics

Panics asserts that the code inside the specified PanicTestFunc panics.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Panics(t, func(){ GoCrazy() })
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: func() { panic("panicking") }
	failure: func() { }
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Panics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Panics) | package-level function |
| [`assert.Panicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Panicsf) | formatted variant |
| [`assert.(*Assertions).Panics(f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Panics) | method variant |
| [`assert.(*Assertions).Panicsf(f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Panicsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Panics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Panics) | package-level function |
| [`require.Panicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Panicsf) | formatted variant |
| [`require.(*Assertions).Panics(f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Panics) | method variant |
| [`require.(*Assertions).Panicsf(f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Panicsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Panics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Panics) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Panics](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L29)
{{% /tab %}}
{{< /tabs >}}

### PanicsWithError

PanicsWithError asserts that the code inside the specified PanicTestFunc
panics, and that the recovered panic value is an error that satisfies the
EqualError comparison.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest.Error(), func() { panic(ErrTest) }
	failure: ErrTest.Error(), func() { }
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.PanicsWithError(t T, errString string, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithError) | package-level function |
| [`assert.PanicsWithErrorf(t T, errString string, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithErrorf) | formatted variant |
| [`assert.(*Assertions).PanicsWithError(errString string, f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithError) | method variant |
| [`assert.(*Assertions).PanicsWithErrorf(errString string, f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithErrorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.PanicsWithError(t T, errString string, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithError) | package-level function |
| [`require.PanicsWithErrorf(t T, errString string, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithErrorf) | formatted variant |
| [`require.(*Assertions).PanicsWithError(errString string, f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithError) | method variant |
| [`require.(*Assertions).PanicsWithErrorf(errString string, f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithErrorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.PanicsWithError(t T, errString string, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#PanicsWithError) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#PanicsWithError](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L82)
{{% /tab %}}
{{< /tabs >}}

### PanicsWithValue

PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
the recovered panic value equals the expected panic value.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: "panicking", func() { panic("panicking") }
	failure: "panicking", func() { }
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.PanicsWithValue(t T, expected any, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithValue) | package-level function |
| [`assert.PanicsWithValuef(t T, expected any, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithValuef) | formatted variant |
| [`assert.(*Assertions).PanicsWithValue(expected any, f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithValue) | method variant |
| [`assert.(*Assertions).PanicsWithValuef(expected any, f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithValuef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.PanicsWithValue(t T, expected any, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithValue) | package-level function |
| [`require.PanicsWithValuef(t T, expected any, f assertions.PanicTestFunc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithValuef) | formatted variant |
| [`require.(*Assertions).PanicsWithValue(expected any, f assertions.PanicTestFunc) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithValue) | method variant |
| [`require.(*Assertions).PanicsWithValuef(expected any, f assertions.PanicTestFunc, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithValuef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.PanicsWithValue(t T, expected any, f assertions.PanicTestFunc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#PanicsWithValue) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#PanicsWithValue](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L53)
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
