---
title: "Equality"
description: "Asserting Two Things Are Equal"
modified: 2026-01-02
weight: 5
domains:
  - "equality"
keywords:
  - "Empty"
  - "Emptyf"
  - "Equal"
  - "Equalf"
  - "EqualExportedValues"
  - "EqualExportedValuesf"
  - "EqualValues"
  - "EqualValuesf"
  - "Exactly"
  - "Exactlyf"
  - "Nil"
  - "Nilf"
  - "NotEmpty"
  - "NotEmptyf"
  - "NotEqual"
  - "NotEqualf"
  - "NotEqualValues"
  - "NotEqualValuesf"
  - "NotNil"
  - "NotNilf"
  - "NotSame"
  - "NotSamef"
  - "Same"
  - "Samef"
---

Asserting Two Things Are Equal

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 12 functionalities.

### Empty

Empty asserts that the given value is "empty".

Zero values are "empty".

Arrays are "empty" if every element is the zero value of the type (stricter than "empty").

Slices, maps and channels with zero length are "empty".

Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Empty(t, obj)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ""
	failure: "not empty"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

[Zero values]: https://go.dev/ref/spec#The_zero_value

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Empty) | package-level function |
| [`assert.Emptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Emptyf) | formatted variant |
| [`assert.(*Assertions).Empty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Empty) | method variant |
| [`assert.(*Assertions).Emptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Emptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Empty) | package-level function |
| [`require.Emptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Emptyf) | formatted variant |
| [`require.(*Assertions).Empty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Empty) | method variant |
| [`require.(*Assertions).Emptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Emptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Empty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Empty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Empty](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L276)
{{% /tab %}}
{{< /tabs >}}

### Equal

Equal asserts that two objects are equal.

Pointer variable equality is determined based on the equality of the
referenced values (as opposed to the memory addresses).

Function equality cannot be determined and will always fail.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Equal(t, 123, 123)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 123, 123
	failure: 123, 456
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) | package-level function |
| [`assert.Equalf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equalf) | formatted variant |
| [`assert.(*Assertions).Equal(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Equal) | method variant |
| [`assert.(*Assertions).Equalf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Equalf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Equal) | package-level function |
| [`require.Equalf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Equalf) | formatted variant |
| [`require.(*Assertions).Equal(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Equal) | method variant |
| [`require.(*Assertions).Equalf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Equalf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Equal(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Equal) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Equal](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L28)
{{% /tab %}}
{{< /tabs >}}

### EqualExportedValues

EqualExportedValues asserts that the types of two objects are equal and their public
fields are also equal. This is useful for comparing structs that have private fields
that could potentially differ.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	 type S struct {
		Exported     	int
		notExported   	int
	 }
	assertions.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
	assertions.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}
	failure:  &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualExportedValues) | package-level function |
| [`assert.EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualExportedValuesf) | formatted variant |
| [`assert.(*Assertions).EqualExportedValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualExportedValues) | method variant |
| [`assert.(*Assertions).EqualExportedValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualExportedValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualExportedValues) | package-level function |
| [`require.EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualExportedValuesf) | formatted variant |
| [`require.(*Assertions).EqualExportedValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualExportedValues) | method variant |
| [`require.(*Assertions).EqualExportedValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualExportedValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualExportedValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualExportedValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L161)
{{% /tab %}}
{{< /tabs >}}

### EqualValues

EqualValues asserts that two objects are equal or convertible to the larger
type and equal.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.EqualValues(t, uint32(123), int32(123))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: uint32(123), int32(123)
	failure: uint32(123), int32(456)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualValues) | package-level function |
| [`assert.EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualValuesf) | formatted variant |
| [`assert.(*Assertions).EqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualValues) | method variant |
| [`assert.(*Assertions).EqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.EqualValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualValues) | package-level function |
| [`require.EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#EqualValuesf) | formatted variant |
| [`require.(*Assertions).EqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualValues) | method variant |
| [`require.(*Assertions).EqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.EqualValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#EqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#EqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L127)
{{% /tab %}}
{{< /tabs >}}

### Exactly

Exactly asserts that two objects are equal in value and type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Exactly(t, int32(123), int64(123))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: int32(123), int32(123)
	failure: int32(123), int64(123)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Exactly) | package-level function |
| [`assert.Exactlyf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Exactlyf) | formatted variant |
| [`assert.(*Assertions).Exactly(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Exactly) | method variant |
| [`assert.(*Assertions).Exactlyf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Exactlyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Exactly) | package-level function |
| [`require.Exactlyf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Exactlyf) | formatted variant |
| [`require.(*Assertions).Exactly(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Exactly) | method variant |
| [`require.(*Assertions).Exactlyf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Exactlyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Exactly(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Exactly) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Exactly](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L198)
{{% /tab %}}
{{< /tabs >}}

### Nil

Nil asserts that the specified object is nil.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Nil(t, err)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: nil
	failure: "not nil"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Nil) | package-level function |
| [`assert.Nilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Nilf) | formatted variant |
| [`assert.(*Assertions).Nil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Nil) | method variant |
| [`assert.(*Assertions).Nilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Nilf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Nil) | package-level function |
| [`require.Nilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Nilf) | formatted variant |
| [`require.(*Assertions).Nil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Nil) | method variant |
| [`require.(*Assertions).Nilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Nilf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Nil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Nil) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Nil](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L245)
{{% /tab %}}
{{< /tabs >}}

