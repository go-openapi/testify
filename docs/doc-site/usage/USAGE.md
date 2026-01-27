---
title: Usage
description: "Introduction Guide"
weight: 1
---

{{% notice primary "TL;DR" "meteor" %}}
> Learn testify's naming conventions (assert vs require, format variants, generic `T` suffix), argument order patterns, and how to navigate
> 76+ assertions organized into 18 domains. Start here to understand the API structure.
{{% /notice %}}

Testify v2 provides **over 40 core assertion types** (76+ functions including inverse variants and all naming styles) organized into clear domains. This guide explains how to navigate the API and use the naming conventions effectively.

## How the API is Organized

Assertions are grouped by domain for easier discovery:

| Domain | Examples | Count |
|--------|----------|-------|
| **Boolean** | `True`, `False` | 2 |
| **Equality** | `Equal`, `NotEqual`, `EqualValues`, `Same`, `Exactly` | 8 |
| **Comparison** | `Greater`, `Less`, `Positive` | 8 |
| **Collection** | `Contains`, `Len`, `Empty`, `ElementsMatch` | 18 |
| **Error** | `Error`, `NoError`, `ErrorIs`, `ErrorAs`, `ErrorContains` | 8 |
| **Type** | `IsType`, `Implements`, `Zero` | 7 |
| **String** | `Regexp`, `NotRegexp` | 4 |
| **Numeric** | `InDelta`, `InEpsilon` | 6 |
| **Ordering** | `IsIncreasing`, `Sorted` | 8 |
| **Panic** | `Panics`, `NotPanics` | 4 |
| **Others** | HTTP, JSON, YAML, Time, File assertions | 12 |

{{% notice style="info" title="Browse by Domain" icon="book" %}}
See the complete [API Reference](../api/_index.md) organized by domain for a detailed documentation of all assertions.
{{% /notice %}}

## Navigating the Documentation

### Quick Reference

