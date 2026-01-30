// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/go-openapi/testify/enable/colors/v2"
	target "github.com/go-openapi/testify/v2/assert"
)

//nolint:gochecknoinits // we want specifically to test here the initialization process
func init() {
	// This forces the test program built when testing this package to run with the testify flags for colorized output.
	//
	// We append the colorization flags (if not already there) to the command line _before_ it gets processed by go test
	// flags parsing. It seems like this is the only way to explore the package initialization code in testify/enable/v2.
	if len(os.Args) < 2 || !strings.Contains(os.Args[len(os.Args)-2], "-testify.colorized") {
		os.Args = append(os.Args, "-testify.colorized", "-testify.colorized.notty")
	}
}

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
