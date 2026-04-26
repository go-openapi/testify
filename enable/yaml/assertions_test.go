// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"strings"
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

// TestYAMLRedact covers the happy-path Redactor (`func() string` /
// `func() []byte`) arms of the [RText] type set for the YAML T-variants,
// which can only be exercised once the YAML feature is enabled. The
// nil-Redactor panic path is also re-checked here for YAMLMarshalAsT,
// since in the internal/assertions package yaml.Marshal panics first
// when the feature is disabled and the asBytes guard is unreachable.
func TestYAMLRedact(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	type dummy struct {
		Hello string `yaml:"hello"`
		Foo   string `yaml:"foo"`
	}
	value := dummy{Hello: "world", Foo: "bar"}
	const doc = `{"hello": "world", "foo": "bar"}`

	t.Run("YAMLEqT with func() string Redactor", func(t *testing.T) {
		t.Parallel()
		red := func() string { return doc }
		target.True(t, target.YAMLEqT(mock, red, doc))
	})

	t.Run("YAMLEqT with func() []byte Redactor on actual", func(t *testing.T) {
		t.Parallel()
		red := func() []byte { return []byte(doc) }
		target.True(t, target.YAMLEqT(mock, doc, red))
	})

	// Real-world scenario: a redactor normalises a non-deterministic field
	// before comparison.
	t.Run("YAMLEqT redactor normalises field", func(t *testing.T) {
		t.Parallel()
		raw := `{"id": 42, "ts": "2026-04-26T15:30:45Z"}`
		red := func() string {
			return strings.Replace(raw, `"2026-04-26T15:30:45Z"`, `"REDACTED"`, 1)
		}
		expected := `{"id": 42, "ts": "REDACTED"}`
		target.True(t, target.YAMLEqT(mock, expected, red))
	})

	t.Run("YAMLUnmarshalAsT with Redactor", func(t *testing.T) {
		t.Parallel()
		red := func() []byte { return []byte(doc) }
		target.True(t, target.YAMLUnmarshalAsT(mock, value, red))
	})

	t.Run("YAMLMarshalAsT with Redactor", func(t *testing.T) {
		t.Parallel()
		red := func() string { return doc }
		target.True(t, target.YAMLMarshalAsT(mock, red, value))
	})

	t.Run("YAMLMarshalAsT with nil Redactor panics with clear message", func(t *testing.T) {
		t.Parallel()
		const wantMsg = "passed Redactor cannot be nil"
		defer func() {
			rec := recover()
			if rec == nil {
				t.Errorf("expected YAMLMarshalAsT to panic with nil Redactor, got no panic")
				return
			}
			s, _ := rec.(string)
			if !strings.Contains(s, wantMsg) {
				t.Errorf("expected panic message to contain %q, got: %v", wantMsg, rec)
			}
		}()
		var nilFn func() string
		_ = target.YAMLMarshalAsT(mock, nilFn, value)
	})
}
