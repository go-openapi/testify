---
title: "Tutorial"
description: "How to write great tests with go and testify"
weight: 3
---

## What makes a good test?

A good test is:

1. **Focused** - Tests one logical concept
2. **Independent** - Can run in any order, in parallel
3. **Repeatable** - Same input always produces same result
4. **Fast** - Runs quickly to encourage frequent execution
5. **Have clear expectations** - Failure messages immediately show what broke

**With testify, you write tests that read like documentation:**

```go
func TestUserCreation(t *testing.T) {
    user := CreateUser("alice@example.com")

    require.NotNil(t, user)
    assert.Equal(t, "alice@example.com", user.Email) // if user is nil, will fail and stop before
    assert.True(t, user.Active)
}
```
{{% notice style="tip" title="tip" icon="meteor" %}}
Adopt a test layout similar to your functionality.

```sh
# ❌ Don't do this - confusing
boolean.go
file.go
all_test.go
```

```go
// ✅ Better - clear mapping between features and tests
boolean.go
boolean_test.go
file.go
file_test.go
```
{{% /notice %}}

The assertions are self-documenting - you can read the test and immediately understand what behavior is being verified.

---

## Patterns

### Simple test logic

Oftentimes, much of the test logic can be replaced by a proper use of `require`.

```go
// ❌ Don't do this - repetitive and hard to maintain
func TestUserCreation(t *testing.T) {
    user := CreateUser("alice@example.com")

    if assert.NotNil(t, user) {
        assert.Equal(t, "alice@example.com", user.Email) // if user is nil, will skip this test
        assert.True(t, user.Active)
    }
}
```

```go
// ✅ Better - linear flow, no indented subcases
func TestUserCreation(t *testing.T) {
    user := CreateUser("alice@example.com")

    require.NotNil(t, user)
    assert.Equal(t, "alice@example.com", user.Email) // if user is nil, will fail and stop before
    assert.True(t, user.Active)
}
```

### Table-Driven Tests with Iterator Pattern

The **iterator pattern** is the idiomatic way to write table-driven tests in Go 1.23+. This repository uses it extensively, and you should too.

#### Why Table-Driven Tests?

Instead of writing separate test functions for each case:

```go
// ❌ Don't do this - repetitive and hard to maintain
func TestAdd_PositiveNumbers(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result)
}

func TestAdd_NegativeNumbers(t *testing.T) {
    result := Add(-2, -3)
    assert.Equal(t, -5, result)
}

func TestAdd_MixedSigns(t *testing.T) {
    result := Add(-2, 3)
    assert.Equal(t, 1, result)
}
```

Write one test function with multiple cases:

```go
// ✅ Better - all cases in one place
func TestAdd(t *testing.T) {
    // All test cases defined once
    // Test logic written once
    // Easy to add new cases
    for c := range addTestCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            result := Add(c.a, c.b)
            assert.Equal(t, c.expected, result)
        })
    }
}

func addTestCases() iter.Seq[addTestCase] {
    ...
}
```

#### The Iterator Pattern

**Structure:**

```go
import (
    "iter"
    "slices"
    "testing"
    "github.com/go-openapi/testify/v2/assert"
)

// 1. Define a test case struct
type addTestCase struct {
    name     string
    a, b     int
    expected int
}

// 2. Create an iterator function returning iter.Seq[T]
func addTestCases() iter.Seq[addTestCase] {
    return slices.Values([]addTestCase{
        {
            name:     "positive numbers",
            a:        2,
            b:        3,
            expected: 5,
        },
        {
            name:     "negative numbers",
            a:        -2,
            b:        -3,
            expected: -5,
        },
        {
            name:     "mixed signs",
            a:        -2,
            b:        3,
            expected: 1,
        },
        {
            name:     "with zero",
            a:        0,
            b:        5,
            expected: 5,
        },
    })
}

// 3. Test function iterates over cases using range
func TestAdd(t *testing.T) {
    t.Parallel()

    for c := range addTestCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            result := Add(c.a, c.b)
            assert.Equal(t, c.expected, result)
        })
    }
}
```

#### Why This Pattern Is Better

**Clean separation of concerns:**
- Test data (in iterator function) separate from test logic (in test function)
- Easy to see all test cases at a glance
- Easy to add new cases without touching test logic

**Type safety:**
- Compiler enforces struct fields
- No risk of wrong number of arguments
- IDE autocomplete works perfectly

