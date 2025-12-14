// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"fmt"
	"iter"
	"runtime"
	"slices"
	"strings"
	"testing"
	"time"
)

const pkg = "github.com/go-openapi/testify/v2/internal/assertions"

func TestCompareGreater(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Greater(mock, 2, 1) {
		t.Error("Greater should return true")
	}

	if Greater(mock, 1, 1) {
		t.Error("Greater should return false")
	}

	if Greater(mock, 1, 2) {
		t.Error("Greater should return false")
	}

	// check error report
	for currCase := range compareIncreasingFixtures() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Greater(out, currCase.less, currCase.greater))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Greater")
	}
}

func TestCompareGreaterOrEqual(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !GreaterOrEqual(mock, 2, 1) {
		t.Error("GreaterOrEqual should return true")
	}

	if !GreaterOrEqual(mock, 1, 1) {
		t.Error("GreaterOrEqual should return true")
	}

	if GreaterOrEqual(mock, 1, 2) {
		t.Error("GreaterOrEqual should return false")
	}

	// check error report
	for currCase := range compareIncreasingFixtures() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, GreaterOrEqual(out, currCase.less, currCase.greater))
		Contains(t, out.buf.String(), strings.ReplaceAll(currCase.msg, "than", "than or equal to"))
		Contains(t, out.helpers, pkg+".GreaterOrEqual")
	}
}

func TestCompareLess(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Less(mock, 1, 2) {
		t.Error("Less should return true")
	}

	if Less(mock, 1, 1) {
		t.Error("Less should return false")
	}

	if Less(mock, 2, 1) {
		t.Error("Less should return false")
	}

	// check error report
	for currCase := range compareIncreasingFixtures3() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Less(out, currCase.greater, currCase.less))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Less")
	}
}

func TestCompareLessOrEqual(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !LessOrEqual(mock, 1, 2) {
		t.Error("LessOrEqual should return true")
	}

	if !LessOrEqual(mock, 1, 1) {
		t.Error("LessOrEqual should return true")
	}

	if LessOrEqual(mock, 2, 1) {
		t.Error("LessOrEqual should return false")
	}

	// check error report
	for currCase := range compareIncreasingFixtures3() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, LessOrEqual(out, currCase.greater, currCase.less))
		Contains(t, out.buf.String(), strings.ReplaceAll(currCase.msg, "than", "than or equal to"))
		Contains(t, out.helpers, pkg+".LessOrEqual")
	}
}

func TestComparePositive(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Positive(mock, 1) {
		t.Error("Positive should return true")
	}

	if !Positive(mock, 1.23) {
		t.Error("Positive should return true")
	}

	if Positive(mock, -1) {
		t.Error("Positive should return false")
	}

	if Positive(mock, -1.23) {
		t.Error("Positive should return false")
	}

	// Check error report
	for currCase := range comparePositiveCases() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Positive(out, currCase.e))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Positive")
	}
}

func TestCompareNegative(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Negative(mock, -1) {
		t.Error("Negative should return true")
	}

	if !Negative(mock, -1.23) {
		t.Error("Negative should return true")
	}

	if Negative(mock, 1) {
		t.Error("Negative should return false")
	}

	if Negative(mock, 1.23) {
		t.Error("Negative should return false")
	}

	// Check error report
	for currCase := range compareNegativeCases() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Negative(out, currCase.e))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Negative")
	}
}

