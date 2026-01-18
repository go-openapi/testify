// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"regexp"

	"golang.org/x/tools/go/packages"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
	"github.com/go-openapi/testify/codegen/v2/internal/scanner/comments"
	parser "github.com/go-openapi/testify/codegen/v2/internal/scanner/comments-parser"
	"github.com/go-openapi/testify/codegen/v2/internal/scanner/signature"
)

// Scanner scans the internal/assertions package and extracts all functions, types, and metadata.
type Scanner struct {
	options

	scanConfig       *packages.Config
	syntaxPackage    []*ast.File
	typedPackage     *types.Package
	typesInfo        *types.Info
	filesMap         map[*token.File]*ast.File
	fileSet          *token.FileSet
	importAliases    map[string]string // import path -> alias name used in source
	sigExtractor     *signature.Extractor
	commentExtractor *comments.Extractor
	result           *model.AssertionPackage
}

// New [Scanner] ready to [Scanner.Scan] code.
func New(opts ...Option) *Scanner {
	o := optionsWithDefaults(opts)

	const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps |
		packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

	cfg := &packages.Config{
		Dir:   o.dir,
		Mode:  pkgLoadMode,
		Tests: false, // we infer generated tests from production code and ignore internal tests
	}

	return &Scanner{
		options:    o,
		scanConfig: cfg,
		result:     model.New(),
	}
}

// Scan a source package an returns a [model.AssertionPackage] data structure suited for code generation.
func (s *Scanner) Scan() (*model.AssertionPackage, error) {
	pkgs, err := packages.Load(s.scanConfig, s.pkg)
	if err != nil {
		return nil, err
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("could not parse package: %v", s.pkg)
	}

	// we consider only one package
	pkg := pkgs[0]
	s.syntaxPackage = pkg.Syntax
	s.typedPackage = pkg.Types
	if s.typedPackage == nil {
		return nil, fmt.Errorf("something is wrong. Could not parse packages: %v", s.pkg)
	}

	// stash everything we need from [packages.Package]
	s.typesInfo = pkg.TypesInfo
	s.filesMap = comments.BuildFilesMap(pkg)
	s.fileSet = pkg.Fset
	s.importAliases = buildImportAliases(pkg)
	s.sigExtractor = signature.New(s.typedPackage, s.importAliases)
	s.commentExtractor = comments.New(s.syntaxPackage, s.fileSet, s.filesMap)

	if err := s.resolveScope(); err != nil {
		return nil, err
	}

	// post-collection enrichments
	s.result.Tool = moduleName()
	s.result.Header = header()
	s.result.Copyright = s.commentExtractor.ExtractCopyright().Text()
	s.result.DocString = s.commentExtractor.ExtractPackageComments().Text()
	s.result.TestDataPath = filepath.Join(pkg.Dir, "testdata")

	if s.collectDoc {
		// extract domain descriptions from package comments
		s.result.ExtraComments = s.commentExtractor.ExtractDomainDescriptions()
	}

	return s.result, nil
}

func (s *Scanner) resolveScope() error {
	s.result.Package = s.typedPackage.Path()

	for _, pkg := range s.typedPackage.Imports() {
		s.addImport(pkg)
	}

	scope := s.typedPackage.Scope()
	for _, symbol := range scope.Names() {
		object := scope.Lookup(symbol)
		if object == nil { // something went wrong
			continue
		}
		if !object.Exported() {
			continue
		}

		switch typedObject := object.(type) {
		case *types.Func:
			s.addFunction(typedObject)
		case *types.Const:
			s.addConst(typedObject)
		case *types.Var:
			s.addVar(typedObject)
		case *types.TypeName:
			s.addNamedType(typedObject)
		// NOTE: known limitation this does not support type or var aliases in the source package
		default:
			continue
		}
	}

	// Accumulate parse errors from test values
	var parseErrors []error
	for _, fn := range s.result.Functions {
		for _, test := range fn.Tests {
			for _, val := range test.TestedValues {
				if val.Error != nil {
					parseErrors = append(parseErrors, fmt.Errorf("function %s: invalid test value %q: %w", fn.Name, val.Raw, val.Error))
				}
			}
		}
	}

	if len(parseErrors) > 0 {
		return errors.Join(parseErrors...)
	}

	return nil
}

