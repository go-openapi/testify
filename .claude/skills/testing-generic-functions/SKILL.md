# Testing Generic Functions with Table-Driven Tests

## The Challenge

Go generics require type parameters to be resolved at compile time. This creates a problem for traditional table-driven tests where test cases are stored in a slice of `any`:

```go
// This DOES NOT work - type parameter cannot be inferred from 'any'
cases := []struct {
    name   string
    value1 any
    value2 any
}{
    {"int", 1, 2},
    {"string", "a", "b"},
}

for _, c := range cases {
    GreaterT(mock, c.value1, c.value2) // Error: cannot infer type parameter
}
```

## The Solution: Test Function Closures

Wrap each test case in a closure where the type parameter is resolved at slice construction time, not iteration time.

### Step 1: Define a Test Case Struct

```go
import "testing"

// genericTestCase wraps a test function with its name for table-driven tests.
type genericTestCase struct {
	name string
	test func(*testing.T)
}
```

### Step 2: Create a Generic Test Helper

The helper function is generic and returns a closure. The type parameter is resolved when the closure is created (at slice construction), not when it's executed.

```go
import "testing"

// testGreaterT creates a test function for GreaterT with specific type V.
// Type parameter V is resolved when this function is called, not when the
// returned closure executes.
func testGreaterT[V Ordered](successE1, successE2, failE1, failE2 V) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		// Success case: e1 > e2
		True(t, GreaterT(mock, successE1, successE2))

		// Failure case: e1 <= e2
		False(t, GreaterT(mock, failE1, failE2))

		// Equal values should fail
		False(t, GreaterT(mock, successE1, successE1))
	}
}
```

### Step 3: Create Iterator Function for Test Cases

Use the iterator pattern (`iter.Seq`) to return test cases:

```go
import (
	"iter"
	"slices"
)

func greaterTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Numeric types - type inferred from arguments
		{"int", testGreaterT[int](2, 1, 1, 2)},
		{"int8", testGreaterT[int8](2, 1, 1, 2)},
		{"int16", testGreaterT[int16](2, 1, 1, 2)},
		{"int32", testGreaterT[int32](2, 1, 1, 2)},
		{"int64", testGreaterT[int64](2, 1, 1, 2)},
		{"uint", testGreaterT[uint](2, 1, 1, 2)},
		{"uint8", testGreaterT[uint8](2, 1, 1, 2)},
		{"uint16", testGreaterT[uint16](2, 1, 1, 2)},
		{"uint32", testGreaterT[uint32](2, 1, 1, 2)},
		{"uint64", testGreaterT[uint64](2, 1, 1, 2)},
		{"uintptr", testGreaterT[uintptr](2, 1, 1, 2)},
		{"float32", testGreaterT[float32](2.0, 1.0, 1.0, 2.0)},
		{"float64", testGreaterT[float64](2.0, 1.0, 1.0, 2.0)},
		{"string", testGreaterT[string]("b", "a", "a", "b")},

		// Special types requiring dedicated setup functions
		{"time.Time", testGreaterTTime()},
		{"[]byte", testGreaterTBytes()},

		// Custom types to verify constraint satisfaction
		{"custom int type", testGreaterTCustomInt()},
	})
}
```

### Step 4: Handle Special Types

Some types need dedicated test functions because they require specific setup:

```go
import (
	"testing"
	"time"
)

// testGreaterTTime tests GreaterT with time.Time values.
func testGreaterTTime() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t0 := time.Now()
		t1 := t0.Add(-time.Second) // t1 is before t0

		True(t, GreaterT(mock, t0, t1))  // t0 > t1
		False(t, GreaterT(mock, t1, t0)) // t1 < t0
		False(t, GreaterT(mock, t0, t0)) // equal
	}
}

// testGreaterTBytes tests GreaterT with []byte values.
func testGreaterTBytes() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		b1 := []byte("b")
		b2 := []byte("a")

		True(t, GreaterT(mock, b1, b2))
		False(t, GreaterT(mock, b2, b1))
		False(t, GreaterT(mock, b1, b1))
	}
}

// testGreaterTCustomInt verifies custom types satisfying the constraint work.
type myInt int

func testGreaterTCustomInt() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, GreaterT(mock, myInt(2), myInt(1)))
		False(t, GreaterT(mock, myInt(1), myInt(2)))
	}
}
```

### Step 5: Write the Test Function

```go
import "testing"

func TestCompareGreaterT(t *testing.T) {
	t.Parallel()

	for tc := range greaterTCases() {
		t.Run(tc.name, tc.test)
	}
}
```

## Complete Pattern Summary

```go
import (
	"iter"
	"slices"
	"testing"
)

// 1. Test case wrapper struct
type genericTestCase struct {
	name string
	test func(*testing.T)
}

// 2. Generic test helper returning closure
func testFunctionUnderTest[V Constraint](args V) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		// Test logic using args with type V
	}
}

// 3. Special type handlers
func testFunctionUnderTestSpecialType() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		// Test logic with special setup
	}
}

// 4. Iterator function
func functionUnderTestCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"type1", testFunctionUnderTest[Type1](value1)},
		{"type2", testFunctionUnderTest[Type2](value2)},
		{"special", testFunctionUnderTestSpecialType()},
	})
}

// 5. Test function
func TestFunctionUnderTest(t *testing.T) {
	t.Parallel()
	for tc := range functionUnderTestCases() {
		t.Run(tc.name, tc.test)
	}
}
```

## Key Insights

1. **Type resolution timing**: Generic type parameters are resolved at compile time. By creating closures at slice construction, each closure has its type parameter already resolved.

2. **Closure captures**: The generic test helper captures the typed arguments in its closure, preserving type information for when the test executes.

3. **Special type handling**: Types like `time.Time` and `[]byte` need dedicated functions because they require specific construction patterns that don't fit the simple value pattern.

4. **Custom type testing**: Always include at least one custom type (e.g., `type myInt int`) to verify the constraint works with user-defined types, not just built-in types.

5. **Parallel execution**: Each closure is independent, enabling parallel test execution with `t.Parallel()`.

## When to Use This Pattern

- Testing generic functions with type constraints
- When you need to verify behavior across multiple types satisfying a constraint
- When traditional table-driven tests fail due to type inference limitations
- Testing functions from `internal/assertions/` that have generic variants (e.g., `GreaterT`, `LessT`, `PositiveT`)
