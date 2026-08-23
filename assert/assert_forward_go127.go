// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.27

package assert

import (
	"iter"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// BlockedT is the same as [BlockedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) BlockedT[E any, CHAN ~chan E](ch CHAN, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.BlockedT[E, CHAN](a.T, ch, msgAndArgs...)
}

// BlockedTf is the same as [Assertions.BlockedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) BlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.BlockedT[E, CHAN](a.T, ch, forwardArgs(msg, args)...)
}

// Consistently is the same as [Consistently], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Consistently[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Consistently[C](a.T, condition, timeout, tick, msgAndArgs...)
}

// Consistentlyf is the same as [Assertions.Consistently], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Consistentlyf[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Consistently[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...)
}

// ElementsMatchT is the same as [ElementsMatchT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatchT[E comparable](listA []E, listB []E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatchT[E](a.T, listA, listB, msgAndArgs...)
}

// ElementsMatchTf is the same as [Assertions.ElementsMatchT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatchTf[E comparable](listA []E, listB []E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatchT[E](a.T, listA, listB, forwardArgs(msg, args)...)
}

// EqualT is the same as [EqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualT[V comparable](expected V, actual V, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualT[V](a.T, expected, actual, msgAndArgs...)
}

// EqualTf is the same as [Assertions.EqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualTf[V comparable](expected V, actual V, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualT[V](a.T, expected, actual, forwardArgs(msg, args)...)
}

// ErrorAsType is the same as [ErrorAsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAsType[E error](err error, target *E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAsType[E](a.T, err, target, msgAndArgs...)
}

// ErrorAsTypef is the same as [Assertions.ErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAsTypef[E error](err error, target *E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAsType[E](a.T, err, target, forwardArgs(msg, args)...)
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventually[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Eventually[C](a.T, condition, timeout, tick, msgAndArgs...)
}

// Eventuallyf is the same as [Assertions.Eventually], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventuallyf[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Eventually[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...)
}

// EventuallyWith is the same as [EventuallyWith], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWith[C CollectibleConditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWith[C](a.T, condition, timeout, tick, msgAndArgs...)
}

// EventuallyWithf is the same as [Assertions.EventuallyWith], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWithf[C CollectibleConditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWith[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...)
}

// FalseT is the same as [FalseT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FalseT[B Boolean](value B, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FalseT[B](a.T, value, msgAndArgs...)
}

// FalseTf is the same as [Assertions.FalseT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FalseTf[B Boolean](value B, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FalseT[B](a.T, value, forwardArgs(msg, args)...)
}

// GreaterOrEqualT is the same as [GreaterOrEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqualT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqualT[Orderable](a.T, e1, e2, msgAndArgs...)
}

// GreaterOrEqualTf is the same as [Assertions.GreaterOrEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqualTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqualT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...)
}

// GreaterT is the same as [GreaterT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterT[Orderable](a.T, e1, e2, msgAndArgs...)
}

// GreaterTf is the same as [Assertions.GreaterT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...)
}

// InDeltaT is the same as [InDeltaT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaT[Number Measurable](expected Number, actual Number, delta Number, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaT[Number](a.T, expected, actual, delta, msgAndArgs...)
}

// InDeltaTf is the same as [Assertions.InDeltaT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaTf[Number Measurable](expected Number, actual Number, delta Number, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaT[Number](a.T, expected, actual, delta, forwardArgs(msg, args)...)
}

// InEpsilonSymmetricT is the same as [InEpsilonSymmetricT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSymmetricT[Number Measurable](x Number, y Number, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSymmetricT[Number](a.T, x, y, epsilon, msgAndArgs...)
}

// InEpsilonSymmetricTf is the same as [Assertions.InEpsilonSymmetricT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSymmetricTf[Number Measurable](x Number, y Number, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSymmetricT[Number](a.T, x, y, epsilon, forwardArgs(msg, args)...)
}

