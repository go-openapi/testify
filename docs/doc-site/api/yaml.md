---
title: "Yaml"
description: "Asserting Yaml Documents"
weight: 18
domains:
  - "yaml"
keywords:
  - "YAMLEq"
  - "YAMLEqf"
  - "YAMLEqBytes"
  - "YAMLEqBytesf"
  - "YAMLEqT"
  - "YAMLEqTf"
  - "YAMLMarshalAsT"
  - "YAMLMarshalAsTf"
  - "YAMLUnmarshalAsT"
  - "YAMLUnmarshalAsTf"
---

Asserting Yaml Documents

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 5 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [YAMLEq](#yamleq) | angles-right
- [YAMLEqBytes](#yamleqbytes) | angles-right
- [YAMLEqT[EDoc, ADoc Text]](#yamleqtedoc-adoc-text) | star | orange
- [YAMLMarshalAsT[EDoc Text]](#yamlmarshalastedoc-text) | star | orange
- [YAMLUnmarshalAsT[Object any, ADoc Text]](#yamlunmarshalastobject-any-adoc-text) | star | orange
```

### YAMLEq{#yamleq}

YAMLEq asserts that two YAML strings are equivalent.

See [YAMLEqBytes](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqBytes).

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

#### Important

By default, this function is disabled and will panic.

To enable it, you should add a blank import like so:

	import(
	  "github.com/go-openapi/testify/enable/yaml/v2"
	)

{{% expand title="Examples" %}}
{{< tabs >}}
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

See [YAMLEqBytes](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEqBytes).

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
| [`assertions.YAMLEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLEqT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLEqT](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L96)
{{% /tab %}}
{{< /tabs >}}

### YAMLMarshalAsT[EDoc Text] {{% icon icon="star" color=orange %}}{#yamlmarshalastedoc-text}

YAMLMarshalAsT wraps [YAMLEq](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLEq) after [yaml.Marshal](https://pkg.go.dev/yaml#Marshal).

The input YAML may be a string or []byte.

It fails if the marshaling returns an error or if the expected YAML bytes differ semantically
from the expected ones.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actual := struct {
		A int `yaml:"a"`
	}{
		A: 10,
	}
	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
```
{{< /tab >}}
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
| [`assert.YAMLMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLMarshalAsT) | package-level function |
| [`assert.YAMLMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLMarshalAsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.YAMLMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLMarshalAsT) | package-level function |
| [`require.YAMLMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLMarshalAsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.YAMLMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLMarshalAsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLMarshalAsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L163)
{{% /tab %}}
{{< /tabs >}}

### YAMLUnmarshalAsT[Object any, ADoc Text] {{% icon icon="star" color=orange %}}{#yamlunmarshalastobject-any-adoc-text}

YAMLUnmarshalAsT wraps [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) after [yaml.Unmarshal](https://pkg.go.dev/yaml#Unmarshal).

The input YAML may be a string or []byte.

It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.

Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
the unmarshaling would take this type to unmarshal as a map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)any.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	expected := struct {
		A int `yaml:"a"`
	}{
		A: 10,
	}
	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
```
{{< /tab >}}
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
| [`assert.YAMLUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLUnmarshalAsT) | package-level function |
| [`assert.YAMLUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#YAMLUnmarshalAsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.YAMLUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLUnmarshalAsT) | package-level function |
| [`require.YAMLUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#YAMLUnmarshalAsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.YAMLUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#YAMLUnmarshalAsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#YAMLUnmarshalAsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/yaml.go#L128)
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
