// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"iter"
	"slices"
	"testing"
)

const (
	pkgAssertions = "assertions"
	pkgAssert     = "assert"
	pkgRequire    = "require"
	myType        = "[myType]"
)

func TestNew(t *testing.T) {
	t.Parallel()

	pkg := New()

	if pkg == nil {
		t.Fatal("New() returned nil")
	}

	if pkg.Functions == nil {
		t.Error("Functions should be initialized")
	}

	if pkg.Types == nil {
		t.Error("Types should be initialized")
	}

	if pkg.Consts == nil {
		t.Error("Consts should be initialized")
	}

	if pkg.Vars == nil {
		t.Error("Vars should be initialized")
	}

	if cap(pkg.Functions) != 100 {
		t.Errorf("Functions capacity: want 100, got %d", cap(pkg.Functions))
	}
}

func TestWithTestPackage(t *testing.T) {
	t.Parallel()

	pkg := New()

	if pkg.TestPackage {
		t.Error("TestPackage should be false initially")
	}

	result := pkg.WithTestPackage()

	if result != pkg {
		t.Error("WithTestPackage should return the same pointer")
	}

	if !pkg.TestPackage {
		t.Error("TestPackage should be true after WithTestPackage")
	}
}

func TestHasHelpers(t *testing.T) {
	t.Parallel()

	for tt := range hasHelpersCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pkg := New()
			pkg.Functions = tt.functions

			if got := pkg.HasHelpers(); got != tt.expected {
				t.Errorf("HasHelpers() = %v, want %v", got, tt.expected)
			}
		})
	}
}

type hasHelpersCase struct {
	name      string
	functions Functions
	expected  bool
}

func hasHelpersCases() iter.Seq[hasHelpersCase] {
	return slices.Values([]hasHelpersCase{
		{
			name:      "empty functions",
			functions: Functions{},
			expected:  false,
		},
		{
			name: "no helpers or constructors",
			functions: Functions{
				{Name: "Equal"},
				{Name: "True"},
			},
			expected: false,
		},
		{
			name: "has helper",
			functions: Functions{
				{Name: "Equal"},
				{Name: "CallerInfo", IsHelper: true},
			},
			expected: true,
		},
		{
			name: "has constructor",
			functions: Functions{
				{Name: "New", IsConstructor: true},
			},
			expected: true,
		},
	})
}

func TestClone(t *testing.T) {
	t.Parallel()

	original := New()
	original.Package = pkgAssertions
	original.Imports = ImportMap{"foo": "bar"}
	original.Functions = Functions{
		{Name: "Equal"},
		{Name: "True"},
	}
	original.Types = []Ident{{Name: "T"}}
	original.Consts = []Ident{{Name: "Const1"}}
	original.Vars = []Ident{{Name: "Var1"}}

	cloned := original.Clone()

	if cloned == original {
		t.Error("Clone should return a different pointer")
	}

	if cloned.Package != pkgAssertions {
		t.Error("Clone should preserve Package")
	}

	// Verify slice independence
	cloned.Functions = append(cloned.Functions, Function{Name: "False"})
	if len(original.Functions) != 2 {
		t.Error("Modifying cloned Functions should not affect original")
	}

	cloned.Types = append(cloned.Types, Ident{Name: "T2"})
	if len(original.Types) != 1 {
		t.Error("Modifying cloned Types should not affect original")
	}
}

func TestNames(t *testing.T) {
	t.Parallel()

	for tt := range namesCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pkg := New()
			pkg.Functions = tt.functions
			pkg.EnableFormat = tt.enableFormat

			names := make([]string, 0, len(tt.expected))
			for name := range pkg.Names() {
				names = append(names, name)
			}

			if len(names) != len(tt.expected) {
				t.Errorf("Names() yielded %d names, want %d", len(names), len(tt.expected))
				return
			}

			for i, name := range names {
				if name != tt.expected[i] {
					t.Errorf("Names()[%d] = %q, want %q", i, name, tt.expected[i])
				}
			}
		})
	}
}

type namesCase struct {
	name         string
	functions    Functions
	enableFormat bool
	expected     []string
}

func namesCases() iter.Seq[namesCase] {
	return slices.Values([]namesCase{
		{
			name:      "empty",
			functions: Functions{},
			expected:  nil,
		},
		{
			name: "without format",
			functions: Functions{
				{Name: "Equal"},
				{Name: "True"},
			},
			enableFormat: false,
			expected:     []string{"Equal", "True"},
		},
		{
			name: "with format",
			functions: Functions{
				{Name: "Equal"},
				{Name: "True"},
			},
			enableFormat: true,
			expected:     []string{"Equal", "Equalf", "True", "Truef"},
		},
	})
}

