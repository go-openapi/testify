// This program reads all assertion functions from the assert package and
// automatically generates the corresponding requires and forwarded assertions

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/doc"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/go-openapi/testify/v2/_codegen/internal/imports"
)

//nolint:gochecknoglobals // ok to register flags as globals
var (
	pkg       = flag.String("assert-path", "github.com/go-openapi/testify/v2/assert", "Path to the assert package")
	includeF  = flag.Bool("include-format-funcs", false, "include format functions such as Errorf and Equalf")
	outputPkg = flag.String("output-package", "", "package for the resulting code")
	tmplFile  = flag.String("template", "", "What file to load the function template from")
	out       = flag.String("out", "", "What file to write the source code to")
)

func main() {
	flag.Parse()

	scope, docs, err := parsePackageSource(*pkg)
	if err != nil {
		log.Fatal(err)
	}

	importer, funcs, err := analyzeCode(scope, docs)
	if err != nil {
		log.Fatal(err)
	}

	if err := generateCode(importer, funcs); err != nil {
		log.Fatal(err)
	}
}

func generateCode(importer imports.Importer, funcs []testFunc) error {
	buff := bytes.NewBuffer(nil)

	tmplHead, tmplFunc, err := parseTemplates()
	if err != nil {
		return err
	}

	// Generate header
	if err := tmplHead.Execute(buff, struct {
		Name    string
		Imports map[string]string
	}{
		*outputPkg,
		importer.Imports(),
	}); err != nil {
		return err
	}

	// Generate funcs
	for _, fn := range funcs {
		buff.WriteString("\n\n")
		if err := tmplFunc.Execute(buff, &fn); err != nil {
			return err
		}
	}

	code, err := format.Source(buff.Bytes())
	if err != nil {
		return err
	}

	// Write file
	output, err := outputFile()
	if err != nil {
		return err
	}
	defer output.Close()
	_, err = io.Copy(output, bytes.NewReader(code))
	return err
}

func parseTemplates() (*template.Template, *template.Template, error) {
	tmplHead, err := template.New("header").Parse(headerTemplate)
	if err != nil {
		return nil, nil, err
	}
	var funcTemplate string
	if *tmplFile != "" {
		f, err := os.ReadFile(*tmplFile)
		if err != nil {
			return nil, nil, err
		}
		funcTemplate = string(f)
	} else {
		funcTemplate = defaultTemplate
	}

	tmpl, err := template.New("function").Funcs(template.FuncMap{
		"replace": strings.ReplaceAll,
	}).Parse(funcTemplate)
	if err != nil {
		return nil, nil, err
	}

	return tmplHead, tmpl, nil
}

func outputFile() (*os.File, error) {
	filename := *out
	if filename == "-" || (filename == "" && *tmplFile == "") {
		return os.Stdout, nil
	}
	if filename == "" {
		filename = strings.TrimSuffix(strings.TrimSuffix(*tmplFile, ".tmpl"), ".go") + ".go"
	}
	return os.Create(filename)
}

// analyzeCode takes the types scope and the docs and returns the import
// information and information about all the assertion functions.
func analyzeCode(scope *types.Scope, docs *doc.Package) (imports.Importer, []testFunc, error) { //nolint:ireturn // ok to use what stdlib does here
	underlying := scope.Lookup("TestingT").Type().Underlying()
	testingT, ok := underlying.(*types.Interface)
	if !ok {
		panic(fmt.Errorf("internal error: expected go type to resolve as *types.Interface but got: %T", underlying))
	}

	importer := imports.New(*outputPkg)
	funcs := make([]testFunc, 0, len(docs.Funcs))
	// Go through all the top level functions
	for _, fdocs := range docs.Funcs {
		// Find the function
		obj := scope.Lookup(fdocs.Name)

		fn, ok := obj.(*types.Func)
		if !ok {
			continue
		}

		// Check function signature has at least two arguments
		sig, ok := fn.Type().(*types.Signature)
		if !ok {
			return nil, nil, fmt.Errorf("internal error: expected go type to resolve as *types.Signature but got: %T", sig)
		}

		const minParams = 2
		if sig.Params().Len() < minParams {
			continue
		}
		// Check first argument is of type testingT
		first, ok := sig.Params().At(0).Type().(*types.Named)
		if !ok {
			continue
		}

		firstType, ok := first.Underlying().(*types.Interface)
		if !ok {
			continue
		}
		if !types.Implements(firstType, testingT) {
			continue
		}

		// Skip functions ending with f
		if strings.HasSuffix(fdocs.Name, "f") && !*includeF {
			continue
		}

		funcs = append(funcs, testFunc{*outputPkg, fdocs, fn})
		importer.AddImportsFrom(sig.Params())
	}

	return importer, funcs, nil
}

