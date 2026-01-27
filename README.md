# Testify/v2

<!-- Badges: s[roadmap]tatus  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vuln scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
<!-- Badges: release & docker images  -->
<!-- Badges: code quality  -->
<!-- Badges: license & compliance -->
[![Release][release-badge]][release-url] [![Go Report Card][gocard-badge]][gocard-url] [![CodeFactor Grade][codefactor-badge]][codefactor-url] [![License][license-badge]][license-url]
<!-- Badges: documentation & support -->
<!-- Badges: others & stats -->
[![Doc][doc-badge]][doc-url] [![GoDoc][godoc-badge]][godoc-url] [![Discord Channel][discord-badge]][discord-url] [![go version][goversion-badge]][goversion-url] ![Top language][top-badge] ![Commits since latest release][commits-badge]

---

**The v2 our tests wanted**

A set of `go` packages that provide tools for testifying (verifying) that your code behaves as you intended.

This is the go-openapi fork of the great [testify](https://github.com/stretchr/testify) package.

> [!NOTE]
> This is the home of `github.com/go-openapi/testify/v2`, an active, opinionated fork of `github.com/stretchr/testify`.

Main features:

* zero external dependencies
* opt-in dependencies for extra features (e.g. asserting YAML, colorized output)
* assertions using generic types (see [a basic example][example-with-generics-url]). [Read the fully story with generics][doc-generics]
* [searchable documentation][doc-url]

## Announcements

* **2025-12-19** : new community chat on discord
  * a new discord community channel is available to be notified of changes and support users
  * our venerable Slack channel remains open, and will be eventually discontinued on **2026-03-31**

You may join the discord community by clicking the invite link on the discord badge (also above). [![Discord Channel][discord-badge]][discord-url]

Or join our Slack channel: [![Slack Channel][slack-logo]![slack-badge]][slack-url]

### Status

Design and exploration phase. Feedback, contributions and proposals are welcome.

> **Recent news**
>
> ✅ Fully refactored how assertions are generated and documented.
>
> ✅ Fixed hangs & panics when using `spew`. Fuzzed `spew`.
>
> ✅  Fixed go routine leaks with `EventuallyWithT` and co.
>
> ✅ Added `Kind` & `NotKind`
>
> ✅ Fix deterministic order of keys in diff
>
> ✅ Fixed edge cases with `InDelta`, `InEpsilon`
>
> ✅ Fixed edge cases with `EqualValues`
>
> ✅ Fixed wrong logic with `IsNonIncreasing`, `InNonDecreasing`
>
> ✅ Added opt-in support for colorized output
>
> ✅ Introduced generics: 38 new type-safe assertions with generic types (doc: added usage guide, examples and benchmark)
>
> See also our [ROADMAP][doc-roadmap].

## Getting started

Import this library in your project like so.

```cmd
go get github.com/go-openapi/testify/v2
```

... and start writing tests. Look at our [examples][doc-examples].

## Basic usage

`testify` simplifies your test assertions like so.

```go
    import (
        "testing"
    )
    ...
    
    const expected = "expected result"

	result := printImports(input)
	if result != expected {
		t.Errorf(
            "Expected: %s. Got: %s", expected, result, 
        )

        return
	}
```

Becomes:

```go
    import (
        "testing"
        "github.com/go-openapi/testify/v2/require"
    )
    ...

    const expected = "expected result"

	require.Equalf(t,
        expected, printImports(input), "Expected: %s. Got: %s",
        expected, result, 
    )
```

## Usage at go-openapi and go-swagger

This fork now fully replaces the original project for all go-openapi projects,
thus reducing their dependencies footprint.

Go-swagger will be adapted over Q1 2026.

Features will be added to support our main use cases there.

## Change log

See <https://github.com/go-openapi/testify/releases>

## Licensing

`SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers`

This library ships under the [SPDX-License-Identifier: Apache-2.0](./LICENSE).

See the license [NOTICE](./NOTICE), which recalls the licensing terms of all the pieces of software
distributed with this fork, including internalized libraries.

## Other documentation

* [Getting started][doc-examples]
* [Usage](https://go-openapi.github.io/testify/usage/)
* [Motivations](https://go-openapi.github.io/testify/project/readme)
* [Roadmap][doc-roadmap]
* [Internal architecture](https://go-openapi.github.io/testify/project/maintainers/architecture)
* [Benchmarks](https://go-openapi.github.io/testify/project/maintainers/benchmarks)

* [All-time contributors](./CONTRIBUTORS.md)
* [Contributing guidelines](https://go-openapi.github.io/testify/project/contributing/)
* [Maintainers documentation](https://go-openapi.github.io/testify/project/maintainers/)
* [Coding style](https://go-openapi.github.io/testify/project/contributing/style)
* [Security policy](https://go-openapi.github.io/testify/project/security)

## Cutting a new release

Maintainers can cut a new release by either:

* running [this workflow][ci-release-workflow] (recommended)
* or :
  1. preparing go.mod files with the next tag, merge
  2. pushing a semver tag
  * signed tags are preferred
  * The tag message is prepended to release notes

<!-- Doc links -->
[doc-roadmap]: https://go-openapi.github.io/testify/project/maintainers/roadmap
<!-- Badges: status  -->
[test-badge]: https://github.com/go-openapi/testify/actions/workflows/go-test.yml/badge.svg
[test-url]: https://github.com/go-openapi/testify/actions/workflows/go-test.yml
[cov-badge]: https://codecov.io/gh/go-openapi/testify/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/go-openapi/testify
[vuln-scan-badge]: https://github.com/go-openapi/testify/actions/workflows/scanner.yml/badge.svg
[vuln-scan-url]: https://github.com/go-openapi/testify/actions/workflows/scanner.yml
[codeql-badge]: https://github.com/go-openapi/testify/actions/workflows/codeql.yml/badge.svg
[codeql-url]: https://github.com/go-openapi/testify/actions/workflows/codeql.yml
<!-- Badges: release & docker images  -->
[release-badge]: https://badge.fury.io/gh/go-openapi%2Ftestify.svg
[release-url]: https://badge.fury.io/gh/go-openapi%2Ftestify
[gomod-badge]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Ftestify.svg
[gomod-url]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Ftestify
<!-- Badges: code quality  -->
[gocard-badge]: https://goreportcard.com/badge/github.com/go-openapi/testify
[gocard-url]: https://goreportcard.com/report/github.com/go-openapi/testify
[codefactor-badge]: https://img.shields.io/codefactor/grade/github/go-openapi/testify
[codefactor-url]: https://www.codefactor.io/repository/github/go-openapi/testify
<!-- Badges: documentation & support -->
[doc-badge]: https://img.shields.io/badge/doc-site-blue?link=https%3A%2F%2Fgo-openapi.github.io%2Ftestify%2F
[doc-url]: https://go-openapi.github.io/testify
[doc-examples]: https://go-openapi.github.io/testify/usage/examples
[doc-generics]: https://go-openapi.github.io/testify/usage/generics
[example-with-generics-url]: https://go-openapi.github.io/testify#usage-with-generics
[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify
[godoc-url]: http://pkg.go.dev/github.com/go-openapi/testify
[slack-logo]: https://a.slack-edge.com/e6a93c1/img/icons/favicon-32.png
[slack-badge]: https://img.shields.io/badge/slack-blue?link=https%3A%2F%2Fgoswagger.slack.com%2Farchives%2FC04R30YM
[slack-url]: https://goswagger.slack.com/archives/C04R30YMU
[discord-badge]: https://img.shields.io/discord/1446918742398341256?logo=discord&label=discord&color=blue
[discord-url]: https://discord.gg/DrafRmZx

<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-openapi/testify/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/go-openapi/testify
[goversion-url]: https://github.com/go-openapi/testify/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/go-openapi/testify
[commits-badge]: https://img.shields.io/github/commits-since/go-openapi/testify/latest
[ci-release-workflow]: https://github.com/go-openapi/testify/actions/workflows/bump-release.yml
