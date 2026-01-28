---
title: "Examples"
description: "Practical examples for using testify v2"
weight: 2
---

{{% notice primary "TL;DR" "meteor" %}}
> If you've already used `github.com/stretchr/testify`, adopting v2 will be straightforward.
{{% /notice %}}

More examples to showcase generic assertions specifically may be found [here](./GENERICS.md).

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestNil(t *testing.T) {
	var ptr *User

	assert.Nil(t, ptr)

	user := &User{Name: "Alice"}
	assert.NotNil(t, user)
}
```

### Boolean and Comparisons

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestPackageLevel(t *testing.T) {
	assert.Equal(t, 42, result)
	require.NotNil(t, user)
}
```

### 2. Formatted Variants (Custom Messages)

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestFormatted(t *testing.T) {
	// Add custom failure message with formatting
	assert.Equalf(t, 42, result, "expected answer to be %d", 42)
	require.NotNilf(t, user, "user %d should exist", userID)
}
```

### 3. Forward Methods (Cleaner Syntax)

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestAdd(t *testing.T) {
	tests := slices.Values([]struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"with zero", 0, 5, 5},
	})

	for tt := range tests {
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
    tests := slices.Values([]struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
    })

    for tt := range tests {
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
	"testing"

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestJSONResponse(t *testing.T) {
	expected := `{"name":"Alice","age":25}`
	actual := `{"age":25,"name":"Alice"}`

	// JSONEq compares JSON semantically (ignores key order, whitespace)
	assert.JSONEq(t, expected, actual)
}
```

### Testing with Subtests

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestWithSetup(t *testing.T) {
	// Setup
	db := setupTestDatabase(t)
	t.Cleanup(func() {
		db.Close() // Teardown
	})

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestUsers(t *testing.T) {
	user := GetUser(123)
	assertUserValid(t, user) // Failures point to this line, not inside helper
}

func assertUserValid(t *testing.T, user *User) {
	t.Helper() // Mark as helper for better error messages

	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Name)
	assert.Greater(t, user.Age, 0)
}
```

### Combining Multiple Assertions

```go
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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

### Asynchronous Testing

Testify provides three assertions for testing asynchronous code: `Eventually`, `Never`, and `EventuallyWith`.

{{% notice warning %}}
> Asynchronous testing may sometimes be unavoidable. It should be avoided whenever possible.
>
> Async tests (with timeouts, ticks etc) may easily become flaky under heavy concurrence on small CI runners.
>
> When you've control over the code you test, always prefer sync tests, possibly with well-designed mocks.
{{% /notice %}}

#### Eventually: Wait for a Condition to Become True

Use `Eventually` when testing code that updates state asynchronously (background goroutines, event loops, caches).

```go
import (
	"sync"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func TestBackgroundProcessor(t *testing.T) {
	// Simulate a background processor that updates state
	var processed bool
	var mu sync.Mutex

	go func() {
		time.Sleep(50 * time.Millisecond)
		mu.Lock()
		processed = true
		mu.Unlock()
	}()

	// Wait up to 200ms for the background task to complete,
	// checking every 10ms
	assert.Eventually(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		return processed
	}, 200*time.Millisecond, 10*time.Millisecond,
		"background processor should have completed")
}

// Real-world example: Testing cache warming
func TestCacheWarming(t *testing.T) {
	cache := NewCache()
	cache.StartWarmup() // Populates cache in background

	// Verify cache becomes ready within 5 seconds
	assert.Eventually(t, func() bool {
		return cache.IsReady() && cache.Size() > 0
	}, 5*time.Second, 100*time.Millisecond,
		"cache should warm up and contain entries")
}
```

#### Never: Ensure a Condition Remains False

Use `Never` to verify that something undesirable never happens during a time window (no data corruption, no invalid state).

```go
import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func TestNoDataCorruption(t *testing.T) {
	var counter atomic.Int32
	stopChan := make(chan struct{})
	defer close(stopChan)

	// Start multiple goroutines incrementing safely
	for i := 0; i < 10; i++ {
		go func() {
			ticker := time.NewTicker(5 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-stopChan:
					return
				case <-ticker.C:
					counter.Add(1)
				}
			}
		}()
	}

	// Verify counter never goes negative (indicating corruption)
	assert.Never(t, func() bool {
		return counter.Load() < 0
	}, 500*time.Millisecond, 20*time.Millisecond,
		"counter should never go negative")
}

// Real-world example: Testing rate limiter doesn't exceed threshold
func TestRateLimiter(t *testing.T) {
	limiter := NewRateLimiter(100) // 100 requests per second max
	stopChan := make(chan struct{})
	defer close(stopChan)

	// Hammer the limiter with requests
	for i := 0; i < 50; i++ {
		go func() {
			ticker := time.NewTicker(1 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-stopChan:
					return
				case <-ticker.C:
					limiter.Allow()
				}
			}
		}()
	}

	// Verify we never exceed the rate limit over 2 seconds
	assert.Never(t, func() bool {
		return limiter.CurrentRate() > 120 // 20% tolerance
	}, 2*time.Second, 50*time.Millisecond,
		"rate limiter should never exceed threshold")
}
```

