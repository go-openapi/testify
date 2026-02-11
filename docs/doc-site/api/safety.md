---
title: "Safety"
description: "Checks Against Leaked Resources (Goroutines, File Descriptors)"
weight: 13
domains:
  - "safety"
keywords:
  - "NoFileDescriptorLeak"
  - "NoFileDescriptorLeakf"
  - "NoGoRoutineLeak"
  - "NoGoRoutineLeakf"
---

Checks Against Leaked Resources (Goroutines, File Descriptors)

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 2 functionalities.

```tree
- [NoFileDescriptorLeak](#nofiledescriptorleak) | angles-right
- [NoGoRoutineLeak](#nogoroutineleak) | angles-right
```

### NoFileDescriptorLeak{#nofiledescriptorleak}
NoFileDescriptorLeak ensures that no file descriptor leaks from inside the tested function.

This assertion works on Linux only (via /proc/self/fd).
On other platforms, the test is skipped.

NOTE: this assertion is not compatible with parallel tests.
File descriptors are a process-wide resource; concurrent tests
opening files would cause false positives.

Sockets, pipes, and anonymous inodes are filtered out by default,
as these are typically managed by the Go runtime.

#### Concurrency

[NoFileDescriptorLeak](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileDescriptorLeak) is not compatible with parallel tests.
File descriptors are a process-wide resource; any concurrent I/O
from other goroutines may cause false positives.

Calls to [NoFileDescriptorLeak](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileDescriptorLeak) are serialized with a mutex
to prevent multiple leak checks from interfering with each other.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	NoFileDescriptorLeak(t, func() {
		// code that should not leak file descriptors
	})
	success: func() {}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNoFileDescriptorLeak(t *testing.T)
