// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"iter"
	"os"
	"os/exec"
	"path"
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

	t.Run("should prepare go.mod to build in tmpDir", func(t *testing.T) {
		if err := os.MkdirAll(targetRoot, 0o700); err != nil {
			t.Fatalf("could not create working dir %s: %v", targetRoot, err)
		}

		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("could not retrieve current dir: %v", err)
		}

		// NOTE: source formatting & imports check requires a proper go mod when building outside of the current GOROOT.
		goModInit(t, targetRoot, filepath.Join(cwd, ".."))
	})

	cfg := &config{
		dir:        "..",
		inputPkg:   "github.com/go-openapi/testify/v2/internal/assertions",
		outputPkgs: "assert",
		targetRoot: targetRoot,
		targetDoc:  "doc",
		includeFmt: true,
		includeFwd: true,
		includeTst: true,
		includeGen: true,
		includeHlp: true,
		includeExa: true,
		includeDoc: true,
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

	t.Run("generated source files should exist", func(t *testing.T) {
		expectedFiles := []string{
			"assert_types.go",
			"assert_assertions.go",
			"assert_format.go",
			"assert_forward.go",
			"assert_helpers.go",
		}

		for _, file := range expectedFiles {
			pth := filepath.Join(assertPkg, file)
			if _, err := os.Stat(pth); os.IsNotExist(err) {
				t.Errorf("Expected file %s was not generated", file)
			}
		}
	})

	t.Run("generated doc files should exist", func(t *testing.T) {
		docLocation := filepath.Join(tmpDir, "output", "doc")
		if _, err := os.Stat(docLocation); os.IsNotExist(err) {
			t.Error("doc directory was not created")
		}

		pth := filepath.Join(docLocation, "number.md") // only check one sample
		if _, err := os.Stat(pth); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not generated", pth)
		}
	})
}

// TestExecuteMultiplePackages verifies that generating multiple packages works.
func TestExecuteMultiplePackages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tmpDir := t.TempDir()
	targetRoot := filepath.Join(tmpDir, "output")
	// NOTE: we don't need all the go.mod preparation work in this simpler test case

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

// TestSmokeRegisterFlags merely passes over the flag registration step to ensure
// we don't have a panic due to flags registered several times.
//
// There is no real assertion to be made.
func TestSmokeRegisterFlags(_ *testing.T) {
	registerFlags(&config{})
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

//nolint:gosec // "tainted" args exec is actually okay in tests
func goModInit(t *testing.T, location, source string) {
	t.Run("should init go.mod", func(t *testing.T) {
		mod := exec.CommandContext(t.Context(), "go", "mod", "init", path.Base(location))
		mod.Dir = location
		output, err := mod.CombinedOutput()
		if err != nil {
			t.Fatalf("go mod init at %s returned: %v: %s", location, err, string(output))
		}

		t.Run("should replace testify/v2 in go.mod by local references", func(t *testing.T) {
			replace := exec.CommandContext(t.Context(), "go", "mod", "edit",
				"-replace=github.com/go-openapi/testify/v2="+source,
			)
			replace.Dir = location
			output, err := replace.CombinedOutput()
			if err != nil {
				t.Fatalf("go mod edit at %s returned: %v: %s", location, err, string(output))
			}

			t.Run("should load required modules", func(t *testing.T) {
				get := exec.CommandContext(t.Context(), "go", "get",
					"github.com/go-openapi/testify/v2/assert",
					"github.com/go-openapi/testify/v2/require",
				)
				get.Dir = location
				output, err := get.CombinedOutput()
				if err != nil {
					t.Fatalf("go get at %s returned: %v: %s", location, err, string(output))
				}
			})
		})
	})
}
