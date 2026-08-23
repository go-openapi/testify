// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"slices"
	"testing"
)

func TestGoBuildTag(t *testing.T) {
	cases := map[string]string{
		"":        "",
		"go1.26":  "go126",
		"go1.27":  "go127",
		"foo bar": "foobar",
	}
	for in, want := range cases {
		if got := GoBuildTag(in); got != want {
			t.Errorf("GoBuildTag(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestBuildVariants(t *testing.T) {
	fns := Functions{
		{Name: "A", GoBuild: ""},
		{Name: "B", GoBuild: "go1.27"},
		{Name: "C", GoBuild: "go1.26"},
		{Name: "D", GoBuild: "go1.26"},
		{Name: "E", GoBuild: ""},
	}

	got := fns.BuildVariants()
	want := []string{"", "go1.26", "go1.27"} // default first, rest sorted

	if !slices.Equal(got, want) {
		t.Errorf("BuildVariants() = %v, want %v", got, want)
	}
}

func TestForwardGoBuild(t *testing.T) {
	cases := []struct {
		name string
		fn   Function
		want string
	}{
		{"plain function", Function{Name: "Equal"}, ""},
		{"guarded function", Function{Name: "ErrorAs", GoBuild: "go1.26"}, "go1.26"},
		{"generic function", Function{Name: "EqualT", IsGeneric: true}, GenericMethodsGoBuild},
		{"generic below the floor", Function{Name: "ErrorAsType", IsGeneric: true, GoBuild: "go1.26"}, GenericMethodsGoBuild},
		{"generic above the floor", Function{Name: "FutureT", IsGeneric: true, GoBuild: "go1.28"}, "go1.28"},
	}

	for _, c := range cases {
		if got := c.fn.ForwardGoBuild(); got != c.want {
			t.Errorf("%s: ForwardGoBuild() = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestForwardBuildVariants(t *testing.T) {
	fns := Functions{
		{Name: "Equal"},
		{Name: "ErrorAs", GoBuild: "go1.26"},
		{Name: "EqualT", IsGeneric: true},
		{Name: "ErrorAsType", IsGeneric: true, GoBuild: "go1.26"},
		{Name: "FutureT", IsGeneric: true, GoBuild: "go1.28"},
		{Name: "CallerInfo", IsHelper: true, IsGeneric: true},
		{Name: "New", IsConstructor: true},
	}

	got := fns.ForwardBuildVariants()
	want := []string{"", "go1.26", "go1.27", "go1.28"} // default first, rest sorted

	if !slices.Equal(got, want) {
		t.Errorf("ForwardBuildVariants() = %v, want %v", got, want)
	}
}

func TestSourceGoBuild(t *testing.T) {
	// generic or not, the source guard is what the source file carries
	fn := Function{Name: "EqualT", IsGeneric: true}
	if got := fn.SourceGoBuild(); got != "" {
		t.Errorf("SourceGoBuild() = %q, want %q", got, "")
	}

	fn = Function{Name: "ErrorAsType", IsGeneric: true, GoBuild: "go1.26"}
	if got := fn.SourceGoBuild(); got != "go1.26" {
		t.Errorf("SourceGoBuild() = %q, want %q", got, "go1.26")
	}
}