// InEpsilonT is the same as [InEpsilonT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonT[Number Measurable](expected Number, actual Number, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonT[Number](a.T, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonTf is the same as [Assertions.InEpsilonT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonTf[Number Measurable](expected Number, actual Number, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonT[Number](a.T, expected, actual, epsilon, forwardArgs(msg, args)...)
}

// IsDecreasingT is the same as [IsDecreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// IsDecreasingTf is the same as [Assertions.IsDecreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// IsIncreasingT is the same as [IsIncreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// IsIncreasingTf is the same as [Assertions.IsIncreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// IsNonDecreasingT is the same as [IsNonDecreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// IsNonDecreasingTf is the same as [Assertions.IsNonDecreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// IsNonIncreasingT is the same as [IsNonIncreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// IsNonIncreasingTf is the same as [Assertions.IsNonIncreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// IsNotOfTypeT is the same as [IsNotOfTypeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotOfTypeT[EType any](object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotOfTypeT[EType](a.T, object, msgAndArgs...)
}

// IsNotOfTypeTf is the same as [Assertions.IsNotOfTypeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotOfTypeTf[EType any](object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotOfTypeT[EType](a.T, object, forwardArgs(msg, args)...)
}

// IsOfTypeT is the same as [IsOfTypeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsOfTypeT[EType any](object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsOfTypeT[EType](a.T, object, msgAndArgs...)
}

// IsOfTypeTf is the same as [Assertions.IsOfTypeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsOfTypeTf[EType any](object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsOfTypeT[EType](a.T, object, forwardArgs(msg, args)...)
}

// JSONEqT is the same as [JSONEqT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqT[EDoc, ADoc RText](expected EDoc, actual ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqT[EDoc, ADoc](a.T, expected, actual, msgAndArgs...)
}

// JSONEqTf is the same as [Assertions.JSONEqT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqTf[EDoc, ADoc RText](expected EDoc, actual ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqT[EDoc, ADoc](a.T, expected, actual, forwardArgs(msg, args)...)
}

// JSONMarshalAsT is the same as [JSONMarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONMarshalAsT[EDoc RText](expected EDoc, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONMarshalAsT[EDoc](a.T, expected, object, msgAndArgs...)
}

// JSONMarshalAsTf is the same as [Assertions.JSONMarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONMarshalAsTf[EDoc RText](expected EDoc, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONMarshalAsT[EDoc](a.T, expected, object, forwardArgs(msg, args)...)
}

// JSONUnmarshalAsT is the same as [JSONUnmarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONUnmarshalAsT[Object any, ADoc RText](expected Object, jazon ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONUnmarshalAsT[Object, ADoc](a.T, expected, jazon, msgAndArgs...)
}

// JSONUnmarshalAsTf is the same as [Assertions.JSONUnmarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONUnmarshalAsTf[Object any, ADoc RText](expected Object, jazon ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONUnmarshalAsT[Object, ADoc](a.T, expected, jazon, forwardArgs(msg, args)...)
}

// LessOrEqualT is the same as [LessOrEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqualT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqualT[Orderable](a.T, e1, e2, msgAndArgs...)
}

// LessOrEqualTf is the same as [Assertions.LessOrEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqualTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqualT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...)
}

// LessT is the same as [LessT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessT[Orderable](a.T, e1, e2, msgAndArgs...)
}

// LessTf is the same as [Assertions.LessT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...)
}

// MapContainsT is the same as [MapContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapContainsT[Map ~map[K]V, K comparable, V any](m Map, key K, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapContainsT[Map, K, V](a.T, m, key, msgAndArgs...)
}

// MapContainsTf is the same as [Assertions.MapContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapContainsTf[Map ~map[K]V, K comparable, V any](m Map, key K, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapContainsT[Map, K, V](a.T, m, key, forwardArgs(msg, args)...)
}

// MapEqualT is the same as [MapEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapEqualT[K, V comparable](listA map[K]V, listB map[K]V, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapEqualT[K, V](a.T, listA, listB, msgAndArgs...)
}

