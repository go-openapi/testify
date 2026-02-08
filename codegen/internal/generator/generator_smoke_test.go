// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// minimalSource builds a minimal but realistic AssertionPackage fixture
// that satisfies all template requirements for code generation.
//
// testDataPath must be an absolute path so that filepath.Rel can compute
// a relative path from the target directory. Use t.TempDir() or os.TempDir().
func minimalSource(testDataPath string) *model.AssertionPackage {
	pkg := model.New()
	pkg.Package = pkgSource
	pkg.Copyright = "Copyright 2025"
	pkg.Tool = "codegen-test"
	pkg.Header = "Test header"
	pkg.Receiver = "Assertions"
	pkg.TestDataPath = testDataPath
	pkg.Imports = model.ImportMap{}

	// A minimal assertion function.
	// Note: Params excludes 't' AND the variadic 'msgAndArgs' (used by format template).
	// AllParams includes everything for the main assertion template.
	pkg.Functions = model.Functions{
		{
			Name:          "Equal",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "Equal asserts that two objects are equal.",
			Domain:        "equality",
			AllParams: model.Parameters{
				{Name: "t", GoType: "T"},
				{Name: "expected", GoType: "any"},
				{Name: "actual", GoType: "any"},
				{Name: "msgAndArgs", GoType: "any", IsVariadic: true},
			},
			Params: model.Parameters{
				{Name: "expected", GoType: "any"},
				{Name: "actual", GoType: "any"},
			},
			Returns: model.Parameters{
				{GoType: "bool"},
			},
			Tests: []model.Test{
				{
					TestedValue:     "123, 123",
					ExpectedOutcome: model.TestSuccess,
					IsFirst:         true,
				},
				{
					TestedValue:     "123, 456",
					ExpectedOutcome: model.TestFailure,
				},
			},
		},
		{
			Name:          "CallerInfo",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "CallerInfo returns caller information.",
			IsHelper:      true,
			AllParams:     model.Parameters{},
			Params:        model.Parameters{},
			Returns: model.Parameters{
				{GoType: "[]string"},
			},
		},
	}

	pkg.Types = []model.Ident{
		{
			Name:          "T",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "T is the testing interface.",
		},
		{
			Name:          "Assertions",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "Assertions provides methods for asserting.",
		},
		{
			Name:          "ComparisonFunc",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "ComparisonFunc is a comparison function.",
			Function: &model.Function{
				AllParams: model.Parameters{
					{Name: "a", GoType: "any"},
					{Name: "b", GoType: "any"},
				},
				Returns: model.Parameters{
					{GoType: "bool"},
				},
			},
		},
	}

	pkg.Consts = []model.Ident{
		{
			Name:          "TestConst",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "TestConst is a test constant.",
		},
	}

	pkg.Vars = []model.Ident{
		{
			Name:          "TestVar",
			SourcePackage: assertions,
			TargetPackage: assertions,
			DocString:     "TestVar is a test variable.",
		},
	}

	return pkg
}

// TestGenerateAssertSmokeTest exercises the full Generate pipeline for the assert package.
func TestGenerateAssertSmokeTest(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)

	gen := New(source)
	err := gen.Generate(
		WithTargetPackage(pkgAssert),
		WithTargetRoot(tmpDir),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
		WithIncludeTests(true),
		WithIncludeHelpers(true),
		WithIncludeDoc(true),
	)
	if err != nil {
		t.Fatalf("Generate(assert) failed: %v", err)
	}

	// Verify expected files exist and are non-empty
	expectedFiles := []string{
		"assert_types.go",
		"assert_assertions.go",
		"assert_format.go",
		"assert_forward.go",
		"assert_helpers.go",
		"assert_assertions_test.go",
		"assert_format_test.go",
		"assert_forward_test.go",
		"assert_helpers_test.go",
	}

	assertDir := filepath.Join(tmpDir, "assert")
	for _, file := range expectedFiles {
		path := filepath.Join(assertDir, file)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Expected file %s not found: %v", file, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("File %s is empty", file)
		}
	}

	// Verify documentation was built
	docs := gen.Documentation()
	if len(docs.Documents) == 0 {
		t.Error("Documentation should have documents")
	}
}

