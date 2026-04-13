// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package spew implements a deep pretty printer for Go data structures to aid in
// debugging.
//
// It exposes the pretty-printer used internally by [github.com/go-openapi/testify/v2].
//
// This a modernized version of the well-known [github.com/davecgh/go-spew].
//
// The original software is Copyright: 2012-2016 Dave Collins, under an ISC license.
package spew

import (
	"io"

	ispew "github.com/go-openapi/testify/v2/internal/spew"
)

// ConfigState houses the configuration options used by spew to format and
// display values.
type ConfigState = ispew.ConfigState

// Config is the active configuration of the top-level functions.
//
// This is an independent copy that does not affect the internal testify configuration.
// The configuration can be changed by modifying the contents of [spew.Config].
var Config = ConfigState{ //nolint:gochecknoglobals
	Indent:             " ",
	EnableTimeStringer: true,
}

// Dump displays the passed parameters to standard out with newlines, customizable
// indentation, and additional debug information such as complete types and all
// pointer addresses used to indirect to the final value.
func Dump(a ...any) {
	Config.Dump(a...)
}

// Fdump formats and displays the passed arguments to [io.Writer] w.  It formats
// exactly the  same as [Dump].
func Fdump(w io.Writer, a ...any) {
	Config.Fdump(w, a...)
}

// Sdump returns a string with the passed arguments formatted exactly the same
// as [Dump].
func Sdump(a ...any) string {
	return Config.Sdump(a...)
}
