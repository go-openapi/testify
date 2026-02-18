# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the go-openapi fork of the testify testing package. The main goal is to remove external dependencies while maintaining a clean, focused API for testing in Go. This fork strips out unnecessary features (mocks, suite) and internalizes dependencies (go-spew, difflib) to ensure `github.com/go-openapi/testify/v2` is the only import needed.

The project uses sophisticated code generation to:
1. Generate assert and require packages from a single source (`internal/assertions/`)
2. Generate comprehensive tests for all assertion variants
3. Generate domain-organized documentation for a Hugo-based documentation site

## Key Architecture

### Single Source of Truth: `internal/assertions/`

All assertion implementations live in `internal/assertions/`, organized by domain:
- **boolean.go** - True, False
- **collection.go** - Contains, Empty, Len, ElementsMatch, Subset, etc.
- **compare.go** - Greater, Less, comparison assertions
- **equal.go** - Equal, EqualValues, NotEqual, Same, etc.
- **error.go** - Error, NoError, ErrorIs, ErrorAs, etc.
- **file.go** - FileExists, DirExists, FileEmpty, FileNotEmpty
- **http.go** - HTTPSuccess, HTTPError, HTTPStatusCode, etc.
- **json.go** - JSONEq
- **number.go** - InDelta, InEpsilon, Positive, Negative
- **panic.go** - Panics, NotPanics, PanicsWithValue
- **string.go** - Regexp, NotRegexp
- **time.go** - WithinDuration
- **type.go** - IsType, Zero, NotZero, Implements
- **yaml.go** - YAMLEq

**Key principle:** Write assertions once in `internal/assertions/` with comprehensive tests. Everything else is generated.

### Core Packages (Generated)
- **assert**: Provides non-fatal test assertions (tests continue after failures)
  - Generated from `internal/assertions/` by `codegen/`
  - Returns `bool` to indicate success/failure
- **require**: Provides fatal test assertions (tests stop immediately on failure via `FailNow()`)
  - Generated from `internal/assertions/` by `codegen/`
  - Void functions that call `FailNow()` on failure

Both packages are 100% generated and maintain API consistency mechanically.

### Code Generation Architecture

The codebase uses sophisticated code generation via the `codegen/` directory:

**Structure:**
```
codegen/
├── internal/
│   ├── scanner/              # Parses internal/assertions using go/packages and go/types
│   │   ├── comments/         # Doc comment extraction
│   │   ├── comments-parser/  # Domain, examples, and tag parsing
│   │   └── signature/        # Function signature analysis
│   ├── generator/            # Template-based code generation engine
│   │   ├── doc_generator.go  # Documentation generator (domain-organized)
│   │   ├── render.go         # Markdown rendering utilities
│   │   └── templates/        # Go templates for code and docs
│   ├── model/                # Data model for assertions
│   │   └── documentation.go  # Documentation structures
├── main.go                   # CLI orchestration
├── docs/                     # Generated documentation output
└── (generated outputs in assert/ and require/)
```

**Generated code files include:**
- **assert/assertion_assertions.go** - Package-level assertion functions
- **assert/assertion_format.go** - Format string variants (Equalf, Truef, etc.)
- **assert/assertion_forward.go** - Forwarded assertion methods for chaining
- **assert/assertion_*_test.go** - Generated tests for all assert variants
- **require/requirement_assertions.go** - Fatal assertion functions
- **require/requirement_format.go** - Fatal format variants
- **require/requirement_forward.go** - Fatal forwarded methods
- **require/requirement_*_test.go** - Generated tests for all require variants

**Generated documentation files include:**
- **docs/doc-site/api/_index.md** - API index page listing all domains
- **docs/doc-site/api/{domain}.md** - Domain-specific pages (boolean.md, collection.md, etc.)
- Documentation is organized by domain (boolean, string, error, etc.) rather than by package
- Each domain page shows both assert and require variants together