### NotEmpty

NotEmpty asserts that the specified object is NOT [Empty].

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	if assert.NotEmpty(t, obj) {
		assertions.Equal(t, "two", obj[1])
	}
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: "not empty"
	failure: ""
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEmpty) | package-level function |
| [`assert.NotEmptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEmptyf) | formatted variant |
| [`assert.(*Assertions).NotEmpty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEmpty) | method variant |
| [`assert.(*Assertions).NotEmptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEmptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEmpty) | package-level function |
| [`require.NotEmptyf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEmptyf) | formatted variant |
| [`require.(*Assertions).NotEmpty(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEmpty) | method variant |
| [`require.(*Assertions).NotEmptyf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEmptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotEmpty(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEmpty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEmpty](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L301)
{{% /tab %}}
{{< /tabs >}}

### NotEqual

NotEqual asserts that the specified values are NOT equal.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotEqual(t, obj1, obj2)
Pointer variable equality is determined based on the equality of the
referenced values (as opposed to the memory addresses).
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 123, 456
	failure: 123, 123
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqual) | package-level function |
| [`assert.NotEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualf) | formatted variant |
| [`assert.(*Assertions).NotEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqual) | method variant |
| [`assert.(*Assertions).NotEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqual) | package-level function |
| [`require.NotEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualf) | formatted variant |
| [`require.(*Assertions).NotEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqual) | method variant |
| [`require.(*Assertions).NotEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEqual) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L327)
{{% /tab %}}
{{< /tabs >}}

### NotEqualValues

NotEqualValues asserts that two objects are not equal even when converted to the same type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotEqualValues(t, obj1, obj2)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: uint32(123), int32(456)
	failure: uint32(123), int32(123)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualValues) | package-level function |
| [`assert.NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotEqualValuesf) | formatted variant |
| [`assert.(*Assertions).NotEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualValues) | method variant |
| [`assert.(*Assertions).NotEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotEqualValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualValues) | package-level function |
| [`require.NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotEqualValuesf) | formatted variant |
| [`require.(*Assertions).NotEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualValues) | method variant |
| [`require.(*Assertions).NotEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotEqualValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotEqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotEqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L354)
{{% /tab %}}
{{< /tabs >}}

### NotNil

NotNil asserts that the specified object is not nil.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
assertions.NotNil(t, err)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: "not nil"
	failure: nil
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotNil) | package-level function |
| [`assert.NotNilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotNilf) | formatted variant |
| [`assert.(*Assertions).NotNil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotNil) | method variant |
| [`assert.(*Assertions).NotNilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotNilf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotNil) | package-level function |
| [`require.NotNilf(t T, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotNilf) | formatted variant |
| [`require.(*Assertions).NotNil(object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotNil) | method variant |
| [`require.(*Assertions).NotNilf(object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotNilf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotNil(t T, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotNil) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotNil](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L224)
{{% /tab %}}
{{< /tabs >}}

### NotSame

NotSame asserts that two pointers do not reference the same object.

Both arguments must be pointer variables. Pointer variable sameness is
determined based on the equality of both type and value.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotSame(t, ptr1, ptr2)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: &staticVar, ptr("static string")
	failure: &staticVar, staticVarPtr
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSame) | package-level function |
| [`assert.NotSamef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotSamef) | formatted variant |
| [`assert.(*Assertions).NotSame(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSame) | method variant |
| [`assert.(*Assertions).NotSamef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotSamef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSame) | package-level function |
| [`require.NotSamef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotSamef) | formatted variant |
| [`require.(*Assertions).NotSame(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSame) | method variant |
| [`require.(*Assertions).NotSamef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotSamef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotSame(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotSame) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotSame](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L96)
{{% /tab %}}
{{< /tabs >}}

### Same

Same asserts that two pointers reference the same object.

Both arguments must be pointer variables. Pointer variable sameness is
determined based on the equality of both type and value.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Same(t, ptr1, ptr2)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: &staticVar, staticVarPtr
	failure: &staticVar, ptr("static string")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Same) | package-level function |
| [`assert.Samef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Samef) | formatted variant |
| [`assert.(*Assertions).Same(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Same) | method variant |
| [`assert.(*Assertions).Samef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Samef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Same) | package-level function |
| [`require.Samef(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Samef) | formatted variant |
| [`require.(*Assertions).Same(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Same) | method variant |
| [`require.(*Assertions).Samef(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Samef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Same(t T, expected any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Same) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Same](https://github.com/go-openapi/testify/blob/master/internal/assertions/equal.go#L62)
{{% /tab %}}
{{< /tabs >}}

---

---

Generated with github.com/go-openapi/testify/v2/codegen

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/v2/codegen DO NOT EDIT.

Generated on 2026-01-02 (version v1.2.2-760-g97c29e3) using codegen version master [sha: 97c29e3dbfc40800a080863ceea81db0cfd6e858]
-->
