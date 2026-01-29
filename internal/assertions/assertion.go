// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "context"

var (
	_ T              = &Assertions{}
	_ H              = &Assertions{}
	_ failNower      = &Assertions{}
	_ namer          = &Assertions{}
	_ contextualizer = &Assertions{}
)

// Assertions is the base type to provide assertion methods around the [T] interface.
//
// It implements [T] and exposes the other useful methods from [testing.T]:
// [testing.T.Helper], [testing.T.FailNow] and [testing.T.Context].
//
// These methods are optional but useful in this package.
type Assertions struct {
	t T
}

func New(t T) *Assertions {
	return &Assertions{
		t: t,
	}
}

func (a *Assertions) Errorf(format string, args ...any) {
	a.t.Errorf(format, args...)
}

func (a *Assertions) Helper() {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
}

func (a *Assertions) FailNow() {
	if f, ok := a.t.(failNower); ok {
		f.FailNow()
	}
}

func (a *Assertions) Name() string {
	if n, ok := a.t.(namer); ok {
		return n.Name()
	}

	return ""
}

func (a *Assertions) Context() context.Context {
	if c, ok := a.t.(contextualizer); ok {
		return c.Context()
	}

	return context.Background()
}
