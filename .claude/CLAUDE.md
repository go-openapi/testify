# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the go-openapi fork of the testify testing package. The main goal is to remove external dependencies while maintaining a clean, focused API for testing in Go. This fork strips out unnecessary features (mocks, suite) and internalizes dependencies (go-spew, difflib) to ensure `github.com/go-openapi/testify/v2` is the only import needed.

## Key Architecture

### Core Packages
- **assert**: Provides non-fatal test assertions (tests continue after failures)
- **require**: Provides fatal test assertions (tests stop immediately on failure via `FailNow()`)
- Both packages share similar APIs, but `require` wraps `assert` functions to make them fatal

### Code Generation
- The codebase uses code generation extensively via `_codegen/main.go`
- Generated files include:
  - `assert/assertion_format.go` - Format string variants of assertions
  - `assert/assertion_forward.go` - Forwarded assertion methods
  - `require/require.go` - Require variants of all assert functions
  - `require/require_forward.go` - Forwarded require methods

### Dependency Isolation Strategy
- **internal/spew**: Internalized copy of go-spew for pretty-printing values
- **internal/difflib**: Internalized copy of go-difflib for generating diffs
- **assert/yaml**: Stub package that panics by default if YAML assertions are used
- **enable/yaml**: Optional module that activates YAML support via init() when imported

The "enable" pattern allows YAML functionality to be opt-in: import `_ "github.com/go-openapi/testify/v2/enable/yaml"` to activate YAML assertions without forcing a dependency on all users.

## Development Commands

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests in a specific package
go test ./assert
go test ./require

# Run a single test
go test ./assert -run TestEqual

# Run tests with verbose output
go test -v ./...
```

### Code Generation
When modifying assertion functions in `assert/assertions.go`, regenerate derived code:
```bash
# Generate all code
go generate ./...

# This runs the codegen tool which:
# 1. Parses assert/assertions.go for TestingT functions
# 2. Generates format variants (e.g., Equalf from Equal)
# 3. Generates require variants (fatal versions)
# 4. Generates forwarded assertion methods
```

The code generator looks for functions with signature `func(TestingT, ...) bool` in the assert package and creates corresponding variants.

### Build and Verify
```bash
# Tidy dependencies
go mod tidy

# Build code generator
cd _codegen && go build

# Format code
go fmt ./...
```

## Important Constraints

### API Stability
The following assertions are guaranteed stable (used by go-openapi/go-swagger):
- Condition, Contains, Empty, Equal, EqualError, EqualValues, Error, ErrorContains, ErrorIs
- Fail, FailNow, False, Greater, Implements, InDelta, IsType, JSONEq, Len
- Nil, NoError, NotContains, NotEmpty, NotEqual, NotNil, NotPanics, NotZero
- Panics, PanicsWithValue, Subset, True, YAMLEq, Zero

Other APIs may change without notice as the project evolves.

### Zero External Dependencies
Do not add external dependencies to the main module. If new functionality requires a dependency:
1. Consider internalizing it (copy into `internal/` with proper licensing)
2. Or create an "enable" package that users import to activate the feature

### YAML Support Pattern
When using YAML assertions (YAMLEq, YAMLEqf):
- Tests must import: `_ "github.com/go-openapi/testify/v2/enable/yaml"`
- Without this import, YAML assertions will panic with helpful error message
- This pattern keeps gopkg.in/yaml.v3 as an optional dependency

## Module Information

- Module path: `github.com/go-openapi/testify/v2`
- Go version: 1.24.0
- v2.0.0 is retracted (see go.mod)
- License: Apache-2.0 (forked from MIT-licensed stretchr/testify)

## Testing Philosophy

Keep tests simple and focused. The assert package provides detailed failure messages automatically, so test code should be minimal and readable. Use `require` when a test cannot continue meaningfully after a failure, and `assert` when subsequent checks might provide additional context.
