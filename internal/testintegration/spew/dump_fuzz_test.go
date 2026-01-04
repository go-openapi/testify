// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"testing"

	"pgregory.net/rapid"
)

// FuzzSdump is the fuzzed equivalent of TestSdump.
//
// Given a high number of values of different types generated randomly,
// the fuzz engine will alter these values and run [spew.Sdump].
//
// # Limitations
//
// FuzzSdump skips circular map cases because Go's fmt package
// (used by rapid's fuzz logging) cannot handle circular references.
//
// This case is covered by rapid.Check in TestSdump (which doesn't log).
func FuzzSdump(f *testing.F) {
	prop := NoPanicProp(f.Context(), Generator(WithSkipCircularMap(true)))
	f.Fuzz(rapid.MakeFuzz(prop))
}
