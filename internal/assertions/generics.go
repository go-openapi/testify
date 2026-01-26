// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"cmp"
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
)