// TestGenerateRequireSmokeTest exercises the full Generate pipeline for the require package.
func TestGenerateRequireSmokeTest(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)

	gen := New(source)
	err := gen.Generate(
		WithTargetPackage(pkgRequire),
		WithTargetRoot(tmpDir),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
		WithIncludeTests(true),
		WithIncludeHelpers(true),
	)
	if err != nil {
		t.Fatalf("Generate(require) failed: %v", err)
	}

	expectedFiles := []string{
		"require_types.go",
		"require_assertions.go",
		"require_format.go",
		"require_forward.go",
		"require_helpers.go",
		"require_assertions_test.go",
		"require_format_test.go",
		"require_forward_test.go",
		"require_helpers_test.go",
	}

	requireDir := filepath.Join(tmpDir, "require")
	for _, file := range expectedFiles {
		path := filepath.Join(requireDir, file)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Expected file %s not found: %v", file, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("File %s is empty", file)
		}
	}
}

// TestGenerateMinimalNoOptionals exercises Generate with all optional generation disabled.
func TestGenerateMinimalNoOptionals(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)

	gen := New(source)
	err := gen.Generate(
		WithTargetPackage(pkgAssert),
		WithTargetRoot(tmpDir),
		WithIncludeFormatFuncs(false),
		WithIncludeForwardFuncs(false),
		WithIncludeTests(false),
		WithIncludeHelpers(false),
		WithIncludeExamples(false),
		WithIncludeDoc(false),
	)
	if err != nil {
		t.Fatalf("Generate(minimal) failed: %v", err)
	}

	assertDir := filepath.Join(tmpDir, "assert")

	// Types and assertions should exist
	for _, file := range []string{"assert_types.go", "assert_assertions.go"} {
		info, err := os.Stat(filepath.Join(assertDir, file))
		if err != nil {
			t.Errorf("Expected file %s not found: %v", file, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("File %s is empty", file)
		}
	}

	// Optional files should NOT exist
	for _, file := range []string{"assert_format.go", "assert_forward.go", "assert_helpers.go"} {
		if _, err := os.Stat(filepath.Join(assertDir, file)); err == nil {
			t.Errorf("File %s should not exist when disabled", file)
		}
	}
}

// TestGenerateDocSmokeTest exercises the DocGenerator pipeline with merged documentation.
func TestGenerateDocSmokeTest(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)

	// Generate assert package with doc enabled
	assertGen := New(source)
	err := assertGen.Generate(
		WithTargetPackage(pkgAssert),
		WithTargetRoot(tmpDir),
		WithIncludeDoc(true),
	)
	if err != nil {
		t.Fatalf("Generate(assert) failed: %v", err)
	}
	assertDocs := assertGen.Documentation()

	// Generate require package with doc enabled
	requireGen := New(source)
	err = requireGen.Generate(
		WithTargetPackage(pkgRequire),
		WithTargetRoot(tmpDir),
		WithIncludeDoc(true),
	)
	if err != nil {
		t.Fatalf("Generate(require) failed: %v", err)
	}
	requireDocs := requireGen.Documentation()

	// Merge documentation
	assertDocs.Merge(requireDocs)

	// Run doc generator
	docDir := filepath.Join(tmpDir, "docs")
	docGen := NewDocGenerator(assertDocs)
	err = docGen.Generate(
		WithTargetDoc(docDir),
		WithTargetRoot(tmpDir),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
	)
	if err != nil {
		t.Fatalf("DocGenerator.Generate() failed: %v", err)
	}

	// Verify index and at least one domain page were created
	indexPath := filepath.Join(tmpDir, docDir, "_index.md")
	info, err := os.Stat(indexPath)
	if err != nil {
		t.Fatalf("Index file not found: %v", err)
	}
	if info.Size() == 0 {
		t.Error("Index file is empty")
	}

	// Verify the domain page for "equality" exists
	equalityPath := filepath.Join(tmpDir, docDir, "equality.md")
	info, err = os.Stat(equalityPath)
	if err != nil {
		t.Fatalf("equality.md not found: %v", err)
	}
	if info.Size() == 0 {
		t.Error("equality.md is empty")
	}
}

