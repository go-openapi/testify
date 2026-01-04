---
title: "Testify v2"
type: home
description: 'Go testing assertions for the rest of us'
weight: 1
---

**Go testing assertions for the rest of us**

{{% notice info %}}
This is the home of `github.com/go-openapi/testify/v2`, an active, opinionated fork of `github.com/stretchr/testify`.
{{% /notice %}}

## Testify v2 - The v2 our tests wanted

A set of `go` packages that provide tools for testifying that your code behaves as you intended.

This is the go-openapi fork of the great [testify](https://github.com/stretchr/testify) package.

### Status

{{% button href="https://github.com/go-openapi/testify/fork" hint="fork me on github" style=primary icon=code-fork %}}Fork me{{% /button %}}
Design and exploration phase. Contributions and proposals are welcome.

### Motivation

From the maintainers of `testify`, it looks like a v2 will eventually be released, but they'll do it at their own pace.

We like all the principles they put forward to build this v2. [See discussion about v2](https://github.com/stretchr/testify/discussions/1560)

However, at `go-openapi` we would like to address the well-known issues in `testify` with different priorities.

With this fork, we want to:
1. remove all external dependencies.
2. make it easy to maintain and extend.
3. pare down some of the chrome that has been added over the years.

{{% notice style="primary" title="Extended hand" icon="hand" %}}
We hope that this endeavor will help the original project with a live-drill of what a v2 could look like.
Hopefully, some of our ideas here will eventually percolate back into the original project.

We are always happy to discuss with people who face the same problems as we do: avoid breaking changes, 
APIs that became bloated over a decade or so, uncontrolled dependencies, difficult choices when it comes to introduce
breaking changes, conflicting demands from users etc.
{{% /notice %}}

More about our motivations in the project's [README](README.md).

### Getting started

To use this package in your projects:

```cmd
    go get github.com/go-openapi/testify/v2
```

... and start writing tests. Look at our [examples](./examples/).

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

	require.Equalf(t,
        expected, printImports(input),
        "Expected: %s. Got: %s",
        expected, result, 
    )
```
{{% /card %}}
{{< /cards >}}

## Licensing

`SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers`

This library ships under the [SPDX-License-Identifier: Apache-2.0](./project/LICENSE.md).

See the license [NOTICE](./project/NOTICE.md), which recalls the licensing terms of all the pieces of software
distributed with this fork, including internalized libraries.

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
Extra credit for those using Testify to write the test code that demonstrates it.

{{% notice style="primary" title="Info" icon="info" %}}
Code generation is used. Run `go generate ./...` to update generated files.
{{% /notice %}}

See also the [CONTRIBUTING guidelines](./project/contributing/CONTRIBUTING.md).

---

{{< children type="card" description="true" >}}
