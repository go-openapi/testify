// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package signature

import (
	"go/types"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// Extractor extracts function signatures and formats types with proper package qualification.
type Extractor struct {
	currentPackage *types.Package    // the package being scanned
	importAliases  map[string]string // import path -> alias name used in source
}

// New creates a signature extractor.
//
// The currentPackage is used to determine when type qualification is needed.
// The importAliases map provides the mapping from import paths to the aliases
// used in the source code (e.g., "net/http/httputil" -> "httputil").
func New(currentPackage *types.Package, importAliases map[string]string) *Extractor {
	return &Extractor{
		currentPackage: currentPackage,
		importAliases:  importAliases,
	}
}

// ExtractFunctionSignature extracts function signature details from a types.Signature.
//
// This is used for both regular functions and function types (var/type declarations).
func (e *Extractor) ExtractFunctionSignature(signature *types.Signature, name string) model.Function {
	function := model.Function{
		Name: name,
	}

	// check if generic and extract type parameters
	typeParams := signature.TypeParams()
	function.IsGeneric = typeParams != nil
	if typeParams != nil {
		function.TypeParams = e.extractTypeParams(typeParams)
	}

	// extract parameters
	params := signature.Params()
	function.AllParams = make(model.Parameters, 0, params.Len())
	function.Params = make(model.Parameters, 0, params.Len())

	for param := range params.Variables() {
		p := model.Parameter{
			Name:      param.Name(),
			GoType:    e.ElidedType(param.Type()),
			Selector:  e.ElidedQualifier(param.Type()),
			IsGeneric: isTypeParam(param.Type()),
		}
		function.AllParams = append(function.AllParams, p)

		if p.Name != "t" && p.Name != "msgAndArgs" {
			// filtered params
			function.Params = append(function.Params, p)
		}
	}

	// handle variadic parameters
	l := params.Len()
	if signature.Variadic() && l > 0 {
		// introduce ellipsis notation
		lastParam := params.At(l - 1)
		lastType := lastParam.Type()
		if asSlice, ok := lastType.(*types.Slice); ok { // should be always ok
			elemGoType := e.ElidedType(asSlice.Elem())
			function.AllParams[len(function.AllParams)-1].GoType = "..." + elemGoType
			function.AllParams[len(function.AllParams)-1].IsVariadic = true

			// check filtered params
			if len(function.Params) > 0 && function.Params[len(function.Params)-1].Name == function.AllParams[len(function.AllParams)-1].Name {
				function.Params[len(function.Params)-1].GoType = "..." + elemGoType
				function.Params[len(function.Params)-1].IsVariadic = true
			}
		}
	}

	// extract return values
	results := signature.Results()
	function.Returns = make(model.Parameters, 0, results.Len())
	for result := range results.Variables() {
		function.Returns = append(function.Returns, model.Parameter{
			Name:     result.Name(),
			GoType:   e.ElidedType(result.Type()),
			Selector: e.ElidedQualifier(result.Type()),
		})
	}

	// detect helpers (not assertions)
	function.IsHelper = l == 0 || (len(function.AllParams) > 0 && function.AllParams[0].Name != "t")
	function.IsConstructor = name == "New"

	return function
}

// Qualifier returns the appropriate package name for type qualification.
//
// It uses import aliases from the source (AST) rather than the package's actual name.
func (e *Extractor) Qualifier(pkg *types.Package) string {
	if pkg == nil {
		return ""
	}

	// if it's the current package being scanned, no qualification needed
	if pkg == e.currentPackage {
		return ""
	}

	// look up the alias used in source imports
	if alias, ok := e.importAliases[pkg.Path()]; ok {
		return alias
	}

	// fallback to the package's actual name
	return pkg.Name()
}

// ElidedType returns a string representation of the type with package names as they appear in source.
//
// Uses import aliases when available (e.g., "httputil.Handler" if imported as "httputil").
func (e *Extractor) ElidedType(t types.Type) string {
	return types.TypeString(t, e.Qualifier)
}

// ElidedQualifier returns the selector used for a type, as its import package alias used in source,
// or the empty string if this is a local declaration.
func (e *Extractor) ElidedQualifier(t types.Type) string {
	const maxInterestingParts = 2
	parts := strings.SplitN(types.TypeString(t, e.Qualifier), ".", maxInterestingParts)
	if len(parts) > 1 {
		return parts[0]
	}

	return ""
}

// isTypeParam checks if a type is a type parameter (generic type variable).
func isTypeParam(t types.Type) bool {
	_, ok := t.(*types.TypeParam)
	return ok
}

// extractTypeParams extracts type parameter names and constraints from a TypeParamList.
func (e *Extractor) extractTypeParams(typeParams *types.TypeParamList) []model.TypeParam {
	result := make([]model.TypeParam, typeParams.Len())
	for i := range typeParams.Len() {
		tp := typeParams.At(i)
		result[i] = model.TypeParam{
			Name:       tp.Obj().Name(),
			Constraint: e.ElidedType(tp.Constraint()),
		}
	}
	return result
}
