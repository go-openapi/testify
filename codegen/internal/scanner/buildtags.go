// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"go/ast"
	"go/build/constraint"
	"go/token"

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
