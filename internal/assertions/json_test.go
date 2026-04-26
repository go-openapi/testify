// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"strings"
	"sync/atomic"
	"testing"
)

func TestJSONEq(t *testing.T) {
	t.Parallel()

	for tc := range jsonCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestJSONMarshalUnmarshalAs(t *testing.T) {
	t.Parallel()

	for tc := range jsonMarshalCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestJSONErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, jsonFailCases())
}

// =======================================
// Test JSONEq variants
// =======================================

// test all JSONEq variants with the same input (possibly converted).
func testAllJSONEq(expected, actual string, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("with JSONEq", testJSONEq(expected, actual, success))
		t.Run("with JSONEqBytes", testJSONEqBytes(expected, actual, success))
		t.Run("with JSONEqT[[]byte,string]", testJSONEqT[[]byte, string](expected, actual, success))
		t.Run("with JSONEqT[string,[]byte]", testJSONEqT[string, []byte](expected, actual, success))
		t.Run("with JSONEqT[byte,byte]", testJSONEqT[[]byte, []byte](expected, actual, success))
		t.Run("with JSONEqT[string,string]", testJSONEqT[string, string](expected, actual, success))
	}
}

func testJSONEq(expected, actual string, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		res := JSONEq(mock, expected, actual)
		if res != success {
			if success {
				croakWantEquiv(t, expected, actual)
				return
			}
			croakWantNotEquiv(t, expected, actual)
		}
	}
}

func testJSONEqBytes(expected, actual string, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		res := JSONEqBytes(mock, []byte(expected), []byte(actual))
		if res != success {
			if success {
				croakWantEquiv(t, expected, actual)
				return
			}
			croakWantNotEquiv(t, expected, actual)
		}
	}
}

func testJSONEqT[EDoc, ADoc Text](expected, actual string, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		res := JSONEqT(mock, EDoc(expected), ADoc(actual))
		if res != success {
			if success {
				croakWantEquiv(t, expected, actual)
				return
			}
			croakWantNotEquiv(t, expected, actual)
		}
	}
}

func croakWantEquiv(t *testing.T, expected, actual string) {
	t.Helper()
	t.Errorf("expected %q to be equivalent to %q", expected, actual)
}

func croakWantNotEquiv(t *testing.T, expected, actual string) {
	t.Helper()
	t.Errorf("expected %q NOT to be equivalent to %q", expected, actual)
}

func jsonCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"should be equivalent JSON", testAllJSONEq(
			`{"hello": "world", "foo": "bar"}`,
			`{"hello": "world", "foo": "bar"}`,
			true,
		)},
		{"should be equivalent but not equal JSON", testAllJSONEq(
			`{"hello": "world", "foo": "bar"}`,
			`{"foo": "bar", "hello": "world"}`,
			true,
		)},
		{"should be equivalent object with array and nested objects", testAllJSONEq(
			"{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],"+
				"\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}",
			"{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", "+
				"\"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}",
			true,
		)},
		{"should be equivalent array (ordered)", testAllJSONEq(
			`["foo", {"hello": "world", "nested": "hash"}]`,
			`["foo", {"nested": "hash", "hello": "world"}]`,
			true,
		)},
		{"should NOT be equivalent array (elements are different)", testAllJSONEq(
			`["foo", {"hello": "world", "nested": "hash"}]`,
			`{"foo": "bar", {"nested": "hash", "hello": "world"}}`,
			false,
		)},
		{"should NOT be equivalent array (order is different)", testAllJSONEq(
			`["foo", {"hello": "world", "nested": "hash"}]`,
			`[{ "hello": "world", "nested": "hash"}, "foo"]`,
			false,
		)},
		{"should NOT be equivalent objects (keys are different)", testAllJSONEq(
			`{"foo": "bar"}`,
			`{"foo": "bar", "hello": "world"}`,
			false,
		)},
		{"should fail with actual invalid JSON", testAllJSONEq(
			`{"foo": "bar"}`,
			"Not JSON",
			false,
		)},
		{"should fail with expected invalid JSON", testAllJSONEq(
			"Not JSON",
			`{"foo": "bar", "hello": "world"}`,
			false,
		)},
		{"should fail with expected and actual invalid JSON", testAllJSONEq(
			"Not JSON",
			"Not JSON",
			false,
		)},
	})
}

// =======================================
// Test JSONMarshalAsT / JSONUnmarshalAsT
// =======================================

func jsonMarshalCases() iter.Seq[genericTestCase] {
	type canMarshalJSON struct {
		A string `json:"a"`
		B int    `json:"b"`
		C int    `json:"c,omitempty"`
	}
	type cantMarshalJSON struct {
		A string `json:"a"`
		B int    `json:"b"`
		C func() `json:"c,omitempty"` // this fails when marshaling
	}

	const jazon = `{"a":"x","b":1}`
	valueCanDo := canMarshalJSON{A: "x", B: 1}
	valueCantDo := cantMarshalJSON{A: "x", B: 1, C: func() {}}

	return slices.Values([]genericTestCase{
		{"can JSON", testAllMarshalAs(valueCanDo, jazon, true)},
		{"can't JSON/marshal-fails", testAllMarshalAs(valueCantDo, jazon, false)},
		{"can JSON/different-values", testAllMarshalAs(valueCanDo, `{"a": 1,"b":"x"}`, false)},
	})
}