**Each assertion function generates 8 variants:**
1. `assert.Equal(t, ...)` - package-level function
2. `assert.Equalf(t, ..., "msg")` - format variant
3. `a.Equal(...)` - forward method (where `a := assert.New(t)`)
4. `a.Equalf(..., "msg")` - forward format variant
5. `require.Equal(t, ...)` - fatal package-level
6. `require.Equalf(t, ..., "msg")` - fatal format variant
7. `r.Equal(...)` - fatal forward method
8. `r.Equalf(..., "msg")` - fatal forward format variant

With 127 assertion functions, this generates 840 functions automatically.

### Dependency Isolation Strategy
- **internal/spew**: Internalized copy of go-spew for pretty-printing values
- **internal/difflib**: Internalized copy of go-difflib for generating diffs
- **internal/assertions/enable**: Internal stubs that panic by default if YAML/color assertions are used
- **enable/stubs**: Public API for enabling optional features (yaml, colors)
- **enable/yaml**: Optional module that activates YAML support via init() when imported
- **enable/colors**: Optional module that activates colorized output via init() when imported

The "enable" pattern allows optional functionality to be opt-in: import `_ "github.com/go-openapi/testify/v2/enable/yaml"` to activate YAML assertions, or `_ "github.com/go-openapi/testify/v2/enable/colors"` to enable colorized output, without forcing dependencies on all users.

## Development Commands

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests in specific packages
go test ./internal/assertions  # Source of truth with exhaustive tests
go test ./assert               # Generated package tests
go test ./require              # Generated package tests
go test ./codegen/internal/... # Scanner and generator tests

# Run a single test
go test ./internal/assertions -run TestEqual

# Run with coverage
go test -cover ./internal/assertions          # Should be 90%+
go test -cover ./assert                       # Should be ~100%
go test -cover ./require                      # Should be ~100%
go test -cover ./codegen/internal/scanner/... # Scanner tests

# Run tests with verbose output
go test -v ./...
```

### Working with Documentation

```bash
# Generate documentation (included in go generate)
go generate ./...

# Or generate documentation explicitly
cd codegen && go run . -output-packages assert,require -include-doc

# Preview documentation site locally
cd hack/doc-site/hugo
go run gendoc.go
# Visit http://localhost:1313/testify/

# The Hugo site auto-reloads on changes to docs/doc-site/
# To see changes, re-run: go generate ./...
```

### Adding a New Assertion

**The entire workflow:**
1. Add function to appropriate file in `internal/assertions/`
2. Add "Examples:" section to doc comment
3. Add "domain:" tag to doc comment for documentation organization
4. Add tests to corresponding `*_test.go` file
5. Run `go generate ./...`
6. Done - all 8 variants generated with tests + documentation

**Example - Adding a new assertion:**
```go
import (
	"fmt"
	"strings"
)

// In internal/assertions/string.go

// StartsWith asserts that the string starts with the given prefix.
//
// domain: string
//
// Examples:
//
//	success: "hello world", "hello"
//	failure: "hello world", "bye"
func StartsWith(t T, str, prefix string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if !strings.HasPrefix(str, prefix) {
		return Fail(t, fmt.Sprintf("Expected %q to start with %q", str, prefix), msgAndArgs...)
	}
	return true
}
```

**Note on placeholder values in Examples:**
- For complex values that can't be easily represented (pointers, structs, etc.), use `// NOT IMPLEMENTED` marker:
  ```go
  // Examples:
  //   success: &customStruct{Field: "value"}, // NOT IMPLEMENTED
  //   failure: complexType{}, // NOT IMPLEMENTED
  ```
- **Never use `// TODO`** - it triggers false positives in code quality analyzers and project management tools

Then add tests in `internal/assertions/string_test.go` and run `go generate ./...`.

This generates:
- `assert.StartsWith(t, str, prefix)`
- `assert.StartsWithf(t, str, prefix, "msg")`
- `a.StartsWith(str, prefix)` (forward method)
- `a.StartsWithf(str, prefix, "msg")`
- `require.StartsWith(t, str, prefix)`
- `require.StartsWithf(t, str, prefix, "msg")`
- `r.StartsWith(str, prefix)` (forward method)
- `r.StartsWithf(str, prefix, "msg")`
- Tests for all 8 variants
- Documentation entry in `docs/doc-site/api/string.md`

