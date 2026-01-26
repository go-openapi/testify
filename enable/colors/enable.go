// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"

	colorstub "github.com/go-openapi/testify/v2/enable/stubs/colors"
)

const (
	envVarColorize = "TESTIFY_COLORIZED"
	envVarTheme    = "TESTIFY_THEME"
	envVarNoTTY    = "TESTIFY_COLORIZED_NOTTY"
)

var flags cliFlags //nolint:gochecknoglobals // it's okay to store the state CLI flags in a package global

type cliFlags struct {
	colorized bool
	theme     string
	notty     bool
}

func init() { //nolint:gochecknoinits // it's okay: we want to declare CLI flags when a blank import references this package
	isTerminal := term.IsTerminal(1)

	flag.BoolVar(&flags.colorized, "testify.colorized", colorizeFromEnv(), "testify: colorized output")
	flag.StringVar(&flags.theme, "testify.theme", themeFromEnv(), "testify: color theme (light,dark)")
	flag.BoolVar(&flags.notty, "testify.colorized.notty", nottyFromEnv(), "testify: force colorization, even if not a tty")

	colorstub.Enable(
		func() []colorstub.Option {
			return []colorstub.Option{
				colorstub.WithEnable(flags.colorized && (isTerminal || flags.notty)),
				colorstub.WithSanitizedTheme(flags.theme),
			}
		})
}

func colorizeFromEnv() bool {
	envColorize := os.Getenv(envVarColorize)
	isEnvConfigured, _ := strconv.ParseBool(envColorize)

	return isEnvConfigured
}

func themeFromEnv() string {
	envTheme := os.Getenv(envVarTheme)

	return strings.ToLower(envTheme)
}

func nottyFromEnv() bool {
	envNoTTY := os.Getenv(envVarNoTTY)
	isEnvNoTTY, _ := strconv.ParseBool(envNoTTY)

	return isEnvNoTTY
}
