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

func TestErrorNoErrorWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	NoError(mock, fmt.Errorf("long: %v", longSlice))
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Received unexpected error:
	            	long: [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated>`)
}

func TestErrorEqualErrorWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	EqualError(mock, fmt.Errorf("long: %v", longSlice), "EOF")
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Error message not equal:
	            	expected: "EOF"
	            	actual  : "long: [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated>`)
}

func TestErrorContainsWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	ErrorContains(mock, fmt.Errorf("long: %v", longSlice), "EOF")
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Error "long: [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated> does not contain "EOF"`)
}

func TestErrorIsWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	ErrorIs(mock, fmt.Errorf("long: %v", longSlice), fmt.Errorf("also: %v", longSlice))
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Target error should be in err chain:
	            	expected: "also: [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated>
	            	in chain: "long: [0 0 0`)
}

func TestErrorNotErrorIsWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	err := fmt.Errorf("long: %v", longSlice)
	NotErrorIs(mock, err, err)
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Target error should not be in err chain:
	            	found: "long: [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated>
	            	in chain: "long: [0 0 0`)
}

func TestErrorAsWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	var target *customError
	ErrorAs(mock, fmt.Errorf("long: %v", longSlice), &target)
	Contains(t, mock.errorString(), fmt.Sprintf(`
	Error Trace:	
	Error:      	Should be in error chain:
	            	expected: *%s.customError`,
		shortpkg))
	Contains(t, mock.errorString(), `
	            	in chain: "long: [0 0 0`)
	Contains(t, mock.errorString(), "<... truncated>")
}

func TestErrorNotErrorAsWithErrorTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	var target *customError
	NotErrorAs(mock, fmt.Errorf("long: %v %w", longSlice, &customError{}), &target)
	Contains(t, mock.errorString(), fmt.Sprintf(`
	Error Trace:	
	Error:      	Target error should not be in err chain:
	            	found: *%s.customError`,
		shortpkg))
	Contains(t, mock.errorString(), `
	            	in chain: "long: [0 0 0`)
	Contains(t, mock.errorString(), "<... truncated>")
}

func TestErrorNotErrorAs(t *testing.T) {
	t.Parallel()

	for tt := range errorNotErrorAsCases() {
		t.Run(fmt.Sprintf("NotErrorAs(%#v,%#v)", tt.err, &customError{}), func(t *testing.T) {
			t.Parallel()
			mock := new(captureT)
			var target *customError

			res := NotErrorAs(mock, tt.err, &target)
			mock.checkResultAndErrMsg(t, tt.result, res, tt.resultErrMsg)
		})
	}
}

func TestErrorIs(t *testing.T) {
	t.Parallel()

	for tt := range errorIsCases() {
		t.Run(fmt.Sprintf("ErrorIs(%#v,%#v)", tt.err, tt.target), func(t *testing.T) {
			t.Parallel()
			mock := new(captureT)

			res := ErrorIs(mock, tt.err, tt.target)
			mock.checkResultAndErrMsg(t, tt.result, res, tt.resultErrMsg)
		})
	}
}

func TestErrorNotErrorIs(t *testing.T) {
	t.Parallel()

	for tt := range errorNotErrorIsCases() {
		t.Run(fmt.Sprintf("NotErrorIs(%#v,%#v)", tt.err, tt.target), func(t *testing.T) {
			t.Parallel()
			mock := new(captureT)

			res := NotErrorIs(mock, tt.err, tt.target)
			mock.checkResultAndErrMsg(t, tt.result, res, tt.resultErrMsg)
		})
	}
}

func TestErrorAs(t *testing.T) {
	t.Parallel()

	for tt := range errorAsCases() {
		t.Run(fmt.Sprintf("ErrorAs(%#v,%#v)", tt.err, &customError{}), func(t *testing.T) {
			t.Parallel()
			mock := new(captureT)
			var target *customError

			res := ErrorAs(mock, tt.err, &target)
			mock.checkResultAndErrMsg(t, tt.result, res, tt.resultErrMsg)
		})
	}
}

func TestErrorNoError(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// start with a nil error
	var err error

	True(t, NoError(mock, err), "NoError should return True for nil arg")

	// now set an error
	err = errors.New("some error")

	False(t, NoError(mock, err), "NoError with error should return False")

	// returning an empty error interface
	err = func() error {
		var err *customError
		return err
	}()

	if err == nil { // err is not nil here!
		t.Errorf("Error should be nil due to empty interface: %s", err)
	}

	False(t, NoError(mock, err), "NoError should fail with empty error interface")
}

func TestError(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// start with a nil error
	var err error

	False(t, Error(mock, err), "Error should return False for nil arg")

	// now set an error
	err = errors.New("some error")

	True(t, Error(mock, err), "Error with error should return True")

	// returning an empty error interface
	err = func() error {
		var err *customError
		return err
	}()

	if err == nil { // err is not nil here!
		t.Errorf("Error should be nil due to empty interface: %s", err)
	}

	True(t, Error(mock, err), "Error should pass with empty error interface")
}

func TestErrorEqualError(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// start with a nil error
	var err error
	False(t, EqualError(mock, err, ""),
		"EqualError should return false for nil arg")

	// now set an error
	err = errors.New("some error")
	False(t, EqualError(mock, err, "Not some error"),
		"EqualError should return false for different error string")
	True(t, EqualError(mock, err, "some error"),
		"EqualError should return true")
}

