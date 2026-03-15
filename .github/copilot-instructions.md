# Copilot Instructions

This is the `go-openapi/testify` repository — a zero-dependency fork of
`stretchr/testify` used across all go-openapi repositories.

## Architecture

All assertion implementations live in `internal/assertions/`, organized by domain.
The `assert/` and `require/` packages are **generated** from that source via
`go generate ./...`. Do not edit generated files directly.

Each assertion function generates up to 8 variants (package-level, format, forward,
forward-format, in both assert and require).

## Adding assertions

1. Add function to appropriate domain file in `internal/assertions/`
2. Add `Examples:` and `domain:` tags in the doc comment
3. Add tests in the corresponding `*_test.go` file
4. Run `go generate ./...`

## Optional features

YAML and color support require blank imports of enable packages:

```go
import _ "github.com/go-openapi/testify/v2/enable/yaml"
import _ "github.com/go-openapi/testify/v2/enable/colors"
```

## Conventions

Coding conventions are found beneath `.github/copilot`

### Summary

- All `.go` files must have SPDX license headers (Apache-2.0).
- Commits require DCO sign-off (`git commit -s`).
- Linting: `golangci-lint run` — config in `.golangci.yml` (posture: `default: all` with explicit disables).
- Every `//nolint` directive **must** have an inline comment explaining why.
- Tests: `go test work ./...` (mono-repo). CI runs on `{ubuntu, macos, windows} x {stable, oldstable}` with `-race`.
- Test framework: this IS the test framework (`github.com/go-openapi/testify/v2`); `testifylint` does not work with this fork.

See `.github/copilot/` (symlinked to `.claude/rules/`) for detailed rules on Go conventions, linting, testing, and contributions.
