---
title: "String"
description: "Asserting Strings"
weight: 13
domains:
  - "string"
keywords:
  - "NotRegexp"
  - "NotRegexpf"
  - "NotRegexpT"
  - "NotRegexpTf"
  - "Regexp"
  - "Regexpf"
  - "RegexpT"
  - "RegexpTf"
---

Asserting Strings

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 4 functionalities.
Generic assertions are marked with a {{% icon icon="star" color=orange %}}.

```tree
- [NotRegexp](#notregexp) | angles-right
- [NotRegexpT[Rex RegExp, ADoc Text]](#notregexptrex-regexp-adoc-text) | star | orange
- [Regexp](#regexp) | angles-right
- [RegexpT[Rex RegExp, ADoc Text]](#regexptrex-regexp-adoc-text) | star | orange
```

### NotRegexp{#notregexp}

NotRegexp asserts that a specified regular expression does not match a string.

See [Regexp](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexp).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
	assertions.NotRegexp(t, "^start", "it's not starting")
	success: "^start", "not starting"
	failure: "^start", "starting"
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotRegexp(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotRegexp(t, "^start", "not starting")
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
| [`assert.NotRegexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexp) | package-level function |
| [`assert.NotRegexpf(t T, rx any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexpf) | formatted variant |
| [`assert.(*Assertions).NotRegexp(rx any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotRegexp) | method variant |
| [`assert.(*Assertions).NotRegexpf(rx any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NotRegexpf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotRegexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexp) | package-level function |
| [`require.NotRegexpf(t T, rx any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexpf) | formatted variant |
| [`require.(*Assertions).NotRegexp(rx any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotRegexp) | method variant |
| [`require.(*Assertions).NotRegexpf(rx any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NotRegexpf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotRegexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotRegexp) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotRegexp](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L90)
{{% /tab %}}
{{< /tabs >}}

### NotRegexpT[Rex RegExp, ADoc Text] {{% icon icon="star" color=orange %}}{#notregexptrex-regexp-adoc-text}

NotRegexpT asserts that a specified regular expression does not match a string.

See [RegexpT](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#RegexpT).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
	assertions.NotRegexp(t, "^start", "it's not starting")
	success: "^start", "not starting"
	failure: "^start", "starting"
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestNotRegexpT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.NotRegexpT(t, "^start", "not starting")
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
| [`assert.NotRegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexpT) | package-level function |
| [`assert.NotRegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NotRegexpTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NotRegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexpT) | package-level function |
| [`require.NotRegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NotRegexpTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NotRegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NotRegexpT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NotRegexpT](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L131)
{{% /tab %}}
{{< /tabs >}}

### Regexp{#regexp}

Regexp asserts that a specified regular expression matches a string.

The regular expression may be passed as a [regexp.Regexp](https://pkg.go.dev/regexp#Regexp), a string or a []byte and will be compiled.

The actual argument to be matched may be a string, []byte or anything that prints as a string with [fmt.Sprint](https://pkg.go.dev/fmt#Sprint).

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.Regexp(t, regexp.MustCompile("start"), "it's starting")
	assertions.Regexp(t, "start...$", "it's not starting")
	success: "^start", "starting"
	failure: "^start", "not starting"
```
{{< /tab >}}
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestRegexp(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.Regexp(t, "^start", "starting")
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
| [`assert.Regexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexp) | package-level function |
| [`assert.Regexpf(t T, rx any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexpf) | formatted variant |
| [`assert.(*Assertions).Regexp(rx any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Regexp) | method variant |
| [`assert.(*Assertions).Regexpf(rx any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.Regexpf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.Regexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Regexp) | package-level function |
| [`require.Regexpf(t T, rx any, actual any, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Regexpf) | formatted variant |
| [`require.(*Assertions).Regexp(rx any, actual any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Regexp) | method variant |
| [`require.(*Assertions).Regexpf(rx any, actual any, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.Regexpf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.Regexp(t T, rx any, actual any, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#Regexp) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#Regexp](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L27)
{{% /tab %}}
{{< /tabs >}}

### RegexpT[Rex RegExp, ADoc Text] {{% icon icon="star" color=orange %}}{#regexptrex-regexp-adoc-text}

RegexpT asserts that a specified regular expression matches a string.

The actual argument to be matched may be a string or []byte.

See [Regexp](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Regexp).

{{% expand title="Examples" %}}
{{< tabs >}}
	success: "^start", "starting"
	failure: "^start", "not starting"
{{% tab title="Testable Examples" %}}
{{% cards %}}
{{% card href="https://go.dev/play/" %}}


*Copy and click to open Go Playground*


```go
// real-world test would inject *testing.T from TestRegexpT(t *testing.T)
package main

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func main() {
	t := new(testing.T)
	require.RegexpT(t, "^start", "starting")
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
| [`assert.RegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#RegexpT) | package-level function |
| [`assert.RegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#RegexpTf) | formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.RegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#RegexpT) | package-level function |
| [`require.RegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#RegexpTf) | formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.RegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#RegexpT) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#RegexpT](https://github.com/go-openapi/testify/blob/master/internal/assertions/string.go#L63)
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
