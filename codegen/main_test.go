// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"iter"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// TestExecute verifies the execute function works with valid configuration.
func TestExecute(t *testing.T) {
	// Skip if we're in a short test run
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	targetRoot := filepath.Join(tmpDir, "output")

	cfg := &config{
		dir:        "..",
		inputPkg:   "github.com/go-openapi/testify/v2/internal/assertions",
		outputPkgs: "assert",
		targetRoot: targetRoot,
		includeFmt: true,
		includeFwd: true,
		includeTst: true,
		includeGen: false,
		includeHlp: true,
		includeExa: true,
		runExa:     true,
	}

	err := execute(cfg)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verify that the assert package was generated
	assertPkg := filepath.Join(targetRoot, "assert")
	if _, err := os.Stat(assertPkg); os.IsNotExist(err) {
		t.Error("Assert package directory was not created")
	}

	// Verify key generated files exist
	expectedFiles := []string{
		"assert_types.go",
		"assert_assertions.go",
		"assert_format.go",
		"assert_forward.go",
		"assert_helpers.go",
	}

	for _, file := range expectedFiles {
		path := filepath.Join(assertPkg, file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not generated", file)
		}
	}
}

// TestExecuteMultiplePackages verifies that generating multiple packages works.
func TestExecuteMultiplePackages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	targetRoot := filepath.Join(tmpDir, "output")

	cfg := &config{
		dir:        "..",
		inputPkg:   "github.com/go-openapi/testify/v2/internal/assertions",
		outputPkgs: "assert,require",
		targetRoot: targetRoot,
		includeFmt: false,
		includeFwd: false,
		includeTst: false,
		includeGen: false,
		includeHlp: false,
		includeExa: false,
		runExa:     false,
	}

	err := execute(cfg)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verify both packages were generated
	for _, pkg := range []string{"assert", "require"} {
		pkgDir := filepath.Join(targetRoot, pkg)
		if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
			t.Errorf("Package %s was not created", pkg)
		}

		// Verify at least the main assertions file exists
		assertionsFile := filepath.Join(pkgDir, pkg+"_assertions.go")
		if _, err := os.Stat(assertionsFile); os.IsNotExist(err) {
			t.Errorf("Assertions file for %s was not generated", pkg)
		}
	}
}

// TestExecuteInvalidConfig verifies error handling for invalid configuration.
func TestExecuteInvalidConfig(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	for tt := range invalidConfigCases(dir) {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := execute(tt.cfg)
			if err == nil {
				t.Error("Expected error but got nil")
			}
		})
	}
}

type invalidConfigCase struct {
	name string
	cfg  *config
}

func invalidConfigCases(dir string) iter.Seq[invalidConfigCase] {
	return slices.Values([]invalidConfigCase{
		{
			name: "invalid input package",
			cfg: &config{
				dir:        "..",
				inputPkg:   "invalid/package/path",
				outputPkgs: "assert",
				targetRoot: dir,
			},
		},
		{
			name: "nonexistent work dir",
			cfg: &config{
				dir:        "/nonexistent/directory",
				inputPkg:   "github.com/go-openapi/testify/v2/internal/assertions",
				outputPkgs: "assert",
				targetRoot: dir,
			},
		},
	})
}
