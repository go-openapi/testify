---
title: README
description: |
  Introducing `go-openapi/testify/v2`.

  - Approach
  - Main features
  - Differences with v1
weight: 1
---

## Motivation

From the maintainers of `testify`, it looks like a v2 is coming up, but they'll do it at their own pace.

We like all the principles they put forward to build this v2. [See discussion about v2](https://github.com/stretchr/testify/discussions/1560)

However, at `go-openapi` we would like to address the well-known issues in `testify` with different priorities.

---

1. We want first to remove all external dependencies.

> For all our libraries and generated test code we don't want test dependencies
> to drill farther than `import github.com/go-openapi/testify/v2`, but on some specific (and controlled)
> occasions.
>
> In this fork, all external stuff is either internalized (`go-spew`, `difflib`),
> removed (`mocks`, `suite`, `http`) or specifically enabled by importing a specific module
> (`github.com/go-openapi/testify/v2/enable/yaml`).

2. Make it easy to maintain and extend.

> For go-openapi, testify should just be yet another part of our toolkit.
> We need it to work, be easily adaptable to our needs and not divert our development effort away from our other repos.
> This big refactor is an investment that has to pay off.

3. We want to pare down some of the chrome that has been added over the years

> The `go-openapi` libraries and the `go-swagger` project make a rather limited use of the vast API provided by `testify`.
>
> With this first version of the fork, we have removed `mocks` and `suite`, which we don't use.
> They might be added later on, with better controlled dependencies.
>
> In the forthcoming maintenance of this fork, much of the "chrome" or "ambiguous" API will be pared down.
> There is no commitment yet on the stability of the API.
>
> Chrome would be added later: we have the "enable" packages just for that.

4. We hope that this endeavor will help the original project with a live-drill of what a v2 could look like.
   We are always happy to discuss with people who face the same problems as we do: avoid breaking changes, 
   APIs that became bloated over a decade or so, uncontrolled dependencies, conflicting demands from users etc.

### The approach with this fork

This fork targets **go1.24**.

* [x] **zero** external dependencies by default
* [x] extra features (and dependencies) are opt-in
* [x] **modernized** code base
* [x] **simplified** maintenance
* [x] can add or remove assertions with ease
* [x] **mitigated API bloat** with comprehensive domain-indexed documentation
* [ ] can leverage generics without backward compatibility concerns

The approach will be **selective and pragmatic** rather than comprehensive:

* **Targeted improvements** where generics provide clear value without compromising existing functionality
* **Focus on eliminating anti-patterns** like dummy value instantiation in `IsType` (see #1805)
* **Preserve reflection-based flexibility** for comparing complex types rather than forcing everything through generic constraints
* **Careful constraint design** to ensure type safety without being overly restrictive or permissive

The goal is to enhance type safety and developer experience where it matters most,
while maintaining the flexibility that makes testify useful for real-world testing scenarios.

## Breaking changes from v1

* `YAMLEq` panics by default: must enable the feature with an additional blank import
* deprecated types and methods have been removed
* removed the `suite` and `http` packages
* replaced internal utility package `_codegen` by `codegen`

### Other (non-breaking) changes

* added `JSONEqBytes`
* adapted & merged fixes: see [ROADMAP](./maintainers/ROADMAP.md).

## Usage at go-openapi

At this moment, we have identified the following usage in our tools. This API shall remain stable.
Currently, there are no guarantees about the entry points not in this list.

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


## [The original README](./maintainers/ORIGINAL.md)

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
