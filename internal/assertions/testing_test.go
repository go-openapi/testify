// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"strings"
	"testing"
)

func TestTestingFailNowWithPlainT(t *testing.T) {
	t.Parallel()
	mock := &mockT{}

	Panics(t, func() {
		FailNow(mock, "failed")
	}, "should panic since mockT is missing FailNow()")
}

func TestTestingFailNowWithFullT(t *testing.T) {
	t.Parallel()
	mock := &mockFailNowT{}

	NotPanics(t, func() {
		FailNow(mock, "failed")
	}, "should call mockT.FailNow() rather than panicking")
}

func TestIndentMessageLines(t *testing.T) {
	t.Parallel()

	for tc := range indentMessageLinesCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := indentMessageLines(tc.message, tc.longestLabel)
			Equal(t, tc.expected, result)
		})
	}
}

func TestParseLabeledOutput(t *testing.T) {
	t.Parallel()

	t.Run("round-trip single label", func(t *testing.T) {
		t.Parallel()

		input := []labeledContent{{"Error", "test message"}}
		output := labeledOutput(input...)
		parsed := parseLabeledOutput(output)

		NotNil(t, parsed)
		Len(t, parsed, 1)
		Equal(t, "Error", parsed[0].label)
		Equal(t, "test message\n", parsed[0].content)
	})

	t.Run("round-trip multiple labels", func(t *testing.T) {
		t.Parallel()

		input := []labeledContent{
			{"Error Trace", "file.go:42"},
			{"Error", "not equal"},
			{"Test", "TestFoo"},
			{"Messages", "extra info"},
		}
		output := labeledOutput(input...)
		parsed := parseLabeledOutput(output)

		NotNil(t, parsed)
		Len(t, parsed, 4)
		Equal(t, "Error Trace", parsed[0].label)
		Equal(t, "Error", parsed[1].label)
		Equal(t, "Test", parsed[2].label)
		Equal(t, "Messages", parsed[3].label)
	})

	t.Run("blank line skipping", func(t *testing.T) {
		t.Parallel()

		// Build output with blank lines injected
		output := "\n\tError:\ttest message\n\n"
		parsed := parseLabeledOutput(output)

		NotNil(t, parsed)
		Len(t, parsed, 1)
		Equal(t, "Error", parsed[0].label)
	})

	t.Run("malformed input returns nil", func(t *testing.T) {
		t.Parallel()

		parsed := parseLabeledOutput("this is not labeled output")
		Nil(t, parsed)
	})
}

func TestMessageFromMsgAndArgs(t *testing.T) {
	t.Parallel()

	for tc := range messageFromMsgAndArgsCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := messageFromMsgAndArgs(tc.args...)
			Equal(t, tc.expected, result)
		})
	}
}

// Tests for envelope infrastructure functions.
func TestLabeledOutput(t *testing.T) {
	t.Parallel()

	for tc := range labeledOutputCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := labeledOutput(tc.input...)
			Equal(t, tc.expected, result)
		})
	}
}

func TestErrorEnvelopeIntegration(t *testing.T) {
	t.Parallel()

	mock := new(captureT)
	Fail(mock, "test message")

	True(t, mock.failed, "Fail should mark test as failed")

	parsed := parseLabeledOutput(mock.msg)
	NotNil(t, parsed, "Fail output should be parseable by parseLabeledOutput")

	var hasErrorTrace, hasError bool
	var errorContent string
	for _, lc := range parsed {
		switch lc.label {
		case "Error Trace":
			hasErrorTrace = true
		case "Error":
			hasError = true
			errorContent = strings.TrimRight(lc.content, "\n")
		}
	}

	True(t, hasErrorTrace, "envelope should contain Error Trace label")
	True(t, hasError, "envelope should contain Error label")
	Equal(t, "test message", errorContent)
}

// =======================================
// TestLabeledOutput
// =======================================

type labeledOutputCase struct {
	name     string
	input    []labeledContent
	expected string
}

func labeledOutputCases() iter.Seq[labeledOutputCase] {
	return slices.Values([]labeledOutputCase{
		{
			name:     "single label",
			input:    []labeledContent{{"Error", "something failed"}},
			expected: "\tError:\tsomething failed\n",
		},
		{
			name: "multiple labels aligned",
			input: []labeledContent{
				{"Error Trace", "file.go:42"},
				{"Error", "not equal"},
				{"Test", "TestFoo"},
			},
			expected: "\tError Trace:\tfile.go:42\n" +
				"\tError:      \tnot equal\n" +
				"\tTest:       \tTestFoo\n",
		},
		{
			name:  "multi-line content indented",
			input: []labeledContent{{"Error", "line1\nline2\nline3"}},
			expected: "\tError:\tline1\n" +
				"\t      \tline2\n" +
				"\t      \tline3\n",
		},
		{
			name:     "empty content",
			input:    []labeledContent{{"Label", ""}},
			expected: "\tLabel:\t\n",
		},
	})
}

// =======================================
// TestIndentMessageLines
// =======================================

type indentMessageLinesCase struct {
	name         string
	message      string
	longestLabel int
	expected     string
}

func indentMessageLinesCases() iter.Seq[indentMessageLinesCase] {
	return slices.Values([]indentMessageLinesCase{
		{
			name:         "single line unchanged",
			message:      "hello world",
			longestLabel: 5,
			expected:     "hello world",
		},
		{
			name:         "multi-line indented",
			message:      "line1\nline2\nline3",
			longestLabel: 5,
			expected:     "line1\n\t      \tline2\n\t      \tline3",
		},
		{
			name:         "empty string",
			message:      "",
			longestLabel: 5,
			expected:     "",
		},
		{
			name:         "trailing newline",
			message:      "hello\n",
			longestLabel: 3,
			expected:     "hello",
		},
	})
}

// =======================================
// TestMessageFromMsgAndArgs
// =======================================

type messageFromMsgAndArgsCase struct {
	name     string
	args     []any
	expected string
}

func messageFromMsgAndArgsCases() iter.Seq[messageFromMsgAndArgsCase] {
	return slices.Values([]messageFromMsgAndArgsCase{
		{
			name:     "no args",
			args:     nil,
			expected: "",
		},
		{
			name:     "single string",
			args:     []any{"hello"},
			expected: "hello",
		},
		{
			name:     "single non-string",
			args:     []any{42},
			expected: "42",
		},
		{
			name:     "format string with args",
			args:     []any{"value is %d", 42},
			expected: "value is 42",
		},
		{
			name:     "format string with multiple args",
			args:     []any{"%s=%d", "x", 5},
			expected: "x=5",
		},
	})
}