#### EventuallyWith: Complex Async Assertions

Use `EventuallyWith` when you need multiple assertions to pass together.
The `CollectT` parameter lets you make regular assertions.

```go
import (
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func TestAPIEventualConsistency(t *testing.T) {
	// Simulate an eventually-consistent API
	api := NewEventuallyConsistentAPI()
	api.CreateUser("alice", "alice@example.com")

	// Wait for the user to be fully replicated across all shards
	// All conditions must pass in the same tick
	assert.EventuallyWith(t, func(c *assert.CollectT) {
		user, err := api.GetUser("alice")

		// All these assertions must pass together
		assert.NoError(c, err, "user should be retrievable")
		assert.NotNil(c, user, "user should exist")
		assert.EqualT(c, "alice@example.com", user.Email, "email should match")
		assert.True(c, user.Replicated, "user should be replicated")
		assert.GreaterOrEqual(c, user.ReplicaCount, 3, "should be on at least 3 replicas")
	}, 10*time.Second, 500*time.Millisecond,
		"user should be eventually consistent across all replicas")
}

// Real-world example: Testing distributed cache sync
func TestDistributedCacheSync(t *testing.T) {
	primary := NewCacheNode("primary")
	replica1 := NewCacheNode("replica1")
	replica2 := NewCacheNode("replica2")

	// Connect nodes for replication
	primary.AddReplica(replica1)
	primary.AddReplica(replica2)

	// Write to primary
	primary.Set("key", "value", 5*time.Minute)

	// Verify value propagates to all replicas with correct TTL
	assert.EventuallyWith(t, func(c *assert.CollectT) {
		val1, ttl1, ok1 := replica1.Get("key")
		val2, ttl2, ok2 := replica2.Get("key")

		// All replicas must have the value
		assert.True(c, ok1, "replica1 should have the key")
		assert.True(c, ok2, "replica2 should have the key")

		// Values must match
		assert.EqualT(c, "value", val1, "replica1 value should match")
		assert.EqualT(c, "value", val2, "replica2 value should match")

		// TTL should be approximately the same (within 1 second)
		assert.InDelta(c, 5*time.Minute, ttl1, float64(time.Second),
			"replica1 TTL should be close to original")
		assert.InDelta(c, 5*time.Minute, ttl2, float64(time.Second),
			"replica2 TTL should be close to original")
	}, 5*time.Second, 100*time.Millisecond,
		"cache value should replicate to all nodes with correct TTL")
}

// Advanced: Using require in EventuallyWith to fail fast
func TestEventuallyWithRequire(t *testing.T) {
	api := NewAPI()

	assert.EventuallyWith(t, func(c *assert.CollectT) {
		resp, err := api.HealthCheck()

		// Use require to stop checking this tick if request fails
		// This prevents nil pointer panics on subsequent assertions
		assert.NoError(c, err, "health check should not error")
		if err != nil {
			return // Skip remaining checks this tick
		}

		// Now safe to check response fields
		assert.EqualT(c, "healthy", resp.Status)
		assert.Greater(c, resp.Uptime, 0)
		assert.NotEmpty(c, resp.Version)
	}, 30*time.Second, 1*time.Second,
		"API should become healthy")
}
```

**Key differences:**
- **Eventually**: Simple boolean condition, use for single checks
- **Never**: Opposite of Eventually, verifies condition stays false
- **EventuallyWith**: Complex checks with multiple assertions, use when you need detailed failure messages

**Best practices:**
1. Choose appropriate timeouts: long enough for async operations, short enough for fast test feedback
2. Choose appropriate tick intervals: frequent enough to catch state changes, infrequent enough to avoid overhead
3. Use `EventuallyWith` when you need to understand *which* assertion failed
4. Use `Eventually` for simple boolean conditions (faster, simpler)
5. Use `Never` to verify invariants over time (no race conditions, no invalid state)

---

## YAML Support (Optional)

YAML assertions require explicit opt-in:

```go
import (
	"testing"

	_ "github.com/go-openapi/testify/enable/yaml/v2" // Enable YAML support
	"github.com/go-openapi/testify/v2/assert"
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

	_ "github.com/go-openapi/testify/enable/colors/v2" // Enable colorized output
	"github.com/go-openapi/testify/v2/assert"
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

See [screenshot](./MIGRATION.md#optional-enable-colorized-output).

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
import "testing"

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
import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

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

---

## See Also

- [Tutorial](./TUTORIAL.md) - Comprehensive guide to writing great tests with testify patterns
- [Usage Guide](./USAGE.md) - API conventions, naming patterns, and how to navigate the documentation
- [Generics Guide](./GENERICS.md) - Type-safe assertions for better compile-time checking
- [Migration Guide](./MIGRATION.md) - Migrating from stretchr/testify v1 to this fork
- [API Reference](../api/_index.md) - Complete assertion catalog organized by domain

