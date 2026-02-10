// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"fmt"
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
	colorstub "github.com/go-openapi/testify/v2/internal/assertions/enable/colors"
)

func TestEnableColors(t *testing.T) {
	// Test execution is serialized on the sensitive package initialization section
	//
	// Proper test coverage of Enable() is assured by the testintegration package.

	for _, opt := range []Option{
		WithSanitizedTheme("light"),
		WithDark(),
		WithLight(),
		WithTheme(ThemeDark),
	} {
		mock := new(mockT)
		target.NotPanics(mock, func() {
			testEnableWithLock(opt) // executed serially

			_ = target.JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `{"hello": "worldwide", "foo": "bar"}`)
		})
		// we may call several times with different options, but only the first setup is going to be used:
		// colorization is set lazily with a sync.Once global by the consuming package.
		t.Log(mock.errorString())
	}
}

func testEnableWithLock(opt Option) {
	colorstub.Enable(
		func() []colorstub.Option {
			return []colorstub.Option{
				colorstub.WithEnable(true),
				opt,
			}
		})
}

type mockT struct {
	errorFmt string
	args     []any
}

func (m *mockT) Errorf(format string, args ...any) {
	m.errorFmt = format
	m.args = args
}

func (m *mockT) errorString() string {
	return fmt.Sprintf(m.errorFmt, m.args...)
}
