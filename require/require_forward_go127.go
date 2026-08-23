// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.27

package require

import (
	"iter"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// BlockedT is the same as [BlockedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) BlockedT[E any, CHAN ~chan E](ch CHAN, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.BlockedT[E, CHAN](a.T, ch, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// BlockedTf is the same as [Assertions.BlockedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) BlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.BlockedT[E, CHAN](a.T, ch, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// Consistently is the same as [Consistently], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Consistently[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Consistently[C](a.T, condition, timeout, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Consistentlyf is the same as [Assertions.Consistently], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Consistentlyf[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Consistently[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// ElementsMatchT is the same as [ElementsMatchT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatchT[E comparable](listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatchT[E](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ElementsMatchTf is the same as [Assertions.ElementsMatchT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatchTf[E comparable](listA []E, listB []E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatchT[E](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// EqualT is the same as [EqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualT[V comparable](expected V, actual V, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualT[V](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EqualTf is the same as [Assertions.EqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualTf[V comparable](expected V, actual V, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualT[V](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// ErrorAsType is the same as [ErrorAsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAsType[E error](err error, target *E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAsType[E](a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ErrorAsTypef is the same as [Assertions.ErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAsTypef[E error](err error, target *E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAsType[E](a.T, err, target, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventually[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Eventually[C](a.T, condition, timeout, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Eventuallyf is the same as [Assertions.Eventually], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventuallyf[C Conditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Eventually[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// EventuallyWith is the same as [EventuallyWith], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWith[C CollectibleConditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith[C](a.T, condition, timeout, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EventuallyWithf is the same as [Assertions.EventuallyWith], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWithf[C CollectibleConditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// FalseT is the same as [FalseT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FalseT[B Boolean](value B, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FalseT[B](a.T, value, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// FalseTf is the same as [Assertions.FalseT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FalseTf[B Boolean](value B, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FalseT[B](a.T, value, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// GreaterOrEqualT is the same as [GreaterOrEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqualT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqualT[Orderable](a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// GreaterOrEqualTf is the same as [Assertions.GreaterOrEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqualTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqualT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// GreaterT is the same as [GreaterT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterT[Orderable](a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// GreaterTf is the same as [Assertions.GreaterT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// InDeltaT is the same as [InDeltaT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaT[Number Measurable](expected Number, actual Number, delta Number, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaT[Number](a.T, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InDeltaTf is the same as [Assertions.InDeltaT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaTf[Number Measurable](expected Number, actual Number, delta Number, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaT[Number](a.T, expected, actual, delta, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonSymmetricT is the same as [InEpsilonSymmetricT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSymmetricT[Number Measurable](x Number, y Number, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSymmetricT[Number](a.T, x, y, epsilon, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonSymmetricTf is the same as [Assertions.InEpsilonSymmetricT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSymmetricTf[Number Measurable](x Number, y Number, epsilon float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSymmetricT[Number](a.T, x, y, epsilon, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonT is the same as [InEpsilonT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonT[Number Measurable](expected Number, actual Number, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonT[Number](a.T, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonTf is the same as [Assertions.InEpsilonT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonTf[Number Measurable](expected Number, actual Number, epsilon float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonT[Number](a.T, expected, actual, epsilon, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsDecreasingT is the same as [IsDecreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsDecreasingTf is the same as [Assertions.IsDecreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsIncreasingT is the same as [IsIncreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsIncreasingTf is the same as [Assertions.IsIncreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsNonDecreasingT is the same as [IsNonDecreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNonDecreasingTf is the same as [Assertions.IsNonDecreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsNonIncreasingT is the same as [IsNonIncreasingT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasingT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNonIncreasingTf is the same as [Assertions.IsNonIncreasingT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasingT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsNotOfTypeT is the same as [IsNotOfTypeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotOfTypeT[EType any](object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNotOfTypeT[EType](a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNotOfTypeTf is the same as [Assertions.IsNotOfTypeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotOfTypeTf[EType any](object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNotOfTypeT[EType](a.T, object, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// IsOfTypeT is the same as [IsOfTypeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsOfTypeT[EType any](object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsOfTypeT[EType](a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsOfTypeTf is the same as [Assertions.IsOfTypeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsOfTypeTf[EType any](object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsOfTypeT[EType](a.T, object, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// JSONEqT is the same as [JSONEqT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqT[EDoc, ADoc RText](expected EDoc, actual ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqT[EDoc, ADoc](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// JSONEqTf is the same as [Assertions.JSONEqT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqTf[EDoc, ADoc RText](expected EDoc, actual ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqT[EDoc, ADoc](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// JSONMarshalAsT is the same as [JSONMarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONMarshalAsT[EDoc RText](expected EDoc, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONMarshalAsT[EDoc](a.T, expected, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// JSONMarshalAsTf is the same as [Assertions.JSONMarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONMarshalAsTf[EDoc RText](expected EDoc, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONMarshalAsT[EDoc](a.T, expected, object, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// JSONUnmarshalAsT is the same as [JSONUnmarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONUnmarshalAsT[Object any, ADoc RText](expected Object, jazon ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONUnmarshalAsT[Object, ADoc](a.T, expected, jazon, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// JSONUnmarshalAsTf is the same as [Assertions.JSONUnmarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONUnmarshalAsTf[Object any, ADoc RText](expected Object, jazon ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONUnmarshalAsT[Object, ADoc](a.T, expected, jazon, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// LessOrEqualT is the same as [LessOrEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqualT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqualT[Orderable](a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// LessOrEqualTf is the same as [Assertions.LessOrEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqualTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqualT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// LessT is the same as [LessT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessT[Orderable Ordered](e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessT[Orderable](a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// LessTf is the same as [Assertions.LessT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessTf[Orderable Ordered](e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessT[Orderable](a.T, e1, e2, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// MapContainsT is the same as [MapContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapContainsT[Map ~map[K]V, K comparable, V any](m Map, key K, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapContainsT[Map, K, V](a.T, m, key, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// MapContainsTf is the same as [Assertions.MapContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapContainsTf[Map ~map[K]V, K comparable, V any](m Map, key K, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapContainsT[Map, K, V](a.T, m, key, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// MapEqualT is the same as [MapEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapEqualT[K, V comparable](listA map[K]V, listB map[K]V, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapEqualT[K, V](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// MapEqualTf is the same as [Assertions.MapEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapEqualTf[K, V comparable](listA map[K]V, listB map[K]V, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapEqualT[K, V](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// MapNotContainsT is the same as [MapNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapNotContainsT[Map ~map[K]V, K comparable, V any](m Map, key K, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapNotContainsT[Map, K, V](a.T, m, key, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// MapNotContainsTf is the same as [Assertions.MapNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapNotContainsTf[Map ~map[K]V, K comparable, V any](m Map, key K, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapNotContainsT[Map, K, V](a.T, m, key, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// MapNotEqualT is the same as [MapNotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapNotEqualT[K, V comparable](listA map[K]V, listB map[K]V, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapNotEqualT[K, V](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// MapNotEqualTf is the same as [Assertions.MapNotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) MapNotEqualTf[K, V comparable](listA map[K]V, listB map[K]V, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.MapNotEqualT[K, V](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NegativeT is the same as [NegativeT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NegativeT[SignedNumber SignedNumeric](e SignedNumber, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NegativeT[SignedNumber](a.T, e, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NegativeTf is the same as [Assertions.NegativeT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NegativeTf[SignedNumber SignedNumeric](e SignedNumber, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NegativeT[SignedNumber](a.T, e, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Never[C NeverConditioner](condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Never[C](a.T, condition, timeout, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Neverf is the same as [Assertions.Never], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Neverf[C NeverConditioner](condition C, timeout time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Never[C](a.T, condition, timeout, tick, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotBlockedT is the same as [NotBlockedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotBlockedT[E any, CHAN ~chan E](ch CHAN, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotBlockedT[E, CHAN](a.T, ch, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotBlockedTf is the same as [Assertions.NotBlockedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotBlockedTf[E any, CHAN ~chan E](ch CHAN, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotBlockedT[E, CHAN](a.T, ch, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotElementsMatchT is the same as [NotElementsMatchT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatchT[E comparable](listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatchT[E](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotElementsMatchTf is the same as [Assertions.NotElementsMatchT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatchTf[E comparable](listA []E, listB []E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatchT[E](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotEqualT is the same as [NotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualT[V comparable](expected V, actual V, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualT[V](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotEqualTf is the same as [Assertions.NotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualTf[V comparable](expected V, actual V, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualT[V](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotErrorAsType is the same as [NotErrorAsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAsType[E error](err error, target *E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAsType[E](a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotErrorAsTypef is the same as [Assertions.NotErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAsTypef[E error](err error, target *E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAsType[E](a.T, err, target, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotRegexpT is the same as [NotRegexpT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexpT[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexpT[Rex, ADoc](a.T, rx, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotRegexpTf is the same as [Assertions.NotRegexpT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexpTf[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexpT[Rex, ADoc](a.T, rx, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotSameT is the same as [NotSameT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSameT[P any](expected *P, actual *P, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSameT[P](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotSameTf is the same as [Assertions.NotSameT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSameTf[P any](expected *P, actual *P, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSameT[P](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// NotSortedT is the same as [NotSortedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSortedT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSortedT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotSortedTf is the same as [Assertions.NotSortedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSortedTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSortedT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// PositiveT is the same as [PositiveT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PositiveT[SignedNumber SignedNumeric](e SignedNumber, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PositiveT[SignedNumber](a.T, e, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// PositiveTf is the same as [Assertions.PositiveT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PositiveTf[SignedNumber SignedNumeric](e SignedNumber, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PositiveT[SignedNumber](a.T, e, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// RegexpT is the same as [RegexpT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) RegexpT[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.RegexpT[Rex, ADoc](a.T, rx, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// RegexpTf is the same as [Assertions.RegexpT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) RegexpTf[Rex RegExp, ADoc Text](rx Rex, actual ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.RegexpT[Rex, ADoc](a.T, rx, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SameT is the same as [SameT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SameT[P any](expected *P, actual *P, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SameT[P](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SameTf is the same as [Assertions.SameT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SameTf[P any](expected *P, actual *P, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SameT[P](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SeqContainsT is the same as [SeqContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SeqContainsT[E comparable](iter iter.Seq[E], element E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SeqContainsT[E](a.T, iter, element, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SeqContainsTf is the same as [Assertions.SeqContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SeqContainsTf[E comparable](iter iter.Seq[E], element E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SeqContainsT[E](a.T, iter, element, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SeqNotContainsT is the same as [SeqNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SeqNotContainsT[E comparable](iter iter.Seq[E], element E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SeqNotContainsT[E](a.T, iter, element, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SeqNotContainsTf is the same as [Assertions.SeqNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SeqNotContainsTf[E comparable](iter iter.Seq[E], element E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SeqNotContainsT[E](a.T, iter, element, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceContainsT is the same as [SliceContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceContainsT[Slice ~[]E, E comparable](s Slice, element E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceContainsT[Slice, E](a.T, s, element, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceContainsTf is the same as [Assertions.SliceContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceContainsTf[Slice ~[]E, E comparable](s Slice, element E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceContainsT[Slice, E](a.T, s, element, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceEqualT is the same as [SliceEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceEqualT[E comparable](listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceEqualT[E](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceEqualTf is the same as [Assertions.SliceEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceEqualTf[E comparable](listA []E, listB []E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceEqualT[E](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceNotContainsT is the same as [SliceNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotContainsT[Slice ~[]E, E comparable](s Slice, element E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotContainsT[Slice, E](a.T, s, element, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceNotContainsTf is the same as [Assertions.SliceNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotContainsTf[Slice ~[]E, E comparable](s Slice, element E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotContainsT[Slice, E](a.T, s, element, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceNotEqualT is the same as [SliceNotEqualT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotEqualT[E comparable](listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotEqualT[E](a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceNotEqualTf is the same as [Assertions.SliceNotEqualT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotEqualTf[E comparable](listA []E, listB []E, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotEqualT[E](a.T, listA, listB, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceNotSubsetT is the same as [SliceNotSubsetT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotSubsetT[Slice ~[]E, E comparable](list Slice, subset Slice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotSubsetT[Slice, E](a.T, list, subset, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceNotSubsetTf is the same as [Assertions.SliceNotSubsetT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceNotSubsetTf[Slice ~[]E, E comparable](list Slice, subset Slice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotSubsetT[Slice, E](a.T, list, subset, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SliceSubsetT is the same as [SliceSubsetT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceSubsetT[Slice ~[]E, E comparable](list Slice, subset Slice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceSubsetT[Slice, E](a.T, list, subset, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SliceSubsetTf is the same as [Assertions.SliceSubsetT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SliceSubsetTf[Slice ~[]E, E comparable](list Slice, subset Slice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SliceSubsetT[Slice, E](a.T, list, subset, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// SortedT is the same as [SortedT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SortedT[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SortedT[OrderedSlice, E](a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// SortedTf is the same as [Assertions.SortedT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) SortedTf[OrderedSlice ~[]E, E Ordered](collection OrderedSlice, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.SortedT[OrderedSlice, E](a.T, collection, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// StringContainsT is the same as [StringContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) StringContainsT[ADoc, EDoc Text](str ADoc, substring EDoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.StringContainsT[ADoc, EDoc](a.T, str, substring, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// StringContainsTf is the same as [Assertions.StringContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) StringContainsTf[ADoc, EDoc Text](str ADoc, substring EDoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.StringContainsT[ADoc, EDoc](a.T, str, substring, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// StringNotContainsT is the same as [StringNotContainsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) StringNotContainsT[ADoc, EDoc Text](str ADoc, substring EDoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.StringNotContainsT[ADoc, EDoc](a.T, str, substring, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// StringNotContainsTf is the same as [Assertions.StringNotContainsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) StringNotContainsTf[ADoc, EDoc Text](str ADoc, substring EDoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.StringNotContainsT[ADoc, EDoc](a.T, str, substring, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// TrueT is the same as [TrueT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) TrueT[B Boolean](value B, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.TrueT[B](a.T, value, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// TrueTf is the same as [Assertions.TrueT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) TrueTf[B Boolean](value B, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.TrueT[B](a.T, value, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// YAMLEqT is the same as [YAMLEqT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqT[EDoc, ADoc RText](expected EDoc, actual ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqT[EDoc, ADoc](a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// YAMLEqTf is the same as [Assertions.YAMLEqT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqTf[EDoc, ADoc RText](expected EDoc, actual ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqT[EDoc, ADoc](a.T, expected, actual, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// YAMLMarshalAsT is the same as [YAMLMarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLMarshalAsT[EDoc RText](expected EDoc, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLMarshalAsT[EDoc](a.T, expected, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// YAMLMarshalAsTf is the same as [Assertions.YAMLMarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLMarshalAsTf[EDoc RText](expected EDoc, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLMarshalAsT[EDoc](a.T, expected, object, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}

// YAMLUnmarshalAsT is the same as [YAMLUnmarshalAsT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLUnmarshalAsT[Object any, ADoc RText](expected Object, yamlDoc ADoc, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLUnmarshalAsT[Object, ADoc](a.T, expected, yamlDoc, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// YAMLUnmarshalAsTf is the same as [Assertions.YAMLUnmarshalAsT], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLUnmarshalAsTf[Object any, ADoc RText](expected Object, yamlDoc ADoc, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLUnmarshalAsT[Object, ADoc](a.T, expected, yamlDoc, forwardArgs(msg, args)...) {
		return
	}

	a.T.FailNow()
}
