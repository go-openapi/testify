/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew_test

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"slices"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// TestSpew executes all of the tests described by spewTestCases.
func TestSpew(t *testing.T) {
	t.Parallel()

	runners := spewTestRunners()

	i := 0
	for tc := range spewTestCases() {
		buf := new(bytes.Buffer)
		runner, ok := runners[tc.f]
		if !ok {
			t.Errorf("test configuration: %v #%d unrecognized function", tc.f, i)
			continue
		}

		if errMsg := runner(t, buf, tc); errMsg != "" {
			t.Errorf("%v #%d %s", tc.f, i, errMsg)
			continue
		}

		if s := buf.String(); tc.want != s {
			t.Errorf("ConfigState #%d\n  got: %s want: %s", i, s, tc.want)
		}

		i++
	}
}

// spewFunc is used to identify which public function of the spew package or
// ConfigState a test applies to.
type spewFunc int

const (
	fCSFdump spewFunc = iota
	fCSFprint
	fCSFprintf
	fCSFprintln
	fCSPrint
	fCSPrintln
	fCSSdump
	fCSSprint
	fCSSprintf
	fCSSprintln
	fCSErrorf
	fCSNewFormatter
	fErrorf
	fFprint
	fFprintln
	fPrint
	fPrintln
	fSdump
	fSprint
	fSprintf
	fSprintln
)

func (f spewFunc) String() string {
	names := map[spewFunc]string{
		fCSFdump:        "ConfigState.Fdump",
		fCSFprint:       "ConfigState.Fprint",
		fCSFprintf:      "ConfigState.Fprintf",
		fCSFprintln:     "ConfigState.Fprintln",
		fCSSdump:        "ConfigState.Sdump",
		fCSPrint:        "ConfigState.Print",
		fCSPrintln:      "ConfigState.Println",
		fCSSprint:       "ConfigState.Sprint",
		fCSSprintf:      "ConfigState.Sprintf",
		fCSSprintln:     "ConfigState.Sprintln",
		fCSErrorf:       "ConfigState.Errorf",
		fCSNewFormatter: "ConfigState.NewFormatter",
		fErrorf:         "spew.Errorf",
		fFprint:         "spew.Fprint",
		fFprintln:       "spew.Fprintln",
		fPrint:          "spew.Print",
		fPrintln:        "spew.Println",
		fSdump:          "spew.Sdump",
		fSprint:         "spew.Sprint",
		fSprintf:        "spew.Sprintf",
		fSprintln:       "spew.Sprintln",
	}

	if s, ok := names[f]; ok {
		return s
	}

	return fmt.Sprintf("Unknown spewFunc (%d)", int(f))
}

// spewTest is used to describe a test to be performed against the public
// functions of the spew package or ConfigState.
type spewTest struct {
	cs     *spew.ConfigState
	f      spewFunc
	format string
	in     any
	want   string
}

