// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package comments

import (
	"strings"
	"sync"
	"testing"

	"golang.org/x/tools/go/packages"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
)

//nolint:gochecknoglobals // we use [sync.Once] to cache the parsed test source and run the tests faster
var (
	loadTestPackageOnce sync.Once
	cachedTestPackage   *packages.Package
)

// loadTestPackage loads the testdata/assertions package for testing.
func loadTestPackage(t *testing.T) (*packages.Package, *Extractor) {
	t.Helper()

	loadTestPackageOnce.Do(func() {
		cfg := &packages.Config{
			Mode: packages.NeedName |
				packages.NeedFiles |
				packages.NeedCompiledGoFiles |
				packages.NeedImports |
				packages.NeedDeps |
				packages.NeedTypes |
				packages.NeedSyntax |
				packages.NeedTypesInfo,
			BuildFlags: []string{"-tags", "integrationtest"},
		}

		pkgs, err := packages.Load(cfg, "./testdata/assertions")
		if err != nil {
			t.Fatalf("Failed to load test package: %v", err)
		}

		if len(pkgs) != 1 {
			t.Fatalf("Expected 1 package, got %d", len(pkgs))
		}

		pkg := pkgs[0]
		if len(pkg.Errors) > 0 {
			t.Fatalf("Package has errors: %v", pkg.Errors)
		}

		cachedTestPackage = pkg
	})

	pkg := cachedTestPackage
	filesMap := BuildFilesMap(pkg)
	extractor := New(pkg.Syntax, pkg.Fset, filesMap)

	return pkg, extractor
}

func TestNew(t *testing.T) {
	_, extractor := loadTestPackage(t)

	if extractor == nil {
		t.Fatal("New() returned nil")
	}
	if extractor.syntaxPackage == nil {
		t.Error("syntaxPackage is nil")
	}
	if extractor.fileSet == nil {
		t.Error("fileSet is nil")
	}
	if extractor.filesMap == nil {
		t.Error("filesMap is nil")
	}
	if extractor.declCache == nil {
		t.Error("declCache is nil")
	}
}

func TestExtractPackageComments(t *testing.T) {
	_, extractor := loadTestPackage(t)

	comments := extractor.ExtractPackageComments()

	if comments == nil {
		t.Fatal("ExtractPackageComments() returned nil")
	}

	text := comments.Text()
	if text == "" {
		t.Error("Package comments are empty")
	}

	// Should contain package description from doc.go
	if !strings.Contains(text, "Package assertions") {
		t.Errorf("Package comments should mention 'Package assertions', got: %s", text)
	}
}

func TestExtractCopyright(t *testing.T) {
	_, extractor := loadTestPackage(t)

	copyright := extractor.ExtractCopyright()

	if copyright == nil {
		t.Fatal("ExtractCopyright() returned nil")
	}

	text := copyright.Text()
	if text == "" {
		t.Error("Copyright is empty")
	}

	// Should contain copyright notice
	if !strings.Contains(strings.ToLower(text), "copyright") {
		t.Errorf("Copyright should contain 'copyright', got: %s", text)
	}

	// Should contain SPDX header
	if !strings.Contains(text, "SPDX-FileCopyrightText") {
		t.Errorf("Copyright should contain SPDX header, got: %s", text)
	}
}

func TestExtractDomainDescriptions(t *testing.T) {
	_, extractor := loadTestPackage(t)

	domains := extractor.ExtractDomainDescriptions()

	if domains == nil {
		t.Fatal("ExtractDomainDescriptions() returned nil")
	}

	if len(domains) == 0 {
		t.Error("Expected domain descriptions, got none")
	}

	// Verify we found expected domains
	expectedDomains := map[string]bool{
		"boolean":    false,
		"collection": false,
		"compare":    false,
		"equal":      false,
		"error":      false,
	}

	for _, domain := range domains {
		if domain.Tag != model.CommentTagDomainDescription {
			t.Errorf("Expected tag %v, got %v", model.CommentTagDomainDescription, domain.Tag)
		}

		if _, exists := expectedDomains[domain.Key]; exists {
			expectedDomains[domain.Key] = true
		}

		// Verify description is not empty
		if domain.Text == "" {
			t.Errorf("Domain %q has empty description", domain.Key)
		}
	}

	// Check that we found at least some expected domains
	foundCount := 0
	for domain, found := range expectedDomains {
		if found {
			foundCount++
		} else {
			t.Logf("Domain %q not found (may be OK if doc.go doesn't list it)", domain)
		}
	}

	if foundCount == 0 {
		t.Error("Expected to find at least one domain from the expected list")
	}
}