func testAllMarshalAs[Doc Text, Object any](value Object, jazon Doc, success bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("with JSONMarshalAsT", testJSONMarshalAsT(value, jazon, success))

		t.Run("with JSONUnmarshalAsT", testJSONUnmarshalAsT(value, jazon, success))
	}
}

func testJSONMarshalAsT[Doc Text, Object any](value Object, jazon Doc, success bool) func(*testing.T) {
	if success {
		return func(t *testing.T) {
			t.Run("should marshal JSON", func(t *testing.T) {
				t.Parallel()

				mock := new(mockT)
				res := JSONMarshalAsT(mock, jazon, value)
				if !res {
					t.Errorf("expected struct to marshal correctly as JSON: %#v <=> %s.Got %s", value, jazon, mock.errorString())
				}
			})
		}
	}

	return func(t *testing.T) {
		t.Run("should not marshal JSON", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := JSONMarshalAsT(mock, jazon, value)
			if res {
				t.Errorf("expected struct NOT to marshal correctly as JSON: %#v != %s", value, jazon)
			}
		})
	}
}

func testJSONUnmarshalAsT[Doc Text, Object any](value Object, jazon Doc, success bool) func(*testing.T) {
	if success {
		return func(t *testing.T) {
			t.Run("should unmarshal JSON", func(t *testing.T) {
				t.Parallel()

				mock := new(mockT)
				res := JSONUnmarshalAsT(mock, value, jazon)
				if !res {
					t.Errorf("expected json string to unmarshal correctly from JSON: %#v <=> %s. Got %s", value, jazon, mock.errorString())
				}
			})
		}
	}

	return func(t *testing.T) {
		t.Run("should not unmarshal JSON", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			res := JSONUnmarshalAsT(mock, value, jazon)
			if res {
				t.Errorf("expected json string NOT to unmarshal correctly from JSON: %#v != %s", value, jazon)
			}
		})
	}
}

// =======================================
// Test JSONErrorMessages
// =======================================

func jsonFailCases() iter.Seq[failCase] {
	const (
		part1 = `{"a":1}`
		part2 = `{"a":2}`
	)
	return slices.Values([]failCase{
		{
			name:         "JSONEq/not-equal",
			assertion:    func(t T) bool { return JSONEq(t, part1, part2) },
			wantContains: []string{"Not equal"},
		},
		{
			name:         "JSONEq/invalid-expected",
			assertion:    func(t T) bool { return JSONEq(t, "not json", part1) },
			wantContains: []string{"is not valid json"},
		},
		{
			name:         "JSONEq/invalid-actual",
			assertion:    func(t T) bool { return JSONEq(t, part1, "not json") },
			wantContains: []string{"needs to be valid json"},
		},
		{
			name: "JSONEqT/redactor-mismatch",
			assertion: func(t T) bool {
				return JSONEqT(t, func() string { return part1 }, part2)
			},
			wantContains: []string{"Not equal"},
		},
	})
}

// =======================================
// Test JSONEqT / JSONUnmarshalAsT / JSONMarshalAsT with Redactor inputs
//
// These tests exercise the Redactor (`func() string` / `func() []byte`)
// arms of the [RText] type set, the named-string/named-bytes default
// branch in asBytes, and the nil-redactor panic guard.
// =======================================

