// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

const (
	pkgRequire      = "require"
	pkgAssert       = "assert"
	assertions      = "assertions"
	mockWithFailNow = "mockFailNowT"
	mock            = "mockT"
)

const (
	dirPermissions  = 0o750
	filePermissions = 0o600
)

//go:embed templates/*.gotmpl
var templatesFS embed.FS

// Generator generates testify packages (assert, require) from the internal assertions package.
type Generator struct {
	options

	source *model.AssertionPackage
	ctx    *genCtx
	docs   model.Documentation
}

type genCtx struct {
	generateOptions

	index         map[string]string
	templates     map[string]*template.Template
	target        *model.AssertionPackage
	docs          *model.Documentation
	targetBase    string
	variantSuffix string          // filename suffix for the current build-variant (e.g. "_go126"), empty for the default
	rendered      map[string]bool // base filenames this run accounted for (written or removed-when-empty)
}

func New(source *model.AssertionPackage, opts ...Option) *Generator {
	return &Generator{
		options: optionsWithDefaults(opts),
		source:  source,
	}
}

func (g *Generator) Generate(opts ...GenerateOption) error {
	if err := g.initContext(opts); err != nil {
		return err
	}
	defer func() {
		if g.ctx.docs != nil {
			g.docs = *g.ctx.docs // before we leave, stash the transformed documentation for the current package
		}
		g.ctx = nil
	}()

	{
		// prepare steps

		if err := g.loadTemplates(); err != nil {
			return err
		}

		if err := g.transformModel(); err != nil {
			return err
		}

		g.buildDocs()
	}

	// Type constraints are not version-guarded (functions-only scope): generate them once,
	// from the full model, into the default (unsuffixed) file.
	if err := g.generateTypes(); err != nil {
		// assertion_types.gotmpl
		return err
	}

	// Build constraints are file-level in Go, so guarded functions must land in their own
	// generated files carrying the same //go:build line. We partition the functions by
	// constraint and render a parallel set of files per variant. Empty category files
	// (e.g. a go1.26 variant with no helpers) are skipped by render().
	full := g.ctx.target
	for _, constraint := range full.Functions.BuildVariants() {
		g.ctx.target = variantTarget(full, constraint)
		g.ctx.variantSuffix = model.GoBuildTag(constraint)
		if g.ctx.variantSuffix != "" {
			g.ctx.variantSuffix = "_" + g.ctx.variantSuffix
		}

		if err := g.generateVariant(); err != nil {
			return err
		}
	}
	g.ctx.target = full // restore the full model for Documentation()

	// Remove generated build-variant files left over from a variant that no longer exists
	// (e.g. the last go1.26-guarded assertion was deleted). Without this, those files would
	// linger and fail to compile against the now-missing source symbols.
	if err := g.sweepOrphanVariants(); err != nil {
		return err
	}

	return nil
}

// orphanVariantRx matches a generated build-variant filename for the given target package,
// e.g. "assert_assertions_go126.go" or "assert_forward_go126_test.go". The "_go<N>" infix is
// what distinguishes a variant file from the default (unsuffixed) ones, which are always
// regenerated and never swept.
func orphanVariantRx(targetBase string) *regexp.Regexp {
	return regexp.MustCompile(`^` + regexp.QuoteMeta(targetBase) + `_.+_go\d+(_test)?\.go$`)
}

// sweepOrphanVariants removes generated build-variant files in the target package directory
// that were not produced by this run. To guard against deleting anything we shouldn't, a file
// is removed only when it (a) matches the variant filename pattern, (b) was not rendered this
// run, and (c) actually carries our generated-code marker. Hand-authored files — even ones
// that happen to match the name pattern — are left untouched, and every removal is logged.
//
// Note on `go generate ./...`: that command eagerly snapshots every package's file list
// before running directives, so the run that removes an orphan may then fail with a benign
// "no such file" as go generate tries to scan the file we just deleted. The removal itself
// succeeds and a rerun is clean; invoking codegen directly (go run ./codegen/main.go) avoids
// the race entirely. This only happens when a whole guarded variant is deleted/renamed.
func (g *Generator) sweepOrphanVariants() error {
	dir := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("can't scan target folder for orphans: %w", err)
	}

	pattern := orphanVariantRx(g.ctx.targetBase)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !pattern.MatchString(name) || g.ctx.rendered[name] {
			continue
		}

		isGen, err := isGeneratedFile(filepath.Join(dir, name))
		if err != nil {
			return err
		}
		if !isGen {
			continue // not ours: never remove a hand-authored file
		}

		if err := os.Remove(filepath.Join(dir, name)); err != nil {
			return fmt.Errorf("can't remove orphaned variant file %q: %w", name, err)
		}
		// Announce every removal: an orphan disappearing should never be silent.
		log.Printf("codegen: removed orphaned build-variant file %s (its guarded source no longer exists)", filepath.Join(g.ctx.targetBase, name))
	}

	return nil
}

