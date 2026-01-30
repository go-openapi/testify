# Integration Testing Module

This is a separate Go module dedicated to property-based and fuzz testing of internal packages.

## Purpose

This module uses external testing libraries (like `rapid`) to perform comprehensive black-box testing
without polluting the main module's dependency tree.

This maintains our zero-dependency goal while enabling powerful testing techniques.

## Structure

```
internal/testintegration/
├── go.mod                    # Separate module with test dependencies
├── go.sum                    # Dependency checksums
├── README.md                 # This file
└── spew/
    ├── doc.go                # Package documentation
    ├── generator.go          # Reflection-based random value generator
    ├── generator_test.go     # Generator validation tests
    ├── edgecases.go          # Hand-crafted edge case generators
    ├── edgecases_test.go     # Edge case focused tests
    ├── dump_test.go          # Main property-based tests (rapid.Check)
    ├── dump_fuzz_test.go     # Go native fuzz tests
    └── testdata/             # Fuzz corpus and rapid failure files
```

## Dependencies

- **pgregory.net/rapid** - Property-based testing library with fuzzing capabilities

## Bugs Fixed in spew

This test suite helped identify and validate fixes for the following issues:

### Circular Reference Hangs

1. **Pointer wrapped as interface** - `self = &self` pattern caused infinite loop
2. **Map containing itself** - `m["key"] = m` pattern caused infinite loop

Both are now correctly handled with `<already shown>` markers.

### Historical Issues Addressed

- **#1828** - Panic on structs with unexported fields
- **#1829** - Time rendering in diffs

## Generator Architecture

### Two-Layer Generator System

The test suite uses two complementary generators:

#### 1. Reflection-Based Generator (`generator.go`)

Generates arbitrary Go values using reflection:

- All primitive types (int, string, bool, float, complex, etc.)
- Container types (slice, map, array, struct)
- Pointers with cyclical reference tracking
- Channels and functions
- Special types (sync.Mutex, atomic.Value, etc.)

**Limitations:**
- No generic types
- No type declarations (all types are anonymous)
- No unexported struct fields (reflect limitation)
- No embedded fields with methods

**Example**

Sample random structures generated.

```go
(struct { Ƃ꘨㝛٠ []struct { string; *chan sync.RWMutex; ℙB୧ float64; complex64 }; Ḥ5᠐ string; E߉🯹J sync.Mutex; Ⱆ᮱G𗚟 interface {}; func(int, string) *string }) {
         Ƃ꘨㝛٠: ([]struct { string; *chan sync.RWMutex; ℙB୧ float64; complex64 }) {
         },
         Ḥ5᠐: (string) (len=4) "~௹",
         E߉🯹J: (sync.Mutex) {
          _: (sync.noCopy) {
          },
          mu: (sync.Mutex) {
           state: (int32) 0,
           sema: (uint32) 0
          }
         },
         Ⱆ᮱G𗚟: (interface {}) <nil>,
         P𤚁: (func(int, string) *string) 0x689760
        }
```

```go
(map[[4]*complex128][]*complex128) (len=1) {
         ([4]*complex128) (len=4 cap=4) {
          (*complex128)(0xc005201cd0)((-2.8853058180580424e-129-1.6336761511385133e-16i)),
          (*complex128)(0xc005201cd0)((-2.8853058180580424e-129-1.6336761511385133e-16i)),
          (*complex128)(0xc005201cd0)((-2.8853058180580424e-129-1.6336761511385133e-16i)),
          (*complex128)(0xc005201cd0)((-2.8853058180580424e-129-1.6336761511385133e-16i))
         }: ([]*complex128) (len=1 cap=1) {
          (*complex128)(0xc005201cd0)((-2.8853058180580424e-129-1.6336761511385133e-16i))
         }
        }
```

#### 2. Edge Case Generator (`edgecases.go`)

Hand-crafted generators for known problematic patterns:

- Structs with unexported fields
- Circular references (various patterns)
- Nil interfaces with typed nil values
- Deep nesting
- Maps with interface keys
- Pointer chains
- time.Time values
- Multi-level pointer indirection
- Pointer-to-interface chains

### Generator Options

The generators support options to customize behavior:

```go
// Skip circular map cases (needed for fuzz testing)
Generator(WithSkipCircularMap())
```

## Running Tests

