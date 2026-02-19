---
title: "Testify v2"
type: home
description: 'The v2 our tests wanted'
weight: 1
---

{{% notice info %}}
This is the home of `github.com/go-openapi/testify/v2`, an active, opinionated fork of `github.com/stretchr/testify`.
{{% /notice %}}

## Testify v2 - The v2 our tests wanted

A set of `go` packages that provide tools for _testifying_ (verifying) that your code behaves as you intended.

This is the go-openapi fork of the great [testify](https://github.com/stretchr/testify) package.

### Status

{{% button href="https://github.com/go-openapi/testify/fork" hint="fork me on github" style=primary icon=code-fork %}}Fork me{{% /button %}}
Design and exploration phase. Feedback, contributions and proposals are welcome.

See our [ROADMAP](./project/maintainers/ROADMAP.md).

### Motivation

See [why we wanted a v2](./MOTIVATION.md).

### Getting started

Import this library in your project like so.

```cmd
go get github.com/go-openapi/testify/v2
```

... and start writing tests. Look at our [examples][doc-examples].

### Basic usage

`testify` simplifies your test assertions like so.

{{< cards >}}
{{% card title="Standard library" %}}
```go
import (
        "testing"
    )
    ...
    
    const expected = "expected result"

	  result := printImports(input)
	  if result != expected {
		  t.Errorf(
        "Expected: %s. Got: %s",
        expected, result, 
      )

      return
	  }
```
{{% /card %}}

{{% card title="testify" %}}
```go
import (
        "testing"

        "github.com/go-openapi/testify/v2/require"
    )
    ...

    const expected = "expected result"

	  result := printImports(input)
	  require.Equalf(t, expected, result,
        "Expected: %s. Got: %s", expected, result, 
    )
```
{{% /card %}}
{{< /cards >}}

### Usage with generics

Assertion functions that support go generic types are suffixed with `T` (for "Type safety").
A formatted variant suffixed with `Tf` is also exposed.

Obviously, the `Assertion` type cannot be extended with generic methods, as of `go1.25`.

{{< cards >}}
{{% card title="EqualT" %}}
```go
import (
        "testing"

        "github.com/go-openapi/testify/v2/require"
    )
    ...
    
    const expected = "Hello World"
    var input := "World"

	result := someRamblingTextGeneration(input)
    require.EqualT(t, expected, result)
```
{{% /card %}}
{{% card title="InDeltaT" %}}
```go
import (
        "testing"

        "github.com/go-openapi/testify/v2/require"
    )
    ...
    
    const (
        expected = 1.00
        delta = 1E-6
    )
    var input = 1.01

	result := someComplexComputation(input)
    require.InDeltaT(t, expected, input, delta)
```
{{% /card %}}
{{< /cards >}}

## Licensing

`SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers`

This library ships under the [SPDX-License-Identifier: Apache-2.0](./project/LICENSE.md).

See the license [NOTICE](./project/NOTICE.md), which recalls the licensing terms of all the pieces of software
distributed with this fork, including internalized libraries.

## Contributing

Feel free to submit issues, fork the repository and send pull requests!

{{% notice style="primary" title="Info" icon="info" %}}
Code generation is used. Run `go generate ./...` to update generated files.
{{% /notice %}}

See also our [CONTRIBUTING guidelines](./project/contributing/CONTRIBUTING.md).

---

## See Also

**Getting Started:**
- [Usage Guide](./usage/USAGE.md) - API conventions and how to navigate the documentation
- [Tutorial](./usage/TUTORIAL.md) - Best practices and patterns for writing great tests
- [Examples](./usage/EXAMPLES.md) - Practical code examples for common testing scenarios

**Advanced Topics:**
- [Generics Guide](./usage/GENERICS.md) - Type-safe assertions with generic functions
- [Migration Guide](./usage/MIGRATION.md) - Migrating from stretchr/testify v1
- [Changes from v1](./usage/CHANGES.md) - All changes and improvements in v2
- [Benchmarks](./project/maintainers/benchmarks.md) - Performance improvements in v2

**Reference:**
- [API Reference](./api/_index.md) - Complete assertion catalog organized by domain

---

{{< children type="card" description="true" >}}

[doc-examples]: ./usage/EXAMPLES.md
