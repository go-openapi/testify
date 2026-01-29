// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestAssertionNew(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	a := New(mock)
	if a == nil {
		FailNow(t, "New should never return nil")

		return
	}
	if a.t == nil {
		FailNow(t, "assertion should contain a T")
	}

	// TODO: check methods
}
