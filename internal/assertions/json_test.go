// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
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
	//nolint:thelper // linter false positive: this is not a helper
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

		mock := new(testing.T)
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

		mock := new(testing.T)
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
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
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

//nolint:thelper // false positive: this is not a helper
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
	return slices.Values([]failCase{
		{
			name:         "JSONEq/not-equal",
			assertion:    func(t T) bool { return JSONEq(t, `{"a":1}`, `{"a":2}`) },
			wantContains: []string{"Not equal"},
		},
		{
			name:         "JSONEq/invalid-expected",
			assertion:    func(t T) bool { return JSONEq(t, "not json", `{"a":1}`) },
			wantContains: []string{"is not valid json"},
		},
		{
			name:         "JSONEq/invalid-actual",
			assertion:    func(t T) bool { return JSONEq(t, `{"a":1}`, "not json") },
			wantContains: []string{"needs to be valid json"},
		},
	})
}