- **[Examples](../examples)** - Practical code examples for common testing scenarios
- **[API Reference](../api/_index.md)** - Complete assertion catalog organized by domain
- **[Generics Guide](../GENERICS.md)** - Using type-safe assertions with the `T` suffix
- **[Changes](../CHANGES.md)** - All changes since fork from stretchr/testify
- **[pkg.go.dev](https://pkg.go.dev/github.com/go-openapi/testify/v2)** - godoc API reference with full signatures

### Finding the Right Assertion

1. Browse the [API Reference](../api/_index.md) by domain (e.g., "Collection" for slice operations)
2. Search in the [API Reference](../api/_index.md) (use search box)
3. Check (or search) the provided [Examples](../examples) for practical usage patterns
4. Check [pkg.go.dev](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert) for alphabetical listing
5. Use your editor's Go to Definition on any assertion
6. Use your IDE's autocomplete - type `assert.` and explore

## API Conventions

Understanding the naming patterns helps you find and use the right assertions quickly.

### Package Choice: `assert` vs `require`

{{< cards >}}
{{% card title="assert - Non-Fatal" %}}
**Use when**: Tests should continue after failures to gather more context

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUser(t *testing.T) {
	user := getUser()

	assert.NotNil(t, user)              // ✓ Returns false
	assert.Equal(t, "Alice", user.Name) // Still runs
	assert.True(t, user.Active)         // Still runs
}
```

**Returns**: `bool` indicating success/failure
{{% /card %}}

{{% card title="require - Fatal" %}}
**Use when**: Test cannot continue meaningfully after failure

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func TestUser(t *testing.T) {
	user := getUser()

	require.NotNil(t, user)              // ✓ Calls t.FailNow() if fails
	require.Equal(t, "Alice", user.Name) // Safe to proceed
	require.True(t, user.Active)         // user is guaranteed non-nil
}
```

**Returns**: Nothing (void) - stops test on failure
{{% /card %}}
{{< /cards >}}

### Function Variants

Each assertion comes in multiple variants following consistent patterns:

| Pattern | Example | Description |
|---------|---------|-------------|
| **Base** | `Equal(t, expected, actual)` | Standard assertion |
| **Format** (`f` suffix) | `Equalf(t, expected, actual, "checking %s", name)` | With custom message |
| **Generic** (`T` suffix) | `EqualT(t, expected, actual)` | Type-safe variant |
| **Generic + Format** (`Tf` suffix) | `EqualTf(t, expected, actual, "checking %s", name)` | Type-safe with message |

The `f` suffix follows Go's standard library convention (like `Printf`, `Errorf`): it accepts a format string followed by arguments for custom failure messages.

{{% notice style="tip" title="When to Use Each Variant" icon="lightbulb" %}}
- **Base/Generic**: Use by default - testify provides detailed failure messages
- **Format variants**: Add context when testing similar values in loops or complex scenarios
- **Generic (`T` suffix)**: Prefer for compile-time type safety and better performance
{{% /notice %}}

One (historical) exception: `EventuallyWithT` is not generic...

### Inverse Assertions

Most assertions come with their opposite variant, typically formed by adding a `Not` prefix:

| Assertion | Inverse | Pattern |
|-----------|---------|---------|
| `Equal` | `NotEqual` | `Not` prefix |
| `Nil` | `NotNil` | `Not` prefix |
| `Empty` | `NotEmpty` | `Not` prefix |
| `Contains` | `NotContains` | `Not` prefix |
| `Zero` | `NotZero` | `Not` prefix |
| `Same` | `NotSame` | `Not` prefix |
| `Panics` | `NotPanics` | `Not` prefix |
| `Regexp` | `NotRegexp` | `Not` prefix |

**Exceptions:** Some assertions use semantic opposites instead of the `Not` prefix:

| Assertion | Inverse | Reason |
|-----------|---------|--------|
| `True` | `False` | Semantic opposites (`NotTrue` doesn't sound natural) |
| `Positive` | `Negative` | Semantic opposites, except for 0 which is neither |
| `Greater` | `LessOrEqual` | Comparative opposites (and not `NotGreater`) |
| `GreaterOrEqual` | `Less` | Comparative opposites |

{{% notice style="info" title="Why Semantic Opposites?" icon="question" %}}
These exceptions follow natural English usage:
- Testing for `False` is clearer than testing for "not true"
- (strictly) `Negative` numbers are semantically opposite to (strictly) `Positive`,  unless when `Zero`, and not "not positive"
- `Less` is the natural opposite of `Greater` in comparisons
{{% /notice %}}

**More semantic opposites:**

| Assertion | Inverse | Reason |
|-----------|---------|--------|
| `Eventually` | `Never` | Semantic opposites for polling conditions |

**Not inverses:** Some assertions are independent checks, not inverses of each other:

| Assertions | Relationship |
|------------|--------------|
| `IsIncreasing` / `IsDecreasing` | Independent checks (a sequence can be neither) |
| `IsNonIncreasing` / `IsNonDecreasing` | Independent checks (a sequence can be neither) |
| `Sorted` / `NotSorted` | True inverse pair using `Not` prefix |

**Generic variants:** All inverse assertions have corresponding generic variants (suffix `T` or `Tf`):
- `NotEqualT`, `FalseT`, `NegativeT`, `IsDecreasingT`, etc.

### Argument Order Patterns

Most assertions follow the **"expected, actual"** pattern, but several categories use different conventions:

#### Standard Pattern: Expected, Actual

The majority of assertions check an actual value against an expected value:

```go
assert.Equal(t, expected, actual)
assert.NotEqual(t, expected, actual)
assert.InDelta(t, expected, actual, delta)
assert.JSONEq(t, expected, actual)
assert.YAMLEq(t, expected, actual)
assert.WithinDuration(t, expected, actual, delta)
assert.Implements(t, (*interface)(nil), object)  // Expected interface, actual object
```

#### Comparison Operators: e1, e2

Comparison assertions express the relationship directly (reads as "assert e1 > e2"):

```go
assert.Greater(t, e1, e2)           // Asserts: e1 > e2
assert.GreaterOrEqual(t, e1, e2)    // Asserts: e1 >= e2
assert.Less(t, e1, e2)              // Asserts: e1 < e2
assert.LessOrEqual(t, e1, e2)       // Asserts: e1 <= e2
```

#### Exceptions using Different Argument Orders

{{% tabs %}}
{{% tab title="Unary checks" color=green %}}

**Unary checks** (test a single value):
```go
assert.True(t, value)
assert.False(t, value)
assert.Nil(t, value)
assert.Empty(t, object)
assert.Zero(t, value)
assert.Positive(t, value)
assert.Negative(t, value)
assert.Error(t, err)
assert.NoError(t, err)
assert.Panics(t, panicFunc)
```
{{% /tab %}}
{{% tab title="Object-first" color=green %}}

**Object-first pattern** (object under test, then expected property):
```go
assert.Len(t, object, expectedLength)  // Object first, expected length second
assert.IsType(t, expectedType, object) // Expected type first, object second
```
{{% /tab %}}
{{% tab title="Container-first" color=green %}}

**Container-first pattern** (container, then element/subset):
```go
assert.Contains(t, container, element)     // Container first, element second
assert.StringContains(t, str, substring)   // String first, substring second
assert.SliceContains(t, slice, element)    // Slice first, element second
assert.Subset(t, list, subset)             // Superset first, subset second
assert.ElementsMatch(t, listA, listB)      // Either order works (symmetric)
```
{{% /tab %}}
{{% tab title="Error assertions" color=green %}}

**Error assertions** (error first, then expected property):
```go
assert.ErrorContains(t, err, substring)    // Error first, expected substring second
assert.ErrorIs(t, err, target)             // Error first, target error second
assert.ErrorAs(t, err, &target)            // Error first, target pointer second
assert.EqualError(t, err, expectedString)  // Error first, expected message second
```
{{% /tab %}}
{{% tab title="Special cases" color=green %}}

**Special cases**:
```go
assert.HTTPSuccess(t, handler, method, url, values) // Handler first, HTTP params follow
assert.Eventually(t, condition, waitFor, tick)      // Condition first, timing params follow
```
{{% /tab %}}
{{% /tabs %}}

{{% notice style="tip" title="Finding Argument Order" icon="lightbulb" %}}
When unsure about argument order:
- Check the [API Reference](../api/_index.md) for detailed signatures
- Use IDE autocomplete to see parameter names
- Consult [pkg.go.dev](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert) for complete documentation
{{% /notice %}}

### Forward Methods

Create an `Assertion` object to reduce repetition in tests with many assertions:

{{< cards >}}
{{% card title="Package-Level Functions" %}}
```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUser(t *testing.T) {
	user := getUser()

	assert.NotNil(t, user)
	assert.Equal(t, "Alice", user.Name)
	assert.True(t, user.Active)
	assert.Greater(t, user.Age, 0)
}
```
{{% /card %}}

{{% card title="Forward Methods" %}}
```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUser(t *testing.T) {
	a := assert.New(t) // Create once
	user := getUser()

	a.NotNil(user) // No 't' needed
	a.Equal("Alice", user.Name)
	a.True(user.Active)
}
```
{{% /card %}}
{{< /cards >}}

**Both styles are equivalent** - choose based on your preference and test structure.

**⚠️ Generic assertions are not available as forward methods** (this is a limitation of go generics).


## Common Usage Patterns

{{% tabs %}}
{{% tab title="Table-driven tests" color=green %}}

**Pattern 1: Table-Driven Tests**

```go
import (
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestCalculation(t *testing.T) {
	tests := slices.Values([]struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 25},
		{"negative", -3, 9},
		{"zero", 0, 0},
	})

	for tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := square(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
```
{{% /tab %}}

{{% tab title="Multiple assertions" color=green %}}

**Pattern 2: Multiple Assertions (assert for context)**

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUserValidation(t *testing.T) {
	user := createUser()

	// Use assert to see all failures
	assert.NotEmpty(t, user.Name)  // Check name
	assert.NotEmpty(t, user.Email) // Check email
	assert.Greater(t, user.Age, 0) // Check age
	// All assertions run - see complete picture
}
```
{{% /tab %}}
{{% tab title="Early exit" color=green %}}

**Pattern 3: Early Exit (use require for prerequisites)**

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestDatabaseQuery(t *testing.T) {
	db := connectDB()
	require.NotNil(t, db) // Stop if no connection

	result := db.Query("SELECT * FROM users")
	require.NoError(t, result.Error) // Stop if query fails

	// Safe to proceed - db and result are valid
	assert.NotEmpty(t, result.Rows)
}
```
{{% /tab %}}
{{% tab title="Type-safe" color=green %}}

**Pattern 4: Type-Safe Generics**

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestTypeSafety(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := getNumbers()

	// Compiler checks types at compile time
	assert.ElementsMatchT(t, expected, actual)
	assert.GreaterT(t, len(actual), 0)

	// If getNumbers() changes return type,
	// compiler catches it immediately
}
```
{{% /tab %}}
{{% /tabs %}}

## Getting Started

1. **Import the package:**
   ```go
   import "github.com/go-openapi/testify/v2/assert"
   // or
   import "github.com/go-openapi/testify/v2/require"
   ```

2. **Choose your style:**
   - Package-level: `assert.Equal(t, expected, actual)`
   - Forward methods: `a := assert.New(t); a.Equal(expected, actual)`

3. **Explore by domain:**
   - Browse [API Reference](../api/_index.md) to discover assertions
   - Check [Examples](./EXAMPLES.md) for practical patterns

4. **Use generics for type safety:**
   - See [Generics Guide](./GENERICS.md) for type-safe assertions
   - Add `T` suffix for compile-time type checking

## Best Practices

✅ **Do:**
- Use `require` for prerequisites that make subsequent assertions meaningless (or will possibly panic)
- Use `assert` when you want to see all failures in a test
- Prefer generic variants (`*T` functions) for compile-time type safety
- Use format variants (`*f`) to add context in complex scenarios
- Browse by domain in the API reference to discover relevant assertions

❌ **Don't:**
- Don't mix `assert` and `require` randomly - be intentional
- Don't add unnecessary format messages - testify provides detailed output
- Don't ignore compiler errors from generic variants - they reveal design issues
- Don't forget that both packages provide the same assertions with different behavior

---

{{% notice style="success" title="Ready to Test" icon="check" %}}
**Next Steps:**
- Explore [Examples](../examples) for practical usage patterns
- Browse the [API Reference](../api/_index.md) to discover assertions
- Read the [Generics Guide](../GENERICS.md) for type-safe testing
- Check [pkg.go.dev](https://pkg.go.dev/github.com/go-openapi/testify/v2) for complete reference
{{% /notice %}}

---

## Customization

### Using a Custom YAML Unmarshaler

By default, testify uses `gopkg.in/yaml.v3` for YAML assertions (e.g. `YAMLEq`) when you import the standard
`enable/yaml/v2` package.

However, you can register a custom YAML unmarshaler to use alternative libraries like
[goccy/go-yaml](https://github.com/goccy/go-yaml), either because you need additional features such as colored error
messages or better performance.

#### How It Works

The YAML support in testify works through a registration mechanism:

1. `internal/assertions/yaml.go` calls `yaml.Unmarshal()` - an abstraction layer
2. The abstraction layer panics if no unmarshaler is registered
3. The `enable/yaml/v2` package registers `gopkg.in/yaml.v3` via `init()` when imported (e.g. on blank import)
4. You can register a custom unmarshaler using `enable/stubs/yaml.EnableYAMLWithUnmarshal()`

#### Example: Using goccy/go-yaml

Create a custom enable package in your test code:

```go
package testutil

import (
	goccyyaml "github.com/goccy/go-yaml"
	yamlstub "github.com/go-openapi/testify/v2/enable/stubs/yaml"
)

func init() {
	// Register goccy/go-yaml as the YAML unmarshaler
	yamlstub.EnableYAMLWithUnmarshal(goccyyaml.Unmarshal)
}
```

Then import your custom enable package in your tests:

```go
// File: mypackage/user_test.go
package mypackage

import (
	"testing"

	_ "yourmodule/internal/testutil" // Register goccy/go-yaml

	"github.com/go-openapi/testify/v2/assert"
)

func TestUserYAML(t *testing.T) {
	expected := `
name: Alice
email: alice@example.com
age: 30
`
	actual := serializeUser(getUser())

	// Now uses goccy/go-yaml under the hood
	assert.YAMLEq(t, expected, actual)
}
```

#### Why Use a Custom YAML Library?

Different YAML libraries offer different trade-offs:

**`gopkg.in/yaml.v3` (default):**
- De facto standard library for Go YAML
- Widely used and well-tested
- Complete YAML 1.2 support

**`github.com/goccy/go-yaml`:**
- Better performance (up to 2-3x faster)
- Colored error messages for debugging
- Better error reporting with line/column numbers
- JSON-like syntax support
- Comment preservation (useful for config testing)

#### Important Notes

1. **Register once:** Call `EnableYAMLWithUnmarshal()` only once, typically in an `init()` function
2. **Not concurrent-safe:** The registration is global and should happen during package or main program initialization
3. **Signature compatibility:** The custom unmarshaler must match the signature `func([]byte, any) error`
4. **No mixing:** Don't import both `github.com/go-openapi/testify/enable/yaml/v2` and your custom enable package - choose one

#### Advanced: Wrapping Unmarshalers

You can also wrap an unmarshaler to add custom behavior:

```go
package testutil

import (
	"fmt"
	"log"

	goccyyaml "github.com/goccy/go-yaml"

	yamlstub "github.com/go-openapi/testify/v2/enable/stubs/yaml"
)

func init() {
	// Wrap the unmarshaler to add logging or validation
	yamlstub.EnableYAMLWithUnmarshal(func(data []byte, v any) error {
		// Custom pre-processing
		if len(data) == 0 {
			return fmt.Errorf("empty YAML document")
		}

		// Call the actual unmarshaler
		err := goccyyaml.Unmarshal(data, v)
		// Custom post-processing
		if err != nil {
			log.Printf("YAML unmarshal error: %v", err)
		}

		return err
	})
}
```

This pattern allows you to add logging, validation, or transformation logic around any YAML library.

---

## See Also

- [Examples](./EXAMPLES.md) - Practical code examples for common testing scenarios
- [Tutorial](./TUTORIAL.md) - Comprehensive guide to writing great tests with testify patterns
- [Generics Guide](./GENERICS.md) - Type-safe assertions with compile-time checking
- [Migration Guide](./MIGRATION.md) - Migrating from stretchr/testify v1
- [API Reference](../api/_index.md) - Complete assertion catalog organized by domain
