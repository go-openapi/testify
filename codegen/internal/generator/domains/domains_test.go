// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package domains

import (
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
)

const (
	testPackage        = "github.com/go-openapi/testify/v2/internal/assertions"
	testRepo           = "github.com/go-openapi/testify/v2"
	testAssertPackage  = "github.com/go-openapi/testify/v2/assert"
	testRequirePackage = "github.com/go-openapi/testify/v2/require"
)

// TestFlattenDocumentation verifies documentation flattening.
func TestFlattenDocumentation(t *testing.T) {
	t.Parallel()

	for c := range flattenDocumentationCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			flattened := FlattenDocumentation(c.doc)

			if len(flattened) != c.expectedCount {
				t.Errorf("Expected %d packages, got %d", c.expectedCount, len(flattened))
			}

			for _, pkg := range c.expectedPackages {
				if _, ok := flattened[pkg]; !ok {
					t.Errorf("Missing package %s", pkg)
				}
			}
		})
	}
}

// TestMakeDomainIndex verifies domain index creation.
func TestMakeDomainIndex(t *testing.T) {
	t.Parallel()

	for c := range makeDomainIndexCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			index := MakeDomainIndex(c.docs)

			// Verify metadata
			if c.checkMetadata != nil {
				c.checkMetadata(t, index)
			}

			// Verify domains
			if c.checkDomains != nil {
				c.checkDomains(t, index)
			}
		})
	}
}

// TestDomainIndexSorting verifies that domains are sorted correctly.
func TestDomainIndexSorting(t *testing.T) {
	t.Parallel()

	for c := range domainIndexSortingCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			index := MakeDomainIndex(c.docs)

			// Collect domain names in order
			var domains []string
			for domain := range index.Entries() {
				domains = append(domains, domain)
			}

			// Verify order
			if len(domains) != len(c.expectedOrder) {
				t.Fatalf("Expected %d domains, got %d: %v", len(c.expectedOrder), len(domains), domains)
			}

			for i, expected := range c.expectedOrder {
				if domains[i] != expected {
					t.Errorf("Position %d: expected %q, got %q", i, expected, domains[i])
				}
			}
		})
	}
}

// TestEntry_AddMethods verifies entry mutation methods.
func TestEntry_AddMethods(t *testing.T) {
	t.Parallel()

	for c := range entryAddMethodsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			var entry Entry

			// Execute the add operation
			c.addFunc(&entry)

			// Verify the result
			c.checkFunc(t, &entry)
		})
	}
}

// TestEntry_ExtraPackages verifies extra package filtering.
func TestEntry_ExtraPackages(t *testing.T) {
	t.Parallel()

	for c := range entryExtraPackagesCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			var entry Entry

			// Add packages
			for _, pkg := range c.packages {
				entry.AddPackage(pkg.name, pkg.pkg)
			}

			extras := entry.ExtraPackages()

			// Verify count
			if len(extras) != c.expectedCount {
				t.Errorf("Expected %d extra packages, got %d", c.expectedCount, len(extras))
			}

			// Verify order and content
			for i, expectedPkg := range c.expectedPackages {
				if i >= len(extras) {
					break
				}
				if extras[i].Package != expectedPkg {
					t.Errorf("Expected package[%d] to be %s, got %s", i, expectedPkg, extras[i].Package)
				}
			}
		})
	}
}

// TestEntry_Functions verifies function filtering and sorting.
func TestEntry_Functions(t *testing.T) {
	t.Parallel()

	for c := range entryFunctionsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			var entry Entry

			// Add functions
			for _, fn := range c.functions {
				entry.AddFunction(fn.pkg, fn.fn)
			}

			funcs := entry.Functions()

			// Verify count
			if len(funcs) != c.expectedCount {
				t.Errorf("Expected %d functions, got %d", c.expectedCount, len(funcs))
			}

			// Verify order
			for i, expectedName := range c.expectedNames {
				if i >= len(funcs) {
					break
				}
				if funcs[i].Name != expectedName {
					t.Errorf("Expected function[%d] to be %s, got %s", i, expectedName, funcs[i].Name)
				}
			}
		})
	}
}

