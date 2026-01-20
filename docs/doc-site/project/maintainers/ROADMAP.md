---
title: "Roadmap"
description: "Let's share our plans."
weight: 4
---

## What's next with this project?

{{< mermaid align="center" zoom="true" >}}
timeline
    title Planned releases
    section Q4 2025
    v2.0 (Nov 2025) : zero dependencies
                    : optional dependencies (YAML)
                    : modernized code (relint)
                    : JSONEqBytes
    section Q1 2026
    v2.1 (Jan 2026) : generated assertions
                    : complete refactoring
                    : documentation site
                    : panic handling fixes
                    : removed deprecated
                    : optional dependencies (colorized)
                    : upstream PRs: Kind/NotKind
    v2.2 (Fev 2026) : : generics
                    : SortedT, NotSortedT
                    : JSON assertions. JSONMarshalsAs...
                    : complete test refactoring
                    : more benchmarks. Perf improvements
    v2.3 (Mar 2026) : other extensions (TBD)
                    : more documentation and examples
                    : export internal tools (spew, difflib, benchviz)
    section Q2 2026
    v2.4 (Apr 2026) : Stabilize API
                    : export internal tools (blackbox)
{{< /mermaid >}}

1. [x] The first release comes with zero dependencies and an unstable API (see below [our use case](#usage-at-go-openapi))
2. [x] This project is going to be injected as the main and sole test dependency of the `go-openapi` libraries
3. [x] Since we have leveled the go requirements to the rest of the go-openapi (currently go1.24) there is quite a bit of relinting lying ahead.
4. [x] Valuable pending pull requests from the original project could be merged (e.g. `JSONEqBytes`) or transformed as "enable" modules (e.g. colorized output)
5. [x] More testing and bug fixes (from upstream or detected during our testing)
6. [x] Introduces colorization (opt-in)
7. [x] Introduces generics
8. [ ] New features following test simplification effort in go-openapi repos (e.g. JSONMarshalsAs ...)
9. [ ] Unclear assertions may be provided an alternative verb (e.g. `InDelta`)
10. [ ] Inject this test dependency into the `go-swagger` tool

### What won't come anytime soon

* mocks: we use [mockery](https://github.com/vektra/mockery) and prefer the simpler `matryer` mocking-style.
  testify-style mocks are thus not going to be supported anytime soon.
* extra convoluted stuff in the like of `InDeltaSlice` (more likely to be removed)

## PRs and issues from the original repo

We monitor github.com/stretchr/testify (upstream) for updates, new issues and proposals.

### Already merged or incorporated / adapted

The following proposed contributions to the original repo have been merged or incorporated with
some adaptations into this fork:

* [x] github.com/stretchr/testify#1147 - General discussion about generics adoption (marked "Not Planned")
* [x] github.com/stretchr/testify#1223 - Display uint values in decimal instead of hex in diffs [merged]
* [x] github.com/stretchr/testify#1232 - Colorized output for expected/actual/errors
* [x] github.com/stretchr/testify#1308 - Comprehensive refactor replacing `interface{}` with generic type parameters across assertions (Draft, v2.0.0 milestone)
* [x] github.com/stretchr/testify#1356 - panic(nil) handling for Go 1.21+
* [x] github.com/stretchr/testify#1467 - Colorized output with terminal detection (most mature implementation)
* [x] github.com/stretchr/testify#1480 - Colorized diffs via TESTIFY_COLORED_DIFF env var
* [x] github.com/stretchr/testify#1513 - JSONEqBytes for byte slice JSON comparison
* [x] github.com/stretchr/testify#1772 - YAML library migration to maintained fork (go.yaml.in/yaml)
* [x] github.com/stretchr/testify#1797 - Codegen package consolidation and licensing
* [x] github.com/stretchr/testify#1816 - Fix panic on unexported struct key in map (internalized go-spew - may need deeper fix)
* [x] github.com/stretchr/testify#1818 - Fix panic on invalid regex in Regexp/NotRegexp assertions [merged]
* [x] github.com/stretchr/testify#1822 - Deterministic map ordering in diffs (internalized go-spew)
* [x] github.com/stretchr/testify#1825 - Fix panic when using EqualValues with uncomparable types [merged]
* [x] github.com/stretchr/testify#1829 - Fix time.Time rendering in diffs (internalized go-spew)
* [x] github.com/stretchr/testify#994 - Colorize expected vs actual values
* [x] github.com/stretchr/testify/issues/1611 - Go routine leak
* [x] github.com/stretchr/testify/issues/1813
* [x] github.com/stretchr/testify/issues/1826 - type safety with spew
* ⛔ https://github.com/stretchr/testify/pull/1824 - No longer relevant in our context

### Merges from upstream under consideration

* [ ] **github.com/stretchr/testify#1685** - Iterator support (`iter.Seq`) for Contains/ElementsMatch assertions (Go 1.23+)
* [ ] **github.com/stretchr/testify#1805** - Proposal for generic `IsOfType[T]()` to avoid dummy value instantiation in type checks

## Generics adoption

The original repository's exploration of generics revealed several design challenges:

1. **Type inference limitations**: Go's type inference struggles with complex generic signatures, often requiring explicit type parameters that burden the API (e.g., `Contains[int, int](arr1, arr2)`)

2. **Overly broad type constraints**: PR #1308's approach used constraints like `ConvertibleToFloat64` that accepted more types than intended, weakening type safety

3. **Loss of flexibility**: Testify currently compares non-comparable types (slices, maps) via `reflect.DeepEqual`. Generic constraints would eliminate this capability, as Go generics require comparable or explicitly constrained types

4. **Breaking changes**: Any comprehensive generics adoption requires a major version bump and Go 1.18+ minimum version

5. **Inconsistent design patterns**: Different assertions would need different constraint strategies, making a uniform approach difficult