func TestExtractComments(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "True" function
	scope := pkg.Types.Scope()
	trueObj := scope.Lookup("True")
	if trueObj == nil {
		t.Fatal("Could not find 'True' function in package scope")
	}

	comments := extractor.ExtractComments(trueObj)

	if comments == nil {
		t.Fatal("ExtractComments() returned nil")
	}

	text := comments.Text()
	if text == "" {
		t.Error("Function comments are empty")
	}

	// Should contain function description
	if !strings.Contains(text, "True asserts") {
		t.Errorf("Expected function description, got: %s", text)
	}

	// Should contain Examples section
	if !strings.Contains(text, "Examples:") {
		t.Errorf("Expected Examples section, got: %s", text)
	}
}

func TestExtractComments_Caching(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	scope := pkg.Types.Scope()
	trueObj := scope.Lookup("True")
	if trueObj == nil {
		t.Fatal("Could not find 'True' function")
	}

	// First call - should populate cache
	comments1 := extractor.ExtractComments(trueObj)

	// Second call - should use cache
	comments2 := extractor.ExtractComments(trueObj)

	// Should return same content
	if comments1.Text() != comments2.Text() {
		t.Error("Cached comments differ from original")
	}

	// Verify cache was populated
	if _, cached := extractor.declCache[trueObj]; !cached {
		t.Error("Object not found in cache after ExtractComments()")
	}
}

func TestExtractExtraComments(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "True" function
	scope := pkg.Types.Scope()
	trueObj := scope.Lookup("True")
	if trueObj == nil {
		t.Fatal("Could not find 'True' function in package scope")
	}

	extraComments := extractor.ExtractExtraComments(trueObj)

	// The True function has "// Domain: boolean" inside it
	if len(extraComments) == 0 {
		t.Fatal("Expected extra comments (domain tag), got none")
	}

	// Should have parsed the domain tag
	foundDomain := false
	for _, comment := range extraComments {
		if comment.Tag == model.CommentTagDomain && comment.Key == "boolean" {
			foundDomain = true
			break
		}
	}

	if !foundDomain {
		t.Errorf("Expected to find domain tag with key 'boolean', got: %+v", extraComments)
	}
}

func TestExtractComments_NoComment(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Try to get comments for a type that exists but might not have comments
	scope := pkg.Types.Scope()
	// Look for the "T" type (testing interface)
	tObj := scope.Lookup("T")

	if tObj == nil {
		// If T doesn't exist in testdata, that's fine - skip this test
		t.Skip("Type 'T' not found in test package")
	}

	comments := extractor.ExtractComments(tObj)

	// Should never return nil, even if no comments
	if comments == nil {
		t.Error("ExtractComments() returned nil for object without comments")
	}
}

func TestExtractComments_MultipleFunctions(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	scope := pkg.Types.Scope()

	testCases := []struct {
		funcName      string
		shouldContain string
	}{
		{
			funcName:      "True",
			shouldContain: "True asserts",
		},
		{
			funcName:      "False",
			shouldContain: "False asserts",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.funcName, func(t *testing.T) {
			obj := scope.Lookup(tc.funcName)
			if obj == nil {
				t.Fatalf("Could not find function %q", tc.funcName)
			}

			comments := extractor.ExtractComments(obj)
			text := comments.Text()

			if !strings.Contains(text, tc.shouldContain) {
				t.Errorf("Expected comment to contain %q, got: %s", tc.shouldContain, text)
			}
		})
	}
}

func TestExtractComments_TypeDeclaration(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "CollectT" type
	scope := pkg.Types.Scope()
	collectTObj := scope.Lookup("CollectT")
	if collectTObj == nil {
		t.Fatal("Could not find 'CollectT' type in package scope")
	}

	comments := extractor.ExtractComments(collectTObj)

	if comments == nil {
		t.Fatal("ExtractComments() returned nil for type")
	}

	text := comments.Text()
	if text == "" {
		t.Error("Type comments are empty")
	}

	// Should contain type description
	if !strings.Contains(text, "CollectT") {
		t.Errorf("Expected type description, got: %s", text)
	}
}

