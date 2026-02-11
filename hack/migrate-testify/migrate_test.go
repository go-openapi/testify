// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"iter"
	"os"
	"slices"
	"strings"
	"testing"

	"golang.org/x/tools/go/ast/astutil"
)

type migrateTestCase struct {
	name     string
	input    string
	expected string
	// warnContains, if non-empty, checks that a warning was emitted containing this string.
	warnContains string
}

func migrateTestCases() iter.Seq[migrateTestCase] {
	return slices.Values([]migrateTestCase{
		{
			name:     "basic import rewrite",
			input:    readTestdata("migrate_basic/input.go.txt"),
			expected: readTestdata("migrate_basic/expected.go.txt"),
		},
		{
			name:     "yaml enable injection",
			input:    readTestdata("migrate_yaml/input.go.txt"),
			expected: readTestdata("migrate_yaml/expected.go.txt"),
		},
		{
			name:     "aliased import with rename",
			input:    readTestdata("migrate_alias/input.go.txt"),
			expected: readTestdata("migrate_alias/expected.go.txt"),
		},
		{
			name:         "incompatible imports warn",
			input:        readTestdata("migrate_incompatible/input.go.txt"),
			warnContains: "mock package is not available",
		},
		{
			name:     "PanicTestFunc replacement",
			input:    readTestdata("migrate_panic_func/input.go.txt"),
			expected: readTestdata("migrate_panic_func/expected.go.txt"),
		},
	})
}

func TestMigrateFile(t *testing.T) {
	t.Parallel()

	for c := range migrateTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			runMigrateSubtest(t, c)
		})
	}
}

func runMigrateSubtest(t *testing.T, c migrateTestCase) {
	t.Helper()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", c.input, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	rpt := &report{}

	// Build aliases before rewriting imports.
	aliases := buildAliasMap(f)

	// Detect incompatible imports.
	for imp, msg := range incompatibleImports {
		if fileImportsPath(f, imp) {
			rpt.warn("test.go", 0, msg)
		}
	}

	// Rewrite imports.
	for old, replacement := range importRewrites {
		astutil.RewriteImport(fset, f, old, replacement)
	}

	// Rename functions.
	renameFunctions(f, aliases, fset, rpt, "test.go", true)

	// Replace PanicTestFunc.
	replacePanicTestFunc(f, aliases, fset, rpt, "test.go", true)

	// YAML injection.
	if needsYAMLEnable(f, aliases) {
		astutil.AddNamedImport(fset, f, "_", goopenapiYAMLEnable)
	}

	// Check warnings.
	if c.warnContains != "" {
		found := false
		for _, d := range rpt.diagnostics {
			if d.kind == "warning" && strings.Contains(d.message, c.warnContains) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected warning containing %q, got: %v", c.warnContains, rpt.diagnostics)
		}
		return
	}

	// Compare output.
	var buf strings.Builder
	cfg := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	if err := cfg.Fprint(&buf, fset, f); err != nil {
		t.Fatalf("print: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	want := strings.TrimSpace(c.expected)

	if got != want {
		t.Errorf("output mismatch:\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func readTestdata(path string) string {
	data, err := os.ReadFile("testdata/" + path)
	if err != nil {
		panic("reading testdata: " + err.Error())
	}
	return string(bytes.ReplaceAll(data, []byte{'\r'}, []byte{})) // on windows, remove the \r from \n\r sequences
}