### Code Generation
```bash
# Generate all code and documentation from internal/assertions
go generate ./...

# Or run the generator directly with all outputs
cd codegen && go run . -output-packages assert,require -include-doc

# Run code generation only (skip documentation)
cd codegen && go run . -output-packages assert,require -include-doc=false

# The generator workflow:
# 1. Scans internal/assertions/ for exported functions and types
# 2. Extracts doc comments, "Examples:", and domain tags
# 3. Generates assert/ package with all variants + tests
# 4. Generates require/ package with all variants + tests
# 5. Reorganizes by domain and generates docs/doc-site/ markdown
# 6. Ensures 100% test coverage via example-driven tests
```

### Example-Driven Test Generation

The generator reads "Examples:" sections from doc comments:

```go
// Equal asserts that two objects are equal.
//
// Examples:
//
//	success: 123, 123
//	failure: 123, 456
func Equal(t T, expected, actual any, msgAndArgs ...any) bool {
	// implementation
}
```

From this, it generates tests that verify:
- Success case works correctly
- Failure case works correctly and calls appropriate failure methods
- Format variants work with message parameter
- Forward methods work with chaining

**Test case types:**
- `success: <args>` - Test should pass
- `failure: <args>` - Test should fail
- `panic: <args>` - Test should panic (followed by assertion message on next line)
  `<expected panic message>`

### Documentation Generation

The codegen also generates domain-organized documentation for a Hugo static site:

**Documentation workflow:**
1. Scanner extracts domain tags from doc comments (e.g., `// domain: string`)
2. Scanner collects domain descriptions from `internal/assertions/doc.go`
3. Generator merges documentation from assert and require packages
4. DocGenerator reorganizes by domain instead of package
5. Markdown files generated in `docs/doc-site/api/` (19 domain pages)

**Domain organization:**
- Functions are grouped by domain (boolean, collection, comparison, equality, error, etc.)
- 19 domains total covering all assertion categories
- Each domain page shows assert and require variants together
- Index page lists all domains with descriptions
- Uses Hugo front matter for metadata

**Hugo static site setup:**
Located in `hack/doc-site/hugo/`:
- **hugo.yaml** - Main Hugo configuration
- **gendoc.go** - Development server script
- **themes/hugo-relearn** - Documentation theme
- Mounts generated content from `docs/doc-site/`

**Running the documentation site locally:**
```bash
# First generate the documentation
go generate ./...

# Then start the Hugo dev server
cd hack/doc-site/hugo
go run gendoc.go

# Visit http://localhost:1313/testify/
```

**Domain tagging in source code:**
To assign a function to a domain, add a domain tag in its doc comment:
```go
// Equal asserts that two objects are equal.
//
// domain: equality
//
// Examples:
//
//	success: 123, 123
//	failure: 123, 456
func Equal(t T, expected, actual any, msgAndArgs ...any) bool {
	// implementation
}
```

Add domain descriptions in `internal/assertions/doc.go`:
```go
// domain-description: equality
// Assertions for comparing values for equality, including deep equality,
// value equality, and pointer equality.
```

