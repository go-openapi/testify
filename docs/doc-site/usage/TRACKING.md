---
title: "Upstream Tracking"
description: "All issues and PRs reviewed for this fork"
weight: 16
---

## Upstream Tracking

We continue to monitor and selectively adopt changes from upstream:

### Implemented from Upstream
- ✅ [#1513] - JSONEqBytes
- ✅ [#1803] - Kind/NotKind assertions
- ✅ [#1805] - IsOfTypeT[T] generic assertions
- ✅ [#1685] - Partial iterator support (SeqContainsT variants)
- ✅ [#1828] - Spew panic fixes
- ✅ [#1825], [#1818], [#1223], [#1813], [#1611], [#1822], [#1829] - Various bug fixes

### Monitoring
- 🔍 [#1087] - Consistently assertion
- 🔍 [#1601] - NoFieldIsZero

### Superseded by Our Implementation
- ⛔ [#1830] - CollectT.Halt() (superseded by context-based pollCondition)
- ⛔ [#1819] - Handle unexpected exits (superseded by context-based pollCondition)
- ⛔ [#1824] - Spew testing (superseded by property-based fuzzing)

[#1087]: https://github.com/stretchr/testify/pull/1087
[#1601]: https://github.com/stretchr/testify/issues/1601
[#1830]: https://github.com/stretchr/testify/pull/1830
[#1824]: https://github.com/stretchr/testify/pull/1824
[#1819]: https://github.com/stretchr/testify/pull/1819

**Review frequency**: Quarterly (next review: April 2026)

---

## Appendix: Upstream References

This table catalogs all upstream PRs and issues from [github.com/stretchr/testify](https://github.com/stretchr/testify) that we have processed.

### Implemented (Adapted or Merged)

| Reference | Type | Summary | Outcome in Fork |
|-----------|------|---------|-----------------|
| [#994] | PR | Colorize expected vs actual values | ✅ Adapted into `enable/color` module with themes and configuration |
| [#1223] | PR | Display uint values in decimal instead of hex | ✅ Merged - Applied to diff output |
| [#1232] | PR | Colorized output for expected/actual/errors | ✅ Adapted into `enable/color` module |
| [#1356] | PR | panic(nil) handling for Go 1.21+ | ✅ Merged - Updated panic assertions |
| [#1467] | PR | Colorized output with terminal detection | ✅ Adapted into `enable/color` module (most mature implementation) |
| [#1480] | PR | Colorized diffs via TESTIFY_COLORED_DIFF env var | ✅ Adapted with env var support in `enable/color` |
| [#1513] | PR | JSONEqBytes for byte slice JSON comparison | ✅ Merged - Added to JSON domain |
| [#1685] | PR | Iterator support (`iter.Seq`) for Contains/ElementsMatch | ✅ Partial - Implemented SeqContainsT and SeqNotContainsT only |
| [#1772] | PR | YAML library migration to maintained fork | ✅ Adapted - Used gopkg.in/yaml.v3 in optional `enable/yaml` module |
| [#1797] | PR | Codegen package consolidation and licensing | ✅ Adapted - Complete rewrite of code generation system |
| [#1803] | PR | Kind/NotKind assertions | ✅ Merged - Added to Type domain |
| [#1805] | Issue | Generic `IsOfType[T]()` without dummy value | ✅ Implemented - IsOfTypeT and IsNotOfTypeT in Type domain |
| [#1816] | Issue | Fix panic on unexported struct key in map | ✅ Fixed in internalized go-spew |
| [#1818] | PR | Fix panic on invalid regex in Regexp/NotRegexp | ✅ Merged - Added graceful error handling |
| [#1822] | Issue | Deterministic map ordering in diffs | ✅ Fixed in internalized go-spew |
| [#1825] | PR | Fix panic using EqualValues with uncomparable types | ✅ Merged - Enhanced type safety in EqualValues |
| [#1826] | Issue | Type safety with spew (meta-issue) | ✅ Addressed through comprehensive fuzzing and fixes |
| [#1828] | PR | Fixed panic with unexported fields in maps | ✅ Merged into internalized go-spew |
| [#1829] | Issue | Fix time.Time rendering in diffs | ✅ Fixed in internalized go-spew |
| [#1611] | Issue | Goroutine leak in Eventually/Never | ✅ Fixed by using context.Context (consolidation into single pollCondition function) |
| [#1813] | Issue | Panic with unexported fields | ✅ Fixed via #1828 in internalized spew |

[#994]: https://github.com/stretchr/testify/pull/994
[#1356]: https://github.com/stretchr/testify/pull/1356
[#1772]: https://github.com/stretchr/testify/pull/1772
[#1797]: https://github.com/stretchr/testify/pull/1797
[#1816]: https://github.com/stretchr/testify/issues/1816
[#1826]: https://github.com/stretchr/testify/issues/1826

### Superseded by Our Implementation

| Reference | Type | Summary | Why Superseded |
|-----------|------|---------|----------------|
| [#1819] | PR | Handle unexpected exits in Eventually | Superseded by context-based pollCondition implementation |
| [#1824] | PR | Spew testing improvements | Superseded by property-based fuzzing with random type generator |
| [#1830] | PR | CollectT.Halt() for stopping tests | Superseded by context-based pollCondition implementation |

[#1819]: https://github.com/stretchr/testify/pull/1819

### Under Consideration (Monitoring)

| Reference | Type | Summary | Status |
|-----------|------|---------|--------|
| [#1087] | PR | Consistently assertion | 🔍 Monitoring - Evaluating usefulness |
| [#1601] | Issue | NoFieldIsZero assertion | 🔍 Monitoring - Considering implementation |

### Informational (Not Implemented)

| Reference | Type | Summary | Outcome |
|-----------|------|---------|---------|
| [#1147] | Issue | General discussion about generics adoption | ℹ️ Marked "Not Planned" upstream - We implemented our own generics approach (38 functions) |
| [#1308] | PR | Comprehensive refactor with generic type parameters | ℹ️ Draft for v2.0.0 upstream - We took a different approach with the same objective |

[#1147]: https://github.com/stretchr/testify/issues/1147
[#1308]: https://github.com/stretchr/testify/pull/1308

### Summary Statistics

| Category | Count |
|----------|-------|
| **Implemented/Merged** | 21 |
| **Superseded** | 3 |
| **Monitoring** | 2 |
| **Informational** | 2 |
| **Total Processed** | 28 |

**Note**: This fork maintains an active relationship with upstream, regularly reviewing new PRs and issues. The quarterly review process ensures we stay informed about upstream developments while maintaining our architectural independence.

---
