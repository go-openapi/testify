// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"iter"
	"slices"
	"testing"
)

const shortpkg = "assertions"

func TestEqualUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("formatUnequalValue", testFormatUnequalValues())
	t.Run("validateEqualArgs", testValidateEqualArgs())
}

func testFormatUnequalValues() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for tt := range formatUnequalCases() {
			t.Run(tt.testName, func(t *testing.T) {
				t.Parallel()

				expected, actual := formatUnequalValues(tt.unequalExpected, tt.unequalActual)
				if tt.expectedExpected != expected {
					t.Errorf("%s: expected formatted expected %q, got %q", tt.testName, tt.expectedExpected, expected)
				}
				if tt.expectedActual != actual {
					t.Errorf("%s: expected formatted actual %q, got %q", tt.testName, tt.expectedActual, actual)
				}
			})
		}
	}
}

func testValidateEqualArgs() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if validateEqualArgs(func() {}, func() {}) == nil {
			t.Error("non-nil functions should error")
		}

		if validateEqualArgs(func() {}, func() {}) == nil {
			t.Error("non-nil functions should error")
		}

		if validateEqualArgs(nil, nil) != nil {
			t.Error("nil functions are equal")
		}
	}
}

type formatUnequalCase struct {
	unequalExpected  any
	unequalActual    any
	expectedExpected string
	expectedActual   string
	testName         string
}

func formatUnequalCases() iter.Seq[formatUnequalCase] {
	type testStructType struct {
		Val string
	}

	return slices.Values([]formatUnequalCase{
		{"foo", "bar", `"foo"`, `"bar"`, "value should not include type"},
		{123, 123, `123`, `123`, "value should not include type"},
		{int64(123), int32(123), `int64(123)`, `int32(123)`, "value should include type"},
		{int64(123), nil, `int64(123)`, `<nil>(<nil>)`, "value should include type"},
		{
			unequalExpected:  &testStructType{Val: "test"},
			unequalActual:    &testStructType{Val: "test"},
			expectedExpected: fmt.Sprintf(`&%s.testStructType{Val:"test"}`, shortpkg),
			expectedActual:   fmt.Sprintf(`&%s.testStructType{Val:"test"}`, shortpkg),
			testName:         "value should not include type annotation",
		},
		{uint(123), uint(124), `123`, `124`, "uint should print clean"},
		{uint8(123), uint8(124), `123`, `124`, "uint8 should print clean"},
		{uint16(123), uint16(124), `123`, `124`, "uint16 should print clean"},
		{uint32(123), uint32(124), `123`, `124`, "uint32 should print clean"},
		{uint64(123), uint64(124), `123`, `124`, "uint64 should print clean"},
	})
}
