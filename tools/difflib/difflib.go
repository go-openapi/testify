// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package difflib provides a unified diff for sequences of strings.
//
// It exposes the unified diff formatter used internally by [github.com/go-openapi/testify/v2].
package difflib

import (
	"io"

	idifflib "github.com/go-openapi/testify/v2/internal/difflib"
)

// SplitLines splits a string on "\n" while preserving them. The output can be used
// as input for the [UnifiedDiff] structure.
func SplitLines(s string) []string {
	return idifflib.SplitLines(s)
}

// UnifiedDiff holds the unified diff parameters.
type UnifiedDiff = idifflib.UnifiedDiff

// Options to fine-tune the rendering of the diff output.
type Options = idifflib.Options

// Formatter is a formatting function like [fmt.Printf].
type Formatter = idifflib.Formatter

// Printer outputs a formatted string, e.g. to some underlying writer.
type Printer idifflib.Printer

// FormatterBuilder is a function that builds a [Formatter] given a buffered [bufio.Writer].
type FormatterBuilder = idifflib.FormatterBuilder

// PrinterBuilder is a function that builds a [Printer] given a buffered [bufio.Writer].
type PrinterBuilder = idifflib.PrinterBuilder

// GetUnifiedDiffString is like [WriteUnifiedDiff] but returns the diff as a string instead of writing to an [io.Writer].
func GetUnifiedDiffString(diff UnifiedDiff) (string, error) {
	return idifflib.GetUnifiedDiffString(diff)
}

// WriteUnifiedDiff writes the comparison between two sequences of lines.
//
// It generates the delta as a unified diff.
func WriteUnifiedDiff(writer io.Writer, diff UnifiedDiff) error {
	return idifflib.WriteUnifiedDiff(writer, diff)
}
