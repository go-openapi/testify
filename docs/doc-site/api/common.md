---
title: "Common"
description: "Other Uncategorized Helpers"
modified: 2026-01-19
weight: 18
domains:
  - "common"
keywords:
  - "CallerInfo"
  - "CallerInfof"
  - "New"
  - "Newf"
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

This domain exposes 4 functionalities.

```tree
```

---

## Other helpers

### CallerInfo

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

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#CallerInfo](https://github.com/go-openapi/testify/blob/master/internal/assertions/testing.go#L19)

> **Maintainer Note**
> it is not necessary to export CallerInfo. This should remain an internal implementation detail.
{{% /tab %}}
{{< /tabs >}}

### New

New makes a new [Assertions] object for the specified [T].


{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.New(t T) *Assertions`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#New) | package-level function |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.New(t T) *Assertions`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#New) | package-level function |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.New(t T) *Assertions`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#New) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#New](https://github.com/go-openapi/testify/blob/master/internal/assertions/assertion.go#L12)
{{% /tab %}}
{{< /tabs >}}

### ObjectsAreEqual

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

### ObjectsAreEqualValues

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

---

Generated with github.com/go-openapi/testify/codegen/v2

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/codegen/v2 DO NOT EDIT.

Generated on 2026-01-19 (version fbbb078) using codegen version v2.1.9-0.20260119215714-fbbb0787fd81+dirty [sha: fbbb0787fd8131d63f280f85b14e47f7c0dc8ee0]
-->