func TestNamesEarlyBreak(t *testing.T) {
	t.Parallel()

	pkg := New()
	pkg.Functions = Functions{
		{Name: "Equal"},
		{Name: "True"},
		{Name: "False"},
	}
	pkg.EnableFormat = true

	var names []string
	for name := range pkg.Names() {
		names = append(names, name)
		if len(names) == 2 {
			break
		}
	}

	if len(names) != 2 {
		t.Errorf("Early break: got %d names, want 2", len(names))
	}
}

func TestImportMapHasImports(t *testing.T) {
	t.Parallel()

	empty := ImportMap{}
	if empty.HasImports() {
		t.Error("Empty map should return false")
	}

	nonEmpty := ImportMap{"foo": "bar"}
	if !nonEmpty.HasImports() {
		t.Error("Non-empty map should return true")
	}
}

func TestFunctionsScope(t *testing.T) {
	t.Parallel()

	for tt := range scopeCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			seq, err := tt.functions.Scope(tt.scope, tt.ctx)

			if tt.expectError {
				if err == nil {
					t.Fatal("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			var names []string
			for fn := range seq {
				names = append(names, fn.Name)
			}

			if len(names) != len(tt.expected) {
				t.Errorf("Got %d functions %v, want %d %v", len(names), names, len(tt.expected), tt.expected)
				return
			}

			for i, name := range names {
				if name != tt.expected[i] {
					t.Errorf("Function[%d] = %q, want %q", i, name, tt.expected[i])
				}
			}
		})
	}
}

type scopeCase struct {
	name        string
	functions   Functions
	scope       ScopeKind
	ctx         *AssertionPackage
	expectError bool
	expected    []string
}

func scopeCases() iter.Seq[scopeCase] {
	funcs := Functions{
		{Name: "Equal"},
		{Name: "GreaterOrEqual", IsGeneric: true},
		{Name: "CallerInfo", IsHelper: true},
		{Name: "New", IsConstructor: true},
		{Name: "True"},
	}

	ctxNoGenerics := &AssertionPackage{EnableGenerics: false}
	ctxWithGenerics := &AssertionPackage{EnableGenerics: true}

	return slices.Values([]scopeCase{
		{
			name:      "with-generics scope, generics disabled",
			functions: funcs,
			scope:     ScopeKindWithGenerics,
			ctx:       ctxNoGenerics,
			expected:  []string{"Equal", "True"},
		},
		{
			name:      "with-generics scope, generics enabled",
			functions: funcs,
			scope:     ScopeKindWithGenerics,
			ctx:       ctxWithGenerics,
			expected:  []string{"Equal", "GreaterOrEqual", "True"},
		},
		{
			name:      "without-generics scope",
			functions: funcs,
			scope:     ScopeKindWithoutGenerics,
			expected:  []string{"Equal", "True"},
		},
		{
			name:      "helpers scope",
			functions: funcs,
			scope:     ScopeKindHelpers,
			expected:  []string{"CallerInfo"},
		},
		{
			name:        "with-generics scope, nil context",
			functions:   funcs,
			scope:       ScopeKindWithGenerics,
			ctx:         nil,
			expectError: true,
		},
		{
			name:        "invalid scope kind",
			functions:   funcs,
			scope:       ScopeKind("invalid"),
			expectError: true,
		},
	})
}

func TestScopeEarlyBreak(t *testing.T) {
	t.Parallel()

	funcs := Functions{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
	}

	// Test early break on iterScopeWithoutGenerics
	seq, err := funcs.Scope(ScopeKindWithoutGenerics, nil)
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	for range seq {
		count++
		if count == 1 {
			break
		}
	}

	if count != 1 {
		t.Errorf("Early break: got %d, want 1", count)
	}

	// Test early break on iterScopeHelpers
	helperFuncs := Functions{
		{Name: "H1", IsHelper: true},
		{Name: "H2", IsHelper: true},
	}

	seq2, err := helperFuncs.Scope(ScopeKindHelpers, nil)
	if err != nil {
		t.Fatal(err)
	}

	count = 0
	for range seq2 {
		count++
		if count == 1 {
			break
		}
	}

	if count != 1 {
		t.Errorf("Early break helpers: got %d, want 1", count)
	}

	// Test early break on iterScopeWithGenerics
	ctx := &AssertionPackage{EnableGenerics: true}
	seq3, err := funcs.Scope(ScopeKindWithGenerics, ctx)
	if err != nil {
		t.Fatal(err)
	}

	count = 0
	for range seq3 {
		count++
		if count == 1 {
			break
		}
	}

	if count != 1 {
		t.Errorf("Early break with-generics: got %d, want 1", count)
	}
}

