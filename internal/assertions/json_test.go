// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

// Proposal for enhancement: load fixtures and assertions from embedded testdata

func TestJSONEq_EqualSONString(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	True(t, JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
}

func TestJSONEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	True(t, JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestJSONEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, JSONEq(mock,
		"{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],"+
			"\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}",
		"{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", "+
			"\"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}",
	))
}

func TestJSONEq_Array(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, JSONEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`))
}

func TestJSONEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`))
}

func TestJSONEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, `{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestJSONEq_ActualIsNotJSON(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, `{"foo": "bar"}`, "Not JSON"))
}

func TestJSONEq_ExpectedIsNotJSON(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, "Not JSON", `{"foo": "bar", "hello": "world"}`))
}

func TestJSONEq_ExpectedAndActualNotJSON(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, "Not JSON", "Not JSON"))
}

func TestJSONEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, JSONEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`))
}

func TestJSONEqBytes_EqualSONString(t *testing.T) {
	mock := new(testing.T)
	True(t, JSONEqBytes(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"hello": "world", "foo": "bar"}`)))
}

func TestJSONEqBytes_EquivalentButNotEqual(t *testing.T) {
	mock := new(testing.T)
	True(t, JSONEqBytes(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`)))
}

func TestJSONEqBytes_HashOfArraysAndHashes(t *testing.T) {
	mock := new(testing.T)
	True(t, JSONEqBytes(mock,
		[]byte("{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],"+
			"\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}"),
		[]byte("{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", "+
			"\"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}",
		)))
}

func TestJSONEqBytes_Array(t *testing.T) {
	mock := new(testing.T)
	True(t, JSONEqBytes(mock, []byte(`["foo", {"hello": "world", "nested": "hash"}]`), []byte(`["foo", {"nested": "hash", "hello": "world"}]`)))
}

func TestJSONEqBytes_HashAndArrayNotEquivalent(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte(`["foo", {"hello": "world", "nested": "hash"}]`), []byte(`{"foo": "bar", {"nested": "hash", "hello": "world"}}`)))
}

func TestJSONEqBytes_HashesNotEquivalent(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte(`{"foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`)))
}

func TestJSONEqBytes_ActualIsNotJSON(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte(`{"foo": "bar"}`), []byte("Not JSON")))
}

func TestJSONEqBytes_ExpectedIsNotJSON(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte("Not JSON"), []byte(`{"foo": "bar", "hello": "world"}`)))
}

func TestJSONEqBytes_ExpectedAndActualNotJSON(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte("Not JSON"), []byte("Not JSON")))
}

func TestJSONEqBytes_ArraysOfDifferentOrder(t *testing.T) {
	mock := new(testing.T)
	False(t, JSONEqBytes(mock, []byte(`["foo", {"hello": "world", "nested": "hash"}]`), []byte(`[{ "hello": "world", "nested": "hash"}, "foo"]`)))
}