func TestCompareMsgAndArgsForwarding(t *testing.T) {
	msgAndArgs := []any{"format %s %x", "this", 0xc001}
	expectedOutput := "format this c001\n"

	funcs := []func(t T){
		func(t T) { Greater(t, 1, 2, msgAndArgs...) },
		func(t T) { GreaterOrEqual(t, 1, 2, msgAndArgs...) },
		func(t T) { Less(t, 2, 1, msgAndArgs...) },
		func(t T) { LessOrEqual(t, 2, 1, msgAndArgs...) },
		func(t T) { Positive(t, 0, msgAndArgs...) },
		func(t T) { Negative(t, 0, msgAndArgs...) },
	}

	for _, f := range funcs {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		f(out)
		Contains(t, out.buf.String(), expectedOutput)
	}
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

type compareFixture struct {
	less    any
	greater any
	msg     string
}

//nolint:dupl // factoring further the message to save a little duplication would make the test harder to read
func compareIncreasingFixtures() iter.Seq[compareFixture] {
	return slices.Values(
		[]compareFixture{
			{less: "a", greater: "b", msg: `"a" is not greater than "b"`},
			{less: int(1), greater: int(2), msg: `"1" is not greater than "2"`},
			{less: int8(1), greater: int8(2), msg: `"1" is not greater than "2"`},
			{less: int16(1), greater: int16(2), msg: `"1" is not greater than "2"`},
			{less: int32(1), greater: int32(2), msg: `"1" is not greater than "2"`},
			{less: int64(1), greater: int64(2), msg: `"1" is not greater than "2"`},
			{less: uint8(1), greater: uint8(2), msg: `"1" is not greater than "2"`},
			{less: uint16(1), greater: uint16(2), msg: `"1" is not greater than "2"`},
			{less: uint32(1), greater: uint32(2), msg: `"1" is not greater than "2"`},
			{less: uint64(1), greater: uint64(2), msg: `"1" is not greater than "2"`},
			{less: float32(1.23), greater: float32(2.34), msg: `"1.23" is not greater than "2.34"`},
			{less: float64(1.23), greater: float64(2.34), msg: `"1.23" is not greater than "2.34"`},
			{less: uintptr(1), greater: uintptr(2), msg: `"1" is not greater than "2"`},
			{less: time.Time{}, greater: time.Time{}.Add(time.Hour), msg: `"0001-01-01 00:00:00 +0000 UTC" is not greater than "0001-01-01 01:00:00 +0000 UTC"`},
			{less: []byte{1, 1}, greater: []byte{1, 2}, msg: `"[1 1]" is not greater than "[1 2]"`},
		},
	)
}

//nolint:dupl // factoring further the message to save a little duplication would make the test harder to read
func compareIncreasingFixtures3() iter.Seq[compareFixture] {
	return slices.Values(
		[]compareFixture{
			{less: "a", greater: "b", msg: `"b" is not less than "a"`},
			{less: int(1), greater: int(2), msg: `"2" is not less than "1"`},
			{less: int8(1), greater: int8(2), msg: `"2" is not less than "1"`},
			{less: int16(1), greater: int16(2), msg: `"2" is not less than "1"`},
			{less: int32(1), greater: int32(2), msg: `"2" is not less than "1"`},
			{less: int64(1), greater: int64(2), msg: `"2" is not less than "1"`},
			{less: uint8(1), greater: uint8(2), msg: `"2" is not less than "1"`},
			{less: uint16(1), greater: uint16(2), msg: `"2" is not less than "1"`},
			{less: uint32(1), greater: uint32(2), msg: `"2" is not less than "1"`},
			{less: uint64(1), greater: uint64(2), msg: `"2" is not less than "1"`},
			{less: float32(1.23), greater: float32(2.34), msg: `"2.34" is not less than "1.23"`},
			{less: float64(1.23), greater: float64(2.34), msg: `"2.34" is not less than "1.23"`},
			{less: uintptr(1), greater: uintptr(2), msg: `"2" is not less than "1"`},
			{less: time.Time{}, greater: time.Time{}.Add(time.Hour), msg: `"0001-01-01 01:00:00 +0000 UTC" is not less than "0001-01-01 00:00:00 +0000 UTC"`},
			{less: []byte{1, 1}, greater: []byte{1, 2}, msg: `"[1 2]" is not less than "[1 1]"`},
		},
	)
}

type compareTestCase struct {
	e   any
	msg string
}

type comparePositiveCase = compareTestCase

type compareNegativeCase = compareTestCase

func comparePositiveCases() iter.Seq[comparePositiveCase] {
	return slices.Values([]comparePositiveCase{
		{e: int(-1), msg: `"-1" is not positive`},
		{e: int8(-1), msg: `"-1" is not positive`},
		{e: int16(-1), msg: `"-1" is not positive`},
		{e: int32(-1), msg: `"-1" is not positive`},
		{e: int64(-1), msg: `"-1" is not positive`},
		{e: float32(-1.23), msg: `"-1.23" is not positive`},
		{e: float64(-1.23), msg: `"-1.23" is not positive`},
	})
}

func compareNegativeCases() iter.Seq[compareNegativeCase] {
	return slices.Values([]compareNegativeCase{
		{e: int(1), msg: `"1" is not negative`},
		{e: int8(1), msg: `"1" is not negative`},
		{e: int16(1), msg: `"1" is not negative`},
		{e: int32(1), msg: `"1" is not negative`},
		{e: int64(1), msg: `"1" is not negative`},
		{e: float32(1.23), msg: `"1.23" is not negative`},
		{e: float64(1.23), msg: `"1.23" is not negative`},
	})
}
