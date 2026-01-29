// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser_test

import (
	"flag"
	"strings"
	"sync"
	"testing"

	parser "github.com/go-openapi/testify/codegen/v2/internal/scanner/examples-parser"
)

//nolint:gochecknoglobals // it is okay to use globals for test flag and [sync.Once] cache for parsed testdata
var (
	render                = flag.Bool("render", false, "log rendered example output")
	parseTestExamplesOnce sync.Once
	cachedTestExamples    parser.Examples
)

const buildTag = "integrationtest"

func TestParse(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	if len(examples) == 0 {
		t.Fatal("expected examples, got none")
	}

	t.Run("Verify known function examples exist", func(t *testing.T) {
		for _, funcName := range []string{"Greet", "Add"} {
			t.Run("with "+funcName, func(t *testing.T) {
				exs, ok := examples[funcName]
				if !ok {
					t.Errorf("expected examples for %q, found none", funcName)
					return
				}
				if len(exs) == 0 {
					t.Errorf("expected at least one example for %q", funcName)
				}
			})
		}
	})

	t.Run("NoExample is exported but has no example", func(t *testing.T) {
		if _, ok := examples["NoExample"]; ok {
			t.Error("expected no examples for NoExample")
		}
	})

	t.Run("unexported must not appear", func(t *testing.T) {
		if _, ok := examples["unexported"]; ok {
			t.Error("expected no examples for unexported")
		}
	})

	if *render {
		t.Logf("found %d symbols with examples", len(examples))
	}
}

func TestParse_TypeExamples(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	exs, ok := examples["Formatter"]
	if !ok {
		t.Fatal("expected example for type Formatter, found none")
	}
	if len(exs) != 1 {
		t.Fatalf("expected 1 example for Formatter, got %d", len(exs))
	}
	if !exs[0].WholeFile {
		t.Error("expected WholeFile=true for Formatter")
	}
}

func TestParse_ExampleDetails(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	greetExamples, ok := examples["Greet"]
	if !ok || len(greetExamples) == 0 {
		t.Fatal("expected examples for Greet")
	}

	ex := greetExamples[0]
	if ex.Name != "Greet" {
		t.Errorf("expected Name=%q, got %q", "Greet", ex.Name)
	}
	if ex.Suffix != "" {
		t.Errorf("expected empty Suffix, got %q", ex.Suffix)
	}
	if ex.Output != "Hello, World!\n" {
		t.Errorf("expected Output=%q, got %q", "Hello, World!\n", ex.Output)
	}
	if ex.WholeFile {
		t.Error("expected WholeFile=false for ExampleGreet")
	}
}

func TestParse_SuffixedExample(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	addExamples, ok := examples["Add"]
	if !ok {
		t.Fatal("expected examples for Add")
	}

	// Add has two examples: ExampleAdd and ExampleAdd_negative.
	if len(addExamples) != 2 {
		t.Fatalf("expected 2 examples for Add, got %d", len(addExamples))
	}

	var found bool
	for _, ex := range addExamples {
		if ex.Suffix == "negative" {
			found = true

			break
		}
	}
	if !found {
		t.Error("expected an example with Suffix=\"negative\" for Add")
	}
}

func TestParse_Render(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	greetExamples := examples["Greet"]
	if len(greetExamples) == 0 {
		t.Fatal("expected examples for Greet")
	}

	rendered := greetExamples[0].Render()
	if rendered == "" {
		t.Fatal("Render() returned empty string")
	}

	t.Run("Should contain the function call", func(t *testing.T) {
		if !strings.Contains(rendered, "Greet") {
			t.Errorf("expected rendered code to contain 'Greet', got:\n%s", rendered)
		}
	})

	t.Run(`Should NOT contain "// Output:" lines`, func(t *testing.T) {
		if strings.Contains(rendered, "// Output:") {
			t.Errorf("expected rendered code to NOT contain '// Output:', got:\n%s", rendered)
		}
	})

	t.Run("Should NOT have outer braces", func(t *testing.T) {
		trimmed := strings.TrimSpace(rendered)
		if strings.HasPrefix(trimmed, "{") {
			t.Errorf("expected rendered code without outer braces, got:\n%s", rendered)
		}
	})

	if *render {
		t.Logf("Rendered ExampleGreet:\n%s", rendered)
	}
}

func TestParse_RenderWholeFile(t *testing.T) {
	t.Parallel()

	examples := loadTestExamples(t)

	exs := examples["Formatter"]
	if len(exs) == 0 {
		t.Fatal("expected example for Formatter")
	}

	rendered := exs[0].Render()
	if rendered == "" {
		t.Fatal("Render() returned empty string")
	}

	t.Run("Should contain a main function", func(t *testing.T) {
		if !strings.Contains(rendered, "func main()") {
			t.Errorf("expected rendered code to contain main function declaration, got:\n%s", rendered)
		}
	})

	t.Run("Should contain the supporting type declaration", func(t *testing.T) {
		if !strings.Contains(rendered, "type helper struct") {
			t.Errorf("expected rendered code to contain 'type helper struct', got:\n%s", rendered)
		}
	})

	t.Run("Should contain package clause or imports", func(t *testing.T) {
		if !strings.Contains(rendered, "package ") {
			t.Errorf("expected rendered code without package clause, got:\n%s", rendered)
		}
		if !strings.Contains(rendered, "import ") {
			t.Errorf("expected rendered code without imports, got:\n%s", rendered)
		}
	})

	if *render {
		t.Logf("Rendered ExampleFormatter:\n%s", rendered)
	}
}

func TestParse_NonExistentPackage(t *testing.T) {
	t.Parallel()

	ext := parser.New("./nonexistent/package/path")

	_, err := ext.Parse()
	if err == nil {
		t.Error("expected error for non-existent package")
	}
}

func TestRender_NilCode(t *testing.T) {
	t.Parallel()

	ex := parser.TestableExample{}
	if got := ex.Render(); got != "" {
		t.Errorf("expected empty string for nil code, got %q", got)
	}
}

// loadTestExamples parses the testdata package once and caches the result.
func loadTestExamples(t *testing.T) parser.Examples {
	t.Helper()

	parseTestExamplesOnce.Do(func() {
		examples, err := parser.New(
			"./testdata/examplespkg",
			parser.WithBuildTags(buildTag),
		).Parse()
		if err != nil {
			t.Fatalf("Parse() error: %v", err)
		}

		cachedTestExamples = examples
	})

	return cachedTestExamples
}
