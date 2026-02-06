// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

func TestYAMLEq_EqualYAMLString(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
}

func TestYAMLEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock, expectedYAML, actualYAML))
}

func TestYAMLEq_Array(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`))
}

func TestYAMLEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`))
}

func TestYAMLEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `{"foo": "bar"}`, "Simple String"))
}

func TestYAMLEq_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, "Simple String", `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock, "Simple String", "Simple String"))
}

func TestYAMLEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`))
}

func TestYAMLEq_OnlyFirstDocument(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.True(t, target.YAMLEq(mock,
		`---
doc1: same
---
doc2: different
`,
		`---
doc1: same
---
doc2: notsame
`,
	))
}

func TestYAMLEq_InvalidIdenticalYAML(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `}`, `}`))
}

func TestYAMLUnmarshalAs(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	type dummy struct {
		Hello string `yaml:"hello"`
		Foo   string `yaml:"foo"`
	}

	value := dummy{Hello: "world", Foo: "bar"}
	target.False(t, target.YAMLUnmarshalAsT(mock, value, `{"hello": "buzz", "foo": "lightyear"}`))
	target.False(t, target.YAMLMarshalAsT(mock, `{"hello": "buzz", "foo": "lightyear"}`, value))

	target.True(t, target.YAMLUnmarshalAsT(mock, value, `{"hello": "world", "foo": "bar"}`))
	target.True(t, target.YAMLMarshalAsT(mock, `{"hello": "world", "foo": "bar"}`, value))
}
