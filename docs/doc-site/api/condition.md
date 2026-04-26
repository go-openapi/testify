---
title: "Condition"
description: "Expressing Assertions Using Conditions"
weight: 4
domains:
  - "condition"
keywords:
  - "Blocked"
  - "Blockedf"
  - "BlockedT"
  - "BlockedTf"
  - "Condition"
  - "Conditionf"
  - "Consistently"
  - "Consistentlyf"
  - "Eventually"
  - "Eventuallyf"
  - "EventuallyWith"
  - "EventuallyWithf"
  - "Never"
  - "Neverf"
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

This domain exposes 9 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [Blocked](#blocked) | angles-right
- [BlockedT[E any, CHAN ~chan E]](#blockedte-any-chan-chan-e) | star | orange
- [Condition](#condition) | angles-right
- [Consistently[C Conditioner]](#consistentlyc-conditioner) | star | orange
- [Eventually[C Conditioner]](#eventuallyc-conditioner) | star | orange
- [EventuallyWith[C CollectibleConditioner]](#eventuallywithc-collectibleconditioner) | star | orange
- [Never[C NeverConditioner]](#neverc-neverconditioner) | star | orange
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
| [`assert.(*Assertions).Blockedf(ch any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Blockedf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Blocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Blocked) | package-level function |
| [`require.Blockedf(t T, ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Blockedf) | formatted variant |
| [`require.(*Assertions).Blocked(ch any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Blocked) | method variant |
| [`require.(*Assertions).Blockedf(ch any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Blockedf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Blocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Blocked) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Blocked](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L56)
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
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#BlockedT) | package-level function |
| [`require.BlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#BlockedTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#BlockedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#BlockedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L104)
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
| [`assert.(*Assertions).Conditionf(comp func() bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Condition(t T, comp func() bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Condition) | package-level function |
| [`require.Conditionf(t T, comp func() bool, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Conditionf) | formatted variant |
| [`require.(*Assertions).Condition(comp func() bool) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Condition) | method variant |
| [`require.(*Assertions).Conditionf(comp func() bool, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Condition(t T, comp func() bool, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Condition) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Condition](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L29)
{{% /tab %}}
{{< /tabs >}}

### Consistently[C Conditioner] {{% icon icon="star" color=orange %}}{#consistentlyc-conditioner}
Consistently asserts that the given condition is always satisfied until timeout,
periodically checking the target function at each tick.

[Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) ("always") imposes a stronger constraint than [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) ("at least once"):
it checks at every tick that every occurrence of the condition is satisfied, whereas
[Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) succeeds on the first occurrence of a successful condition.

#### Alternative condition signature

The simplest form of condition is:

	func() bool

The semantics of the assertion are "always returns true".

To build more complex cases, a condition may also be defined as:

	func(context.Context) error

It fails as soon as an error is returned before timeout expressing "always returns no error (nil)"

This is consistent with [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) expressing "eventually returns no error (nil)".

It will be executed with the context of the assertion, which inherits the [testing.T.Context](https://pkg.go.dev/testing#T.Context) and
is cancelled on timeout.

#### Panic recovery

A panicking condition is treated as an error, causing [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) to fail immediately.
See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details.

#### Concurrency

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Attention point

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Synctest (opt-in)

Wrap the condition with [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) (or [WithSynctestContext](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctestContext)) to run
the polling loop inside a [testing/synctest] bubble, which uses a fake
clock. This eliminates timing-induced flakiness and makes the tick count
deterministic. See [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) for the constraints (no real I/O in
the condition, requires [*testing.T]).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Consistently(t, func() bool { return true }, time.Second, 10*time.Millisecond)
See also [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details about using context, concurrency, and panic recovery.
	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestConsistently(t *testing.T)
	success := assert.Consistently(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that stays healthy.
	healthCheck := func(_ context.Context) error {
		return nil // always healthy
	}

	result := assert.Consistently(t, healthCheck, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently healthy: %t", result)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A counter that stays within bounds during the test.
	var counter atomic.Int32
	counter.Store(5)

	result := assert.Consistently(t, func() bool {
		return counter.Load() < 10
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently under limit: %t", result)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// An invariant that must hold throughout the observation period.
	var counter atomic.Int32
	counter.Store(5)
	invariant := func() bool { return counter.Load() < 10 }

	result := assert.Consistently(t, assert.WithSynctest(invariant), 1*time.Hour, 1*time.Minute)

	fmt.Printf("invariant held: %t", result)

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
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestConsistently(t *testing.T)
	require.Consistently(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that stays healthy.
	healthCheck := func(_ context.Context) error {
		return nil // always healthy
	}

	require.Consistently(t, healthCheck, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently healthy: %t", !t.Failed())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A counter that stays within bounds during the test.
	var counter atomic.Int32
	counter.Store(5)

	require.Consistently(t, func() bool {
		return counter.Load() < 10
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently under limit: %t", !t.Failed())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestConsistently(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// An invariant that must hold throughout the observation period.
	var counter atomic.Int32
	counter.Store(5)
	invariant := func() bool { return counter.Load() < 10 }

	require.Consistently(t, require.WithSynctest(invariant), 1*time.Hour, 1*time.Minute)

	fmt.Printf("invariant held: %t", !t.Failed())

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
| [`assert.Consistently[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) | package-level function |
| [`assert.Consistentlyf[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistentlyf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Consistently[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Consistently) | package-level function |
| [`require.Consistentlyf[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Consistentlyf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Consistently[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Consistently) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Consistently](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L414)
{{% /tab %}}
{{< /tabs >}}

### Eventually[C Conditioner] {{% icon icon="star" color=orange %}}{#eventuallyc-conditioner}
Eventually asserts that the given condition will be met before timeout,
periodically checking the target function on each tick.

[Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) waits until the condition returns true, at most until timeout,
or until the parent context of the test is cancelled.

If the condition takes longer than the timeout to complete, [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) fails
but waits for the current condition execution to finish before returning.

For long-running conditions to be interrupted early, check [testing.T.Context](https://pkg.go.dev/testing#T.Context)
which is cancelled on test failure.

#### Alternative condition signature

The simplest form of condition is:

	func() bool

To build more complex cases, a condition may also be defined as:

	func(context.Context) error

It fails when an error has always been returned up to timeout (equivalent semantics to func() bool returns false),
expressing "eventually returns no error (nil)".

It will be executed with the context of the assertion, which inherits the [testing.T.Context](https://pkg.go.dev/testing#T.Context) and
is cancelled on timeout.

The semantics of the three available async assertions read as follows.

  - [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) (func() bool) : "eventually returns true"

  - [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) (func() bool) : "never returns true"

  - [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) (func() bool): "always returns true"

  - [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) (func(ctx) error) : "eventually returns nil"

  - [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) (func(ctx) error) : not supported, use [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) instead (avoids confusion with double negation)

  - [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) (func(ctx) error): "always returns nil"

#### Concurrency

The condition function is always executed serially by a single goroutine. It is always executed at least once.

It may thus write to variables outside its scope without triggering race conditions.

A blocking condition will cause [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) to hang until it returns.

Notice that time ticks may be skipped if the condition takes longer than the tick interval.

#### Panic recovery

If the condition panics, the panic is recovered and treated as a failed tick
(equivalent to returning false or a non-nil error). For [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually), this means
the poller retries on the next tick — if a later tick succeeds, the assertion
succeeds. For [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) and [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently), a panic is treated as the condition
erroring, which causes immediate failure.

The recovered panic is wrapped as an error with the sentinel [errConditionPanicked](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#errConditionPanicked),
detectable with [errors.Is](https://pkg.go.dev/errors#Is).

#### Attention point

Time-based tests may be flaky in a resource-constrained environment such as a CI runner and may produce
counter-intuitive results, such as ticks or timeouts not firing in time as expected.

To avoid flaky tests, always make sure that ticks and timeouts differ by at least an order of magnitude (tick <<
timeout).

#### Synctest (opt-in)

Wrap the condition with [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) (or [WithSynctestContext](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctestContext)) to run
the polling loop inside a [testing/synctest] bubble, which uses a fake
clock. This eliminates timing-induced flakiness and makes the tick count
deterministic. See [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) for the constraints (no real I/O in
the condition, requires `*testing.T`).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Eventually(t, func() bool { return true }, time.Second, 10*time.Millisecond)
	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEventually(t *testing.T)
	success := assert.Eventually(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate an async operation that completes after a short delay.
	var ready atomic.Bool
	go func() {
		time.Sleep(30 * time.Millisecond)
		ready.Store(true)
	}()

	result := assert.Eventually(t, ready.Load, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually ready: %t", result)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that becomes healthy after a few attempts.
	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	result := assert.Eventually(t, healthCheck, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually healthy: %t", result)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	result := assert.Eventually(t, assert.WithSynctestContext(healthCheck), 1*time.Hour, 1*time.Minute)

	fmt.Printf("healthy: %t, attempts: %d", result, attempts.Load())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A counter that converges on the 5th poll — no external time pressure.
	var attempts atomic.Int32
	cond := func() bool {
		return attempts.Add(1) == 5
	}

	// 1-hour/1-minute: under fake time this is instantaneous and
	// deterministic — exactly 5 calls to the condition.
	result := assert.Eventually(t, assert.WithSynctest(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("ready: %t, attempts: %d", result, attempts.Load())

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
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEventually(t *testing.T)
	require.Eventually(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate an async operation that completes after a short delay.
	var ready atomic.Bool
	go func() {
		time.Sleep(30 * time.Millisecond)
		ready.Store(true)
	}()

	require.Eventually(t, ready.Load, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually ready: %t", !t.Failed())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that becomes healthy after a few attempts.
	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	require.Eventually(t, healthCheck, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually healthy: %t", !t.Failed())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	require.Eventually(t, require.WithSynctestContext(healthCheck), 1*time.Hour, 1*time.Minute)

	fmt.Printf("healthy: %t, attempts: %d", !t.Failed(), attempts.Load())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventually(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A counter that converges on the 5th poll — no external time pressure.
	var attempts atomic.Int32
	cond := func() bool {
		return attempts.Add(1) == 5
	}

	// 1-hour/1-minute: under fake time this is instantaneous and
	// deterministic — exactly 5 calls to the condition.
	require.Eventually(t, require.WithSynctest(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("ready: %t, attempts: %d", !t.Failed(), attempts.Load())

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
| [`assert.Eventually[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) | package-level function |
| [`assert.Eventuallyf[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventuallyf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Eventually[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Eventually) | package-level function |
| [`require.Eventuallyf[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Eventuallyf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Eventually[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Eventually) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Eventually](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L290)
{{% /tab %}}
{{< /tabs >}}

### EventuallyWith[C CollectibleConditioner] {{% icon icon="star" color=orange %}}{#eventuallywithc-collectibleconditioner}
EventuallyWith asserts that the given condition will be met before the timeout,
periodically checking the target function at each tick.

In contrast to [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually), the condition function is supplied with a [CollectT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#CollectT)
to accumulate errors from calling other assertions.

The condition is considered "met" if no errors are raised in a tick.
The supplied [CollectT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#CollectT) collects all errors from one tick.

If the condition is not met before the timeout, the collected errors from the
last tick are copied to t.

Calling [CollectT.FailNow](https://pkg.go.dev/CollectT#FailNow) (directly, or transitively through [require](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#require) assertions)
fails the current tick only: the poller will retry on the next tick. This means
[require](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#require)-style assertions inside [EventuallyWith](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EventuallyWith) behave naturally — they abort
the current evaluation and let the polling loop converge.

To abort the whole assertion immediately (e.g. when the condition can no longer
be expected to succeed), call [CollectT.Cancel](https://pkg.go.dev/CollectT#Cancel).

#### Concurrency

The condition function is never executed in parallel: only one goroutine executes it.
It may write to variables outside its scope without triggering race conditions.

The condition is wrapped in its own goroutine, so a call to [runtime.Goexit](https://pkg.go.dev/runtime#Goexit)
(e.g. via [require](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#require) assertions or [CollectT.FailNow](https://pkg.go.dev/CollectT#FailNow)) cleanly aborts only the
current tick.

#### Panic recovery

If the condition panics, the panic is recovered and recorded as an error in the
[CollectT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#CollectT) for that tick. The poller treats it as a failed tick and retries on the
next one. If the assertion times out, the panic error is included in the collected
errors reported on the parent t.

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for the general panic recovery semantics.

#### Synctest (opt-in)

Wrap the condition with [WithSynctestCollect](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctestCollect) (or [WithSynctestCollectContext](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctestCollectContext))
to run the polling loop inside a [testing/synctest] bubble, which uses
a fake clock. This eliminates timing-induced flakiness and makes the
tick count deterministic. See [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) for the constraints (no
real I/O in the condition, requires [*testing.T]).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	externalValue := false
	go func() {
		time.Sleep(8*time.Second)
		externalValue = true
	}()
	assertions.EventuallyWith(t, func(c *assertions.CollectT) {
		// add assertions as needed; any assertion failure will fail the current tick
		assertions.True(c, externalValue, "expected 'externalValue' to be true")
	},
	10*time.Second,
	1*time.Second,
	"external state has not changed to 'true'; still false",
	)
	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
	failure: func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventuallyWith(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEventuallyWith(t *testing.T)
	success := assert.EventuallyWith(t, func(c *assert.CollectT) {
		assert.True(c, true)
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventuallyWith(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	cond := func(c *assert.CollectT) {
		n := attempts.Add(1)
		assert.Equal(c, int32(3), n, "not yet converged")
	}

	result := assert.EventuallyWith(t, assert.WithSynctestCollect(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("converged: %t, attempts: %d", result, attempts.Load())

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
// real-world test would inject *testing.T from TestEventuallyWith(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEventuallyWith(t *testing.T)
	require.EventuallyWith(t, func(c *assert.CollectT) {
		assert.True(c, true)
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEventuallyWith(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	cond := func(c *require.CollectT) {
		n := attempts.Add(1)
		require.Equal(c, int32(3), n, "not yet converged")
	}

	require.EventuallyWith(t, require.WithSynctestCollect(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("converged: %t, attempts: %d", !t.Failed(), attempts.Load())

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
| [`assert.EventuallyWith[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EventuallyWith) | package-level function |
| [`assert.EventuallyWithf[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EventuallyWithf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EventuallyWith[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EventuallyWith) | package-level function |
| [`require.EventuallyWithf[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EventuallyWithf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.EventuallyWith[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EventuallyWith) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EventuallyWith](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L491)
{{% /tab %}}
{{< /tabs >}}

### Never[C NeverConditioner] {{% icon icon="star" color=orange %}}{#neverc-neverconditioner}
Never asserts that the given condition is never satisfied until timeout,
periodically checking the target function at each tick.

[Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) is the opposite of [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) ("at least once").
It succeeds if the timeout is reached without the condition ever returning true.

If the parent context is cancelled before the timeout, [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) fails.

#### Alternative condition signature

The simplest form of condition is:

	func() bool

Use [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) instead if you want to use a condition returning an error.

#### Panic recovery

A panicking condition is treated as an error, causing [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) to fail immediately.
See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details.

#### Concurrency

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Attention point

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Synctest (opt-in)

Wrap the condition with [WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) to run the polling loop inside a
[testing/synctest] bubble, which uses a fake clock. This eliminates
timing-induced flakiness and makes the tick count deterministic. See
[WithSynctest](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctest) for the constraints (no real I/O in the condition,
requires [*testing.T]). Note: [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) does not accept the context/error
form of condition, so [WithSynctestContext](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithSynctestContext) does not apply here.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
See also [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details about using context, concurrency, and panic recovery.
	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNever(t *testing.T)
	success := assert.Never(t, func() bool {
		return false
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A channel that should remain empty during the test.
	events := make(chan struct{}, 1)

	result := assert.Never(t, func() bool {
		select {
		case <-events:
			return true // event received = condition becomes true = Never fails
		default:
			return false
		}
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("never received: %t", result)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A flag that should remain false across the whole observation period.
	var flipped atomic.Bool
	result := assert.Never(t, assert.WithSynctest(flipped.Load), 1*time.Hour, 1*time.Minute)

	fmt.Printf("never flipped: %t", result)

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
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNever(t *testing.T)
	require.Never(t, func() bool {
		return false
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A channel that should remain empty during the test.
	events := make(chan struct{}, 1)

	require.Never(t, func() bool {
		select {
		case <-events:
			return true // event received = condition becomes true = Never fails
		default:
			return false
		}
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("never received: %t", !t.Failed())

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNever(t *testing.T)
package main

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // normally provided by test

	// A flag that should remain false across the whole observation period.
	var flipped atomic.Bool
	require.Never(t, require.WithSynctest(flipped.Load), 1*time.Hour, 1*time.Minute)

	fmt.Printf("never flipped: %t", !t.Failed())

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
| [`assert.Never[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) | package-level function |
| [`assert.Neverf[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Neverf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Never[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Never) | package-level function |
| [`require.Neverf[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Neverf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Never[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Never) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Never](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L348)
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
| [`assert.(*Assertions).NotBlockedf(ch any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotBlockedf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotBlocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlocked) | package-level function |
| [`require.NotBlockedf(t T, ch any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedf) | formatted variant |
| [`require.(*Assertions).NotBlocked(ch any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlocked) | method variant |
| [`require.(*Assertions).NotBlockedf(ch any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotBlockedf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotBlocked(t T, ch any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotBlocked) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotBlocked](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L141)
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
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedT) | package-level function |
| [`require.NotBlockedTf[E any, CHAN ~chan E](t T, ch CHAN, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotBlockedTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotBlockedT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotBlockedT](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L188)
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
