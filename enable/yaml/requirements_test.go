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

	mockT := new(MockT)
	target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, expectedYAML, actualYAML)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_Array(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `{"foo": "bar"}`, "Simple String")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, "Simple String", `{"foo": "bar", "hello": "world"}`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestRequireYAMLEq_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, "Simple String", "Simple String")
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestRequireYAMLEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

type MockT struct {
	Failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (MockT) Helper() {}

func (t *MockT) FailNow() {
	t.Failed = true
}

func (t *MockT) Errorf(format string, args ...any) {
	_, _ = format, args
}
