// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
)

func TestYAMLEq_EqualYAMLString(t *testing.T) {
	t.Parallel()

	enableYAML()
	mockT := new(testing.T)
	target.True(t, target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
	target.False(t, target.YAMLEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "buzz", "foo": "lightyear"}`))
}
