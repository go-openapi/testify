---
title: "Migration Guide"
description: "Migrating from testify/v1"
weight: 20
---

## Migration Guide from stretchr/testify v1

### 1. Update Import Path

```go
// Old
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"

// New
import "github.com/go-openapi/testify/v2/assert"
import "github.com/go-openapi/testify/v2/require"
```

### 2. Optional: Enable YAML Support

If you use `YAMLEq` assertions:

```go
import _ "github.com/go-openapi/testify/enable/yaml/v2"
```

Without this import, YAML assertions will panic with a helpful error message.

### 3. Optional: Enable Colorized Output

```go
import _ "github.com/go-openapi/testify/enable/colors/v2"
```

Use go additional test flags or environment variables: `TESTIFY_COLORIZED=true`, `TESTIFY_THEME=dark|light`

Example:

```
go test -v -testify.colorized -testify.theme=light .
```

![Colorized Test](colorized.png)

### 4. Optional: Adopt Generic Assertions

For better type safety and performance, consider migrating to generic assertion variants. This is entirely optional—reflection-based assertions continue to work as before.

#### Step 1: Identify Generic-Capable Assertions

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

#### Step 2: Add Type Suffix

Simply add `T` to the function name. The compiler will check types automatically:

```go
// Before
assert.Equal(t, expected, actual)
assert.ElementsMatch(t, slice1, slice2)

// After
assert.EqualT(t, expected, actual)
assert.ElementsMatchT(t, slice1, slice2)
```

#### Step 3: Fix Type Mismatches

The compiler will now catch type errors. This is a feature—it reveals bugs:

```go
// Compiler catches this
assert.EqualT(t, int64(42), int32(42))
// ❌ Error: mismatched types int64 and int32

// Fix: Use same type
assert.EqualT(t, int64(42), int64(actual))

// Or: Use reflection if cross-type comparison is intentional
assert.Equal(t, int64(42), int32(42))  // Still works
```

#### Benefits of Migration

- **Compile-time type safety**: Catch errors when writing tests
- **Performance**: 1.2x to 81x faster (see [Benchmarks](../project/maintainers/BENCHMARKS.md))
- **IDE support**: Better autocomplete with type constraints
- **Refactoring safety**: Type changes break tests at compile time, not runtime

See the [Generics Guide](./GENERICS.md) for detailed usage patterns and best practices.

### 5. Remove Suite/Mock Usage

Replace testify mocks with: 
- [mockery](https://github.com/vektra/mockery) for mocking
Replace testify suites with:
- Standard Go subtests for test organization
- or wait until we reintroduce this feature (possible, but not certain)

### 6. Remove HTTP Assertion Usage

If you were still using the deprecated package `github.com/stretchr/testitfy/http`,
you'll need to replace with standard HTTP testing. We won't reintroduce this package ever.

## Breaking Changes Summary

### Removed Packages

- ❌ `suite` - Use standard Go subtests
- ❌ `mock` - Use [mockery](https://github.com/vektra/mockery)
- ❌ `http` - May be reintroduced later

### Removed Functions

- ❌ All deprecated functions from v1 removed

### Behavior Changes

Make sure to check the [behavior changes](./CHANGES.md) as we have fixed a few quirks in the existing API
(mostly edge cases handling).

---

## See Also

- [Changes from v1](./CHANGES.md) - Complete list of all changes, fixes, and new features
- [Examples](./EXAMPLES.md) - Practical examples showing v2 usage patterns
- [Generics Guide](./GENERICS.md) - Learn about the 38 new type-safe generic assertions
- [Usage Guide](./USAGE.md) - API conventions and how to navigate the documentation
- [Tutorial](./TUTORIAL.md) - Best practices for writing tests with testify v2

