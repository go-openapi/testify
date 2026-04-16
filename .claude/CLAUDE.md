# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Zero-dependency fork of [stretchr/testify](https://github.com/stretchr/testify) used across
all go-openapi repositories. The fork strips external dependencies (go-spew, difflib are
internalized), removes mock/suite packages, and uses extensive code generation to maintain
consistency across assertion variants.

**Key difference:** `testifylint` does **not** work with this fork — it only recognizes
`stretchr/testify` imports.

### Mono-repo structure

This is a mono-repo with multiple Go modules tied together by `go.work`.

| Module | Purpose |
|--------|---------|
| `.` (root) | Main module `github.com/go-openapi/testify/v2` — package-level doc, `go:generate` directive |
| `codegen/` | Code generator: scans `internal/assertions/`, generates `assert/`, `require/`, and docs |
| `enable/yaml` | Optional module — blank-import to activate YAML assertions |
| `enable/colors` | Optional module — blank-import to activate colorized test output |
| `hack/migrate-testify` | Migration tool for converting from `stretchr/testify` |
| `internal/testintegration` | Integration tests |

### Key packages (within root module)

| Package | Contents |
|---------|----------|
| `assert/` | Non-fatal test assertions (generated from `internal/assertions/`) |
| `require/` | Fatal test assertions via `FailNow()` (generated from `internal/assertions/`) |
| `internal/assertions/` | **Single source of truth** — all assertion implementations, organized by domain |
| `internal/spew/` | Internalized go-spew for pretty-printing values |
| `internal/difflib/` | Internalized go-difflib for generating diffs |
| `internal/fdleak/` | File descriptor leak detection |
| `internal/leak/` | Goroutine leak detection |
| `enable/stubs/` | Public API for enabling optional features (yaml, colors) |

### Code generation architecture

All assertion implementations live in `internal/assertions/`, organized by domain files
(boolean.go, collection.go, equal.go, error.go, etc.). Each assertion function generates
up to 8 variants:

1. `assert.Equal(t, ...)` — package-level function
2. `assert.Equalf(t, ..., "msg")` — format variant
3. `a.Equal(...)` — forward method (`a := assert.New(t)`)
4. `a.Equalf(..., "msg")` — forward format variant
5–8. Same four variants in `require/` (fatal on failure)

Run `go generate ./...` to regenerate `assert/` and `require/` packages from source.

### Dependencies

The root module has **zero external dependencies** by design. Do not add any.
Optional functionality uses the "enable" pattern: users blank-import a sub-module
(e.g., `_ "github.com/go-openapi/testify/v2/enable/yaml"`) to activate features
that require external deps.

### Key API

- `assert.Equal(t, expected, actual)` — non-fatal equality check
- `require.Equal(t, expected, actual)` — fatal equality check
- `assert.New(t)` / `require.New(t)` — create forwarded assertion objects
- All assertions in `internal/assertions/` follow the pattern: `func Name(t T, args..., msgAndArgs ...any) bool`

### Notable design decisions

- **Single source of truth** — write assertions once in `internal/assertions/`, generate everything else.
- **Example-driven test generation** — `Examples:` sections in doc comments drive generated tests for all 8 variants.
- **Domain tagging** — `// domain: equality` tags in doc comments organize documentation by concern.
- **Enable pattern** — optional features (YAML, colors) activated via blank imports, keeping the core dependency-free.
- **v2.0.0 is retracted** — see `go.mod`.

## Module Information

- Module path: `github.com/go-openapi/testify/v2`
- Go version: 1.25.0
- License: Apache-2.0
