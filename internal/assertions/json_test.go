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