**Excellent for parallel execution:**
- Both the outer test and subtests can run in parallel
- `t.Parallel()` catches race conditions early

**Reusable:**
- Iterator functions can be reused across multiple test functions
- Share test cases between related tests

**Maintainable:**
- Adding a case: just append to the slice
- Changing test logic: edit one place
- Renaming fields: IDE refactoring works

#### Comparison with Traditional Pattern

**Traditional inline pattern:**

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        // Test data mixed with test function
        // Hard to reuse
        // No named fields - order matters
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

**Iterator pattern:**

```go
// Test logic separate and clean
func TestAdd(t *testing.T) {
    t.Parallel()
    for c := range addTestCases() {  // Clean iteration
        // ...
    }
}

type addTestCase struct {
    name     string
    a, b     int
    expected int
}

// Test data in separate function - clean, reusable
func addTestCases() iter.Seq[addTestCase] {
    return slices.Values([]addTestCase{
        {
            name:     "positive numbers",  // Named fields
            a:        2,                   // Self-documenting
            b:        3,
            expected: 5,
        },
        // More cases...
    })
}
```

#### When to Use Iterator Pattern

**Always use it for:**
- Any test with 2+ test cases
- Tests requiring complex setup per case
- Tests that benefit from parallel execution
- Any table-driven test scenario

**Example - complex setup:**

```go
func TestUserValidation(t *testing.T) {
    t.Parallel()

    for c := range userValidationCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            err := ValidateUser(c.user)

            if c.shouldErr {
                assert.Error(t, err)
                assert.ErrorContains(t, err, c.errMsg)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

type userValidationCase struct {
    name      string
    user      User
    shouldErr bool
    errMsg    string
}

func userValidationCases() iter.Seq[userValidationCase] {
    return slices.Values([]userValidationCase{
        {
            name: "valid user",
            user: User{
                Name:  "Alice",
                Email: "alice@example.com",
                Age:   25,
            },
            shouldErr: false,
        },
        {
            name: "missing email",
            user: User{
                Name: "Bob",
                Age:  30,
            },
            shouldErr: true,
            errMsg:    "email is required",
        },
        {
            name: "invalid age",
            user: User{
                Name:  "Charlie",
                Email: "charlie@example.com",
                Age:   -5,
            },
            shouldErr: true,
            errMsg:    "age must be positive",
        },
    })
}
```

---

### Using testify with Iterator Pattern

The iterator pattern works beautifully with testify's forward methods:

```go
func TestUserOperations(t *testing.T) {
    t.Parallel()

    for c := range userOperationCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()
            a := assert.New(t)  // Forward assertion object

            user := PerformOperation(c.input)

            // Clean assertions without repeating 't'
            a.NotNil(user)
            a.Equal(c.expectedName, user.Name)
            a.Greater(user.ID, 0)
        })
    }
}
```

---

### Helper Functions with t.Helper()

When extracting common assertions into helper functions, use `t.Helper()` to get better error messages:

```go
func assertUserValid(t *testing.T, user *User) {
    t.Helper()  // Makes test failures point to the caller

    assert.NotNil(t, user)
    assert.NotEmpty(t, user.Name)
    assert.NotEmpty(t, user.Email)
    assert.Greater(t, user.Age, 0)
}

func TestUserCreation(t *testing.T) {
    user := CreateUser("alice@example.com")

    // If this fails, error points HERE, not inside assertUserValid
    assertUserValid(t, user)
}
```

Without `t.Helper()`, failures would show the line number inside `assertUserValid`, making it harder to find the actual failing test.

---

### Parallel Test Execution

Always use `t.Parallel()` unless you have a specific reason not to:

```go
func TestAdd(t *testing.T) {
    t.Parallel()  // Outer test runs in parallel

    for c := range addTestCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()  // Each subtest runs in parallel

            result := Add(c.a, c.b)
            assert.Equal(t, c.expected, result)
        })
    }
}
```

**Benefits:**
- Tests run faster
- Catches race conditions and shared state bugs
- Encourages writing independent tests

**When NOT to use parallel:**
- Tests that modify global state
- Tests that use the same external resource (file, database, etc.)
- Integration tests with shared setup

---

### Setup and Teardown

Use `defer` for cleanup:

