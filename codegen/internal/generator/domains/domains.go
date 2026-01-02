// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package domains

import (
	"iter"
	"path"
	"slices"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

const (
	nodomain   = "common"
	assertions = "assertions"
)

// FlattenDocumentation flattens a nested documentation structure into a map of packages.
func FlattenDocumentation(documentation model.Documentation) map[string]model.Document {
	index := make(map[string]model.Document, len(documentation.Documents))

	flattenDocuments(index, documentation.Documents)

	return index
}

func flattenDocuments(index map[string]model.Document, docs []model.Document) {
	for _, doc := range docs {
		key := doc.Package.Package
		if _, ok := index[key]; !ok {
			index[key] = doc
		}

		flattenDocuments(index, doc.Documents)
	}
}

// Index represents a domain-based index of assertions organized by domain.
type Index struct {
	index       map[string]Entry
	tool        string
	copyright   string
	receiver    string
	rootPackage string
	header      string
}

// mapEntry is used for sorting domain entries.
type mapEntry struct {
	key   string
	value Entry
}

// Entries returns an iterator over domain entries, sorted alphabetically with "common" domain last.
func (idx Index) Entries() iter.Seq2[string, Entry] {
	list := make([]mapEntry, 0, len(idx.index))
	for k, v := range idx.index {
		list = append(list, mapEntry{key: k, value: v})
	}

	slices.SortFunc(list, compareMapEntries)

	return func(yield func(string, Entry) bool) {
		for _, entry := range list {
			if !yield(entry.key, entry.value) {
				return
			}
		}
	}
}

// compareMapEntries compares two mapEntry items by their key field.
//
// The "common" domain (nodomain) always sorts last, all others sort alphabetically.
func compareMapEntries(a, b mapEntry) int {
	if a.key == nodomain {
		if b.key == nodomain {
			return 0
		}

		return 1
	}
	if b.key == nodomain {
		return -1
	}

	return strings.Compare(a.key, b.key)
}

// Tool returns the tool name.
func (idx Index) Tool() string {
	return idx.tool
}

// Copyright returns the copyright string.
func (idx Index) Copyright() string {
	return idx.copyright
}

// Receiver returns the receiver name.
func (idx Index) Receiver() string {
	return idx.receiver
}

// RootPackage returns the root package path.
func (idx Index) RootPackage() string {
	return idx.rootPackage
}

// Header returns the header prepared by the Scanner.
func (idx Index) Header() string {
	return idx.header
}

// MakeDomainIndex creates a domain-based index from flattened documentation.
func MakeDomainIndex(docs map[string]model.Document) Index {
	// collect all functions from all packages (assert, require)
	describedDomains := make(map[string]string)
	discoveredDomains := make(map[string]Entry)

	var tool, receiver, copyright, rootPackage, header string

	for _, doc := range docs {
		data := doc.Package
		findDescribedDomains(data, describedDomains)
		pkg := data.Package // the generated package

		// unique values resolved once
		if tool == "" {
			tool = data.Tool
		}
		if receiver == "" {
			receiver = data.Receiver
		}
		if copyright == "" {
			copyright = data.Copyright
		}
		if rootPackage == "" {
			rootPackage = path.Dir(path.Dir(data.Imports[assertions]))
		}
		if header == "" {
			header = data.Header
		}

		discoverDomainsInFunctions(pkg, data, discoveredDomains)
		discoverDomainsInTypes(pkg, data, discoveredDomains)
		discoverDomainsInVariables(pkg, data, discoveredDomains)
		discoverDomainsInConstants(pkg, data, discoveredDomains)
	}

	// now add descriptions to domains
	for domain, description := range describedDomains {
		entry, ok := discoveredDomains[domain]
		if !ok {
			continue
		}
		entry.description = description
		discoveredDomains[domain] = entry
	}

	entry, ok := discoveredDomains[nodomain]
	if ok {
		entry.description = "Other uncategorized helpers"
		discoveredDomains[nodomain] = entry
	}

	return Index{
		index:       discoveredDomains,
		tool:        tool,
		receiver:    receiver,
		copyright:   copyright,
		header:      header,
		rootPackage: rootPackage,
	}
}

// findDescribedDomains look for annotations in the package comments
// that describe assertion domains.
//
// Example: github.com/go-openapi/testify/v2/internal/assertions/doc.go
//
// Comment format:
//
//	domain: description
func findDescribedDomains(data *model.AssertionPackage, describedDomains map[string]string) {
	for _, tagged := range data.ExtraComments {
		if tagged.Tag != model.CommentTagDomainDescription {
			continue
		}
		description := describedDomains[tagged.Key]
		if description == "" {
			describedDomains[tagged.Key] = tagged.Text
		}
	}
}

func discoverDomainsInFunctions(pkg string, data *model.AssertionPackage, discoveredDomains map[string]Entry) {
	for _, fn := range data.Functions {
		domain := fn.Domain
		if domain == "" {
			entry := discoveredDomains[nodomain]
			entry.AddPackage(pkg, data)
			fn.Domain = nodomain
			entry.AddFunction(pkg, fn)
			discoveredDomains[nodomain] = entry

			continue
		}

		entry := discoveredDomains[domain]
		entry.AddPackage(pkg, data)
		entry.AddFunction(pkg, fn)
		discoveredDomains[domain] = entry
	}
}

func discoverDomainsInTypes(pkg string, data *model.AssertionPackage, discoveredDomains map[string]Entry) {
	for _, ty := range data.Types {
		domain := ty.Domain
		if domain == "" {
			entry := discoveredDomains[nodomain]
			entry.AddPackage(pkg, data)
			ty.Domain = nodomain
			entry.AddType(pkg, ty)
			discoveredDomains[nodomain] = entry

			continue
		}

		entry := discoveredDomains[domain]
		entry.AddPackage(pkg, data)
		entry.AddType(pkg, ty)
		discoveredDomains[domain] = entry
	}
}

func discoverDomainsInVariables(pkg string, data *model.AssertionPackage, discoveredDomains map[string]Entry) {
	for _, va := range data.Vars {
		domain := va.Domain
		if domain == "" {
			entry := discoveredDomains[nodomain]
			entry.AddPackage(pkg, data)
			va.Domain = nodomain
			entry.AddVariable(pkg, va)
			discoveredDomains[nodomain] = entry

			continue
		}

		entry := discoveredDomains[domain]
		entry.AddPackage(pkg, data)
		entry.AddVariable(pkg, va)
		discoveredDomains[domain] = entry
	}
}

func discoverDomainsInConstants(pkg string, data *model.AssertionPackage, discoveredDomains map[string]Entry) {
	for _, co := range data.Consts {
		domain := co.Domain
		if domain == "" {
			entry := discoveredDomains[nodomain]
			entry.AddPackage(pkg, data)
			co.Domain = nodomain
			entry.AddConst(pkg, co)
			discoveredDomains[nodomain] = entry

			continue
		}

		entry := discoveredDomains[domain]
		entry.AddPackage(pkg, data)
		entry.AddConst(pkg, co)
		discoveredDomains[domain] = entry
	}
}

type key struct {
	pkg  string
	name string
}

func makeKey(pkg, name string) key {
	return key{
		pkg:  pkg,
		name: name,
	}
}

// Entry represents a discovered domain entry with associated functions, types, variables, and constants.
type Entry struct {
	description string
	packages    map[key]*model.AssertionPackage
	funcs       map[key]*model.Function
	typeDecls   map[key]*model.Ident
	varDecls    map[key]*model.Ident
	constDecls  map[key]*model.Ident
}

func makeEntry() Entry {
	return Entry{
		packages:   make(map[key]*model.AssertionPackage),
		funcs:      make(map[key]*model.Function),
		typeDecls:  make(map[key]*model.Ident),
		varDecls:   make(map[key]*model.Ident),
		constDecls: make(map[key]*model.Ident),
	}
}

// AddPackage adds a package to the entry.
func (e *Entry) AddPackage(pkg string, apkg *model.AssertionPackage) {
	if e.packages == nil {
		*e = makeEntry()
	}
	e.packages[makeKey(pkg, apkg.Package)] = apkg
}

// AddFunction adds a function to the entry.
func (e *Entry) AddFunction(pkg string, fn model.Function) {
	val := fn
	if e.packages == nil {
		*e = makeEntry()
	}
	e.funcs[makeKey(pkg, fn.Name)] = &val
}

// AddType adds a type declaration to the entry.
func (e *Entry) AddType(pkg string, id model.Ident) {
	val := id
	if e.packages == nil {
		*e = makeEntry()
	}
	e.typeDecls[makeKey(pkg, id.Name)] = &val
}

// AddVariable adds a variable declaration to the entry.
func (e *Entry) AddVariable(pkg string, id model.Ident) {
	val := id
	if e.packages == nil {
		*e = makeEntry()
	}
	e.varDecls[makeKey(pkg, id.Name)] = &val
}

// AddConst adds a constant declaration to the entry.
func (e *Entry) AddConst(pkg string, id model.Ident) {
	val := id
	if e.packages == nil {
		*e = makeEntry()
	}
	e.constDecls[makeKey(pkg, id.Name)] = &val
}

// ExtraPackages returns the list of extra packages (excluding the assertions package).
//
// This should correspond to the generated packages, e.g. "assert" and "require".
func (e Entry) ExtraPackages() model.ExtraPackages {
	if len(e.packages) == 0 {
		return nil
	}

	result := make(model.ExtraPackages, 0, len(e.packages)-1)
	for _, pkg := range e.packages {
		if pkg == nil {
			// safeguard
			continue
		}
		if path.Base(pkg.Package) == assertions {
			continue
		}
		result = append(result, pkg)
	}
	slices.SortFunc(result, comparePackages)

	return result
}

// comparePackages compares two AssertionPackages by their Package field.
func comparePackages(a, b *model.AssertionPackage) int {
	return strings.Compare(a.Package, b.Package)
}

// Functions returns the list of functions in this domain from the assertions package.
func (e Entry) Functions() []model.Function {
	result := make([]model.Function, 0, len(e.funcs))
	for key, fn := range e.funcs {
		if fn == nil {
			// safeguard
			continue
		}
		if path.Base(key.pkg) != assertions {
			continue
		}
		result = append(result, *fn)
	}
	slices.SortFunc(result, compareFunctions)

	return slices.Clip(result)
}

// Len returns the number of (filtered) functions for this entry.
func (e Entry) Len() (l int) {
	for key, fn := range e.funcs {
		if fn == nil {
			// safeguard
			continue
		}
		if path.Base(key.pkg) != assertions {
			continue
		}
		l++
	}

	return l
}

// compareFunctions compares two Functions by their Name field.
func compareFunctions(a, b model.Function) int {
	return strings.Compare(a.Name, b.Name)
}

// Types returns the list of type declarations in this domain.
func (e Entry) Types() []model.Ident {
	return e.sortedIdents(e.typeDecls)
}

// Vars returns the list of variable declarations in this domain.
func (e Entry) Vars() []model.Ident {
	return e.sortedIdents(e.varDecls)
}

// Consts returns the list of constant declarations in this domain.
func (e Entry) Consts() []model.Ident {
	return e.sortedIdents(e.constDecls)
}

// Description returns the description of this domain.
func (e Entry) Description() string {
	return e.description
}

func (e Entry) sortedIdents(idents map[key]*model.Ident) []model.Ident {
	result := make([]model.Ident, 0, len(idents))
	for key, id := range idents {
		if id == nil {
			// safeguard
			continue
		}
		if path.Base(key.pkg) != assertions {
			continue
		}
		result = append(result, *id)
	}
	slices.SortFunc(result, compareIdents)

	return result
}

// compareIdents compares two Idents by their Name field.
func compareIdents(a, b model.Ident) int {
	return strings.Compare(a.Name, b.Name)
}
