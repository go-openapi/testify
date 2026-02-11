// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"
)

// loadPackages loads all Go packages under dir with full type information.
// If a go.work file is present, it loads packages from all workspace modules.
// Otherwise, it loads packages from the single module at dir.
func loadPackages(dir string) ([]*packages.Package, *token.FileSet, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, nil, fmt.Errorf("resolving path: %w", err)
	}

	patterns := workspacePatterns(absDir)

	fset := token.NewFileSet()
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedImports |
			packages.NeedDeps,
		Dir:   absDir,
		Fset:  fset,
		Tests: true,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, nil, fmt.Errorf("loading packages: %w", err)
	}

	// Report loading errors but don't fail — partial results are useful.
	var errs []string
	for _, pkg := range pkgs {
		for _, e := range pkg.Errors {
			errs = append(errs, e.Error())
		}
	}
	if len(errs) > 0 {
		// Log but continue — we can still transform files that loaded.
		fmt.Printf("warning: %d package loading errors (some files may be skipped):\n", len(errs)) //nolint:forbidigo // CLI output
		for _, e := range errs {
			fmt.Printf("  %s\n", e) //nolint:forbidigo // CLI output
		}
	}

	return pkgs, fset, nil
}

// workspacePatterns returns the load patterns for packages.Load.
// If go.work exists, it returns a pattern per workspace module (e.g. "./conv/...", "./mangling/...").
// If go.work does not exist, it returns ["./..."] and warns if sub-modules are detected.
func workspacePatterns(absDir string) []string {
	goworkPath := filepath.Join(absDir, "go.work")
	data, err := os.ReadFile(goworkPath)
	if err != nil {
		// No go.work — check for sub-modules and warn.
		warnIfSubModules(absDir)
		return []string{"./..."}
	}

	wf, err := modfile.ParseWork(goworkPath, data, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to parse go.work: %v; falling back to ./...\n", err)
		return []string{"./..."}
	}

	patterns := make([]string, 0, len(wf.Use))
	for _, use := range wf.Use {
		dir := use.Path
		if dir == "." {
			patterns = append(patterns, "./...")
		} else {
			// Normalize: strip leading "./" if present, ensure trailing "/...".
			dir = strings.TrimPrefix(dir, "./")
			patterns = append(patterns, "./"+dir+"/...")
		}
	}

	if len(patterns) == 0 {
		return []string{"./..."}
	}

	fmt.Printf("go.work: loading %d workspace modules\n", len(patterns)) //nolint:forbidigo // CLI output
	return patterns
}

// warnIfSubModules scans for nested go.mod files and warns if found.
func warnIfSubModules(absDir string) {
	var subModules []string

	_ = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil //nolint:nilerr // intentionally swallow walk errors for best-effort scan
		}
		if info.IsDir() {
			base := filepath.Base(path)
			if base == "vendor" || base == ".git" || base == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if info.Name() == "go.mod" && path != filepath.Join(absDir, "go.mod") {
			rel, _ := filepath.Rel(absDir, filepath.Dir(path))
			subModules = append(subModules, rel)
		}
		return nil
	})

	if len(subModules) == 0 {
		return
	}

	fmt.Fprintf(os.Stderr, "warning: found %d sub-modules without go.work; pass 2 will only cover the root module\n", len(subModules))
	fmt.Fprintf(os.Stderr, "  → Create go.work with: go work init . %s\n", strings.Join(prefixDot(subModules), " "))
}

// prefixDot adds "./" prefix to each path.
func prefixDot(paths []string) []string {
	result := make([]string, len(paths))
	for i, p := range paths {
		result[i] = "./" + p
	}
	return result
}

// fileImportsPath reports whether the given file has an import with the specified path.
func fileImportsPath(f *ast.File, path string) bool {
	for _, imp := range f.Imports {
		if importPath(imp) == path {
			return true
		}
	}
	return false
}

// importPath returns the unquoted import path from an ImportSpec.
func importPath(imp *ast.ImportSpec) string {
	return strings.Trim(imp.Path.Value, `"`)
}

// fileImportsAny reports whether the file imports any path with the given prefix.
func fileImportsAny(f *ast.File, prefix string) bool {
	for _, imp := range f.Imports {
		if strings.HasPrefix(importPath(imp), prefix) {
			return true
		}
	}
	return false
}

// isVendorPath reports whether the file path is inside a vendor/ directory.
func isVendorPath(path string) bool {
	return strings.Contains(path, "/vendor/") || strings.HasPrefix(path, "vendor/")
}
