// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/require"
)

func TestRequireYAMLEqWrapper_EqualYAMLString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

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
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_Array(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`{"foo": "bar"}`, "Simple String")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq("Simple String", `{"foo": "bar", "hello": "world"}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEqWrapper_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq("Simple String", "Simple String")
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEqWrapper_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	mockRequire := target.New(mockT)

	mockRequire.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}
