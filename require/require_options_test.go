// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require

import (
	"fmt"
	"strings"
	"testing"
)

// capturingT is a [T] that keeps the last failure message and absorbs FailNow.
type capturingT struct {
	message string
	failed  bool
}

func (capturingT) Helper() {}

func (m *capturingT) Errorf(format string, args ...any) {
	m.message = fmt.Sprintf(format, args...)
}

func (m *capturingT) FailNow() {
	m.failed = true
}

// TestNewWithHunkSize demonstrates how [WithHunkSize] widens the diff reported by [Assertions.Equal].
//
// The two values differ on their last element only, so the diff holds a single hunk and the
// option decides how many unchanged lines precede the change.
func TestNewWithHunkSize(t *testing.T) {
	t.Parallel()

	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	actual := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "z"}

	for _, toPin := range []struct {
		name        string
		options     []Option
		wantContext int    // unchanged lines of the diff
		wantFirst   string // first unchanged line of the hunk
	}{
		{name: "default", options: nil, wantContext: 2, wantFirst: `(string) (len=1) "i",`},
		{name: "hunk size 1", options: []Option{WithHunkSize(1)}, wantContext: 2, wantFirst: `(string) (len=1) "i",`},
		{name: "hunk size 3", options: []Option{WithHunkSize(3)}, wantContext: 4, wantFirst: `(string) (len=1) "g",`},
		{name: "hunk size 5", options: []Option{WithHunkSize(5)}, wantContext: 6, wantFirst: `(string) (len=1) "e",`},
		{name: "hunk size 100", options: []Option{WithHunkSize(100)}, wantContext: 11, wantFirst: `([]string) (len=10) {`},
	} {
		t.Run(toPin.name, func(t *testing.T) {
			t.Parallel()

			mock := new(capturingT)
			a := New(mock, toPin.options...)

			a.Equal(expected, actual)

			if !mock.failed {
				t.Fatal("Equal should call FailNow() on different values")
			}

			if got := countDiffContextLines(mock.message); got != toPin.wantContext {
				t.Errorf("expected %d unchanged lines in the diff, got %d in:\n%s",
					toPin.wantContext, got, mock.message)
			}

			if got := firstDiffContextLine(mock.message); got != toPin.wantFirst {
				t.Errorf("expected the hunk to start at %q, got %q in:\n%s",
					toPin.wantFirst, got, mock.message)
			}

			if strings.Contains(mock.message, "hunkSize") || strings.Contains(mock.message, "Messages:") {
				t.Errorf("expected the option not to leak into the message, got:\n%s", mock.message)
			}
		})
	}

	t.Run("the format variant should keep its message", func(t *testing.T) {
		t.Parallel()

		mock := new(capturingT)
		a := New(mock, WithHunkSize(5))

		a.Equalf(expected, actual, "values differ at index %d", 9)

		if !mock.failed {
			t.Fatal("Equalf should call FailNow() on different values")
		}

		if !strings.Contains(mock.message, "values differ at index 9") {
			t.Errorf("expected the formatted message to be reported, got:\n%s", mock.message)
		}

		if got := countDiffContextLines(mock.message); got != 6 {
			t.Errorf("expected 6 unchanged lines in the diff, got %d in:\n%s", got, mock.message)
		}
	})
}

// diffLines returns the lines of the "Diff:" section of a failure message, stripped from
// the "\t<padding>\t" indentation that labeledOutput adds to continuation lines.
func diffLines(message string) []string {
	_, unified, found := strings.Cut(message, "Diff:\n")
	if !found {
		return nil
	}

	var lines []string

	for line := range strings.SplitSeq(unified, "\n") {
		parts := strings.SplitN(line, "\t", 3)
		if len(parts) < 3 {
			continue
		}

		lines = append(lines, parts[2])
	}

	return lines
}

// countDiffContextLines counts the unchanged lines of the diff, i.e. those prefixed with a space.
func countDiffContextLines(message string) int {
	var count int

	for _, line := range diffLines(message) {
		if strings.HasPrefix(line, " ") && strings.TrimSpace(line) != "" {
			count++
		}
	}

	return count
}

// firstDiffContextLine returns the first unchanged line of the diff, which tells how far
// back the hunk reaches.
func firstDiffContextLine(message string) string {
	for _, line := range diffLines(message) {
		if strings.HasPrefix(line, " ") && strings.TrimSpace(line) != "" {
			return strings.TrimSpace(line)
		}
	}

	return ""
}
