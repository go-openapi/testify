// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package comments provides utilities for extracting comments from Go source code.
//
// This package handles the bridge between go/types (semantic view) and go/ast (syntactic view)
// using position-based lookup. It extracts:
//
//   - Package-level comments and copyright headers
//   - Doc comments for functions, types, constants, and variables
//   - Body comments from within functions and type declarations
//   - Domain descriptions from package comments
//
// The Extractor type maintains the context needed for position-based lookup,
// including the file set, files map, and declaration cache.
package comments
