// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
	"golang.org/x/tools/imports"
)

const (
	pkgRequire = "require"
	pkgAssert  = "assert"
	assertions = "assertions"
)

const (
	dirPermissions  = 0750
	filePermissions = 0600
)

var (
	//go:embed templates/*.gotmpl
	templatesFS embed.FS
)

// Generator generates testify packages (assert, require) from the internal assertions package.
type Generator struct {
	options

	source *model.AssertionPackage
	ctx    *genCtx
}

type genCtx struct {
	generateOptions

	index      map[string]string
	templates  map[string]*template.Template
	target     *model.AssertionPackage
	targetBase string
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
		g.ctx = nil
	}()

	if err := g.loadTemplates(); err != nil {
		return err
	}

	if err := g.transformModel(); err != nil {
		return err
	}

	{
		// auto-generated assertions

		if err := g.generateTypes(); err != nil {
			// assertion_types.gotmpl
			return err
		}

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

		return g.generateHelpersTests()
	}
}

func (g *Generator) initContext(opts []GenerateOption) error {
	// prepare options
	g.ctx = &genCtx{
		generateOptions: generateOptionsWithDefaults(opts),
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
	const expectedTemplates = 10
	index := make(map[string]string, expectedTemplates)

	switch g.ctx.targetBase {
	case pkgAssert:
		g.loadAssertTemplates(index)
	case pkgRequire:
		g.loadRequireTemplates(index)
	default:
		panic(fmt.Errorf("internal error: invalid targetBase: %q", g.ctx.targetBase))
	}

	g.ctx.index = index
	needed := make([]string, 0, len(index))
	for _, v := range index {
		needed = append(needed, v)
	}
	sort.Strings(needed)

	g.ctx.templates = make(map[string]*template.Template, len(needed))
	for _, name := range needed {
		file := name + ".gotmpl"
		tpl, err := template.New(file).Funcs(funcMap()).ParseFS(templatesFS, path.Join("templates", file))
		if err != nil {
			return fmt.Errorf("failed to load template %q from %q: %w", name, file, err)
		}

		g.ctx.templates[name] = tpl
	}

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
	tgt.RunnableExamples = g.ctx.runnableExamples
	if tgt.Imports == nil {
		tgt.Imports = make(model.ImportMap, 1)
	}
	tgt.Imports[assertions] = g.source.Package // add the import of our internal assertions package
	absRoot, err := filepath.Abs(g.ctx.targetRoot)
	if err != nil {
		return err
	}

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
		fn.UseMock = "mockFailNowT"
	} else {
		fn.UseMock = "mockT"
	}

	return fn
}

func (g *Generator) generateTypes() error {
	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_types.go")

	return g.render("types", file, g.ctx.target)
}

func (g *Generator) generateAssertions() error {
	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_assertions.go")

	return g.render("assertions", file, g.ctx.target)
}

func (g *Generator) generateFormatFuncs() error {
	if !g.ctx.enableFormat {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_format.go")

	return g.render("format", file, g.ctx.target)
}

func (g *Generator) generateForwardFuncs() error {
	if !g.ctx.enableForward {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_forward.go")

	return g.render("forward", file, g.ctx.target)
}

func (g *Generator) generateHelpers() error {
	if !g.ctx.generateHelpers {
		return nil
	}
	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_helpers.go")

	return g.render("helpers", file, g.ctx.target)
}

func (g *Generator) generateAssertionsTests() error {
	if !g.ctx.generateTests {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_assertions_test.go")

	return g.render("assertions_test", file, g.ctx.target)
}

func (g *Generator) generateFormatTests() error {
	if !g.ctx.enableFormat || !g.ctx.generateTests {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_format_test.go")

	return g.render("format_test", file, g.ctx.target)
}

func (g *Generator) generateForwardTests() error {
	if !g.ctx.enableForward || !g.ctx.generateTests {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_forward_test.go")

	return g.render("forward_test", file, g.ctx.target)
}

func (g *Generator) generateExampleTests() error {
	if !g.ctx.generateExamples {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_examples_test.go")

	return g.render("examples_test", file, g.ctx.target)
}

func (g *Generator) generateHelpersTests() error {
	if !g.ctx.generateHelpers || !g.ctx.generateTests {
		return nil
	}

	file := filepath.Join(g.ctx.targetRoot, g.ctx.targetBase, g.ctx.targetBase+"_helpers_test.go")

	return g.render("helpers_test", file, g.ctx.target)
}

func (g *Generator) render(name string, target string, data any) error {
	tplName, ok := g.ctx.index[name]
	if !ok {
		panic(fmt.Errorf("internal error: expect template name %q", name))
	}

	tpl, ok := g.ctx.templates[tplName]
	if !ok {
		panic(fmt.Errorf("internal error: expect template %q", name))
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("error executing template %q (%s): %w", name, tplName, err)
	}

	formatted, err := imports.Process(target, buffer.Bytes(), g.ctx.formatOptions)
	if err != nil {
		_ = os.WriteFile(target, buffer.Bytes(), filePermissions)
		return fmt.Errorf("error formatting go code: %w:%w", err, fmt.Errorf("details available at: %v", target))
	}

	return os.WriteFile(target, formatted, filePermissions)
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
