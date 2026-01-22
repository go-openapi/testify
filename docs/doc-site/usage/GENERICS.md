---
title: Generics
description: Using generic assertions.
weight: 10
---

# Using Generic Assertions

Testify v2 provides **38 generic assertion functions** that offer compile-time type safety alongside the traditional reflection-based assertions. Generic variants are identified by the `T` suffix (e.g., `EqualT`, `GreaterT`, `ElementsMatchT`).

{{% notice style="success" title="Type Safety First" icon="check" %}}
Generic assertions catch type mismatches **when writing tests**, not when running them. The performance improvements (1.2x-81x faster) are a bonus on top of this primary benefit.
{{% /notice %}}

## Quick Start

Generic assertions work exactly like their reflection-based counterparts, but with compile-time type checking:

{{< cards >}}
{{% card title="Reflection-based" %}}
```go
import "github.com/go-openapi/testify/v2/assert"

func TestUser(t *testing.T) {
    expected := 42
    actual := getUserAge()

    // Compiles, but type errors appear at runtime
    assert.Equal(t, expected, actual)
}
```
{{% /card %}}

{{% card title="Generic (Type-safe)" %}}
```go
import "github.com/go-openapi/testify/v2/assert"

func TestUser(t *testing.T) {
    expected := 42
    actual := getUserAge()

    // Compiler checks types immediately
    assert.EqualT(t, expected, actual)
}
```
{{% /card %}}
{{< /cards >}}

## When to Use Generic Variants

### ✅ Use Generic Variants (`*T` functions) When:

1. **Testing with known concrete types** - The most common case
   ```go
   assert.EqualT(t, 42, result)              // int comparison
   assert.GreaterT(t, count, 0)              // numeric comparison
   assert.ElementsMatchT(t, expected, actual) // slice comparison
   ```

2. **You want refactoring safety** - Compiler catches broken tests immediately
   ```go
   // If getUserIDs() changes from []int to []string,
   // the compiler flags this line immediately
   assert.ElementsMatchT(t, expectedIDs, getUserIDs())
   ```

3. **IDE assistance matters** - Autocomplete suggests only correctly-typed variables
   ```go
   // Typing: assert.EqualT(t, expectedUser, actual
   //                                              ^
   // IDE suggests: actualUser ✓  (correct type)
   //               actualOrder ✗ (wrong type - grayed out)
   ```

4. **Performance-critical tests** - See [benchmarks](../../project/maintainers/BENCHMARKS.md) for 1.2-81x speedups

### 🔄 Use Reflection Variants (no suffix) When:

1. **Intentionally comparing different types** - Especially with `EqualValues`
   ```go
   // Comparing int and int64 for semantic equality
   assert.EqualValues(t, int64(42), int32(42))  // ✓ Reflection handles this
   assert.EqualT(t, int64(42), int32(42))       // ❌ Compiler error
   ```

2. **Working with heterogeneous collections** - `[]any` or `interface{}` slices
   ```go
   mixed := []any{1, "string", true}
   assert.Contains(t, mixed, "string")  // ✓ Reflection works
   ```

3. **Dynamic type scenarios** - Where compile-time type is unknown
   ```go
   var result interface{} = getResult()
   assert.Equal(t, expected, result)  // ✓ Reflection handles dynamic types
   ```

4. **Backward compatibility** - Existing test code using reflection-based assertions

## Type Safety Benefits

### Catching Refactoring Errors

Generic assertions act as a safety net during refactoring:

{{< cards >}}
{{% card title="Without Generics ❌" %}}
```go
// Original code
type UserID int
var userIDs []UserID

assert.ElementsMatch(t, userIDs, getActiveUsers())

// Later: UserID changes to string
type UserID string
var userIDs []UserID

// Test still compiles!
// Fails mysteriously at runtime or passes with wrong comparison
assert.ElementsMatch(t, userIDs, getActiveUsers())
```
{{% /card %}}

{{% card title="With Generics ✅" %}}
```go
// Original code
type UserID int
var userIDs []UserID

assert.ElementsMatchT(t, userIDs, getActiveUsers())

// Later: UserID changes to string
type UserID string
var userIDs []UserID

// Compiler immediately flags the error
assert.ElementsMatchT(t, userIDs, getActiveUsers())
// ❌ Compile error: type mismatch!
```
{{% /card %}}
{{< /cards >}}

### Preventing Wrong Comparisons

Generic assertions force you to think about what you're comparing:

{{< cards >}}
{{% card title="Pointer vs Value Comparison" %}}
```go
expected := &User{ID: 1, Name: "Alice"}
actual := &User{ID: 1, Name: "Alice"}

// Reflection: Compares pointer addresses (probably wrong)
assert.Equal(t, expected, actual)  // ✗ Fails (different addresses)

// Generic: Makes the intent explicit
assert.EqualT(t, expected, actual)   // Compares pointers
assert.EqualT(t, *expected, *actual) // Compares values ✓
```
{{% /card %}}