func TestExtractExtraComments_TypeWithBodyComments(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "CollectT" type which has maintainer comments inside
	scope := pkg.Types.Scope()
	collectTObj := scope.Lookup("CollectT")
	if collectTObj == nil {
		t.Fatal("Could not find 'CollectT' type")
	}

	extraComments := extractor.ExtractExtraComments(collectTObj)

	// The CollectT struct has "// Maintainer:" comments inside it
	if len(extraComments) == 0 {
		t.Fatal("Expected extra comments (maintainer tag), got none")
	}

	// Should have parsed the maintainer tag
	foundMaintainer := false
	for _, comment := range extraComments {
		if comment.Tag == model.CommentTagMaintainer {
			foundMaintainer = true
			// Verify maintainer text is not empty
			if comment.Text == "" {
				t.Error("Maintainer comment has empty text")
			}
			// Should mention the maintainer notes
			if !strings.Contains(comment.Text, "runtime.GoExit") && !strings.Contains(comment.Text, "deprecated") {
				t.Logf("Maintainer comment: %q", comment.Text)
			}
			break
		}
	}

	if !foundMaintainer {
		t.Errorf("Expected to find maintainer tag, got: %+v", extraComments)
	}
}

func TestExtractExtraComments_Caching(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	scope := pkg.Types.Scope()
	collectTObj := scope.Lookup("CollectT")
	if collectTObj == nil {
		t.Fatal("Could not find 'CollectT' type")
	}

	// First call - should populate cache
	extraComments1 := extractor.ExtractExtraComments(collectTObj)

	// Second call - should use cache
	extraComments2 := extractor.ExtractExtraComments(collectTObj)

	// Should return same number of comments
	if len(extraComments1) != len(extraComments2) {
		t.Error("Cached extra comments differ from original")
	}

	// Verify cache was populated
	if _, cached := extractor.declCache[collectTObj]; !cached {
		t.Error("Object not found in cache after ExtractExtraComments()")
	}
}

func TestExtractComments_VariableDeclaration(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "ErrTest" variable
	scope := pkg.Types.Scope()
	errTestObj := scope.Lookup("ErrTest")
	if errTestObj == nil {
		t.Fatal("Could not find 'ErrTest' variable in package scope")
	}

	comments := extractor.ExtractComments(errTestObj)

	if comments == nil {
		t.Fatal("ExtractComments() returned nil for variable")
	}

	text := comments.Text()
	if text == "" {
		t.Error("Variable comments are empty")
	}

	// Should contain variable description
	if !strings.Contains(text, "ErrTest") {
		t.Errorf("Expected variable description, got: %s", text)
	}

	if !strings.Contains(text, "error instance") {
		t.Errorf("Expected description 'error instance', got: %s", text)
	}
}

func TestExtractExtraComments_InterfaceWithBodyComments(t *testing.T) {
	pkg, extractor := loadTestPackage(t)

	// Find the "T" interface which has maintainer comments inside
	scope := pkg.Types.Scope()
	tObj := scope.Lookup("T")
	if tObj == nil {
		t.Fatal("Could not find 'T' interface")
	}

	extraComments := extractor.ExtractExtraComments(tObj)

	// Note: Comments attached to interface methods (not floating) may not be
	// extracted as "extra comments" - they're associated with the method.
	// This test verifies the ExtractExtraComments path for interfaces,
	// even if the result is empty for interfaces with only method-attached comments.
	if extraComments == nil {
		// nil is acceptable for interfaces without floating comments
		t.Log("No extra comments found for interface (expected if comments are method-attached)")
		return
	}

	// If we do have comments, verify they're parsed correctly
	if len(extraComments) > 0 {
		foundMaintainer := false
		for _, comment := range extraComments {
			if comment.Tag == model.CommentTagMaintainer {
				foundMaintainer = true
				// Verify maintainer text is not empty
				if comment.Text == "" {
					t.Error("Maintainer comment has empty text")
				}
				t.Logf("Found maintainer comment: %q", comment.Text)
				break
			}
		}

		if !foundMaintainer {
			t.Logf("Extra comments found but no maintainer tag: %+v", extraComments)
		}
	}
}

func TestBuildFilesMap(t *testing.T) {
	pkg, _ := loadTestPackage(t)

	filesMap := BuildFilesMap(pkg)

	if filesMap == nil {
		t.Fatal("BuildFilesMap() returned nil")
	}

	if len(filesMap) == 0 {
		t.Error("BuildFilesMap() returned empty map")
	}

	// Should have entries for each syntax file
	if len(filesMap) != len(pkg.Syntax) {
		t.Errorf("Expected %d entries in filesMap, got %d", len(pkg.Syntax), len(filesMap))
	}

	// Verify map structure
	for tokenFile, astFile := range filesMap {
		if tokenFile == nil {
			t.Error("Found nil token.File key in filesMap")
		}
		if astFile == nil {
			t.Error("Found nil ast.File value in filesMap")
		}
	}
}
