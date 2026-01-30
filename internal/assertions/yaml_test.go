// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestYAML(t *testing.T) {
	t.Parallel()

	t.Run("should panic", testAllYAMLEq())
}

// =======================================
// TestYAML: all YAML assertions
// =======================================

func testAllYAMLEq() func(*testing.T) {
	return func(t *testing.T) {
		const (
			actual = `
---
a: 1
`
			expected = ""
			success  = false
		)
		a := struct {
			A string `json:"a"`
		}{
			A: "x",
		}

		t.Run("with YAMLEq", testYAMLEq(expected, actual, success))
		t.Run("with YAMLEqBytes", testYAMLEqBytes(expected, actual, success))
		t.Run("with YAMLEqT[string,string]", testYAMLEqT[string, string](expected, actual, success))
		t.Run("with YAMLEqT[[]byte,string]", testYAMLEqT[[]byte, string](expected, actual, success))
		t.Run("with YAMLEqT[string,[]byte]", testYAMLEqT[string, []byte](expected, actual, success))
		t.Run("with YAMLEqT[[]byte,[]byte]", testYAMLEqT[[]byte, []byte](expected, actual, success))
		t.Run("with YAMLMarshalAsT[[]byte,struct{}]", testYAMLMarshalAsT(expected, a, success))
		t.Run("with YAMLUnmarshalAsT[struct{},[]byte]", testYAMLUnmarshalAsT(a, actual, success))
	}
}

func testYAMLEq(expected, actual string, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Panics(t, func() {
			_ = YAMLEq(mock, expected, actual)
		}) {
			croakWantPanic(t, "YAMLEq")
		}
	}
}

func testYAMLEqBytes(expected, actual string, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Panics(t, func() {
			_ = YAMLEqBytes(mock, []byte(expected), []byte(actual))
		}) {
			croakWantPanic(t, "YAMLEqBytes")
		}
	}
}

//nolint:thelper // linter false positive: this is not a helper
func testYAMLEqT[ADoc, EDoc Text](expected, actual string, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Panics(t, func() {
			_ = YAMLEqT(mock, EDoc(expected), ADoc(actual))
		}) {
			croakWantPanic(t, "YAMLEqT")
		}
	}
}

func testYAMLUnmarshalAsT[ADoc Text, Object any](expected Object, actual ADoc, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Panics(t, func() {
			_ = YAMLUnmarshalAsT(mock, expected, actual)
		}) {
			croakWantPanic(t, "YAMLUnmarshalAsT")
		}
	}
}

func testYAMLMarshalAsT[EDoc Text](expected EDoc, actual any, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Panics(t, func() {
			_ = YAMLMarshalAsT(mock, expected, actual)
		}) {
			croakWantPanic(t, "YAMLMarshalAsT")
		}
	}
}

func croakWantPanic(t *testing.T, fn string) {
	t.Helper()
	t.Errorf("expected %q to panic with default settings", fn)
}
