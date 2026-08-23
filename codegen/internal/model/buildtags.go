// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"slices"
	"strconv"
	"strings"
)

// GenericMethodsGoBuild is the //go:build constraint under which generic assertions may be
// exposed as methods of the Assertions receiver. go1.27 is the first release that accepts
// type parameters on methods.
const GenericMethodsGoBuild = "go1.27"

// GoBuildTag converts a //go:build version expression into a filename-safe suffix tag,
// e.g. "go1.26" -> "go126". Returns "" for an empty constraint (the default partition).
//
// Only simple single-term go-version constraints are supported for now; dots and spaces
// are stripped to keep the suffix free of any implicit GOOS/GOARCH filename semantics.
func GoBuildTag(expr string) string {
	if expr == "" {
		return ""
	}

	r := strings.NewReplacer(".", "", " ", "")

	return r.Replace(expr)
}

// goBuildMinor extracts N from a "go1.N" constraint. It returns 0 for an empty constraint
// and for anything that isn't a plain go-version term.
func goBuildMinor(expr string) int {
	rest, ok := strings.CutPrefix(expr, "go1.")
	if !ok {
		return 0
	}

	minor, err := strconv.Atoi(rest)
	if err != nil {
		return 0
	}

	return minor
}

// maxGoBuild returns the stricter of two go-version constraints, i.e. the one a file must
// carry to satisfy both. "" (no constraint) is weaker than any go1.N.
func maxGoBuild(a, b string) string {
	if goBuildMinor(a) >= goBuildMinor(b) {
		return a
	}

	return b
}

// BuildVariants returns the distinct //go:build constraints across the functions,
// in deterministic order, always starting with the default (empty) partition.
func (f Functions) BuildVariants() []string {
	seen := map[string]struct{}{"": {}}
	variants := []string{""} // default partition is always present

	for _, fn := range f {
		if fn.GoBuild == "" {
			continue
		}
		if _, ok := seen[fn.GoBuild]; ok {
			continue
		}
		seen[fn.GoBuild] = struct{}{}
		variants = append(variants, fn.GoBuild)
	}

	slices.Sort(variants[1:]) // keep "" first, sort the rest for stable output

	return variants
}

// SourceGoBuild returns the //go:build constraint of the source file this function was
// declared in. It exists as a method so it can be passed around as a partitioning key
// alongside [Function.ForwardGoBuild].
func (f Function) SourceGoBuild() string {
	return f.GoBuild
}

// ForwardGoBuild returns the //go:build constraint the forward method of this function must
// carry. A generic assertion becomes a method only from [GenericMethodsGoBuild] onwards, so
// its method is guarded by whichever of the two constraints is stricter: the guard of the
// source file it was declared in, or the generic-methods floor.
func (f Function) ForwardGoBuild() string {
	if !f.IsGeneric {
		return f.GoBuild
	}

	return maxGoBuild(f.GoBuild, GenericMethodsGoBuild)
}

// ForwardBuildVariants returns the distinct //go:build constraints across the forward
// methods, in deterministic order, always starting with the default (empty) partition.
//
// It is the counterpart of [Functions.BuildVariants] for the forward files, which partition
// on [Function.ForwardGoBuild] instead of [Function.GoBuild]: the generic assertions carry a
// go1.27 guard as methods even though their source files are unguarded.
func (f Functions) ForwardBuildVariants() []string {
	seen := map[string]struct{}{"": {}}
	variants := []string{""} // default partition is always present

	for _, fn := range f {
		if fn.IsHelper || fn.IsConstructor {
			continue
		}

		constraint := fn.ForwardGoBuild()
		if _, ok := seen[constraint]; ok {
			continue
		}
		seen[constraint] = struct{}{}
		variants = append(variants, constraint)
	}

	slices.Sort(variants[1:]) // keep "" first, sort the rest for stable output

	return variants
}
