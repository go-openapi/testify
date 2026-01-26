// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"bufio"
	"bytes"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/internal/difflib"
)

func TestMakeDiffOptions(t *testing.T) {
	t.Parallel()

	for c := range makeDiffOptionsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := makeDiffOptions(c.opts)

			if c.expectNil {
				if result != nil {
					t.Errorf("expected nil, got %+v", result)
				}
				return
			}

			if !c.expectPrinters {
				t.Fatal("test case must specify expectPrinters or expectNil")
			}

			if result == nil {
				t.Fatal("expected non-nil result")
			}

			if c.validatePrinters != nil {
				c.validatePrinters(t, result)
			}
		})
	}
}

func TestSetColorizers(t *testing.T) {
	t.Parallel()

	for c := range setColorizerCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := setColorizers(c.opts)

			if c.validateColorizers != nil {
				c.validateColorizers(t, result)
			}
		})
	}
}

// ==============================
// Test makeDiffOptions function
// ==============================

type makeDiffOptionsCase struct {
	name             string
	opts             options
	expectNil        bool
	expectPrinters   bool
	validatePrinters func(*testing.T, *difflib.Options)
}

const testPrinter = "test"

// printerSpec defines the expected ANSI marks for each printer type in a theme.
type printerSpec struct {
	equal  string
	delete string
	update string
	insert string
}

// validateAllPrinters validates all four printer types against the expected spec.
func validateAllPrinters(t *testing.T, o *difflib.Options, spec printerSpec) {
	t.Helper()

	validatePrinter(t, o.EqualPrinter, spec.equal, "EqualPrinter")
	validatePrinter(t, o.DeletePrinter, spec.delete, "DeletePrinter")
	validatePrinter(t, o.UpdatePrinter, spec.update, "UpdatePrinter")
	validatePrinter(t, o.InsertPrinter, spec.insert, "InsertPrinter")
}

// validatePrinter tests that a printer builder produces the expected ANSI-wrapped output.
func validatePrinter(t *testing.T, builder difflib.PrinterBuilder, expectedMark, printerName string) {
	t.Helper()

	if builder == nil {
		t.Errorf("%s should not be nil", printerName)
		return
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)

	printer := builder(w)
	if err := printer(testPrinter); err != nil {
		t.Errorf("%s error: %v", printerName, err)
		return
	}

	if err := w.Flush(); err != nil {
		t.Errorf("%s flush error: %v", printerName, err)
		return
	}

	expected := expectedMark + testPrinter + endMark
	if buf.String() != expected {
		t.Errorf("%s expected %q, got %q", printerName, expected, buf.String())
	}
}

func makeDiffOptionsCases() iter.Seq[makeDiffOptionsCase] {
	return slices.Values([]makeDiffOptionsCase{
		{
			name: "disabled colors should return nil",
			opts: options{
				enabled: false,
				theme:   ThemeDark,
			},
			expectNil: true,
		},
		{
			name: "enabled with ThemeLight should return light color printers",
			opts: options{
				enabled: true,
				theme:   ThemeLight,
			},
			expectPrinters: true,
			validatePrinters: func(t *testing.T, o *difflib.Options) {
				t.Helper()
				spec := printerSpec{
					equal:  greenMark,
					delete: redMark,
					update: cyanMark,
					insert: yellowMark,
				}
				validateAllPrinters(t, o, spec)
			},
		},
		{
			name: "enabled with ThemeDark should return bright color printers",
			opts: options{
				enabled: true,
				theme:   ThemeDark,
			},
			expectPrinters: true,
			validatePrinters: func(t *testing.T, o *difflib.Options) {
				t.Helper()
				spec := printerSpec{
					equal:  brightGreenMark,
					delete: brightRedMark,
					update: brightCyanMark,
					insert: brightYellowMark,
				}
				validateAllPrinters(t, o, spec)
			},
		},
		{
			name: "enabled with unknown theme should return nil",
			opts: options{
				enabled: true,
				theme:   Theme("unknown"),
			},
			expectNil: true,
		},
	})
}

// ==============================
// Test setColorizers function
// ==============================

type setColorizerCase struct {
	name               string
	opts               options
	validateColorizers func(*testing.T, colorizers)
}

// validateColorizer tests that a colorizer produces the expected output.
func validateColorizer(t *testing.T, colorizer StringColorizer, expectedOutput, colorizerName string) {
	t.Helper()

	got := colorizer(testPrinter)
	if got != expectedOutput {
		t.Errorf("%s expected %q, got %q", colorizerName, expectedOutput, got)
	}
}

// validateNoopColorizer tests that a colorizer returns input unchanged.
func validateNoopColorizer(t *testing.T, colorizer StringColorizer, colorizerName string) {
	t.Helper()

	got := colorizer(testPrinter)
	if got != testPrinter {
		t.Errorf("%s should be noop, got %q for input %q", colorizerName, got, testPrinter)
	}
}

func setColorizerCases() iter.Seq[setColorizerCase] {
	return slices.Values([]setColorizerCase{
		{
			name: "disabled colors should return noop colorizers",
			opts: options{
				enabled: false,
				theme:   ThemeDark,
			},
			validateColorizers: func(t *testing.T, c colorizers) {
				t.Helper()
				validateNoopColorizer(t, c.expected, "expected")
				validateNoopColorizer(t, c.actual, "actual")
			},
		},
		{
			name: "ThemeLight should return green/red colorizers",
			opts: options{
				enabled: true,
				theme:   ThemeLight,
			},
			validateColorizers: func(t *testing.T, c colorizers) {
				t.Helper()
				validateColorizer(t, c.expected, greenMark+testPrinter+endMark, "expected")
				validateColorizer(t, c.actual, redMark+testPrinter+endMark, "actual")
			},
		},
		{
			name: "ThemeDark should return bright green/red colorizers",
			opts: options{
				enabled: true,
				theme:   ThemeDark,
			},
			validateColorizers: func(t *testing.T, c colorizers) {
				t.Helper()
				validateColorizer(t, c.expected, brightGreenMark+testPrinter+endMark, "expected")
				validateColorizer(t, c.actual, brightRedMark+testPrinter+endMark, "actual")
			},
		},
		{
			name: "unknown theme should return noop colorizers",
			opts: options{
				enabled: true,
				theme:   Theme("unknown"),
			},
			validateColorizers: func(t *testing.T, c colorizers) {
				t.Helper()
				validateNoopColorizer(t, c.expected, "expected")
				validateNoopColorizer(t, c.actual, "actual")
			},
		},
	})
}
