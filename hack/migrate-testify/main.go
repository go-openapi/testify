// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package main implements a migration tool for converting stretchr/testify
// usage to go-openapi/testify/v2, and upgrading reflection-based assertions
// to generic variants where type information permits.
//
// Usage:
//
//	go run ./hack/migrate-testify [flags] [directory]
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type options struct {
	migrate         bool
	upgradeGenerics bool
	all             bool
	dryRun          bool
	verbose         bool
	skipGomod       bool
	skipVendor      bool
	version         string
}

func main() {
	opts := &options{}

	flag.BoolVar(&opts.migrate, "migrate", false, "Run pass 1: stretchr/testify → go-openapi/testify/v2")
	flag.BoolVar(&opts.upgradeGenerics, "upgrade-generics", false, "Run pass 2: reflection → generic assertions")
	flag.BoolVar(&opts.all, "all", false, "Run both passes sequentially")
	flag.BoolVar(&opts.dryRun, "dry-run", false, "Show diffs without modifying files")
	flag.BoolVar(&opts.verbose, "verbose", false, "Print detailed transformation info")
	flag.BoolVar(&opts.skipGomod, "skip-gomod", false, "Skip go.mod changes")
	flag.BoolVar(&opts.skipVendor, "skip-vendor", true, "Skip vendor/ directory")
	flag.StringVar(&opts.version, "version", "v2.3.0", "Target testify version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go run ./hack/migrate-testify [flags] [directory]\n\n")
		fmt.Fprintf(os.Stderr, "Migrate stretchr/testify to go-openapi/testify/v2 and upgrade to generic assertions.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nAt least one of --migrate, --upgrade-generics, or --all is required.\n")
		fmt.Fprintf(os.Stderr, "\nMono-repo support:\n")
		fmt.Fprintf(os.Stderr, "  Pass 1 walks the filesystem and works across module boundaries.\n")
		fmt.Fprintf(os.Stderr, "  Pass 2 requires type information and uses go/packages to load code.\n")
		fmt.Fprintf(os.Stderr, "  For multi-module repos, a go.work file must be present so that pass 2\n")
		fmt.Fprintf(os.Stderr, "  can load all workspace modules. Create one with:\n")
		fmt.Fprintf(os.Stderr, "    go work init . ./sub/module1 ./sub/module2 ...\n")
		fmt.Fprintf(os.Stderr, "\nPost-migration checklist:\n")
		fmt.Fprintf(os.Stderr, "  - Run your linter: the migration may surface pre-existing unchecked linting issues.\n")
		fmt.Fprintf(os.Stderr, "  - Run your test suite to verify all tests still pass.\n")
	}

	flag.Parse()

	if opts.all {
		opts.migrate = true
		opts.upgradeGenerics = true
	}

	if !opts.migrate && !opts.upgradeGenerics {
		flag.Usage()
		os.Exit(2) //nolint:mnd // standard exit code for usage error
	}

	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	// Pre-flight: warn if git is dirty.
	checkGitDirty(dir)

	if opts.migrate {
		fmt.Println("=== Pass 1: Migration (stretchr/testify → go-openapi/testify/v2) ===") //nolint:forbidigo // CLI output
		if err := runMigration(dir, opts); err != nil {
			fmt.Fprintf(os.Stderr, "error: migration: %v\n", err)
			os.Exit(1)
		}
	}

	if opts.upgradeGenerics {
		fmt.Println("=== Pass 2: Generic Upgrade (reflection → generic assertions) ===") //nolint:forbidigo // CLI output
		if err := runGenericUpgrade(dir, opts); err != nil {
			fmt.Fprintf(os.Stderr, "error: generic upgrade: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Done.") //nolint:forbidigo // CLI output
}

// checkGitDirty warns if the working directory has uncommitted changes.
func checkGitDirty(dir string) {
	cmd := exec.CommandContext(context.Background(), "git", "status", "--porcelain")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return // Not a git repo or git not available — skip check.
	}
	output := strings.TrimSpace(string(out))
	if output != "" {
		fmt.Fprintf(os.Stderr, "warning: working directory has uncommitted changes\n")
	}
}
