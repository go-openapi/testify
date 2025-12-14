// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

const (
	successPrefix = "success:"
	failurePrefix = "failure:"
	panicPrefix   = "panic:"
)

// Scanner scans the internal/assertions package and extracts all functions, types, and metadata.
type Scanner struct {
	options

	scanConfig    *packages.Config
	syntaxPackage []*ast.File
	typedPackage  *types.Package
	typesInfo     *types.Info
	filesMap      map[*token.File]*ast.File
	fileSet       *token.FileSet
	importAliases map[string]string // import path -> alias name used in source
	result        *model.AssertionPackage
}

// New [Scanner] ready to [Scanner.Scan] code.
func New(opts ...Option) *Scanner {
	o := optionsWithDefaults(opts)

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
	s.typesInfo = pkg.TypesInfo
	s.filesMap = buildFilesMap(pkg)
	s.fileSet = pkg.Fset
	s.importAliases = buildImportAliases(pkg)

	s.resolveScope()
	s.result.Tool = moduleName()
	s.result.Copyright = s.extractCopyright().Text()
	s.result.DocString = s.extractPackageComments().Text()
	s.result.TestDataPath = filepath.Join(pkg.Dir, "testdata")

	return s.result, nil
}

func (s *Scanner) resolveScope() {
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
}

// buildFilesMap constructs a lookup index to help bridge token position vs ast.File.
func buildFilesMap(pkg *packages.Package) map[*token.File]*ast.File {
	filesMap := make(map[*token.File]*ast.File, len(pkg.Syntax))

	for _, astFile := range pkg.Syntax {
		tokenFile := pkg.Fset.File(astFile.Pos()) // o(log n) lookup
		filesMap[tokenFile] = astFile
	}

	return filesMap
}

// buildImportAliases scans import declarations to find aliases used in the source.
//
// This bridges the AST view (import aliases) with the types view (package paths).
func buildImportAliases(pkg *packages.Package) map[string]string {
	aliases := make(map[string]string)

	for _, astFile := range pkg.Syntax {
		for _, importSpec := range astFile.Imports {
			// Get the import path (remove quotes)
			importPath := strings.Trim(importSpec.Path.Value, `"`)

			var alias string
			if importSpec.Name != nil {
				// Explicit alias: import foo "bar/baz"
				alias = importSpec.Name.Name

				// Skip blank imports and dot imports for type qualification
				if alias == "_" || alias == "." {
					continue
				}
			} else {
				// No explicit alias - need to determine the package name
				// Try to find it in the loaded imports
				for _, imported := range pkg.Imports {
					if imported.PkgPath == importPath {
						alias = imported.Name
						break
					}
				}

				// Fallback: use last segment of import path
				if alias == "" {
					parts := strings.Split(importPath, "/")
					alias = parts[len(parts)-1]
				}
			}

			// Store the mapping (later imports override earlier ones if there are conflicts)
			aliases[importPath] = alias
		}
	}

	return aliases
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
	docComment := s.extractComments(object)

	// Extract the base function signature
	function := s.extractFunctionSignature(object.Signature(), object.Name())

	// Add function-specific metadata
	function.ID = object.Id()
	function.SourcePackage = object.Pkg().Path() // full package name
	function.TargetPackage = object.Pkg().Name() // short package name
	function.DocString = docComment.Text()
	function.Tests = parseTestExamples(docComment) // extract test values from Example section

	s.result.Functions = append(s.result.Functions, function)
}

func (s *Scanner) addConst(object *types.Const) {
	constant := model.Ident{
		ID:            object.Id(),
		Name:          object.Name(),
		SourcePackage: object.Pkg().Path(),
		TargetPackage: object.Pkg().Name(), // short package name
		DocString:     s.extractComments(object).Text(),
	}
	s.result.Consts = append(s.result.Consts, constant)
}

func (s *Scanner) addVar(object *types.Var) {
	variable := model.Ident{
		ID:            object.Id(),
		Name:          object.Name(),
		SourcePackage: object.Pkg().Path(),
		TargetPackage: object.Pkg().Name(), // short package name
		DocString:     s.extractComments(object).Text(),
	}

	// Check if this variable is a function type
	if sig, ok := object.Type().(*types.Signature); ok {
		fn := s.extractFunctionSignature(sig, object.Name())
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
		DocString:     s.extractComments(object).Text(),
	}

	// Check if this named type is a function type
	// For named types, we need to check the underlying type
	if named, ok := object.Type().(*types.Named); ok {
		if sig, ok := named.Underlying().(*types.Signature); ok {
			fn := s.extractFunctionSignature(sig, object.Name())
			namedType.Function = &fn
		}
	} else if sig, ok := object.Type().(*types.Signature); ok {
		// Type alias to a function type
		fn := s.extractFunctionSignature(sig, object.Name())
		namedType.Function = &fn
	}

	s.result.Types = append(s.result.Types, namedType)
}

