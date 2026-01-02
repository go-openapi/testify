// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package main provides the testify code generator.
//
// This program reads all assertion functions from the internal/assertions package and
// automatically generates the corresponding packages assert and require with all variants.
package main

import (
	"flag"
	"log"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/generator"
	"github.com/go-openapi/testify/codegen/v2/internal/model"
	"github.com/go-openapi/testify/codegen/v2/internal/scanner"
)

type config struct {
	dir        string
	targetRoot string
	targetDoc  string

	inputPkg   string
	outputPkgs string

	includeFmt bool
	includeFwd bool
	includeTst bool
	includeGen bool
	includeHlp bool
	includeExa bool
	includeDoc bool
	runExa     bool
}

func main() {
	cfg := new(config)
	registerFlags(cfg)
	flag.Parse() // exits if invalid flags

	if err := execute(cfg); err != nil {
		log.Fatal(err)
	}
}

func registerFlags(cfg *config) {
	flag.StringVar(&cfg.dir, "work-dir", "..", "working directory to scan for package (default is the parent folder)")
	flag.StringVar(&cfg.inputPkg, "input-package", "github.com/go-openapi/testify/v2/internal/assertions", "source package to scan to produce generated output")
	flag.StringVar(&cfg.outputPkgs, "output-packages", "assert,require", "package(s) to generate")
	flag.StringVar(&cfg.targetRoot, "target-root", "..", "where to place the generated packages (default is in the parent folder)")
	flag.StringVar(&cfg.targetDoc, "target-doc", "docs/doc-site/api", "where to place the generated documentation, relative to the target root")

	flag.BoolVar(&cfg.includeFmt, "include-format-funcs", true, "include format functions such as Errorf and Equalf")
	flag.BoolVar(&cfg.includeFwd, "include-forward-funcs", true, "include forward assertions functions as methods of the Assertions object")
	flag.BoolVar(&cfg.includeTst, "include-tests", true, "generate the tests in the target package(s)")
	flag.BoolVar(&cfg.includeGen, "include-generics", false, "include generic functions")
	flag.BoolVar(&cfg.includeHlp, "include-helpers", true, "include helper functions that are not assertions")
	flag.BoolVar(&cfg.includeExa, "include-examples", true, "include generated testable examples")
	flag.BoolVar(&cfg.runExa, "runnable-examples", true, "include Output to generated testable examples, so they are run as tests")
	flag.BoolVar(&cfg.includeDoc, "include-doc", true, "include generated markdown documentation")
}

func execute(cfg *config) error {
	scanner := scanner.New(
		scanner.WithWorkDir(cfg.dir),
		scanner.WithPackage(cfg.inputPkg),
		scanner.WithCollectDoc(cfg.includeDoc), // collects extra documentation from comments
	)
	scanned, err := scanner.Scan()
	if err != nil {
		return err
	}

	builder := generator.New(scanned)
	var doc model.Documentation

	for targetPkg := range strings.SplitSeq(cfg.outputPkgs, ",") {
		err = builder.Generate(
			// where to generate
			generator.WithTargetRoot(cfg.targetRoot),
			generator.WithTargetPackage(targetPkg),
			// what to generate
			generator.WithIncludeFormatFuncs(cfg.includeFmt),
			generator.WithIncludeForwardFuncs(cfg.includeFwd),
			generator.WithIncludeGenerics(cfg.includeGen),
			generator.WithIncludeTests(cfg.includeTst),
			generator.WithIncludeHelpers(cfg.includeHlp),
			generator.WithIncludeExamples(cfg.includeExa),
			generator.WithRunnableExamples(cfg.runExa),
			generator.WithIncludeDoc(cfg.includeDoc),
		)
		if err != nil {
			return err
		}

		if cfg.includeDoc {
			// stash the transformed doc
			doc.Merge(builder.Documentation())
		}
	}
	if !cfg.includeDoc {
		// we're done with codegen
		return nil
	}

	// and now for something completely different: generating the documentation
	documentalist := generator.NewDocGenerator(doc)
	err = documentalist.Generate(
		// where to generate
		generator.WithTargetRoot(cfg.targetRoot),
		generator.WithTargetDoc(cfg.targetDoc),
		// what to generate
		generator.WithIncludeFormatFuncs(cfg.includeFmt),
		generator.WithIncludeForwardFuncs(cfg.includeFwd),
		generator.WithIncludeGenerics(cfg.includeGen),
		generator.WithIncludeHelpers(cfg.includeHlp),
	)
	if err != nil {
		return err
	}

	return nil
}
