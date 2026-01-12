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
