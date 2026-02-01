---
title: "Panic"
description: "Asserting A Panic Behavior"
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

```tree
- [NotPanics](#notpanics) | angles-right
- [Panics](#panics) | angles-right
- [PanicsWithError](#panicswitherror) | angles-right
- [PanicsWithValue](#panicswithvalue) | angles-right
```

### NotPanics{#notpanics}

NotPanics asserts that the code inside the specified function does NOT panic.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotPanics(t, func(){ RemainCalm() })
	success: func() { }
	failure: func() { panic("panicking") }
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotPanics(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.NotPanics(t, func() {
	})
	fmt.Printf("success: %t\n", success)

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
| [`assert.NotPanics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotPanics) | package-level function |
| [`assert.NotPanicsf(t T, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotPanicsf) | formatted variant |
| [`assert.(*Assertions).NotPanics(f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotPanics) | method variant |
| [`assert.(*Assertions).NotPanicsf(f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotPanicsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotPanics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotPanics) | package-level function |
| [`require.NotPanicsf(t T, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotPanicsf) | formatted variant |
| [`require.(*Assertions).NotPanics(f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotPanics) | method variant |
| [`require.(*Assertions).NotPanicsf(f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotPanicsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotPanics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotPanics) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotPanics](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L111)
{{% /tab %}}
{{< /tabs >}}

### Panics{#panics}

Panics asserts that the code inside the specified function panics.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Panics(t, func(){ GoCrazy() })
	success: func() { panic("panicking") }
	failure: func() { }
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestPanics(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.Panics(t, func() {
		panic("panicking")
	})
	fmt.Printf("success: %t\n", success)

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
| [`assert.Panics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Panics) | package-level function |
| [`assert.Panicsf(t T, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Panicsf) | formatted variant |
| [`assert.(*Assertions).Panics(f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Panics) | method variant |
| [`assert.(*Assertions).Panicsf(f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Panicsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Panics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Panics) | package-level function |
| [`require.Panicsf(t T, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Panicsf) | formatted variant |
| [`require.(*Assertions).Panics(f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Panics) | method variant |
| [`require.(*Assertions).Panicsf(f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Panicsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Panics(t T, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Panics) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Panics](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L25)
{{% /tab %}}
{{< /tabs >}}

### PanicsWithError{#panicswitherror}

PanicsWithError asserts that the code inside the specified function panics,
and that the recovered panic value is an error that satisfies the EqualError comparison.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
	success: ErrTest.Error(), func() { panic(ErrTest) }
	failure: ErrTest.Error(), func() { }
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestPanicsWithError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.PanicsWithError(t, assert.ErrTest.Error(), func() {
		panic(assert.ErrTest)
	})
	fmt.Printf("success: %t\n", success)

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
| [`assert.PanicsWithError(t T, errString string, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithError) | package-level function |
| [`assert.PanicsWithErrorf(t T, errString string, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithErrorf) | formatted variant |
| [`assert.(*Assertions).PanicsWithError(errString string, f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithError) | method variant |
| [`assert.(*Assertions).PanicsWithErrorf(errString string, f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithErrorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.PanicsWithError(t T, errString string, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithError) | package-level function |
| [`require.PanicsWithErrorf(t T, errString string, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithErrorf) | formatted variant |
| [`require.(*Assertions).PanicsWithError(errString string, f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithError) | method variant |
| [`require.(*Assertions).PanicsWithErrorf(errString string, f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithErrorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.PanicsWithError(t T, errString string, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#PanicsWithError) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#PanicsWithError](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L77)
{{% /tab %}}
{{< /tabs >}}

### PanicsWithValue{#panicswithvalue}

PanicsWithValue asserts that the code inside the specified function panics,
and that the recovered panic value equals the expected panic value.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
	success: "panicking", func() { panic("panicking") }
	failure: "panicking", func() { }
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestPanicsWithValue(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.PanicsWithValue(t, "panicking", func() {
		panic("panicking")
	})
	fmt.Printf("success: %t\n", success)

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
| [`assert.PanicsWithValue(t T, expected any, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithValue) | package-level function |
| [`assert.PanicsWithValuef(t T, expected any, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#PanicsWithValuef) | formatted variant |
| [`assert.(*Assertions).PanicsWithValue(expected any, f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithValue) | method variant |
| [`assert.(*Assertions).PanicsWithValuef(expected any, f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.PanicsWithValuef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.PanicsWithValue(t T, expected any, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithValue) | package-level function |
| [`require.PanicsWithValuef(t T, expected any, f func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#PanicsWithValuef) | formatted variant |
| [`require.(*Assertions).PanicsWithValue(expected any, f func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithValue) | method variant |
| [`require.(*Assertions).PanicsWithValuef(expected any, f func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.PanicsWithValuef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.PanicsWithValue(t T, expected any, f func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#PanicsWithValue) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#PanicsWithValue](https://github.com/go-openapi/testify/blob/master/internal/assertions/panic.go#L49)
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