// extractFunctionSignature extracts function signature details from a types.Signature.
//
// This is used for both regular functions and function types (var/type declarations).
func (s *Scanner) extractFunctionSignature(signature *types.Signature, name string) model.Function {
	function := model.Function{
		Name: name,
	}

	// Check if generic
	isGeneric := signature.TypeParams() != nil
	function.IsGeneric = isGeneric

	// Extract parameters
	params := signature.Params()
	function.AllParams = make(model.Parameters, 0, params.Len())
	function.Params = make(model.Parameters, 0, params.Len())

	for param := range params.Variables() {
		p := model.Parameter{
			Name:     param.Name(),
			GoType:   s.elidedType(param.Type()),
			Selector: s.elidedQualifier(param.Type()),
		}
		function.AllParams = append(function.AllParams, p)

		if p.Name != "t" && p.Name != "msgAndArgs" {
			// filtered params
			function.Params = append(function.Params, p)
		}
	}

	// Handle variadic parameters
	l := params.Len()
	if signature.Variadic() && l > 0 {
		// introduce ellipsis notation
		lastParam := params.At(l - 1)
		lastType := lastParam.Type()
		if asSlice, ok := lastType.(*types.Slice); ok { // should be always ok
			elemGoType := s.elidedType(asSlice.Elem())
			function.AllParams[len(function.AllParams)-1].GoType = "..." + elemGoType
			function.AllParams[len(function.AllParams)-1].IsVariadic = true

			// check filtered params
			if len(function.Params) > 0 && function.Params[len(function.Params)-1].Name == function.AllParams[len(function.AllParams)-1].Name {
				function.Params[len(function.Params)-1].GoType = "..." + elemGoType
				function.Params[len(function.Params)-1].IsVariadic = true
			}
		}
	}

	// Extract return values
	results := signature.Results()
	function.Returns = make(model.Parameters, 0, results.Len())
	for result := range results.Variables() {
		function.Returns = append(function.Returns, model.Parameter{
			Name:     result.Name(),
			GoType:   s.elidedType(result.Type()),
			Selector: s.elidedQualifier(result.Type()),
		})
	}

	// Detect helpers (not assertions)
	function.IsHelper = l == 0 || (len(function.AllParams) > 0 && function.AllParams[0].Name != "t")
	function.IsConstructor = name == "New"

	return function
}

// qualifier returns the appropriate package name for type qualification.
//
// It uses import aliases from the source (AST) rather than the package's actual name.
func (s *Scanner) qualifier(pkg *types.Package) string {
	if pkg == nil {
		return ""
	}

	// If it's the current package being scanned, no qualification needed
	if pkg == s.typedPackage {
		return ""
	}

	// Look up the alias used in source imports
	if alias, ok := s.importAliases[pkg.Path()]; ok {
		return alias
	}

	// Fallback to the package's actual name
	return pkg.Name()
}

// elidedType returns a string representation of the type with package names as they appear in source.
//
// Uses import aliases when available (e.g., "httputil.Handler" if imported as "httputil").
func (s *Scanner) elidedType(t types.Type) string {
	return types.TypeString(t, s.qualifier)
}

// elidedQualifier returns the selector used for a type, as its import package alias used in source,
// or the empty string if this is a local declaration.
func (s *Scanner) elidedQualifier(t types.Type) string {
	const maxInterestingParts = 2
	parts := strings.SplitN(types.TypeString(t, s.qualifier), ".", maxInterestingParts)
	if len(parts) > 1 {
		return parts[0]
	}

	return ""
}

// extractPackageComments finds the package-level comment that
// describes the source package.
func (s *Scanner) extractPackageComments() *ast.CommentGroup {
	for _, file := range s.syntaxPackage {
		if file.Doc == nil {
			continue
		}

		return file.Doc
	}

	// safeguard: never returns nil
	return &ast.CommentGroup{}
}

var copyrightRex = regexp.MustCompile(`(?i)copyright`)

// extractCopyright scans source files and expects that at least one of those contains
// a leading comment line with the string "copyright". This comment group is assigned
// to the copyright heading that will be retained.
func (s *Scanner) extractCopyright() *ast.CommentGroup {
FILE:
	for _, file := range s.syntaxPackage {
		if len(file.Comments) == 0 {
			continue
		}

		firstGroup := file.Comments[0]

		for _, line := range firstGroup.List {
			if line.Slash >= file.Package {
				// copyright comment must be before the "package" stanza
				continue FILE
			}

			if copyrightRex.MatchString(line.Text) {
				// first file found with leading comments
				return firstGroup
			}
		}
	}

	// safeguard: never returns nil
	return &ast.CommentGroup{}
}

