// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"iter"
	"slices"
	"testing"
)

type trackAPIUsageCase struct {
	name     string
	calls    []string
	expected map[string]int
}

func trackAPIUsageCases() iter.Seq[trackAPIUsageCase] {
	return slices.Values([]trackAPIUsageCase{
		{
			name:  "single call",
			calls: []string{"assert.Equal"},
			expected: map[string]int{
				"assert.Equal": 1,
			},
		},
		{
			name:  "multiple calls same function",
			calls: []string{"assert.Equal", "assert.Equal", "assert.Equal"},
			expected: map[string]int{
				"assert.Equal": 3,
			},
		},
		{
			name:  "multiple different functions",
			calls: []string{"assert.Equal", "require.NoError", "assert.True"},
			expected: map[string]int{
				"assert.Equal":    1,
				"require.NoError": 1,
				"assert.True":     1,
			},
		},
	})
}

func TestTrackAPIUsage(t *testing.T) {
	t.Parallel()

	for c := range trackAPIUsageCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			rpt := &report{}
			for _, call := range c.calls {
				rpt.trackAPIUsage(call)
			}

			if len(rpt.apiUsage) != len(c.expected) {
				t.Errorf("expected %d entries, got %d", len(c.expected), len(rpt.apiUsage))
			}
			for k, v := range c.expected {
				if rpt.apiUsage[k] != v {
					t.Errorf("expected %s=%d, got %d", k, v, rpt.apiUsage[k])
				}
			}
		})
	}
}

type trackUpgradeCase struct {
	name     string
	upgrades [][2]string // from, to pairs
	expected map[string]int
}

func trackUpgradeCases() iter.Seq[trackUpgradeCase] {
	return slices.Values([]trackUpgradeCase{
		{
			name:     "single upgrade",
			upgrades: [][2]string{{"Equal", "EqualT"}},
			expected: map[string]int{
				"Equal → EqualT": 1,
			},
		},
		{
			name: "multiple upgrades",
			upgrades: [][2]string{
				{"Equal", "EqualT"},
				{"Equal", "EqualT"},
				{"True", "TrueT"},
			},
			expected: map[string]int{
				"Equal → EqualT": 2,
				"True → TrueT":   1,
			},
		},
	})
}

func TestTrackUpgrade(t *testing.T) {
	t.Parallel()

	for c := range trackUpgradeCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			rpt := &report{}
			for _, u := range c.upgrades {
				rpt.trackUpgrade(u[0], u[1])
			}

			if len(rpt.upgraded) != len(c.expected) {
				t.Errorf("expected %d entries, got %d", len(c.expected), len(rpt.upgraded))
			}
			for k, v := range c.expected {
				if rpt.upgraded[k] != v {
					t.Errorf("expected %s=%d, got %d", k, v, rpt.upgraded[k])
				}
			}
		})
	}
}

type trackSkipCase struct {
	name            string
	skips           []skipReason
	expectedSkipped map[string]int
	verbose         bool
	expectDiag      bool
}

func trackSkipCases() iter.Seq[trackSkipCase] {
	return slices.Values([]trackSkipCase{
		{
			name:  "single skip",
			skips: []skipReason{skipPointerSemantics},
			expectedSkipped: map[string]int{
				string(skipPointerSemantics): 1,
			},
			verbose:    false,
			expectDiag: false,
		},
		{
			name:  "verbose emits diagnostic",
			skips: []skipReason{skipInterfaceType},
			expectedSkipped: map[string]int{
				string(skipInterfaceType): 1,
			},
			verbose:    true,
			expectDiag: true,
		},
		{
			name:  "multiple skips aggregated",
			skips: []skipReason{skipPointerSemantics, skipPointerSemantics, skipTypeMismatch},
			expectedSkipped: map[string]int{
				string(skipPointerSemantics): 2,
				string(skipTypeMismatch):     1,
			},
			verbose:    false,
			expectDiag: false,
		},
	})
}

func TestTrackSkip(t *testing.T) {
	t.Parallel()

	for c := range trackSkipCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			rpt := &report{}
			for _, reason := range c.skips {
				rpt.trackSkip("test.go", 42, "Equal", reason, c.verbose, "*float64")
			}

			if len(rpt.skipped) != len(c.expectedSkipped) {
				t.Errorf("expected %d skip entries, got %d", len(c.expectedSkipped), len(rpt.skipped))
			}
			for k, v := range c.expectedSkipped {
				if rpt.skipped[k] != v {
					t.Errorf("expected %s=%d, got %d", k, v, rpt.skipped[k])
				}
			}

			hasDiag := len(rpt.diagnostics) > 0
			if c.expectDiag && !hasDiag {
				t.Error("expected diagnostic to be emitted in verbose mode")
			}
			if !c.expectDiag && hasDiag {
				t.Errorf("expected no diagnostic in non-verbose mode, got %d", len(rpt.diagnostics))
			}
		})
	}
}

func TestReportInitializesLazily(t *testing.T) {
	t.Parallel()

	rpt := &report{}

	// All maps should be nil initially.
	if rpt.apiUsage != nil {
		t.Error("apiUsage should be nil initially")
	}
	if rpt.upgraded != nil {
		t.Error("upgraded should be nil initially")
	}
	if rpt.skipped != nil {
		t.Error("skipped should be nil initially")
	}

	// After tracking, maps should be initialized.
	rpt.trackAPIUsage("assert.Equal")
	rpt.trackUpgrade("Equal", "EqualT")
	rpt.trackSkip("test.go", 1, "Equal", skipPointerSemantics, false, "")

	if rpt.apiUsage == nil {
		t.Error("apiUsage should be initialized after tracking")
	}
	if rpt.upgraded == nil {
		t.Error("upgraded should be initialized after tracking")
	}
	if rpt.skipped == nil {
		t.Error("skipped should be initialized after tracking")
	}
}
