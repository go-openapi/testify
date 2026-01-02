// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package comments

import (
	"go/ast"
	"go/token"
	"go/types"
	"regexp"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
	parser "github.com/go-openapi/testify/codegen/v2/internal/scanner/comments-parser"
)

var copyrightRex = regexp.MustCompile(`(?i)copyright`)

type cacheEntry struct {
	Decl ast.Decl
	File *ast.File
}

// Extractor extracts comments from Go source code using position-based AST/types bridging.
type Extractor struct {
	syntaxPackage []*ast.File
	fileSet       *token.FileSet
	filesMap      map[*token.File]*ast.File
	declCache     map[types.Object]cacheEntry
}

// New creates a comment extractor.
//
// The syntaxPackage provides access to all AST files in the package.
// The fileSet and filesMap enable position-based lookup to bridge types.Object to ast.Decl.
func New(syntaxPackage []*ast.File, fileSet *token.FileSet, filesMap map[*token.File]*ast.File) *Extractor {
	return &Extractor{
		syntaxPackage: syntaxPackage,
		fileSet:       fileSet,
		filesMap:      filesMap,
		declCache:     make(map[types.Object]cacheEntry),
	}
}

// ExtractPackageComments finds the package-level comment that describes the source package.
//
// This is a single comment group from the first file that has package documentation.
// Never returns nil.
func (e *Extractor) ExtractPackageComments() *ast.CommentGroup {
	for _, file := range e.syntaxPackage {
		if file.Doc == nil {
			continue
		}

		return file.Doc
	}

	// safeguard: never returns nil
	return &ast.CommentGroup{}
}

// ExtractCopyright scans source files and expects that at least one contains
// a leading comment line with the string "copyright".
//
// Never returns nil.
func (e *Extractor) ExtractCopyright() *ast.CommentGroup {
FILE:
	for _, file := range e.syntaxPackage {
		if len(file.Comments) == 0 {
			continue
		}

		// Check all comment groups before the package declaration
		// (needed for files with build tags, where copyright is not the first group)
		for _, group := range file.Comments {
			// Skip comment groups that come after the package declaration
			if group.Pos() >= file.Package {
				continue FILE
			}

			// Check if this group contains a copyright notice
			for _, line := range group.List {
				if copyrightRex.MatchString(line.Text) {
					return group
				}
			}
		}
	}

	// safeguard: never returns nil
	return &ast.CommentGroup{}
}

// ExtractComments retrieves the docstring for an object.
//
// Never returns nil: if no comment is present an empty but valid [ast.CommentGroup] is returned.
//
// NOTE: comment prioritization rule is:
//  1. GenDecl.Doc = leading comment for the entire declaration block
//  2. Spec.Doc = individual doc for that specific spec
func (e *Extractor) ExtractComments(object types.Object) *ast.CommentGroup {
	entry, ok := e.declCache[object]
	if ok {
		// short-circuit already resolved declarations
		if entry.Decl == nil {
			return &ast.CommentGroup{}
		}

		return extractCommentFromDecl(entry.Decl, object)
	}

	declaration, astFile := resolveDeclFromObject(object, e.fileSet, e.filesMap)
	if declaration == nil {
		e.declCache[object] = cacheEntry{}

		return &ast.CommentGroup{}
	}

	e.declCache[object] = cacheEntry{
		Decl: declaration,
		File: astFile,
	}

	return extractCommentFromDecl(declaration, object)
}

// ExtractExtraComments retrieves body comments (inside functions or type declarations).
//
// Returns nil if no extra comments are found.
func (e *Extractor) ExtractExtraComments(object types.Object) []model.ExtraComment {
	entry, ok := e.declCache[object]
	if ok {
		// short-circuit already resolved declarations
		if entry.Decl == nil {
			return nil
		}

		return extractExtraCommentsFromDecl(entry.Decl, entry.File, object)
	}

	declaration, astFile := resolveDeclFromObject(object, e.fileSet, e.filesMap)
	if declaration == nil || astFile == nil {
		e.declCache[object] = cacheEntry{}

		return nil
	}

	e.declCache[object] = cacheEntry{
		Decl: declaration,
		File: astFile,
	}

	return extractExtraCommentsFromDecl(declaration, astFile, object)
}