// MapEqualTf is the same as [Assertions.MapEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapEqualTf[K, V comparable](listA map[K]V, listB map[K]V, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapEqualT[K, V](a.T, listA, listB, forwardArgs(msg, args)...)
}

// MapNotContainsT is the same as [MapNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapNotContainsT[Map ~map[K]V, K comparable, V any](m Map, key K, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapNotContainsT[Map, K, V](a.T, m, key, msgAndArgs...)
}

// MapNotContainsTf is the same as [Assertions.MapNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapNotContainsTf[Map ~map[K]V, K comparable, V any](m Map, key K, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapNotContainsT[Map, K, V](a.T, m, key, forwardArgs(msg, args)...)
}

// MapNotEqualT is the same as [MapNotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapNotEqualT[K, V comparable](listA map[K]V, listB map[K]V, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapNotEqualT[K, V](a.T, listA, listB, msgAndArgs...)
}

// MapNotEqualTf is the same as [Assertions.MapNotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) MapNotEqualTf[K, V comparable](listA map[K]V, listB map[K]V, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.MapNotEqualT[K, V](a.T, listA, listB, forwardArgs(msg, args)...)
}

// NegativeT is the same as [NegativeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NegativeT[SignedNumber SignedNumeric](e SignedNumber, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NegativeT[SignedNumber](a.T, e, msgAndArgs...)
}

// NegativeTf is the same as [Assertions.NegativeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NegativeTf[SignedNumber SignedNumeric](e SignedNumber, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NegativeT[SignedNumber](a.T, e, forwardArgs(msg, args)...)
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Never[C NeverConditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Never[C](a.T, condition, timeout, tick, msgAndArgs...)
}

// Neverf is the same as [Assertions.Never], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Neverf[C NeverConditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Never[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...)
}

// NotBlockedT is the same as [NotBlockedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotBlockedT[E any, CHAN ~chan E](ch CHAN, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotBlockedT[E, CHAN](a.T, ch, msgAndArgs...)
}

// NotBlockedTf is the same as [Assertions.NotBlockedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotBlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotBlockedT[E, CHAN](a.T, ch, forwardArgs(msg, args)...)
}

// NotElementsMatchT is the same as [NotElementsMatchT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatchT[E comparable](listA []E, listB []E, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatchT[E](a.T, listA, listB, msgAndArgs...)
}

// NotElementsMatchTf is the same as [Assertions.NotElementsMatchT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatchTf[E comparable](listA []E, listB []E, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatchT[E](a.T, listA, listB, forwardArgs(msg, args)...)
}

// NotEqualT is the same as [NotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualT[V comparable](expected V, actual V, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualT[V](a.T, expected, actual, msgAndArgs...)
}

// NotEqualTf is the same as [Assertions.NotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualTf[V comparable](expected V, actual V, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualT[V](a.T, expected, actual, forwardArgs(msg, args)...)
}

// NotErrorAsType is the same as [NotErrorAsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAsType[E error](err error, target *E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAsType[E](a.T, err, target, msgAndArgs...)
}

// NotErrorAsTypef is the same as [Assertions.NotErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAsTypef[E error](err error, target *E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAsType[E](a.T, err, target, forwardArgs(msg, args)...)
}

// NotRegexpT is the same as [NotRegexpT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexpT[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexpT[Rex, ADoc](a.T, rx, actual, msgAndArgs...)
}

// NotRegexpTf is the same as [Assertions.NotRegexpT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexpTf[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexpT[Rex, ADoc](a.T, rx, actual, forwardArgs(msg, args)...)
}

// NotSameT is the same as [NotSameT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSameT[P any](expected *P, actual *P, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSameT[P](a.T, expected, actual, msgAndArgs...)
}

// NotSameTf is the same as [Assertions.NotSameT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSameTf[P any](expected *P, actual *P, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSameT[P](a.T, expected, actual, forwardArgs(msg, args)...)
}

