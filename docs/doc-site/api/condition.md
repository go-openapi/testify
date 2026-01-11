---
title: "Condition"
description: "Expressing Assertions Using Conditions"
modified: 2026-01-11
weight: 4
domains:
  - "condition"
keywords:
  - "Condition"
  - "Conditionf"
  - "Eventually"
  - "Eventuallyf"
  - "EventuallyWithT"
  - "EventuallyWithTf"
  - "Never"
  - "Neverf"
---

Expressing Assertions Using Conditions

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.

### Condition

Condition uses a [Comparison] to assert a complex condition.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Condition(t, func() bool { return myCondition })
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success:  func() bool { return true }
	failure:  func() bool { return false }
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Condition(t T, comp Comparison, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Condition) | package-level function |
| [`assert.Conditionf(t T, comp Comparison, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Conditionf) | formatted variant |
| [`assert.(*Assertions).Condition(comp Comparison) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Condition) | method variant |
| [`assert.(*Assertions).Conditionf(comp Comparison, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Condition(t T, comp Comparison, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Condition) | package-level function |
| [`require.Conditionf(t T, comp Comparison, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Conditionf) | formatted variant |
| [`require.(*Assertions).Condition(comp Comparison) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Condition) | method variant |
| [`require.(*Assertions).Conditionf(comp Comparison, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Conditionf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Condition(t T, comp Comparison, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Condition) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Condition](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L26)
{{% /tab %}}
{{< /tabs >}}

### Eventually

Eventually asserts that the given condition will be met in waitFor time,
periodically checking the target function on each tick.

[Eventually] waits until the condition returns true, for at most waitFor,
or until the parent context of the test is cancelled.

If the condition takes longer than waitFor to complete, [Eventually] fails
but waits for the current condition execution to finish before returning.

For long-running conditions to be interrupted early, check [testing.T.Context](https://pkg.go.dev/testing#T.Context)
which is cancelled on test failure.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Eventually(t, func() bool { return true }, time.Second, 10*time.Millisecond)
```
{{< /tab >}}
{{% tab title="Concurrency" %}}
```go
The condition function is never executed in parallel: only one goroutine executes it.
It may write to variables outside its scope without triggering race conditions.
A blocking condition will cause [Eventually] to hang until it returns.
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventually) | package-level function |
| [`assert.Eventuallyf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Eventuallyf) | formatted variant |
| [`assert.(*Assertions).Eventually(condition func() bool, waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Eventually) | method variant |
| [`assert.(*Assertions).Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Eventuallyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Eventually) | package-level function |
| [`require.Eventuallyf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Eventuallyf) | formatted variant |
| [`require.(*Assertions).Eventually(condition func() bool, waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Eventually) | method variant |
| [`require.(*Assertions).Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Eventuallyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Eventually) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Eventually](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L67)
{{% /tab %}}
{{< /tabs >}}

### EventuallyWithT

EventuallyWithT asserts that the given condition will be met in waitFor time,
periodically checking the target function at each tick.

In contrast to [Eventually], the condition function is supplied with a [CollectT]
to accumulate errors from calling other assertions.

The condition is considered "met" if no errors are raised in a tick.
The supplied [CollectT] collects all errors from one tick.

If the condition is not met before waitFor, the collected errors from the
last tick are copied to t.

Calling [CollectT.FailNow](https://pkg.go.dev/CollectT#FailNow) cancels the condition immediately and fails the assertion.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	externalValue := false
	go func() {
		time.Sleep(8*time.Second)
		externalValue = true
	}()
	assertions.EventuallyWithT(t, func(c *assertions.CollectT) {
		// add assertions as needed; any assertion failure will fail the current tick
		assertions.True(c, externalValue, "expected 'externalValue' to be true")
	}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```
{{< /tab >}}
{{% tab title="Concurrency" %}}
```go
The condition function is never executed in parallel: only one goroutine executes it.
It may write to variables outside its scope without triggering race conditions.
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EventuallyWithT) | package-level function |
| [`assert.EventuallyWithTf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EventuallyWithTf) | formatted variant |
| [`assert.(*Assertions).EventuallyWithT(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EventuallyWithT) | method variant |
| [`assert.(*Assertions).EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EventuallyWithTf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EventuallyWithT) | package-level function |
| [`require.EventuallyWithTf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EventuallyWithTf) | formatted variant |
| [`require.(*Assertions).EventuallyWithT(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EventuallyWithT) | method variant |
| [`require.(*Assertions).EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EventuallyWithTf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EventuallyWithT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EventuallyWithT](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L144)
{{% /tab %}}
{{< /tabs >}}

### Never

Never asserts that the given condition is never satisfied within waitFor time,
periodically checking the target function at each tick.

[Never] is the opposite of [Eventually]. It succeeds if the waitFor timeout
is reached without the condition ever returning true.

If the parent context is cancelled before the timeout, [Never] fails.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
```
{{< /tab >}}
{{% tab title="Concurrency" %}}
```go
The condition function is never executed in parallel: only one goroutine executes it.
It may write to variables outside its scope without triggering race conditions.
A blocking condition will cause [Never] to hang until it returns.
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Never) | package-level function |
| [`assert.Neverf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Neverf) | formatted variant |
| [`assert.(*Assertions).Never(condition func() bool, waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Never) | method variant |
| [`assert.(*Assertions).Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Neverf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Never) | package-level function |
| [`require.Neverf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Neverf) | formatted variant |
| [`require.(*Assertions).Never(condition func() bool, waitFor time.Duration, tick time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Never) | method variant |
| [`require.(*Assertions).Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Neverf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Never) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Never](https://github.com/go-openapi/testify/blob/master/internal/assertions/condition.go#L99)
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
