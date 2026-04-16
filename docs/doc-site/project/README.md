---
title: README
description: |
  Introducing go-openapi/testify/v2.

  - Why choose testify/v2
  - Approach
  - Main features
  - Differences with v1
weight: 2
---

**The v2 our tests wanted**

## Why choose `go-openapi/testify/v2`

* 95% compatible with `stretchr/testify` — if you already use it, our migration tool automates the switch
* Actively maintained: regular fixes and evolutions, many PRs proposed upstream are already in
* Zero external dependencies — you import what you need, with opt-in modules for extras (e.g. YAML, colorized output)
* Modernized codebase targeting go1.25+
* Go routine leak detection built in: zero-setup, no false positives, works with parallel tests (unlike `go.uber.org/goleak`)
* File descriptor leak detection (linux-only)
* Type-safe assertions with generics (see [a basic example](../usage/EXAMPLES.md)) — migration to generics can be automated too. [Read the full story](../usage/GENERICS.md)
* Safe async assertions, extended JSON & YAML assertions
* Coming in `v2.5.0`: non-flaky async assertions using `synctest`, and internal tools exposed as standalone modules (spew, unified diff, goleak)
* We take documentation seriously: [searchable doc site](../../api/) with testable examples and a complete tutorial, plus detailed [godoc][godoc-url] for every assertion

### This fork isn't for everyone

* You need the `mock` package — we removed it and won't bring it back. For suites, we're [open to discussion](https://github.com/go-openapi/testify/discussions/75) about a redesigned approach
* Your project must support Go versions older than 1.25
* You rely on `testifylint` or other tooling that expects the `stretchr/testify` import path
* You need 100% API compatibility — we're at 95%, and the remaining 5% are intentional removals

## Motivation

See [why we wanted a v2](./MOTIVATION.md).

### Approach with this fork

This fork targets **go1.25**.

* [x] **zero** external dependencies by default
* [x] extra features (and dependencies) are opt-in
* [x] **modernized** code base
* [x] **simplified** maintenance
* [x] can add or remove assertions with ease
* [x] **mitigated API bloat** with comprehensive domain-indexed documentation
* [x] can leverage generics without backward compatibility concerns

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
* removed the `suite`, `mocks` and `http` packages
* replaced internal utility package `_codegen` by `codegen`

See [all changes from v1](../usage/CHANGES.md) and check out our [ROADMAP](./maintainers/ROADMAP.md).

## API Stability Guarantee

The assertions currently used by go-openapi projects constitute our **stable API**.
These entry points will remain backward compatible. Other assertions may evolve as we refine the v2 API.

---

## See Also

**Getting Started:**
- [Examples](../usage/EXAMPLES.md) - Practical code examples for using testify v2
- [Usage Guide](../usage/USAGE.md) - API conventions and navigation guide
- [Migration Guide](../usage/MIGRATION.md) - Migrating from stretchr/testify v1

**Project Documentation:**
- [Changes from v1](../usage/CHANGES.md) - Complete list of changes and new features
- [Roadmap](./maintainers/ROADMAP.md) - Future development plans
- [Architecture](./maintainers/ARCHITECTURE.md) - Technical architecture and design decisions
- [Contributing](./contributing/CONTRIBUTING.md) - How to contribute to this project
- [The original README](./maintainers/ORIGINAL.md)

---

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
