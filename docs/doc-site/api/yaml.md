---
title: "Yaml"
description: "Asserting Yaml Documents"
modified: 2026-01-24
weight: 17
domains:
  - "yaml"
keywords:
  - "YAMLEq"
  - "YAMLEqf"
  - "YAMLEqBytes"
  - "YAMLEqBytesf"
  - "YAMLEqT"
  - "YAMLEqTf"
---

Asserting Yaml Documents

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 3 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}

```tree
- [YAMLEq](#yamleq) | angles-right
- [YAMLEqBytes](#yamleqbytes) | angles-right
- [YAMLEqT[EDoc, ADoc Text]](#yamleqtedoc-adoc-text) | star | orange
```

### YAMLEq{#yamleq}

YAMLEq asserts that two YAML strings are equivalent.

See [YAMLEqBytes].

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Examples" %}}
```go
	panic: "key: value", "key: value"
	should panic without the yaml feature enabled.
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.YAMLEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEq) | package-level function |
| [`assert.YAMLEqf(t T, expected string, actual string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqf) | formatted variant |
| [`assert.(*Assertions).YAMLEq(expected string, actual string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.YAMLEq) | method variant |
| [`assert.(*Assertions).YAMLEqf(expected string, actual string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.YAMLEqf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.YAMLEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEq) | package-level function |
| [`require.YAMLEqf(t T, expected string, actual string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEqf) | formatted variant |
| [`require.(*Assertions).YAMLEq(expected string, actual string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.YAMLEq) | method variant |
| [`require.(*Assertions).YAMLEqf(expected string, actual string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.YAMLEqf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.YAMLEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLEq) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLEq](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L77)
{{% /tab %}}
{{< /tabs >}}

### YAMLEqBytes{#yamleqbytes}

YAMLEqBytes asserts that two YAML slices of bytes are equivalent.

Expected and actual must be valid YAML.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Important" %}}
```go
By default, this function is disabled and will panic.
To enable it, you should add a blank import like so:
	import(
	  "github.com/go-openapi/testify/enable/yaml/v2"
	)
```
{{< /tab >}}
{{% tab title="Usage" %}}
```go
	expected := `---
	key: value
	---
	key: this is a second document, it is not evaluated
	`
	actual := `---
	key: value
	---
	key: this is a subsequent document, it is not evaluated
	`
	assertions.YAMLEq(t, expected, actual)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	panic: []byte("key: value"), []byte("key: value")
	should panic without the yaml feature enabled.
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.YAMLEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqBytes) | package-level function |
| [`assert.YAMLEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqBytesf) | formatted variant |
| [`assert.(*Assertions).YAMLEqBytes(expected []byte, actual []byte) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.YAMLEqBytes) | method variant |
| [`assert.(*Assertions).YAMLEqBytesf(expected []byte, actual []byte, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.YAMLEqBytesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.YAMLEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEqBytes) | package-level function |
| [`require.YAMLEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEqBytesf) | formatted variant |
| [`require.(*Assertions).YAMLEqBytes(expected []byte, actual []byte) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.YAMLEqBytes) | method variant |
| [`require.(*Assertions).YAMLEqBytesf(expected []byte, actual []byte, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.YAMLEqBytesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.YAMLEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLEqBytes) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLEqBytes](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L46)
{{% /tab %}}
{{< /tabs >}}

### YAMLEqT[EDoc, ADoc Text] {{% icon icon="star" color=orange %}}{#yamleqtedoc-adoc-text}

YAMLEqT asserts that two YAML documents are equivalent.

The expected and actual arguments may be string or []byte. They do not need to be of the same type.

See [YAMLEqBytes].

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Examples" %}}
```go
	panic: "key: value", "key: value"
	should panic without the yaml feature enabled.
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.YAMLEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqT) | package-level function |
| [`assert.YAMLEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.YAMLEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEqT) | package-level function |
| [`require.YAMLEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLEqTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.YAMLEqT(t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLEqT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLEqT](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L96)
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

Generated on 2026-01-24 (version 178304f) using codegen version v2.1.9-0.20260123222731-178304f36678+dirty [sha: 178304f366789315d4db6b11c89786c43d916247]
-->
