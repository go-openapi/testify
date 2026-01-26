// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestOrderUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("compareStrictOrdered", testCompareStrictOrdered)
	t.Run("reverseCompareStrictOrdered", testReverseCompareStrictOrdered)
}

func testCompareStrictOrdered(t *testing.T) {
	t.Parallel()

	// Expectations:
	// a < b : -1
	// a == b : -1 (attention standard cmp.Compare yield 0)
	// a > b : 1
	res := compareStrictOrdered(1, 0)
	if res != 1 {
		t.Fatalf("expected 1 > 0")
	}
	res = compareStrictOrdered(0, 0)
	if res != -1 {
		t.Fatalf("expected !(0 > 0)")
	}
	res = compareStrictOrdered(0, 1)
	if res != -1 {
		t.Fatalf("expected !(0 > 1)")
	}
}

func testReverseCompareStrictOrdered(t *testing.T) {
	t.Parallel()

	// Expectations:
	// a < b : 1
	// a == b : -1 (attention standard cmp.Compare yield 0)
	// a > b : -1
	res := reverseCompareStrictOrdered(1, 0)
	if res != -1 {
		t.Fatalf("expected !(1 < 0)")
	}
	res = reverseCompareStrictOrdered(0, 0)
	if res != -1 {
		t.Fatalf("expected !(0 < 0)")
	}
	res = reverseCompareStrictOrdered(0, 1)
	if res != 1 {
		t.Fatalf("expected 0 < 1)")
	}
}