// extractComments retrieves the docstring for an object.
//
// It never returns nil: if no comment is present an empty but valid [ast.CommentGroup] is returned.
//
// NOTE: comment prioritization rule is:
//  1. GenDecl.Doc = leading comment for the entire declaration block
//  2. Spec.Doc = individual doc for that specific spec
func (s *Scanner) extractComments(object types.Object) *ast.CommentGroup {
	pos := object.Pos()
	if !pos.IsValid() {
		return &ast.CommentGroup{}
	}

	tokenFile := s.fileSet.File(pos) // o(log n) lookup
	if tokenFile == nil {
		return &ast.CommentGroup{}
	}

	astFile := s.filesMap[tokenFile] // o(1) lookup
	if astFile == nil {
		return &ast.CommentGroup{}
	}

	path, _ := astutil.PathEnclosingInterval(astFile, pos, pos)
	for _, node := range path {
		declaration, ok := node.(ast.Decl)
		if !ok {
			continue
		}

		return extractCommentFromDecl(declaration, object)
	}

	return &ast.CommentGroup{}
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

// parseTestExamples extracts test examples from doc comments.
//
// It looks for an "Examples:" or "Example:" section and parses lines like:
//   - success: <test values>
//   - failure: <test values>
//   - panic: <assertion message>
//
// Assumption: we use regular "//" comments, not "/* ... */".
func parseTestExamples(docComment *ast.CommentGroup) []model.Test {
	if docComment == nil {
		return nil
	}

	const usualNumberOfExamples = 2
	tests := make([]model.Test, 0, usualNumberOfExamples)

	inExamplesSection := false // comments are expected to contain an Example[s] or # Example[s] section
	inValueSection := false

	for _, comment := range docComment.List {
		text := strings.TrimPrefix(comment.Text, "//")
		text = strings.TrimSpace(text)

		if inValueSection && len(tests) > 0 && text != "" && !startAnotherSection(text) && !startExampleValue(text) {
			// check if a comment line follows an example value: this would the assertion message
			tests[len(tests)-1].AssertionMessage = text
			inValueSection = false

			continue
		}

		inValueSection = false

		// Check if we're entering the Examples section
		if startExampleSection(text) {
			inExamplesSection = true
			continue
		}

		// Skip until we find the Examples section
		if !inExamplesSection {
			continue
		}

		// Stop if we hit another section (starts with capital letter and ends with colon)
		if startAnotherSection(text) {
			break
		}

		// parse test lines: "success: { values to put in the test }", "failure: ..." or "panic: ..."
		// After each value line, we may found an additional text to be used as an assertion message.
		switch {
		case strings.HasPrefix(text, successPrefix):
			inValueSection = true
			if testcase, ok := parseTestValue(successPrefix, model.TestSuccess, text); ok {
				tests = append(tests, testcase)
			}
		case strings.HasPrefix(text, failurePrefix):
			inValueSection = true
			if testcase, ok := parseTestValue(failurePrefix, model.TestFailure, text); ok {
				tests = append(tests, testcase)
			}
		case strings.HasPrefix(text, panicPrefix):
			inValueSection = true
			if testcase, ok := parseTestValue(panicPrefix, model.TestPanic, text); ok {
				tests = append(tests, testcase)
			}
		}
	}

	return tests
}

// moduleName returns the main module name of the caller.
//
// This identifies the tool currently running the analysis.
func moduleName() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	var (
		modVersion string
		gitVersion string
	)

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			gitVersion = setting.Value
		}
	}

	if info.Main.Version == "(devel)" {
		modVersion = "master"
	} else {
		modVersion = info.Main.Version
	}

	final := info.Main.Path + "@" + modVersion
	if gitVersion != "" {
		final += " [sha: " + gitVersion + "]"
	}

	return final
}

func startExampleSection(text string) bool {
	return strings.HasPrefix(text, "Examples:") ||
		strings.HasPrefix(text, "Example:") ||
		strings.HasPrefix(text, "# Example") ||
		strings.HasPrefix(text, "# Examples")
}

func startAnotherSection(text string) bool {
	return len(text) > 0 && (text[0] >= 'A' && text[0] <= 'Z' && strings.HasSuffix(text, ":")) || strings.HasPrefix(text, "# ")
}

func startExampleValue(text string) bool {
	return strings.HasPrefix(text, successPrefix) || strings.HasPrefix(text, failurePrefix) || strings.HasPrefix(text, panicPrefix)
}

func parseTestValue(placeHolder string, outcome model.TestExpectedOutcome, text string) (model.Test, bool) {
	value := strings.TrimSpace(strings.TrimPrefix(text, placeHolder))
	_, isTodo := strings.CutPrefix(value, "// TODO")

	if value != "" && !isTodo {
		return model.Test{
			TestedValue:     value,
			ExpectedOutcome: outcome,
		}, true
	}

	return model.Test{}, false
}
