# Testify/v2

<!-- Badges: s[roadmap]tatus  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vulnerability scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
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

## Why choose `go-openapi/testify/v2`

* 95% compatible with `stretchr/testify` — if you already use it, our migration tool automates the switch
* Actively maintained: regular fixes and evolutions, many PRs proposed upstream are already in
* Zero external dependencies — you import what you need, with opt-in modules for extras (e.g. YAML, colorized output)
* Modernized codebase targeting go1.25+
* Go routine leak detection built in: zero-setup, no false positives, works with parallel tests (unlike `go.uber.org/goleak`)
* File descriptor leak detection (linux-only)
* Type-safe assertions with generics (see [a basic example][example-with-generics-url]) — migration to generics can be automated too. [Read the full story][doc-generics]
* Safe async assertions, extended JSON & YAML assertions
* Coming in `v2.5.0`: non-flaky async assertions using `synctest`, and internal tools exposed as standalone modules (spew, unified diff, goleak)
* We take documentation seriously: [searchable doc site][doc-url] with testable examples and a complete tutorial, plus detailed [godoc][godoc-url] for every assertion

### This fork isn't for everyone

* You need the `mock` package — we removed it and won't bring it back. For suites, we're [open to discussion][suite-discussion] about a redesigned approach
* Your project must support Go versions older than 1.25
* You rely on `testifylint` or other tooling that expects the `stretchr/testify` import path
* You need 100% API compatibility — we're at 95%, and the remaining 5% are intentional removals

## Announcements

### Status

Design and exploration phase completed. The published API is now stable:
moving forward, API changes will remain backward-compatible with v2.4.0.

Feedback, contributions and proposals are welcome.

> **Recent news**
>
> ✅ Preparing v2.5.0: new features: support for synctest, NoFileDescriptorLeak for macos,
> plus a few fixes (`EventuallyWithT`, `Subset`).
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

Go-swagger has also adopted it. Now the work is to generalize the use of generics (leveraging our migration tool).

Features might be added to support our main use cases there.

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
[suite-discussion]: https://github.com/go-openapi/testify/discussions/75
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
[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2
[discord-badge]: https://img.shields.io/discord/1446918742398341256?logo=discord&label=discord&color=blue
[discord-url]: https://discord.gg/FfnFYaC3k5

<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-openapi/testify/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/go-openapi/testify
[goversion-url]: https://github.com/go-openapi/testify/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/go-openapi/testify
[commits-badge]: https://img.shields.io/github/commits-since/go-openapi/testify/latest
[ci-release-workflow]: https://github.com/go-openapi/testify/actions/workflows/bump-release.yml
