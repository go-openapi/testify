// Package main provides the testify code generator.
//
// This program reads all assertion functions from the internal/assertions package and
// automatically generates the corresponding packages assert and require with all variants.
package main

import (
	"flag"
	"log"
	"strings"

	"github.com/go-openapi/testify/v2/codegen/internal/generator"
	"github.com/go-openapi/testify/v2/codegen/internal/scanner"
)

type config struct {
	dir        string
	targetRoot string
	inputPkg   string
	outputPkgs string

	includeFmt bool
	includeFwd bool
	includeTst bool
	includeGen bool
	includeHlp bool
	includeExa bool
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

	flag.BoolVar(&cfg.includeFmt, "include-format-funcs", true, "include format functions such as Errorf and Equalf")
	flag.BoolVar(&cfg.includeFwd, "include-forward-funcs", true, "include forward assertions functions as methods of the Assertions object")
	flag.BoolVar(&cfg.includeTst, "include-tests", true, "generate the tests in the target package(s)")
	flag.BoolVar(&cfg.includeGen, "include-generics", false, "include generic functions")
	flag.BoolVar(&cfg.includeHlp, "include-helpers", true, "include helper functions that are not assertions")
	flag.BoolVar(&cfg.includeExa, "include-examples", true, "include generated testable examples")
	flag.BoolVar(&cfg.runExa, "runnable-examples", true, "include Output to generated testable examples, so they are run as tests")
}

func execute(cfg *config) error {
	scanner := scanner.New(
		scanner.WithWorkDir(cfg.dir),
		scanner.WithPackage(cfg.inputPkg),
	)
	scanned, err := scanner.Scan()
	if err != nil {
		return err
	}

	builder := generator.New(scanned)

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
		)
		if err != nil {
			return err
		}
	}

	return nil
}
