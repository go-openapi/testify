// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"fmt"
	"iter"
	"regexp"
	"strings"
	"testing"
)

var (
	_ T         = &mockT{}
	_ T         = &mockFailNowT{}
	_ failNower = &mockFailNowT{}
	_ T         = &captureT{}
	_ T         = &errorsCapturingT{}
)

type mockT struct {
	errorFmt string
	args     []any
	failed   bool
}

const (
	errString = "Error"
	errTrace  = "Error Trace"
)

// Helper is like [testing.T.Helper] but does nothing.
func (mockT) Helper() {}

func (m *mockT) Errorf(format string, args ...any) {
	m.errorFmt = format
	m.args = args
	m.failed = true
}

func (m *mockT) Failed() bool {
	return m.errorFmt != "" || m.failed
}

func (m *mockT) errorString() string {
	return fmt.Sprintf(m.errorFmt, m.args...)
}

type mockFailNowT struct {
	mockT
}

func (m *mockFailNowT) FailNow() {
	m.failed = true
}

// errorsCapturingT is a mock implementation of TestingT that captures multiple errors reported with Errorf.
//
// It may be equipped with a [context.Context] for tests that check on the [testing.T.Context].
type errorsCapturingT struct {
	errors []error
	ctx    context.Context //nolint:containedctx // this is ok to support context injection tests
}

// Helper is like [testing.T.Helper] but does nothing.
func (errorsCapturingT) Helper() {}

func (t errorsCapturingT) Context() context.Context {
	if t.ctx == nil {
		return context.Background()
	}

	return t.ctx
}

func (t *errorsCapturingT) WithContext(ctx context.Context) *errorsCapturingT {
	t.ctx = ctx

	return t
}

func (t *errorsCapturingT) Errorf(format string, args ...any) {
	t.errors = append(t.errors, fmt.Errorf(format, args...))
}

type captureT struct {
	failed bool
	msg    string
}

// Helper is like [testing.T.Helper] but does nothing.
func (captureT) Helper() {}

func (ctt *captureT) Errorf(format string, args ...any) {
	ctt.msg = fmt.Sprintf(format, args...)
	ctt.failed = true
}

func shouldPassOrFail(t *testing.T, mock *mockT, result, shouldPass bool) {
	t.Helper()

	if shouldPass {
		if !result || mock.Failed() {
			t.Error("expected to pass")
		}

		return
	}

	if result || !mock.Failed() {
		t.Error("expected to fail")
	}
}

func ptr(i int) *int {
	return &i
}

// failCase defines a test case for verifying assertion error messages.
//
// Only one of wantError, wantMatch, or wantContains should be set per case.
type failCase struct {
	name         string
	assertion    func(t T) bool // assertion call with bad inputs baked in
	wantError    string         // exact match on errString label content
	wantMatch    string         // regexp match on errString label content
	wantContains []string       // substring checks on errString label content
}

// runFailCases runs a set of failCase tests using the standard harness.
func runFailCases(t *testing.T, cases iter.Seq[failCase]) {
	t.Helper()

	for tc := range cases {
		t.Run(tc.name, runFailCase(tc))
	}
}

func runFailCase(tc failCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(captureT)
		result := tc.assertion(mock)

		// 1. Verify assertion returned false
		if result {
			t.Error("expected assertion to return false")
			return
		}

		// 2. Verify mock recorded a failure
		if !mock.failed {
			t.Error("expected mock to record a failure")
			return
		}

		// 3. Parse envelope
		parsed := parseLabeledOutput(mock.msg)
		if parsed == nil {
			t.Errorf("could not parse labeled output: %q", mock.msg)
			return
		}

		// 4. Validate envelope structure
		var hasErrorTrace, hasError bool
		var errorContent string
		for _, lc := range parsed {
			switch lc.label {
			case errTrace:
				hasErrorTrace = true
			case errString:
				hasError = true
				errorContent = strings.TrimRight(lc.content, "\n")
			}
		}
		if !hasErrorTrace {
			t.Error("envelope missing Error Trace label")
		}
		if !hasError {
			t.Error("envelope missing Error label")
			return
		}

		// 5. Match based on which want* field is set
		switch {
		case tc.wantError != "":
			if errorContent != tc.wantError {
				t.Errorf("error content mismatch:\n  want: %q\n  got:  %q", tc.wantError, errorContent)
			}
		case tc.wantMatch != "":
			matched, err := regexp.MatchString(tc.wantMatch, errorContent)
			if err != nil {
				t.Errorf("invalid regexp %q: %v", tc.wantMatch, err)

				return
			}

			if !matched {
				t.Errorf("error content does not match pattern %q:\n  got: %q", tc.wantMatch, errorContent)
			}
		case len(tc.wantContains) > 0:
			for _, sub := range tc.wantContains {
				if !strings.Contains(errorContent, sub) {
					t.Errorf("error content missing substring %q:\n  got: %q", sub, errorContent)
				}
			}
		}
	}
}

// truncationCase is a convenience constructor for a failCase that checks for truncation.
func truncationCase(name string, assertion func(t T) bool) failCase {
	return failCase{
		name:         name,
		assertion:    assertion,
		wantContains: []string{"<... truncated>"},
	}
}
