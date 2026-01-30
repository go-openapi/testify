// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build testcolorized

package colors

import (
	"fmt"
	"testing"

	_ "github.com/go-openapi/testify/enable/colors/v2"
	target "github.com/go-openapi/testify/v2/assert"
)

// This test executes correctly with build tag "testcolorized".
//
// This tag enables an "escape hatch" in testify/enable/colors/v2 to run the init with colors enabled
// without the normally required command line arguments or environment variables.
//
// This setup allows us to run a full integration test on colorized output targeted only for this test only.

func TestColorsAssertJSONEq(t *testing.T) {
	t.Parallel()

	mockT := new(mockT)
	res := target.JSONEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "worldwide", "foo": "bar"}`)

	target.False(t, res)

	output := mockT.errorString()
	t.Log(output) // best to visualize the output
	target.Contains(t, output, "\x1b")
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
