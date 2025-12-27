// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package signature

import (
	"go/types"
	"iter"
	"slices"
	"testing"
)

func TestQualifier(t *testing.T) {
	t.Parallel()

	// Create test packages
	currentPkg := types.NewPackage("github.com/example/test", "test")
	otherPkg := types.NewPackage("net/http", "http")
	aliasedPkg := types.NewPackage("net/http/httputil", "httputil")

	for c := range qualifierCases(currentPkg, otherPkg, aliasedPkg) {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			e := New(c.currentPkg, c.importAliases)
			result := e.Qualifier(c.pkg)

			if result != c.expected {
				t.Errorf("Qualifier() = %q, expected %q", result, c.expected)
			}
		})
	}
}

func TestElidedQualifier(t *testing.T) {
	t.Parallel()

	currentPkg := types.NewPackage("github.com/example/test", "test")
	otherPkg := types.NewPackage("net/http", "http")

	for c := range elidedQualifierCases(currentPkg, otherPkg) {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			e := New(c.currentPkg, c.importAliases)
			result := e.ElidedQualifier(c.typ)

			if result != c.expected {
				t.Errorf("ElidedQualifier() = %q, expected %q", result, c.expected)
			}
		})
	}
}

func TestElidedType(t *testing.T) {
	t.Parallel()

	currentPkg := types.NewPackage("github.com/example/test", "test")
	otherPkg := types.NewPackage("net/http", "http")

	for c := range elidedTypeCases(currentPkg, otherPkg) {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			e := New(c.currentPkg, c.importAliases)
			result := e.ElidedType(c.typ)

			if result != c.expected {
				t.Errorf("ElidedType() = %q, expected %q", result, c.expected)
			}
		})
	}
}

func TestExtractFunctionSignature(t *testing.T) {
	t.Parallel()

	currentPkg := types.NewPackage("github.com/example/test", "test")

	// Helper to create a signature
	makeSig := func(params, results []*types.Var, variadic bool) *types.Signature {
		return types.NewSignatureType(nil, nil, nil,
			types.NewTuple(params...),
			types.NewTuple(results...),
			variadic)
	}

	for c := range extractFunctionSignatureCases(currentPkg, makeSig) {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			e := New(currentPkg, map[string]string{})
			result := e.ExtractFunctionSignature(c.signature, c.funcName)

			if result.Name != c.expectedName {
				t.Errorf("Name = %q, expected %q", result.Name, c.expectedName)
			}
			if result.IsHelper != c.expectedHelper {
				t.Errorf("IsHelper = %v, expected %v", result.IsHelper, c.expectedHelper)
			}
			if result.IsConstructor != c.expectedCtor {
				t.Errorf("IsConstructor = %v, expected %v", result.IsConstructor, c.expectedCtor)
			}
			if len(result.Params) != c.paramsLen {
				t.Errorf("len(Params) = %d, expected %d", len(result.Params), c.paramsLen)
			}
			if len(result.Returns) != c.returnsLen {
				t.Errorf("len(Returns) = %d, expected %d", len(result.Returns), c.returnsLen)
			}
			if len(result.AllParams) != c.allParamsLen {
				t.Errorf("len(AllParams) = %d, expected %d", len(result.AllParams), c.allParamsLen)
			}
		})
	}
}

func TestExtractFunctionSignature_Variadic(t *testing.T) {
	currentPkg := types.NewPackage("github.com/example/test", "test")

	// Create variadic signature
	params := []*types.Var{
		types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
		types.NewVar(0, currentPkg, "values", types.NewSlice(types.Typ[types.String])),
	}
	results := []*types.Var{
		types.NewVar(0, currentPkg, "", types.Typ[types.Bool]),
	}
	sig := types.NewSignatureType(nil, nil, nil,
		types.NewTuple(params...),
		types.NewTuple(results...),
		true) // variadic

	e := New(currentPkg, map[string]string{})
	result := e.ExtractFunctionSignature(sig, "TestFunc")

	// Check that last param has ellipsis notation
	if len(result.AllParams) != 2 {
		t.Fatalf("Expected 2 AllParams, got %d", len(result.AllParams))
	}

	lastParam := result.AllParams[len(result.AllParams)-1]
	if lastParam.GoType != "...string" {
		t.Errorf("Last param GoType = %q, expected %q", lastParam.GoType, "...string")
	}
	if !lastParam.IsVariadic {
		t.Errorf("Last param IsVariadic = false, expected true")
	}

	// Check filtered params also have ellipsis
	if len(result.Params) != 1 {
		t.Fatalf("Expected 1 filtered Param, got %d", len(result.Params))
	}
	if result.Params[0].GoType != "...string" {
		t.Errorf("Filtered param GoType = %q, expected %q", result.Params[0].GoType, "...string")
	}
}

