// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package domains

import (
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
)

// TestCompareIdents verifies Ident comparison by Name.
func TestCompareIdents(t *testing.T) {
	t.Parallel()

	for c := range compareIdentsCase() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := compareIdents(c.a, c.b)

			if (result < 0 && !c.aLessThanB) || (result > 0 && c.aLessThanB) || (result == 0 && c.aLessThanB) {
				t.Errorf("compareIdents(%q, %q) = %d, expected a < b: %v", c.a.Name, c.b.Name, result, c.aLessThanB)
			}
		})
	}
}

// TestCompareFunctions verifies Function comparison by Name.
func TestCompareFunctions(t *testing.T) {
	t.Parallel()

	for c := range compareFunctionsCase() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := compareFunctions(c.a, c.b)

			if (result < 0 && !c.aLessThanB) || (result > 0 && c.aLessThanB) || (result == 0 && c.aLessThanB) {
				t.Errorf("compareFunctions(%q, %q) = %d, expected a < b: %v", c.a.Name, c.b.Name, result, c.aLessThanB)
			}
		})
	}
}

// TestComparePackages verifies AssertionPackage comparison by Package field.
func TestComparePackages(t *testing.T) {
	t.Parallel()

	for c := range comparePackagesCase() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := comparePackages(c.a, c.b)

			if (result < 0 && !c.aLessThanB) || (result > 0 && c.aLessThanB) || (result == 0 && c.aLessThanB) {
				t.Errorf("comparePackages(%q, %q) = %d, expected a < b: %v", c.a.Package, c.b.Package, result, c.aLessThanB)
			}
		})
	}
}

// TestCompareMapEntries verifies mapEntry comparison with special "common" domain handling.
func TestCompareMapEntries(t *testing.T) {
	t.Parallel()

	for c := range compareMapEntriesCase() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := compareMapEntries(c.a, c.b)

			if (result < 0 && !c.aLessThanB) || (result > 0 && c.aLessThanB) || (result == 0 && c.aLessThanB) {
				t.Errorf("compareMapEntries(%q, %q) = %d, expected a < b: %v", c.a.key, c.b.key, result, c.aLessThanB)
			}
		})
	}
}

/* Test case iterators */

type compareIdentsTestCase struct {
	name       string
	a          model.Ident
	b          model.Ident
	aLessThanB bool
}

func compareIdentsCase() iter.Seq[compareIdentsTestCase] {
	return slices.Values([]compareIdentsTestCase{
		{
			name:       "a < b alphabetically",
			a:          model.Ident{Name: "Alpha"},
			b:          model.Ident{Name: "Beta"},
			aLessThanB: true,
		},
		{
			name:       "a > b alphabetically",
			a:          model.Ident{Name: "Zebra"},
			b:          model.Ident{Name: "Alpha"},
			aLessThanB: false,
		},
		{
			name:       "a == b",
			a:          model.Ident{Name: "Equal"},
			b:          model.Ident{Name: "Equal"},
			aLessThanB: false,
		},
	})
}

type compareFunctionsTestCase struct {
	name       string
	a          model.Function
	b          model.Function
	aLessThanB bool
}

func compareFunctionsCase() iter.Seq[compareFunctionsTestCase] {
	return slices.Values([]compareFunctionsTestCase{
		{
			name:       "a < b alphabetically",
			a:          model.Function{Name: "Contains"},
			b:          model.Function{Name: "Equal"},
			aLessThanB: true,
		},
		{
			name:       "a > b alphabetically",
			a:          model.Function{Name: "True"},
			b:          model.Function{Name: "False"},
			aLessThanB: false,
		},
		{
			name:       "a == b",
			a:          model.Function{Name: "Empty"},
			b:          model.Function{Name: "Empty"},
			aLessThanB: false,
		},
	})
}

type comparePackagesTestCase struct {
	name       string
	a          *model.AssertionPackage
	b          *model.AssertionPackage
	aLessThanB bool
}

func comparePackagesCase() iter.Seq[comparePackagesTestCase] {
	return slices.Values([]comparePackagesTestCase{
		{
			name: "a < b alphabetically",
			a: &model.AssertionPackage{
				Package: "github.com/go-openapi/testify/v2/assert",
			},
			b: &model.AssertionPackage{
				Package: "github.com/go-openapi/testify/v2/require",
			},
			aLessThanB: true,
		},
		{
			name: "a > b alphabetically",
			a: &model.AssertionPackage{
				Package: "github.com/some/pkg/z",
			},
			b: &model.AssertionPackage{
				Package: "github.com/some/pkg/a",
			},
			aLessThanB: false,
		},
		{
			name: "a == b",
			a: &model.AssertionPackage{
				Package: "github.com/go-openapi/testify/v2/assert",
			},
			b: &model.AssertionPackage{
				Package: "github.com/go-openapi/testify/v2/assert",
			},
			aLessThanB: false,
		},
	})
}

type compareMapEntriesTestCase struct {
	name       string
	a          mapEntry
	b          mapEntry
	aLessThanB bool
}

func compareMapEntriesCase() iter.Seq[compareMapEntriesTestCase] {
	return slices.Values([]compareMapEntriesTestCase{
		{
			name:       "both regular domains, a < b",
			a:          mapEntry{key: "alpha"},
			b:          mapEntry{key: "beta"},
			aLessThanB: true,
		},
		{
			name:       "both regular domains, a > b",
			a:          mapEntry{key: "zebra"},
			b:          mapEntry{key: "alpha"},
			aLessThanB: false,
		},
		{
			name:       "both regular domains, a == b",
			a:          mapEntry{key: "equal"},
			b:          mapEntry{key: "equal"},
			aLessThanB: false,
		},
		{
			name:       "a is common domain, b is regular (common sorts last)",
			a:          mapEntry{key: "common"},
			b:          mapEntry{key: "alpha"},
			aLessThanB: false, // common > alpha (sorts after)
		},
		{
			name:       "a is regular, b is common domain (common sorts last)",
			a:          mapEntry{key: "zebra"},
			b:          mapEntry{key: "common"},
			aLessThanB: true, // zebra < common (common sorts after)
		},
		{
			name:       "both are common domain",
			a:          mapEntry{key: "common"},
			b:          mapEntry{key: "common"},
			aLessThanB: false,
		},
	})
}
