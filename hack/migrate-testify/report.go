// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"slices"
	"strings"
)

// diagnostic represents a single warning or info message from the migration tool.
type diagnostic struct {
	file    string
	line    int
	message string
	kind    string // "warning", "info", "error"
}

func (d diagnostic) String() string {
	if d.line > 0 {
		return fmt.Sprintf("%s:%d: %s: %s", d.file, d.line, d.kind, d.message)
	}
	return fmt.Sprintf("%s: %s: %s", d.file, d.kind, d.message)
}

// report collects diagnostics and file changes during a migration run.
type report struct {
	diagnostics  []diagnostic
	filesChanged int
	totalChanges int
	filesScanned int
	// Pass 1 stats: funcName → call count (e.g. "assert.Equal" → 42).
	apiUsage map[string]int
	// Pass 2 stats: "Equal→EqualT" → count.
	upgraded map[string]int
	// Pass 2 stats: skipReason → count.
	skipped map[string]int
}

func (r *report) warn(file string, line int, msg string) {
	r.diagnostics = append(r.diagnostics, diagnostic{file: file, line: line, message: msg, kind: "warning"})
}

func (r *report) info(file string, line int, msg string) {
	r.diagnostics = append(r.diagnostics, diagnostic{file: file, line: line, message: msg, kind: "info"})
}

func (r *report) errorf(file string, line int, msg string) {
	r.diagnostics = append(r.diagnostics, diagnostic{file: file, line: line, message: msg, kind: "error"})
}

// trackAPIUsage increments the usage counter for a testify API call.
func (r *report) trackAPIUsage(qualifiedName string) {
	if r.apiUsage == nil {
		r.apiUsage = make(map[string]int)
	}
	r.apiUsage[qualifiedName]++
}

// trackUpgrade increments the upgrade counter for a successful generic upgrade.
func (r *report) trackUpgrade(from, to string) {
	if r.upgraded == nil {
		r.upgraded = make(map[string]int)
	}
	r.upgraded[from+" → "+to]++
}

// trackSkip records a skipped generic upgrade and emits an info diagnostic.
func (r *report) trackSkip(file string, line int, funcName string, reason skipReason, verbose bool, typeInfo string) {
	if r.skipped == nil {
		r.skipped = make(map[string]int)
	}
	r.skipped[string(reason)]++
	if verbose {
		msg := fmt.Sprintf("skipped %s: %s", funcName, reason)
		if typeInfo != "" {
			msg += " [" + typeInfo + "]"
		}
		r.info(file, line, msg)
	}
}

// print outputs all collected diagnostics.
func (r *report) print(verbose bool) {
	if verbose {
		for _, d := range r.diagnostics {
			fmt.Println(d) //nolint:forbidigo // CLI output
		}
	} else {
		// Only print warnings and errors, not info diagnostics.
		for _, d := range r.diagnostics {
			if d.kind != "info" {
				fmt.Println(d) //nolint:forbidigo // CLI output
			}
		}
	}
}

// printPass1Summary outputs the Pass 1 structured summary.
func (r *report) printPass1Summary() {
	fmt.Printf("\n=== Pass 1 Summary ===\n")                                        //nolint:forbidigo // CLI output
	fmt.Printf("Files scanned: %d  |  Files changed: %d  |  Transformations: %d\n", //nolint:forbidigo // CLI output
		r.filesScanned, r.filesChanged, r.totalChanges)

	if len(r.apiUsage) > 0 {
		fmt.Printf("\nAPI usage across migrated scope:\n") //nolint:forbidigo // CLI output
		printCountTable(r.apiUsage)
	}

	warnings := 0
	for _, d := range r.diagnostics {
		if d.kind == "warning" {
			warnings++
		}
	}
	fmt.Printf("\nWarnings: %d\n", warnings) //nolint:forbidigo // CLI output
}

// printPass2Summary outputs the Pass 2 structured summary.
func (r *report) printPass2Summary() {
	fmt.Printf("\n=== Pass 2 Summary ===\n")                                 //nolint:forbidigo // CLI output
	fmt.Printf("Files scanned: %d  |  Files changed: %d  |  Upgrades: %d\n", //nolint:forbidigo // CLI output
		r.filesScanned, r.filesChanged, r.totalChanges)

	if len(r.upgraded) > 0 {
		fmt.Printf("\nUpgraded assertions:\n") //nolint:forbidigo // CLI output
		printCountTable(r.upgraded)
	}

	if len(r.skipped) > 0 {
		total := 0
		for _, v := range r.skipped {
			total += v
		}
		fmt.Printf("\nSkipped (generic alternative exists but cannot upgrade): %d\n", total) //nolint:forbidigo // CLI output
		printCountTable(r.skipped)
	}
}

// printCountTable prints a map of label→count as a right-aligned table with two columns.
func printCountTable(m map[string]int) {
	type entry struct {
		label string
		count int
	}
	entries := make([]entry, 0, len(m))
	for k, v := range m {
		entries = append(entries, entry{k, v})
	}
	slices.SortFunc(entries, func(a, b entry) int {
		if c := cmp.Compare(b.count, a.count); c != 0 {
			return c
		}
		return cmp.Compare(a.label, b.label)
	})

	// Find max label width for alignment.
	maxLabel := 0
	for _, e := range entries {
		if len(e.label) > maxLabel {
			maxLabel = len(e.label)
		}
	}

	// Print in two columns if there are enough entries.
	for i := 0; i < len(entries); i += 2 {
		left := entries[i]
		if i+1 < len(entries) {
			right := entries[i+1]
			fmt.Printf("  %-*s %5d    %-*s %5d\n", maxLabel, left.label, left.count, maxLabel, right.label, right.count) //nolint:forbidigo // CLI output
		} else {
			fmt.Printf("  %-*s %5d\n", maxLabel, left.label, left.count) //nolint:forbidigo // CLI output
		}
	}
}

// writeFile writes a modified AST back to disk via go/printer.
func writeFile(fset *token.FileSet, f *ast.File, filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating %s: %w", filename, err)
	}
	defer out.Close()

	cfg := &printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 8, //nolint:mnd // standard Go tabwidth
	}
	if err := cfg.Fprint(out, fset, f); err != nil {
		return fmt.Errorf("writing %s: %w", filename, err)
	}
	return nil
}

// showDiff displays a simple diff-like output showing what would change.
func showDiff(fset *token.FileSet, f *ast.File, filename string) error {
	original, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var buf strings.Builder
	cfg := &printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 8, //nolint:mnd // standard Go tabwidth
	}
	if err := cfg.Fprint(&buf, fset, f); err != nil {
		return err
	}

	modified := buf.String()
	if string(original) == modified {
		return nil
	}

	fmt.Printf("--- %s\n+++ %s (modified)\n", filename, filename) //nolint:forbidigo // CLI output

	origLines := strings.Split(string(original), "\n")
	modLines := strings.Split(modified, "\n")

	// Simple line-by-line diff — not a real unified diff, but helpful for dry-run.
	maxLines := max(len(origLines), len(modLines))

	for i := range maxLines {
		var origLine, modLine string
		if i < len(origLines) {
			origLine = origLines[i]
		}
		if i < len(modLines) {
			modLine = modLines[i]
		}
		if origLine != modLine {
			if origLine != "" {
				fmt.Printf("-%s\n", origLine) //nolint:forbidigo // CLI output
			}
			if modLine != "" {
				fmt.Printf("+%s\n", modLine) //nolint:forbidigo // CLI output
			}
		}
	}
	fmt.Println() //nolint:forbidigo // CLI output
	return nil
}
