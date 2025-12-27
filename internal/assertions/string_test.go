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

func TestSringEqualFormatting(t *testing.T) {
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
	mock := new(testing.T)

	for tc := range stringRegexpCases() {
		True(t, Regexp(mock, tc.rx, tc.str))
		True(t, Regexp(mock, regexp.MustCompile(tc.rx), tc.str))
		True(t, Regexp(mock, regexp.MustCompile(tc.rx), []byte(tc.str)))

		False(t, NotRegexp(mock, tc.rx, tc.str))
		False(t, NotRegexp(mock, tc.rx, []byte(tc.str)))
		False(t, NotRegexp(mock, regexp.MustCompile(tc.rx), tc.str))
	}

	for tc := range stringNotRegexpCases() {
		False(t, Regexp(mock, tc.rx, tc.str), "Expected %q to not match %q", tc.rx, tc.str)
		False(t, Regexp(mock, regexp.MustCompile(tc.rx), tc.str))
		False(t, Regexp(mock, regexp.MustCompile(tc.rx), []byte(tc.str)))

		True(t, NotRegexp(mock, tc.rx, tc.str))
		True(t, NotRegexp(mock, tc.rx, []byte(tc.str)))
		True(t, NotRegexp(mock, regexp.MustCompile(tc.rx), tc.str))
	}
}

// Verifies that invalid patterns no longer cause a panic when using Regexp/NotRegexp.
// Instead, the assertion should fail and return false.
func TestStringRegexp_InvalidPattern(t *testing.T) {
	t.Parallel()

	const (
		invalidPattern = "\\C"
		msg            = "whatever"
	)

	t.Run("Regexp should not panic on invalid patterns", func(t *testing.T) {
		result := NotPanics(t, func() {
			mockT := new(testing.T)
			False(t, Regexp(mockT, invalidPattern, msg))
		})
		if !result {
			t.Failed()
		}
	})

	t.Run("NoRegexp should not panic on invalid patterns", func(t *testing.T) {
		result := NotPanics(t, func() {
			mockT := new(testing.T)
			False(t, NotRegexp(mockT, invalidPattern, msg))
		})
		if !result {
			t.Failed()
		}
	})
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

type stringRegexpCase struct {
	rx, str string
}

func stringRegexpCases() iter.Seq[stringRegexpCase] {
	return slices.Values([]stringRegexpCase{
		{"^start", "start of the line"},
		{"end$", "in the end"},
		{"end$", "in the end"},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12.34"},
	})
}

func stringNotRegexpCases() iter.Seq[stringRegexpCase] {
	return slices.Values([]stringRegexpCase{
		{"^asdfastart", "Not the start of the line"},
		{"end$", "in the end."},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12a.34"},
	})
}
