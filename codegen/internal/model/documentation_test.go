// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"testing"
)

func TestNewDocumentation(t *testing.T) {
	t.Parallel()

	doc := NewDocumentation()

	if doc == nil {
		t.Fatal("NewDocumentation() returned nil")
	}

	if doc.Package != nil {
		t.Error("Package should be nil")
	}

	if doc.Documents != nil {
		t.Error("Documents should be nil")
	}
}

func TestDocumentationMerge_EmptyTarget(t *testing.T) {
	t.Parallel()

	target := NewDocumentation()

	incoming := Documentation{
		Package: &AssertionPackage{Package: pkgAssert},
		Documents: []Document{
			{Title: "doc1"},
		},
	}

	target.Merge(incoming)

	if target.Package == nil || target.Package.Package != pkgAssert {
		t.Error("Merge into empty should replace Package")
	}

	if len(target.Documents) != 1 || target.Documents[0].Title != "doc1" {
		t.Error("Merge into empty should replace Documents")
	}
}

func TestDocumentationMerge_SingleDocument(t *testing.T) {
	t.Parallel()

	target := &Documentation{
		Package: &AssertionPackage{Package: pkgAssert},
		Documents: []Document{
			{Title: "assertDoc"},
		},
	}

	incoming := Documentation{
		Package: &AssertionPackage{Package: pkgRequire},
		Documents: []Document{
			{Title: "requireDoc"},
		},
	}

	target.Merge(incoming)

	// Should create a hierarchy with 2 folders
	if target.Package != nil {
		t.Error("Package should be nil after hierarchy creation")
	}

	if len(target.Documents) != 2 {
		t.Fatalf("Expected 2 documents, got %d", len(target.Documents))
	}

	if target.Documents[0].Title != pkgAssert {
		t.Errorf("First folder title: want %q, got %q", pkgAssert, target.Documents[0].Title)
	}

	if target.Documents[0].Kind != KindFolder {
		t.Error("First document should be KindFolder")
	}

	if target.Documents[1].Title != pkgRequire {
		t.Errorf("Second folder title: want %q, got %q", pkgRequire, target.Documents[1].Title)
	}
}

func TestDocumentationMerge_MultipleDocuments(t *testing.T) {
	t.Parallel()

	// Start with a hierarchy already in place (2 folders)
	target := &Documentation{
		Documents: []Document{
			{Title: pkgAssert, Kind: KindFolder},
			{Title: pkgRequire, Kind: KindFolder},
		},
	}

	incoming := Documentation{
		Package: &AssertionPackage{Package: "extra"},
		Documents: []Document{
			{Title: "extraDoc"},
		},
	}

	target.Merge(incoming)

	if len(target.Documents) != 3 {
		t.Fatalf("Expected 3 documents, got %d", len(target.Documents))
	}

	if target.Documents[2].Title != "extra" {
		t.Errorf("Third folder title: want %q, got %q", "extra", target.Documents[2].Title)
	}

	if target.Documents[2].Kind != KindFolder {
		t.Error("Third document should be KindFolder")
	}
}

func TestDocumentHasGenerics(t *testing.T) {
	t.Parallel()

	// nil package
	doc := Document{}
	if doc.HasGenerics() {
		t.Error("nil package should return false")
	}

	// no generic functions
	doc = Document{
		Package: &AssertionPackage{
			Functions: Functions{
				{Name: "Equal"},
				{Name: "True"},
			},
		},
	}
	if doc.HasGenerics() {
		t.Error("No generic functions should return false")
	}

	// has generic function
	doc = Document{
		Package: &AssertionPackage{
			Functions: Functions{
				{Name: "Equal"},
				{Name: "Greater", IsGeneric: true},
			},
		},
	}
	if !doc.HasGenerics() {
		t.Error("With generic function should return true")
	}
}

func TestExtraPackagesLookupFunction(t *testing.T) {
	t.Parallel()

	assertPkg := &AssertionPackage{
		Package:        pkgAssert,
		Receiver:       "Assertions",
		EnableFormat:   true,
		EnableForward:  true,
		EnableGenerics: false,
		EnableExamples: true,
		Functions: Functions{
			{Name: "Equal"},
			{Name: "True"},
		},
	}

	requirePkg := &AssertionPackage{
		Package:        pkgRequire,
		Receiver:       "Assertions",
		EnableFormat:   true,
		EnableForward:  true,
		EnableGenerics: true,
		EnableExamples: false,
		Functions: Functions{
			{Name: "Equal"},
			{Name: "False"},
		},
	}

	pkgs := ExtraPackages{assertPkg, requirePkg}

	// Found in both packages
	results := pkgs.LookupFunction("Equal")
	if len(results) != 2 {
		t.Fatalf("LookupFunction(Equal): want 2 results, got %d", len(results))
	}
	if results[0].Package != pkgAssert {
		t.Errorf("First result package: want %q, got %q", pkgAssert, results[0].Package)
	}
	if results[1].Package != pkgRequire {
		t.Errorf("Second result package: want %q, got %q", pkgRequire, results[1].Package)
	}
	if !results[0].EnableExamples {
		t.Error("First result should have EnableExamples true")
	}
	if results[1].EnableExamples {
		t.Error("Second result should have EnableExamples false")
	}

	// Found in one package
	results = pkgs.LookupFunction("True")
	if len(results) != 1 {
		t.Fatalf("LookupFunction(True): want 1 result, got %d", len(results))
	}

	// Not found
	results = pkgs.LookupFunction("NonExistent")
	if len(results) != 0 {
		t.Errorf("LookupFunction(NonExistent): want 0 results, got %d", len(results))
	}

	// Empty extra packages
	empty := ExtraPackages{}
	results = empty.LookupFunction("Equal")
	if len(results) != 0 {
		t.Errorf("LookupFunction on empty: want 0 results, got %d", len(results))
	}
}