func TestGenericName(t *testing.T) {
	t.Parallel()

	for tt := range genericNameCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.fn.GenericName(tt.suffixes...); got != tt.expected {
				t.Errorf("GenericName(%v) = %q, want %q", tt.suffixes, got, tt.expected)
			}
		})
	}
}

type genericNameCase struct {
	name     string
	fn       Function
	suffixes []string
	expected string
}

func genericNameCases() iter.Seq[genericNameCase] {
	return slices.Values([]genericNameCase{
		{
			name:     "non-generic, no suffix",
			fn:       Function{Name: "Equal"},
			expected: "Equal",
		},
		{
			name:     "non-generic, with suffix",
			fn:       Function{Name: "Equal"},
			suffixes: []string{"f"},
			expected: "Equalf",
		},
		{
			name: "generic, single type param",
			fn: Function{
				Name:      "Greater",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "T", Constraint: "cmp.Ordered"},
				},
			},
			expected: "Greater[T cmp.Ordered]",
		},
		{
			name: "generic, with suffix",
			fn: Function{
				Name:      "Greater",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "T", Constraint: "cmp.Ordered"},
				},
			},
			suffixes: []string{"f"},
			expected: "Greaterf[T cmp.Ordered]",
		},
		{
			name: "generic, two type params, same constraint elided",
			fn: Function{
				Name:      "Compare",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "A", Constraint: "cmp.Ordered"},
					{Name: "B", Constraint: "cmp.Ordered"},
				},
			},
			expected: "Compare[A, B cmp.Ordered]",
		},
		{
			name: "generic, two type params, different constraints",
			fn: Function{
				Name:      "Compare",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "A", Constraint: "cmp.Ordered"},
					{Name: "B", Constraint: "any"},
				},
			},
			expected: "Compare[A cmp.Ordered, B any]",
		},
		{
			name: "generic, three type params, mixed constraints",
			fn: Function{
				Name:      "Multi",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "A", Constraint: "cmp.Ordered"},
					{Name: "B", Constraint: "cmp.Ordered"},
					{Name: "C", Constraint: "any"},
				},
			},
			expected: "Multi[A, B cmp.Ordered, C any]",
		},
	})
}

func TestGenericCallName(t *testing.T) {
	t.Parallel()

	for tt := range genericCallNameCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.fn.GenericCallName(tt.suffixes...); got != tt.expected {
				t.Errorf("GenericCallName(%v) = %q, want %q", tt.suffixes, got, tt.expected)
			}
		})
	}
}

type genericCallNameCase struct {
	name     string
	fn       Function
	suffixes []string
	expected string
}

func genericCallNameCases() iter.Seq[genericCallNameCase] {
	return slices.Values([]genericCallNameCase{
		{
			name:     "non-generic",
			fn:       Function{Name: "Equal"},
			expected: "Equal",
		},
		{
			name:     "non-generic, with suffix",
			fn:       Function{Name: "Equal"},
			suffixes: []string{"f"},
			expected: "Equalf",
		},
		{
			name: "generic, single type param",
			fn: Function{
				Name:      "Greater",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "T", Constraint: "cmp.Ordered"},
				},
			},
			expected: "Greater[T]",
		},
		{
			name: "generic, two type params",
			fn: Function{
				Name:      "Compare",
				IsGeneric: true,
				TypeParams: []TypeParam{
					{Name: "A", Constraint: "cmp.Ordered"},
					{Name: "B", Constraint: "any"},
				},
			},
			expected: "Compare[A, B]",
		},
	})
}

func TestHasTest(t *testing.T) {
	t.Parallel()

	fn := Function{Name: "Equal"}
	if fn.HasTest() {
		t.Error("HasTest() should be false with no tests")
	}

	fn.Tests = []Test{{ExpectedOutcome: TestSuccess}}
	if !fn.HasTest() {
		t.Error("HasTest() should be true with tests")
	}
}

func TestHasSuccessTest(t *testing.T) {
	t.Parallel()

	fn := Function{
		Name:  "Equal",
		Tests: []Test{{ExpectedOutcome: TestFailure}},
	}
	if fn.HasSuccessTest() {
		t.Error("HasSuccessTest() should be false with only failure tests")
	}

	fn.Tests = append(fn.Tests, Test{ExpectedOutcome: TestSuccess})
	if !fn.HasSuccessTest() {
		t.Error("HasSuccessTest() should be true with a success test")
	}
}

func TestGenericSuffix(t *testing.T) {
	t.Parallel()

	fn := Function{Name: "Equal"}
	if got := fn.GenericSuffix(); got != "" {
		t.Errorf("GenericSuffix() = %q, want empty", got)
	}

	fn.Name = "IsOfTypeT"
	if got := fn.GenericSuffix(); got != myType {
		t.Errorf("GenericSuffix() = %q, want %q", got, myType)
	}

	fn.Name = "OfTypeT"
	if got := fn.GenericSuffix(); got != myType {
		t.Errorf("GenericSuffix() = %q, want %q", got, myType)
	}
}

