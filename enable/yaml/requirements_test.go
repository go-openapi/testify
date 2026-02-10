// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/require"
)

const (
	expectedYAML = `
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

	actualYAML = `
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
)

func TestRequireYAMLEq_EqualYAMLString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, expectedYAML, actualYAML)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_Array(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`)
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `{"foo": "bar"}`, "Simple String")
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, "Simple String", `{"foo": "bar", "hello": "world"}`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, "Simple String", "Simple String")
	if mock.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	target.YAMLEq(mock, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`)
	if !mock.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLUnmarshalAsWrapper(t *testing.T) {
	t.Parallel()

	type dummy struct {
		Hello string `yaml:"hello"`
		Foo   string `yaml:"foo"`
	}

	t.Run("should pass", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		value := dummy{Hello: "world", Foo: "bar"}
		target.YAMLUnmarshalAsT(mock, value, `{"hello": "world", "foo": "bar"}`)
		if mock.Failed {
			t.Error("Check should pass")
		}

		target.YAMLMarshalAsT(mock, `{"hello": "world", "foo": "bar"}`, value)
		if mock.Failed {
			t.Error("Check should pass")
		}
	})

	t.Run("should fail", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		value := dummy{Hello: "world", Foo: "bar"}
		target.YAMLUnmarshalAsT(mock, value, `{"hello": "world", "foo": "yay"}`)
		if !mock.Failed {
			t.Error("Check should fail with FailNow")
		}

		target.YAMLMarshalAsT(mock, `{"hello": "world", "foo": "yay"}`, value)
		if !mock.Failed {
			t.Error("Check should fail with FailNow")
		}
	})
}

type mockT struct {
	Failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockT) Helper() {}

func (t *mockT) FailNow() {
	t.Failed = true
}

func (t *mockT) Errorf(format string, args ...any) {
	_, _ = format, args
}
