// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"sync/atomic"
	"testing"
)

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
		panicked := func() (didPanic bool) {
			defer func() {
				if recover() != nil {
					didPanic = true
				}
			}()
			_ = YAMLEq(mock, expected, actual)
			return false
		}()
		if !panicked {
			croakWantPanic(t, "YAMLEq")
		}
	}
}

func testYAMLEqBytes(expected, actual string, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		panicked := func() (didPanic bool) {
			defer func() {
				if recover() != nil {
					didPanic = true
				}
			}()
			_ = YAMLEqBytes(mock, []byte(expected), []byte(actual))
			return false
		}()
		if !panicked {
			croakWantPanic(t, "YAMLEqBytes")
		}
	}
}

func testYAMLEqT[ADoc, EDoc Text](expected, actual string, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		panicked := func() (didPanic bool) {
			defer func() {
				if recover() != nil {
					didPanic = true
				}
			}()
			_ = YAMLEqT(mock, EDoc(expected), ADoc(actual))
			return false
		}()
		if !panicked {
			croakWantPanic(t, "YAMLEqT")
		}
	}
}

func testYAMLUnmarshalAsT[ADoc Text, Object any](expected Object, actual ADoc, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		panicked := func() (didPanic bool) {
			defer func() {
				if recover() != nil {
					didPanic = true
				}
			}()
			_ = YAMLUnmarshalAsT(mock, expected, actual)
			return false
		}()
		if !panicked {
			croakWantPanic(t, "YAMLUnmarshalAsT")
		}
	}
}

func testYAMLMarshalAsT[EDoc Text](expected EDoc, actual any, success bool) func(*testing.T) {
	_ = success
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		panicked := func() (didPanic bool) {
			defer func() {
				if recover() != nil {
					didPanic = true
				}
			}()
			_ = YAMLMarshalAsT(mock, expected, actual)
			return false
		}()
		if !panicked {
			croakWantPanic(t, "YAMLMarshalAsT")
		}
	}
}

func croakWantPanic(t *testing.T, fn string) {
	t.Helper()
	t.Errorf("expected %q to panic with default settings", fn)
}

// =======================================
// Test YAML T-variants accept Redactor inputs
//
// In this package the YAML feature is disabled, so a successful redactor
// call still ends in a panic from YAMLEqBytes. These tests verify that:
//
//   1. The Redactor (`func() string` / `func() []byte`) arms of the [RText]
//      constraint compile and route through asBytes correctly.
//   2. The redactor body IS invoked (proving asBytes runs before the
//      YAML-disabled panic).
//   3. Nil redactors panic with the standard "passed Redactor cannot be nil"
//      message — same guard as the JSON path.
// =======================================

func TestYAMLRedact(t *testing.T) {
	t.Parallel()

	t.Run("YAMLEqT/redactor-invoked-then-yaml-panics", func(t *testing.T) {
		t.Parallel()
		var calls atomic.Int32
		red := func() string {
			calls.Add(1)
			return "key: value"
		}
		didPanic := func() (panicked bool) {
			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()
			_ = YAMLEqT(new(mockT), red, "key: value")
			return false
		}()
		if !didPanic {
			croakWantPanic(t, "YAMLEqT (with Redactor)")
		}
		// asBytes must have run at least once (one side is the redactor).
		if got := calls.Load(); got != 1 {
			t.Errorf("expected redactor to be invoked exactly once before YAML-disabled panic, got %d", got)
		}
	})

	t.Run("YAMLUnmarshalAsT/redactor-bytes-invoked-then-yaml-panics", func(t *testing.T) {
		t.Parallel()
		var calls atomic.Int32
		red := func() []byte {
			calls.Add(1)
			return []byte("key: value")
		}
		didPanic := func() (panicked bool) {
			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()
			_ = YAMLUnmarshalAsT(new(mockT), struct{}{}, red)
			return false
		}()
		if !didPanic {
			croakWantPanic(t, "YAMLUnmarshalAsT (with Redactor)")
		}
		if got := calls.Load(); got != 1 {
			t.Errorf("expected redactor to be invoked exactly once before YAML-disabled panic, got %d", got)
		}
	})

	// YAMLMarshalAsT runs `yaml.Marshal(object)` BEFORE evaluating
	// `asBytes(expected)`, so when the YAML feature is disabled the
	// redactor body is never reached. We can only verify that the call
	// still panics; happy-path Redactor coverage for YAMLMarshalAsT
	// belongs in the enable/yaml integration tests.
	t.Run("YAMLMarshalAsT/with-redactor-still-panics", func(t *testing.T) {
		t.Parallel()
		red := func() string { return "key: value" }
		didPanic := func() (panicked bool) {
			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()
			_ = YAMLMarshalAsT(new(mockT), red, struct{}{})
			return false
		}()
		if !didPanic {
			croakWantPanic(t, "YAMLMarshalAsT (with Redactor)")
		}
	})

	// Nil redactors panic with the standard message before reaching YAMLEqBytes,
	// so this asserts the asBytes guard fires regardless of the YAML feature.
	t.Run("YAMLEqT/nil-func-string-panics-with-redactor-message", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "YAMLEqT", func() {
			var nilFn func() string
			_ = YAMLEqT(new(mockT), nilFn, "key: value")
		})
	})

	t.Run("YAMLUnmarshalAsT/nil-func-bytes-panics-with-redactor-message", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "YAMLUnmarshalAsT", func() {
			var nilFn func() []byte
			_ = YAMLUnmarshalAsT(new(mockT), struct{}{}, nilFn)
		})
	})

	// YAMLMarshalAsT/nil-redactor case omitted: yaml.Marshal panics first
	// when YAML is disabled, so the nil-Redactor guard inside asBytes is
	// unreachable here. Covered in the enable/yaml integration tests.
}