// isGeneratedFile reports whether the file carries our "DO NOT EDIT" generated-code marker.
func isGeneratedFile(path string) (bool, error) {
	data, err := os.ReadFile(path) //nolint:gosec // path is built from a controlled target directory listing
	if err != nil {
		return false, fmt.Errorf("can't read candidate orphan %q: %w", path, err)
	}

	return bytes.Contains(data, []byte("DO NOT EDIT.")), nil
}

// generateVariant renders every per-function artifact (functions, format, forward,
// helpers and their tests + examples) for the currently selected build-variant.
func (g *Generator) generateVariant() error {
	{
		// auto-generated assertions

		if err := g.generateAssertions(); err != nil {
			return err
		}

		if err := g.generateFormatFuncs(); err != nil {
			// assertion_format.gotmpl
			// requirement_format.gotmpl
			return err
		}

		if err := g.generateForwardFuncs(); err != nil {
			// assertion_forward.gotmpl
			// requirement_forward.gotmpl
			return err
		}

		if err := g.generateHelpers(); err != nil {
			return err
		}
	}

	{
		// auto-generated tests

		if err := g.generateAssertionsTests(); err != nil {
			return err
		}

		if err := g.generateFormatTests(); err != nil {
			return err
		}

		if err := g.generateForwardTests(); err != nil {
			return err
		}

		if err := g.generateExampleTests(); err != nil {
			return err
		}

		if err := g.generateHelpersTests(); err != nil {
			return err
		}
	}

	return nil
}

// variantTarget clones the transformed model, keeping only the functions belonging to the
// given build constraint, and stamps the constraint so templates emit the //go:build line.
func variantTarget(base *model.AssertionPackage, constraint string) *model.AssertionPackage {
	tgt := base.Clone()
	tgt.BuildConstraint = constraint

	filtered := tgt.Functions[:0:0]
	for _, fn := range base.Functions {
		if fn.GoBuild == constraint {
			filtered = append(filtered, fn)
		}
	}
	tgt.Functions = filtered

	return tgt
}

// Documentation yields the transformed package model as a [model.Documentation]
// usable by a generation step by the [DocGenerator].
func (g *Generator) Documentation() model.Documentation {
	return g.docs
}

func (g *Generator) initContext(opts []GenerateOption) error {
	// prepare options
	g.ctx = &genCtx{
		generateOptions: generateOptionsWithDefaults(opts),
		rendered:        make(map[string]bool),
	}
	if g.ctx.targetPkg == "" {
		return errors.New("a target package is required")
	}
	g.ctx.targetBase = path.Base(g.ctx.targetPkg) // perhaps find a better name

	if g.ctx.targetBase != pkgAssert && g.ctx.targetBase != pkgRequire {
		return fmt.Errorf(`unsupported target package. Expect pkgAssert or pkgRequire but got: %q`, g.ctx.targetBase)
	}

	if err := os.MkdirAll(filepath.Join(g.ctx.targetRoot, g.ctx.targetBase), dirPermissions); err != nil {
		return fmt.Errorf("can't make target folder: %w", err)
	}

	return nil
}

func (g *Generator) loadTemplates() error {
	const (
		tplExt            = ".gotmpl"
		expectedTemplates = 10
	)

	index := make(map[string]string, expectedTemplates)

	switch g.ctx.targetBase {
	case pkgAssert:
		g.loadAssertTemplates(index)
	case pkgRequire:
		g.loadRequireTemplates(index)
	default:
		panic(fmt.Errorf("internal error: invalid targetBase: %q", g.ctx.targetBase))
	}

	templates, err := loadTemplatesFromIndex(index, tplExt, templatesFS)
	if err != nil {
		return err
	}

	g.ctx.index = index
	g.ctx.templates = templates

	return nil
}

