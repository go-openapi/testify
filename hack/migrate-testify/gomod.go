// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// updateGoMod updates the target project's go.mod to replace stretchr/testify
// with go-openapi/testify/v2.
func updateGoMod(dir, version string, dryRun, verbose bool) error {
	// Check go.mod exists.
	if _, err := os.Stat(dir + "/go.mod"); err != nil {
		return fmt.Errorf("no go.mod found in %s", dir)
	}

	if dryRun {
		fmt.Println("would run: go mod edit -droprequire github.com/stretchr/testify")               //nolint:forbidigo // CLI output
		fmt.Printf("would run: go mod edit -require github.com/go-openapi/testify/v2@%s\n", version) //nolint:forbidigo // CLI output
		fmt.Println("would run: go mod tidy")                                                        //nolint:forbidigo // CLI output
		return nil
	}

	commands := []struct {
		name string
		args []string
	}{
		{
			name: "drop stretchr/testify require",
			args: []string{"go", "mod", "edit", "-droprequire", "github.com/stretchr/testify"},
		},
		{
			name: "add go-openapi/testify/v2 require",
			args: []string{"go", "mod", "edit", "-require", "github.com/go-openapi/testify/v2@" + version},
		},
		{
			name: "go mod tidy",
			args: []string{"go", "mod", "tidy"},
		},
	}

	for _, c := range commands {
		if verbose {
			fmt.Printf("running: %s\n", strings.Join(c.args, " ")) //nolint:forbidigo // CLI output
		}

		cmd := exec.CommandContext(context.Background(), c.args[0], c.args[1:]...) //nolint:gosec // args are constructed from controlled constants
		cmd.Dir = dir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("%s: %w", c.name, err)
		}
	}

	return nil
}
