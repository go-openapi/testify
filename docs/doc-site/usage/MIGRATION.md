---
title: "Migration Guide"
description: "Migrating from testify/v1"
weight: 20
---

## Migration Guide from stretchr/testify v1

This guide covers migrating from `stretchr/testify` to `go-openapi/testify/v2`.
You can use the [automated migration tool](#automated-migration-tool) or migrate [manually](#manual-migration).

### Automated Migration Tool

`migrate-testify` automates both the import migration (pass 1) and the generic
upgrade (pass 2). It uses `go/packages` and `go/types` for type-checked,
semantics-preserving transformations.

#### Installation

```bash
go install github.com/go-openapi/testify/hack/migrate-testify/v2@latest
```

This installs the `migrate-testify` binary into your `$GOBIN`.

#### Quick Start

```bash
# Run both passes on the current directory (preview first, then apply)
migrate-testify --all --dry-run .
migrate-testify --all .

# Or run each pass separately
migrate-testify --migrate .
migrate-testify --upgrade-generics .
```

#### Pass 1: Import Migration (`--migrate`)

Rewrites `stretchr/testify` imports to `go-openapi/testify/v2`:

```bash
# Dry-run to preview changes
migrate-testify --migrate --dry-run .

# Apply changes
migrate-testify --migrate .
```

This pass handles:
- Import path rewriting (`assert`, `require`, root package)
- Function renames (`EventuallyWithT` to `EventuallyWith`, `NoDirExists` to `DirNotExists`, etc.)
- Type replacement (`PanicTestFunc` to `func()`)
- YAML enable import injection (adds `_ "github.com/go-openapi/testify/v2/enable/yaml"` when `YAMLEq` is used)
- Incompatible import detection (`mock`, `suite`, `http` packages emit warnings with guidance)
- `go.mod` update (drops `stretchr/testify`, adds `go-openapi/testify/v2`)

#### Pass 2: Generic Upgrade (`--upgrade-generics`)

Upgrades reflection-based assertions to generic variants where types are statically
resolvable and the semantics are preserved:

```bash
# Dry-run to preview changes
migrate-testify --upgrade-generics --dry-run .

# Apply changes
migrate-testify --upgrade-generics .
```

The tool is conservative: it only upgrades when:
- Argument types are statically known (no `any`, no `interface{}`)
- Types satisfy the required constraint (`comparable`, `Ordered`, `Text`, etc.)
- For `Equal`/`NotEqual`: types are "deeply comparable" (no pointers or structs with pointer fields)
- For `Contains`: the container type disambiguates to `StringContainsT`, `SliceContainsT`, or `MapContainsT`
- `IsType` is flagged for manual review (argument count changes)

Assertions that cannot be safely upgraded are tracked and reported in the summary with
a specific reason (e.g., "pointer type", "interface{}/any", "type mismatch").
Use `--verbose` to see the file and line of each skipped assertion.

#### Reference

```
Usage: migrate-testify [flags] [directory]

Migrate stretchr/testify to go-openapi/testify/v2 and upgrade to generic assertions.

Flags:
  -all                Run both passes sequentially
  -dry-run            Show diffs without modifying files
  -migrate            Run pass 1: stretchr/testify -> go-openapi/testify/v2
  -upgrade-generics   Run pass 2: reflection -> generic assertions
  -verbose            Print detailed transformation info
  -skip-gomod         Skip go.mod changes
  -skip-vendor        Skip vendor/ directory (default true)
  -version string     Target testify version (default "v2.3.0")

At least one of --migrate, --upgrade-generics, or --all is required.

Mono-repo support:
  Pass 1 walks the filesystem and works across module boundaries.
  Pass 2 requires type information and uses go/packages to load code.
  For multi-module repos, a go.work file must be present so that pass 2
  can load all workspace modules. Create one with:
    go work init . ./sub/module1 ./sub/module2 ...

Post-migration checklist:
  - Run your linter: the migration may surface pre-existing unchecked linting issues.
  - Run your test suite to verify all tests still pass.
```

---

### Manual Migration

#### 1. Update Import Paths

```go
// Old
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"

// New
import "github.com/go-openapi/testify/v2/assert"
import "github.com/go-openapi/testify/v2/require"
```

#### 2. Optional: Enable YAML Support

If you use `YAMLEq` assertions: this feature is now opt-in.

```go
import _ "github.com/go-openapi/testify/enable/yaml/v2"
```

Without this import, YAML assertions will panic with a helpful error message.

#### 3. Optional: Enable Colorized Output

```go
import _ "github.com/go-openapi/testify/enable/colors/v2"
```

Use go additional test flags or environment variables: `TESTIFY_COLORIZED=true`, `TESTIFY_THEME=dark|light`

Example:

```
go test -v -testify.colorized -testify.theme=light .
```

![Colorized Test](colorized.png)

#### 4. Optional: Adopt Generic Assertions

For better type safety and performance, consider migrating to generic assertion variants.
This is entirely optional: reflection-based assertions continue to work as before.

##### Identify Generic-Capable Assertions

Look for these common assertions in your tests:

```go
// Equality
assert.Equal → assert.EqualT
assert.NotEqual → assert.NotEqualT

// Comparisons
assert.Greater → assert.GreaterT
assert.Less → assert.LessT
assert.Positive → assert.PositiveT
assert.Negative → assert.NegativeT

// Collections
assert.Contains → assert.ContainsT (or SliceContainsT/MapContainsT/StringContainsT)
assert.ElementsMatch → assert.ElementsMatchT
assert.Subset → assert.SubsetT

// Ordering
assert.IsIncreasing → assert.IsIncreasingT
assert.IsDecreasing → assert.IsDecreasingT

// Type checks
assert.IsType(t, User{}, v) → assert.IsOfTypeT[User](t, v)  // No dummy value!
```

Simply add `T` to the function name. The compiler will check types automatically:

```go
// Before
assert.Equal(t, expected, actual)
assert.ElementsMatch(t, slice1, slice2)

// After
assert.EqualT(t, expected, actual)
assert.ElementsMatchT(t, slice1, slice2)
```

##### Fix Type Mismatches

The compiler will now catch type errors. This is a feature—it reveals bugs:

```go
// Compiler catches this
assert.EqualT(t, int64(42), int32(42))
// Error: mismatched types int64 and int32

// Fix: Use same type
assert.EqualT(t, int64(42), int64(actual))

// Or: Use reflection if cross-type comparison is intentional
assert.Equal(t, int64(42), int32(42))  // Still works
```

##### Pointer Semantics: When NOT to Upgrade

Generic assertions use Go's `==` operator, while reflection-based assertions use `reflect.DeepEqual`.
For most types these are equivalent, but **they differ for pointers and structs containing pointers**:

```go
a := &MyStruct{Name: "alice"}
b := &MyStruct{Name: "alice"}

assert.Equal(t, a, b)   // PASSES (reflect.DeepEqual compares pointed-to values)
assert.EqualT(t, a, b)  // FAILS  (== compares pointer addresses)
```

**Do not upgrade to generic variants when:**
- Arguments are pointer types (`*T`) — `EqualT` compares addresses, not values
- Arguments are structs with pointer fields — `==` compares field addresses, `DeepEqual` compares field values
- You intentionally rely on cross-type comparison (`int64` vs `int32`)

The automated migration tool handles this automatically by only upgrading
assertions where the argument types are "deeply comparable" — types where `==` and
`reflect.DeepEqual` produce the same result.

##### Benefits of Generic Assertions

- **Compile-time type safety**: Catch errors when writing tests
- **Performance**: 1.2x to 81x faster (see [Benchmarks](../project/maintainers/BENCHMARKS.md))
- **IDE support**: Better autocomplete with type constraints
- **Refactoring safety**: Type changes break tests at compile time, not runtime

See the [Generics Guide](./GENERICS.md) for detailed usage patterns and best practices.

#### 5. Remove Suite/Mock Usage

Replace testify mocks with:
- [mockery](https://github.com/vektra/mockery) for mocking
Replace testify suites with:
- Standard Go subtests for test organization
- or wait until we reintroduce this feature (possible, but not certain)

#### 6. Replace `go.uber.org/goleak` with `NoGoRoutineLeak`

If you use `go.uber.org/goleak` to detect goroutine leaks in tests, consider replacing it
with `assert.NoGoRoutineLeak` (or `require.NoGoRoutineLeak`), which is built into testify v2.

```go
// Before (with goleak)
import "go.uber.org/goleak"

func TestNoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	// ... test code ...
}

// After (with testify v2)
import "github.com/go-openapi/testify/v2/assert"

func TestNoLeak(t *testing.T) {
	assert.NoGoRoutineLeak(t, func() {
		// ... test code ...
	})
}
```

This removes the `go.uber.org/goleak` dependency. This step is not automated by the
migration tool.

#### 7. Remove use of the `testify/http` package

If you were still using the deprecated package `github.com/stretchr/testitfy/http`,
you'll need to replace it by the standard `net/http/httptest` package.

We won't reintroduce this package ever.

---

## Breaking Changes Summary

### Removed Packages

- `suite` - Use standard Go subtests
- `mock` - Use [mockery](https://github.com/vektra/mockery)
- `http` - May be reintroduced later

### Removed Functions and Types

- All deprecated functions from v1 removed
- Removed extraneous "helper" types: `PanicTestFunc` (`func()`)

### Behavior Changes

Make sure to check the [behavior changes](./CHANGES.md) as we have fixed a few quirks in the existing API
(mostly edge cases handling).

---

## See Also

- [Changes from v1](./CHANGES.md) - Complete list of all changes, fixes, and new features
- [Examples](./EXAMPLES.md) - Practical examples showing v2 usage patterns
- [Generics Guide](./GENERICS.md) - Learn about the 43 new type-safe generic assertions
- [Usage Guide](./USAGE.md) - API conventions and how to navigate the documentation
- [Tutorial](./TUTORIAL.md) - Best practices for writing tests with testify v2
