// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"fmt"
	"iter"
	"os"
	"path"
	"path/filepath"

	"github.com/go-openapi/testify/codegen/v2/internal/generator/domains"
	"github.com/go-openapi/testify/codegen/v2/internal/generator/funcmaps"
	"github.com/go-openapi/testify/codegen/v2/internal/model"
	exparser "github.com/go-openapi/testify/codegen/v2/internal/scanner/examples-parser"
)

const (
	// index page metadata.
	indexTitle       = "Assertions index"
	indexDescription = "Index of assertion domains"
	indexFile        = "_index.md"

	// sensible default preallocated slots.
	allocatedEntries = 15
)

type DocGenerator struct {
	options

	ctx *genCtx
	doc model.Documentation
}

func NewDocGenerator(doc model.Documentation, opts ...Option) *DocGenerator {
	return &DocGenerator{
		options: optionsWithDefaults(opts),
		doc:     doc,
	}
}

func (d *DocGenerator) Generate(opts ...GenerateOption) error {
	// prepare options
	d.ctx = &genCtx{
		generateOptions: generateOptionsWithDefaults(opts),
	}
	if d.ctx.targetDoc == "" {
		return errors.New("a target directory is required for docs")
	}

	if err := d.loadTemplates(); err != nil {
		return err
	}

	// capture testable examples from generated packages and attach them to
	// the model so templates may render their source code.
	if err := d.populateExamples(); err != nil {
		return err
	}

	// reorganize accumulated package-based docs into domain-based docs
	//
	// This iterator renders all domains in the desired order.
	domainDocs, extraUniqueValues := d.reorganizeByDomain()

	// generate an index page with all domains
	indexDoc := d.buildIndexDocument(domainDocs, extraUniqueValues)
	if err := d.generateDomainIndex(indexDoc); err != nil {
		return err
	}

	// generate one page per domain. Each document knows about the domain.
	for _, document := range domainDocs {
		if err := d.generateDomainPage(document); err != nil {
			return err
		}
	}

	return nil
}

type uniqueValues struct {
	tool        string
	receiver    string
	copyright   string
	header      string
	githubURL   string
	pkggodevURL string
}

func (d *DocGenerator) reorganizeByDomain() (iter.Seq2[string, model.Document], uniqueValues) {
	docs := domains.FlattenDocumentation(d.doc)
	discoveredDomains := domains.MakeDomainIndex(docs)
	githubURL := "https://" + path.Dir(discoveredDomains.RootPackage())
	pkggodevURL := "https://pkg.go.dev/" + discoveredDomains.RootPackage()

	return func(yield func(string, model.Document) bool) {
			weight := 1
			for domain, entry := range discoveredDomains.Entries() {
				doc := model.Document{
					Title:       funcmaps.Titleize(domain),
					Domain:      domain,
					Description: entry.Description(),
					Kind:        model.KindPage,
					File:        domain + ".md",
					Package: &model.AssertionPackage{
						Package:          assertions, // package that is the single source of truth
						Tool:             discoveredDomains.Tool(),
						Copyright:        discoveredDomains.Copyright(),
						Receiver:         discoveredDomains.Receiver(),
						Header:           discoveredDomains.Header(),
						EnableFormat:     d.ctx.enableFormat,
						EnableForward:    d.ctx.enableForward,
						EnableGenerics:   d.ctx.enableGenerics,
						EnableExamples:   d.ctx.generateExamples,
						RunnableExamples: d.ctx.runnableExamples,
						// skip package-level docstring
						// skip other package-level extra comments
						// filtered functions and types for this domain across all packages
						Functions: entry.Functions(),
						Types:     entry.Types(),
						Vars:      entry.Vars(),
						Consts:    entry.Consts(),
					},
					ExtraPackages: entry.ExtraPackages(),
					GitHubURL:     githubURL,
					PkgGoDevURL:   pkggodevURL,
					RefCount:      entry.Len(),
					Weight:        weight,
				}
				weight++

				// populate document context in all children
				doc.Package.Context = &doc
				for i, fn := range doc.Package.Functions {
					fn.Context = &doc
					doc.Package.Functions[i] = fn
				}

				if !yield(doc.Domain, doc) {
					return
				}
			}
		}, uniqueValues{
			// metadata that are unique
			tool:        discoveredDomains.Tool(),
			receiver:    discoveredDomains.Receiver(),
			copyright:   discoveredDomains.Copyright(),
			header:      discoveredDomains.Header(),
			githubURL:   githubURL,
			pkggodevURL: pkggodevURL,
		}
}

