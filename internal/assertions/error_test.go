// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"io"
	"iter"
	"slices"
	"testing"
)

func TestErrorNotErrorAs(t *testing.T) {
	t.Parallel()

	for tt := range errorNotErrorAsCases() {
		t.Run(fmt.Sprintf("NotErrorAs(%#v,%#v)", tt.err, &customError{}), func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)
			var target *customError

			res := NotErrorAs(mock, tt.err, &target)
			shouldPassOrFail(t, mock, res, tt.result)
		})
	}
}

func TestErrorErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, errorFailCases())
}

func TestErrorIs(t *testing.T) {
	t.Parallel()

	for tt := range errorIsCases() {
		t.Run(fmt.Sprintf("ErrorIs(%#v,%#v)", tt.err, tt.target), func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			res := ErrorIs(mock, tt.err, tt.target)
			shouldPassOrFail(t, mock, res, tt.result)
		})
	}
}

func TestErrorNotErrorIs(t *testing.T) {
	t.Parallel()

	for tt := range errorNotErrorIsCases() {
		t.Run(fmt.Sprintf("NotErrorIs(%#v,%#v)", tt.err, tt.target), func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			res := NotErrorIs(mock, tt.err, tt.target)
			shouldPassOrFail(t, mock, res, tt.result)
		})
	}
}

func TestErrorAs(t *testing.T) {
	t.Parallel()

	for tt := range errorAsCases() {
		t.Run(fmt.Sprintf("ErrorAs(%#v,%#v)", tt.err, &customError{}), func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)
			var target *customError

			res := ErrorAs(mock, tt.err, &target)
			shouldPassOrFail(t, mock, res, tt.result)
		})
	}
}

func TestErrorNoError(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	// start with a nil error
	var err error

	if !NoError(mock, err) {
		t.Error("NoError should return true for nil arg")
	}

	// now set an error
	err = errors.New("some error")

	if NoError(mock, err) {
		t.Error("NoError with error should return false")
	}

	// returning an empty error interface
	err = func() error {
		var err *customError
		return err
	}()

	if err == nil { // err is not nil here!
		t.Errorf("Error should be nil due to empty interface: %s", err)
	}

	if NoError(mock, err) {
		t.Error("NoError should fail with empty error interface")
	}
}

func TestError(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	// start with a nil error
	var err error

	if Error(mock, err) {
		t.Error("Error should return false for nil arg")
	}

	// now set an error
	err = errors.New("some error")

	if !Error(mock, err) {
		t.Error("Error with error should return true")
	}

	// returning an empty error interface
	err = func() error {
		var err *customError
		return err
	}()

	if err == nil { // err is not nil here!
		t.Errorf("Error should be nil due to empty interface: %s", err)
	}

	if !Error(mock, err) {
		t.Error("Error should pass with empty error interface")
	}
}

func TestErrorEqualError(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	// start with a nil error
	var err error
	if EqualError(mock, err, "") {
		t.Error("EqualError should return false for nil arg")
	}

	// now set an error
	err = errors.New("some error")
	if EqualError(mock, err, "Not some error") {
		t.Error("EqualError should return false for different error string")
	}
	if !EqualError(mock, err, "some error") {
		t.Error("EqualError should return true")
	}
}

func TestErrorContains(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	// start with a nil error
	var err error
	if ErrorContains(mock, err, "") {
		t.Error("ErrorContains should return false for nil arg")
	}

	// now set an error
	err = errors.New("some error: another error")
	if ErrorContains(mock, err, "bad error") {
		t.Error("ErrorContains should return false for different error string")
	}
	if !ErrorContains(mock, err, "some error") {
		t.Error("ErrorContains should return true for 'some error'")
	}
	if !ErrorContains(mock, err, "another error") {
		t.Error("ErrorContains should return true for 'another error'")
	}
}

// ============================================================================
// TestNotErrorAs
// ============================================================================

type errorNotErrorAsCase struct {
	err    error
	result bool
}