/* Test case iterators */

type flattenDocumentationCase struct {
	name             string
	doc              model.Documentation
	expectedCount    int
	expectedPackages []string
}

func flattenDocumentationCases() iter.Seq[flattenDocumentationCase] {
	return slices.Values([]flattenDocumentationCase{
		{
			name: "nested documentation structure",
			doc: func() model.Documentation {
				pkg1 := model.New()
				pkg1.Package = "pkg1"

				pkg2 := model.New()
				pkg2.Package = "pkg2"

				pkg3 := model.New()
				pkg3.Package = "pkg3"

				doc := model.NewDocumentation()
				doc.Documents = []model.Document{
					{
						Package: pkg1,
						Documents: []model.Document{
							{Package: pkg2},
						},
					},
					{Package: pkg3},
				}

				return *doc
			}(),
			expectedCount:    3,
			expectedPackages: []string{"pkg1", "pkg2", "pkg3"},
		},
	})
}

type makeDomainIndexCase struct {
	name          string
	docs          map[string]model.Document
	checkMetadata func(*testing.T, Index)
	checkDomains  func(*testing.T, Index)
}

//nolint:gocognit,gocyclo,cyclop // this is temporary accepted extra complexity. Should refactor with externalized asserting functions
func makeDomainIndexCases() iter.Seq[makeDomainIndexCase] {
	return slices.Values([]makeDomainIndexCase{
		{
			name: "complete domain index with metadata and domains",
			docs: func() map[string]model.Document {
				assertPkg := model.New()
				assertPkg.Package = testPackage
				assertPkg.Tool = "testify-codegen"
				assertPkg.Copyright = "Copyright 2025"
				assertPkg.Receiver = "Assertions"
				assertPkg.Imports = model.ImportMap{
					"assertions": testPackage,
				}

				assertPkg.Functions = []model.Function{
					{Name: "Equal", Domain: "equal"},
					{Name: "True", Domain: "boolean"},
					{Name: "Helper"}, // no domain
				}

				assertPkg.Types = []model.Ident{
					{Name: "TestingT", Domain: "testing"},
					{Name: "H"}, // no domain
				}

				assertPkg.Vars = []model.Ident{
					{Name: "SomeVar", Domain: "boolean"},
					{Name: "SomeOtherVar"}, // no domain
				}

				assertPkg.Consts = []model.Ident{
					{Name: "SomeConst", Domain: "equal"},
					{Name: "SomeOtherConst"}, // no domain
				}

				assertPkg.ExtraComments = []model.ExtraComment{
					{
						Tag:  model.CommentTagDomainDescription,
						Key:  "equal",
						Text: "Equality assertions",
					},
					{
						Tag:  model.CommentTagDomainDescription,
						Key:  "boolean",
						Text: "Boolean assertions",
					},
					{
						Tag:  model.CommentTagMaintainer,
						Key:  "author",
						Text: "Some Author", // Should be ignored by domain indexer
					},
				}

				return map[string]model.Document{
					testPackage: {Package: assertPkg},
				}
			}(),
			checkMetadata: func(t *testing.T, index Index) {
				t.Helper()

				if index.Tool() != "testify-codegen" {
					t.Errorf("Expected tool 'testify-codegen', got %q", index.Tool())
				}
				if index.Copyright() != "Copyright 2025" {
					t.Errorf("Expected copyright 'Copyright 2025', got %q", index.Copyright())
				}
				if index.Receiver() != "Assertions" {
					t.Errorf("Expected receiver 'Assertions', got %q", index.Receiver())
				}
				if index.RootPackage() != testRepo {
					t.Errorf("Expected root package '%s', got %q", testRepo, index.RootPackage())
				}
			},
			checkDomains: func(t *testing.T, index Index) {
				t.Helper()

				domainCount := 0
				for domain, entry := range index.Entries() {
					domainCount++

					switch domain {
					case "equal":
						if entry.Description() != "Equality assertions" {
							t.Errorf("Expected description 'Equality assertions', got %q", entry.Description())
						}
						if len(entry.Functions()) != 1 || entry.Functions()[0].Name != "Equal" {
							t.Error("Expected Equal function in equal domain")
						}
						if len(entry.Consts()) != 1 || entry.Consts()[0].Name != "SomeConst" {
							t.Error("Expected SomeConst in equal domain")
						}

					case "boolean":
						if entry.Description() != "Boolean assertions" {
							t.Errorf("Expected description 'Boolean assertions', got %q", entry.Description())
						}
						if len(entry.Functions()) != 1 || entry.Functions()[0].Name != "True" {
							t.Error("Expected True function in boolean domain")
						}
						if len(entry.Vars()) != 1 || entry.Vars()[0].Name != "SomeVar" {
							t.Error("Expected SomeVar in boolean domain")
						}

					case "testing":
						if len(entry.Types()) != 1 || entry.Types()[0].Name != "TestingT" {
							t.Error("Expected TestingT type in testing domain")
						}

					case "common":
						if entry.Description() != "Other uncategorized helpers" {
							t.Errorf("Expected description 'Other uncategorized helpers', got %q", entry.Description())
						}
						if len(entry.Functions()) != 1 || entry.Functions()[0].Name != "Helper" {
							t.Error("Expected Helper function in common domain")
						}
						if len(entry.Types()) != 1 || entry.Types()[0].Name != "H" {
							t.Error("Expected H type in common domain")
						}
						if len(entry.Vars()) != 1 || entry.Vars()[0].Name != "SomeOtherVar" {
							t.Error("Expected SomeOtherVar in common domain")
						}
						if len(entry.Consts()) != 1 || entry.Consts()[0].Name != "SomeOtherConst" {
							t.Error("Expected SomeOtherConst in common domain")
						}

					default:
						t.Errorf("Unexpected domain: %s", domain)
					}
				}

				if domainCount != 4 {
					t.Errorf("Expected 4 domains, got %d", domainCount)
				}
			},
		},
		{
			name: "dangling domain description without declarations",
			docs: func() map[string]model.Document {
				pkg := model.New()
				pkg.Package = testPackage
				pkg.Imports = model.ImportMap{
					"assertions": testPackage,
				}

				pkg.Functions = []model.Function{
					{Name: "Equal", Domain: "equal"},
				}

				pkg.ExtraComments = []model.ExtraComment{
					{
						Tag:  model.CommentTagDomainDescription,
						Key:  "equal",
						Text: "Equality assertions",
					},
					{
						Tag:  model.CommentTagDomainDescription,
						Key:  "phantom",
						Text: "This domain has no associated declarations", // Dangling description
					},
				}

				return map[string]model.Document{
					testPackage: {Package: pkg},
				}
			}(),
			checkDomains: func(t *testing.T, index Index) {
				t.Helper()

				domainCount := 0
				for domain := range index.Entries() {
					domainCount++
					if domain == "phantom" {
						t.Error("Dangling domain 'phantom' should not appear in entries")
					}
				}

				if domainCount != 1 {
					t.Errorf("Expected 1 domain (equal), got %d", domainCount)
				}
			},
		},
	})
}