// NotSortedT is the same as [NotSortedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSortedT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSortedT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// NotSortedTf is the same as [Assertions.NotSortedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSortedTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSortedT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// PositiveT is the same as [PositiveT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PositiveT[SignedNumber SignedNumeric](e SignedNumber, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PositiveT[SignedNumber](a.T, e, msgAndArgs...)
}

// PositiveTf is the same as [Assertions.PositiveT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PositiveTf[SignedNumber SignedNumeric](e SignedNumber, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PositiveT[SignedNumber](a.T, e, forwardArgs(msg, args)...)
}

// RegexpT is the same as [RegexpT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) RegexpT[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.RegexpT[Rex, ADoc](a.T, rx, actual, msgAndArgs...)
}

// RegexpTf is the same as [Assertions.RegexpT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) RegexpTf[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.RegexpT[Rex, ADoc](a.T, rx, actual, forwardArgs(msg, args)...)
}

// SameT is the same as [SameT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SameT[P any](expected *P, actual *P, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SameT[P](a.T, expected, actual, msgAndArgs...)
}

// SameTf is the same as [Assertions.SameT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SameTf[P any](expected *P, actual *P, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SameT[P](a.T, expected, actual, forwardArgs(msg, args)...)
}

// SeqContainsT is the same as [SeqContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SeqContainsT[E comparable](iter iter.Seq[E], element E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SeqContainsT[E](a.T, iter, element, msgAndArgs...)
}

// SeqContainsTf is the same as [Assertions.SeqContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SeqContainsTf[E comparable](iter iter.Seq[E], element E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SeqContainsT[E](a.T, iter, element, forwardArgs(msg, args)...)
}

// SeqNotContainsT is the same as [SeqNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SeqNotContainsT[E comparable](iter iter.Seq[E], element E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SeqNotContainsT[E](a.T, iter, element, msgAndArgs...)
}

// SeqNotContainsTf is the same as [Assertions.SeqNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SeqNotContainsTf[E comparable](iter iter.Seq[E], element E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SeqNotContainsT[E](a.T, iter, element, forwardArgs(msg, args)...)
}

// SliceContainsT is the same as [SliceContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceContainsT[Slice ~[]E, E comparable](s Slice, element E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceContainsT[Slice, E](a.T, s, element, msgAndArgs...)
}

// SliceContainsTf is the same as [Assertions.SliceContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceContainsTf[Slice ~[]E, E comparable](s Slice, element E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceContainsT[Slice, E](a.T, s, element, forwardArgs(msg, args)...)
}

// SliceEqualT is the same as [SliceEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceEqualT[E comparable](listA []E, listB []E, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceEqualT[E](a.T, listA, listB, msgAndArgs...)
}

// SliceEqualTf is the same as [Assertions.SliceEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceEqualTf[E comparable](listA []E, listB []E, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceEqualT[E](a.T, listA, listB, forwardArgs(msg, args)...)
}

// SliceNotContainsT is the same as [SliceNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotContainsT[Slice ~[]E, E comparable](s Slice, element E, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotContainsT[Slice, E](a.T, s, element, msgAndArgs...)
}

// SliceNotContainsTf is the same as [Assertions.SliceNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotContainsTf[Slice ~[]E, E comparable](s Slice, element E, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotContainsT[Slice, E](a.T, s, element, forwardArgs(msg, args)...)
}

// SliceNotEqualT is the same as [SliceNotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotEqualT[E comparable](listA []E, listB []E, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotEqualT[E](a.T, listA, listB, msgAndArgs...)
}

// SliceNotEqualTf is the same as [Assertions.SliceNotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotEqualTf[E comparable](listA []E, listB []E, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotEqualT[E](a.T, listA, listB, forwardArgs(msg, args)...)
}