func (d *DocGenerator) buildIndexDocument(docsByDomain iter.Seq2[string, model.Document], extras uniqueValues) model.Document {
	doc := model.Document{
		Title:       indexTitle,
		Description: indexDescription,
		Kind:        model.KindIndex,
		File:        indexFile,
		Index:       buildIndexEntries(docsByDomain),
		Package: &model.AssertionPackage{
			Tool:      extras.tool,
			Copyright: extras.copyright,
			Receiver:  extras.receiver,
			Header:    extras.header,
			// skip everything else
		},
		GitHubURL:   extras.githubURL,
		PkgGoDevURL: extras.pkggodevURL,
	}

	doc.RefCount = len(doc.Index)

	return doc
}

func buildIndexEntries(docsByDomain iter.Seq2[string, model.Document]) []model.IndexEntry {
	entries := make([]model.IndexEntry, 0, allocatedEntries)

	weight := 1
	for domain, doc := range docsByDomain {
		entries = append(entries, model.IndexEntry{
			Name:        domain,
			Title:       doc.Title,
			Description: doc.Description,
			Link:        "./" + doc.File,
			RefCount:    len(doc.Package.Functions),
			Weight:      weight,
		})
		weight++
	}

	return entries
}

func (d *DocGenerator) generateDomainIndex(document model.Document) error {
	base := filepath.Join(d.ctx.targetRoot, d.ctx.targetDoc, document.Path)
	if err := os.MkdirAll(base, dirPermissions); err != nil {
		return fmt.Errorf("can't make target folder: %w", err)
	}
	return d.render("doc_index", filepath.Join(base, document.File), document)
}

func (d *DocGenerator) generateDomainPage(document model.Document) error {
	base := filepath.Join(d.ctx.targetRoot, d.ctx.targetDoc, document.Path)
	if err := os.MkdirAll(base, dirPermissions); err != nil {
		return fmt.Errorf("can't make target folder: %w", err)
	}

	return d.render("doc_page", filepath.Join(base, document.File), document)
}

func (d *DocGenerator) loadTemplates() error {
	const (
		tplExt            = ".md.gotmpl"
		expectedTemplates = 10
	)

	index := make(map[string]string, expectedTemplates)
	index["doc_index"] = "doc_index"
	index["doc_page"] = "doc_page"

	templates, err := loadTemplatesFromIndex(index, tplExt, templatesFS)
	if err != nil {
		return err
	}

	d.ctx.index = index
	d.ctx.templates = templates

	return nil
}

// populateExamples runs the examples-parser against all generated packages in the
// merged Documentation and attaches the discovered testable examples to the
// corresponding Function and Ident objects.
//
// This must run before [reorganizeByDomain] because domain discovery copies
// functions and types into domain entries.
func (d *DocGenerator) populateExamples() error {
	if !d.ctx.runnableExamples {
		return nil
	}

	docs := domains.FlattenDocumentation(d.doc)

	// derive the module root from the assertions import that every generated
	// package carries (e.g. "github.com/go-openapi/testify/v2").
	var rootPkg string
	for _, doc := range docs {
		if doc.Package != nil && doc.Package.Imports != nil {
			if assertionsPath, ok := doc.Package.Imports[assertions]; ok {
				rootPkg = path.Dir(path.Dir(assertionsPath))

				break
			}
		}
	}
	if rootPkg == "" {
		return nil // nothing to do
	}

	workDir, err := filepath.Abs(d.ctx.targetRoot)
	if err != nil {
		return fmt.Errorf("resolving target root: %w", err)
	}

	for _, doc := range docs {
		pkg := doc.Package
		if pkg == nil {
			continue
		}

		// Skip the internal assertions package: testable examples live in the
		// generated packages (assert, require), not in the source package.
		if path.Base(pkg.Package) == assertions {
			continue
		}

		importPath := rootPkg + "/" + pkg.Package
		examples, parseErr := exparser.New(importPath, exparser.WithWorkDir(workDir)).Parse()
		if parseErr != nil {
			return fmt.Errorf("parsing examples for %s: %w", pkg.Package, parseErr)
		}

		populateFunctionExamples(pkg, examples)
		populateIdentExamples(pkg.Types, examples)
	}

	return nil
}

func populateFunctionExamples(pkg *model.AssertionPackage, examples exparser.Examples) {
	for i, fn := range pkg.Functions {
		exs, ok := examples[fn.Name]
		if !ok {
			continue
		}
		renderables := make([]model.Renderable, len(exs))
		for j := range exs {
			renderables[j] = exs[j]
		}
		pkg.Functions[i].Examples = renderables
	}
}

func populateIdentExamples(idents []model.Ident, examples exparser.Examples) {
	for i, id := range idents {
		exs, ok := examples[id.Name]
		if !ok {
			continue
		}
		renderables := make([]model.Renderable, len(exs))
		for j := range exs {
			renderables[j] = exs[j]
		}
		idents[i].Examples = renderables
	}
}

func (d *DocGenerator) render(name string, target string, data any) error {
	return renderTemplate(
		d.ctx.index,
		d.ctx.templates,
		name,
		target,
		data,
		renderMD,
	)
}
