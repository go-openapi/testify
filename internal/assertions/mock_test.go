// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var (
	_ T         = &mockT{}
	_ T         = &mockFailNowT{}
	_ failNower = &mockFailNowT{}
	_ T         = &captureT{}
	_ T         = &bufferT{}
	_ T         = &dummyT{}
	_ T         = &errorsCapturingT{}
	_ T         = &outputT{}
)

type mockT struct {
	errorFmt string
	args     []any
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockT) Helper() {}

func (m *mockT) Errorf(format string, args ...any) {
	m.errorFmt = format
	m.args = args
}

func (m *mockT) Failed() bool {
	return m.errorFmt != ""
}

func (m *mockT) errorString() string {
	return fmt.Sprintf(m.errorFmt, m.args...)
}

type mockFailNowT struct {
	failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockFailNowT) Helper() {}

func (m *mockFailNowT) Errorf(format string, args ...any) {
	_ = format
	_ = args
}

func (m *mockFailNowT) FailNow() {
	m.failed = true
}

type dummyT struct{}

func (dummyT) Errorf(string, ...any) {}

func (dummyT) Context() context.Context {
	return context.Background()
}

// errorsCapturingT is a mock implementation of TestingT that captures errors reported with Errorf.
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

type outputT struct {
	buf     *bytes.Buffer
	helpers map[string]struct{}
}

// Implements T.
func (t *outputT) Errorf(format string, args ...any) {
	s := fmt.Sprintf(format, args...)
	t.buf.WriteString(s)
}

func (t *outputT) Helper() {
	if t.helpers == nil {
		t.helpers = make(map[string]struct{})
	}
	t.helpers[callerName(1)] = struct{}{}
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

func (ctt *captureT) checkResultAndErrMsg(t *testing.T, expectedRes, res bool, expectedErrMsg string) {
	t.Helper()
	if res != expectedRes {
		t.Errorf("Should return %t", expectedRes)
		return
	}
	if res == ctt.failed {
		t.Errorf("The test result (%t) should be reflected in the testing.T type (%t)", res, !ctt.failed)
		return
	}
	contents := parseLabeledOutput(ctt.msg)
	if res == true {
		if contents != nil {
			t.Errorf("Should not log an error. Log output: %q", ctt.msg)
		}
		return
	}
	if contents == nil {
		t.Errorf("Should log an error. Log output: %q", ctt.msg)
		return
	}
	for _, content := range contents {
		if content.label == "Error" {
			if expectedErrMsg == content.content {
				return
			}
			t.Errorf("Recorded Error: %q", content.content)
		}
	}
	t.Errorf("Expected Error: %q", expectedErrMsg)
}

// bufferT implements TestingT. Its implementation of Errorf writes the output that would be produced by
// testing.T.Errorf to an internal bytes.Buffer.
type bufferT struct {
	buf bytes.Buffer
}

// Helper is like [testing.T.Helper] but does nothing.
func (bufferT) Helper() {}

func (t *bufferT) Errorf(format string, args ...any) {
	// implementation of decorate is copied from testing.T
	decorate := func(s string) string {
		_, file, line, ok := runtime.Caller(3) // decorate + log + public function.
		if ok {
			// Truncate file name at last file name separator.
			if index := strings.LastIndex(file, "/"); index >= 0 {
				file = file[index+1:]
			} else if index = strings.LastIndex(file, "\\"); index >= 0 {
				file = file[index+1:]
			}
		} else {
			file = "???"
			line = 1
		}
		buf := new(bytes.Buffer)
		// Every line is indented at least one tab.
		buf.WriteByte('\t')
		fmt.Fprintf(buf, "%s:%d: ", file, line)
		lines := strings.Split(s, "\n")
		if l := len(lines); l > 1 && lines[l-1] == "" {
			lines = lines[:l-1]
		}
		for i, line := range lines {
			if i > 0 {
				// Second and subsequent lines are indented an extra tab.
				buf.WriteString("\n\t\t")
			}
			buf.WriteString(line)
		}
		buf.WriteByte('\n')
		return buf.String()
	}
	t.buf.WriteString(decorate(fmt.Sprintf(format, args...)))
}

// parseLabeledOutput does the inverse of labeledOutput - it takes a formatted
// output string and turns it back into a slice of labeledContent.
func parseLabeledOutput(output string) []labeledContent {
	labelPattern := regexp.MustCompile(`^\t([^\t]*): *\t(.*)$`)
	contentPattern := regexp.MustCompile(`^\t *\t(.*)$`)
	var contents []labeledContent
	lines := strings.Split(output, "\n")
	i := -1
	for _, line := range lines {
		if line == "" {
			// skip blank lines
			continue
		}
		matches := labelPattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			// a label
			contents = append(contents, labeledContent{
				label:   matches[1],
				content: matches[2] + "\n",
			})
			i++
			continue
		}
		matches = contentPattern.FindStringSubmatch(line)
		if len(matches) == 2 {
			// just content
			if i >= 0 {
				contents[i].content += matches[1] + "\n"
				continue
			}
		}
		// Couldn't parse output
		return nil
	}
	return contents
}

// callerName gives the function name (qualified with a package path)
// for the caller after skip frames (where 0 means the current function).
func callerName(skip int) string {
	// Make room for the skip PC.
	var pc [1]uintptr
	n := runtime.Callers(skip+2, pc[:]) // skip + runtime.Callers + callerName
	if n == 0 {
		panic("testing: zero callers found")
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}

type testCase struct {
	expected any
	actual   any
	result   bool
}
