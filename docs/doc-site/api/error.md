---
title: "Error"
description: "Asserting Errors"
weight: 6
domains:
  - "error"
keywords:
  - "EqualError"
  - "EqualErrorf"
  - "Error"
  - "Errorf"
  - "ErrorAs"
  - "ErrorAsf"
  - "ErrorContains"
  - "ErrorContainsf"
  - "ErrorIs"
  - "ErrorIsf"
  - "NoError"
  - "NoErrorf"
  - "NotErrorAs"
  - "NotErrorAsf"
  - "NotErrorIs"
  - "NotErrorIsf"
---

Asserting Errors

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 8 functionalities.

```tree
- [EqualError](#equalerror) | angles-right
- [Error](#error) | angles-right
- [ErrorAs](#erroras) | angles-right
- [ErrorContains](#errorcontains) | angles-right
- [ErrorIs](#erroris) | angles-right
- [NoError](#noerror) | angles-right
- [NotErrorAs](#noterroras) | angles-right
- [NotErrorIs](#noterroris) | angles-right
```

### EqualError{#equalerror}
EqualError asserts that a function returned a non-nil error (i.e. an error)
and that it is equal to the provided error.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	assertions.EqualError(t, err,  expectedErrorString)
	success: ErrTest, "assert.ErrTest general error for testing"
	failure: ErrTest, "wrong error message"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestEqualError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualError(t *testing.T)
	success := assert.EqualError(t, assert.ErrTest, "assert.ErrTest general error for testing")
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
// real-world test would inject *testing.T from TestEqualError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualError(t *testing.T)
	require.EqualError(t, require.ErrTest, "assert.ErrTest general error for testing")
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
| [`assert.EqualError(t T, err error, errString string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualError) | package-level function |
| [`assert.EqualErrorf(t T, err error, errString string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualErrorf) | formatted variant |
| [`assert.(*Assertions).EqualError(err error, errString string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualError) | method variant |
| [`assert.(*Assertions).EqualErrorf(err error, errString string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualErrorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualError(t T, err error, errString string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualError) | package-level function |
| [`require.EqualErrorf(t T, err error, errString string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualErrorf) | formatted variant |
| [`require.(*Assertions).EqualError(err error, errString string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualError) | method variant |
| [`require.(*Assertions).EqualErrorf(err error, errString string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualErrorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.EqualError(t T, err error, errString string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualError) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualError](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L89)
{{% /tab %}}
{{< /tabs >}}

### Error{#error}
Error asserts that a function returned a non-nil error (i.e. an error).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	assertions.Error(t, err)
	success: ErrTest
	failure: nil
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestError(t *testing.T)
	success := assert.Error(t, assert.ErrTest)
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
// real-world test would inject *testing.T from TestError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestError(t *testing.T)
	require.Error(t, require.ErrTest)
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
| [`assert.Error(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Error) | package-level function |
| [`assert.Errorf(t T, err error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Errorf) | formatted variant |
| [`assert.(*Assertions).Error(err error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Error) | method variant |
| [`assert.(*Assertions).Errorf(err error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Errorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Error(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Error) | package-level function |
| [`require.Errorf(t T, err error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Errorf) | formatted variant |
| [`require.(*Assertions).Error(err error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Error) | method variant |
| [`require.(*Assertions).Errorf(err error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Errorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.Error(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Error) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Error](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L65)
{{% /tab %}}
{{< /tabs >}}

### ErrorAs{#erroras}
ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.

This is a wrapper for [errors.As](https://pkg.go.dev/errors#As).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.ErrorAs(t, err, &target)
	success: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
	failure: ErrTest, new(*dummyError)
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestErrorAs(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorAs(t *testing.T)
	success := assert.ErrorAs(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
	fmt.Printf("success: %t\n", success)

}

type dummyError struct {
}

func (d *dummyError) Error() string {
	return "dummy error"
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
// real-world test would inject *testing.T from TestErrorAs(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorAs(t *testing.T)
	require.ErrorAs(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
	fmt.Println("passed")

}

type dummyError struct {
}

func (d *dummyError) Error() string {
	return "dummy error"
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
| [`assert.ErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorAs) | package-level function |
| [`assert.ErrorAsf(t T, err error, target any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorAsf) | formatted variant |
| [`assert.(*Assertions).ErrorAs(err error, target any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorAs) | method variant |
| [`assert.(*Assertions).ErrorAsf(err error, target any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorAsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorAs) | package-level function |
| [`require.ErrorAsf(t T, err error, target any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorAsf) | formatted variant |
| [`require.(*Assertions).ErrorAs(err error, target any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorAs) | method variant |
| [`require.(*Assertions).ErrorAsf(err error, target any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorAsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ErrorAs) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorAs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L220)
{{% /tab %}}
{{< /tabs >}}

### ErrorContains{#errorcontains}
ErrorContains asserts that a function returned a non-nil error (i.e. an
error) and that the error contains the specified substring.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	assertions.ErrorContains(t, err,  expectedErrorSubString)
	success: ErrTest, "general error"
	failure: ErrTest, "not in message"
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestErrorContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorContains(t *testing.T)
	success := assert.ErrorContains(t, assert.ErrTest, "general error")
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
// real-world test would inject *testing.T from TestErrorContains(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorContains(t *testing.T)
	require.ErrorContains(t, require.ErrTest, "general error")
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
| [`assert.ErrorContains(t T, err error, contains string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorContains) | package-level function |
| [`assert.ErrorContainsf(t T, err error, contains string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorContainsf) | formatted variant |
| [`assert.(*Assertions).ErrorContains(err error, contains string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorContains) | method variant |
| [`assert.(*Assertions).ErrorContainsf(err error, contains string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorContainsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ErrorContains(t T, err error, contains string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorContains) | package-level function |
| [`require.ErrorContainsf(t T, err error, contains string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorContainsf) | formatted variant |
| [`require.(*Assertions).ErrorContains(err error, contains string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorContains) | method variant |
| [`require.(*Assertions).ErrorContainsf(err error, contains string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorContainsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ErrorContains(t T, err error, contains string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ErrorContains) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L120)
{{% /tab %}}
{{< /tabs >}}

### ErrorIs{#erroris}
ErrorIs asserts that at least one of the errors in err's chain matches target.

This is a wrapper for [errors.Is](https://pkg.go.dev/errors#Is).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.ErrorIs(t, err, io.EOF)
	success: fmt.Errorf("wrap: %w", io.EOF), io.EOF
	failure: ErrTest, io.EOF
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestErrorIs(t *testing.T)
package main

import (
	"fmt"
	"io"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorIs(t *testing.T)
	success := assert.ErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
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
// real-world test would inject *testing.T from TestErrorIs(t *testing.T)
package main

import (
	"fmt"
	"io"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorIs(t *testing.T)
	require.ErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
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
| [`assert.ErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorIs) | package-level function |
| [`assert.ErrorIsf(t T, err error, target error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ErrorIsf) | formatted variant |
| [`assert.(*Assertions).ErrorIs(err error, target error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorIs) | method variant |
| [`assert.(*Assertions).ErrorIsf(err error, target error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ErrorIsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorIs) | package-level function |
| [`require.ErrorIsf(t T, err error, target error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ErrorIsf) | formatted variant |
| [`require.(*Assertions).ErrorIs(err error, target error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorIs) | method variant |
| [`require.(*Assertions).ErrorIsf(err error, target error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ErrorIsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ErrorIs) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorIs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L149)
{{% /tab %}}
{{< /tabs >}}

### NoError{#noerror}
NoError asserts that a function returned a nil error (i.e. no error).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	if assert.NoError(t, err) {
		assertions.Equal(t, expectedObj, actualObj)
	}
	success: nil
	failure: ErrTest
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNoError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNoError(t *testing.T)
	success := assert.NoError(t, nil)
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
// real-world test would inject *testing.T from TestNoError(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNoError(t *testing.T)
	require.NoError(t, nil)
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
| [`assert.NoError(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoError) | package-level function |
| [`assert.NoErrorf(t T, err error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoErrorf) | formatted variant |
| [`assert.(*Assertions).NoError(err error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoError) | method variant |
| [`assert.(*Assertions).NoErrorf(err error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoErrorf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NoError(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoError) | package-level function |
| [`require.NoErrorf(t T, err error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoErrorf) | formatted variant |
| [`require.(*Assertions).NoError(err error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoError) | method variant |
| [`require.(*Assertions).NoErrorf(err error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoErrorf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NoError(t T, err error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NoError) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoError](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L42)
{{% /tab %}}
{{< /tabs >}}

### NotErrorAs{#noterroras}
NotErrorAs asserts that none of the errors in err's chain matches target,
but if so, sets target to that error value.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotErrorAs(t, err, &target)
	success: ErrTest, new(*dummyError)
	failure: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotErrorAs(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorAs(t *testing.T)
	success := assert.NotErrorAs(t, assert.ErrTest, new(*dummyError))
	fmt.Printf("success: %t\n", success)

}

type dummyError struct {
}

func (d *dummyError) Error() string {
	return "dummy error"
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
// real-world test would inject *testing.T from TestNotErrorAs(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorAs(t *testing.T)
	require.NotErrorAs(t, require.ErrTest, new(*dummyError))
	fmt.Println("passed")

}

type dummyError struct {
}

func (d *dummyError) Error() string {
	return "dummy error"
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
| [`assert.NotErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotErrorAs) | package-level function |
| [`assert.NotErrorAsf(t T, err error, target any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotErrorAsf) | formatted variant |
| [`assert.(*Assertions).NotErrorAs(err error, target any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotErrorAs) | method variant |
| [`assert.(*Assertions).NotErrorAsf(err error, target any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotErrorAsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotErrorAs) | package-level function |
| [`require.NotErrorAsf(t T, err error, target any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotErrorAsf) | formatted variant |
| [`require.(*Assertions).NotErrorAs(err error, target any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotErrorAs) | method variant |
| [`require.(*Assertions).NotErrorAsf(err error, target any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotErrorAsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotErrorAs(t T, err error, target any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotErrorAs) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotErrorAs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L254)
{{% /tab %}}
{{< /tabs >}}

### NotErrorIs{#noterroris}
NotErrorIs asserts that none of the errors in err's chain matches target.

This is a wrapper for [errors.Is](https://pkg.go.dev/errors#Is).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotErrorIs(t, err, io.EOF)
	success: ErrTest, io.EOF
	failure: fmt.Errorf("wrap: %w", io.EOF), io.EOF
```
{{< /tab >}}
{{% tab title="Testable Examples (assert)" %}}
{{% cards %}}
{{% card %}}


*[Copy and click to open Go Playground](https://go.dev/play/)*


```go
// real-world test would inject *testing.T from TestNotErrorIs(t *testing.T)
package main

import (
	"fmt"
	"io"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorIs(t *testing.T)
	success := assert.NotErrorIs(t, assert.ErrTest, io.EOF)
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
// real-world test would inject *testing.T from TestNotErrorIs(t *testing.T)
package main

import (
	"fmt"
	"io"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorIs(t *testing.T)
	require.NotErrorIs(t, require.ErrTest, io.EOF)
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
| [`assert.NotErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotErrorIs) | package-level function |
| [`assert.NotErrorIsf(t T, err error, target error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotErrorIsf) | formatted variant |
| [`assert.(*Assertions).NotErrorIs(err error, target error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotErrorIs) | method variant |
| [`assert.(*Assertions).NotErrorIsf(err error, target error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotErrorIsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotErrorIs) | package-level function |
| [`require.NotErrorIsf(t T, err error, target error, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotErrorIsf) | formatted variant |
| [`require.(*Assertions).NotErrorIs(err error, target error) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotErrorIs) | method variant |
| [`require.(*Assertions).NotErrorIsf(err error, target error, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotErrorIsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.NotErrorIs(t T, err error, target error, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotErrorIs) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotErrorIs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L186)
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