func TestErrorContains(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// start with a nil error
	var err error
	False(t, ErrorContains(mock, err, ""),
		"ErrorContains should return false for nil arg")

	// now set an error
	err = errors.New("some error: another error")
	False(t, ErrorContains(mock, err, "bad error"),
		"ErrorContains should return false for different error string")
	True(t, ErrorContains(mock, err, "some error"),
		"ErrorContains should return true")
	True(t, ErrorContains(mock, err, "another error"),
		"ErrorContains should return true")
}

type errorNotErrorAsCase struct {
	err          error
	result       bool
	resultErrMsg string
}

func errorNotErrorAsCases() iter.Seq[errorNotErrorAsCase] {
	return slices.Values([]errorNotErrorAsCase{
		{
			err:    fmt.Errorf("wrap: %w", &customError{}),
			result: false,
			resultErrMsg: "" +
				"Target error should not be in err chain:\n" +
				fmt.Sprintf("found: *%[1]s.customError\n", shortpkg) +
				"in chain: \"wrap: fail\" (*fmt.wrapError)\n" +
				fmt.Sprintf("\t\"fail\" (*%[1]s.customError)\n", shortpkg),
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
	err          error
	target       error
	result       bool
	resultErrMsg string
}

func errorIsCases() iter.Seq[errorIsCase] {
	return slices.Values([]errorIsCase{
		{
			err:    io.EOF,
			target: io.EOF,
			result: true,
		},
		{
			err:    fmt.Errorf("wrap: %w", io.EOF),
			target: io.EOF,
			result: true,
		},
		{
			err:    io.EOF,
			target: io.ErrClosedPipe,
			result: false,
			resultErrMsg: "" +
				"Target error should be in err chain:\n" +
				"expected: \"io: read/write on closed pipe\"\n" +
				"in chain: \"EOF\"\n",
		},
		{
			err:          nil,
			target:       io.EOF,
			result:       false,
			resultErrMsg: "Expected error with \"EOF\" in chain but got nil.\n",
		},
		{
			err:    io.EOF,
			target: nil,
			result: false,
			resultErrMsg: "" +
				"Target error should be in err chain:\n" +
				"expected: \"\"\n" +
				"in chain: \"EOF\"\n",
		},
		{
			err:    nil,
			target: nil,
			result: true,
		},
		{
			err:    fmt.Errorf("abc: %w", errors.New("def")),
			target: io.EOF,
			result: false,
			resultErrMsg: "" +
				"Target error should be in err chain:\n" +
				"expected: \"EOF\"\n" +
				"in chain: \"abc: def\"\n" +
				"\t\"def\"\n",
		},
	})
}

type errorNotErrorIsCase struct {
	err          error
	target       error
	result       bool
	resultErrMsg string
}

func errorNotErrorIsCases() iter.Seq[errorNotErrorIsCase] {
	return slices.Values([]errorNotErrorIsCase{
		{
			err:    io.EOF,
			target: io.EOF,
			result: false,
			resultErrMsg: "" +
				"Target error should not be in err chain:\n" +
				"found: \"EOF\"\n" +
				"in chain: \"EOF\"\n",
		},
		{
			err:    fmt.Errorf("wrap: %w", io.EOF),
			target: io.EOF,
			result: false,
			resultErrMsg: "" +
				"Target error should not be in err chain:\n" +
				"found: \"EOF\"\n" +
				"in chain: \"wrap: EOF\"\n" +
				"\t\"EOF\"\n",
		},
		{
			err:    io.EOF,
			target: io.ErrClosedPipe,
			result: true,
		},
		{
			err:    nil,
			target: io.EOF,
			result: true,
		},
		{
			err:    io.EOF,
			target: nil,
			result: true,
		},
		{
			err:    nil,
			target: nil,
			result: false,
			resultErrMsg: "" +
				"Target error should not be in err chain:\n" +
				"found: \"\"\n" +
				"in chain: \n",
		},
		{
			err:    fmt.Errorf("abc: %w", errors.New("def")),
			target: io.EOF,
			result: true,
		},
	})
}

type errorAsCase struct {
	err          error
	result       bool
	resultErrMsg string
}

func errorAsCases() iter.Seq[errorAsCase] {
	return slices.Values([]errorAsCase{
		{
			err:    fmt.Errorf("wrap: %w", &customError{}),
			result: true,
		},
		{
			err:    io.EOF,
			result: false,
			resultErrMsg: "" +
				"Should be in error chain:\n" +
				fmt.Sprintf("expected: *%[1]s.customError\n", shortpkg) +
				"in chain: \"EOF\" (*errors.errorString)\n",
		},
		{
			err:    nil,
			result: false,
			resultErrMsg: "" +
				"An error is expected but got nil.\n" +
				fmt.Sprintf(`expected: *%s.customError`, shortpkg) + "\n",
		},
		{
			err:    fmt.Errorf("abc: %w", errors.New("def")),
			result: false,
			resultErrMsg: "" +
				"Should be in error chain:\n" +
				fmt.Sprintf("expected: *%[1]s.customError\n", shortpkg) +
				"in chain: \"abc: def\" (*fmt.wrapError)\n" +
				"\t\"def\" (*errors.errorString)\n",
		},
	})
}

type customError struct{}

func (*customError) Error() string { return "fail" }