func (s *Scanner) addImport(pkg *types.Package) {
	if s.result.Imports == nil {
		s.result.Imports = make(model.ImportMap)
	}

	// use the alias if one exists, otherwise use the package name.
	// (this ensures the key matches what elidedQualifier() returns)
	key := pkg.Name() // default to package name
	if alias, ok := s.importAliases[pkg.Path()]; ok {
		key = alias // use the alias from source
	}

	s.result.Imports[key] = pkg.Path()
}

func (s *Scanner) addFunction(object *types.Func) {
	docComment := s.commentExtractor.ExtractComments(object)

	// extract the base function signature
	function := s.sigExtractor.ExtractFunctionSignature(object.Signature(), object.Name())

	// add function-specific metadata
	function.ID = object.Id()
	function.SourcePackage = object.Pkg().Path() // full package name
	function.TargetPackage = object.Pkg().Name() // short package name
	function.DocString = docComment.Text()
	function.Tests = parser.ParseTestExamples(docComment.Text())
	if s.collectDoc {
		function.ExtraComments = s.commentExtractor.ExtractExtraComments(object)
		pos := s.fileSet.Position(object.Pos())
		function.SourceLink = &pos
		function.Domain = parser.DomainFromExtraComments(function.ExtraComments)
	}
	function.IsDeprecated = isDeprecated(function.DocString)

	s.result.Functions = append(s.result.Functions, function)
}

func (s *Scanner) addConst(object *types.Const) {
	constant := model.Ident{
		ID:            object.Id(),
		Name:          object.Name(),
		SourcePackage: object.Pkg().Path(),
		TargetPackage: object.Pkg().Name(), // short package name
		DocString:     s.commentExtractor.ExtractComments(object).Text(),
	}
	constant.IsDeprecated = isDeprecated(constant.DocString)

	if s.collectDoc {
		pos := s.fileSet.Position(object.Pos())
		constant.SourceLink = &pos
	}
	s.result.Consts = append(s.result.Consts, constant)
}

func (s *Scanner) addVar(object *types.Var) {
	variable := model.Ident{
		ID:            object.Id(),
		Name:          object.Name(),
		SourcePackage: object.Pkg().Path(),
		TargetPackage: object.Pkg().Name(), // short package name
		DocString:     s.commentExtractor.ExtractComments(object).Text(),
	}
	variable.IsDeprecated = isDeprecated(variable.DocString)

	if s.collectDoc {
		pos := s.fileSet.Position(object.Pos())
		variable.SourceLink = &pos
	}

	// Check if this variable is a function type
	if sig, ok := object.Type().(*types.Signature); ok {
		fn := s.sigExtractor.ExtractFunctionSignature(sig, object.Name())
		variable.Function = &fn
	}

	s.result.Vars = append(s.result.Vars, variable)
}

func (s *Scanner) addNamedType(object *types.TypeName) {
	namedType := model.Ident{
		ID:            object.Id(),
		Name:          object.Name(),
		SourcePackage: object.Pkg().Path(),
		TargetPackage: object.Pkg().Name(), // short package name
		IsAlias:       object.IsAlias(),
		DocString:     s.commentExtractor.ExtractComments(object).Text(),
	}
	namedType.IsDeprecated = isDeprecated(namedType.DocString)

	// check if this named type is a function type.
	// For named types, we need to check the underlying type
	if named, ok := object.Type().(*types.Named); ok {
		if sig, ok := named.Underlying().(*types.Signature); ok {
			fn := s.sigExtractor.ExtractFunctionSignature(sig, object.Name())
			namedType.Function = &fn
		}
	} else if sig, ok := object.Type().(*types.Signature); ok {
		// type alias to a function type
		fn := s.sigExtractor.ExtractFunctionSignature(sig, object.Name())
		namedType.Function = &fn
	}

	if s.collectDoc {
		pos := s.fileSet.Position(object.Pos())
		namedType.SourceLink = &pos
		namedType.ExtraComments = s.commentExtractor.ExtractExtraComments(object)
		namedType.Domain = parser.DomainFromExtraComments(namedType.ExtraComments)
	}

	s.result.Types = append(s.result.Types, namedType)
}

// This is the regexp used by pkgsite.
// See: https://cs.opensource.google/go/x/pkgsite/+/master:internal/godoc/dochtml/deprecated.go
//
// "Deprecated:" at the start of a paragraph.
var deprecatedRx = regexp.MustCompile(`(^|\n\s*\n)\s*Deprecated:`)

func isDeprecated(comment string) bool {
	return deprecatedRx.MatchString(comment)
}
