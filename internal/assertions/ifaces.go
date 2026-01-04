// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "context"

// T is an interface wrapper around [testing.T].
type T interface {
	Errorf(format string, args ...any)
}

// H is an interface for types that implement the Helper method.
// This allows marking functions as test helpers.
type H interface {
	Helper()
}

type failNower interface {
	FailNow()
}

type namer interface {
	Name() string
}

type contextualizer interface {
	Context() context.Context
}
