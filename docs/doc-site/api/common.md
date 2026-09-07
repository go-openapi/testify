---
title: "Common"
description: "Other Uncategorized Helpers"
weight: 20
domains:
  - "common"
keywords:
  - "CallerInfo"
  - "ObjectsAreEqual"
  - "ObjectsAreEqualValues"
  - "WithHunkSize"
---

Other Uncategorized Helpers

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.

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
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#CallerInfo) | package-level function |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.CallerInfo() []string`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#CallerInfo) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#CallerInfo](https://github.com/go-openapi/testify/blob/master/internal/assertions/testing.go#L71)

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
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ObjectsAreEqual(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqual) | package-level function |
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
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#ObjectsAreEqualValues) | package-level function |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.ObjectsAreEqualValues(expected any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqualValues) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#ObjectsAreEqualValues](https://github.com/go-openapi/testify/blob/master/internal/assertions/object.go#L38)
{{% /tab %}}
{{< /tabs >}}

### WithHunkSize{#withhunksize}
WithHunkSize sets how many unchanged lines the diff shows around each change.

The diff appears in the failure message of [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal), [EqualT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualT), [EqualValues](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualValues),
[EqualExportedValues](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#EqualExportedValues) and [Exactly](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Exactly), whenever both values are a struct, map, slice, array
or string. The default is 1. A value below 1 is clamped to 1, and a value larger than the
rendered value prints it whole.

Pass it to [New](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#New), which is the only place options are read:

	a := assert.New(t, assert.WithHunkSize(4))
	a.Equal(expected, actual)


{{< tabs >}}
  
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.WithHunkSize(n int) Option`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#WithHunkSize) | package-level function |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.WithHunkSize(n int) Option`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#WithHunkSize) | package-level function |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--|
| [`assertions.WithHunkSize(n int) Option`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#WithHunkSize) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#WithHunkSize](https://github.com/go-openapi/testify/blob/master/internal/assertions/options.go#L45)
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