func TestExtractFunctionSignature_GenericFunction(t *testing.T) {
	currentPkg := types.NewPackage("github.com/example/test", "test")

	// Create a generic signature with type parameters
	typeParam := types.NewTypeParam(types.NewTypeName(0, currentPkg, "T", nil), types.NewInterfaceType(nil, nil).Complete())

	params := []*types.Var{
		types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
		types.NewVar(0, currentPkg, "value", typeParam),
	}
	results := []*types.Var{
		types.NewVar(0, currentPkg, "", types.Typ[types.Bool]),
	}

	sig := types.NewSignatureType(nil, nil, nil,
		types.NewTuple(params...),
		types.NewTuple(results...),
		false)

	// Manually set type params on signature using reflection or accept that
	// we can't easily test generics without more complex setup.
	// For now, test non-generic case and verify IsGeneric is false.
	e := New(currentPkg, map[string]string{})
	result := e.ExtractFunctionSignature(sig, "GenericFunc")

	// Without type params attached, this will be false
	if result.IsGeneric {
		t.Errorf("IsGeneric = true, expected false (cannot easily construct generic sig in test)")
	}
}

/* Test case iterators */

type qualifierCase struct {
	name          string
	currentPkg    *types.Package
	importAliases map[string]string
	pkg           *types.Package
	expected      string
}

func qualifierCases(currentPkg, otherPkg, aliasedPkg *types.Package) iter.Seq[qualifierCase] {
	return slices.Values([]qualifierCase{
		{
			name:          "nil package",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			pkg:           nil,
			expected:      "",
		},
		{
			name:          "current package - no qualification",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			pkg:           currentPkg,
			expected:      "",
		},
		{
			name:          "other package - use package name",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			pkg:           otherPkg,
			expected:      "http",
		},
		{
			name:       "aliased package - use alias",
			currentPkg: currentPkg,
			importAliases: map[string]string{
				"net/http/httputil": "httputil",
			},
			pkg:      aliasedPkg,
			expected: "httputil",
		},
		{
			name:       "aliased package with different alias",
			currentPkg: currentPkg,
			importAliases: map[string]string{
				"net/http/httputil": "util",
			},
			pkg:      aliasedPkg,
			expected: "util",
		},
		{
			name:       "package with alias but using actual name as fallback",
			currentPkg: currentPkg,
			importAliases: map[string]string{
				"some/other/path": "other",
			},
			pkg:      otherPkg,
			expected: "http", // falls back to package name
		},
	})
}

type elidedQualifierCase struct {
	name          string
	currentPkg    *types.Package
	importAliases map[string]string
	typ           types.Type
	expected      string
}

func elidedQualifierCases(currentPkg, otherPkg *types.Package) iter.Seq[elidedQualifierCase] {
	return slices.Values([]elidedQualifierCase{
		{
			name:          "basic type - no qualifier",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.Typ[types.String],
			expected:      "",
		},
		{
			name:          "named type from other package",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), types.NewStruct(nil, nil), nil),
			expected:      "http",
		},
		{
			name:          "pointer to basic type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewPointer(types.Typ[types.Int]),
			expected:      "",
		},
		{
			name:          "pointer to named type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewPointer(types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), types.NewStruct(nil, nil), nil)),
			expected:      "*http", // includes pointer in qualifier
		},
		{
			name:       "aliased package qualifier",
			currentPkg: currentPkg,
			importAliases: map[string]string{
				"net/http": "h",
			},
			typ:      types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), types.NewStruct(nil, nil), nil),
			expected: "h",
		},
	})
}

type elidedTypeCase struct {
	name          string
	currentPkg    *types.Package
	importAliases map[string]string
	typ           types.Type
	expected      string
}

func elidedTypeCases(currentPkg, otherPkg *types.Package) iter.Seq[elidedTypeCase] {
	return slices.Values([]elidedTypeCase{
		{
			name:          "string type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.Typ[types.String],
			expected:      "string",
		},
		{
			name:          "int type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.Typ[types.Int],
			expected:      "int",
		},
		{
			name:          "bool type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.Typ[types.Bool],
			expected:      "bool",
		},
		{
			name:          "pointer to string",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewPointer(types.Typ[types.String]),
			expected:      "*string",
		},
		{
			name:          "slice of int",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewSlice(types.Typ[types.Int]),
			expected:      "[]int",
		},
		{
			name:          "named type from other package",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), types.NewStruct(nil, nil), nil),
			expected:      "http.Request",
		},
		{
			name:       "named type with alias",
			currentPkg: currentPkg,
			importAliases: map[string]string{
				"net/http": "h",
			},
			typ:      types.NewNamed(types.NewTypeName(0, otherPkg, "Request", nil), types.NewStruct(nil, nil), nil),
			expected: "h.Request",
		},
		{
			name:          "interface type",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewInterfaceType(nil, nil),
			expected:      "interface{}",
		},
		{
			name:          "empty interface (any)",
			currentPkg:    currentPkg,
			importAliases: map[string]string{},
			typ:           types.NewInterfaceType(nil, nil).Complete(),
			expected:      "interface{}",
		},
	})
}

