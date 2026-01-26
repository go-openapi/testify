---
title: "Time"
description: "Asserting Times And Durations"
modified: 2026-01-26
weight: 15
domains:
  - "time"
keywords:
  - "WithinDuration"
  - "WithinDurationf"
  - "WithinRange"
  - "WithinRangef"
---

Asserting Times And Durations

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 2 functionalities.

```tree
- [WithinDuration](#withinduration) | angles-right
- [WithinRange](#withinrange) | angles-right
```

### WithinDuration{#withinduration}

WithinDuration asserts that the two times are within duration delta of each other.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.WithinDuration(t, time.Now(), 10*time.Second)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second
	failure: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.WithinDuration(t T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithinDuration) | package-level function |
| [`assert.WithinDurationf(t T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithinDurationf) | formatted variant |
| [`assert.(*Assertions).WithinDuration(expected time.Time, actual time.Time, delta time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.WithinDuration) | method variant |
| [`assert.(*Assertions).WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.WithinDurationf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.WithinDuration(t T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#WithinDuration) | package-level function |
| [`require.WithinDurationf(t T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#WithinDurationf) | formatted variant |
| [`require.(*Assertions).WithinDuration(expected time.Time, actual time.Time, delta time.Duration) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.WithinDuration) | method variant |
| [`require.(*Assertions).WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.WithinDurationf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.WithinDuration(t T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#WithinDuration) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#WithinDuration](https://github.com/go-openapi/testify/blob/master/internal/assertions/time.go#L21)
{{% /tab %}}
{{< /tabs >}}

### WithinRange{#withinrange}

WithinRange asserts that a time is within a time range (inclusive).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
	failure: time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.WithinRange(t T, actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithinRange) | package-level function |
| [`assert.WithinRangef(t T, actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithinRangef) | formatted variant |
| [`assert.(*Assertions).WithinRange(actual time.Time, start time.Time, end time.Time) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.WithinRange) | method variant |
| [`assert.(*Assertions).WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.WithinRangef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.WithinRange(t T, actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#WithinRange) | package-level function |
| [`require.WithinRangef(t T, actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#WithinRangef) | formatted variant |
| [`require.(*Assertions).WithinRange(actual time.Time, start time.Time, end time.Time) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.WithinRange) | method variant |
| [`require.(*Assertions).WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.WithinRangef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.WithinRange(t T, actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#WithinRange) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#WithinRange](https://github.com/go-openapi/testify/blob/master/internal/assertions/time.go#L45)
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

Generated on 2026-01-26 (version 43574c8) using codegen version v2.2.1-0.20260126160846-43574c83eea9+dirty [sha: 43574c83eea9c46dc5bb573128a4038e90e2f44b]
-->