// spewTestCases returns the tests to be performed against the public functions
// of the spew package and ConfigState.
//
// These tests are only intended to ensure the public functions are exercised
// and are intentionally not exhaustive of types.  The exhaustive type
// tests are handled in the dump and format tests.
func spewTestCases() iter.Seq[spewTest] {
	// Config states with various settings.
	scsDefault := spew.NewDefaultConfig()
	scsNoMethods := &spew.ConfigState{Indent: " ", DisableMethods: true}
	scsNoMethodsButTimeStringer := &spew.ConfigState{Indent: " ", DisableMethods: true, EnableTimeStringer: true}
	scsNoPmethods := &spew.ConfigState{Indent: " ", DisablePointerMethods: true}
	scsMaxDepth := &spew.ConfigState{Indent: " ", MaxDepth: 1}
	scsContinue := &spew.ConfigState{Indent: " ", ContinueOnMethod: true}
	scsNoPtrAddr := &spew.ConfigState{DisablePointerAddresses: true}
	scsNoCap := &spew.ConfigState{DisableCapacities: true}

	// Variables for tests on types which implement Stringer interface with and
	// without a pointer receiver.
	ts := stringer("test")
	tps := pstringer("test")

	type ptrTester struct {
		s *struct{}
	}
	tptr := &ptrTester{s: &struct{}{}}

	// depthTester is used to test max depth handling for structs, array, slices
	// and maps.
	type depthTester struct {
		ic    indirCir1
		arr   [1]string
		slice []string
		m     map[string]int
	}

	dt := depthTester{
		indirCir1{nil},
		[1]string{"arr"},
		[]string{"slice"},
		map[string]int{"one": 1},
	}

	// Variable for tests on types which implement error interface.
	te := customError(10)

	// Variables for testing time.Time behavior
	tm := time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)
	tmAddr := fmt.Sprintf("%p", &tm)
	em := embeddedTime{Time: tm}
	emptr := embeddedTimePtr{Time: &tm}
	er := embeddedRedeclaredTime{
		redeclaredTime: redeclaredTime(tm),
	}

	return slices.Values([]spewTest{
		// default config
		{scsDefault, fCSFdump, "", int8(127), "(int8) 127\n"},
		{scsDefault, fCSFprint, "", int16(32767), "32767"},
		{scsDefault, fCSFprintf, "%v", int32(2147483647), "2147483647"},
		{scsDefault, fCSFprintln, "", int(2147483647), "2147483647\n"},
		{scsDefault, fCSPrint, "", int64(9223372036854775807), "9223372036854775807"},
		{scsDefault, fCSPrintln, "", uint8(255), "255\n"},
		{scsDefault, fCSSdump, "", uint8(64), "(uint8) 64\n"},
		{scsDefault, fCSSprint, "", complex(1, 2), "(1+2i)"},
		{scsDefault, fCSSprintf, "%v", complex(float32(3), 4), "(3+4i)"},
		{scsDefault, fCSSprintln, "", complex(float64(5), 6), "(5+6i)\n"},
		{scsDefault, fCSErrorf, "%#v", uint16(65535), "(uint16)65535"},
		{scsDefault, fCSNewFormatter, "%v", uint32(4294967295), "4294967295"},
		{scsDefault, fErrorf, "%v", uint64(18446744073709551615), "18446744073709551615"},
		{scsDefault, fFprint, "", float32(3.14), "3.14"},
		{scsDefault, fFprintln, "", float64(6.28), "6.28\n"},
		{scsDefault, fPrint, "", true, "true"},
		{scsDefault, fPrintln, "", false, "false\n"},
		{scsDefault, fSdump, "", complex(-10, -20), "(complex128) (-10-20i)\n"},
		{scsDefault, fSprint, "", complex(-1, -2), "(-1-2i)"},
		{scsDefault, fSprintf, "%v", complex(float32(-3), -4), "(-3-4i)"},
		{scsDefault, fSprintln, "", complex(float64(-5), -6), "(-5-6i)\n"},
		// config with no methods
		{scsNoMethods, fCSFprint, "", ts, "test"},
		{scsNoMethods, fCSFprint, "", &ts, "<*>test"},
		{scsNoMethods, fCSFprint, "", tps, "test"},
		{scsNoMethods, fCSFprint, "", &tps, "<*>test"},
		{scsNoPmethods, fCSFprint, "", ts, "stringer test"},
		{scsNoPmethods, fCSFprint, "", &ts, "<*>stringer test"},
		{scsNoPmethods, fCSFprint, "", tps, "test"},
		{scsNoPmethods, fCSFprint, "", &tps, "<*>stringer test"},
		// config with maxdepth
		{scsMaxDepth, fCSFprint, "", dt, "{{<max>} [<max>] [<max>] map[<max>]}"},
		{scsMaxDepth, fCSFdump, "", dt, "(spew_test.depthTester) {\n" +
			" ic: (spew_test.indirCir1) {\n  <max depth reached>\n },\n" +
			" arr: ([1]string) (len=1 cap=1) {\n  <max depth reached>\n },\n" +
			" slice: ([]string) (len=1 cap=1) {\n  <max depth reached>\n },\n" +
			" m: (map[string]int) (len=1) {\n  <max depth reached>\n }\n}\n"},
		// config with continue on method
		{scsContinue, fCSFprint, "", ts, "(stringer test) test"},
		{scsContinue, fCSFdump, "", ts, "(spew_test.stringer) (len=4) (stringer test) \"test\"\n"},
		{scsContinue, fCSFprint, "", te, "(error: 10) 10"},
		{scsContinue, fCSFdump, "", te, "(spew_test.customError) (error: 10) 10\n"},
		{scsNoPtrAddr, fCSFprint, "", tptr, "<*>{<*>{}}"},
		{scsNoPtrAddr, fCSSdump, "", tptr, "(*spew_test.ptrTester)({\ns: (*struct {})({\n})\n})\n"},
		{scsNoCap, fCSSdump, "", make([]string, 0, 10), "([]string) {\n}\n"},
		{scsNoCap, fCSSdump, "", make([]string, 1, 10), "([]string) (len=1) {\n(string) \"\"\n}\n"},
		//
		// time.Time formatting with all configs
		{scsDefault, fCSFprint, "", tm, "2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsDefault, fCSFdump, "", tm, "(time.Time) 2006-01-02 15:04:05.999999999 +0000 UTC\n"},
		{scsDefault, fCSFprint, "", &tm, "<*>2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsDefault, fCSFdump, "", &tm, "(*time.Time)(" + tmAddr + ")(2006-01-02 15:04:05.999999999 +0000 UTC)\n"},
		{scsDefault, fCSFprint, "", em, "2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsDefault, fCSFdump, "", em, "(spew_test.embeddedTime) 2006-01-02 15:04:05.999999999 +0000 UTC\n"},
		{scsDefault, fCSFprint, "", emptr, "2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsDefault, fCSFdump, "", emptr, "(spew_test.embeddedTimePtr) 2006-01-02 15:04:05.999999999 +0000 UTC\n"},
		{scsNoMethods, fCSFprint, "", tm, "{999999999 63271811045 <nil>}"},
		{scsNoMethods, fCSFprint, "", &tm, "<*>{999999999 63271811045 <nil>}"},
		{scsNoMethods, fCSFprint, "", em, "{{999999999 63271811045 <nil>}}"},
		{scsNoMethods, fCSFprint, "", emptr, "{<*>{999999999 63271811045 <nil>}}"}, // NOTE: the type of the embedded pointer is not rendered, on purpose
		{
			scsContinue, fCSFdump, "", tm, // this prints time and continues digging the struct
			"(time.Time) (2006-01-02 15:04:05.999999999 +0000 UTC)" +
				" {\n wall: (uint64) 999999999,\n ext: (int64) 63271811045,\n loc: (*time.Location)(<nil>)\n}\n",
		},
		{
			scsContinue, fCSFdump, "", &tm,
			"(*time.Time)(" + tmAddr + ")((2006-01-02 15:04:05.999999999 +0000 UTC)" +
				" {\n wall: (uint64) 999999999,\n ext: (int64) 63271811045,\n loc: (*time.Location)(<nil>)\n})\n",
		},
		{
			scsContinue, fCSFdump, "", em,
			"(spew_test.embeddedTime) (2006-01-02 15:04:05.999999999 +0000 UTC) {\n" +
				" Time: (time.Time) (2006-01-02 15:04:05.999999999 +0000 UTC) {\n" +
				"  wall: (uint64) 999999999,\n  ext: (int64) 63271811045,\n  loc: (*time.Location)(<nil>)\n }\n}\n",
		},
		{
			scsContinue, fCSFdump, "", emptr,
			"(spew_test.embeddedTimePtr) (2006-01-02 15:04:05.999999999 +0000 UTC) {\n" +
				" Time: (*time.Time)(" + tmAddr + ")((2006-01-02 15:04:05.999999999 +0000 UTC) {\n" +
				"  wall: (uint64) 999999999,\n  ext: (int64) 63271811045,\n  loc: (*time.Location)(<nil>)\n })\n}\n",
		},
		{
			scsContinue, fCSFdump, "", er,
			"(spew_test.embeddedRedeclaredTime) {\n" +
				" redeclaredTime: (spew_test.redeclaredTime) (2006-01-02 15:04:05.999999999 +0000 UTC) {\n" +
				"  wall: (uint64) 999999999,\n  ext: (int64) 63271811045,\n  loc: (*time.Location)(<nil>)\n }\n}\n",
		},
		{scsNoMethodsButTimeStringer, fCSFprint, "", tm, "2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsNoMethodsButTimeStringer, fCSFdump, "", tm, "(time.Time) 2006-01-02 15:04:05.999999999 +0000 UTC\n"},
		{scsNoMethodsButTimeStringer, fCSFprint, "", &tm, "<*>2006-01-02 15:04:05.999999999 +0000 UTC"},
		{scsNoMethodsButTimeStringer, fCSFprint, "", em, "{2006-01-02 15:04:05.999999999 +0000 UTC}"},
		{scsNoMethodsButTimeStringer, fCSFdump, "", em, "(spew_test.embeddedTime) {\n Time: (time.Time) 2006-01-02 15:04:05.999999999 +0000 UTC\n}\n"},
		{scsNoMethodsButTimeStringer, fCSFprint, "", emptr, "{<*>2006-01-02 15:04:05.999999999 +0000 UTC}"},
		{scsNoMethodsButTimeStringer, fCSFdump, "", emptr, "(spew_test.embeddedTimePtr) {\n Time: (*time.Time)(" + tmAddr + ")(2006-01-02 15:04:05.999999999 +0000 UTC)\n}\n"},
	})
}

// spewTestRunner is a function that executes a spew test case, writing output
// to the buffer. It returns an error string if something goes wrong, or empty
// string on success.
type spewTestRunner func(t *testing.T, buf *bytes.Buffer, test spewTest) string

// spewTestRunners returns the dispatch table mapping each spewFunc to its
// test execution logic.
func spewTestRunners() map[spewFunc]spewTestRunner {
	return map[spewFunc]spewTestRunner{
		fCSFdump:  func(_ *testing.T, buf *bytes.Buffer, test spewTest) string { test.cs.Fdump(buf, test.in); return "" },
		fCSFprint: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string { test.cs.Fprint(buf, test.in); return "" },
		fCSFprintf: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			test.cs.Fprintf(buf, test.format, test.in)
			return ""
		},
		fCSFprintln: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			test.cs.Fprintln(buf, test.in)
			return ""
		},
		fCSPrint:   runWithStdoutRedirect(func(test spewTest) { test.cs.Print(test.in) }),
		fCSPrintln: runWithStdoutRedirect(func(test spewTest) { test.cs.Println(test.in) }),
		fCSSdump: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(test.cs.Sdump(test.in))
			return ""
		},
		fCSSprint: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(test.cs.Sprint(test.in))
			return ""
		},
		fCSSprintf: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(test.cs.Sprintf(test.format, test.in))
			return ""
		},
		fCSSprintln: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(test.cs.Sprintln(test.in))
			return ""
		},
		fCSErrorf: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(test.cs.Errorf(test.format, test.in).Error())
			return ""
		},
		fCSNewFormatter: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			fmt.Fprintf(buf, test.format, test.cs.NewFormatter(test.in))
			return ""
		},
		fErrorf: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(spew.Errorf(test.format, test.in).Error())
			return ""
		},
		fFprint:   func(_ *testing.T, buf *bytes.Buffer, test spewTest) string { spew.Fprint(buf, test.in); return "" },
		fFprintln: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string { spew.Fprintln(buf, test.in); return "" },
		fPrint:    runWithStdoutRedirect(func(test spewTest) { spew.Print(test.in) }),
		fPrintln:  runWithStdoutRedirect(func(test spewTest) { spew.Println(test.in) }),
		fSdump: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(spew.Sdump(test.in))
			return ""
		},
		fSprint: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(spew.Sprint(test.in))
			return ""
		},
		fSprintf: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(spew.Sprintf(test.format, test.in))
			return ""
		},
		fSprintln: func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
			buf.WriteString(spew.Sprintln(test.in))
			return ""
		},
	}
}

// runWithStdoutRedirect returns a spewTestRunner that captures stdout output.
func runWithStdoutRedirect(fn func(spewTest)) spewTestRunner {
	return func(_ *testing.T, buf *bytes.Buffer, test spewTest) string {
		b, err := redirStdout(func() { fn(test) })
		if err != nil {
			return err.Error()
		}

		buf.Write(b)
		return ""
	}
}

// redirStdout is a helper function to return the standard output from f as a
// byte slice.
func redirStdout(f func()) ([]byte, error) {
	tempFile, err := os.CreateTemp("", "ss-test")
	if err != nil {
		return nil, err
	}
	fileName := tempFile.Name()
	defer os.Remove(fileName) // Ignore error

	origStdout := os.Stdout
	os.Stdout = tempFile
	f()
	os.Stdout = origStdout
	tempFile.Close()

	return os.ReadFile(fileName)
}
