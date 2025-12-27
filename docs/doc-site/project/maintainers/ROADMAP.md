---
title: "Testify v2 roadmap"
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
    v2.2 (Fev 2026) : extension w/ generics
                    : optional dependencies (colorized)
    v2.3 (Mar 2026) : other extensions
    section Q2 2026
    2006 : Twitter
{{< /mermaid >}}

1. [x] The first release comes with zero dependencies and an unstable API (see below [our use case](#usage-at-go-openapi))
2. [x] This project is going to be injected as the main and sole test dependency of the `go-openapi` libraries
2. [ ] ... and the `go-swagger` tool
3. [x] Valuable pending pull requests from the original project could be merged (e.g. `JSONEqBytes`) or transformed as "enable" modules (e.g. colorized output)
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

## PRs from the original repo

### Already merged or incorporated

The following proposed contributions to the original repo have been merged or incorporated with
some adaptations into this fork:

* github.com/stretchr/testify#1513 - JSONEqBytes for byte slice JSON comparison
* github.com/stretchr/testify#1772 - YAML library migration to maintained fork (go.yaml.in/yaml)
* github.com/stretchr/testify#1797 - Codegen package consolidation and licensing
* github.com/stretchr/testify#1356 - panic(nil) handling for Go 1.21+
* github.com/stretchr/testify#1825 - Fix panic when using EqualValues with uncomparable types [merged]
* github.com/stretchr/testify#1818 - Fix panic on invalid regex in Regexp/NotRegexp assertions [merged]
* github.com/stretchr/testify#1223 - Display uint values in decimal instead of hex in diffs [merged]

### Planned merges

#### Critical safety fixes (high priority)

* Follow / adapt https://github.com/stretchr/testify/pull/1824

Not PRs, but reported issues in the original repo (need to investigate):

* https://github.com/stretchr/testify/issues/1826
* https://github.com/stretchr/testify/issues/1611
* https://github.com/stretchr/testify/issues/1813

#### Leveraging internalized dependencies (go-spew, difflib)

These improvements apply to the internalized and modernized copies of dependencies in this fork:

* github.com/stretchr/testify#1829 - Fix time.Time rendering in diffs (internalized go-spew)
* github.com/stretchr/testify#1822 - Deterministic map ordering in diffs (internalized go-spew)
* github.com/stretchr/testify#1816 - Fix panic on unexported struct key in map (internalized go-spew - may need deeper fix)

#### UX improvements

* diff rendering

### Under consideration

#### Colorized output

Several PRs propose colorized terminal output with different approaches and dependencies.
If implemented, this would be provided as an optional `enable/color` module:

* github.com/stretchr/testify#1467 - Colorized output with terminal detection (most mature implementation)
* github.com/stretchr/testify#1480 - Colorized diffs via TESTIFY_COLORED_DIFF env var
* github.com/stretchr/testify#1232 - Colorized output for expected/actual/errors
* github.com/stretchr/testify#994 - Colorize expected vs actual values
