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
    ✅ v2.0 (Nov 2025) : Zero dependencies
                    : optional dependencies (YAML)
                    : modernized code (relint)
                    : JSONEqBytes
    section Q1 2026
    ✅ v2.1 (Jan 2026) : Generated assertions
                    : complete refactoring
                    : documentation site
                    : panic handling fixes
                    : removed deprecated
    📝 v2.2 (Fev 2026) : Generics
                    : Kind/NotKind
                    : SortedT, NotSortedT
                    : complete test refactoring
                    : more benchmarks. Perf improvements
                    : optional dependencies (colorized)
    ⏳ v2.3 (Fev 2026) : Other extensions
                    : Extensible Assertion type
                    : JSON & YAML assertions: JSONMarshalsAs...
                    : NoGoroutineLeak
                    : more documentation and examples
    📝 v2.4 (Mar 2026) : Stabilize API (no more removals)
                    : NoFileDescriptorLeak (unix)
                    : async: Eventually/Never to accept error and context
                    : JSONPointerT
                    : export internal tools (spew, difflib,
    section Q2 2026
    v2.5 (May 2026) : New candidate features from upstream
                    : NoFileDescriptorLeak (windows port)
                    : export internal tools (blackbox)
{{< /mermaid >}}

## Notes

1. [x] The first release comes with zero dependencies and an unstable API (see below [our use case](#usage-at-go-openapi))
2. [x] This project is going to be injected as the main and sole test dependency of the `go-openapi` libraries
3. [x] Since we have leveled the go requirements to the rest of the go-openapi (currently go1.24) there is quite a bit of relinting lying ahead.
4. [x] Valuable pending pull requests from the original project could be merged (e.g. `JSONEqBytes`) or transformed as "enable" modules (e.g. colorized output)
5. [x] More testing and bug fixes (from upstream or detected during our testing)
6. [x] Introduces colorization (opt-in)
7. [x] Introduces generics
8. [x] Realign behavior re quirks, bugs, unexpected logics ... (e.g. IsNonDecreasing, EventuallyWithT...)
10. [x] Unclear assertions might be provided an alternative verb (e.g. `EventuallyWithT`)

### Adoption timeline at go-openapi

1. [x]  Jan 2026: all go-openapi projects adopts the forked testify
2. [ ] Feb 2026: all go-openapi projects transition to generics
3. [ ] Mar 2026: go-swagger transitions to the forked testify

### What won't come anytime soon

* mocks: we use [mockery](https://https://github.com/vektra/mockery) and prefer the simpler `matryer` mocking-style.
  testify-style mocks are thus not going to be supported anytime soon.
* extra convoluted stuff in the like of `InDeltaSlice` (more likely to be removed)

## Upstream Tracking

We actively monitor [github.com/stretchr/testify](https://github.com/stretchr/testify) for updates, new issues, and proposals.

**Review frequency**: Quarterly (next review: April 2026)

**Processed items**: 28 upstream PRs and issues have been reviewed, with 21 implemented/merged, 5 superseded by our implementation or merely marked as informational, and 2 currently under consideration.

For a complete catalog of all upstream PRs and issues we've processed (implemented, adapted, superseded, or monitoring), see the [Upstream Tracking](../../usage/TRACKING.md).

