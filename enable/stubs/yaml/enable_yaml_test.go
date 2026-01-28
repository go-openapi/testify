// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"fmt"
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

// TestEnableYAML is merely a smoke test to validate that the chain of calls is resolved properly.
func TestEnableYAML(t *testing.T) {
	t.Parallel()

	EnableYAMLWithUnmarshal(func(_ []byte, _ any) error {
		return fmt.Errorf("called: %w", target.ErrTest)
	})
	EnableYAMLWithMarshal(func(_ any) ([]byte, error) {
		return nil, fmt.Errorf("called: %w", target.ErrTest)
	})
	type dummy struct {
		Hello string `yaml:"hello"`
		Foo   string `yaml:"hello"`
	}
	value := dummy{Hello: "world", Foo: "bar"}

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))

	// the struct is correct, but we inject a serializer that always fails
	target.False(t, target.YAMLUnmarshalAsT(mock, value, `{"hello": "world", "foo": "bar"}`))
	target.False(t, target.YAMLMarshalAsT(mock, `{"hello": "world", "foo": "bar"}`, value))
}
