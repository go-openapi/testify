// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"testing"

	"pgregory.net/rapid"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// FuzzAssertionsNilSafety is the fuzzed equivalent of TestNilSafetyUnary
// and TestNilSafetyBinary.
//
// Given a high number of values of different types generated randomly,
// the fuzz engine will alter these values and run assertion functions.
//
// # Property
//
// No assertion function should ever panic, regardless of the inputs.
func FuzzAssertionsNilSafety(f *testing.F) {
	prop := func(rt *rapid.T) {
		value := genAny().Draw(rt, "value")
		other := genAny().Draw(rt, "other")
		mock := silentT{}

		noPanic(rt, "Nil", func() { _ = assertions.Nil(mock, value) })
		noPanic(rt, "NotNil", func() { _ = assertions.NotNil(mock, value) })
		noPanic(rt, "Empty", func() { _ = assertions.Empty(mock, value) })
		noPanic(rt, "NotEmpty", func() { _ = assertions.NotEmpty(mock, value) })
		noPanic(rt, "Zero", func() { _ = assertions.Zero(mock, value) })
		noPanic(rt, "NotZero", func() { _ = assertions.NotZero(mock, value) })
		noPanic(rt, "Len", func() { _ = assertions.Len(mock, value, 0) })
		noPanic(rt, "Equal", func() { _ = assertions.Equal(mock, value, other) })
		noPanic(rt, "NotEqual", func() { _ = assertions.NotEqual(mock, value, other) })
		noPanic(rt, "EqualValues", func() { _ = assertions.EqualValues(mock, value, other) })
		noPanic(rt, "Contains", func() { _ = assertions.Contains(mock, value, other) })
		noPanic(rt, "NotContains", func() { _ = assertions.NotContains(mock, value, other) })
		noPanic(rt, "IsType", func() { _ = assertions.IsType(mock, value, other) })
		noPanic(rt, "IsNotType", func() { _ = assertions.IsNotType(mock, value, other) })
	}

	f.Fuzz(rapid.MakeFuzz(prop))
}

// noPanic wraps a function call and reports a test error if it panics.
func noPanic(t *rapid.T, name string, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%s panicked: %v", name, r)
		}
	}()
	fn()
}