// parsePackageSource returns the types scope and the package documentation from the package.
func parsePackageSource(pkg string) (*types.Scope, *doc.Package, error) {
	pd, err := build.Import(pkg, ".", 0)
	if err != nil {
		return nil, nil, err
	}

	fset := token.NewFileSet()
	files := make(map[string]*ast.File)
	fileList := make([]*ast.File, len(pd.GoFiles))
	for i, fname := range pd.GoFiles {
		src, err := os.ReadFile(path.Join(pd.Dir, fname))
		if err != nil {
			return nil, nil, err
		}
		f, err := parser.ParseFile(fset, fname, src, parser.ParseComments|parser.AllErrors)
		if err != nil {
			return nil, nil, err
		}
		files[fname] = f
		fileList[i] = f
	}

	cfg := types.Config{
		Importer: importer.ForCompiler(token.NewFileSet(), "source", nil),
	}
	info := types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	tp, err := cfg.Check(pkg, fset, fileList, &info)
	if err != nil {
		return nil, nil, err
	}

	scope := tp.Scope()

	ap, _ := ast.NewPackage(fset, files, nil, nil) //nolint:staticcheck // will need more work to upgrade
	docs := doc.New(ap, pkg, 0)

	return scope, docs, nil
}

type testFunc struct {
	CurrentPkg string
	DocInfo    *doc.Func
	TypeInfo   *types.Func
}

func (f *testFunc) Qualifier(p *types.Package) string {
	if p == nil || p.Name() == f.CurrentPkg {
		return ""
	}
	return p.Name()
}

func (f *testFunc) Params() string {
	sig, ok := f.TypeInfo.Type().(*types.Signature)
	if !ok {
		panic(fmt.Errorf("internal error: expected go type to resolve as *types.Interface but got: %T", f.TypeInfo.Type()))
	}
	params := sig.Params()
	var p strings.Builder
	comma := ""
	to := params.Len()
	var i int

	if sig.Variadic() {
		to--
	}
	for i = 1; i < to; i++ {
		param := params.At(i)
		p.WriteString(fmt.Sprintf("%s%s %s", comma, param.Name(), types.TypeString(param.Type(), f.Qualifier)))
		comma = ", "
	}
	if sig.Variadic() {
		param := params.At(params.Len() - 1)
		slice, ok := param.Type().(*types.Slice)
		if !ok {
			panic(fmt.Errorf("internal error: expected go type to resolve as *types.Slice but got: %T", param.Type()))
		}
		p.WriteString(fmt.Sprintf("%s%s ...%s", comma, param.Name(), types.TypeString(slice.Elem(), f.Qualifier)))
	}

	return p.String()
}

func (f *testFunc) ForwardedParams() string {
	sig, ok := f.TypeInfo.Type().(*types.Signature)
	if !ok {
		panic(fmt.Errorf("internal error: expected go type to resolve as *types.Signature but got: %T", sig))
	}
	params := sig.Params()
	var p strings.Builder
	comma := ""
	to := params.Len()
	var i int

	if sig.Variadic() {
		to--
	}
	for i = 1; i < to; i++ {
		param := params.At(i)
		p.WriteString(fmt.Sprintf("%s%s", comma, param.Name()))
		comma = ", "
	}
	if sig.Variadic() {
		param := params.At(params.Len() - 1)
		p.WriteString(fmt.Sprintf("%s%s...", comma, param.Name()))
	}
	return p.String()
}

func (f *testFunc) ParamsFormat() string {
	return strings.Replace(f.Params(), "msgAndArgs", "msg string, args", 1)
}

func (f *testFunc) ForwardedParamsFormat() string {
	return strings.Replace(f.ForwardedParams(), "msgAndArgs", "append([]any{msg}, args...)", 1)
}

func (f *testFunc) Comment() string {
	return "// " + strings.ReplaceAll(strings.TrimSpace(f.DocInfo.Doc), "\n", "\n// ")
}

func (f *testFunc) CommentFormat() string {
	search := f.DocInfo.Name
	replace := f.DocInfo.Name + "f"
	comment := strings.ReplaceAll(f.Comment(), search, replace)
	exp := regexp.MustCompile(replace + `\(((\(\)|[^\n])+)\)`)
	return exp.ReplaceAllString(comment, replace+`($1, "error message %s", "formatted")`)
}

func (f *testFunc) CommentWithoutT(receiver string) string {
	search := fmt.Sprintf("assert.%s(t, ", f.DocInfo.Name)
	replace := fmt.Sprintf("%s.%s(", receiver, f.DocInfo.Name)
	return strings.ReplaceAll(f.Comment(), search, replace)
}

// Standard header https://go.dev/s/generatedcode.
const headerTemplate = `// Code generated with github.com/go-openapi/testify/v2/_codegen; DO NOT EDIT.

package {{.Name}}

import (
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)
`

const defaultTemplate = `{{.Comment}}
func (fwd *AssertionsForwarder) {{.DocInfo.Name}}({{.Params}}) bool {
	return assert.{{.DocInfo.Name}}({{.ForwardedParams}})
}`