```go
func TestDatabaseOperations(t *testing.T) {
    db := setupTestDatabase(t)
    t.Cleanup(func() {
        _ = db.Close()  // Always runs, even if test fails
    }

    user := &User{Name: "Alice"}
    err := db.Save(user)
    require.NoError(t, err)  // Stop if save fails

    loaded, err := db.Find(user.ID)
    require.NoError(t, err)
    assert.Equal(t, "Alice", loaded.Name)
}
```

**Pattern for resources:**
1. Create resource
2. Immediately defer cleanup
3. Use the resource
4. Cleanup happens automatically

---

### Edge Cases to Test

Always include these test categories:

#### 1. Empty/Zero Values

```go
{
    name:     "empty string",
    input:    "",
    expected: defaultValue,
},
{
    name:     "nil slice",
    input:    nil,
    expected: emptyResult,
},
```

#### 2. Single Element

```go
{
    name:     "single item",
    input:    []string{"only"},
    expected: "only",
},
```

#### 3. Multiple Elements

```go
{
    name:     "multiple items",
    input:    []string{"first", "second", "third"},
    expected: "first,second,third",
},
```

#### 4. Boundary Conditions

```go
{
    name:     "maximum value",
    input:    math.MaxInt64,
    expected: overflow,
},
{
    name:     "special characters",
    input:    "hello@#$%world",
    expected: sanitized,
},
```

---

### Testing Errors

**Bad practice - checking error string:**

```go
// ❌ Fragile - breaks if error message changes
if err == nil || err.Error() != "division by zero" {
    t.Error("wrong error")
}
```

**Good practice - checking error chain:**

```go
// ✅ Semantic error checking
assert.Error(t, err)
assert.ErrorContains(t, err, "division")

// ✅ Check error type (possibly wrapped)
assert.ErrorIs(t, err, ErrDivisionByZero)

// ✅ Check for specific error message
assert.EqualError(t, err, "division by zero")
```

---

### Complete Example

Here's a complete example showing all patterns together:

```go
package calculator_test

import (
    "iter"
    "slices"
    "testing"
    "github.com/go-openapi/testify/v2/assert"
    "github.com/go-openapi/testify/v2/require"
)

type divideTestCase struct {
    name      string
    a, b      float64
    expected  float64
    shouldErr bool
}

func divideTestCases() iter.Seq[divideTestCase] {
    return slices.Values([]divideTestCase{
        {
            name:      "positive numbers",
            a:         10,
            b:         2,
            expected:  5,
            shouldErr: false,
        },
        {
            name:      "negative dividend",
            a:         -10,
            b:         2,
            expected:  -5,
            shouldErr: false,
        },
        {
            name:      "division by zero",
            a:         10,
            b:         0,
            shouldErr: true,
        },
        {
            name:      "zero dividend",
            a:         0,
            b:         5,
            expected:  0,
            shouldErr: false,
        },
    })
}

func TestDivide(t *testing.T) {
    t.Parallel()

    for c := range divideTestCases() {
        t.Run(c.name, func(t *testing.T) {
            t.Parallel()

            result, err := Divide(c.a, c.b)

            if c.shouldErr {
                assert.Error(t, err)
                assert.ErrorIs(t, err, ErrDivisionByZero)
            } else {
                require.NoError(t, err)
                assert.Equal(t, c.expected, result)
            }
        })
    }
}
```

---

## Best Practices Summary

1. **Use the iterator pattern** - `iter.Seq[T]` for all table-driven tests
2. **Separate test data from test logic** - Iterator functions are your test data
3. **Use `t.Parallel()`** - Both outer tests and subtests
4. **Use `t.Helper()`** - In assertion helper functions
5. **Use `require` for preconditions** - Stop test if setup fails
6. **Use `assert` for checks** - See all failures
7. **Test edge cases** - Empty, single, multiple, boundary conditions
8. **Use forward methods** - `assert.New(t)` for tests with many assertions
9. **Use semantic error checking** - `ErrorIs`, `ErrorContains`, not string comparison
10. **Keep tests focused** - One logical concept per test

---

## Examples in This Repository

See real-world usage of these patterns:

- **Iterator Pattern**: `codegen/internal/generator/funcmaps/funcmaps_test.go`
- **Domain Tests**: `codegen/internal/generator/domains/domains_test.go`
- **Assertion Tests**: `internal/assertions/*_test.go`
- **Comprehensive Coverage**: `codegen/internal/scanner/comments-parser/` (all test files)

**Study these to see the patterns in action!**
