# Integration Testing Module

This is a separate Go module dedicated to cross-module integration testing of features that require
external dependencies.

## Purpose

Some testify features are opt-in and require importing external dependencies (YAML support, colorized
output). These features cannot be tested in the main module without breaking the zero-dependency
guarantee. This module provides a place to:

- **Exercise opt-in features** that activate via the `enable/` import pattern (YAML, colors)
- **Run property-based and fuzz testing** using external testing libraries (rapid)

This maintains our zero-dependency goal while enabling thorough testing of the full feature set.

## Structure

```
internal/testintegration/
├── go.mod                    # Separate module with test dependencies
├── go.sum                    # Dependency checksums
├── doc.go                    # Package documentation
├── README.md                 # This file
├── colors/
│   ├── doc.go                # Package documentation
│   └── assertions_test.go    # Tests for colorized assertion output
├── yaml/
│   ├── doc.go                # Package documentation
│   ├── enable.go             # YAML enablement via stubs
│   └── assertions_test.go    # Tests for YAML assertions (YAMLEq)
└── spew/
    ├── doc.go                # Package documentation
    ├── generator.go          # Reflection-based random value generator
    ├── generator_test.go     # Generator validation tests
    ├── edgecases.go          # Hand-crafted edge case generators
    ├── edgecases_test.go     # Edge case focused tests
    ├── dump_test.go          # Main property-based tests (rapid.Check)
    ├── dump_fuzz_test.go     # Go native fuzz tests
    └── options.go            # Generator options (e.g., WithSkipCircularMap)
```

## Dependencies

- **`pgregory.net/rapid`** - Property-based testing library with fuzzing capabilities
- **`go.yaml.in/yaml/v3`** - YAML parsing (for YAML assertion integration tests)
- **`github.com/go-openapi/testify/enable/colors/v2`** - Colorized output activation

## Test Packages

### `colors/` — Colorized Output

Tests that the `enable/colors` import pattern correctly activates ANSI color codes in assertion
failure messages.

- Imports `_ "github.com/go-openapi/testify/enable/colors/v2"` to activate colors
- Forces the `-testify.colorized` and `-testify.colorized.notty` flags via `init()`
- Verifies that assertion output contains ANSI escape sequences (`\x1b`)

### `yaml/` — YAML Assertions

Tests that YAML assertions work when the YAML feature is enabled via the stubs mechanism.

- Calls `yamlstub.EnableYAMLWithUnmarshal(yaml.Unmarshal)` to wire in the real YAML parser
- Exercises `YAMLEq` with matching and non-matching YAML documents

### `spew/` — Property-Based and Fuzz Testing

Uses `pgregory.net/rapid` to perform comprehensive black-box testing of `internal/spew`.

#### Bugs Fixed

1. **Pointer wrapped as interface** — `self = &self` pattern caused infinite loop
2. **Map containing itself** — `m["key"] = m` pattern caused infinite loop

Both are now correctly handled with `<already shown>` markers.

Historical issues addressed:
- **#1828** — Panic on structs with unexported fields
- **#1829** — Time rendering in diffs

## Generator Architecture (spew)

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

# Run all integration tests
go test ./...

# Run feature-specific tests
go test -v ./colors          # Colorized output tests
go test -v ./yaml            # YAML assertion tests
go test -v ./spew            # Property-based spew tests

# Run specific test
go test -v ./spew -run TestSdump

# Run fuzz tests
go test -fuzz=FuzzSdump ./spew -fuzztime=30s

# Run with custom rapid iterations
go test ./spew -rapid.checks=1000000
```

> [!NOTE]
>
> In CI, rapid's value generator generates 10,000 values only (local tests: 100,000)
> and fuzz tests run for 5min

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

To add integration tests for a new feature or internal package:

1. Create a new subdirectory under `internal/testintegration/`
2. Add a `doc.go` with package documentation
3. Add test files exercising the feature

For features that require the `enable/` import pattern, follow the `yaml/` or `colors/` examples.
For property-based or fuzz testing, follow the `spew/` example using `pgregory.net/rapid`.

## Why a Separate Module?

This approach:

- **Isolates test dependencies** — rapid, yaml, and colors dependencies stay out of the main module
- **Maintains zero dependencies** — Main module keeps its zero-dependency guarantee
- **Tests the enable pattern end-to-end** — Verifies that opt-in features work when activated
- **Enables powerful testing** — Property-based and fuzz testing with best-in-class tools
- **Clear separation** — Test infrastructure vs production code
- **Flexible versioning** — Can update test tools independently

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
