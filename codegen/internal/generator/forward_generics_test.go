// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// genericSource extends the minimal fixture with a generic assertion, so the forward
// generation has something to place in the go1.27 variant file.
func genericSource(testDataPath string) *model.AssertionPackage {
	pkg := minimalSource(testDataPath)
	pkg.Functions = append(pkg.Functions, model.Function{
		Name:          "EqualT",
		SourcePackage: assertions,
		TargetPackage: assertions,
		DocString:     "EqualT asserts that two values of the same type are equal.",
		Domain:        "equality",
		IsGeneric:     true,
		TypeParams:    []model.TypeParam{{Name: "V", Constraint: "comparable"}},
		AllParams: model.Parameters{
			{Name: "t", GoType: "T"},
			{Name: "expected", GoType: "V"},
			{Name: "actual", GoType: "V"},
			{Name: "msgAndArgs", GoType: "any", IsVariadic: true},
		},
		Params: model.Parameters{
			{Name: "expected", GoType: "V"},
			{Name: "actual", GoType: "V"},
		},
		Returns: model.Parameters{
			{GoType: "bool"},
		},
	})

	return pkg
}

func generateForward(t *testing.T, targetPkg string, withGenerics bool) string {
	t.Helper()

	tmpDir := t.TempDir()
	gen := New(genericSource(tmpDir))
	err := gen.Generate(
		WithTargetPackage(targetPkg),
		WithTargetRoot(tmpDir),
		WithIncludeFormatFuncs(true),
		WithIncludeForwardFuncs(true),
		WithIncludeGenerics(withGenerics),
		WithIncludeTests(true),
		WithIncludeHelpers(true),
	)
	if err != nil {
		t.Fatalf("Generate(%s) failed: %v", targetPkg, err)
	}

	return filepath.Join(tmpDir, targetPkg)
}

func readGenerated(t *testing.T, dir, name string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}

	return string(data)
}

// TestForwardGenericsAssert checks that a generic assertion becomes a method in the
// go1.27 variant file, and stays out of the unguarded forward file.
func TestForwardGenericsAssert(t *testing.T) {
	t.Parallel()

	dir := generateForward(t, pkgAssert, true)

	guarded := readGenerated(t, dir, "assert_forward_go127.go")
	for _, want := range []string{
		"//go:build go1.27",
		"func (a *Assertions) EqualT[V comparable](expected V, actual V, msgAndArgs ...any) bool {",
		"return assertions.EqualT[V](a.T, expected, actual, msgAndArgs...)",
		"func (a *Assertions) EqualTf[V comparable](expected V, actual V, msg string, args ...any) bool {",
	} {
		if !strings.Contains(guarded, want) {
			t.Errorf("assert_forward_go127.go does not contain %q", want)
		}
	}

	// The receiver type and its constructor are declared once, in the unguarded file.
	if strings.Contains(guarded, "type Assertions struct") {
		t.Error("assert_forward_go127.go redeclares the Assertions type")
	}

	plain := readGenerated(t, dir, "assert_forward.go")
	if strings.Contains(plain, "EqualT") {
		t.Error("assert_forward.go carries the generic method, which does not compile before go1.27")
	}
	if !strings.Contains(plain, "type Assertions struct") {
		t.Error("assert_forward.go should declare the Assertions type")
	}

	tests := readGenerated(t, dir, "assert_forward_go127_test.go")
	if !strings.Contains(tests, "//go:build go1.27") {
		t.Error("assert_forward_go127_test.go misses the //go:build guard")
	}
	if !strings.Contains(tests, "func TestAssertionsEqualT(t *testing.T) {") {
		t.Error("assert_forward_go127_test.go misses the generic method test")
	}
}

// TestForwardGenericsRequire checks the require variant, which fails the test rather
// than returning a value.
func TestForwardGenericsRequire(t *testing.T) {
	t.Parallel()

	dir := generateForward(t, pkgRequire, true)

	guarded := readGenerated(t, dir, "require_forward_go127.go")
	for _, want := range []string{
		"//go:build go1.27",
		"func (a *Assertions) EqualT[V comparable](expected V, actual V, msgAndArgs ...any) {",
		"if assertions.EqualT[V](a.T, expected, actual, msgAndArgs...) {",
		"a.T.FailNow()",
	} {
		if !strings.Contains(guarded, want) {
			t.Errorf("require_forward_go127.go does not contain %q", want)
		}
	}
}

// TestForwardGenericsDisabled checks that -include-generics=false leaves no guarded
// forward file behind at all.
func TestForwardGenericsDisabled(t *testing.T) {
	t.Parallel()

	dir := generateForward(t, pkgAssert, false)

	for _, name := range []string{"assert_forward_go127.go", "assert_forward_go127_test.go"} {
		if _, err := os.Stat(filepath.Join(dir, name)); !os.IsNotExist(err) {
			t.Errorf("%s should not be generated when generics are disabled (err: %v)", name, err)
		}
	}
}
