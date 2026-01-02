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
// - comparison: comparing ordered values
// - condition: expressing assertions using conditions
// - equality: asserting two things are equal
// - error: asserting errors
// - file: asserting OS files
// - http: asserting HTTP response and body
// - json: asserting JSON documents
// - number: asserting numbers
// - ordering: asserting how collections are ordered
// - panic: asserting a panic behavior
// - string: asserting strings
// - testing: mimicks methods from the testing standard library
// - time: asserting times and durations
// - type: asserting types rather than values
// - yaml: asserting yaml documents
package assertions