func errorNotErrorAsCases() iter.Seq[errorNotErrorAsCase] {
	return slices.Values([]errorNotErrorAsCase{
		{
			err:    fmt.Errorf("wrap: %w", &customError{}),
			result: false,
		},
		{
			err:    io.EOF,
			result: true,
		},
		{
			err:    nil,
			result: true,
		},
	})
}

type errorIsCase struct {
	err    error
	target error
	result bool
}

func errorIsCases() iter.Seq[errorIsCase] {
	return slices.Values([]errorIsCase{
		{err: io.EOF, target: io.EOF, result: true},
		{err: fmt.Errorf("wrap: %w", io.EOF), target: io.EOF, result: true},
		{err: io.EOF, target: io.ErrClosedPipe, result: false},
		{err: nil, target: io.EOF, result: false},
		{err: io.EOF, target: nil, result: false},
		{err: nil, target: nil, result: true},
		{err: fmt.Errorf("abc: %w", errors.New("def")), target: io.EOF, result: false},
	})
}

type errorNotErrorIsCase struct {
	err    error
	target error
	result bool
}

func errorNotErrorIsCases() iter.Seq[errorNotErrorIsCase] {
	return slices.Values([]errorNotErrorIsCase{
		{err: io.EOF, target: io.EOF, result: false},
		{err: fmt.Errorf("wrap: %w", io.EOF), target: io.EOF, result: false},
		{err: io.EOF, target: io.ErrClosedPipe, result: true},
		{err: nil, target: io.EOF, result: true},
		{err: io.EOF, target: nil, result: true},
		{err: nil, target: nil, result: false},
		{err: fmt.Errorf("abc: %w", errors.New("def")), target: io.EOF, result: true},
	})
}

type errorAsCase struct {
	err    error
	result bool
}

func errorAsCases() iter.Seq[errorAsCase] {
	return slices.Values([]errorAsCase{
		{err: fmt.Errorf("wrap: %w", &customError{}), result: true},
		{err: io.EOF, result: false},
		{err: nil, result: false},
		{err: fmt.Errorf("abc: %w", errors.New("def")), result: false},
	})
}

type customError struct{}

func (*customError) Error() string { return "fail" }

// ============================================================================
// TestErrorErrorMessages
// ============================================================================

