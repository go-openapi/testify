// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"pgregory.net/rapid"
)

func TestMain(m *testing.M) {
	// set flags for rapid:
	// "rapid.checks", 100, "rapid: number of checks to perform"
	// "rapid.steps", 30, "rapid: average number of Repeat actions to execute"
	// "rapid.failfile", "", "rapid: fail file to use to reproduce test failure"
	// "rapid.nofailfile", false, "rapid: do not write fail files on test failures"
	// "rapid.seed", 0, "rapid: PRNG seed to start with (0 to use a random one)"
	// "rapid.log", false, "rapid: eager verbose output to stdout (to aid with unrecoverable test failures)"
	// "rapid.v", false, "rapid: verbose output"
	// "rapid.debug", false, "rapid: debugging output"
	// "rapid.debugvis", false, "rapid: debugging visualization"
	// "rapid.shrinktime", 30*time.Second, "rapid: maximum time to spend on test case minimization"
	os.Args = append(os.Args, "-rapid.checks", strconv.Itoa(testLoad()))
	flag.Parse()

	os.Exit(m.Run())
}

// TestSdump uses property-based testing to ensure Dump never panics or hangs
// with arbitrary Go values, including edge cases that historically caused issues.
func TestSdump(t *testing.T) {
	t.Parallel()

	rapid.Check(t, NoPanicProp(t.Context(), Generator()))
}
