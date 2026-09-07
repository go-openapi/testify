// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"strings"
	"testing"
)

// hunkFixture is a 10-element slice with a single difference on the last element.
// A diff of expected against actual therefore has one hunk, whose size is driven by [WithHunkSize].
func hunkFixture() (expected, actual []string) {
	expected = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	actual = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "z"}

	return expected, actual
}

// countContextLines counts the unchanged lines of a unified diff, i.e. the lines prefixed with a space.
func countContextLines(unified string) int {
	var count int

	for line := range strings.SplitSeq(unified, "\n") {
		if strings.HasPrefix(line, " ") && strings.TrimSpace(line) != "" {
			count++
		}
	}

	return count
}

func TestWithHunkSize(t *testing.T) {
	t.Parallel()

	expected, actual := hunkFixture()

	t.Run("hunk size should widen the diff context", func(t *testing.T) {
		t.Parallel()

		// the changed element is last, so the widening only shows up before it:
		// context lines = min(hunkSize, 9 preceding lines) + the closing "}" line.
		for _, toPin := range []struct {
			name         string
			hunkSize     int
			wantContext  int
			wantLastLine string
		}{
			{name: "default (unset) clamps to 1", hunkSize: 0, wantContext: 2, wantLastLine: `(string) (len=1) "i",`},
			{name: "negative clamps to 1", hunkSize: -5, wantContext: 2, wantLastLine: `(string) (len=1) "i",`},
			{name: "explicit 1", hunkSize: 1, wantContext: 2, wantLastLine: `(string) (len=1) "i",`},
			{name: "3 shows two more lines", hunkSize: 3, wantContext: 4, wantLastLine: `(string) (len=1) "g",`},
			{name: "5 shows four more lines", hunkSize: 5, wantContext: 6, wantLastLine: `(string) (len=1) "e",`},
			{name: "larger than the value shows it whole", hunkSize: 100, wantContext: 11, wantLastLine: `([]string) (len=10) {`},
		} {
			t.Run(toPin.name, func(t *testing.T) {
				t.Parallel()

				o := buildOptions(WithHunkSize(toPin.hunkSize))
				unified := diff(expected, actual, o)

				if got := countContextLines(unified); got != toPin.wantContext {
					t.Errorf("expected %d context lines with hunk size %d, got %d in:\n%s",
						toPin.wantContext, toPin.hunkSize, got, unified)
				}

				// the first context line tells how far back the hunk reaches
				lines := strings.Split(strings.TrimPrefix(unified, "\n\nDiff:\n"), "\n")
				first := strings.TrimSpace(lines[3]) // skip "--- Expected", "+++ Actual", "@@ ... @@"
				if first != toPin.wantLastLine {
					t.Errorf("expected the hunk to start at %q, got %q in:\n%s", toPin.wantLastLine, first, unified)
				}
			})
		}
	})

	t.Run("a wider hunk should widen the reported failure", func(t *testing.T) {
		t.Parallel()

		narrow := new(mockT)
		Equal(narrow, expected, actual, buildOptions(WithHunkSize(1)))

		wide := new(mockT)
		Equal(wide, expected, actual, buildOptions(WithHunkSize(5)))

		if !narrow.Failed() || !wide.Failed() {
			t.Fatal("expected both assertions to fail")
		}

		narrowLines := strings.Count(narrow.errorString(), "\n")
		wideLines := strings.Count(wide.errorString(), "\n")

		if wideLines <= narrowLines {
			t.Errorf("expected the wider hunk to report more lines, got %d with hunk size 5 and %d with hunk size 1",
				wideLines, narrowLines)
		}
	})

	t.Run("the injected option should not show up in the message", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		Equal(mock, expected, actual, buildOptions(WithHunkSize(5)))

		if strings.Contains(mock.errorString(), "hunkSize") {
			t.Errorf("expected the option to be stripped from the message, got:\n%s", mock.errorString())
		}

		if strings.Contains(mock.errorString(), "Messages:") {
			t.Errorf("expected no Messages section when only an option is passed, got:\n%s", mock.errorString())
		}
	})

	t.Run("a message passed alongside an option should survive", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		Equal(mock, expected, actual, "context %d", 42, buildOptions(WithHunkSize(5)))

		if !strings.Contains(mock.errorString(), "context 42") {
			t.Errorf("expected the formatted message to be preserved, got:\n%s", mock.errorString())
		}
	})
}

func TestSplitArgs(t *testing.T) {
	t.Parallel()

	t.Run("should leave plain arguments alone", func(t *testing.T) {
		t.Parallel()

		msgAndArgs, opts := splitArgs([]any{"msg", 1, true})

		if len(msgAndArgs) != 3 {
			t.Errorf("expected 3 arguments, got %d: %v", len(msgAndArgs), msgAndArgs)
		}

		if opts.hunkSize != 0 {
			t.Errorf("expected a zero hunk size, got %d", opts.hunkSize)
		}
	})

	t.Run("should extract the option and keep the message", func(t *testing.T) {
		t.Parallel()

		args := []any{"msg %d", 1, buildOptions(WithHunkSize(7))}
		msgAndArgs, opts := splitArgs(args)

		if opts.hunkSize != 7 {
			t.Errorf("expected a hunk size of 7, got %d", opts.hunkSize)
		}

		if len(msgAndArgs) != 2 || msgAndArgs[0] != "msg %d" || msgAndArgs[1] != 1 {
			t.Errorf("expected the message arguments to be preserved, got %v", msgAndArgs)
		}

		if len(args) != 3 {
			t.Errorf("expected the input slice to be left untouched, got %v", args)
		}
	})

	t.Run("should extract an option in any position", func(t *testing.T) {
		t.Parallel()

		msgAndArgs, opts := splitArgs([]any{buildOptions(WithHunkSize(2)), "msg"})

		if opts.hunkSize != 2 {
			t.Errorf("expected a hunk size of 2, got %d", opts.hunkSize)
		}

		if len(msgAndArgs) != 1 || msgAndArgs[0] != "msg" {
			t.Errorf("expected only the message to remain, got %v", msgAndArgs)
		}
	})

	t.Run("should keep the last option when several are injected", func(t *testing.T) {
		t.Parallel()

		msgAndArgs, opts := splitArgs([]any{buildOptions(WithHunkSize(2)), buildOptions(WithHunkSize(9))})

		if opts.hunkSize != 9 {
			t.Errorf("expected the last option to win with a hunk size of 9, got %d", opts.hunkSize)
		}

		if len(msgAndArgs) != 0 {
			t.Errorf("expected all options to be removed, got %v", msgAndArgs)
		}
	})
}

// buildOptions is the test-side shorthand for [BuildOptions], which returns an opaque any.
func buildOptions(opts ...Option) options {
	return BuildOptions(opts).(options) //nolint:forcetypeassert // BuildOptions always returns options
}
