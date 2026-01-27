// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package colors enables colorized tests with basic and portable ANSI terminal codes.
//
// Colorization is disabled by default when the standard output is not a terminal.
//
// Colors are somewhat limited. We want the package to work on unix and windows without any extra dependencies.
//
// # Usage
//
// To enable the test colorization feature, use a blank import like so:
//
//	import (
//			_ "github.com/go-openapi/testify/enable/colors/v2"
//	)
//
// # Command line arguments
//
//   - testify.colorized={true|false}
//   - testify.theme={dark|light}
//   - testify.colorized.notty={true|false}  (enable colorization even when the output is not a terminal)
//
// The default theme used is dark.
//
// To run tests on a terminal with colorized output:
//
//   - run: go test -v -testify.colorized ./...
//
// # Environment variables
//
// Colorization may be enabled from environment:
//
//   - TESTIFY_COLORIZED=true
//   - TESTIFY_THEME=dark
//   - TESTIFY_COLORIZED_NOTTY=true
//
// Command line arguments take precedence over environment.
package colors
