---
title: "Boolean"
description: "Asserting Boolean Values"
weight: 1
domains:
  - "boolean"
keywords:
  - "False"
  - "Falsef"
  - "FalseT"
  - "FalseTf"
  - "True"
  - "Truef"
  - "TrueT"
  - "TrueTf"
---

Asserting Boolean Values

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [False](#false) | angles-right
- [FalseT[B Boolean]](#falsetb-boolean) | star | orange
- [True](#true) | angles-right
- [TrueT[B Boolean]](#truetb-boolean) | star | orange
```

### False{#false}

False asserts that the specified value is false.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.False(t, myBool)
	success: 1 == 0
	failure: 1 == 1
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestFalse(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.False(t, 1 == 0)
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
| [`assert.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#False) | package-level function |
| [`assert.Falsef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Falsef) | formatted variant |
| [`assert.(*Assertions).False(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.False) | method variant |
| [`assert.(*Assertions).Falsef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Falsef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#False) | package-level function |
| [`require.Falsef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Falsef) | formatted variant |
| [`require.(*Assertions).False(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.False) | method variant |
| [`require.(*Assertions).Falsef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Falsef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.False(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#False) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#False](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L65)
{{% /tab %}}
{{< /tabs >}}

### FalseT[B Boolean] {{% icon icon="star" color=orange %}}{#falsetb-boolean}

FalseT asserts that the specified value is false.

The type constraint [Boolean](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Boolean) accepts any type which underlying type is bool.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	 type B bool
	 var b B = true
		assertions.FalseT(t, b)
	success: 1 == 0
	failure: 1 == 1
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestFalseT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.FalseT(t, 1 == 0)
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
| [`assert.FalseT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FalseT) | package-level function |
| [`assert.FalseTf[B Boolean](t T, value B, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FalseTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.FalseT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FalseT) | package-level function |
| [`require.FalseTf[B Boolean](t T, value B, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FalseTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.FalseT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#FalseT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#FalseT](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L92)
{{% /tab %}}
{{< /tabs >}}

### True{#true}

True asserts that the specified value is true.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.True(t, myBool)
	success: 1 == 1
	failure: 1 == 0
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestTrue(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.True(t, 1 == 1)
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
| [`assert.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#True) | package-level function |
| [`assert.Truef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Truef) | formatted variant |
| [`assert.(*Assertions).True(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.True) | method variant |
| [`assert.(*Assertions).Truef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Truef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#True) | package-level function |
| [`require.Truef(t T, value bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Truef) | formatted variant |
| [`require.(*Assertions).True(value bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.True) | method variant |
| [`require.(*Assertions).Truef(value bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Truef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.True(t T, value bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#True) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#True](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L16)
{{% /tab %}}
{{< /tabs >}}

### TrueT[B Boolean] {{% icon icon="star" color=orange %}}{#truetb-boolean}

TrueT asserts that the specified value is true.

The type constraint [Boolean](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Boolean) accepts any type which underlying type is bool.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	type B bool
	var b B = true
	assertions.True(t, b)
	success: 1 == 1
	failure: 1 == 0
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestTrueT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T)
	success := assert.TrueT(t, 1 == 1)
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
| [`assert.TrueT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#TrueT) | package-level function |
| [`assert.TrueTf[B Boolean](t T, value B, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#TrueTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.TrueT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#TrueT) | package-level function |
| [`require.TrueTf[B Boolean](t T, value B, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#TrueTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.TrueT[B Boolean](t T, value B, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#TrueT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#TrueT](https://github.com/go-openapi/testify/blob/master/internal/assertions/boolean.go#L43)
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
