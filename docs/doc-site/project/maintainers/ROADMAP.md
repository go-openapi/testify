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
    ✅ v2.0 (Nov 2025) : zero dependencies
                    : optional dependencies (YAML)
                    : modernized code (relint)
                    : JSONEqBytes
    section Q1 2026
    ✅ v2.1 (Jan 2026) : generated assertions
                    : complete refactoring
                    : documentation site
                    : panic handling fixes
                    : removed deprecated
    📝 v2.2 (Fev 2026) : generics
                    : Kind/NotKind
                    : SortedT, NotSortedT
                    : complete test refactoring
                    : more benchmarks. Perf improvements
                    : optional dependencies (colorized)
    ⏳ v2.3 (Mar 2026) : other extensions
                    : JSON assertions. JSONMarshalsAs...
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
8. [x] Realign behavior re quirks, bugs, unexpected logics ... (e.g. IsNonDecreasing, EventuallyWithT...)
9. [ ] New features following test simplification effort in go-openapi repos (e.g. JSONMarshalsAs ...)
10. [ ] Unclear assertions might be provided an alternative verb (e.g. `InDelta`)
11. [ ] Inject this test dependency into the `go-swagger` tool

### What won't come anytime soon

* mocks: we use [mockery](https://https://github.com/vektra/mockery) and prefer the simpler `matryer` mocking-style.
  testify-style mocks are thus not going to be supported anytime soon.
* extra convoluted stuff in the like of `InDeltaSlice` (more likely to be removed)

## Upstream Tracking

We actively monitor [github.com/stretchr/testify](https://github.com/stretchr/testify) for updates, new issues, and proposals.

**Review frequency**: Quarterly (next review: April 2026)

**Processed items**: 28 upstream PRs and issues have been reviewed, with 21 implemented/merged, 3 superseded by our implementation, and 2 currently under consideration.

For a complete catalog of all upstream PRs and issues we've processed (implemented, adapted, superseded, or monitoring), see the [Upstream Tracking](../../usage/TRACKING.md).

