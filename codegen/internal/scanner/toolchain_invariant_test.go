// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

// repoRootFromScanner is the path from this test package to the repository root.
const repoRootFromScanner = "../../.."

var goMinorRx = regexp.MustCompile(`go1\.(\d+)`)

// TestToolchainFloorCoversGuards enforces the invariant that every //go:build go1.N guard
// used in internal/assertions is covered by the go.work toolchain floor.
//
// codegen runs in workspace mode, where the go.work toolchain line selects the toolchain.
// A guard above that floor could be silently dropped (go/packages would not even load the
// file), producing incomplete generated output. Bumping the floor must therefore happen in
// lockstep with introducing a higher guard.
func TestToolchainFloorCoversGuards(t *testing.T) {
	floor := workToolchainMinor(t, filepath.Join(repoRootFromScanner, "go.work"))
	maxGuard := maxAssertionGuardMinor(t, filepath.Join(repoRootFromScanner, "internal", "assertions"))

	t.Logf("go.work toolchain floor: go1.%d, highest internal/assertions guard: go1.%d", floor, maxGuard)

	if maxGuard > floor {
		t.Fatalf(
			"internal/assertions uses //go:build go1.%d but the go.work toolchain floor is go1.%d; "+
				"bump the go.work toolchain line to at least go1.%d so codegen observes the guarded file",
			maxGuard, floor, maxGuard,
		)
	}
}

// workToolchainMinor returns the minor version of the go.work toolchain floor, falling back
// to the workspace `go` directive when no explicit toolchain line is present.
func workToolchainMinor(t *testing.T, path string) int {
	t.Helper()

	data, err := os.ReadFile(path) //nolint:gosec // test reads a fixed in-repo file
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	// Prefer the toolchain directive (e.g. "toolchain go1.26.0"); otherwise the go directive
	// (e.g. "go 1.25.0") acts as the effective floor.
	if m := regexp.MustCompile(`(?m)^toolchain go1\.(\d+)`).FindSubmatch(data); m != nil {
		return mustAtoi(t, string(m[1]))
	}
	if m := regexp.MustCompile(`(?m)^go 1\.(\d+)`).FindSubmatch(data); m != nil {
		return mustAtoi(t, string(m[1]))
	}

	t.Fatalf("could not find a toolchain or go directive in %s", path)

	return 0
}

// maxAssertionGuardMinor textually scans every Go file in dir for //go:build go1.N guards
// and returns the highest minor version found (0 when none).
//
// The scan is textual (via go/parser, not go/packages) on purpose: a guard above the
// running toolchain would be excluded from a typed load, which is exactly the situation
// this invariant must detect.
func maxAssertionGuardMinor(t *testing.T, dir string) int {
	t.Helper()

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read dir %s: %v", dir, err)
	}

	fset := token.NewFileSet()
	maxMinor := 0
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || filepath.Ext(name) != ".go" {
			continue
		}

		file, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, parser.ParseComments|parser.SkipObjectResolution)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}

		constraintExpr := fileBuildConstraint(file)
		if constraintExpr == "" {
			continue
		}

		for _, m := range goMinorRx.FindAllStringSubmatch(constraintExpr, -1) {
			if minor := mustAtoi(t, m[1]); minor > maxMinor {
				maxMinor = minor
			}
		}
	}

	return maxMinor
}

func mustAtoi(t *testing.T, s string) int {
	t.Helper()

	n, err := strconv.Atoi(s)
	if err != nil {
		t.Fatalf("invalid integer %q: %v", s, err)
	}

	return n
}