{{% card title="Type Confusion Prevention" %}}
```go
userID := 42
orderID := "ORD-123"

// Reflection: Compiles, wrong comparison
assert.Equal(t, userID, orderID)  // Runtime failure

// Generic: Compiler catches the mistake
assert.EqualT(t, userID, orderID)  // ❌ Compile error!
```
{{% /card %}}
{{< /cards >}}

## Available Generic Functions

Testify v2 provides generic variants across all major domains:

### Equality (4 functions)
- `EqualT[V comparable]` - Type-safe equality for comparable types
- `NotEqualT[V comparable]` - Type-safe inequality
- `SameT[V comparable]` - Pointer identity check
- `NotSameT[V comparable]` - Different pointer check

### Comparison (6 functions)
- `GreaterT[V Ordered]` - Type-safe greater-than comparison
- `GreaterOrEqualT[V Ordered]` - Type-safe >=
- `LessT[V Ordered]` - Type-safe less-than comparison
- `LessOrEqualT[V Ordered]` - Type-safe <=
- `PositiveT[V SignedNumeric]` - Assert value > 0
- `NegativeT[V SignedNumeric]` - Assert value < 0

### Collection (12 functions)
- `StringContainsT[S Text]` - String/byte slice contains substring
- `SliceContainsT[E comparable]` - Slice contains element
- `MapContainsT[K comparable, V any]` - Map contains key
- `SeqContainsT[E comparable]` - Iterator contains element (Go 1.23+)
- `ElementsMatchT[E comparable]` - Slices have same elements (any order)
- `SliceSubsetT[E comparable]` - Slice is subset of another
- Plus negative variants: `*NotContainsT`, `NotElementsMatchT`, `SliceNotSubsetT`

### Ordering (6 functions)
- `IsIncreasingT[E Ordered]` - Slice elements strictly increasing
- `IsDecreasingT[E Ordered]` - Slice elements strictly decreasing
- `IsNonIncreasingT[E Ordered]` - Slice elements non-increasing (allows equal)
- `IsNonDecreasingT[E Ordered]` - Slice elements non-decreasing (allows equal)
- `SortedT[E Ordered]` - Slice is sorted (generic-only function)
- `NotSortedT[E Ordered]` - Slice is not sorted (generic-only function)

### Numeric (2 functions)
- `InDeltaT[V Measurable]` - Numeric comparison with absolute delta (supports integers and floats)
- `InEpsilonT[V Measurable]` - Numeric comparison with relative epsilon (supports integers and floats)

### Boolean (2 functions)
- `TrueT[B Boolean]` - Assert boolean is true
- `FalseT[B Boolean]` - Assert boolean is false

### String (2 functions)
- `RegexpT[S Text]` - String matches regex (string or []byte)
- `NotRegexpT[S Text]` - String doesn't match regex

### Type (2 functions)
- `IsOfTypeT[EType any]` - Assert value is of type EType (no dummy value needed!)
- `IsNotOfTypeT[EType any]` - Assert value is not of type EType

### JSON & YAML (2 functions)
- `JSONEqT[S Text]` - JSON strings are semantically equal
- `YAMLEqT[S Text]` - YAML strings are semantically equal

{{% notice style="info" title="See Complete API" icon="book" %}}
For detailed documentation of all generic functions, see the [API Reference](../../api/_index.md) organized by domain.
{{% /notice %}}

## Practical Examples

### Example 1: Collection Testing

{{< cards >}}
{{% card title="Type-Safe Collection Assertions" %}}
```go
func TestUserPermissions(t *testing.T) {
    user := getUser(123)

    expectedPerms := []string{"read", "write"}
    actualPerms := user.Permissions

    // Compiler ensures both slices are []string
    assert.ElementsMatchT(t, expectedPerms, actualPerms)

    // Check subset relationship
    assert.SliceSubsetT(t, []string{"read"}, actualPerms)
}
```
{{% /card %}}

{{% card title="Iterator Support (Go 1.23+)" %}}
```go
func TestSequenceContains(t *testing.T) {
    // iter.Seq[int] from Go 1.23
    numbers := slices.Values([]int{1, 2, 3, 4, 5})

    // Type-safe iterator checking
    assert.SeqContainsT(t, numbers, 3)
    assert.SeqNotContainsT(t, numbers, 99)
}
```
{{% /card %}}
{{< /cards >}}

### Example 2: Numeric Comparisons

