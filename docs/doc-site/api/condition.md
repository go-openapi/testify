---
title: "Condition"
description: "Expressing Assertions Using Conditions"
weight: 5
domains:
  - "condition"
keywords:
  - "Blocked"
  - "Blockedf"
  - "BlockedT"
  - "BlockedTf"
  - "Condition"
  - "Conditionf"
  - "NotBlocked"
  - "NotBlockedf"
  - "NotBlockedT"
  - "NotBlockedTf"
---

Expressing Assertions Using Conditions

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 5 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.
Their method variants carry a {{% goversion "go1.27" %}} badge: methods take type
parameters only from go1.27 onwards, so on an older toolchain a generic assertion is available as a
package-level function alone.

```tree
- [Blocked](#blocked) | angles-right
- [BlockedT[E any, CHAN ~chan E]](#blockedte-any-chan-chan-e) | star | orange
- [Condition](#condition) | angles-right
- [NotBlocked](#notblocked) | angles-right
- [NotBlockedT[E any, CHAN ~chan E]](#notblockedte-any-chan-chan-e) | star | orange
```

### Blocked{#blocked}
Blocked asserts that a channel is blocked on receive.

It always fails if the operand is not a channel, or if the channel is send-only.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	ch := make(chan struct{})
	assertions.Blocked(t, ch)
	success:  make(chan struct{})
	failure:  sendChanMessage()
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestBlocked(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestBlocked(t *testing.T)
	success := assert.Blocked(t, make(chan struct{}))
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestBlocked(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestBlocked(t *testing.T)
	require.Blocked(t, make(chan struct{}))
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
| [`assert.Blocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Blocked) | package-level function |
| [`assert.Blockedf(t T, ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Blockedf) | formatted variant |
| [`assert.(*Assertions).Blocked(ch any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Blocked) | method variant |
| [`assert.(*Assertions).Blockedf(ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Blockedf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Blocked(t T, ch any, msgAndArgs ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Blocked) | package-level function |
| [`require.Blockedf(t T, ch any, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Blockedf) | formatted variant |
| [`require.(*Assertions).Blocked(ch any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Blocked) | method variant |
| [`require.(*Assertions).Blockedf(ch any, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Blockedf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Blocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Blocked) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Blocked](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L48)
{{% /tab %}}
{{< /tabs >}}

### BlockedT[E any, CHAN ~chan E] {{% icon icon="star" color=orange %}}{#blockedte-any-chan-chan-e}
BlockedT asserts that a channel is blocked on receive.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	ch := make(chan struct{})
	assertions.BlockedT(t, ch)
	success:  make(chan struct{})
	failure:  sendChanMessage()
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestBlockedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestBlockedT(t *testing.T)
	success := assert.BlockedT(t, make(chan struct{}))
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestBlockedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestBlockedT(t *testing.T)
	require.BlockedT(t, make(chan struct{}))
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
| [`assert.BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#BlockedT) | package-level function |
| [`assert.BlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#BlockedTf) | formatted variant |
| [`assert.(*Assertions).BlockedT[E any, CHAN ~chan E](ch CHAN) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.BlockedT) | method variant {{% goversion "go1.27" %}} |
| [`assert.(*Assertions).BlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.BlockedTf) | method formatted variant {{% goversion "go1.27" %}} |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#BlockedT) | package-level function |
| [`require.BlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#BlockedTf) | formatted variant |
| [`require.(*Assertions).BlockedT[E any, CHAN ~chan E](ch CHAN)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.BlockedT) | method variant {{% goversion "go1.27" %}} |
| [`require.(*Assertions).BlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.BlockedTf) | method formatted variant {{% goversion "go1.27" %}} |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#BlockedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#BlockedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L96)
{{% /tab %}}
{{< /tabs >}}

### Condition{#condition}
Condition uses a comparison function to assert a complex condition.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Condition(t, func() bool { return myCondition })
	success:  func() bool { return true }
	failure:  func() bool { return false }
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestCondition(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestCondition(t *testing.T)
	success := assert.Condition(t, func() bool {
		return true
	})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestCondition(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestCondition(t *testing.T)
	require.Condition(t, func() bool {
		return true
	})
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
| [`assert.Condition(t T, comp func() bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Condition) | package-level function |
| [`assert.Conditionf(t T, comp func() bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Conditionf) | formatted variant |
| [`assert.(*Assertions).Condition(comp func() bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Condition) | method variant |
| [`assert.(*Assertions).Conditionf(comp func() bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Condition(t T, comp func() bool, msgAndArgs ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Condition) | package-level function |
| [`require.Conditionf(t T, comp func() bool, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Conditionf) | formatted variant |
| [`require.(*Assertions).Condition(comp func() bool)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Condition) | method variant |
| [`require.(*Assertions).Conditionf(comp func() bool, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Condition(t T, comp func() bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Condition) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Condition](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L21)
{{% /tab %}}
{{< /tabs >}}

### NotBlocked{#notblocked}
NotBlocked asserts that a channel is not blocked on receive.

It always fails if the operand is not a channel, or if the channel is send-only.

A closed channel doesn't block and returns true.
Notice that this consumes any message available in the channel.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	assertions.NotBlocked(t, ch)
	success:  sendChanMessage()
	failure:  make(chan struct{})
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotBlocked(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotBlocked(t *testing.T)
	success := assert.NotBlocked(t, sendChanMessage())
	fmt.Printf("success: %t\n", success)

}

func sendChanMessage() chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	return ch
}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotBlocked(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotBlocked(t *testing.T)
	require.NotBlocked(t, sendChanMessage())
	fmt.Println("passed")

}

func sendChanMessage() chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	return ch
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
| [`assert.NotBlocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotBlocked) | package-level function |
| [`assert.NotBlockedf(t T, ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotBlockedf) | formatted variant |
| [`assert.(*Assertions).NotBlocked(ch any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotBlocked) | method variant |
| [`assert.(*Assertions).NotBlockedf(ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotBlockedf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotBlocked(t T, ch any, msgAndArgs ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlocked) | package-level function |
| [`require.NotBlockedf(t T, ch any, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedf) | formatted variant |
| [`require.(*Assertions).NotBlocked(ch any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlocked) | method variant |
| [`require.(*Assertions).NotBlockedf(ch any, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlockedf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotBlocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotBlocked) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotBlocked](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L133)
{{% /tab %}}
{{< /tabs >}}

### NotBlockedT[E any, CHAN ~chan E] {{% icon icon="star" color=orange %}}{#notblockedte-any-chan-chan-e}
NotBlockedT asserts that a channel is not blocked on receive.

A closed channel doesn't block and returns true.
Notice that this consumes any message available in the channel.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	assertions.NotBlockedT(t, ch)
	success:  sendChanMessage()
	failure:  make(chan struct{})
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotBlockedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotBlockedT(t *testing.T)
	success := assert.NotBlockedT(t, sendChanMessage())
	fmt.Printf("success: %t\n", success)

}

func sendChanMessage() chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	return ch
}

```
{{% /card %}}


{{% /cards %}}
{{< /tab >}}


{{% tab title="Testable Examples (require)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotBlockedT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotBlockedT(t *testing.T)
	require.NotBlockedT(t, sendChanMessage())
	fmt.Println("passed")

}

func sendChanMessage() chan struct{} {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}

	return ch
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
| [`assert.NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotBlockedT) | package-level function |
| [`assert.NotBlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotBlockedTf) | formatted variant |
| [`assert.(*Assertions).NotBlockedT[E any, CHAN ~chan E](ch CHAN) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotBlockedT) | method variant {{% goversion "go1.27" %}} |
| [`assert.(*Assertions).NotBlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotBlockedTf) | method formatted variant {{% goversion "go1.27" %}} |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedT) | package-level function |
| [`require.NotBlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedTf) | formatted variant |
| [`require.(*Assertions).NotBlockedT[E any, CHAN ~chan E](ch CHAN)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlockedT) | method variant {{% goversion "go1.27" %}} |
| [`require.(*Assertions).NotBlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlockedTf) | method formatted variant {{% goversion "go1.27" %}} |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotBlockedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotBlockedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L180)
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
