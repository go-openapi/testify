// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
	colorstub "github.com/go-openapi/testify/v2/internal/assertions/enable/colors"
)

func TestEnableColors(t *testing.T) {
	t.Parallel()

	colorstub.Enable(
		func() []colorstub.Option {
			return []colorstub.Option{
				colorstub.WithEnable(true),
				colorstub.WithSanitizedTheme("light"),
			}
		})

	mock := new(testing.T)
	target.NotPanics(mock, func() {
		_ = target.JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "worldwide", "foo": "bar"}`)
	})
}
