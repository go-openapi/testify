// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// guardedSource returns the minimal fixture plus a go1.26-guarded assertion.
func guardedSource(testDataPath string) *model.AssertionPackage {
	pkg := minimalSource(testDataPath)
	pkg.Functions = append(pkg.Functions, model.Function{
		Name:          "GuardedAssert",
		SourcePackage: assertions,
		TargetPackage: assertions,
		DocString:     "GuardedAssert is a go1.26-guarded assertion.",
		GoBuild:       "go1.26",
		AllParams: model.Parameters{
			{Name: "t", GoType: "T"},
			{Name: "value", GoType: "bool"},
			{Name: "msgAndArgs", GoType: "any", IsVariadic: true},
		},
		Params: model.Parameters{
			{Name: "value", GoType: "bool"},
		},
		Returns: model.Parameters{{GoType: "bool"}},
		Tests: []model.Test{
			{TestedValue: "true", ExpectedOutcome: model.TestSuccess, IsFirst: true},
			{TestedValue: "false", ExpectedOutcome: model.TestFailure},
		},
	})

	return pkg
}

func generateAssert(t *testing.T, source *model.AssertionPackage, root string) {
	t.Helper()

	err := New(source).Generate(
		WithTargetPackage(pkgAssert),
		WithTargetRoot(root),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
		WithIncludeTests(true),
		WithIncludeHelpers(true),
		WithIncludeExamples(true),
	)
	if err != nil {
		t.Fatalf("Generate(assert) failed: %v", err)
	}
}

// TestOrphanVariantCleanup verifies that build-variant files for a variant that no longer
// exists are removed on the next run, while hand-authored lookalikes are preserved.
func TestOrphanVariantCleanup(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()
	assertDir := filepath.Join(tmpDir, "assert")

	// 1. Generate WITH the guarded assertion: go126 variant files appear.
	generateAssert(t, guardedSource(tmpDir), tmpDir)

	guardedFile := filepath.Join(assertDir, "assert_assertions_go126.go")
	if _, err := os.Stat(guardedFile); err != nil {
		t.Fatalf("expected guarded file to be generated first: %v", err)
	}

	// 2. Drop a hand-authored file matching the variant pattern but WITHOUT our marker.
	//    The sweep must never touch it.
	decoy := filepath.Join(assertDir, "assert_handwritten_go126.go")
	const decoyBody = "package assert\n\n// hand-authored, not generated\nfunc handwritten() {}\n"
	if err := os.WriteFile(decoy, []byte(decoyBody), 0o600); err != nil {
		t.Fatal(err)
	}

	// 3. Regenerate WITHOUT the guarded assertion: the variant disappears entirely.
	generateAssert(t, minimalSource(tmpDir), tmpDir)

	// All generated go126 files must be gone.
	leftovers, _ := filepath.Glob(filepath.Join(assertDir, "*_go126*.go"))
	for _, f := range leftovers {
		if f != decoy {
			t.Errorf("orphaned generated variant file not removed: %s", filepath.Base(f))
		}
	}

	// The hand-authored decoy must survive (guard against undue deletion).
	if _, err := os.Stat(decoy); err != nil {
		t.Errorf("hand-authored lookalike was unduly removed: %v", err)
	}

	// And the default (unsuffixed) files must still be present.
	if _, err := os.Stat(filepath.Join(assertDir, "assert_assertions.go")); err != nil {
		t.Errorf("default assertions file should remain: %v", err)
	}
}
