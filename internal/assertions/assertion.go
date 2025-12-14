// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

// Assertions provides assertion methods around the [T] interface.
type Assertions struct {
	t T
}

// New makes a new [Assertions] object for the specified [T].
func New(t T) *Assertions {
	return &Assertions{
		t: t,
	}
}
