// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"iter"
	"maps"
	"slices"
	"strings"
)

type Renderable interface {
	Render() string
}

// AssertionPackage describes the internal/assertions package.
type AssertionPackage struct {
	Tool             string
	Header           string
	Package          string
	DocString        string
	Copyright        string
	Receiver         string
	Imports          ImportMap
	EnableFormat     bool
	EnableForward    bool
	EnableGenerics   bool
	EnableExamples   bool
	RunnableExamples bool

	Functions Functions
	Types     []Ident
	Consts    []Ident
	Vars      []Ident

	// extraneous information when scanning in collectDoc mode
	ExtraComments []ExtraComment
	Context       *Document

	// Overridable context for test generation
	TestDataPath string
	TestPackage  bool
}

func (a *AssertionPackage) WithTestPackage() *AssertionPackage {
	a.TestPackage = true
	return a
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
		Functions: make(Functions, 0, allocatedFuncs),
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

func (a *AssertionPackage) Names() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, fn := range a.Functions {
			if !yield(fn.Name) {
				return
			}

			if !a.EnableFormat {
				continue
			}

			if !yield(fn.Name + "f") {
				return
			}
		}
	}
}

// ImportMap represents the imports for the analyzed package.
type ImportMap map[string]string

func (m ImportMap) HasImports() bool {
	return len(m) > 0
}

type ScopeKind string

const (
	ScopeKindWithGenerics    ScopeKind = "include-generics"
	ScopeKindWithoutGenerics ScopeKind = "exclude-generics"
	ScopeKindHelpers         ScopeKind = "only-helpers"
)

type Functions []Function

func (f Functions) Scope(scope ScopeKind, ctx *AssertionPackage) (iter.Seq[Function], error) {
	switch scope {
	case ScopeKindWithGenerics:
		if ctx == nil {
			return nil, errors.New(`nil package context passed to scope: please pass "." as context to the Scope function`)
		}
		return f.iterScopeWithGenerics(ctx), nil
	case ScopeKindWithoutGenerics:
		return f.iterScopeWithoutGenerics, nil
	case ScopeKindHelpers:
		return f.iterScopeHelpers, nil
	default:
		return nil, fmt.Errorf("invalid scope kind in template. Expected one of include-generics, only-helpers, but got %q", scope)
	}
}

func (f Functions) iterScopeWithGenerics(ctx *AssertionPackage) func(func(Function) bool) {
	return func(yield func(Function) bool) {
		for _, fn := range f {
			if fn.IsConstructor || fn.IsHelper || (fn.IsGeneric && !ctx.EnableGenerics) {
				continue
			}

			if !yield(fn) {
				return
			}
		}
	}
}

func (f Functions) iterScopeWithoutGenerics(yield func(Function) bool) {
	for _, fn := range f {
		if fn.IsConstructor || fn.IsHelper || fn.IsGeneric {
			continue
		}

		if !yield(fn) {
			return
		}
	}
}

func (f Functions) iterScopeHelpers(yield func(Function) bool) {
	for _, fn := range f {
		if !fn.IsHelper {
			continue
		}

		if !yield(fn) {
			return
		}
	}
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
	Examples      []Renderable // testable examples as a collection of [Renderable] examples
	Context       *Document
	TestData
}

// TestData holds test variabilization parameters for templates
type TestData struct {
	TestCall         string
	TestMock         string
	TestErrorPrefix  string
	TestMockFailure  string
	TestPanicWrapper string
	TestMsg          string
}

// GenericName renders the function name with one or more suffixes,
// accounting for any type parameter for generic functions.
func (f Function) GenericName(suffixes ...string) string {
	suffix := strings.Join(suffixes, "")
	if !f.IsGeneric { // means len(f.TypeParams) == 0
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
		if len(f.TypeParams) <= i+1+1 || f.TypeParams[i+1+1].Constraint != p.Constraint {
			w.WriteByte(' ')
			w.WriteString(p.Constraint)
		}
	}
	w.WriteByte(']')

	return w.String()
}

