// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"os"
	"testing"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// TestScanner exercises the [Scanner] without asserting anything else than running without error.
//
// You may use this smoke test to verify the content of the constructed model.
func TestScanner(t *testing.T) {
	s := New()

	r, err := s.Scan()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if isCI := os.Getenv("CI"); isCI != "" {
		// skip output when run on CI
		return
	}
	spew.Config = spew.ConfigState{
		DisableMethods: true,
	}

	spew.Dump(r)
}