package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	if runtime.GOOS != "linux" {
		// This example is only runnable on linux. On other platforms, the assertion skips the test.
		// We force the expected output below, so that tests don't fail on other platforms.
		fmt.Println("success: true")

		return
	}

	t := new(testing.T) // should come from testing, e.g. func TestNoFileDescriptorLeak(t *testing.T)
	success := assert.NoFileDescriptorLeak(t, func() {
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
// real-world test would inject *testing.T from TestNoFileDescriptorLeak(t *testing.T)
package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	if runtime.GOOS != "linux" {
		// This example is only runnable on linux. On other platforms, the assertion skips the test.
		// We force the expected output below, so that tests don't fail on other platforms.
		fmt.Println("passed")

		return
	}

	t := new(testing.T) // should come from testing, e.g. func TestNoFileDescriptorLeak(t *testing.T)
	require.NoFileDescriptorLeak(t, func() {
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
| [`assert.NoFileDescriptorLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileDescriptorLeak) | package-level function |
| [`assert.NoFileDescriptorLeakf(t T, tested func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileDescriptorLeakf) | formatted variant |
| [`assert.(*Assertions).NoFileDescriptorLeak(tested func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoFileDescriptorLeak) | method variant |
| [`assert.(*Assertions).NoFileDescriptorLeakf(tested func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoFileDescriptorLeakf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NoFileDescriptorLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoFileDescriptorLeak) | package-level function |
| [`require.NoFileDescriptorLeakf(t T, tested func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoFileDescriptorLeakf) | formatted variant |
| [`require.(*Assertions).NoFileDescriptorLeak(tested func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoFileDescriptorLeak) | method variant |
| [`require.(*Assertions).NoFileDescriptorLeakf(tested func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoFileDescriptorLeakf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NoFileDescriptorLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NoFileDescriptorLeak) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoFileDescriptorLeak](https://github.com/go-openapi/testify/blob/master/internal/assertions/safety.go#L98)
{{% /tab %}}
{{< /tabs >}}

### NoGoRoutineLeak{#nogoroutineleak}
NoGoRoutineLeak ensures that no goroutine did leak from inside the tested function.

NOTE: only the go routines spawned from inside the tested function are checked for leaks.
No filter or configuration is needed to exclude "known go routines".

Resource cleanup should be done inside the tested function, and not using [testing.T.Cleanup](https://pkg.go.dev/testing#T.Cleanup),
as t.Cleanup is called after the leak check.

#### Edge cases

  - if the tested function panics leaving behind leaked goroutines, these are detected.
  - if the tested function calls runtime.Goexit (e.g. from [testing.T.FailNow](https://pkg.go.dev/testing#T.FailNow)) leaving behind leaked goroutines,
    these are detected.
  - if a panic occurs in one of the leaked go routines, it cannot be recovered with certainty and
    the calling program will usually panic.

#### Concurrency

[NoGoRoutineLeak](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoGoRoutineLeak) may be used safely in parallel tests.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	NoGoRoutineLeak(t, func() {
		...
	},
	"should not leak any go routine",
	)
	success: func() {}
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNoGoRoutineLeak(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNoGoRoutineLeak(t *testing.T)
	success := assert.NoGoRoutineLeak(t, func() {
	})
	fmt.Printf("success: %t\n", success)

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNoGoRoutineLeak(t *testing.T)
package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // normally provided by test

	blocker := make(chan struct{})
	var wg sync.WaitGroup

	defer func() {
		// clean resources _after_ the test
		close(blocker)
		wg.Wait()
	}()

	wg.Add(1)
	// This examplifies how a function that leaks a goroutine is detected.
	result := assert.NoGoRoutineLeak(t, func() { // true when there is no leak
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks until cleanup
		}()
	})

	// Error message from test would typically return the leaked goroutine, e.g.:
	// #	0x69c8e8	github.com/go-openapi/testify/v2/assert_test.ExampleNoGoRoutineLeak.func2.1+0x48	.../assert_adhoc_example_7_test.go:30
	fmt.Printf("has leak: %t", !result)
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
// real-world test would inject *testing.T from TestNoGoRoutineLeak(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNoGoRoutineLeak(t *testing.T)
	require.NoGoRoutineLeak(t, func() {
	})
	fmt.Println("passed")

}

```
{{% /card %}}


{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNoGoRoutineLeak(t *testing.T)
// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"sync"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(mockFailNowT) // normally provided by test
	// Since this test is failing and calls [runtime.Goexit], we need a mock to
	// avoid the example trigger a panick.

	blocker := make(chan struct{})
	var wg sync.WaitGroup

	defer func() {
		// clean resources _after_ the test
		close(blocker)
		wg.Wait()
	}()

	wg.Add(1)
	// This examplifies how a function that leaks a goroutine is detected.
	require.NoGoRoutineLeak(t, func() { // true when there is no leak
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks until cleanup
		}()
	})

	// Error message from test would typically return the leaked goroutine, e.g.:
	// #	0x69c8e8	github.com/go-openapi/testify/v2/assert_test.ExampleNoGoRoutineLeak.func2.1+0x48	.../assert_adhoc_example_7_test.go:30
	fmt.Printf("passed: %t", !t.Failed())

}

type mockFailNowT struct {
	failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockFailNowT) Helper() {}

func (m *mockFailNowT) Errorf(format string, args ...any) {
	_ = format
	_ = args
}

func (m *mockFailNowT) FailNow() {
	m.failed = true
}

func (m *mockFailNowT) Failed() bool {
	return m.failed
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
| [`assert.NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoGoRoutineLeak) | package-level function |
| [`assert.NoGoRoutineLeakf(t T, tested func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoGoRoutineLeakf) | formatted variant |
| [`assert.(*Assertions).NoGoRoutineLeak(tested func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoGoRoutineLeak) | method variant |
| [`assert.(*Assertions).NoGoRoutineLeakf(tested func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoGoRoutineLeakf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoGoRoutineLeak) | package-level function |
| [`require.NoGoRoutineLeakf(t T, tested func(), msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoGoRoutineLeakf) | formatted variant |
| [`require.(*Assertions).NoGoRoutineLeak(tested func()) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoGoRoutineLeak) | method variant |
| [`require.(*Assertions).NoGoRoutineLeakf(tested func(), msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoGoRoutineLeakf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NoGoRoutineLeak) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoGoRoutineLeak](https://github.com/go-openapi/testify/blob/master/internal/assertions/safety.go#L45)
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
