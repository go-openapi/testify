//go:build integrationtest

// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package assertions holds the internal implementation
// of all the helper functions exposed by testify.
//
// It serves as a base to develop and test testify,
// whereas the publicly exposed API (in packages assert and require)
// is generated.
//
// # Domains
//
// - boolean: asserting boolean values
// - collection: asserting slices and maps
// - compare: comparing ordered values
// - condition: expressing assertions using conditions
// - equal: asserting two things are equal
// - error: asserting errors
// - file: asserting os files
// - helpers: extra helpers that are not assertions
// - http: asserting http response and body
// - json: asserting JSON documents
// - number: asserting numbers
// - order: asserting how collection are ordered
// - panic: asserting a panic behavior
// - string: asserting strings
// - time: asserting times and durations
// - type: asserting types rather than values
// - yaml: asserting yaml documents
package assertions
