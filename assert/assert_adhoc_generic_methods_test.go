// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build go1.27

package assert_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

// TestGenericMethods exercises the generic assertions as methods against a real
// [testing.T], the way a user writes them. The generated tests in this package drive
// the same methods through a mock T; this one checks that they hold together outside
// codegen, with the type parameter inferred, spelled out, or bound in a method value.
func TestGenericMethods(t *testing.T) {
	t.Parallel()

	a := assert.New(t)

	// inferred from the arguments
	a.EqualT("expected", "expected")
	a.NotEqualT(1, 2)
	a.SliceContainsT([]string{"a", "b"}, "b")
	a.MapContainsT(map[string]int{"a": 1}, "a")
	a.GreaterT(2, 1)
	a.InDeltaT(1.0, 1.000001, 1e-3)

	// spelled out, for a type parameter no argument carries
	a.IsOfTypeT[int](42)
	a.IsNotOfTypeT[string](42)

	// formatted variant
	a.EqualTf(3, 3, "expected %d", 3)

	// a method value, instantiated once and reused
	equalInt := a.EqualT[int]
	equalInt(1, 1)
	equalInt(2, 2)
}
