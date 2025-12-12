// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package difflib is a partial port of Python difflib module.
//
// It provides tools to compare sequences of strings and generate textual diffs.
//
// The following class and functions have been ported:
//
// - SequenceMatcher
//
// - unified_diff
//
// Getting unified diffs was the main goal of the port. Keep in mind this code
// is mostly suitable to output text differences in a human friendly way, there
// are no guarantees generated diffs are consumable by patch(1).
//
// This package was adopted from [github.com/pmezard/go-difflib] which
// is no longer maintained.
//
// [github.com/pmezard/go-difflib]: https://github.com/pmezard/go-difflib
package difflib
