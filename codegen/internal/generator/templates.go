// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"embed"
	"fmt"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/go-openapi/testify/codegen/v2/internal/generator/funcmaps"
)

// buildTemplateIndex extracts template names from the index and returns them sorted.
//
// This helper reduces duplication between Generator and DocGenerator template loading.
func buildTemplateIndex(index map[string]string) []string {
	needed := make([]string, 0, len(index))
	for _, v := range index {
		needed = append(needed, v)
	}
	sort.Strings(needed)

	return needed
}

// loadTemplatesFromIndex loads templates from an embedded filesystem using a name-to-file index.
//
// Parameters:
//   - index: map from logical template name to template file name (without extension)
//   - tplExt: template file extension (e.g., ".gotmpl", ".md.gotmpl")
//   - fs: embedded filesystem containing template files in "templates/" directory
//
// Returns a map from template name to parsed template, or an error if any template fails to load.
//
// This function consolidates the common template loading logic shared between Generator
// and DocGenerator, eliminating ~25 lines of duplicate code.
func loadTemplatesFromIndex(
	index map[string]string,
	tplExt string,
	fs embed.FS,
) (map[string]*template.Template, error) {
	needed := buildTemplateIndex(index)

	templates := make(map[string]*template.Template, len(needed))
	for _, name := range needed {
		file := name + tplExt
		files := []string{path.Join("templates", file)}
		if strings.Contains(name, "_test") { // test templates use a set of shared definitions
			files = append(files, path.Join("templates", "assertion_test_shared.gotmpl"))
		}
		tpl, err := template.New(file).Funcs(funcmaps.FuncMap()).ParseFS(fs, files...)
		if err != nil {
			return nil, fmt.Errorf("failed to load template %q from %q: %w", name, file, err)
		}
		templates[name] = tpl
	}

	return templates, nil
}

// renderFunc is a function that executes a template and writes the result to a file.
type renderFunc func(*template.Template, string, any) error

// renderTemplate executes a template by name and writes the result using the provided render function.
//
// Parameters:
//   - index: map from logical template name to template file name
//   - templates: map from template file name to parsed template
//   - name: logical template name to render
//   - target: output file path
//   - data: data to pass to the template
//   - renderFn: function that executes the template and writes output (e.g., render or renderMD)
//
// Returns an error if the template name is not found in the index, the template is not loaded,
// or the render function fails.
//
// This function consolidates the common rendering logic shared between Generator and DocGenerator,
// eliminating ~15 lines of duplicate code. The only difference between the two generators is
// the final render function (render vs renderMD), which is passed as a parameter.
func renderTemplate(
	index map[string]string,
	templates map[string]*template.Template,
	name string,
	target string,
	data any,
	renderFn renderFunc,
) error {
	tplName, ok := index[name]
	if !ok {
		return fmt.Errorf("internal error: expect template name %q", name)
	}

	tpl, ok := templates[tplName]
	if !ok {
		return fmt.Errorf("internal error: expect template %q", name)
	}

	return renderFn(tpl, target, data)
}