// SliceNotSubsetT is the same as [SliceNotSubsetT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotSubsetT[Slice ~[]E, E comparable](list Slice, subset Slice, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotSubsetT[Slice, E](a.T, list, subset, msgAndArgs...)
}

// SliceNotSubsetTf is the same as [Assertions.SliceNotSubsetT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceNotSubsetTf[Slice ~[]E, E comparable](list Slice, subset Slice, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotSubsetT[Slice, E](a.T, list, subset, forwardArgs(msg, args)...)
}

// SliceSubsetT is the same as [SliceSubsetT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceSubsetT[Slice ~[]E, E comparable](list Slice, subset Slice, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceSubsetT[Slice, E](a.T, list, subset, msgAndArgs...)
}

// SliceSubsetTf is the same as [Assertions.SliceSubsetT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SliceSubsetTf[Slice ~[]E, E comparable](list Slice, subset Slice, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SliceSubsetT[Slice, E](a.T, list, subset, forwardArgs(msg, args)...)
}

// SortedT is the same as [SortedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SortedT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SortedT[OrderedSlice, E](a.T, collection, msgAndArgs...)
}

// SortedTf is the same as [Assertions.SortedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) SortedTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.SortedT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...)
}

// StringContainsT is the same as [StringContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) StringContainsT[ADoc, EDoc Text](str ADoc, substring EDoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.StringContainsT[ADoc, EDoc](a.T, str, substring, msgAndArgs...)
}

// StringContainsTf is the same as [Assertions.StringContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) StringContainsTf[ADoc, EDoc Text](str ADoc, substring EDoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.StringContainsT[ADoc, EDoc](a.T, str, substring, forwardArgs(msg, args)...)
}

// StringNotContainsT is the same as [StringNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) StringNotContainsT[ADoc, EDoc Text](str ADoc, substring EDoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.StringNotContainsT[ADoc, EDoc](a.T, str, substring, msgAndArgs...)
}

// StringNotContainsTf is the same as [Assertions.StringNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) StringNotContainsTf[ADoc, EDoc Text](str ADoc, substring EDoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.StringNotContainsT[ADoc, EDoc](a.T, str, substring, forwardArgs(msg, args)...)
}

// TrueT is the same as [TrueT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) TrueT[B Boolean](value B, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.TrueT[B](a.T, value, msgAndArgs...)
}

// TrueTf is the same as [Assertions.TrueT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) TrueTf[B Boolean](value B, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.TrueT[B](a.T, value, forwardArgs(msg, args)...)
}

// YAMLEqT is the same as [YAMLEqT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqT[EDoc, ADoc RText](expected EDoc, actual ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqT[EDoc, ADoc](a.T, expected, actual, msgAndArgs...)
}

// YAMLEqTf is the same as [Assertions.YAMLEqT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqTf[EDoc, ADoc RText](expected EDoc, actual ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqT[EDoc, ADoc](a.T, expected, actual, forwardArgs(msg, args)...)
}

// YAMLMarshalAsT is the same as [YAMLMarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLMarshalAsT[EDoc RText](expected EDoc, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLMarshalAsT[EDoc](a.T, expected, object, msgAndArgs...)
}

// YAMLMarshalAsTf is the same as [Assertions.YAMLMarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLMarshalAsTf[EDoc RText](expected EDoc, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLMarshalAsT[EDoc](a.T, expected, object, forwardArgs(msg, args)...)
}

// YAMLUnmarshalAsT is the same as [YAMLUnmarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLUnmarshalAsT[Object any, ADoc RText](expected Object, yamlDoc ADoc, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLUnmarshalAsT[Object, ADoc](a.T, expected, yamlDoc, msgAndArgs...)
}

// YAMLUnmarshalAsTf is the same as [Assertions.YAMLUnmarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLUnmarshalAsTf[Object any, ADoc RText](expected Object, yamlDoc ADoc, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLUnmarshalAsT[Object, ADoc](a.T, expected, yamlDoc, forwardArgs(msg, args)...)
}