func errorFailCases() iter.Seq[failCase] {
	longSlice := make([]int, 1_000_000)

	return slices.Values([]failCase{
		// --- truncation cases ---
		truncationCase("NoError/truncation", func(t T) bool {
			return NoError(t, fmt.Errorf("long: %v", longSlice))
		}),
		truncationCase("EqualError/truncation", func(t T) bool {
			return EqualError(t, fmt.Errorf("long: %v", longSlice), "EOF")
		}),
		truncationCase("ErrorContains/truncation", func(t T) bool {
			return ErrorContains(t, fmt.Errorf("long: %v", longSlice), "EOF")
		}),
		truncationCase("ErrorIs/truncation", func(t T) bool {
			return ErrorIs(t, fmt.Errorf("long: %v", longSlice), fmt.Errorf("also: %v", longSlice))
		}),
		truncationCase("NotErrorIs/truncation", func(t T) bool {
			err := fmt.Errorf("long: %v", longSlice)
			return NotErrorIs(t, err, err)
		}),
		truncationCase("ErrorAs/truncation", func(t T) bool {
			var target *customError
			return ErrorAs(t, fmt.Errorf("long: %v", longSlice), &target)
		}),
		truncationCase("NotErrorAs/truncation", func(t T) bool {
			var target *customError
			return NotErrorAs(t, fmt.Errorf("long: %v %w", longSlice, &customError{}), &target)
		}),

		// --- ErrorIs message cases ---
		{
			name: "ErrorIs/not_in_chain",
			assertion: func(t T) bool {
				return ErrorIs(t, io.EOF, io.ErrClosedPipe)
			},
			wantError: "" +
				"Target error should be in err chain:\n" +
				"expected: \"io: read/write on closed pipe\"\n" +
				"in chain: \"EOF\"",
		},
		{
			name: "ErrorIs/nil_err",
			assertion: func(t T) bool {
				return ErrorIs(t, nil, io.EOF)
			},
			wantError: "Expected error with \"EOF\" in chain but got nil.",
		},
		{
			name: "ErrorIs/nil_target",
			assertion: func(t T) bool {
				return ErrorIs(t, io.EOF, nil)
			},
			wantError: "" +
				"Target error should be in err chain:\n" +
				"expected: \"\"\n" +
				"in chain: \"EOF\"",
		},
		{
			name: "ErrorIs/wrapped_not_in_chain",
			assertion: func(t T) bool {
				return ErrorIs(t, fmt.Errorf("abc: %w", errors.New("def")), io.EOF)
			},
			wantError: "" +
				"Target error should be in err chain:\n" +
				"expected: \"EOF\"\n" +
				"in chain: \"abc: def\"\n" +
				"\t\"def\"",
		},

		// --- NotErrorIs message cases ---
		{
			name: "NotErrorIs/same_error",
			assertion: func(t T) bool {
				return NotErrorIs(t, io.EOF, io.EOF)
			},
			wantError: "" +
				"Target error should not be in err chain:\n" +
				"found: \"EOF\"\n" +
				"in chain: \"EOF\"",
		},
		{
			name: "NotErrorIs/wrapped_in_chain",
			assertion: func(t T) bool {
				return NotErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
			},
			wantError: "" +
				"Target error should not be in err chain:\n" +
				"found: \"EOF\"\n" +
				"in chain: \"wrap: EOF\"\n" +
				"\t\"EOF\"",
		},
		{
			name: "NotErrorIs/both_nil",
			assertion: func(t T) bool {
				return NotErrorIs(t, nil, nil)
			},
			wantError: "" +
				"Target error should not be in err chain:\n" +
				"found: \"\"\n" +
				"in chain: ",
		},

		// --- ErrorAs message cases ---
		{
			name: "ErrorAs/not_in_chain",
			assertion: func(t T) bool {
				var target *customError
				return ErrorAs(t, io.EOF, &target)
			},
			wantError: "" +
				"Should be in error chain:\n" +
				fmt.Sprintf("expected: *%s.customError\n", shortpkg) +
				"in chain: \"EOF\" (*errors.errorString)",
		},
		{
			name: "ErrorAs/nil_err",
			assertion: func(t T) bool {
				var target *customError
				return ErrorAs(t, nil, &target)
			},
			wantError: "" +
				"An error is expected but got nil.\n" +
				fmt.Sprintf("expected: *%s.customError", shortpkg),
		},
		{
			name: "ErrorAs/wrapped_not_in_chain",
			assertion: func(t T) bool {
				var target *customError
				return ErrorAs(t, fmt.Errorf("abc: %w", errors.New("def")), &target)
			},
			wantError: "" +
				"Should be in error chain:\n" +
				fmt.Sprintf("expected: *%s.customError\n", shortpkg) +
				"in chain: \"abc: def\" (*fmt.wrapError)\n" +
				"\t\"def\" (*errors.errorString)",
		},

		// --- NotErrorAs message cases ---
		{
			name: "NotErrorAs/found_in_chain",
			assertion: func(t T) bool {
				var target *customError
				return NotErrorAs(t, fmt.Errorf("wrap: %w", &customError{}), &target)
			},
			wantError: "" +
				"Target error should not be in err chain:\n" +
				fmt.Sprintf("found: *%s.customError\n", shortpkg) +
				"in chain: \"wrap: fail\" (*fmt.wrapError)\n" +
				fmt.Sprintf("\t\"fail\" (*%s.customError)", shortpkg),
		},
		// -- TestExample error
		{
			name: "NotError/TestExampleError",
			assertion: func(t T) bool {
				return NoError(t, ErrTest)
			},
			wantError: "Received unexpected error:\n" +
				"assert.ErrTest general error for testing",
		},
	})
}