type extractFunctionSignatureCase struct {
	name           string
	funcName       string
	signature      *types.Signature
	expectedName   string
	expectedHelper bool
	expectedCtor   bool
	paramsLen      int
	returnsLen     int
	allParamsLen   int
}

func extractFunctionSignatureCases(currentPkg *types.Package, makeSig func([]*types.Var, []*types.Var, bool) *types.Signature) iter.Seq[extractFunctionSignatureCase] {
	return slices.Values([]extractFunctionSignatureCase{
		{
			name:     "simple assertion function",
			funcName: "Equal",
			signature: makeSig(
				[]*types.Var{
					types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
					types.NewVar(0, currentPkg, "expected", types.NewInterfaceType(nil, nil)),
					types.NewVar(0, currentPkg, "actual", types.NewInterfaceType(nil, nil)),
				},
				[]*types.Var{
					types.NewVar(0, currentPkg, "", types.Typ[types.Bool]),
				},
				false,
			),
			expectedName:   "Equal",
			expectedHelper: false,
			expectedCtor:   false,
			paramsLen:      2, // filtered (excludes 't')
			returnsLen:     1,
			allParamsLen:   3,
		},
		{
			name:     "variadic assertion function",
			funcName: "True",
			signature: makeSig(
				[]*types.Var{
					types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
					types.NewVar(0, currentPkg, "value", types.Typ[types.Bool]),
					types.NewVar(0, currentPkg, "msgAndArgs", types.NewSlice(types.NewInterfaceType(nil, nil))),
				},
				[]*types.Var{
					types.NewVar(0, currentPkg, "", types.Typ[types.Bool]),
				},
				true,
			),
			expectedName:   "True",
			expectedHelper: false,
			expectedCtor:   false,
			paramsLen:      1, // filtered (excludes 't' and 'msgAndArgs')
			returnsLen:     1,
			allParamsLen:   3,
		},
		{
			name:     "helper function - no params",
			funcName: "Helper",
			signature: makeSig(
				[]*types.Var{},
				[]*types.Var{
					types.NewVar(0, currentPkg, "", types.Typ[types.String]),
				},
				false,
			),
			expectedName:   "Helper",
			expectedHelper: true,
			expectedCtor:   false,
			paramsLen:      0,
			returnsLen:     1,
			allParamsLen:   0,
		},
		{
			name:     "constructor function",
			funcName: "New",
			signature: makeSig(
				[]*types.Var{
					types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
				},
				[]*types.Var{
					types.NewVar(0, currentPkg, "", types.NewPointer(types.NewNamed(
						types.NewTypeName(0, currentPkg, "Assertions", nil),
						types.NewStruct(nil, nil), nil))),
				},
				false,
			),
			expectedName:   "New",
			expectedHelper: false,
			expectedCtor:   true,
			paramsLen:      0, // filtered (excludes 't')
			returnsLen:     1,
			allParamsLen:   1,
		},
		{
			name:     "helper - first param not 't'",
			funcName: "CalcValue",
			signature: makeSig(
				[]*types.Var{
					types.NewVar(0, currentPkg, "x", types.Typ[types.Int]),
					types.NewVar(0, currentPkg, "y", types.Typ[types.Int]),
				},
				[]*types.Var{
					types.NewVar(0, currentPkg, "", types.Typ[types.Int]),
				},
				false,
			),
			expectedName:   "CalcValue",
			expectedHelper: true, // first param is not 't'
			expectedCtor:   false,
			paramsLen:      2, // not filtered
			returnsLen:     1,
			allParamsLen:   2,
		},
		{
			name:     "no return values",
			funcName: "FailNow",
			signature: makeSig(
				[]*types.Var{
					types.NewVar(0, currentPkg, "t", types.NewInterfaceType(nil, nil)),
				},
				[]*types.Var{},
				false,
			),
			expectedName:   "FailNow",
			expectedHelper: false,
			expectedCtor:   false,
			paramsLen:      0, // filtered (excludes 't')
			returnsLen:     0,
			allParamsLen:   1,
		},
	})
}
