// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"regexp"
	"slices"
	"testing"
)

func TestStringRegexp(t *testing.T) {
	t.Parallel()

	// run test cases with all combinations of supported types
	//
	// NOTE: testing pattern, focused on the expected result (true/false) and _NOT_ the content of the returned message.
	// - stringRegexpCases: loop over generic test cases
	//    - testAllRegexpWithTypes: dispatch over type combinations of values
	//      - testAllRegexp: dispatch over the assertion variants (reflection-based, generic, X vs NotX semantics)
	//        Single assertion test functions:
	//        - testRegexp
	//        - testRegexpT
	//        - testNotRegexp
	//        - testNotRegexpT
	for tc := range stringRegexpCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestStringErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, stringFailCases())
}

// =======================================
// TestStringRegexp
// =======================================

// Values to populate the test harness:
//
// - valid and invalid patterns
// - matching and not matching expressions.
func stringRegexpCases() iter.Seq[genericTestCase] {
	const (
		numeric = 1234
		numRex  = "^[0-9]+$"
	)
	invalidRex := struct{ a string }{a: "invalid"}
	nilRex := []byte(nil)

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
		// invalid type
		{"invalid regexp type/struct", testRegexpWithAny(
			invalidRex, "whatever", false, false,
		)},
		{"invalid regexp type/nil", testRegexpWithAny(
			nilRex, "whatever", false, false,
		)},
		{"invalid regexp type/slice-int", testRegexpWithAny(
			[]int{1, 2}, "whatever", false, false,
		)},
		// types that uses fmt.Print
		{"use-fmt-print/slice-int", testRegexpWithAny(
			"^\\[1", []int{1, 2}, true, true,
		)},
		{"use-fmt-print/slice-rune", testRegexpWithAny(
			"^\\[65", []rune{'A', 'B'}, true, true,
		)},
		{"use-fmt-print/numeric", testRegexpWithAny(
			numRex, numeric, true, true,
		)},
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

// testRegexpWithAny tests edge cases that are specific to the reflection-based variants.
//
//nolint:thelper // linter false positive: this is not a helper
func testRegexpWithAny(rx any, actual any, success, valid bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if !valid {
			t.Run("should fail", func(t *testing.T) {
				t.Run("with Regexp", testRegexp(rx, actual, false))
				t.Run("with NoRegexp", testNotRegexp(rx, actual, false))
			})

			return
		}

		if success {
			t.Run("should match", func(t *testing.T) {
				t.Run("with Regexp", testRegexp(rx, actual, true))
			})
			t.Run("should fail", func(t *testing.T) {
				t.Run("with NoRegexp", testNotRegexp(rx, actual, false))
			})
			return
		}

		t.Run("should NOT match", func(t *testing.T) {
			t.Run("with NoRegexp", testNotRegexp(rx, actual, true))
		})

		t.Run("should fail", func(t *testing.T) {
			t.Run("with Regexp", testRegexp(rx, actual, false))
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
				t.Run("with NoRegexp", testNotRegexp(rx, actual, false))

				t.Run("with RegexpT", testRegexpT(rx, actual, false))
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

			return
		}

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

func testRegexp(rx any, str any, success bool) func(*testing.T) {
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

func testNotRegexp(rx any, str any, success bool) func(*testing.T) {
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

// =======================================
// TestStringErrorMessages
// =======================================

func stringFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "Regexp/no-match",
			assertion:    func(t T) bool { return Regexp(t, "^start", "no match") },
			wantContains: []string{"Expect", "to match"},
		},
		{
			name:         "NotRegexp/unexpected-match",
			assertion:    func(t T) bool { return NotRegexp(t, "^start", "starting") },
			wantContains: []string{"Expect", "to NOT match"},
		},
		{
			name:         "Regexp/invalid-regexp",
			assertion:    func(t T) bool { return Regexp(t, "\\C", "whatever") },
			wantContains: []string{"invalid error expression"},
		},
		{
			name:         "Regexp/invalid-type",
			assertion:    func(t T) bool { return Regexp(t, struct{ a string }{a: "x"}, "whatever") },
			wantContains: []string{"type for regexp is not supported"},
		},
	})
}
