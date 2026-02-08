---
title: "Common"
description: "Other Uncategorized Helpers"
weight: 19
domains:
  - "common"
keywords:
  - "CallerInfo"
  - "CallerInfof"
  - "ObjectsAreEqual"
  - "ObjectsAreEqualf"
  - "ObjectsAreEqualValues"
  - "ObjectsAreEqualValuesf"
---

Other Uncategorized Helpers

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 3 functionalities.

```tree
```

---

## Other helpers

### CallerInfo{#callerinfo}
CallerInfo returns an array of strings containing the file and line number
of each stack frame leading from the current test to the assert call that
failed.


{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#CallerInfo) | package-level function |
| [`assert.CallerInfof(t T, , msg string, args ...any) []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#CallerInfof) | formatted variant |
| [`assert.(*Assertions).CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.CallerInfo) | method variant |
| [`assert.(*Assertions).CallerInfof(, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.CallerInfof) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#CallerInfo) | package-level function |
| [`require.CallerInfof(t T, , msg string, args ...any) []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#CallerInfof) | formatted variant |
| [`require.(*Assertions).CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.CallerInfo) | method variant |
| [`require.(*Assertions).CallerInfof(, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.CallerInfof) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#CallerInfo) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#CallerInfo](https://github.com/go-openapi/testify/blob/master/internal/assertions/testing.go#L70)

> **Maintainer Note**
>
> it is not necessary to export CallerInfo. This should remain an internal implementation detail.
>
{{% /tab %}}
{{< /tabs >}}

### ObjectsAreEqual{#objectsareequal}
ObjectsAreEqual determines if two objects are considered equal.

This function does no assertion of any kind.


{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ObjectsAreEqual) | package-level function |
| [`assert.ObjectsAreEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ObjectsAreEqualf) | formatted variant |
| [`assert.(*Assertions).ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ObjectsAreEqual) | method variant |
| [`assert.(*Assertions).ObjectsAreEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ObjectsAreEqualf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqual) | package-level function |
| [`require.ObjectsAreEqualf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqualf) | formatted variant |
| [`require.(*Assertions).ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ObjectsAreEqual) | method variant |
| [`require.(*Assertions).ObjectsAreEqualf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ObjectsAreEqualf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqual) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqual](https://github.com/go-openapi/testify/blob/master/internal/assertions/object.go#L14)
{{% /tab %}}
{{< /tabs >}}

### ObjectsAreEqualValues{#objectsareequalvalues}
ObjectsAreEqualValues gets whether two objects are equal, or if their
values are equal.


{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ObjectsAreEqualValues) | package-level function |
| [`assert.ObjectsAreEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#ObjectsAreEqualValuesf) | formatted variant |
| [`assert.(*Assertions).ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ObjectsAreEqualValues) | method variant |
| [`assert.(*Assertions).ObjectsAreEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.ObjectsAreEqualValuesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqualValues) | package-level function |
| [`require.ObjectsAreEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqualValuesf) | formatted variant |
| [`require.(*Assertions).ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ObjectsAreEqualValues) | method variant |
| [`require.(*Assertions).ObjectsAreEqualValuesf(expected any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.ObjectsAreEqualValuesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/object.go#L38)
{{% /tab %}}
{{< /tabs >}}

---

Generated with github.com/go-openapi/testify/codegen/v2

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/codegen/v2 DO NOT EDIT.
-->
