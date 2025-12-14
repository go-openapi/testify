// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"maps"
	"slices"
	"strings"
)

// AssertionPackage describes the internal/assertions package.
type AssertionPackage struct {
	Tool             string
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
	IsGeneric     bool
	IsHelper      bool
	IsConstructor bool
	Tests         []Test
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

func (vars Parameters) String() string {
	var b strings.Builder
	l := len(vars)

	if l == 0 {
		return ""
	}

	if l > 1 || vars[0].Name != "" {
		b.WriteByte('(')
	}

	if vars[0].Name != "" {
		b.WriteString(vars[0].Name)
		b.WriteByte(' ')
	}
	b.WriteString(vars[0].GoType)

	for _, v := range vars[1:] {
		b.WriteByte(',')
		b.WriteByte(' ')
		if v.Name != "" {
			b.WriteString(v.Name)
			b.WriteByte(' ')
		}
		b.WriteString(v.GoType)

	}

	if l > 1 || vars[0].Name != "" {
		b.WriteByte(')')
	}

	return b.String()
}

// Parameter represents a function parameter or return value.
type Parameter struct {
	Name       string
	GoType     string
	Selector   string
	IsVariadic bool
}

// Ident represents an exported identifier (type, constant, or variable) from the source package.
type Ident struct {
	ID            string
	Name          string
	SourcePackage string
	TargetPackage string
	DocString     string
	IsAlias       bool
	Function      *Function // for function types
}

// Test captures simple test values to use with generated tests.
type Test struct {
	TestedValue      string
	ExpectedOutcome  TestExpectedOutcome
	AssertionMessage string
}

type TestExpectedOutcome uint8

const (
	TestSuccess TestExpectedOutcome = iota
	TestFailure
	TestPanic
)
