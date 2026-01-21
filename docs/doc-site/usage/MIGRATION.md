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
import _ "github.com/go-openapi/testify/v2/enable/yaml"
```

Without this import, YAML assertions will panic with a helpful error message.

### 3. Optional: Enable Colorized Output

```go
import _ "github.com/go-openapi/testify/v2/enable/color"
```

Use go additional test flags or environment variables: `TESTIFY_COLORIZED=true`, `TESTIFY_THEME=dark|light`

### 4. Optional: Adopt Generic Assertions

For better type safety and performance:

```go
// Reflection-based (still works)
assert.Equal(t, expected, actual)
assert.ElementsMatch(t, slice1, slice2)

// Generic (type-safe + faster)
assert.EqualT(t, expected, actual)
assert.ElementsMatchT(t, slice1, slice2)
```

See [Generics Guide](../GENERICS.md) for complete migration guide.

### 5. Remove Suite/Mock Usage

Replace testify suites and mocks with:
- [mockery](https://github.com/vektra/mockery) for mocking
- Standard Go subtests for test organization

### 6. Remove HTTP Assertion Usage

If you used `assert.HTTPSuccess`, `assert.HTTPStatusCode`, etc., you'll need to replace with standard HTTP testing or wait for potential reintroduction.

## Breaking Changes Summary

### Removed Packages

- ❌ `suite` - Use standard Go subtests
- ❌ `mock` - Use [mockery](https://github.com/vektra/mockery)
- ❌ `http` - May be reintroduced later

### Removed Functions

- ❌ All deprecated functions from v1 removed

### Behavior Changes

- `EqualValues` now fails with function types (consistency with `Equal`)
- `IsNonDecreasing`/`IsNonIncreasing` logic corrected (previously inverted)
- Nil pointer identity handling corrected (`Same`/`NotSame`)

### Dependency Changes

- ✅ Zero external dependencies
- ✅ YAML support requires opt-in via `enable/yaml` module

