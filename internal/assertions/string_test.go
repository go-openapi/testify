// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"regexp"
	"slices"
	"testing"
)

func TestStringEqual(t *testing.T) {
	t.Parallel()

	i := 0
	for currCase := range stringEqualCases() {
		mock := &bufferT{}

		Equal(mock, currCase.equalWant, currCase.equalGot, currCase.msgAndArgs...)
		Regexp(t, regexp.MustCompile(currCase.want), mock.buf.String(), "Case %d", i)
		i++
	}
}

func TestStringEqualFormatting(t *testing.T) {
	t.Parallel()

	i := 0
	for currCase := range stringEqualFormattingCases() {
		mock := &bufferT{}

		Equal(mock, currCase.equalWant, currCase.equalGot, currCase.msgAndArgs...)
		Regexp(t, regexp.MustCompile(currCase.want), mock.buf.String(), "Case %d", i)
		i++
	}
}

func TestStringRegexp(t *testing.T) {
	t.Parallel()

	// run test cases with all combinations of supported types
	for tc := range stringRegexpCases() {
		t.Run(tc.name, tc.test)
	}

	t.Run("with edge cases", func(t *testing.T) {
		t.Run("with unsupported regexp type", func(t *testing.T) {
			t.Parallel()

			const (
				str = "whatever"
				msg = "expected this invalid call to fail (regexp=%v)"
			)

			mock := new(mockT)

			t.Run("should fail (invalid regexp type)", func(t *testing.T) {
				invalidRex := struct{ a string }{a: "invalid"}

				if Regexp(mock, invalidRex, str) {
					t.Errorf(msg, invalidRex)
				}
				if NotRegexp(mock, invalidRex, str) {
					t.Errorf(msg, invalidRex)
				}
			})

			t.Run("should fail (nil regexp)", func(t *testing.T) {
				invalidRex := []byte(nil)

				if Regexp(mock, invalidRex, str) {
					t.Errorf(msg, invalidRex)
				}
				if NotRegexp(mock, invalidRex, str) {
					t.Errorf(msg, invalidRex)
				}
			})
		})

		t.Run("with fmt.Sprint conversion", func(t *testing.T) {
			t.Parallel()

			const (
				numeric = 1234
				msg     = "expected %q to match %q"
				rex     = "^[0-9]+$"
			)

			mock := new(mockT)

			t.Run("should match string representation of a number", func(t *testing.T) {
				if !Regexp(mock, rex, numeric) {
					t.Errorf(msg, numeric, rex)
				}
				if NotRegexp(mock, rex, numeric) {
					t.Errorf(msg, numeric, rex)
				}
			})
		})
	})
}

// test all Regexp variants with the same input (possibly converted).
//
//nolint:thelper // linter false positive: this is not a helper
func testAllRegexpWithTypes(regString, str string, success, valid bool) func(*testing.T) {
	type (
		// redefined types to check for ~string and ~[]byte type constraints
		MyString string
		MyBytes  []byte
	)

	return func(t *testing.T) {
		t.Run("with all type combinations", func(t *testing.T) {
			// generic version : 5 x 4 combinations of input types
			t.Run("with [string,string]", testAllRegexp[string, string](regString, str, success, valid))
			t.Run("with [string,[]byte]", testAllRegexp[string, []byte](regString, []byte(str), success, valid))
			t.Run("with [string,~string]", testAllRegexp[string, MyString](regString, MyString(str), success, valid))
			t.Run("with [string,~[]byte]", testAllRegexp[string, MyBytes](regString, MyBytes(str), success, valid))
			//
			t.Run("with [[]byte,string]", testAllRegexp[[]byte, string]([]byte(regString), str, success, valid))
			t.Run("with [[]byte,[]byte]", testAllRegexp[[]byte, []byte]([]byte(regString), []byte(str), success, valid))
			t.Run("with [[]byte,~string]", testAllRegexp[[]byte, MyString]([]byte(regString), MyString(str), success, valid))
			t.Run("with [[]byte,~[]byte]", testAllRegexp[[]byte, MyBytes]([]byte(regString), MyBytes(str), success, valid))
			//
			t.Run("with [~string,string]", testAllRegexp[MyString, string](MyString(regString), str, success, valid))
			t.Run("with [~string,[]byte]", testAllRegexp[MyString, []byte](MyString(regString), []byte(str), success, valid))
			t.Run("with [~string,~string]", testAllRegexp[MyString, MyString](MyString(regString), MyString(str), success, valid))
			t.Run("with [~string,~[]byte]", testAllRegexp[MyString, MyBytes](MyString(regString), MyBytes(str), success, valid))
			//
			t.Run("with [~[]byte,string]", testAllRegexp[MyBytes, string](MyBytes(regString), str, success, valid))
			t.Run("with [~[]byte,[]byte]", testAllRegexp[MyBytes, []byte](MyBytes(regString), []byte(str), success, valid))
			t.Run("with [~[]byte,~string]", testAllRegexp[MyBytes, MyString](MyBytes(regString), MyString(str), success, valid))
			t.Run("with [~[]byte,~[]byte]", testAllRegexp[MyBytes, MyBytes](MyBytes(regString), MyBytes(str), success, valid))
			//
			t.Run("with [*regexp.Regexp,string]", testAllRegexp[*regexp.Regexp, string](testRex(regString), str, success, valid))
			t.Run("with [*regexp.Regexp,[]byte]", testAllRegexp[*regexp.Regexp, []byte](testRex(regString), []byte(str), success, valid))
			t.Run("with [*regexp.Regexp,~string]", testAllRegexp[*regexp.Regexp, MyString](testRex(regString), MyString(str), success, valid))
			t.Run("with [*regexp.Regexp,~[]byte]", testAllRegexp[*regexp.Regexp, MyBytes](testRex(regString), MyBytes(str), success, valid))
		})
	}
}

