// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package require_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleErrorAsType() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorAsType(t *testing.T)
	require.ErrorAsType(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotErrorAsType() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorAsType(t *testing.T)
	require.NotErrorAsType(t, require.ErrTest, new(*dummyError))
	fmt.Println("passed")

	// Output: passed
}