{{< cards >}}
{{% card title="Ordered Types" %}}
```go
func TestPricing(t *testing.T) {
    price := calculatePrice(item)
    discount := calculateDiscount(item)

    // Type-safe numeric comparisons
    assert.PositiveT(t, price)
    assert.GreaterT(t, price, discount)
    assert.LessOrEqualT(t, discount, price)
}
```
{{% /card %}}

{{% card title="Float Comparisons" %}}
```go
func TestPhysicsCalculation(t *testing.T) {
    result := calculateVelocity(mass, force)
    expected := 42.0

    // Type-safe float comparison with delta
    assert.InDeltaT(t, expected, result, 1e-6)

    // Or with epsilon (relative error)
    assert.InEpsilonT(t, expected, result, 0.001)
}
```
{{% /card %}}
{{< /cards >}}

### Example 3: Type Checking Without Dummy Values

The `IsOfTypeT` function eliminates the need for dummy values:

{{< cards >}}
{{% card title="Old Way (Reflection)" %}}
```go
func TestGetUser(t *testing.T) {
    result := getUser(123)

    // Need to create a dummy User instance
    assert.IsType(t, User{}, result)

    // Or use a pointer dummy
    assert.IsType(t, (*User)(nil), result)
}
```
{{% /card %}}

{{% card title="New Way (Generic)" %}}
```go
func TestGetUser(t *testing.T) {
    result := getUser(123)

    // No dummy value needed!
    assert.IsOfTypeT[User](t, result)

    // For pointer types
    assert.IsOfTypeT[*User](t, result)
}
```
{{% /card %}}
{{< /cards >}}

### Example 4: Sorting and Ordering

{{< cards >}}
{{% card title="Ordering Checks" %}}
```go
func TestSortedData(t *testing.T) {
    timestamps := []int64{
        1640000000,
        1640000100,
        1640000200,
    }

    // Type-safe ordering assertions
    assert.IsIncreasingT(t, timestamps)
    assert.SortedT(t, timestamps)  // Generic-only function
}
```
{{% /card %}}

{{% card title="Custom Ordered Types" %}}
```go
type Priority int

const (
    Low Priority = iota
    Medium
    High
)

func TestPriorities(t *testing.T) {
    tasks := []Priority{Low, Medium, High}

    // Works with Ordered types (custom types supported)
    assert.IsNonDecreasingT(t, tasks)
}
```
{{% /card %}}
{{< /cards >}}

## Migration Guide

### Step 1: Identify High-Value Targets

Start with the most common assertions that benefit most from type safety:

```go
// High value: Collection operations (also get big performance wins)
assert.Equal → assert.EqualT
assert.ElementsMatch → assert.ElementsMatchT
assert.Contains → assert.ContainsT (SliceContainsT/MapContainsT/StringContainsT)

// High value: Comparisons (eliminate allocations)
assert.Greater → assert.GreaterT
assert.Less → assert.LessT
assert.Positive → assert.PositiveT

// High value: Type checks (cleaner API)
assert.IsType(t, User{}, v) → assert.IsOfTypeT[User](t, v)
```

### Step 2: Automated Search & Replace

Use your IDE or tools to find and replace systematically:

```bash
# Find all Equal assertions
grep -r "assert\.Equal(" . --include="*_test.go"

# Find all require.Greater assertions
grep -r "require\.Greater(" . --include="*_test.go"
```

### Step 3: Fix Compiler Errors

The compiler will catch type mismatches. This is a feature, not a bug:

{{< cards >}}
{{% card title="Compiler Error" %}}
```go
// Original code
assert.EqualT(t, int64(result), count)
// ❌ Error: mismatched types int64 and int
```
{{% /card %}}

{{% card title="Fix Option 1: Same Type" %}}
```go
// Convert to same type
assert.EqualT(t, int64(result), int64(count))
```
{{% /card %}}

{{% card title="Fix Option 2: Use Reflection" %}}
```go
// If cross-type comparison is intentional
assert.Equal(t, int64(result), count)
```
{{% /card %}}
{{< /cards >}}

### Step 4: Incremental Adoption

You don't need to migrate everything at once:

```go
func TestMixedAssertions(t *testing.T) {
    // Use generic where types are known
    assert.EqualT(t, 42, getAge())
    assert.GreaterT(t, count, 0)

    // Keep reflection for dynamic types
    var result interface{} = getResult()
    assert.Equal(t, expected, result)

    // Both styles coexist peacefully
}
```

## Performance Benefits

Generic assertions provide significant performance improvements, especially for collection operations:

| Operation | Speedup | When It Matters |
|-----------|---------|-----------------|
| **ElementsMatchT** | **21-81x faster** | Large collections, hot test paths |
| **EqualT** | **10-13x faster** | Most common assertion |
| **GreaterT/LessT** | **10-22x faster** | Numeric comparisons |
| **SliceContainsT** | **16x faster** | Collection membership tests |