### Build and Verify
```bash
# Tidy dependencies
go mod tidy

# Build code generator
cd codegen && go build

# Format code
go fmt ./...

# Run all tests
go test ./...

# Check coverage
go test -cover ./internal/assertions
go test -cover ./assert
go test -cover ./require
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

### Testing Strategy: Layered Coverage

**Layer 1: Exhaustive Tests in `internal/assertions/`** (94% coverage)
- Comprehensive table-driven tests using Go 1.23 `iter.Seq` patterns
- Error message content and format validation
- Edge cases, nil handling, type coercion scenarios
- Domain-organized test files mirroring implementation
- Source of truth for assertion correctness

**Layer 2: Generated Smoke Tests in `assert/` and `require/`** (~100% coverage)
- Minimal mechanical tests proving functions exist and work
- Success case: verify correct return value / no FailNow
- Failure case: verify correct return value / FailNow called
- Generated from "Examples:" in doc comments
- No error message testing (already covered in Layer 1)

**Layer 3: Meta Tests for Generator** (future)
- Test that code generation produces correct output
- Verify function signatures, imports, structure
- Optional golden file testing

This layered approach ensures:
- Deep testing where it matters (source implementation)
- Complete coverage of generated forwarding code
- Simple, maintainable test generation
- No duplication of complex test logic

### Iterator Pattern for Table-Driven Tests

**This repository uses a signature testing pattern** based on Go 1.23's `iter.Seq` for all table-driven tests:

```go
import (
	"iter"
	"slices"
	"testing"
)

// Define test case struct
type parseTestExamplesCase struct {
	name     string
	input    string
	expected []model.Test
}

// Create iterator function returning iter.Seq[caseType]
func parseTestExamplesCases() iter.Seq[parseTestExamplesCase] {
	return slices.Values([]parseTestExamplesCase{
		{
			name: "success and failure examples",
			input: `Examples:
  success: 123, 456
  failure: 789, 012`,
			expected: []model.Test{
				{TestedValue: "123, 456", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "789, 012", ExpectedOutcome: model.TestFailure},
			},
		},
		// More cases...
	})
}

