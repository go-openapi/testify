// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bufio"
	"strings"
	"testing"
)

// a few colors for testing.
//
// The complete set may be found in ../assertions/enable/colors.
const (
	redMark    = "\033[0;31m"
	greenMark  = "\033[0;32m"
	yellowMark = "\033[0;33m"
	cyanMark   = "\033[0;36m"
	endMark    = "\033[0m"
)

func TestOptions(t *testing.T) {
	const (
		a = "(map[spew_test.stringer]int) (len=3) {\n" +
			"(spew_test.stringer) (len=1) stringer 1: (int) 1,\n" +
			"(spew_test.stringer) (len=1) stringer 2: (int) 2,\n" +
			"(spew_test.stringer) (len=1) stringer 3: (int) 3\n" +
			"(spew_test.stringer) (len=1) stringer 5: (int) 3\n" +
			"}\n"
		b = "(map[spew_test.stringer]int) (len=3) {\n" +
			"(spew_test.stringer) (len=1) stringer 1: (int) 1,\n" +
			"(spew_test.stringer) (len=1) stringer 2: (int) 3,\n" +
			"(spew_test.stringer) (len=1) stringer 3: (int) 3\n" +
			"(spew_test.stringer) (len=1) stringer 4: (int) 8\n" +
			"(spew_test.stringer) (len=1) stringer 6: (int) 9\n" +
			"}\n"
	)
	greenPrinterBuilder := ansiPrinterBuilder(greenMark)
	cyanPrinterBuilder := ansiPrinterBuilder(cyanMark)
	redPrinterBuilder := ansiPrinterBuilder(redMark)
	yellowPrinterBuilder := ansiPrinterBuilder(yellowMark)

	diff, err := GetUnifiedDiffString(UnifiedDiff{
		A:        SplitLines(a),
		B:        SplitLines(b),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
		Options: &Options{
			EqualPrinter:  greenPrinterBuilder,
			DeletePrinter: redPrinterBuilder,
			UpdatePrinter: cyanPrinterBuilder,
			InsertPrinter: yellowPrinterBuilder,
		},
	})
	if err != nil {
		t.Fatalf("did not expect an error, but got: %v", err)
	}

	//nolint:staticcheck // ST1018: for this test we specifically want to check escape sequences
	if !strings.Contains(diff, `[0;32m (spew_test.stringer) (len=1) stringer 1: (int) 1,
[0m[0;31m-(spew_test.stringer) (len=1) stringer 2: (int) 2,
[0m[0;33m+(spew_test.stringer) (len=1) stringer 2: (int) 3,
[0m[0;32m (spew_test.stringer) (len=1) stringer 3: (int) 3
[0m[0;31m-(spew_test.stringer) (len=1) stringer 5: (int) 3`,
	) {
		t.Errorf("expected matching ansi color sequences for diff")
	}

	// a visualization is better in this case...
	t.Log("\n\nDiff:\n" + diff)
}

func ansiPrinterBuilder(mark string) PrinterBuilder {
	return func(w *bufio.Writer) Printer {
		return func(str string) (err error) {
			_, err = w.WriteString(mark)
			if err != nil {
				return
			}
			_, err = w.WriteString(str)
			if err != nil {
				return
			}
			_, err = w.WriteString(endMark)
			if err != nil {
				return
			}

			return nil
		}
	}
}

// TestDefaultPrinterBuilder tests the DefaultPrinterBuilder function.
func TestDefaultPrinterBuilder(t *testing.T) {
	var buf strings.Builder
	w := bufio.NewWriter(&buf)

	printer := DefaultPrinterBuilder(w)
	err := printer("hello world")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	w.Flush()

	if buf.String() != "hello world" {
		t.Errorf("expected 'hello world', got %q", buf.String())
	}
}

// TestOptionsWithAllCustomPrinters tests that all custom printers are applied.
func TestOptionsWithAllCustomPrinters(t *testing.T) {
	const (
		a = "line1\nline2\nline3\n"
		b = "line1\nmodified\nline3\nnewline\n"
	)

	// Create custom printers for all types including OtherPrinter and Formatter
	customPrinter := func(prefix string) PrinterBuilder {
		return func(w *bufio.Writer) Printer {
			return func(str string) error {
				_, err := w.WriteString(prefix + str)
				return err
			}
		}
	}

	// customFormatter is a simplified formatter for testing purposes only.
	// It only handles string arguments with %s placeholders.
	// This is sufficient for the diff header format strings used in the library.
	customFormatter := func(w *bufio.Writer) Formatter {
		return func(format string, args ...any) error {
			s := "[FMT]" + format
			for _, arg := range args {
				if str, ok := arg.(string); ok {
					s = strings.Replace(s, "%s", str, 1)
				}
			}
			_, err := w.WriteString(s)
			return err
		}
	}

	diff, err := GetUnifiedDiffString(UnifiedDiff{
		A:        SplitLines(a),
		B:        SplitLines(b),
		FromFile: "Original",
		ToFile:   "Modified",
		Context:  1,
		Options: &Options{
			EqualPrinter:  customPrinter("[EQ]"),
			DeletePrinter: customPrinter("[DEL]"),
			UpdatePrinter: customPrinter("[UPD]"),
			InsertPrinter: customPrinter("[INS]"),
			OtherPrinter:  customPrinter("[OTH]"),
			Formatter:     customFormatter,
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the custom formatter was used for headers
	if !strings.Contains(diff, "[FMT]") {
		t.Error("expected custom formatter to be used")
	}

	// Verify custom printers were used
	if !strings.Contains(diff, "[EQ]") {
		t.Error("expected equal printer to be used")
	}
	if !strings.Contains(diff, "[DEL]") {
		t.Error("expected delete printer to be used")
	}
	if !strings.Contains(diff, "[INS]") {
		t.Error("expected insert printer to be used")
	}
}

// TestOptionsWithDefaults tests that default options are applied correctly.
func TestOptionsWithDefaults(t *testing.T) {
	// Test with nil options
	opts := optionsWithDefaults(nil)
	if opts == nil {
		t.Fatal("expected non-nil options")
	}
	if opts.EqualPrinter == nil {
		t.Error("expected EqualPrinter to be set")
	}
	if opts.DeletePrinter == nil {
		t.Error("expected DeletePrinter to be set")
	}
	if opts.UpdatePrinter == nil {
		t.Error("expected UpdatePrinter to be set")
	}
	if opts.InsertPrinter == nil {
		t.Error("expected InsertPrinter to be set")
	}
	if opts.OtherPrinter == nil {
		t.Error("expected OtherPrinter to be set")
	}
	if opts.Formatter == nil {
		t.Error("expected Formatter to be set")
	}

	// Test with partial options (only some fields set)
	partialOpts := &Options{
		EqualPrinter: ansiPrinterBuilder(greenMark),
	}
	opts = optionsWithDefaults(partialOpts)
	if opts.EqualPrinter == nil {
		t.Error("expected EqualPrinter to be preserved")
	}
	if opts.DeletePrinter == nil {
		t.Error("expected DeletePrinter to have default")
	}
}
