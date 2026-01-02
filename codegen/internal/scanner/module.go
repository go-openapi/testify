// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime/debug"
	"time"
	"unicode"
)

// header returns a codegen header, more detailed than the Tool information.
func header() string {
	moduleVersion := moduleVersion()
	repoVersion := gitDescribe()
	ts := time.Now().Format(time.DateOnly)

	return fmt.Sprintf(
		"Generated on %s (version %s) using codegen version %s",
		ts, repoVersion, moduleVersion,
	)
}

func gitDescribe() string {
	const cmdTimeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()

	gitDesc, err := exec.CommandContext(ctx, "git", "rev-parse", "--short", "HEAD").CombinedOutput()
	if err != nil {
		log.Printf("warning: attempted to run git describe, but got: %v. Skipped", err)
		return ""
	}

	return string(bytes.TrimRightFunc(gitDesc, func(r rune) bool {
		return r == '\n' || r == '\r' || unicode.IsSpace(r)
	}))
}

// moduleName returns the main module name of the caller.
//
// This identifies the tool currently running the analysis.
//
// NOTE: no versioning here - the header provides more information.
func moduleName() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	return info.Main.Path
}

func moduleVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	var (
		modVersion string
		gitVersion string
	)

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			gitVersion = setting.Value
		}
	}

	if info.Main.Version == "(devel)" {
		modVersion = "master"
	} else {
		modVersion = info.Main.Version
	}

	final := modVersion
	if gitVersion != "" {
		final += " [sha: " + gitVersion + "]"
	}

	return final
}