func (g *Generator) loadAssertTemplates(index map[string]string) {
	index["types"] = "assertion_types"
	index["assertions"] = "assertion_assertions"

	if g.ctx.generateTests {
		index["assertions_test"] = "assertion_assertions_test"
	}

	if g.ctx.generateHelpers {
		index["helpers"] = "assertion_helpers"
		if g.ctx.generateTests {
			index["helpers_test"] = "assertion_helpers_test"
		}
	}

	if g.ctx.generateExamples {
		index["examples_test"] = "assertion_examples_test"
	}

	if g.ctx.enableForward {
		index["forward"] = "assertion_forward"
		if g.ctx.generateTests {
			index["forward_test"] = "assertion_forward_test"
		}
	}

	if g.ctx.enableFormat {
		index["format"] = "assertion_format"
		if g.ctx.generateTests {
			index["format_test"] = "assertion_format_test"
		}
	}
}

func (g *Generator) loadRequireTemplates(index map[string]string) {
	index["types"] = "assertion_types"
	index["assertions"] = "requirement_assertions"

	if g.ctx.generateTests {
		index["assertions_test"] = "assertion_assertions_test"
	}

	if g.ctx.generateHelpers {
		index["helpers"] = "assertion_helpers"
		if g.ctx.generateTests {
			index["helpers_test"] = "assertion_helpers_test"
		}
	}

	if g.ctx.generateExamples {
		index["examples_test"] = "assertion_examples_test"
	}

	if g.ctx.enableForward {
		index["forward"] = "requirement_forward"
		if g.ctx.generateTests {
			index["forward_test"] = "assertion_forward_test"
		}
	}

	if g.ctx.enableFormat {
		index["format"] = "requirement_format"
		if g.ctx.generateTests {
			index["format_test"] = "assertion_format_test"
		}
	}
}

func (g *Generator) transformModel() error {
	tgt := g.source.Clone()

	tgt.Package = g.ctx.targetBase
	tgt.Receiver = "Assertions"
	tgt.EnableFormat = g.ctx.enableFormat
	tgt.EnableForward = g.ctx.enableForward
	tgt.EnableGenerics = g.ctx.enableGenerics
	tgt.EnableExamples = g.ctx.generateExamples
	tgt.RunnableExamples = g.ctx.runnableExamples /// instructs the doc generator to scan the generated packages to collect runnable examples
	if tgt.Imports == nil {
		tgt.Imports = make(model.ImportMap, 1)
	}
	tgt.Imports[assertions] = g.source.Package // add the import of our internal assertions package
	absRoot, err := filepath.Abs(g.ctx.targetRoot)
	if err != nil {
		return err
	}

	// NOTE: the use of [filepath.Rel] here imposes a constraint that the target resides on the same
	// drive as the source. This may cause issues, e.g. on windows (currently, our CI ensures any temp file
	// resides on the same drive as the source).
	testdata, err := filepath.Rel(filepath.Join(absRoot, g.ctx.targetBase), g.source.TestDataPath)
	if err != nil {
		return err
	}
	tgt.TestDataPath = testdata

	for i, fn := range tgt.Functions {
		tgt.Functions[i] = g.transformFunc(fn)
	}

	for i, vr := range tgt.Vars {
		if vr.Function == nil {
			continue
		}

		fn := g.transformFunc(*vr.Function)
		vr.Function = &fn
		tgt.Vars[i] = vr
	}

	for i, typ := range tgt.Types {
		if typ.Function == nil {
			continue
		}

		fn := g.transformFunc(*typ.Function)
		typ.Function = &fn
		tgt.Types[i] = typ
	}

	g.ctx.target = tgt

	return nil
}

func (g *Generator) transformFunc(fn model.Function) model.Function {
	fn.Params = g.transformArgs(fn.Params)
	fn.AllParams = g.transformArgs(fn.AllParams)
	fn.Returns = g.transformArgs(fn.Returns)
	if fn.Name == "FailNow" || g.ctx.targetBase == pkgRequire {
		fn.UseMock = mockWithFailNow
	} else {
		fn.UseMock = mock
	}

	for i, test := range fn.Tests {
		test.Pkg = g.ctx.targetBase // override the target package for this test
		fn.Tests[i] = test
	}

	return fn
}

