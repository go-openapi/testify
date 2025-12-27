// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package signature provides utilities for extracting and formatting Go function signatures and types.
//
// This package handles:
//
//   - Extracting function signatures from types.Signature
//   - Type qualification using import aliases
//   - Converting types to strings with proper package selectors
//
// The Extractor type maintains the context needed for proper type qualification,
// including the current package and import aliases from the source code.
package signature
