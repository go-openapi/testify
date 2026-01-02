package model

import (
	"slices"
)

type Documentation struct {
	Package   *AssertionPackage // the complete source package being documented
	Documents []Document        // Documents is a collection of markdown pages or folders
}

// NewDocumentation builds an empty [Documentation].
func NewDocumentation() *Documentation {
	return &Documentation{}
}

func (d *Documentation) Merge(doc Documentation) {
	if len(d.Documents) == 0 {
		// just replace
		*d = doc

		return
	}

	if len(d.Documents) == 1 {
		// creates a hierarchy
		current := Document{
			Title:     d.Package.Package,
			Path:      d.Package.Package,
			Kind:      KindFolder,
			Package:   d.Package,
			Documents: d.Documents,
		}

		next := Document{
			Title:     doc.Package.Package,
			Path:      doc.Package.Package,
			Kind:      KindFolder,
			Package:   doc.Package,
			Documents: doc.Documents,
		}

		d.Package = nil
		d.Documents = []Document{
			current,
			next,
		}

		return
	}

	// simply append a new folder
	d.Documents = append(d.Documents, Document{
		Title:     doc.Package.Package,
		Path:      doc.Package.Package,
		Kind:      KindFolder,
		Package:   doc.Package,
		Documents: doc.Documents,
	})
}

type KindDoc uint8

const (
	KindPage KindDoc = iota
	KindIndex
	KindFolder
)

type Document struct {
	Title         string
	Domain        string
	Description   string
	Path          string // document folder, relative to its parent if any
	File          string // document name, e.g. _index.md, string.md
	GitHubURL     string
	PkgGoDevURL   string
	Kind          KindDoc           // page or index or folder
	Documents     []Document        // for folders, their content [Document]
	Index         []IndexEntry      // for indexes, their entry list
	Package       *AssertionPackage // subset of the package that pertains to this document
	ExtraPackages ExtraPackages
	RefCount      int
	Weight        int
}

type ExtraPackages []*AssertionPackage

func (pkgs ExtraPackages) LookupFunction(name string) []FunctionWithContext {
	const defaultVariants = 8
	result := make([]FunctionWithContext, 0, defaultVariants)

	for _, pkg := range pkgs {
		for _, fn := range pkg.Functions {
			if fn.Name == name {
				result = append(result, FunctionWithContext{
					Function:       fn,
					Package:        pkg.Package,
					Receiver:       pkg.Receiver,
					EnableFormat:   pkg.EnableFormat,
					EnableForward:  pkg.EnableForward,
					EnableGenerics: pkg.EnableGenerics,
					EnableExamples: pkg.EnableExamples,
				})
				break
			}
		}
	}

	return slices.Clip(result)
}

type FunctionWithContext struct {
	Function

	Package        string
	Receiver       string
	EnableFormat   bool
	EnableForward  bool
	EnableGenerics bool
	EnableExamples bool
}

type IndexEntry struct {
	Name        string
	Title       string
	Description string
	Link        string
	RefCount    int
	Weight      int
}
