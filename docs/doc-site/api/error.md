---
title: "Error"
description: "Asserting Errors"
modified: 2026-01-20
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest, "assert.ErrTest general error for testing"
	failure: ErrTest, "wrong error message"
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualError](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L80)
{{% /tab %}}
{{< /tabs >}}

### Error{#error}

Error asserts that a function returned a non-nil error (ie. an error).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	assertions.Error(t, err)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest
	failure: nil
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Error](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L56)
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
	failure: ErrTest, new(*dummyError)
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorAs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L211)
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest, "general error"
	failure: ErrTest, "not in message"
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorContains](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L111)
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: fmt.Errorf("wrap: %w", io.EOF), io.EOF
	failure: ErrTest, io.EOF
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ErrorIs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L140)
{{% /tab %}}
{{< /tabs >}}

### NoError{#noerror}

NoError asserts that a function returned a nil error (ie. no error).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actualObj, err := SomeFunction()
	if assert.NoError(t, err) {
		assertions.Equal(t, expectedObj, actualObj)
	}
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: nil
	failure: ErrTest
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoError](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L33)
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest, new(*dummyError)
	failure: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotErrorAs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L245)
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
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ErrTest, io.EOF
	failure: fmt.Errorf("wrap: %w", io.EOF), io.EOF
```
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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotErrorIs](https://github.com/go-openapi/testify/blob/master/internal/assertions/error.go#L177)
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

Generated on 2026-01-20 (version 74d5686) using codegen version v2.1.9-0.20260119232631-74d5686313f0+dirty [sha: 74d5686313f0820ae0e2758b95d598f646cd7ad5]
-->
