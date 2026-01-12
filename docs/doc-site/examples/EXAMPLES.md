---
title: "Examples"
description: "Practical examples for using testify v2"
weight: 1
---

{{% notice primary "TL;DR" "meteor" %}}
> If you've already used `github.com/stretchr/testify`, adopting v2 will be straightforward.
{{% /notice %}}

## Quick Start

The simplest way to get started with testify is using the `assert` package:

```go
import (
    "testing"
    "github.com/go-openapi/testify/v2/assert"
)

func TestCalculator(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result)
}
```

## assert vs require

**Use `assert`** when you want tests to continue after a failure:

```go
func TestUser(t *testing.T) {
    user := GetUser(123)

    // All three checks run, even if the first fails
    assert.NotNil(t, user)
    assert.Equal(t, "Alice", user.Name)
    assert.Equal(t, 25, user.Age)
}
```

**Use `require`** when subsequent checks depend on earlier ones:

```go
import "github.com/go-openapi/testify/v2/require"

func TestUser(t *testing.T) {
    user := GetUser(123)

    // Stop immediately if user is nil (prevents panic on next line)
    require.NotNil(t, user)

    // Only runs if user is not nil
    assert.Equal(t, "Alice", user.Name)
}
```

**Rule of thumb:** Use `require` for preconditions, `assert` for actual checks.

---

## Common Assertions

### Equality

```go
func TestEquality(t *testing.T) {
    // Basic equality
    assert.Equal(t, 42, actualValue)

    // Deep equality for slices, maps, structs
    assert.Equal(t, []int{1, 2, 3}, result)

    // Check inequality
    assert.NotEqual(t, 0, result)

    // Type-converting equality (123 == int64(123))
    assert.EqualValues(t, 123, int64(123))
}
```

### Collections

```go
func TestCollections(t *testing.T) {
    list := []string{"apple", "banana", "cherry"}

    // Check if collection contains an element
    assert.Contains(t, list, "banana")
    assert.NotContains(t, list, "orange")

    // Check length
    assert.Len(t, list, 3)

    // Check if empty
    assert.NotEmpty(t, list)
    assert.Empty(t, []string{})

    // Check if all elements match (order doesn't matter)
    assert.ElementsMatch(t, []int{3, 1, 2}, []int{1, 2, 3})
}
```

### Errors

```go
func TestErrors(t *testing.T) {
    // Check if function returns an error
    err := DoSomething()
    assert.Error(t, err)

    // Check if function succeeds
    err = DoSomethingElse()
    assert.NoError(t, err)

    // Check specific error message
    err = Divide(10, 0)
    assert.EqualError(t, err, "division by zero")

    // Check if error contains text
    assert.ErrorContains(t, err, "division")

    // Check error type with errors.Is
    assert.ErrorIs(t, err, ErrDivisionByZero)
}
```

### Nil Checks

```go
func TestNil(t *testing.T) {
    var ptr *User

    assert.Nil(t, ptr)

    user := &User{Name: "Alice"}
    assert.NotNil(t, user)
}
```

### Boolean and Comparisons

```go
func TestBooleans(t *testing.T) {
    assert.True(t, isValid)
    assert.False(t, hasErrors)

    // Numeric comparisons
    assert.Greater(t, 10, 5)
    assert.GreaterOrEqual(t, 10, 10)
    assert.Less(t, 5, 10)
    assert.LessOrEqual(t, 5, 5)
}
```

---

## Assertion Variants

Testify provides multiple ways to call assertions:

### 1. Package-Level Functions

```go
func TestPackageLevel(t *testing.T) {
    assert.Equal(t, 42, result)
    require.NotNil(t, user)
}
```

### 2. Formatted Variants (Custom Messages)

```go
func TestFormatted(t *testing.T) {
    // Add custom failure message with formatting
    assert.Equalf(t, 42, result, "expected answer to be %d", 42)
    require.NotNilf(t, user, "user %d should exist", userID)
}
```

### 3. Forward Methods (Cleaner Syntax)

```go
func TestForward(t *testing.T) {
    a := assert.New(t)
    r := require.New(t)

    // No need to pass 't' each time
    a.Equal(42, result)
    a.NotEmpty(list)

    r.NotNil(user)
    r.NoError(err)
}
```

### 4. Forward Methods with Formatting

```go
func TestForwardFormatted(t *testing.T) {
    a := assert.New(t)

    a.Equalf(42, result, "expected answer to be %d", 42)
    a.Lenf(list, 3, "expected %d items", 3)
}
```

**Recommendation:** Use package-level functions for simple tests, forward methods for tests with many assertions.

---

## Table-Driven Tests

The idiomatic Go pattern for testing multiple cases should be:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed signs", -2, 3, 1},
        {"with zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

With forward methods for cleaner syntax:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a := assert.New(t)

            result := Add(tt.a, tt.b)
            a.Equal(tt.expected, result)
            a.Greater(result, tt.a)
        })
    }
}
```

---

## Real-World Examples

### Testing HTTP Handlers

```go
import (
    "net/http"
    "net/http/httptest"
    "github.com/go-openapi/testify/v2/assert"
    "github.com/go-openapi/testify/v2/require"
)

func TestUserHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/123", nil)
    w := httptest.NewRecorder()

    handler := NewUserHandler()
    handler.ServeHTTP(w, req)

    require.Equal(t, http.StatusOK, w.Code)

    // Check response body contains expected data
    body := w.Body.String()
    assert.Contains(t, body, `"name":"Alice"`)
    assert.Contains(t, body, `"id":123`)
}
```

### Testing JSON

```go
func TestJSONResponse(t *testing.T) {
    expected := `{"name":"Alice","age":25}`
    actual := `{"age":25,"name":"Alice"}`

    // JSONEq compares JSON semantically (ignores key order, whitespace)
    assert.JSONEq(t, expected, actual)
}
```

### Testing with Subtests

```go
func TestUserOperations(t *testing.T) {
    user := &User{ID: 123, Name: "Alice"}

    t.Run("creation", func(t *testing.T) {
        assert.NotNil(t, user)
        assert.Equal(t, 123, user.ID)
    })

    t.Run("update", func(t *testing.T) {
        user.Name = "Bob"
        assert.Equal(t, "Bob", user.Name)
    })

    t.Run("deletion", func(t *testing.T) {
        err := DeleteUser(user.ID)
        assert.NoError(t, err)
    })
}
```

### Testing Panics

```go
func TestPanics(t *testing.T) {
    // Function should panic
    assert.Panics(t, func() {
        Divide(10, 0)
    })

    // Function should NOT panic
    assert.NotPanics(t, func() {
        Divide(10, 2)
    })

    // Function should panic with specific value
    assert.PanicsWithValue(t, "division by zero", func() {
        Divide(10, 0)
    })
}
```

---

## Advanced Patterns

### Setup and Teardown

```go
func TestWithSetup(t *testing.T) {
    // Setup
    db := setupTestDatabase(t)
    defer db.Close() // Teardown

    // Test
    user := &User{Name: "Alice"}
    err := db.Save(user)
    require.NoError(t, err)

    // Verify
    loaded, err := db.Find(user.ID)
    require.NoError(t, err)
    assert.Equal(t, "Alice", loaded.Name)
}
```

### Helper Functions

```go
func assertUserValid(t *testing.T, user *User) {
    t.Helper() // Mark as helper for better error messages

    assert.NotNil(t, user)
    assert.NotEmpty(t, user.Name)
    assert.Greater(t, user.Age, 0)
}

func TestUsers(t *testing.T) {
    user := GetUser(123)
    assertUserValid(t, user) // Failures point to this line, not inside helper
}
```

### Combining Multiple Assertions

```go
func TestUserCompleteness(t *testing.T) {
    a := assert.New(t)
    user := GetUser(123)

    // Chain multiple checks cleanly
    a.NotNil(user)
    a.NotEmpty(user.Name)
    a.NotEmpty(user.Email)
    a.Greater(user.Age, 0)
    a.True(user.Active)
}
```

---

## YAML Support (Optional)

YAML assertions require explicit opt-in:

```go
import (
    "testing"
    "github.com/go-openapi/testify/v2/assert"
    _ "github.com/go-openapi/testify/v2/enable/yaml" // Enable YAML support
)

func TestYAML(t *testing.T) {
    expected := `
name: Alice
age: 25
`
    actual := `
age: 25
name: Alice
`

    // YAMLEq compares YAML semantically
    assert.YAMLEq(t, expected, actual)
}
```

**Note:** Without the `enable/yaml` import, YAML assertions will panic with a helpful message.

---

## Colorized Output (Optional)

Testify can colorize test failure output for better readability. This is an opt-in feature.

### Enabling Colors

```go
import (
    "testing"
    "github.com/go-openapi/testify/v2/assert"
    _ "github.com/go-openapi/testify/enable/colors/v2" // Enable colorized output
)

func TestExample(t *testing.T) {
    assert.Equal(t, "expected", "actual") // Failure will be colorized
}
```

### Activation

Colors are activated via command line flag or environment variable:

```bash
# Via flag
go test -v -testify.colorized ./...

# Via environment variable
TESTIFY_COLORIZED=true go test -v ./...
```

### Themes

Two themes are available for different terminal backgrounds:

```bash
# Dark theme (default) - bright colors for dark terminals
go test -v -testify.colorized ./...

# Light theme - normal colors for light terminals
go test -v -testify.colorized -testify.theme=light ./...

# Or via environment
TESTIFY_COLORIZED=true TESTIFY_THEME=light go test -v ./...
```

### CI Environments

By default, colorization is disabled when output is not a terminal. To force colors in CI environments that support ANSI codes:

```bash
TESTIFY_COLORIZED=true TESTIFY_COLORIZED_NOTTY=true go test -v ./...
```

### What Gets Colorized

- **Expected values** in assertion failures (green)
- **Actual values** in assertion failures (red)
- **Diff output**:
  - Deleted lines (red)
  - Inserted lines (yellow)
  - Context lines (green)

**Note:** Without the `enable/colors` import, output remains uncolored (no panic, just no colors).

---

## Best Practices

1. **Use `require` for preconditions** - Stop test immediately if setup fails
2. **Use `assert` for actual checks** - See all failures in one test run
3. **Add custom messages for complex checks** - Use formatted variants when assertion failure needs context
4. **Prefer table-driven tests** - Test multiple cases systematically
5. **Use forward methods for many assertions** - Reduces repetition in long tests
6. **Keep tests focused** - One logical concept per test function
7. **Use subtests for related scenarios** - Group related checks with `t.Run()`
8. **Mark helpers with `t.Helper()`** - Get better error locations

---

## Migration from stdlib testing

**Before (stdlib):**
```go
func TestOld(t *testing.T) {
    result := Calculate(5)
    if result != 10 {
        t.Errorf("Expected 10, got %d", result)
    }
    if len(items) == 0 {
        t.Error("Expected non-empty items")
    }
}
```

**After (testify):**
```go
func TestNew(t *testing.T) {
    a := assert.New(t)

    result := Calculate(5)
    a.Equal(10, result)
    a.NotEmpty(items)
}
```

**Benefits:**
- More readable - assertions read like English
- Better error messages - shows expected vs actual automatically
- Less boilerplate - no manual formatting
- More assertions - Contains, ElementsMatch, JSONEq, etc.