func TestFailMsg(t *testing.T) {
	t.Parallel()

	for tt := range failMsgCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.fn.FailMsg(tt.args...); got != tt.expected {
				t.Errorf("FailMsg(%v) = %q, want %q", tt.args, got, tt.expected)
			}
		})
	}
}

type failMsgCase struct {
	name     string
	fn       Function
	args     []string
	expected string
}

func failMsgCases() iter.Seq[failMsgCase] {
	return slices.Values([]failMsgCase{
		{
			name:     "mockT, no args",
			fn:       Function{Name: "Equal", UseMock: "mockT"},
			expected: "Equal should mark test as failed",
		},
		{
			name:     "mockFailNowT, no args",
			fn:       Function{Name: "Equal", UseMock: "mockFailNowT"},
			expected: "Equal should call FailNow()",
		},
		{
			name:     "mockT, with suffix",
			fn:       Function{Name: "Equal", UseMock: "mockT"},
			args:     []string{"f"},
			expected: "Equalf should mark test as failed",
		},
		{
			name:     "mockT, with prefix and suffix",
			fn:       Function{Name: "Equal", UseMock: "mockT"},
			args:     []string{"Assertions", "f"},
			expected: "Assertions.Equalf should mark test as failed",
		},
		{
			name:     "mockFailNowT, with prefix and suffix",
			fn:       Function{Name: "Equal", UseMock: "mockFailNowT"},
			args:     []string{"Assertions", ""},
			expected: "Assertions.Equal should call FailNow()",
		},
	})
}

func TestTestPredicates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		test      Test
		success   bool
		failure   bool
		panics    bool
		isRequire bool
	}{
		{
			name:    "success test",
			test:    Test{ExpectedOutcome: TestSuccess},
			success: true,
		},
		{
			name:    "failure test",
			test:    Test{ExpectedOutcome: TestFailure},
			failure: true,
		},
		{
			name:   "panic test",
			test:   Test{ExpectedOutcome: TestPanic},
			panics: true,
		},
		{
			name:      "require test",
			test:      Test{Pkg: pkgRequire},
			isRequire: true,
		},
		{
			name: "none outcome",
			test: Test{ExpectedOutcome: TestNone},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.test.IsSuccess(); got != tt.success {
				t.Errorf("IsSuccess() = %v, want %v", got, tt.success)
			}
			if got := tt.test.IsFailure(); got != tt.failure {
				t.Errorf("IsFailure() = %v, want %v", got, tt.failure)
			}
			if got := tt.test.IsPanic(); got != tt.panics {
				t.Errorf("IsPanic() = %v, want %v", got, tt.panics)
			}
			if got := tt.test.IsKindRequire(); got != tt.isRequire {
				t.Errorf("IsKindRequire() = %v, want %v", got, tt.isRequire)
			}
		})
	}
}

func TestCommentTagString(t *testing.T) {
	t.Parallel()

	for tt := range commentTagCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.tag.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

type commentTagCase struct {
	name     string
	tag      CommentTag
	expected string
}

func commentTagCases() iter.Seq[commentTagCase] {
	return slices.Values([]commentTagCase{
		{name: "none", tag: CommentTagNone, expected: "comment-tag-none"},
		{name: "domain", tag: CommentTagDomain, expected: "comment-tag-domain"},
		{name: "maintainer", tag: CommentTagMaintainer, expected: "comment-tag-maintainer"},
		{name: "mention", tag: CommentTagMention, expected: "comment-tag-mention"},
		{name: "note", tag: CommentTagNote, expected: "comment-tag-note"},
		{name: "domain-description", tag: CommentTagDomainDescription, expected: "comment-tag-domain-description"},
		{name: "invalid", tag: CommentTag(99), expected: "invalid-value"},
	})
}

func TestExtraCommentPredicates(t *testing.T) {
	t.Parallel()

	maintainer := ExtraComment{Tag: CommentTagMaintainer}
	if !maintainer.IsTagMaintainer() {
		t.Error("Expected IsTagMaintainer() true")
	}
	if maintainer.IsTagMention() || maintainer.IsTagNote() {
		t.Error("Other predicates should be false")
	}

	mention := ExtraComment{Tag: CommentTagMention}
	if !mention.IsTagMention() {
		t.Error("Expected IsTagMention() true")
	}

	note := ExtraComment{Tag: CommentTagNote}
	if !note.IsTagNote() {
		t.Error("Expected IsTagNote() true")
	}
}
