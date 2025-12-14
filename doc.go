// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package testify provides comprehensive assertion packages for Go testing.
//
// # Overview
//
// This is the go-openapi fork of testify, designed with zero external dependencies
// and a focus on maintainability. The codebase uses code generation more extensively
// than the original testify. This helps maintain consistency across all assertion variants.
//
// # Packages
//
// The [assert] package provides a comprehensive set of assertion functions that
// integrate with Go's testing framework. Assertions return boolean values allowing
// tests to continue after failures.
//
// The [require] package provides the same assertions but with fatal checks that stop
// test execution immediately on failure via [testing.T.FailNow].
//
// # Key Differences from stretchr/testify
//
// This fork prioritizes:
//   - Zero external dependencies (go-spew and difflib are internalized)
//   - Removed mock and suite packages (favor the use mockery or similar specialized tools instead)
//   - Optional features via enable packages (e.g., enable/yaml for YAML assertions)
//   - Code generation ensures consistency across 76 assertion functions × 8 variants
//
// # Optional Features
//
// YAML assertions require importing the enable package:
//
//	import _ "github.com/go-openapi/testify/v2/enable/yaml"
//
// Without this import, YAMLEq and YAMLEqf will panic loudly with a helpful error message.
//
// # Note on testifylint
//
// The golangci-lint compatible linter [testifylint] is designed for stretchr/testify
// and will not work with this fork as it checks only for the original dependency.
//
// [assert]: https://pkg.go.dev/github.com/go-openapi/testify/v2/assert
// [require]: https://pkg.go.dev/github.com/go-openapi/testify/v2/require
// [testifylint]: https://github.com/Antonboom/testifylint
package testify

//go:generate go run ./codegen/main.go -target-root . -work-dir .
