// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package stubs provides public APIs for enabling optional features in testify.
//
// This package exports stub implementations that delegate to internal packages,
// maintaining a clean separation between internal implementation and public API.
//
// Subpackages:
//
//   - yaml: API for enabling YAML assertions
//   - colors: API for enabling colorized output
//
// These stubs are used by the enable/{yaml,colors} modules to activate optional features.
package stubs
