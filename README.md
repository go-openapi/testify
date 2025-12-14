# Testify

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/testify/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/testify.svg)](https://pkg.go.dev/github.com/go-openapi/testify)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/testify)](https://goreportcard.com/report/github.com/go-openapi/testify)

<!-- Badges: status  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vuln scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
<!-- Badges: release & docker images  -->
<!-- Badges: code quality  -->
<!-- Badges: license & compliance -->
[![Release][release-badge]][release-url] [![Go Report Card][gocard-badge]][gocard-url] [![CodeFactor Grade][codefactor-badge]][codefactor-url] [![License][license-badge]][license-url]
<!-- Badges: documentation & support -->
<!-- Badges: others & stats -->
<!-- Slack badge disabled until I am able to restore a valid link to the chat -->
[![GoDoc][godoc-badge]][godoc-url] <!-- [![Slack Channel][slack-badge]][slack-url] -->[![go version][goversion-badge]][goversion-url] ![Top language][top-badge] ![Commits since latest release][commits-badge]

---

## Testify - Thou Shalt Write Tests

A Go set of packages that provide tools for testifying that your code will behave as you intend.

This is the go-openapi fork of the great [testify](https://github.com/stretchr/testify) package.

## Why this fork?

From the maintainers of `testify`, it looks like a v2 is coming up, but they'll do it at their own pace.

We like all the principles they put forward to build this v2. [See discussion about v2](https://github.com/stretchr/testify/discussions/1560)

However, at `go-openapi` we would like to address the well-known issues in `testify` with different priorities.

1. We want first to remove all external dependencies.

> For all our libraries and generated test code we don't want test dependencies
> to drill farther than `import github.com/go-openapi/testify/v2`, but on some specific (and controlled)
> occasions.
>
> In this fork, all external stuff is either internalized (`go-spew`, `difflib`),
> removed (`mocks`, `suite`, `http`) or specifically enabled by importing a specific module
> (`github.com/go-openapi/testify/v2/enable/yaml`).

2. We want to remove most of the chrome that has been added over the years

> The `go-openapi` libraries and the `go-swagger` project make a rather limited use of the vast API provided by `testify`.
>
> With this first version of the fork, we have removed `mocks` and `suite`, which we don't use.
> They might be added later on, with better controlled dependencies.
>
> In the forthcoming maintenance of this fork, much of the "chrome" or "ambiguous" API will be pared down.
> There is no commitment yet on the stability of the API.
>
> Chrome would be added later: we have the "enable" packages just for that.

3. We hope that this endeavor will help the original project with a live-drill of what a v2 could look like.
   We are always happy to discuss with people who face the same problems as we do: avoid breaking changes, 
   APIs that became bloated over a decade or so, uncontrolled dependencies, conflicting demands from users etc.

## What's next with this project?

1. [x] The first release comes with zero dependencies and an unstable API (see below [our use case](#usage-at-go-openapi))
2. |x] This project is going to be injected as the main and sole test dependency of the `go-openapi` libraries
2. [ ] ... and the `go-swagger` tool
3. [ ) Valuable pending pull requests from the original project could be merged (e.g. `JSONEqBytes`) or transformed as "enable" modules (e.g. colorized output)
4. [ ] Unclear assertions may be provided an alternative verb (e.g. `InDelta`)
5. [ ] Since we have leveled the go requirements to the rest of the go-openapi (currently go1.24) there is quite a bit of relinting lying ahead.

### What won't come anytime soon

* mocks: we use [mockery](https://github.com/vektra/mockery) and prefer the simpler `matryer` mocking-style.
  testify-style mocks are thus not going to be supported anytime soon.
* extra convoluted stuff in the like of `InDeltaSlice`

## Generics adoption

### Context from the original repository

Several attempts have been made to introduce generics in the original stretchr/testify repository:

* **github.com/stretchr/testify#1308** - Comprehensive refactor replacing `interface{}` with generic type parameters across assertions (Draft, v2.0.0 milestone)
* **github.com/stretchr/testify#1805** - Proposal for generic `IsOfType[T]()` to avoid dummy value instantiation in type checks
* **github.com/stretchr/testify#1685** - Iterator support (`iter.Seq`) for Contains/ElementsMatch assertions (Go 1.23+)
* **github.com/stretchr/testify#1147** - General discussion about generics adoption (marked "Not Planned")

### Challenges identified

The original repository's exploration of generics revealed several design challenges:

1. **Type inference limitations**: Go's type inference struggles with complex generic signatures, often requiring explicit type parameters that burden the API (e.g., `Contains[int, int](arr1, arr2)`)

2. **Overly broad type constraints**: PR #1308's approach used constraints like `ConvertibleToFloat64` that accepted more types than intended, weakening type safety

3. **Loss of flexibility**: Testify currently compares non-comparable types (slices, maps) via `reflect.DeepEqual`. Generic constraints would eliminate this capability, as Go generics require comparable or explicitly constrained types

4. **Breaking changes**: Any comprehensive generics adoption requires a major version bump and Go 1.18+ minimum version

5. **Inconsistent design patterns**: Different assertions would need different constraint strategies, making a uniform approach difficult

### Approach in this fork

This fork targets **go1.24** and can leverage generics without backward compatibility concerns.

The approach will be **selective and pragmatic** rather than comprehensive:

* **Targeted improvements** where generics provide clear value without compromising existing functionality
* **Focus on eliminating anti-patterns** like dummy value instantiation in `IsType` (see #1805)
* **Preserve reflection-based flexibility** for comparing complex types rather than forcing everything through generic constraints
* **Careful constraint design** to ensure type safety without being overly restrictive or permissive

The goal is to enhance type safety and developer experience where it matters most, while maintaining the flexibility that makes testify useful for real-world testing scenarios.

**Status**: Design and exploration phase. Contributions and proposals welcome.

## Usage at go-openapi

At this moment, we have identified the following usage in our tools. This API shall remain stable.
Currently, there are no guarantees about the entry points not in this list.

TODO: extend the list with usage by go-swagger.

```
Condition
Contains,Containsf
Empty,Emptyf
Equal,Equalf
EqualError,EqualErrorf
EqualValues,EqualValuesf
Error,Errorf
ErrorContains
ErrorIs
Fail,Failf
FailNow
False,Falsef
Greater
Implements
InDelta,InDeltaf
IsType,IsTypef
JSONEq,JSONEqf
Len,Lenf
Nil,Nilf
NoError,NoErrorf
NotContains,NotContainsf
NotEmpty,NotEmptyf
NotEqual
NotNil,NotNilf
NotPanics
NotZeroG
Panics,PanicsWithValue
Subset
True,Truef
YAMLEq,YAMLEqf
Zero,Zerof
```

## Installation

To use this package in your projects:

```cmd
    go get github.com/go-openapi/testify/v2
```

## Get started

Features include:

  * [Easy assertions](./docs/ORIGINAL.md#assert-package)
  * ~[Mocking](./docs/ORIGINAL.md#mock-package)~ removed
  * ~[Testing suite interfaces and functions](./docs/ORIGINAL.md#suite-package)~ removed

## Examples

See [the original README](./docs/ORIGINAL.md)

## Licensing

`SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers`

This library ships under the [SPDX-License-Identifier: Apache-2.0](./LICENSE).

See the license [NOTICE](./NOTICE), which recalls the licensing terms of all the pieces of software
distributed with this fork, including internalized libraries.

* stretchr/testify [SPDX-License-Identifier: MIT](./NOTICE)
* github.com/davecgh/go-spew [SPDX-License-Identifier: ISC](./internal/spew/LICENSE)
* github.com/pmezard/go-difflib [SPDX-License-Identifier: MIT-like](./internal/difflib/LICENSE)

## PRs from the original repo

### Already merged or incorporated

The following proposed contributions to the original repo have been merged or incorporated with
some adaptations into this fork:

* github.com/stretchr/testify#1513 - JSONEqBytes for byte slice JSON comparison
* github.com/stretchr/testify#1772 - YAML library migration to maintained fork (go.yaml.in/yaml)
* github.com/stretchr/testify#1797 - Codegen package consolidation and licensing
* github.com/stretchr/testify#1356 - panic(nil) handling for Go 1.21+

### Planned merges

#### Critical safety fixes (high priority)

* github.com/stretchr/testify#1825 - Fix panic when using EqualValues with uncomparable types
* github.com/stretchr/testify#1818 - Fix panic on invalid regex in Regexp/NotRegexp assertions

#### Leveraging internalized dependencies (go-spew, difflib)

These improvements apply to the internalized and modernized copies of dependencies in this fork:

* github.com/stretchr/testify#1829 - Fix time.Time rendering in diffs (internalized go-spew)
* github.com/stretchr/testify#1822 - Deterministic map ordering in diffs (internalized go-spew)
* github.com/stretchr/testify#1816 - Fix panic on unexported struct key in map (internalized go-spew - may need deeper fix)

#### UX improvements

* github.com/stretchr/testify#1223 - Display uint values in decimal instead of hex in diffs

### Under consideration

#### Colorized output

Several PRs propose colorized terminal output with different approaches and dependencies.
If implemented, this would be provided as an optional `enable/color` module:

* github.com/stretchr/testify#1467 - Colorized output with terminal detection (most mature implementation)
* github.com/stretchr/testify#1480 - Colorized diffs via TESTIFY_COLORED_DIFF env var
* github.com/stretchr/testify#1232 - Colorized output for expected/actual/errors
* github.com/stretchr/testify#994 - Colorize expected vs actual values

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
Extra credit for those using Testify to write the test code that demonstrates it.

Code generation is used. Run `go generate ./...` to update generated files.

See also the [CONTRIBUTING guidelines](.github/CONTRIBUTING.md).


## [The original README](./original.md)

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
[release-badge]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Ftestify.svg
[release-url]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Ftestify
<!-- Badges: code quality  -->
[gocard-badge]: https://goreportcard.com/badge/github.com/go-openapi/testify
[gocard-url]: https://goreportcard.com/report/github.com/go-openapi/testify
[codefactor-badge]: https://img.shields.io/codefactor/grade/github/go-openapi/testify
[codefactor-url]: https://www.codefactor.io/repository/github/go-openapi/testify
<!-- Badges: documentation & support -->
[doc-badge]: https://img.shields.io/badge/doc-site-blue?link=https%3A%2F%2Fgoswagger.io%2Fgo-openapi%2F
[doc-url]: https://goswagger.io/go-openapi
[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify
[godoc-url]: http://pkg.go.dev/github.com/go-openapi/testify
[slack-badge]: https://slackin.goswagger.io/badge.svg
[slack-url]: https://slackin.goswagger.io
<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-openapi/testify/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/go-openapi/testify
[goversion-url]: https://github.com/go-openapi/testify/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/go-openapi/testify
[commits-badge]: https://img.shields.io/github/commits-since/go-openapi/testify/latest
