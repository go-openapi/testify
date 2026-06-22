// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"os"
	"testing"
)

// TestBuildConstraintDetection verifies the scanner attaches the //go:build version
// constraint of the source file to each guarded function.
//
// It relies on ErrorAsType, guarded behind //go:build go1.26 (see internal/assertions/
// error_go126.go), and the unguarded ErrorAs in the same domain.
// NOTE: this requires running codegen on a toolchain >= go1.26 (latest stable).
func TestBuildConstraintDetection(t *testing.T) {
	s := New()

	pkg, err := s.Scan()
	if err != nil {
		t.Fatalf("scan: %v", err)
	}

	var guarded, unguarded int
	for _, fn := range pkg.Functions {
		switch fn.Name {
		case "ErrorAsType":
			if fn.GoBuild != "go1.26" {
				t.Errorf("ErrorAsType: expected GoBuild %q, got %q", "go1.26", fn.GoBuild)
			}
			guarded++
		case "ErrorAs":
			if fn.GoBuild != "" {
				t.Errorf("ErrorAs: expected no GoBuild, got %q", fn.GoBuild)
			}
			unguarded++
		}
	}

	if guarded == 0 {
		if os.Getenv("CI") == "" {
			// local tests spew error
			t.Error("did not observe the guarded ErrorAsType function; is codegen running on go1.26+?")
		} else {
			// expected error when CI is running with GOTOOLCHAIN=local on go1.25
			t.Log("WARN: did not observe the guarded ErrorAsType function; is codegen running on go1.26+?")
		}
	}
	if unguarded == 0 {
		t.Error("did not observe the unguarded ErrorAs function")
	}
}
