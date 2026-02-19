---
title: "Condition"
description: "Expressing Assertions Using Conditions"
weight: 4
domains:
  - "condition"
keywords:
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
---

Expressing Assertions Using Conditions

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 5 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [Condition](#condition) | angles-right
- [Consistently[C Conditioner]](#consistentlyc-conditioner) | star | orange
- [Eventually[C Conditioner]](#eventuallyc-conditioner) | star | orange
- [EventuallyWith[C CollectibleConditioner]](#eventuallywithc-collectibleconditioner) | star | orange
- [Never](#never) | angles-right
```

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Condition](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L26)
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

#### Concurrency

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Attention point

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Consistently(t, func() bool { return true }, time.Second, 10*time.Millisecond)
See also [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details about using context and concurrency.
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Consistently](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L204)
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

#### Attention point

Time-based tests may be flaky in a resource-constrained environment such as a CI runner and may produce
counter-intuitive results, such as ticks or timeouts not firing in time as expected.

To avoid flaky tests, always make sure that ticks and timeouts differ by at least an order of magnitude (tick <<
timeout).

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Eventually](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L108)
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

Calling [CollectT.FailNow](https://pkg.go.dev/CollectT#FailNow) cancels the condition immediately and causes the assertion to fail.

#### Concurrency

The condition function is never executed in parallel: only one goroutine executes it.
It may write to variables outside its scope without triggering race conditions.

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EventuallyWith](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L253)
{{% /tab %}}
{{< /tabs >}}

### Never{#never}
Never asserts that the given condition is never satisfied until timeout,
periodically checking the target function at each tick.

[Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) is the opposite of [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) ("at least once").
It succeeds if the timeout is reached without the condition ever returning true.

If the parent context is cancelled before the timeout, [Never](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) fails.

#### Alternative condition signature

The simplest form of condition is:

	func() bool

Use [Consistently](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Consistently) instead if you want to use a condition returning an error.

#### Concurrency

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

#### Attention point

See [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
See also [Eventually](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) for details about using context and concurrency.
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


{{% /cards %}}
{{< /tab >}}


{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Never(t T, condition func() bool, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) | package-level function |
| [`assert.Neverf(t T, condition func() bool, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Neverf) | formatted variant |
| [`assert.(*Assertions).Never(condition func() bool, timeout time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Never) | method variant |
| [`assert.(*Assertions).Neverf(condition func() bool, timeout time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Neverf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Never(t T, condition func() bool, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Never) | package-level function |
| [`require.Neverf(t T, condition func() bool, timeout time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Neverf) | formatted variant |
| [`require.(*Assertions).Never(condition func() bool, timeout time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Never) | method variant |
| [`require.(*Assertions).Neverf(condition func() bool, timeout time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Neverf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Never(t T, condition func() bool, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Never) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Never](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L151)
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
