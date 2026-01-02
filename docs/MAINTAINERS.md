# Maintainer's guide

## Repo structure

Monorepo with multiple go modules.

## Repo configuration

* default branch: master
* protected branches: master
* branch protection rules:
  * require pull requests and approval
  * required status checks: 
    - DCO (simple email sign-off)
    - Lint
    - tests completed
* auto-merge enabled (used for dependabot updates)

## Continuous Integration

### Code Quality checks

* meta-linter: golangci-lint
* linter config: [`.golangci.yml`](../.golangci.yml) (see our [posture](./STYLE.md) on linters)

* Code quality assessment: [CodeFactor](https://www.codefactor.io/dashboard)
* Code quality badges
  * go report card: <https://goreportcard.com/>
  * CodeFactor: <https://goreportcard.com/>

> **NOTES**
>
> codefactor inherits roles from github. There is no need to create a dedicated account.
>
> The codefactor app is installed at the organization level (`github.com/go-openapi`).
>
> There is no special token to setup in github for CI usage.

### Testing

* Test reports
  * Uploaded to codecov: <https://app.codecov.io/analytics/gh/go-openapi>
* Test coverage reports
  * Uploaded to codecov: <https://app.codecov.io/gh/go-openapi>

* Fuzz testing
  * Fuzz tests are handled separately by CI and may reuse a cached version of the fuzzing corpus.
    At this moment, cache may not be shared between feature branches or feature branch and master.
    The minimized corpus produced on failure is uploaded as an artifact and should be added manually
    to `testdata/fuzz/...`.

Coverage threshold status is informative and not blocking.
This is because the thresholds are difficult to tune and codecov oftentimes reports false negatives
or may fail to upload coverage.

All tests across `go-openapi` use our fork of `stretchr/testify` (this repo): `github.com/go-openapi/testify`.
This allows for minimal test dependencies.

> **NOTES**
>
> codecov inherits roles from github. There is no need to create a dedicated account.
> However, there is only 1 maintainer allowed to be the admin of the organization on codecov
> with their free plan.
>
> The codecov app is installed at the organization level (`github.com/go-openapi`).
>
> There is no special token to setup in github for CI usage.
> A organization-level token used to upload coverage and test reports is managed at codecov:
> no setup is required on github.

### Automated updates

* dependabot
  * configuration: [`dependabot.yaml`](../.github/dependabot.yaml)

  Principle:

  * codecov applies updates and security patches to the github-actions and golang ecosystems.
  * all updates from "trusted" dependencies (github actions, golang.org packages, go-openapi packages
    are auto-merged if they successfully pass CI.

* go version udpates

  Principle:

  * we support the 2 latest minor versions of the go compiler (`stable`, `oldstable`)
  * `go.mod` should be updated (manually) whenever there is a new go minor release
    (e.g. every 6 months).

* contributors
  * a [`CONTRIBUTORS.md`](../CONTRIBUTORS.md) file is updated weekly, with all-time contributors to the repository
  * the `github-actions[bot]` posts a pull request to do that automatically
  * at this moment, this pull request is not auto-approved/auto-merged (bot cannot approve its own PRs)

### Vulnerability scanners

There are 3 complementary scanners - obviously, there is some overlap, but each has a different focus.

* github `CodeQL`
* `trivy` <https://trivy.dev/docs/latest/getting-started>
* `govulnscan` <https://go.dev/blog/govulncheck>

None of these tools require an additional account or token.

Github CodeQL configuration is set to "Advanced", so we may collect a CI status for this check (e.g. for badges).

Scanners run on every commit to master and at least once a week.

Reports are centralized in github security reports for code scanning tools.

## Releases

The release process is minimalist:

* push a semver tag (i.e v{major}.{minor}.{patch}) to the master branch.
* the CI handles this to generate a github release with release notes

* release notes generator: git-cliff <https://git-cliff.org/docs/>
* configuration: [`cliff.toml`](../.cliff.toml)

Tags are preferably PGP-signed.

The tag message introduces the release notes (e.g. a summary of this release).

The release notes generator does not assume that commits are necessarily "conventional commits".

## Other files

Standard documentation:

* [`CONTRIBUTING.md`](../.github/CONTRIBUTING.md) guidelines
* [`DCO.md`](../.github/DCO.md) terms for first-time contributors to read
* [`CODE_OF_CONDUCT.md`](../CODE_OF_CONDUCT.md)
* [`SECURIY.md`](../SECURITY.md) policy: how to report vulnerabilities privately
* [`LICENSE`](../LICENSE) terms
* [`NOTICE`](../NOTICE) on supplementary license terms (original authors, copied code etc)

Reference documentation (released):

* [godoc](https://pkg.go/dev/go-openapi/testify)

## Maintaining Generated Code

This repository uses code generation extensively to maintain consistency across assertion packages.

### Architecture Overview

**Source of Truth: `internal/assertions/`**

All assertion implementations live in `internal/assertions/`, organized by domain:
- Functions are implemented once with comprehensive tests
- Doc comments include "Examples:" sections that drive test generation
- Both `assert/` and `require/` packages are 100% generated from this source

**Code Generator: `codegen/`**

The generator scans `internal/assertions/` and produces:
- 76 assertion functions × 8 variants = 608 generated functions
- Package-level functions (`assert.Equal`, `require.Equal`)
- Format variants (`assert.Equalf`, `require.Equalf`)
- Forward methods (`a.Equal()`, `r.Equal()`)
- Tests for all variants
- Testable examples for godoc

**Generated Packages: `assert/` and `require/`**

Everything in these packages is generated. Never edit generated files directly.

Exceptions:
* `doc.go` is not generated
* the `assert` package contains a small `enable` package to enable features. This is not generated.

### Adding a New Assertion

**Complete workflow:**

1. **Add function to `internal/assertions/<domain>.go`:**

   The following example would like go to `string.go`, next to the `Regexp` assertion.

   ```go
   // StartsWith asserts that the string starts with the given prefix.
   //
   // Examples:
   //
   //   success: "hello world", "hello"
   //   failure: "hello world", "bye"
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

2. **Add tests to `internal/assertions/<domain>_test.go`:**
   Write comprehensive table-driven tests covering edge cases.

3. **Run code generation:**
   ```bash
   go generate ./...
   ```

4. **Done!** All 8 variants are generated with tests and examples:
   - `assert.StartsWith(t, str, prefix)`
   - `assert.StartsWithf(t, str, prefix, "msg")`
   - `a.StartsWith(str, prefix)` (forward method)
   - `a.StartsWithf(str, prefix, "msg")`
   - `require.StartsWith(t, str, prefix)`
   - `require.StartsWithf(t, str, prefix, "msg")`
   - `r.StartsWith(str, prefix)` (forward method)
   - `r.StartsWithf(str, prefix, "msg")`

### Example Annotations Format

The "Examples:" section in doc comments drives test and example generation:

```go
// Examples:
//
//   success: <test arguments that should succeed>
//   failure: <test arguments that should fail>
//   panic: <test arguments that cause panic>
//          <expected panic message>
```

**Rules:**
- Use valid Go expressions that can be directly inserted into test code
- `success:` and `failure:` are required for most assertions
- `panic:` is optional (used for assertions like Panics, YAMLEq)
- Multiple examples of the same type are allowed (e.g., multiple `success:` lines)
- Examples are extracted by the scanner and used to generate:
  - Unit tests (success + failure cases)
  - Testable examples (success cases only for simplicity)

**Example with multiple success cases:**
```go
// Examples:
//
//   success: []string{"a", "b"}, 2        // slice
//   success: map[string]int{"a": 1}, 1    // map
//   success: "hello", 5                    // string
//   failure: []string{"a"}, 5
```

**Important: Placeholder marker for complex values**

When example values are too complex to represent inline (pointers, large structs, etc.), use `// NOT IMPLEMENTED` as the placeholder marker:

```go
// Examples:
//
//   success: &customStruct{Field: "value"}, // NOT IMPLEMENTED
//   failure: complexType{}, // NOT IMPLEMENTED
```

**Never use `// TODO`** - it triggers false positives in:
- Code quality analyzers (linters scanning for TODO items)
- Project management tools (TODO trackers)
- IDE warnings about unfinished work

The `// NOT IMPLEMENTED` marker clearly indicates a placeholder without triggering these tools.

### Special Cases in Generated Tests

For complex assertions requiring special setup, the test templates support conditional logic. See `codegen/internal/generator/templates/assertion_assertions_test.gotmpl` for examples of:
- Custom mock selection based on function behavior (mockT vs mockFailNowT)
- Package-specific test helpers (testDataPath, httpOK, etc.)
- Handling functions without test examples (generates `t.Skip()`)

Some go expressions won't fit nicely for examples (examples use an external package, e.g. `assert_test`).
To cover these edge cases, a `relocate` function map currently rewrites the example values to be used
from an external package. Most transforms there are hard-coded specifically for the 3 edges cases identified so far
(see `codegen/internal/generator/funcmap.go`).

### Generator Flags

```bash
go run ./codegen/main.go \
  -work-dir=.. \
  -input-package=github.com/go-openapi/testify/v2/internal/assertions \
  -output-packages=assert,require \
  -target-root=.. \
  -include-format-funcs=true \
  -include-forward-funcs=true \
  -include-tests=true \
  -include-examples=true \
  -runnable-examples=true \
  -include-helpers=true \
  -include-generics=false
```

Current usage with `go generate` (see `doc.go`):

```go
//go:generate go run ./codegen/main.go -target-root . -work-dir .
```

**Note:** Generic functions are planned but not yet implemented.

### Verification

After generation, verify:
```bash
# All tests pass
go test ./...

# Coverage remains high
go test -cover ./internal/assertions  # Should be ~94%
go test -cover ./assert               # Should be ~99.5%
go test -cover ./require              # Should be ~99.5%

# Examples are valid
go test -run Example ./assert
go test -run Example ./require
```

The 0.5% coverage gap comes from helper functions (non-assertion functions) that don't have "Examples:" annotations.

## TODOs & other ideas

A few things remain ahead to ease a bit a maintainer's job:

* [ ] reuse CI workflows (e.g. in `github.com/go-openapi/workflows`)
* [ ] reusable actions with custom tools pinned  (e.g. in `github.com/go-openapi/gh-actions`)
* open-source license checks
* [ ] auto-merge for CONTRIBUTORS.md (requires a github app to produce tokens)
* [x] more automated code renovation / relinting work (possibly built with CLAUDE)
* organization-level documentation web site
* ...
