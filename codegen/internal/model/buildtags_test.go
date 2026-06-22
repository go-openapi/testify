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