{{% notice style="success" title="Learn More" icon="chart-line" %}}
See the complete [Performance Benchmarks](../../project/maintainers/BENCHMARKS.md) for detailed analysis and real benchmark results.
{{% /notice %}}

## Best Practices

### ✅ Do

1. **Prefer generic variants by default** - Type safety is always valuable
   ```go
   assert.EqualT(t, expected, actual)  // ✓ Type safe
   ```

2. **Let the compiler guide you** - Type errors reveal design issues
   ```go
   // Compiler error reveals you're comparing wrong types
   assert.EqualT(t, userID, orderID)  // ❌ Good - catches mistake!
   ```

3. **Use explicit types for clarity**
   ```go
   assert.IsOfTypeT[*User](t, result)  // ✓ Clear intent
   ```

4. **Leverage performance wins in hot paths** - Generic assertions are faster
   ```go
   // Table-driven tests with many iterations
   for _, tc := range testCases {
       assert.EqualT(t, tc.expected, tc.actual)  // ✓ Fast
   }
   ```

### ❌ Don't

1. **Don't force generics for dynamic types**
   ```go
   var result interface{} = getResult()
   assert.Equal(t, expected, result)  // ✓ Reflection is fine here
   ```

2. **Don't use reflection to avoid fixing types**
   ```go
   // Bad: Using reflection to bypass type safety
   assert.Equal(t, expected, actual)  // ✗ Defeats the purpose

   // Good: Fix the types or use EqualValues if intentional
   assert.EqualT(t, expected, actual)  // ✓ Type safe
   ```

3. **Don't create unnecessary type conversions**
   ```go
   // Bad: Unnecessary conversion
   assert.EqualT(t, int64(42), int64(result))

   // Good: Work with natural types
   assert.EqualT(t, 42, result)
   ```

## Type Constraints Reference

Generic assertions use custom type constraints defined in `internal/assertions/generics.go`:

| Constraint | Definition | Description | Example Types |
|------------|------------|-------------|---------------|
| `comparable` | Go built-in | Types that support `==` and `!=` | `int`, `string`, `bool`, pointers, structs (if all fields are comparable) |
| `Boolean` | `~bool` | Boolean and named bool types | `bool`, `type MyBool bool` |
| `Text` | `~string \| ~[]byte` | String or byte slice types | `string`, `[]byte`, custom string/byte types |
| `Ordered` | `cmp.Ordered \| []byte \| time.Time` | **Extends** `cmp.Ordered` with byte slices and time | Standard ordered types plus `[]byte` and `time.Time` |
| `SignedNumeric` | `~int... \| ~float32 \| ~float64` | Signed integers and floats | `int`, `int8`-`int64`, `float32`, `float64` |
| `UnsignedNumeric` | `~uint...` | Unsigned integers | `uint`, `uint8`-`uint64` |
| `Measurable` | `SignedNumeric \| UnsignedNumeric` | All numeric types (for delta comparisons) | Used by `InDeltaT`/`InEpsilonT` - supports **integers AND floats** |
| `RegExp` | `Text \| *regexp.Regexp` | Regex pattern or compiled regexp | `string`, `[]byte`, `*regexp.Regexp` |

{{% notice style="primary" title="Key Differences from Standard Go Constraints" icon="info" %}}
- **`Ordered` is extended**: Adds `[]byte` and `time.Time` to `cmp.Ordered` for seamless `bytes.Compare()` and `time.Time.Compare()` support
- **`Measurable` supports integers**: `InDeltaT` and `InEpsilonT` work with both integers and floats, not just floating-point types
- **Custom type support**: All constraints use the `~` operator to support custom types (e.g., `type UserID int`)
{{% /notice %}}

## Summary

**Generic assertions in testify v2 provide:**

✅ **Type Safety**: Catch errors when writing tests, not when running them
✅ **Performance**: 1.2x to 81x faster than reflection-based assertions
✅ **Better IDE Support**: Autocomplete suggests correctly-typed values
✅ **Refactoring Safety**: Compiler catches broken tests immediately
✅ **Zero Downside**: Always as fast or faster than reflection variants

**Start using generic assertions today** - add the `T` suffix to your existing assertions and let the compiler guide you to better, safer tests.

---

{{% notice style="tip" title="Quick Reference" icon="lightbulb" %}}
- **Generic functions**: Add `T` suffix (e.g., `EqualT`, `GreaterT`, `ElementsMatchT`)
- **Format variants**: Add `Tf` suffix (e.g., `EqualTf`, `GreaterTf`)
- **When to use**: Prefer generics for known concrete types
- **When not to**: Keep reflection for dynamic types and cross-type comparisons
- **Performance**: See [benchmarks](../../project/maintainers/BENCHMARKS.md) for dramatic speedups
{{% /notice %}}