//nolint:thelper // linter false positive: this is not a helper
func testAllRegexp[Rex RegExp, ADoc Text](rx Rex, actual ADoc, success, valid bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if !valid {
			// all assertions fail on invalid regexp
			t.Run("should fail", func(t *testing.T) {
				t.Run("with Regexp", testRegexp(rx, actual, false))
				t.Run("with RegexpT", testRegexpT(rx, actual, false))
				t.Run("with NoRegexp", testNotRegexp(rx, actual, false))
				t.Run("with NoRegexpT", testNotRegexpT(rx, actual, false))
			})

			return
		}

		if success {
			t.Run("should match", func(t *testing.T) {
				t.Run("with Regexp", testRegexp(rx, actual, true))
				t.Run("with RegexpT", testRegexpT(rx, actual, true))
			})

			t.Run("should fail", func(t *testing.T) {
				t.Run("with NoRegexp", testNotRegexp(rx, actual, false))
				t.Run("with NoRegexpT", testNotRegexpT(rx, actual, false))
			})
		} else {
			t.Run("should NOT match", func(t *testing.T) {
				t.Run("with NoRegexp", testNotRegexp(rx, actual, true))
				t.Run("with NoRegexpT", testNotRegexpT(rx, actual, true))
			})

			t.Run("should fail", func(t *testing.T) {
				t.Run("with Regexp", testRegexp(rx, actual, false))
				t.Run("with RegexpT", testRegexpT(rx, actual, false))
			})
		}
	}
}

func testRegexp[Rex RegExp, ADoc Text](rx Rex, str ADoc, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
		res := Regexp(mock, rx, str)
		if res != success {
			if success {
				croakWantMatch(t, rx, str)
				return
			}
			croakWantNotMatch(t, rx, str)
		}
	}
}

func testNotRegexp[Rex RegExp, ADoc Text](rx Rex, str ADoc, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
		res := NotRegexp(mock, rx, str)
		if res != success {
			if success {
				croakWantMatch(t, rx, str)
				return
			}
			croakWantNotMatch(t, rx, str)
		}
	}
}

func testRegexpT[Rex RegExp, ADoc Text](rx Rex, str ADoc, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
		res := RegexpT(mock, rx, str)
		if res != success {
			if success {
				croakWantMatch(t, rx, str)
				return
			}
			croakWantNotMatch(t, rx, str)
		}
	}
}

func testNotRegexpT[Rex RegExp, ADoc Text](rx Rex, str ADoc, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
		res := NotRegexpT(mock, rx, str)
		if res != success {
			if success {
				croakWantMatch(t, rx, str)
				return
			}
			croakWantNotMatch(t, rx, str)
		}
	}
}

func croakWantMatch(t *testing.T, rx any, str any) {
	t.Helper()
	t.Errorf("expected %q to match %q", str, rx)
}

func croakWantNotMatch(t *testing.T, rx, str any) {
	t.Helper()
	t.Errorf("expected %q NOT to match %q", str, rx)
}

func testRex(rex string) *regexp.Regexp {
	rx, _ := compileRegex(rex)
	return rx
}

type stringEqualCase struct {
	equalWant  string
	equalGot   string
	msgAndArgs []any
	want       string
}

func stringEqualCases() iter.Seq[stringEqualCase] {
	return slices.Values([]stringEqualCase{
		{
			equalWant: "hi, \nmy name is",
			equalGot:  "what,\nmy name is",
			want: "\t[a-z]+.go:\\d+: \n" + // NOTE: the exact file name reported should be asserted in integration tests
				"\t+Error Trace:\t\n+" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"hi, \\\\nmy name is\"\n" +
				"\\s+actual\\s+: " + "\"what,\\\\nmy name is\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n\\s+\\++ " +
				"Actual\n" +
				"\\s+@@ -1,2 \\+1,2 @@\n" +
				"\\s+-hi, \n\\s+\\+what,\n" +
				"\\s+my name is",
		},
	})
}

func stringEqualFormattingCases() iter.Seq[stringEqualCase] {
	return slices.Values([]stringEqualCase{
		{
			equalWant: "want",
			equalGot:  "got",
			want: "\t[a-z]+.go:\\d+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n\\s+-+ Expected\n\\s+\\++ " +
				"Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n",
		},
		{
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{"hello, %v!", "world"},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+hello, world!\n",
		},
		{
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{123},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+123\n",
		},
		{
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{struct{ a string }{"hello"}},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+{a:hello}\n",
		},
	})
}

// Values to populate the test harness:
//
// - valid and invalid patterns
// - matching and not matching expressions.
func stringRegexpCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// successful matches
		{"^start (match)", testAllRegexpWithTypes(
			"^start", "start of the line", true, true,
		)},
		{"end$ (match)", testAllRegexpWithTypes(
			"end$", "in the end", true, true,
		)},
		{"end$ (match)", testAllRegexpWithTypes(
			"end$", "in the end", true, true,
		)},
		{"phone number (match)", testAllRegexpWithTypes(
			"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12.34", true, true,
		)},
		// failed matches
		{"start (no match)", testAllRegexpWithTypes(
			"^asdfastart", "Not the start of the line", false, true,
		)},
		{"end$ (no match)", testAllRegexpWithTypes(
			"end$", "in the end.", false, true,
		)},
		{"phone number (no match)", testAllRegexpWithTypes(
			"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12a.34", false, true,
		)},
		// invalid pattern
		{"invalid regexp", testAllRegexpWithTypes(
			"\\C", "whatever", false, false,
		)},
	})
}
