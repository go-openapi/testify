// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"testing"

	"pgregory.net/rapid"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// TestSdump_EdgeCases focuses on known problematic patterns.
func TestSdump_EdgeCases(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		value := edgeCaseGenerator().Draw(rt, "edge-case-generator")
		func() {
			defer func() {
				if r := recover(); r != nil {
					rt.Fatalf("Dump panicked on edge case: %v\nValue type: %T", r, value)
				}
			}()

			_ = spew.Sdump(value)
		}()
	})
}