// ExtractDomainDescriptions extracts domain descriptions from package-level comments.
//
// Returns nil if no domain descriptions are found.
func (e *Extractor) ExtractDomainDescriptions() []model.ExtraComment {
	// find the documented file (e.g. doc.go) that contains package-level comments
	var documentedFile *ast.File

	for _, file := range e.syntaxPackage {
		if file.Doc == nil {
			continue
		}

		documentedFile = file
		break
	}

	if documentedFile == nil {
		return nil
	}

	// collect all comment text from the documented file
	var allComments strings.Builder
	for _, group := range documentedFile.Comments {
		if group != nil {
			allComments.WriteString(group.Text())
			allComments.WriteString("\n")
		}
	}

	return parser.ParseDomainDescriptions(allComments.String())
}

func extractCommentFromDecl(declaration ast.Decl, object types.Object) *ast.CommentGroup {
	switch typedDeclaration := declaration.(type) {
	case *ast.GenDecl: // const, var, type declaration
		for _, spec := range typedDeclaration.Specs {
			comment := extractCommentFromTypeSpec(typedDeclaration, spec, object)
			if comment != nil {
				return comment
			}
		}
	case *ast.FuncDecl: // a function
		if typedDeclaration.Doc != nil {
			return typedDeclaration.Doc
		}
	}

	return &ast.CommentGroup{}
}

func extractExtraCommentsFromDecl(declaration ast.Decl, astFile *ast.File, object types.Object) []model.ExtraComment {
	var comments *ast.CommentGroup

	switch typedDeclaration := declaration.(type) {
	case *ast.GenDecl: // const, var, type declaration
		for _, spec := range typedDeclaration.Specs {
			comments = extractExtraCommentFromTypeSpec(astFile, spec, object)
			if comments != nil {
				break
			}
		}
	case *ast.FuncDecl: // a function
		bodyStart := typedDeclaration.Body.Lbrace
		bodyEnd := typedDeclaration.Body.Rbrace
		comments = findCommentInFile(astFile, bodyStart, bodyEnd)
	}

	if comments == nil {
		return nil
	}

	return parser.ParseTaggedComments(comments.Text())
}

func extractCommentFromTypeSpec(typedDeclaration *ast.GenDecl, spec ast.Spec, object types.Object) *ast.CommentGroup {
	switch typedSpec := spec.(type) {
	case *ast.TypeSpec:
		// for TypeSpec, check if it matches (though usually only 1 spec per GenDecl for types)
		if typedSpec.Name.Name != object.Name() {
			return nil
		}

		// return Doc, checking both GenDecl.Doc and TypeSpec.Doc
		if typedDeclaration.Doc != nil {
			return typedDeclaration.Doc
		}

		if typedSpec.Doc != nil {
			return typedSpec.Doc
		}
	case *ast.ValueSpec:
		for _, ident := range typedSpec.Names {
			if ident.Name != object.Name() {
				return nil
			}

			// return Doc, checking both GenDecl.Doc and ValueSpec.Doc
			if typedDeclaration.Doc != nil {
				return typedDeclaration.Doc
			}

			if typedSpec.Doc != nil {
				return typedSpec.Doc
			}
		}
	}

	return nil
}

func extractExtraCommentFromTypeSpec(astFile *ast.File, spec ast.Spec, object types.Object) *ast.CommentGroup {
	typedSpec, isTypeSpec := spec.(*ast.TypeSpec)
	if !isTypeSpec {
		// case *ast.ValueSpec: // not supported for now: only types
		return nil
	}

	// for TypeSpec, check if it matches (though usually only 1 spec per GenDecl for types)
	if typedSpec.Name.Name != object.Name() {
		return nil
	}

	var bodyStart, bodyEnd token.Pos

	switch kind := typedSpec.Type.(type) {
	case *ast.StructType:
		bodyStart = kind.Fields.Opening // position of '{'
		bodyEnd = kind.Fields.Closing   // position of '}'

	case *ast.InterfaceType:
		bodyStart = kind.Methods.Opening // position of '{'
		bodyEnd = kind.Methods.Closing   // position of '}'
		// other type declarations (maps, functions, ...): not supported for now
	}

	if bodyEnd > bodyStart {
		comments := findCommentInFile(astFile, bodyStart, bodyEnd)
		if comments != nil {
			return comments
		}
	}

	return nil
}
