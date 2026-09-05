// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build/constraint"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

// buildFileConstraints maps every syntax file of a package to its //go:build expression
// (e.g. "go1.26"), or "" when the file carries no build constraint.
//
// Build constraints are file-level in Go, so this is the unit at which we detect a guard
// and replicate it across generated files.
func buildFileConstraints(pkg *packages.Package) map[*token.File]string {
	constraints := make(map[*token.File]string, len(pkg.Syntax))

	for _, astFile := range pkg.Syntax {
		tokenFile := pkg.Fset.File(astFile.Pos())
		constraints[tokenFile] = fileBuildConstraint(astFile)
	}

	return constraints
}

// fileBuildConstraint returns the raw //go:build expression of a file (e.g. "go1.26"),
// or "" when the file has none.
//
// Only comment groups appearing before the package clause are considered, per the Go
// build-constraint placement rule.
func fileBuildConstraint(f *ast.File) string {
	for _, group := range f.Comments {
		if group.Pos() >= f.Package {
			break // build constraints must precede the package clause
		}

		for _, comment := range group.List {
			if !constraint.IsGoBuild(comment.Text) {
				continue
			}

			expr, err := constraint.Parse(comment.Text)
			if err != nil {
				return "" // malformed constraint: treat as unguarded (the compiler will complain)
			}

			return expr.String()
		}
	}

	return ""
}

// ErrGuardedFileDropped signals that a go-version-guarded source file never reached the scan.
var ErrGuardedFileDropped = errors.New("guarded source file missing from the scan")

// goVersionGuard returns the minor version N of a plain "go1.N" build constraint.
//
// Only a constraint that is exactly one go-version term qualifies. A file selected on
// something else — an OS, an architecture, a negation — is legitimately absent on some
// machines, and must not be mistaken for a file the toolchain dropped.
func goVersionGuard(expr string) (int, bool) {
	rest, ok := strings.CutPrefix(expr, "go1.")
	if !ok {
		return 0, false
	}

	minor, err := strconv.Atoi(rest)
	if err != nil {
		return 0, false
	}

	return minor, true
}

// guardedFilesOnDisk maps every non-test Go file in dir carrying a plain "//go:build go1.N"
// line to that constraint.
//
// The scan reads the directory with go/parser rather than the typed package on purpose: a
// file guarded above the toolchain running the scan never reaches [packages.Package], so the
// typed view cannot report what it is missing.
func guardedFilesOnDisk(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("can't read %q to look for build constraints: %w", dir, err)
	}

	guarded := make(map[string]string)
	fset := token.NewFileSet()

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue // the scan loads production files only (packages.Config.Tests is false)
		}

		astFile, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, parser.PackageClauseOnly|parser.ParseComments)
		if err != nil {
			continue // unparseable: the compiler reports it, we don't
		}

		expr := fileBuildConstraint(astFile)
		if _, ok := goVersionGuard(expr); ok {
			guarded[name] = expr
		}
	}

	return guarded, nil
}

// verifyGuardedFilesLoaded returns an error when a "//go:build go1.N" file sits in the scanned
// package directory but never made it into the typed load.
//
// That happens when the go command running the scan is older than the guard: it drops the file
// before the scanner sees it, and generating from that view leaves out every assertion the file
// declares, with nothing in the output to show for it. The toolchain that built this binary
// says nothing about the matter — [packages.Load] shells out to the go command on PATH.
func verifyGuardedFilesLoaded(pkg *packages.Package) error {
	if pkg.Dir == "" {
		return nil
	}

	guarded, err := guardedFilesOnDisk(pkg.Dir)
	if err != nil {
		return err
	}

	loaded := make(map[string]struct{}, len(pkg.GoFiles))
	for _, path := range pkg.GoFiles {
		loaded[filepath.Base(path)] = struct{}{}
	}

	dropped := make([]string, 0, len(guarded))
	highest := 0

	for name, expr := range guarded {
		if _, ok := loaded[name]; ok {
			continue
		}

		dropped = append(dropped, fmt.Sprintf("%s (//go:build %s)", name, expr))
		if minor, ok := goVersionGuard(expr); ok && minor > highest {
			highest = minor
		}
	}

	if len(dropped) == 0 {
		return nil
	}

	slices.Sort(dropped)

	return fmt.Errorf(
		"%w: %s. The go command running the scan is older than these guards, so the assertions they declare "+
			"would be left out of the generated packages and documentation. Rerun with GOTOOLCHAIN=go1.%d.0 "+
			"(see docs/doc-site/project/maintainers/CODEGEN.md, \"Regenerating\")",
		ErrGuardedFileDropped, strings.Join(dropped, ", "), highest,
	)
}
