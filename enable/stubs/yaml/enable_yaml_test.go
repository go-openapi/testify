// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"fmt"
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

func TestEnableYAML(t *testing.T) {
	t.Parallel()

	EnableYAMLWithUnmarshal(func(_ []byte, _ any) error {
		return fmt.Errorf("called: %w", target.ErrTest)
	})

	mock := new(testing.T)
	target.False(t, target.YAMLEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
}
