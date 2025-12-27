// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package comments

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

// BuildFilesMap constructs a lookup index to help bridge token position vs ast.File.
//
// This creates a map from token.File (file metadata) to ast.File (syntax tree),
// enabling O(1) lookup after O(n) construction.
func BuildFilesMap(pkg *packages.Package) map[*token.File]*ast.File {
	filesMap := make(map[*token.File]*ast.File, len(pkg.Syntax))

	for _, astFile := range pkg.Syntax {
		tokenFile := pkg.Fset.File(astFile.Pos()) // o(log n) lookup
		filesMap[tokenFile] = astFile
	}

	return filesMap
}

// resolveDeclFromObject finds the AST declaration for a types.Object using position-based lookup.
//
// Returns both the declaration and the file containing it, or (nil, nil) if not found.
func resolveDeclFromObject(object types.Object, fileSet *token.FileSet, filesMap map[*token.File]*ast.File) (ast.Decl, *ast.File) {
	pos := object.Pos()
	if !pos.IsValid() {
		return nil, nil
	}

	tokenFile := fileSet.File(pos) // o(log n) lookup
	if tokenFile == nil {
		return nil, nil
	}

	astFile := filesMap[tokenFile] // o(1) lookup
	if astFile == nil {
		return nil, nil
	}

	path, _ := astutil.PathEnclosingInterval(astFile, pos, pos)
	for _, node := range path {
		declaration, ok := node.(ast.Decl)
		if !ok {
			continue
		}

		return declaration, astFile
	}

	return nil, astFile
}

// findCommentInFile finds a comment group within a position range.
func findCommentInFile(astFile *ast.File, minPos, maxPos token.Pos) *ast.CommentGroup {
	for _, commentGroup := range astFile.Comments {
		pos := commentGroup.Pos()
		if pos > minPos && pos < maxPos {
			return commentGroup
		}
	}

	return nil
}