// GenericCallName renders the function name with explicit type parameters.
// This is used when forwarding type parameters, as all type parameters may not be always inferred from the arguments.
func (f Function) GenericCallName(suffixes ...string) string {
	suffix := strings.Join(suffixes, "")
	if !f.IsGeneric { // means len(f.TypeParams) == 0
		return f.Name + suffix
	}

	var w strings.Builder
	w.WriteString(f.Name)
	w.WriteString(suffix)
	w.WriteByte('[')
	c := f.TypeParams[0]
	w.WriteString(c.Name)

	for _, p := range f.TypeParams[1:] {
		w.WriteString(", ")
		w.WriteString(p.Name)
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

// GenericSuffix provides a type parameter instantiation for methods which cannot infer
// type parameters from their arguments. At this moment, only one such case exists: OfTypeT().
func (f Function) GenericSuffix() string {
	if strings.HasSuffix(f.Name, "OfTypeT") {
		return "[myType]"
	}

	return ""
}

// FailMsg returns an error message to report in tests.
//
// If the function should use the "mockFailNowT" mock, we should report that [testing.FailNow] should be
// called, not just marking the test as failed.
//
// An optional suffix may be used to apply to the function name and express a formatted variant.
//
// If to arguments are passed, the first one is interpreted as a prefix, followed by "." and the second as a suffix.
// Further passed args are ignored.
func (f Function) FailMsg(args ...string) string {
	var prefix, suffix string

	switch len(args) {
	case 1:
		suffix = args[0]
	case 2:
		prefix = args[0] + "."
		suffix = args[1]
	}

	if f.UseMock == "mockFailNowT" {
		return prefix + f.Name + suffix + " should call FailNow()"
	}

	return prefix + f.Name + suffix + " should mark test as failed"
}

func (f Function) WithTestCall(in string) Function {
	f.TestCall = in
	return f
}

func (f Function) WithTestMock(in string) Function {
	f.TestMock = in
	return f
}

func (f Function) WithTestMockFailure(in string) Function {
	f.TestMockFailure = in
	return f
}

func (f Function) WithTestErrorPrefix(in string) Function {
	f.TestErrorPrefix = in
	return f
}

func (f Function) WithTestPanicWrapper(in string) Function {
	f.TestPanicWrapper = in
	return f
}

func (f Function) WithTestMsg(in string) Function {
	f.TestMsg = in
	return f
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
	Examples      []Renderable // testable examples as a collection of [Renderable] examples
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
	TestedValues     []TestValue         // Parsed test value expressions
	TestedValue      string              // Original raw string, kept for auditability
	ExpectedOutcome  TestExpectedOutcome // Expected test outcome (success/failure/panic)
	AssertionMessage string              // Optional assertion message for panic tests
	IsFirst          bool
	Pkg              string
}

func (t Test) IsSuccess() bool {
	return t.ExpectedOutcome == TestSuccess
}

func (t Test) IsFailure() bool {
	return t.ExpectedOutcome == TestFailure
}

func (t Test) IsPanic() bool {
	return t.ExpectedOutcome == TestPanic
}

func (t Test) IsKindRequire() bool {
	return t.Pkg == "require"
}

type TestExpectedOutcome uint8

const (
	TestNone TestExpectedOutcome = iota
	TestSuccess
	TestFailure
	TestPanic
)

type CommentTag uint8

// String representation of a comment tag, mostly useful for debugging purpose.
func (t CommentTag) String() string {
	switch t {
	case CommentTagNone:
		return "comment-tag-none"
	case CommentTagDomain:
		return "comment-tag-domain"
	case CommentTagMaintainer:
		return "comment-tag-maintainer"
	case CommentTagMention:
		return "comment-tag-mention"
	case CommentTagNote:
		return "comment-tag-note"
	case CommentTagDomainDescription:
		return "comment-tag-domain-description"
	default:
		return "invalid-value"
	}
}

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

func (c ExtraComment) IsTagMaintainer() bool {
	return c.Tag == CommentTagMaintainer
}

func (c ExtraComment) IsTagMention() bool {
	return c.Tag == CommentTagMention
}

func (c ExtraComment) IsTagNote() bool {
	return c.Tag == CommentTagNote
}