// Test function iterates over cases
func TestParseTestExamples(t *testing.T) {
	t.Parallel()

	for c := range parseTestExamplesCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTestExamples(c.input)
			// Assertions...
		})
	}
}
```

**Benefits of this pattern:**
- Clean separation between test data (iterator functions) and test logic (test functions)
- Easy to add new test cases by appending to the slice
- Test case structs can include helper functions or complex setup logic
- Subtests run in parallel automatically with `t.Parallel()`
- Iterator functions can be reused across multiple tests
- Type-safe test case definitions

**When to use this pattern:**
- Any test with 2+ test cases
- Tests that need complex setup or helper functions
- Tests that benefit from parallel execution
- Any table-driven test scenario

**Examples in codebase:**
- `codegen/internal/scanner/comments-parser/examples_test.go`
- `codegen/internal/generator/domains/domains_test.go`
- `internal/assertions/*_test.go` (most assertion tests)

## Architecture Benefits

### Why This Design Wins

**For Contributors:**
- Add assertion in focused, domain-organized file
- Write tests once in single location
- Run `go generate` and get all variants for free
- Clear separation: source vs generated code

**For Maintainers:**
- Mechanical consistency across 608 generated functions
- Template changes affect all functions uniformly
- Easy to add new variants (e.g., generics)
- Single source of truth prevents drift

**For Users:**
- Comprehensive API with 76 assertions
- All expected variants (package, format, forward, require)
- Zero external dependencies
- Drop-in replacement for stretchr/testify

**The Math:**
- 127 assertion functions × 4-8 variants = 840 functions
- Old model: Manually maintain 840 functions across multiple packages
- New model: Write 127 functions once, generate the rest
- Result: 85% reduction in manual code maintenance

### Technical Innovations

**Go AST/Types Integration:**
- Scanner uses `go/packages` and `go/types` for semantic analysis
- Position-based lookup bridges AST and type information
- Import alias resolution for accurate code generation
- Handles complex Go constructs (generics, interfaces, variadic args)

**Scanner Architecture (Refactored):**
- Modular sub-packages for different extraction concerns:
  - `comments/` - Doc comment extraction with integration tests
  - `comments-parser/` - Domain tags, examples, and metadata parsing
  - `signature/` - Function signature analysis
- Comprehensive test coverage (~1,435 lines across 4 test files)
- Extracts domain tags, domain descriptions, examples, and other metadata
- Integration tests validate end-to-end comment extraction

**Example-Driven Testing:**
- "Examples:" sections in doc comments drive test generation
- success/failure/panic cases extracted automatically
- Tests generated for all 8 variants per function
- Achieves 100% coverage with minimal test complexity

**Template Architecture:**
- Separate templates for code (assert/require) and documentation
- Code templates handle return values vs void functions
- Doc templates (doc_index.md.gotmpl, doc_page.md.gotmpl) generate Hugo markdown
- Mock selection based on FailNow requirements
- Consistent formatting and structure across all output

**Documentation Generator:**
- Merges package-based documentation into domain-based organization
- Reorganizes functions from assert/require packages by domain
- Generates Hugo-compatible markdown with front matter
- Creates navigable API reference organized by concern (boolean, string, error, etc.)

## Documentation Site

A Hugo-based documentation site is automatically generated from the source code:
- **Generated content**: `docs/doc-site/api/` - Domain-organized markdown files (19 domain pages)
- **Hugo site**: `hack/doc-site/hugo/` - Hugo configuration and theme (temporary location)
- **Target URL**: `https://go-openapi.github.io/testify/`

**Site workflow:**
1. `codegen` generates markdown in `docs/doc-site/api/`
2. Hugo builds static site from the generated markdown
3. Hugo mounts the `docs/doc-site/` directory as content

**Generated domain pages in `docs/doc-site/api/`:**
- `_index.md` - API index page
- `boolean.md`, `collection.md`, `common.md`, `comparison.md`, `condition.md`
- `equality.md`, `error.md`, `file.md`, `http.md`, `json.md`
- `number.md`, `ordering.md`, `panic.md`, `string.md`, `testing.md`
- `time.md`, `type.md`, `yaml.md`

**Hugo configuration in `hack/doc-site/hugo/`:**
```
hack/doc-site/hugo/           # Note: Temporary location
├── hugo.yaml                 # Main Hugo configuration
├── testify.yaml              # Generated config with version info
├── testify.yaml.template     # Template for testify.yaml
├── gendoc.go                 # Development server launcher
├── README.md, TODO.md        # Documentation and planning
├── themes/
│   └── hugo-relearn/         # Documentation theme
├── layouts/                  # Custom layout overrides
└── content/                  # Mounted from docs/doc-site/
```

## Example Coverage Status

Most assertion functions now have "Examples:" sections in their doc comments. The generator extracts these to create both tests and documentation examples.

**Coverage notes:**
- Basic assertions (Equal, Error, Contains, Len, True, False) have complete examples
- Complex values that can't be easily represented should use `// NOT IMPLEMENTED` as placeholder marker
  - **Never use `// TODO`** - it triggers false positives in code quality analyzers
  - Use: `success: &structValue{}, // NOT IMPLEMENTED`
  - Not: `success: &structValue{}, // TODO`
- All new assertions should include Examples and domain tags before merging
- Domain tags organize assertions into logical groups for documentation

For the complete guide on adding examples, see `docs/MAINTAINERS.md` section "Maintaining Generated Code".

## Scanner Testing and Architecture

The scanner package has comprehensive test coverage to ensure reliable code generation:

**Test files (~1,435 lines total):**
- `comments-parser/examples_test.go` (350 lines) - Tests for example extraction from doc comments
- `comments-parser/matchers_test.go` (171 lines) - Tests for pattern matching in comments
- `comments-parser/tags_test.go` (407 lines) - Tests for domain and metadata tag extraction
- `comments/extractor_integration_test.go` (507 lines) - End-to-end comment extraction tests

**Scanner responsibilities:**
1. Parse Go packages using `go/packages` and `go/types`
2. Extract doc comments and parse structured metadata
3. Identify assertion functions and their signatures
4. Extract domain tags for documentation organization
5. Parse "Examples:" sections for test generation
6. Collect domain descriptions from special comment tags
7. Build the model used by code and doc generators

**Testing strategy:**
- Unit tests for individual parsers (examples, tags, matchers)
- Integration tests validate complete extraction pipeline
- Tests use real Go code samples from `testdata/`
- Ensures generated code and docs stay in sync with source