// TestDocGeneratorErrors verifies error paths.
func TestDocGeneratorErrors(t *testing.T) {
	t.Parallel()

	t.Run("missing target doc", func(t *testing.T) {
		t.Parallel()

		doc := model.Documentation{}
		gen := NewDocGenerator(doc)
		err := gen.Generate(WithTargetDoc(""))
		if err == nil {
			t.Error("Expected error for empty target doc")
		}
	})
}

// TestGenerateErrors verifies error paths.
func TestGenerateErrors(t *testing.T) {
	t.Parallel()

	t.Run("missing target package", func(t *testing.T) {
		t.Parallel()

		source := model.New()
		source.Package = pkgSource
		gen := New(source)
		err := gen.Generate()
		if err == nil {
			t.Error("Expected error for missing target package")
		}
	})

	t.Run("invalid target package", func(t *testing.T) {
		t.Parallel()

		source := model.New()
		source.Package = pkgSource
		gen := New(source)
		err := gen.Generate(
			WithTargetPackage("invalid"),
			WithTargetRoot(t.TempDir()),
		)
		if err == nil {
			t.Error("Expected error for invalid target package")
		}
	})
}

// TestBuildDocs verifies doc building is skipped when disabled.
func TestBuildDocsDisabled(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)
	gen := New(source)
	err := gen.initContext([]GenerateOption{
		WithTargetPackage(pkgAssert),
		WithTargetRoot(tmpDir),
		WithIncludeDoc(false),
	})
	if err != nil {
		t.Fatal(err)
	}

	_ = gen.loadTemplates()
	_ = gen.transformModel()
	gen.buildDocs()

	if gen.ctx.docs != nil {
		t.Error("buildDocs should not set docs when disabled")
	}
}

// TestBuildDocsEnabled verifies doc building when enabled.
func TestBuildDocsEnabled(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	source := minimalSource(tmpDir)
	gen := New(source)
	err := gen.initContext([]GenerateOption{
		WithTargetPackage(pkgAssert),
		WithTargetRoot(tmpDir),
		WithIncludeDoc(true),
	})
	if err != nil {
		t.Fatal(err)
	}

	_ = gen.loadTemplates()
	_ = gen.transformModel()
	gen.buildDocs()

	if gen.ctx.docs == nil {
		t.Fatal("buildDocs should set docs when enabled")
	}

	if len(gen.ctx.docs.Documents) != 2 {
		t.Errorf("Expected 2 documents (target + source), got %d", len(gen.ctx.docs.Documents))
	}
}

// TestRenderTemplateErrors verifies renderTemplate error handling.
func TestRenderTemplateErrors(t *testing.T) {
	t.Parallel()

	t.Run("missing index entry", func(t *testing.T) {
		t.Parallel()

		err := renderTemplate(
			map[string]string{},
			nil,
			"nonexistent",
			"target.go",
			nil,
			nil,
		)
		if err == nil {
			t.Error("Expected error for missing index entry")
		}
	})

	t.Run("missing template", func(t *testing.T) {
		t.Parallel()

		err := renderTemplate(
			map[string]string{"test": "test_template"},
			nil,
			"test",
			"target.go",
			nil,
			nil,
		)
		if err == nil {
			t.Error("Expected error for missing template")
		}
	})
}

// TestPopulateFunctionExamples verifies example population logic.
func TestPopulateFunctionExamples(t *testing.T) {
	t.Parallel()

	pkg := &model.AssertionPackage{
		Functions: model.Functions{
			{Name: "Equal"},
			{Name: "True"},
		},
	}

	// No examples at all
	populateFunctionExamples(pkg, nil)
	if len(pkg.Functions[0].Examples) != 0 {
		t.Error("No examples should be set")
	}
}