type domainIndexSortingCase struct {
	name          string
	docs          map[string]model.Document
	expectedOrder []string
}

func domainIndexSortingCases() iter.Seq[domainIndexSortingCase] {
	return slices.Values([]domainIndexSortingCase{
		{
			name: "alphabetical with common last",
			docs: func() map[string]model.Document {
				pkg := model.New()
				pkg.Package = testPackage
				pkg.Imports = model.ImportMap{
					"assertions": testPackage,
				}

				pkg.Functions = []model.Function{
					{Name: "ZFunc", Domain: "zebra"},
					{Name: "AFunc", Domain: "alpha"},
					{Name: "MFunc", Domain: "middle"},
					{Name: "NoFunc"}, // no domain -> common
				}

				return map[string]model.Document{
					testPackage: {Package: pkg},
				}
			}(),
			expectedOrder: []string{"alpha", "middle", "zebra", "common"},
		},
	})
}

type entryAddMethodsCase struct {
	name      string
	addFunc   func(*Entry)
	checkFunc func(*testing.T, *Entry)
}

func entryAddMethodsCases() iter.Seq[entryAddMethodsCase] {
	return slices.Values([]entryAddMethodsCase{
		{
			name: "add package",
			addFunc: func(e *Entry) {
				pkg := model.New()
				pkg.Package = "test"
				e.AddPackage("pkg", pkg)
			},
			checkFunc: func(t *testing.T, e *Entry) {
				t.Helper()
				if len(e.packages) != 1 {
					t.Error("Package not added")
				}
			},
		},
		{
			name: "add function",
			addFunc: func(e *Entry) {
				fn := model.Function{Name: "TestFunc"}
				e.AddFunction("pkg", fn)
			},
			checkFunc: func(t *testing.T, e *Entry) {
				t.Helper()
				if len(e.funcs) != 1 {
					t.Error("Function not added")
				}
			},
		},
		{
			name: "add type",
			addFunc: func(e *Entry) {
				ty := model.Ident{Name: "TestType"}
				e.AddType("pkg", ty)
			},
			checkFunc: func(t *testing.T, e *Entry) {
				t.Helper()
				if len(e.typeDecls) != 1 {
					t.Error("Type not added")
				}
			},
		},
		{
			name: "add variable",
			addFunc: func(e *Entry) {
				vr := model.Ident{Name: "TestVar"}
				e.AddVariable("pkg", vr)
			},
			checkFunc: func(t *testing.T, e *Entry) {
				t.Helper()
				if len(e.varDecls) != 1 {
					t.Error("Variable not added")
				}
			},
		},
		{
			name: "add const",
			addFunc: func(e *Entry) {
				co := model.Ident{Name: "TestConst"}
				e.AddConst("pkg", co)
			},
			checkFunc: func(t *testing.T, e *Entry) {
				t.Helper()
				if len(e.constDecls) != 1 {
					t.Error("Const not added")
				}
			},
		},
	})
}

