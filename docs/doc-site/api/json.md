---
title: "Json"
description: "Asserting JSON Documents"
weight: 9
domains:
  - "json"
keywords:
  - "JSONEq"
  - "JSONEqf"
  - "JSONEqBytes"
  - "JSONEqBytesf"
  - "JSONEqT"
  - "JSONEqTf"
  - "JSONMarshalAsT"
  - "JSONMarshalAsTf"
  - "JSONUnmarshalAsT"
  - "JSONUnmarshalAsTf"
---

Asserting JSON Documents

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 5 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [JSONEq](#jsoneq) | angles-right
- [JSONEqBytes](#jsoneqbytes) | angles-right
- [JSONEqT[EDoc, ADoc Text]](#jsoneqtedoc-adoc-text) | star | orange
- [JSONMarshalAsT[EDoc Text]](#jsonmarshalastedoc-text) | star | orange
- [JSONUnmarshalAsT[ADoc Text, Object any]](#jsonunmarshalastadoc-text-object-any) | star | orange
```

### JSONEq{#jsoneq}

JSONEq asserts that two JSON strings are semantically equivalent.

Expected and actual must be valid JSON.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`
	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.JSONEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEq) | package-level function |
| [`assert.JSONEqf(t T, expected string, actual string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEqf) | formatted variant |
| [`assert.(*Assertions).JSONEq(expected string, actual string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.JSONEq) | method variant |
| [`assert.(*Assertions).JSONEqf(expected string, actual string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.JSONEqf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.JSONEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEq) | package-level function |
| [`require.JSONEqf(t T, expected string, actual string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEqf) | formatted variant |
| [`require.(*Assertions).JSONEq(expected string, actual string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.JSONEq) | method variant |
| [`require.(*Assertions).JSONEqf(expected string, actual string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.JSONEqf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.JSONEq(t T, expected string, actual string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#JSONEq) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#JSONEq](https://github.com/go-openapi/testify/blob/master/internal/assertions/json.go#L62)
{{% /tab %}}
{{< /tabs >}}

### JSONEqBytes{#jsoneqbytes}

JSONEqBytes asserts that two JSON slices of bytes are semantically equivalent.

Expected and actual must be valid JSON.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`)
	failure: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`)
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.JSONEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEqBytes) | package-level function |
| [`assert.JSONEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEqBytesf) | formatted variant |
| [`assert.(*Assertions).JSONEqBytes(expected []byte, actual []byte) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.JSONEqBytes) | method variant |
| [`assert.(*Assertions).JSONEqBytesf(expected []byte, actual []byte, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.JSONEqBytesf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.JSONEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEqBytes) | package-level function |
| [`require.JSONEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEqBytesf) | formatted variant |
| [`require.(*Assertions).JSONEqBytes(expected []byte, actual []byte) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.JSONEqBytes) | method variant |
| [`require.(*Assertions).JSONEqBytesf(expected []byte, actual []byte, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.JSONEqBytesf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.JSONEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#JSONEqBytes) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#JSONEqBytes](https://github.com/go-openapi/testify/blob/master/internal/assertions/json.go#L24)

> **Maintainer Note**
>
> Proposal for enhancement.
>
> We could use and indirection for users to inject their favorite JSON
> library like we do for YAML.
>
{{% /tab %}}
{{< /tabs >}}

### JSONEqT[EDoc, ADoc Text] {{% icon icon="star" color=orange %}}{#jsoneqtedoc-adoc-text}

JSONEqT asserts that two JSON documents are semantically equivalent.

The expected and actual arguments may be string or []byte. They do not need to be of the same type.

Expected and actual must be valid JSON.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`)
	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.JSONEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEqT) | package-level function |
| [`assert.JSONEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEqTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.JSONEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEqT) | package-level function |
| [`require.JSONEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONEqTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.JSONEqT(t T, expected EDoc, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#JSONEqT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#JSONEqT](https://github.com/go-openapi/testify/blob/master/internal/assertions/json.go#L85)
{{% /tab %}}
{{< /tabs >}}

### JSONMarshalAsT[EDoc Text] {{% icon icon="star" color=orange %}}{#jsonmarshalastedoc-text}

JSONMarshalAsT wraps [JSONEq](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONEq) after [json.Marshal](https://pkg.go.dev/json#Marshal).

The input JSON may be a string or []byte.

It fails if the marshaling returns an error or if the expected JSON bytes differ semantically
from the expected ones.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	actual := struct {
		A int `json:"a"`
	}{
		A: 10
	}
	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: []byte(`{"A": "a"}`), dummyStruct{A: "a"}
	failure: `[{"foo": "bar"}, {"hello": "world"}]`, 1
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.JSONMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONMarshalAsT) | package-level function |
| [`assert.JSONMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONMarshalAsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.JSONMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONMarshalAsT) | package-level function |
| [`require.JSONMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONMarshalAsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.JSONMarshalAsT(t T, expected EDoc, object any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#JSONMarshalAsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#JSONMarshalAsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/json.go#L152)
{{% /tab %}}
{{< /tabs >}}

### JSONUnmarshalAsT[ADoc Text, Object any] {{% icon icon="star" color=orange %}}{#jsonunmarshalastadoc-text-object-any}

JSONUnmarshalAsT wraps [Equal](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Equal) after [json.Unmarshal](https://pkg.go.dev/json#Unmarshal).

The input JSON may be a string or []byte.

It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.

Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
the unmarshaling would take this type to unmarshal as a map[string](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#string)any.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	expected := struct {
		A int `json:"a"`
	}{
		A: 10
	}
	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: dummyStruct{A: "a"} , []byte(`{"A": "a"}`)
	failure: 1, `[{"foo": "bar"}, {"hello": "world"}]`
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.JSONUnmarshalAsT[ADoc Text, Object any](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONUnmarshalAsT) | package-level function |
| [`assert.JSONUnmarshalAsTf[ADoc Text, Object any](t T, expected Object, jazon ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#JSONUnmarshalAsTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.JSONUnmarshalAsT[ADoc Text, Object any](t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONUnmarshalAsT) | package-level function |
| [`require.JSONUnmarshalAsTf[ADoc Text, Object any](t T, expected Object, jazon ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#JSONUnmarshalAsTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.JSONUnmarshalAsT(t T, expected Object, jazon ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#JSONUnmarshalAsT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#JSONUnmarshalAsT](https://github.com/go-openapi/testify/blob/master/internal/assertions/json.go#L117)
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
