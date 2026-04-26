// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"cmp"
	"context"
	"regexp"
	"time"
)

// Type constraint definitions for generic variants of assertions.
type (
	// Boolean is a bool or any type that can be converted to a bool.
	Boolean interface {
		~bool
	}

	// Text is any type of underlying type string or []byte.
	//
	// This is used by [RegexpT], [NotRegexpT], [JSONEqT], and [YAMLEqT].
	//
	// NOTE: unfortunately, []rune is not supported.
	Text interface {
		~string | ~[]byte
	}

	// RText extends [Text] by supporting dynamic construction of the
	// expected or actual value, e.g. "redact" functions.
	RText interface {
		Text | Redactor
	}

	// Redactor allows dynamic construction of expected or actual values, e.g. "redacting" values dynamically.
	//
	// This is used by json and yaml assertions.
	Redactor interface {
		func() string | func() []byte
	}

	// Ordered is a standard ordered type (i.e. types that support "<": [cmp.Ordered]) plus []byte and [time.Time].
	//
	// This is used by [GreaterT], [GreaterOrEqualT], [LessT], [LessOrEqualT], [IsIncreasingT], [IsDecreasingT].
	//
	// NOTE: since [time.Time] is a struct, custom types which redeclare [time.Time] are not supported.
	Ordered interface {
		cmp.Ordered | []byte | time.Time
	}

	// SignedNumeric is a signed integer or a floating point number or any type that can be converted to one of these.
	SignedNumeric interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~float32 | ~float64
	}

	// UnsignedNumeric is an unsigned integer.
	//
	// NOTE: there are no unsigned floating point numbers.
	UnsignedNumeric interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	}

	// Measurable is any number for which we can compute a delta (floats or integers).
	//
	// This is used by [InDeltaT] and [InEpsilonT].
	//
	// NOTE: unfortunately complex64 and complex128 are not supported.
	Measurable interface {
		SignedNumeric | UnsignedNumeric | ~float32 | ~float64
	}

	// RegExp is either a text containing a regular expression to compile (string or []byte), or directly the compiled regexp.
	//
	// This is used by [RegexpT] and [NotRegexpT].
	RegExp interface {
		Text | *regexp.Regexp
	}

	// Conditioner is a function used in asynchronous condition assertions.
	//
	// This type constraint allows for "overloaded" versions of the condition assertions ([Eventually], [Consistently]).
	//
	// The [WithSynctest] and [WithSynctestContext] wrappers opt a call into
	// fake-time polling via [testing/synctest]. See [WithSynctest] for details.
	Conditioner interface {
		func() bool | func(context.Context) error | WithSynctest | WithSynctestContext
	}

	// NeverConditioner is a function used by [Never].
	//
	// Unlike [Conditioner], [Never] does not accept the context-returning-error
	// form to avoid the double-negation confusion ("never returns no error").
	//
	// The [WithSynctest] wrapper opts a call into fake-time polling.
	NeverConditioner interface {
		func() bool | WithSynctest
	}

	// CollectibleConditioner is a function used in asynchronous condition assertions that use [CollectT].
	//
	// This type constraint allows for "overloaded" versions of the condition assertions ([EventuallyWith]).
	//
	// The [WithSynctestCollect] and [WithSynctestCollectContext] wrappers opt a
	// call into fake-time polling. See [WithSynctest] for details.
	CollectibleConditioner interface {
		func(*CollectT) | func(context.Context, *CollectT) |
			WithSynctestCollect | WithSynctestCollectContext
	}

	// WithSynctest wraps a [func() bool] condition to run [Eventually] /
	// [Never] / [Consistently] polling inside a [testing/synctest] bubble,
	// so `time.Ticker`, `time.After`, and `context.WithTimeout` use a fake
	// clock. Activation requires the caller to pass a real `*testing.T`;
	// with mocks or other [T] implementations, the wrapper falls back to
	// real-time polling.
	//
	// # When to use
	//
	// Use when the condition is pure compute, relies on `time.Sleep`, or
	// coordinates via channels created inside the condition. Fake time
	// eliminates timing-induced flakiness and enables deterministic tick
	// counts.
	//
	// # When not to use
	//
	// Do NOT use when the condition performs real I/O (network, filesystem,
	// syscalls): those block goroutines non-durably, so the fake clock
	// stalls and the timeout may not fire. Also do NOT use inside a test
	// that is already running in a [synctest.Test] bubble — nested bubbles
	// are forbidden and will panic.
	//
	// # Shared state
	//
	// The condition may read and write variables captured from the enclosing
	// scope; condition execution is serialized by design (see [Eventually]'s
	// Concurrency section). Avoid sharing channels or mutexes with goroutines
	// outside the bubble, as this will stall the fake clock.
	WithSynctest func() bool

	// WithSynctestContext is the [func(context.Context) error] counterpart
	// of [WithSynctest]. See [WithSynctest] for details.
	WithSynctestContext func(context.Context) error

	// WithSynctestCollect is the [func(*CollectT)] counterpart of
	// [WithSynctest] for use with [EventuallyWith]. See [WithSynctest] for details.
	WithSynctestCollect func(*CollectT)

	// WithSynctestCollectContext is the [func(context.Context, *CollectT)]
	// counterpart of [WithSynctest] for use with [EventuallyWith]. See
	// [WithSynctest] for details.
	WithSynctestCollectContext func(context.Context, *CollectT)
)
