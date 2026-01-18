---
title: "Type"
description: "Asserting Types Rather Than Values"
modified: 2026-01-18
weight: 16
domains:
  - "type"
keywords:
  - "Implements"
  - "Implementsf"
  - "IsNotType"
  - "IsNotTypef"
  - "IsType"
  - "IsTypef"
  - "Kind"
  - "Kindf"
  - "NotImplements"
  - "NotImplementsf"
  - "NotKind"
  - "NotKindf"
  - "NotZero"
  - "NotZerof"
  - "Zero"
  - "Zerof"
---

Asserting Types Rather Than Values

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 8 functionalities.

```tree
- [Implements](#implements) | angles-right
- [IsNotType](#isnottype) | angles-right
- [IsType](#istype) | angles-right
- [Kind](#kind) | angles-right
- [NotImplements](#notimplements) | angles-right
- [NotKind](#notkind) | angles-right
- [NotZero](#notzero) | angles-right
- [Zero](#zero) | angles-right
```

### Implements{#implements}

Implements asserts that an object is implemented by the specified interface.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Implements(t, (*MyInterface)(nil), new(MyObject))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: ptr(dummyInterface), new(testing.T)
	failure: (*error)(nil), new(testing.T)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Implements) | package-level function |
| [`assert.Implementsf(t T, interfaceObject any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Implementsf) | formatted variant |
| [`assert.(*Assertions).Implements(interfaceObject any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Implements) | method variant |
| [`assert.(*Assertions).Implementsf(interfaceObject any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Implementsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Implements) | package-level function |
| [`require.Implementsf(t T, interfaceObject any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Implementsf) | formatted variant |
| [`require.(*Assertions).Implements(interfaceObject any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Implements) | method variant |
| [`require.(*Assertions).Implementsf(interfaceObject any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Implementsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Implements) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Implements](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L21)
{{% /tab %}}
{{< /tabs >}}

### IsNotType{#isnottype}

IsNotType asserts that the specified objects are not of the same type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsNotType(t, &NotMyStruct{}, &MyStruct{})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: int32(123), int64(456)
	failure: 123, 456
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsNotType(t T, theType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNotType) | package-level function |
| [`assert.IsNotTypef(t T, theType any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsNotTypef) | formatted variant |
| [`assert.(*Assertions).IsNotType(theType any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNotType) | method variant |
| [`assert.(*Assertions).IsNotTypef(theType any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsNotTypef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsNotType(t T, theType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNotType) | package-level function |
| [`require.IsNotTypef(t T, theType any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsNotTypef) | formatted variant |
| [`require.(*Assertions).IsNotType(theType any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNotType) | method variant |
| [`require.(*Assertions).IsNotTypef(theType any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsNotTypef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsNotType(t T, theType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsNotType) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsNotType](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L96)
{{% /tab %}}
{{< /tabs >}}

### IsType{#istype}

IsType asserts that the specified objects are of the same type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.IsType(t, &MyStruct{}, &MyStruct{})
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 123, 456
	failure: int32(123), int64(456)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.IsType(t T, expectedType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsType) | package-level function |
| [`assert.IsTypef(t T, expectedType any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#IsTypef) | formatted variant |
| [`assert.(*Assertions).IsType(expectedType any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsType) | method variant |
| [`assert.(*Assertions).IsTypef(expectedType any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.IsTypef) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.IsType(t T, expectedType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsType) | package-level function |
| [`require.IsTypef(t T, expectedType any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#IsTypef) | formatted variant |
| [`require.(*Assertions).IsType(expectedType any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsType) | method variant |
| [`require.(*Assertions).IsTypef(expectedType any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.IsTypef) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.IsType(t T, expectedType any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#IsType) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#IsType](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L75)
{{% /tab %}}
{{< /tabs >}}

### Kind{#kind}

Kind asserts that the [reflect.Kind](https://pkg.go.dev/reflect#Kind) of a given object matches the expected [reflect.Kind](https://pkg.go.dev/reflect#Kind).

Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
are comparable to [reflect.Invalid](https://pkg.go.dev/reflect#Invalid). See also [reflect.Value.Kind](https://pkg.go.dev/reflect#Value.Kind).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Kind(t, reflect.String, "Hello World")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: reflect.String, "hello"
	failure: reflect.String, 0
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Kind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Kind) | package-level function |
| [`assert.Kindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Kindf) | formatted variant |
| [`assert.(*Assertions).Kind(expectedKind reflect.Kind, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Kind) | method variant |
| [`assert.(*Assertions).Kindf(expectedKind reflect.Kind, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Kindf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Kind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Kind) | package-level function |
| [`require.Kindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Kindf) | formatted variant |
| [`require.(*Assertions).Kind(expectedKind reflect.Kind, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Kind) | method variant |
| [`require.(*Assertions).Kindf(expectedKind reflect.Kind, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Kindf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Kind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Kind) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Kind](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L162)
{{% /tab %}}
{{< /tabs >}}

### NotImplements{#notimplements}

NotImplements asserts that an object does not implement the specified interface.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotImplements(t, (*MyInterface)(nil), new(MyObject))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: (*error)(nil), new(testing.T)
	failure: ptr(dummyInterface), new(testing.T)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotImplements) | package-level function |
| [`assert.NotImplementsf(t T, interfaceObject any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotImplementsf) | formatted variant |
| [`assert.(*Assertions).NotImplements(interfaceObject any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotImplements) | method variant |
| [`assert.(*Assertions).NotImplementsf(interfaceObject any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotImplementsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotImplements) | package-level function |
| [`require.NotImplementsf(t T, interfaceObject any, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotImplementsf) | formatted variant |
| [`require.(*Assertions).NotImplements(interfaceObject any, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotImplements) | method variant |
| [`require.(*Assertions).NotImplementsf(interfaceObject any, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotImplementsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotImplements) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotImplements](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L48)
{{% /tab %}}
{{< /tabs >}}

### NotKind{#notkind}

NotKind asserts that the [reflect.Kind](https://pkg.go.dev/reflect#Kind) of a given object does not match the expected [reflect.Kind](https://pkg.go.dev/reflect#Kind).

Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
are comparable to [reflect.Invalid](https://pkg.go.dev/reflect#Invalid). See also [reflect.Value.Kind](https://pkg.go.dev/reflect#Value.Kind).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotKind(t, reflect.Int, "Hello World")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: reflect.String, 0
	failure: reflect.String, "hello"
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotKind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotKind) | package-level function |
| [`assert.NotKindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotKindf) | formatted variant |
| [`assert.(*Assertions).NotKind(expectedKind reflect.Kind, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotKind) | method variant |
| [`assert.(*Assertions).NotKindf(expectedKind reflect.Kind, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotKindf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotKind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotKind) | package-level function |
| [`require.NotKindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotKindf) | formatted variant |
| [`require.(*Assertions).NotKind(expectedKind reflect.Kind, object any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotKind) | method variant |
| [`require.(*Assertions).NotKindf(expectedKind reflect.Kind, object any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotKindf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotKind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotKind) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotKind](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L195)
{{% /tab %}}
{{< /tabs >}}

### NotZero{#notzero}

NotZero asserts that i is not the zero value for its type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotZero(t, obj)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 1
	failure: 0
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NotZero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotZero) | package-level function |
| [`assert.NotZerof(t T, i any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotZerof) | formatted variant |
| [`assert.(*Assertions).NotZero(i any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotZero) | method variant |
| [`assert.(*Assertions).NotZerof(i any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotZerof) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotZero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotZero) | package-level function |
| [`require.NotZerof(t T, i any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotZerof) | formatted variant |
| [`require.(*Assertions).NotZero(i any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotZero) | method variant |
| [`require.(*Assertions).NotZerof(i any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotZerof) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotZero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotZero) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotZero](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L138)
{{% /tab %}}
{{< /tabs >}}

### Zero{#zero}

Zero asserts that i is the zero value for its type.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Zero(t, obj)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: 0
	failure: 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.Zero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Zero) | package-level function |
| [`assert.Zerof(t T, i any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Zerof) | formatted variant |
| [`assert.(*Assertions).Zero(i any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Zero) | method variant |
| [`assert.(*Assertions).Zerof(i any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Zerof) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Zero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Zero) | package-level function |
| [`require.Zerof(t T, i any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Zerof) | formatted variant |
| [`require.(*Assertions).Zero(i any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Zero) | method variant |
| [`require.(*Assertions).Zerof(i any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Zerof) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Zero(t T, i any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Zero) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Zero](https://github.com/go-openapi/testify/blob/master/internal/assertions/type.go#L117)
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

Generated on 2026-01-18 (version e12affe) using codegen version v2.1.9-0.20260118112101-e12affef2419+dirty [sha: e12affef24198e72ee13eb6d25018d2c3232629f]
-->
