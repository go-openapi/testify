// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

func TestParseTestExamples(t *testing.T) {
	t.Parallel()

	for c := range parseTestExamplesCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTestExamples(c.input)

			if len(result) != len(c.expected) {
				t.Fatalf("ParseTestExamples() returned %d tests, expected %d\nGot: %+v\nExpected: %+v",
					len(result), len(c.expected), result, c.expected)
			}

			for i, expected := range c.expected {
				got := result[i]
				if got.TestedValue != expected.TestedValue {
					t.Errorf("Test[%d].TestedValue = %q, expected %q", i, got.TestedValue, expected.TestedValue)
				}
				if got.ExpectedOutcome != expected.ExpectedOutcome {
					t.Errorf("Test[%d].ExpectedOutcome = %v, expected %v", i, got.ExpectedOutcome, expected.ExpectedOutcome)
				}
				if got.AssertionMessage != expected.AssertionMessage {
					t.Errorf("Test[%d].AssertionMessage = %q, expected %q", i, got.AssertionMessage, expected.AssertionMessage)
				}
			}
		})
	}
}

/* Test case iterators */

type parseTestExamplesCase struct {
	name     string
	input    string
	expected []model.Test
}

func parseTestExamplesCases() iter.Seq[parseTestExamplesCase] {
	return slices.Values([]parseTestExamplesCase{
		{
			name: "success and failure examples",
			input: `Some function description.

Examples:
  success: 123, 456
  failure: 789, 012`,
			expected: []model.Test{
				{TestedValue: "123, 456", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "789, 012", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "panic example",
			input: `Examples:
  panic: nil, "something"
  should panic with nil pointer`,
			expected: []model.Test{
				{
					TestedValue:      "nil, \"something\"",
					ExpectedOutcome:  model.TestPanic,
					AssertionMessage: "should panic with nil pointer",
				},
			},
		},
		{
			name: "mixed examples with assertion messages",
			input: `Examples:
  success: "hello", "hello"
  failure: "hello", "world"
  strings don't match
  panic: invalid, args
  Expected to panic`,
			expected: []model.Test{
				{TestedValue: "\"hello\", \"hello\"", ExpectedOutcome: model.TestSuccess},
				{
					TestedValue:      "\"hello\", \"world\"",
					ExpectedOutcome:  model.TestFailure,
					AssertionMessage: "strings don't match",
				},
				{
					TestedValue:      "invalid, args",
					ExpectedOutcome:  model.TestPanic,
					AssertionMessage: "Expected to panic",
				},
			},
		},
		{
			name: "case insensitive success/failure/panic",
			input: `Examples:
  SUCCESS: 1, 1
  FAILURE: 1, 2
  PANIC: bad`,
			expected: []model.Test{
				{TestedValue: "1, 1", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "1, 2", ExpectedOutcome: model.TestFailure},
				{TestedValue: "bad", ExpectedOutcome: model.TestPanic},
			},
		},
		{
			name: "markdown header style",
			input: `# Examples

  success: true, true
  failure: true, false`,
			expected: []model.Test{
				{TestedValue: "true, true", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "true, false", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "plural form",
			input: `Examples:
  success: 42`,
			expected: []model.Test{
				{TestedValue: "42", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "singular form",
			input: `Example:
  success: 100`,
			expected: []model.Test{
				{TestedValue: "100", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "skip NOT IMPLEMENTED values",
			input: `Examples:
  success: 123
  failure: // NOT IMPLEMENTED
  success: 456`,
			expected: []model.Test{
				{TestedValue: "123", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "456", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "skip empty values",
			input: `Examples:
  success: 123
  failure:
  success: 456`,
			expected: []model.Test{
				{TestedValue: "123", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "456", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "skip whitespace-only values",
			input: `Examples:
  success: 123
  failure:
  success: 456`,
			expected: []model.Test{
				{TestedValue: "123", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "456", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "stop at next section",
			input: `Examples:
  success: 123
  failure: 456

Usage:
  success: this should not be parsed`,
			expected: []model.Test{
				{TestedValue: "123", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "456", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "complex values with quotes and commas",
			input: `Examples:
  success: "hello, world", "hello, world"
  failure: map[string]int{"a": 1}, map[string]int{"b": 2}`,
			expected: []model.Test{
				{TestedValue: `"hello, world", "hello, world"`, ExpectedOutcome: model.TestSuccess},
				{TestedValue: `map[string]int{"a": 1}, map[string]int{"b": 2}`, ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "assertion message on line after value",
			input: `Examples:
  success: 1, 1
  Both values are equal
  failure: 1, 2
  Values differ`,
			expected: []model.Test{
				{
					TestedValue:      "1, 1",
					ExpectedOutcome:  model.TestSuccess,
					AssertionMessage: "Both values are equal",
				},
				{
					TestedValue:      "1, 2",
					ExpectedOutcome:  model.TestFailure,
					AssertionMessage: "Values differ",
				},
			},
		},
		{
			name: "no assertion message when next line is value",
			input: `Examples:
  success: 1, 1
  failure: 1, 2`,
			expected: []model.Test{
				{TestedValue: "1, 1", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "1, 2", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "assertion message not consumed by next value",
			input: `Examples:
  success: 1, 1
  Equal values
  success: 2, 2`,
			expected: []model.Test{
				{
					TestedValue:      "1, 1",
					ExpectedOutcome:  model.TestSuccess,
					AssertionMessage: "Equal values",
				},
				{TestedValue: "2, 2", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name:     "no examples section",
			input:    `This is just a description without examples.`,
			expected: []model.Test{},
		},
		{
			name: "empty examples section",
			input: `Examples:

Next section.`,
			expected: []model.Test{},
		},
		{
			name: "examples section with no valid values",
			input: `Examples:
  Just some text here.
  No actual test values.`,
			expected: []model.Test{},
		},
		{
			name: "leading and trailing whitespace in values",
			input: `Examples:
  success:   123, 456
  failure:  "hello"  `,
			expected: []model.Test{
				{TestedValue: "123, 456", ExpectedOutcome: model.TestSuccess},
				{TestedValue: `"hello"`, ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "multiple assertion messages only first consumed",
			input: `Examples:
  success: 1
  First message
  Second message should not be consumed
  failure: 2`,
			expected: []model.Test{
				{
					TestedValue:      "1",
					ExpectedOutcome:  model.TestSuccess,
					AssertionMessage: "First message",
				},
				{TestedValue: "2", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "assertion message with special characters",
			input: `Examples:
  failure: nil, "test"
  Expected: non-nil value, got: <nil>`,
			expected: []model.Test{
				{
					TestedValue:      `nil, "test"`,
					ExpectedOutcome:  model.TestFailure,
					AssertionMessage: "Expected: non-nil value, got: <nil>",
				},
			},
		},
		{
			name: "case insensitive example header",
			input: `EXAMPLES:
  success: 1`,
			expected: []model.Test{
				{TestedValue: "1", ExpectedOutcome: model.TestSuccess},
			},
		},
		{
			name: "values with leading whitespace",
			input: `Examples:
    success: 123
    failure: 456`,
			expected: []model.Test{
				{TestedValue: "123", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "456", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "TODO with description",
			input: `Examples:
  success: // NOT IMPLEMENTED: implement this test
  failure: 456`,
			expected: []model.Test{
				{TestedValue: "456", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "empty line between examples",
			input: `Examples:
  success: 1

  failure: 2`,
			expected: []model.Test{
				{TestedValue: "1", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "2", ExpectedOutcome: model.TestFailure},
			},
		},
		{
			name: "assertion message separated by empty line",
			input: `Examples:
  success: 1

  Some text here should not be an assertion message
  failure: 2`,
			expected: []model.Test{
				{TestedValue: "1", ExpectedOutcome: model.TestSuccess},
				{TestedValue: "2", ExpectedOutcome: model.TestFailure},
			},
		},
	})
}
