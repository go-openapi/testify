// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/require"
)

func TestRequireYAMLEqWrapper_EqualYAMLString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	expected := `
numeric: 1.5
array:
  - foo: bar
  - 1
  - "string"
  - ["nested", "array", 5.5]
hash:
  nested: hash
  nested_slice: [this, is, nested]
string: "foo"
`

	actual := `
numeric: 1.5
hash:
  nested: hash
  nested_slice: [this, is, nested]
string: "foo"
array:
  - foo: bar
  - 1
  - "string"
  - ["nested", "array", 5.5]
`

	mockRequire.YAMLEq(expected, actual)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_Array(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`{"foo": "bar"}`, "Simple String")
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq("Simple String", `{"foo": "bar", "hello": "world"}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq("Simple String", "Simple String")
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	mockRequire := target.New(mock)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}