// TestPopulateIdentExamples verifies ident example population.
func TestPopulateIdentExamples(t *testing.T) {
	t.Parallel()

	idents := []model.Ident{
		{Name: "T"},
		{Name: "Assertions"},
	}

	// No examples at all
	populateIdentExamples(idents, nil)
	if len(idents[0].Examples) != 0 {
		t.Error("No examples should be set")
	}
}

// TestTransformFunc verifies function transformation for assert vs require.
func TestTransformFunc(t *testing.T) {
	t.Parallel()

	source := minimalSource(t.TempDir())

	t.Run("assert package uses mockT", func(t *testing.T) {
		t.Parallel()

		gen := New(source)
		err := gen.initContext([]GenerateOption{
			WithTargetPackage(pkgAssert),
			WithTargetRoot(t.TempDir()),
		})
		if err != nil {
			t.Fatal(err)
		}

		fn := model.Function{
			Name: "Equal",
			Params: model.Parameters{
				{Name: "expected", GoType: "any"},
			},
		}
		result := gen.transformFunc(fn)

		if result.UseMock != mock {
			t.Errorf("Assert Equal should use mockT, got %q", result.UseMock)
		}
	})

	t.Run("require package uses mockFailNowT", func(t *testing.T) {
		t.Parallel()

		gen := New(source)
		err := gen.initContext([]GenerateOption{
			WithTargetPackage(pkgRequire),
			WithTargetRoot(t.TempDir()),
		})
		if err != nil {
			t.Fatal(err)
		}

		fn := model.Function{
			Name: "Equal",
			Params: model.Parameters{
				{Name: "expected", GoType: "any"},
			},
		}
		result := gen.transformFunc(fn)

		if result.UseMock != mockWithFailNow {
			t.Errorf("Require Equal should use mockFailNowT, got %q", result.UseMock)
		}
	})

	t.Run("FailNow always uses mockFailNowT", func(t *testing.T) {
		t.Parallel()

		gen := New(source)
		err := gen.initContext([]GenerateOption{
			WithTargetPackage(pkgAssert),
			WithTargetRoot(t.TempDir()),
		})
		if err != nil {
			t.Fatal(err)
		}

		fn := model.Function{Name: "FailNow"}
		result := gen.transformFunc(fn)

		if result.UseMock != mockWithFailNow {
			t.Errorf("FailNow in assert should use mockFailNowT, got %q", result.UseMock)
		}
	})
}

// TestAllOptions exercises all option functions.
func TestAllOptions(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	opts := generateOptionsWithDefaults([]GenerateOption{
		WithTargetRoot(tmpDir),
		WithTargetPackage("assert"),
		WithTargetDoc("docs"),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
		WithIncludeTests(true),
		WithIncludeGenerics(true),
		WithIncludeHelpers(true),
		WithIncludeExamples(true),
		WithRunnableExamples(true),
		WithIncludeDoc(true),
	})

	if opts.targetRoot != tmpDir {
		t.Errorf("targetRoot: want %q, got %q", tmpDir, opts.targetRoot)
	}
	if opts.targetPkg != "assert" {
		t.Errorf("targetPkg: want assert, got %q", opts.targetPkg)
	}
	if opts.targetDoc != "docs" {
		t.Errorf("targetDoc: want docs, got %q", opts.targetDoc)
	}
	if !opts.enableFormat {
		t.Error("enableFormat should be true")
	}
	if !opts.enableForward {
		t.Error("enableForward should be true")
	}
	if !opts.generateTests {
		t.Error("generateTests should be true")
	}
	if !opts.enableGenerics {
		t.Error("enableGenerics should be true")
	}
	if !opts.generateHelpers {
		t.Error("generateHelpers should be true")
	}
	if !opts.generateExamples {
		t.Error("generateExamples should be true")
	}
	if !opts.runnableExamples {
		t.Error("runnableExamples should be true")
	}
	if !opts.generateDoc {
		t.Error("generateDoc should be true")
	}
}
