// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

func TestYAMLEqWrapper_EqualYAMLString(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if !assert.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`) {
		t.Error("YAMLEq should return true")
	}
}

func TestYAMLEqWrapper_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if !assert.YAMLEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`) {
		t.Error("YAMLEq should return true")
	}
}

func TestYAMLEqWrapper_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
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
	if !assert.YAMLEq(expected, actual) {
		t.Error("YAMLEq should return true")
	}
}

func TestYAMLEqWrapper_Array(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if !assert.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`) {
		t.Error("YAMLEq should return true")
	}
}

func TestYAMLEqWrapper_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if assert.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`) {
		t.Error("YAMLEq should return false")
	}
}

func TestYAMLEqWrapper_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if assert.YAMLEq(`{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`) {
		t.Error("YAMLEq should return false")
	}
}

func TestYAMLEqWrapper_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if assert.YAMLEq(`{"foo": "bar"}`, "Simple String") {
		t.Error("YAMLEq should return false")
	}
}

func TestYAMLEqWrapper_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if assert.YAMLEq("Simple String", `{"foo": "bar", "hello": "world"}`) {
		t.Error("YAMLEq should return false")
	}
}

func TestYAMLEqWrapper_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if !assert.YAMLEq("Simple String", "Simple String") {
		t.Error("YAMLEq should return true")
	}
}

func TestYAMLEqWrapper_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	assert := target.New(new(testing.T))
	if assert.YAMLEq(`["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`) {
		t.Error("YAMLEq should return false")
	}
}
