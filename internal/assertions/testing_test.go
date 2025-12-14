// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestTestingFailNowWithPlainT(t *testing.T) {
	t.Parallel()
	mock := &mockT{}

	Panics(t, func() {
		FailNow(mock, "failed")
	}, "should panic since mockT is missing FailNow()")
}

func TestTestingFailNowWithFullT(t *testing.T) {
	t.Parallel()
	mock := &mockFailNowT{}

	NotPanics(t, func() {
		FailNow(mock, "failed")
	}, "should call mockT.FailNow() rather than panicking")
}