type packageToAdd struct {
	name string
	pkg  *model.AssertionPackage
}

type entryExtraPackagesCase struct {
	name             string
	packages         []packageToAdd
	expectedCount    int
	expectedPackages []string
}

func entryExtraPackagesCases() iter.Seq[entryExtraPackagesCase] {
	return slices.Values([]entryExtraPackagesCase{
		{
			name: "filter assertions package and sort",
			packages: func() []packageToAdd {
				assertionsPkg := model.New()
				assertionsPkg.Package = testPackage

				assertPkg := model.New()
				assertPkg.Package = testAssertPackage

				requirePkg := model.New()
				requirePkg.Package = testRequirePackage

				return []packageToAdd{
					{name: "assertions", pkg: assertionsPkg},
					{name: "assert", pkg: assertPkg},
					{name: "require", pkg: requirePkg},
				}
			}(),
			expectedCount: 2,
			expectedPackages: []string{
				testAssertPackage,
				testRequirePackage,
			},
		},
	})
}

type functionToAdd struct {
	pkg string
	fn  model.Function
}

type entryFunctionsCase struct {
	name          string
	functions     []functionToAdd
	expectedCount int
	expectedNames []string
}

func entryFunctionsCases() iter.Seq[entryFunctionsCase] {
	return slices.Values([]entryFunctionsCase{
		{
			name: "filter by assertions package and sort alphabetically",
			functions: []functionToAdd{
				{pkg: "assertions", fn: model.Function{Name: "Zebra"}},
				{pkg: "assertions", fn: model.Function{Name: "Alpha"}},
				{pkg: "assert", fn: model.Function{Name: "Other"}}, // filtered out
			},
			expectedCount: 2,
			expectedNames: []string{"Alpha", "Zebra"},
		},
	})
}
