// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"slices"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// ParseTestExamples extracts test examples from doc comment text.
//
// It looks for an "Examples:" or "Example:" section and parses lines like:
//
//   - success: <test values>
//   - failure: <test values>
//   - panic: <assertion message>
//
// After each value line, an optional comment line can provide an assertion message.
// Empty values or values marked "// NOT IMPLEMENTED" are skipped.
func ParseTestExamples(text string) []model.Test {
	const usualNumberOfExamples = 2
	tests := make([]model.Test, 0, usualNumberOfExamples)

	inExamplesSection := false // we expect an Example[s] or # Example[s] section
	inValueSection := false

	const (
		successPrefix = "success"
		failurePrefix = "failure"
		panicPrefix   = "panic"
		examplePrefix = "example"
	)

	startExampleSection := StartSectionFunc(examplePrefix) // e.g. Example:, examples:, # Example, ...
	startValueSuccess := StartValueFunc(successPrefix)     // e.g. success: , Success:, "SuCCess  :"
	startValueFailure := StartValueFunc(failurePrefix)
	startValuePanic := StartValueFunc(panicPrefix)
	startExampleValue := func(line string) (string, model.TestExpectedOutcome, bool) {
		val, ok := startValueSuccess(line)
		if ok {
			return val, model.TestSuccess, true
		}
		val, ok = startValueFailure(line)
		if ok {
			return val, model.TestFailure, true
		}
		val, ok = startValuePanic(line)
		if ok {
			return val, model.TestPanic, true
		}
		return "", model.TestNone, false
	}

	for line := range strings.SplitSeq(text, "\n") {
		line = strings.TrimSpace(line)
		value, outcome, isExampleValue := startExampleValue(line)
		if inValueSection && len(tests) > 0 && line != "" && !StartAnotherSection(line) && !isExampleValue {
			// check if a comment line follows an example value: this would be the assertion message
			tests[len(tests)-1].AssertionMessage = line
			inValueSection = false

			continue
		}
		inValueSection = false

		// check if we're entering the Examples section
		if startExampleSection(line) {
			inExamplesSection = true
			continue
		}

		// skip until we find the Examples section
		if !inExamplesSection {
			continue
		}

		// stop if we hit another section (starts with capital letter and ends with colon)
		if StartAnotherSection(line) {
			break
		}

		// parse test lines: "success: { values to put in the test }", "failure: ..." or "panic: ..."
		// after each value line, we may found an additional text to be used as an assertion message.
		if !isExampleValue {
			continue
		}

		inValueSection = true // an extra comment may appear right after the value line
		testcase, ok := parseTestValue(outcome, value)
		if !ok {
			continue // there is a value, but it's empty or marked "// NOT IMPLEMENTED"
		}

		tests = append(tests, testcase)
	}

	return slices.Clip(tests)
}

func parseTestValue(outcome model.TestExpectedOutcome, value string) (model.Test, bool) {
	value = strings.TrimSpace(value)
	_, isTodo := strings.CutPrefix(value, "// NOT IMPLEMENTED")

	if value == "" || isTodo {
		return model.Test{}, false
	}

	// Parse test values as Go expressions
	parsedValues := ParseTestValues(value)

	return model.Test{
		TestedValue:     value,        // Keep original for backward compatibility
		TestedValues:    parsedValues, // New parsed representation
		ExpectedOutcome: outcome,
	}, true
}