func TestJSONRedact(t *testing.T) {
	t.Parallel()

	t.Run("JSONEqT/redactor-string-on-expected", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		red := func() string { return `{"a":1}` }
		if !JSONEqT(mock, red, `{"a":1}`) {
			t.Errorf("expected redactor output to match literal; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONEqT/redactor-bytes-on-actual", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		red := func() []byte { return []byte(`{"a":1}`) }
		if !JSONEqT(mock, []byte(`{"a":1}`), red) {
			t.Errorf("expected literal to match redactor output; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONEqT/redactor-on-both-sides", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		left := func() string { return `{"a":1,"b":2}` }
		right := func() []byte { return []byte(`{"b":2,"a":1}`) }
		if !JSONEqT(mock, left, right) {
			t.Errorf("expected both redactor outputs to be JSON-equivalent; mock: %s", mock.errorString())
		}
	})

	// Real-world scenario: a redactor normalises a non-deterministic field
	// (timestamp) before comparison.
	t.Run("JSONEqT/redactor-normalises-timestamp", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		raw := `{"id":42,"timestamp":"2026-04-26T15:30:45Z"}`
		red := func() string {
			return strings.Replace(raw, `"2026-04-26T15:30:45Z"`, `"REDACTED"`, 1)
		}
		expected := `{"id":42,"timestamp":"REDACTED"}`
		if !JSONEqT(mock, expected, red) {
			t.Errorf("expected redactor to normalise timestamp; mock: %s", mock.errorString())
		}
	})

	// Side-effect proof: the redactor IS invoked (asBytes runs before the
	// comparison) and is invoked exactly once per call, regardless of which
	// position it's in.
	t.Run("JSONEqT/redactor-invoked-once", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		var calls atomic.Int32
		red := func() string {
			calls.Add(1)
			return `{"a":1}`
		}
		if !JSONEqT(mock, red, `{"a":1}`) {
			t.Errorf("expected match; mock: %s", mock.errorString())
		}
		if got := calls.Load(); got != 1 {
			t.Errorf("expected redactor to be called exactly once, got %d", got)
		}
	})

	t.Run("JSONUnmarshalAsT/with-redactor-string", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		type doc struct {
			A int `json:"a"`
		}
		red := func() string { return `{"a":42}` }
		if !JSONUnmarshalAsT(mock, doc{A: 42}, red) {
			t.Errorf("expected unmarshal of redactor output to match expected struct; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONUnmarshalAsT/with-redactor-bytes", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		type doc struct {
			A int `json:"a"`
		}
		red := func() []byte { return []byte(`{"a":42}`) }
		if !JSONUnmarshalAsT(mock, doc{A: 42}, red) {
			t.Errorf("expected unmarshal of redactor output to match expected struct; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONMarshalAsT/with-redactor-string", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		type doc struct {
			A int `json:"a"`
		}
		red := func() string { return `{"a":42}` }
		if !JSONMarshalAsT(mock, red, doc{A: 42}) {
			t.Errorf("expected marshal of struct to match redactor output; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONMarshalAsT/with-redactor-bytes", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		type doc struct {
			A int `json:"a"`
		}
		red := func() []byte { return []byte(`{"a":42}`) }
		if !JSONMarshalAsT(mock, red, doc{A: 42}) {
			t.Errorf("expected marshal of struct to match redactor output; mock: %s", mock.errorString())
		}
	})

	// Named string/[]byte types reach the default reflect branch in asBytes.
	t.Run("JSONEqT/named-string-type", func(t *testing.T) {
		t.Parallel()
		type myJSON string
		mock := new(mockT)
		var v myJSON = `{"a":1}`
		if !JSONEqT(mock, v, `{"a":1}`) {
			t.Errorf("expected named string type to convert to []byte via reflect; mock: %s", mock.errorString())
		}
	})

	t.Run("JSONEqT/named-bytes-type", func(t *testing.T) {
		t.Parallel()
		type myRaw []byte
		mock := new(mockT)
		v := myRaw(`{"a":1}`)
		if !JSONEqT(mock, v, []byte(`{"a":1}`)) {
			t.Errorf("expected named []byte type to convert via reflect; mock: %s", mock.errorString())
		}
	})

	// Nil redactors are a programming error and panic with a clear message
	// rather than the cryptic nil-pointer dereference users would otherwise see.
	t.Run("JSONEqT/nil-func-string-panics", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "JSONEqT", func() {
			var nilFn func() string
			_ = JSONEqT(new(mockT), nilFn, `{"a":1}`)
		})
	})

	t.Run("JSONEqT/nil-func-bytes-panics", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "JSONEqT", func() {
			var nilFn func() []byte
			_ = JSONEqT(new(mockT), nilFn, []byte(`{"a":1}`))
		})
	})

	t.Run("JSONUnmarshalAsT/nil-func-string-panics", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "JSONUnmarshalAsT", func() {
			var nilFn func() string
			_ = JSONUnmarshalAsT(new(mockT), struct{}{}, nilFn)
		})
	})

	t.Run("JSONMarshalAsT/nil-func-bytes-panics", func(t *testing.T) {
		t.Parallel()
		mustPanicWithRedactor(t, "JSONMarshalAsT", func() {
			var nilFn func() []byte
			_ = JSONMarshalAsT(new(mockT), nilFn, struct{}{})
		})
	})
}

// mustPanicWithRedactor asserts that fn panics with the standard
// nil-Redactor message. Used by both JSON and YAML redactor tests.
func mustPanicWithRedactor(t *testing.T, fn string, body func()) {
	t.Helper()
	const wantMsg = "passed Redactor cannot be nil"
	defer func() {
		rec := recover()
		if rec == nil {
			t.Errorf("expected %s to panic with nil Redactor, got no panic", fn)
			return
		}
		s, _ := rec.(string)
		if !strings.Contains(s, wantMsg) {
			t.Errorf("expected %s panic message to contain %q, got: %v", fn, wantMsg, rec)
		}
	}()
	body()
}
