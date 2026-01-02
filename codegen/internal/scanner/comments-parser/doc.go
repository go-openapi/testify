// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package parser provides text parsing utilities for extracting structured information from comments.
//
// This package contains pure text parsing logic with no AST dependencies. It parses:
//
//   - Test values from "Examples:" sections in doc comments
//   - Tagged comments (domain, maintainer, note, mention)
//   - Domain descriptions from package-level comments
//
// The parsers use regular expressions to identify sections and extract values.
// All parsing is case-insensitive and supports multiple formatting styles.
package parser