// codeFile builds the path of a generated code file for a category, honoring the
// current build-variant suffix (e.g. "assert/assert_assertions_go126.go").
func (g *Generator) codeFile(category string) string {
	name := g.ctx.targetBase + "_" + category + g.ctx.variantSuffix + ".go"

	return filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, name)
}

// testFile builds the path of a generated test file for a category, honoring the
// current build-variant suffix (e.g. "assert/assert_assertions_go126_test.go").
func (g *Generator) testFile(category string) string {
	name := g.ctx.targetBase + "_" + category + g.ctx.variantSuffix + "_test.go"

	return filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, name)
}

func (g *Generator) generateTypes() error {
	// type constraints are unguarded: always written to the default (unsuffixed) file
	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_types.go")

	return g.render("types", file, g.ctx.target)
}

func (g *Generator) generateAssertions() error {
	return g.render("assertions", g.codeFile("assertions"), g.ctx.target)
}

func (g *Generator) generateFormatFuncs() error {
	if !g.ctx.enableFormat {
		return nil
	}

	return g.render("format", g.codeFile("format"), g.ctx.target)
}

func (g *Generator) generateForwardFuncs() error {
	if !g.ctx.enableForward {
		return nil
	}

	return g.render("forward", g.codeFile("forward"), g.ctx.target)
}

func (g *Generator) generateHelpers() error {
	if !g.ctx.generateHelpers {
		return nil
	}

	return g.render("helpers", g.codeFile("helpers"), g.ctx.target)
}

func (g *Generator) generateAssertionsTests() error {
	if !g.ctx.generateTests {
		return nil
	}

	return g.render("assertions_test", g.testFile("assertions"), g.ctx.target)
}

func (g *Generator) generateFormatTests() error {
	if !g.ctx.enableFormat || !g.ctx.generateTests {
		return nil
	}

	return g.render("format_test", g.testFile("format"), g.ctx.target)
}

func (g *Generator) generateForwardTests() error {
	if !g.ctx.enableForward || !g.ctx.generateTests {
		return nil
	}

	return g.render("forward_test", g.testFile("forward"), g.ctx.target)
}

func (g *Generator) generateExampleTests() error {
	if !g.ctx.generateExamples {
		return nil
	}

	return g.render("examples_test", g.testFile("examples"), g.ctx.target)
}

func (g *Generator) generateHelpersTests() error {
	if !g.ctx.generateHelpers || !g.ctx.generateTests {
		return nil
	}

	return g.render("helpers_test", g.testFile("helpers"), g.ctx.target)
}

func (g *Generator) render(name string, target string, data any) error {
	g.ctx.rendered[filepath.Base(target)] = true // account for this file even if render() drops it when empty

	return renderTemplate(
		g.ctx.index,
		g.ctx.templates,
		name,
		target,
		data,
		func(tpl *template.Template, target string, data any) error {
			return render(tpl, target, data, g.ctx.formatOptions)
		},
	)
}

func (g *Generator) transformArgs(in model.Parameters) model.Parameters {
	for j, arg := range in {
		if arg.Selector == "" {
			arg.Selector = assertions
		}
		in[j] = arg
	}

	return in
}

// buildDocs builds a hierarchy of documents for the current package.
//
// Documents are NOT rendered by the [Generator].
func (g *Generator) buildDocs() {
	if !g.ctx.generateDoc {
		return
	}

	docs := model.NewDocumentation()
	docs.Package = g.ctx.target

	// create a single document representing this package's contribution.
	//
	// This document will be reorganized by domain by the [DocGenerator] later on.
	// The [DocGenerator] carries out the heavy lifting when building a layout-ready Document.
	doc := model.Document{
		Title:   g.ctx.target.Package,
		Path:    g.ctx.targetBase, // "assert" or "require"
		Kind:    model.KindFolder,
		Package: g.ctx.target,
	}

	source := model.Document{
		Title:   g.source.Package,
		Path:    assertions,
		Kind:    model.KindFolder,
		Package: g.source,
	}

	docs.Documents = []model.Document{
		doc,
		source,
	}

	g.ctx.docs = docs
}
