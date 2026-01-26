// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"fmt"
	"os"
	"strings"
	"testing"

	target "github.com/go-openapi/testify/v2/assert"
	colorstub "github.com/go-openapi/testify/v2/enable/stubs/colors"
)

func TestMain(m *testing.M) {
	// we can't easily simulate arg flags in CI (uses gotestsum etc).
	// Similarly, env vars are evaluated too early.
	colorstub.Enable(
		func() []colorstub.Option {
			return []colorstub.Option{
				colorstub.WithEnable(true),
				colorstub.WithSanitizedTheme(flags.theme),
			}
		})

	os.Exit(m.Run())
}

func TestAssertJSONEq(t *testing.T) {
	t.Parallel()

	mockT := new(mockT)
	res := target.JSONEq(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "worldwide", "foo": "bar"}`)

	target.False(t, res)

	output := mockT.errorString()
	t.Log(output) // best to visualize the output
	target.Contains(t, neuterize(output), neuterize(expectedColorizedDiff))
}

func TestAssertJSONEq_Array(t *testing.T) {
	t.Parallel()

	mockT := new(mockT)
	res := target.JSONEq(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `["bar", {"nested": "hash", "hello": "world"}]`)

	target.False(t, res)
	output := mockT.errorString()
	t.Log(output) // best to visualize the output
	target.Contains(t, neuterize(output), neuterize(expectedColorizedArrayDiff))
}

func neuterize(str string) string {
	// remove blanks and replace escape sequences for readability
	blankRemover := strings.NewReplacer("\t", "", " ", "", "\x1b", "^[")
	return blankRemover.Replace(str)
}

type mockT struct {
	errorFmt string
	args     []any
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockT) Helper() {}

func (m *mockT) Errorf(format string, args ...any) {
	m.errorFmt = format
	m.args = args
}

func (m *mockT) Failed() bool {
	return m.errorFmt != ""
}

func (m *mockT) errorString() string {
	return fmt.Sprintf(m.errorFmt, m.args...)
}

// captured output (indentation is not checked)
//
//nolint:staticcheck // indeed we want to check the escape sequences in this test
const (
	expectedColorizedDiff = ` Not equal:
	            	expected: [0;92mmap[string]interface {}{"foo":"bar", "hello":"world"}[0m
	            	actual  : [0;91mmap[string]interface {}{"foo":"bar", "hello":"worldwide"}[0m

	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -2,3 +2,3 @@
	            	[0;92m  (string) (len=3) "foo": (string) (len=3) "bar",
	            	[0m[0;91m- (string) (len=5) "hello": (string) (len=5) "world"
	            	[0m[0;93m+ (string) (len=5) "hello": (string) (len=9) "worldwide"
	            	[0m[0;92m }
	            	[0m
`

	expectedColorizedArrayDiff = `Not equal:
        	            	expected: [0;92m[]interface {}{"foo", map[string]interface {}{"hello":"world", "nested":"hash"}}[0m
        	            	actual  : [0;91m[]interface {}{"bar", map[string]interface {}{"hello":"world", "nested":"hash"}}[0m
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -1,3 +1,3 @@
        	            	[0;92m ([]interface {}) (len=2) {
        	            	[0m[0;91m- (string) (len=3) "foo",
        	            	[0m[0;93m+ (string) (len=3) "bar",
        	            	[0m[0;92m  (map[string]interface {}) (len=2) {
        	            	[0m
`
)
