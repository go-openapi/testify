// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"go/ast"
	"go/token"
	"maps"
	"slices"
	"strings"
)

// AssertionPackage describes the internal/assertions package.
type AssertionPackage struct {
	Tool             string
	Header           string
	Package          string
	DocString        string
	Copyright        string
	Receiver         string
	TestDataPath     string
	Imports          ImportMap
	EnableFormat     bool
	EnableForward    bool
	EnableGenerics   bool
	EnableExamples   bool
	RunnableExamples bool

	Functions []Function
	Types     []Ident
	Consts    []Ident
	Vars      []Ident

	// extraneous information when scanning in collectDoc mode
	ExtraComments []ExtraComment
}

func (a *AssertionPackage) HasHelpers() (ok bool) {
	for _, fn := range a.Functions {
		if fn.IsHelper || fn.IsConstructor {
			return true
		}
	}

	return false
}

// New empty [AssertionPackage].
func New() *AssertionPackage {
	const (
		allocatedFuncs  = 100
		allocatedIdents = 50
	)

	return &AssertionPackage{
		// preallocate with sensible defaults for our package
		Functions: make([]Function, 0, allocatedFuncs),
		Types:     make([]Ident, 0, allocatedIdents),
		Consts:    make([]Ident, 0, allocatedIdents),
		Vars:      make([]Ident, 0, allocatedIdents),
	}
}

func (a *AssertionPackage) Clone() *AssertionPackage {
	b := *a
	maps.Copy(b.Imports, a.Imports)

	return &b
}

// ImportMap represents the imports for the analyzed package.
type ImportMap map[string]string

func (m ImportMap) HasImports() bool {
	return len(m) > 0
}

// Function represents an assertion function extracted from the source package.
type Function struct {
	ID            string
	Name          string
	SourcePackage string
	TargetPackage string
	DocString     string
	UseMock       string
	Params        Parameters
	AllParams     Parameters
	Returns       Parameters
	TypeParams    []TypeParam
	IsGeneric     bool
	IsHelper      bool
	IsDeprecated  bool
	IsConstructor bool
	Tests         []Test
	// extraneous information when scanning in collectDoc mode
	Domain        string
	SourceLink    *token.Position
	ExtraComments []ExtraComment
}

// GenericName renders the function name with one or more suffixes,
// accounting for any type parameter for generic functions.
func (f Function) GenericName(suffixes ...string) string {
	suffix := strings.Join(suffixes, "")
	if !f.IsGeneric { // means len(f.TypeParams) > 0
		return f.Name + suffix
	}

	var w strings.Builder
	w.WriteString(f.Name)
	w.WriteString(suffix)
	w.WriteByte('[')
	c := f.TypeParams[0]
	w.WriteString(c.Name)
	if len(f.TypeParams) <= 1 || f.TypeParams[1].Constraint != c.Constraint {
		// constraint is elided if next param has the same constraint
		w.WriteByte(' ')
		w.WriteString(c.Constraint)
	}

	for i, p := range f.TypeParams[1:] {
		w.WriteString(", ")
		w.WriteString(p.Name)
		if len(f.TypeParams) <= i+1+1 || f.TypeParams[i+1].Constraint != p.Constraint {
			w.WriteByte(' ')
			w.WriteString(p.Constraint)
		}
	}
	w.WriteByte(']')

	return w.String()
}

func (f Function) HasTest() bool {
	return len(f.Tests) > 0
}

func (f Function) HasSuccessTest() bool {
	return slices.ContainsFunc(f.Tests, func(e Test) bool {
		return e.ExpectedOutcome == TestSuccess
	})
}

type Parameters []Parameter

// Parameter represents a function parameter or return value.
type Parameter struct {
	Name       string
	GoType     string
	Selector   string
	IsVariadic bool
	IsGeneric  bool
}

// TypeParam represents a type parameter in a generic function.
type TypeParam struct {
	Name       string // type parameter name (e.g., "B")
	Constraint string // constraint type (e.g., "Boolean", "cmp.Ordered")
}

// Ident represents an exported identifier (type, constant, or variable) from the source package.
type Ident struct {
	ID            string
	Name          string
	SourcePackage string
	TargetPackage string
	DocString     string
	IsAlias       bool
	IsDeprecated  bool
	Function      *Function // for function types (or vars)

	// extraneous information when scanning in collectDoc mode
	Domain        string
	SourceLink    *token.Position
	ExtraComments []ExtraComment
}

// TestValue represents a single parsed test value expression.
//
// It stores both the original string (for debugging/audit) and the parsed AST.
type TestValue struct {
	Raw   string   // Original string from doc comment
	Expr  ast.Expr // Parsed Go expression (nil if parse failed)
	Error error    // Parse error if any
}

// Test captures test values to use with generated tests.
//
// Test values are parsed as Go expressions and stored with their AST representation.
type Test struct {
	TestedValue      string              // DEPRECATED: Original raw string, kept for backward compatibility
	TestedValues     []TestValue         // Parsed test value expressions
	ExpectedOutcome  TestExpectedOutcome // Expected test outcome (success/failure/panic)
	AssertionMessage string              // Optional assertion message for panic tests
}

type TestExpectedOutcome uint8

const (
	TestNone TestExpectedOutcome = iota
	TestSuccess
	TestFailure
	TestPanic
)

type CommentTag uint8

const (
	CommentTagNone CommentTag = iota
	CommentTagDomain
	CommentTagMaintainer
	CommentTagMention
	CommentTagNote
	CommentTagDomainDescription
)

type ExtraComment struct {
	Tag  CommentTag
	Key  string
	Text string
}