```bash
cd internal/testintegration

# Run all tests (100,000 rapid checks by default)
go test ./...

# Run with verbose output
go test -v ./spew

# Run specific test
go test -v ./spew -run TestSdump

# Run fuzz tests
go test -fuzz=FuzzSdump ./spew -fuzztime=30s

# Run with custom rapid iterations
go test ./spew -rapid.checks=1000000
```

## Test Types

### Property-Based Tests (`dump_test.go`)

Uses `rapid.Check` with 100,000 iterations to verify `spew.Sdump` never panics or hangs:

```go
rapid.Check(t, NoPanicProp(t.Context(), Generator()))
```

The `NoPanicProp` function:
- Draws a random value from the generator
- Runs `spew.Sdump` with a 1-second timeout
- Fails on panic or timeout (hang detection)

### Edge Case Tests (`edgecases_test.go`)

Focused testing of known problematic patterns using the edge case generator.

### Fuzz Tests (`dump_fuzz_test.go`)

Go native fuzzing integrated with rapid:

```go
import (
	"testing"

	"pgregory.net/rapid"
)

func FuzzSdump(f *testing.F) {
	prop := NoPanicProp(f.Context(), Generator(WithSkipCircularMap()))
	f.Fuzz(rapid.MakeFuzz(prop))
}
```

## Known Limitations and Workarounds

### Circular Maps in Fuzz Tests

**Issue:** Go's standard library `fmt` package cannot handle circular references with maps. When rapid's fuzz integration logs drawn values using `fmt`, it causes a stack overflow before `spew.Sdump` is even called.

**Workaround:** Fuzz tests use `WithSkipCircularMap()` to exclude the self-referencing map case. This case is still covered by `rapid.Check` tests which don't trigger the logging issue.

**Root cause:** This is a limitation in Go's `fmt` package, not in rapid or spew.

```go
// This pattern causes fmt to stack overflow:
m := map[string]any{"key": "value"}
m["self"] = m
fmt.Printf("%v", m)  // stack overflow
```

### Generator Limitations

Values that cannot be generated via reflection:
- Structs with unexported fields (use edge case generator)
- Named types with methods on embedded fields
- Generic types
- C types for CGO

## Adding New Tests

To add fuzz tests for other internal packages:

1. Create a new subdirectory under `internal/testintegration/`
2. Add test files with generators specific to that package
3. Use `NoPanicProp` pattern for hang detection

Example structure:

```go
package mypackage

import (
	"context"
	"testing"
	"time"

	"pgregory.net/rapid"
)

func TestMyFunction(t *testing.T) {
	rapid.Check(t, func(rt *rapid.T) {
		input := myGenerator().Draw(rt, "input")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		done := make(chan struct{})
		go func() {
			defer close(done)
			_ = mypackage.MyFunction(input)
		}()

		select {
		case <-done:
			// success
		case <-ctx.Done():
			rt.Fatal("function timed out")
		}
	})
}
```

## Why a Separate Module?

This approach:

- **Isolates test dependencies** - rapid is only needed for integration testing
- **Maintains zero dependencies** - Main module stays clean
- **Enables powerful testing** - Use best-in-class testing tools
- **Clear separation** - Test infrastructure vs production code
- **Flexible versioning** - Can update test tools independently

## Rapid Quick Reference

```go
// Basic generators
rapid.Int()                          // any int
rapid.IntRange(0, 100)               // 0 to 100
rapid.String()                       // any string
rapid.Bool()                         // true or false
rapid.SliceOf(rapid.Int())           // []int
rapid.MapOf(rapid.String(), rapid.Int()) // map[string]int

// Combinators
rapid.OneOf(gen1, gen2, gen3)        // Choose one generator
rapid.Just(value)                    // Constant value
rapid.Deferred(func() *Generator)    // Recursive definitions

// Drawing values
value := rapid.Int().Draw(t, "label")

// Custom generators
rapid.Custom(func(t *rapid.T) MyType {
    return MyType{
        Field: rapid.String().Draw(t, "field"),
    }
})
```

## Further Reading

- [rapid documentation](https://pkg.go.dev/pgregory.net/rapid)
- [Property-based testing introduction](https://hypothesis.works/articles/what-is-property-based-testing/)
- [Fuzzing in Go](https://go.dev/doc/security/fuzz/)
