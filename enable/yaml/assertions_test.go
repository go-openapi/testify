package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

func TestYAMLEq_EqualYAMLString(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
}

func TestYAMLEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
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
	target.True(t, target.YAMLEq(mockT, expected, actual))
}

func TestYAMLEq_Array(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`))
}

func TestYAMLEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`))
}

func TestYAMLEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, `{"foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_ActualIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, `{"foo": "bar"}`, "Simple String"))
}

func TestYAMLEq_ExpectedIsSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, "Simple String", `{"foo": "bar", "hello": "world"}`))
}

func TestYAMLEq_ExpectedAndActualSimpleString(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT, "Simple String", "Simple String"))
}

func TestYAMLEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`))
}

func TestYAMLEq_OnlyFirstDocument(t *testing.T) {
	t.Parallel()

	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT,
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

	mockT := new(testing.T)
	target.False(t, target.YAMLEq(mockT, `}`, `}`))
}
