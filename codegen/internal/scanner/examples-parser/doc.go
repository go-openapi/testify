// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package parser scans a package of go source to extract testable examples.
//
// The outcome of the parser is an index of testable examples by exported function or type.
//
// Provided testable examples are a structure that may be rendered as go code.
//
// This package is freely inspired by go team's pkgsite tool (github.com/golang/pkgsite).
package parser
