// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package spew runs integration tests against [github.com/go-openapi/testify/v2/internal/spew].
//
// It knows how to generate random data structures and how to fuzz that package.
//
// # Motivation
//
// Even though spew is a very well tested library, its heavy use of [reflect] makes it
// highly sensitive to panicking issues.
//
// This package is here to make sure that all edge-cases are eventually explored
// and tested.
//
// The implementation depends on package [pgregory.net/rapid].
//
// This internal testing functionality is therefore defined as an independent module and
// does not affect the dependencies of [github.com/go-openapi/testify/v2].
package spew
