// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"fmt"
	"go/token"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
	parser "github.com/go-openapi/testify/codegen/v2/internal/scanner/comments-parser"
)

const (
	testPackage        = "github.com/go-openapi/testify/v2/internal/assertions"
	testRepo           = "github.com/go-openapi/testify/v2"
	testAssertPackage  = "github.com/go-openapi/testify/v2/assert"
	testRequirePackage = "github.com/go-openapi/testify/v2/require"
	githubRepo         = "https://github.com/go-openapi/testify"
)

// TestFuncMap verifies that FuncMap returns all expected functions.
func TestFuncMap(t *testing.T) {
	t.Parallel()

	fm := FuncMap()

	expectedFuncs := []string{
		"imports", "comment", "date", "params", "forward", "docStringFor", "docStringPackage",
		"returns", "concat", "pathparts", "relocate", "hasSuffix", "sourceLink",
		"titleize", "quote", "mdformat", "godocbadge", "debug",
	}

	for _, name := range expectedFuncs {
		if _, ok := fm[name]; !ok {
			t.Errorf("FuncMap missing expected function: %q", name)
		}
	}
}

// TestPrintImports verifies import formatting.
func TestPrintImports(t *testing.T) {
	t.Parallel()

	for c := range printImportsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := printImports(c.input)

			if result != c.expected {
				t.Errorf("Expected:\n%s\n\nGot:\n%s", c.expected, result)
			}
		})
	}
}

// TestComment verifies comment formatting.
func TestComment(t *testing.T) {
	t.Parallel()

	for c := range commentCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := comment(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestParams verifies parameter formatting.
func TestParams(t *testing.T) {
	t.Parallel()

	for c := range paramsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := params(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestForward verifies forward parameter formatting.
func TestForward(t *testing.T) {
	t.Parallel()

	for c := range forwardCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := forward(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestPrintReturns verifies return value formatting.
func TestPrintReturns(t *testing.T) {
	t.Parallel()

	for c := range printReturnsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := PrintReturns(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestDocStringFor verifies usage-specific doc string generation.
func TestDocStringFor(t *testing.T) {
	t.Parallel()

	for c := range docStringForCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := docStringFor(c.usage, c.funcName)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestDocStringPackage verifies package-specific doc string generation.
func TestDocStringPackage(t *testing.T) {
	t.Parallel()

	for c := range docStringPackageCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := docStringPackage(c.pkg)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestConcatStrings verifies string concatenation.
func TestConcatStrings(t *testing.T) {
	t.Parallel()

	for c := range concatStringsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := concatStrings(c.input...)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestPathParts verifies path splitting and quoting.
func TestPathParts(t *testing.T) {
	t.Parallel()

	for c := range pathPartsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := pathParts(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestRelocate verifies special value relocation for testable examples.
func TestRelocate(t *testing.T) {
	t.Parallel()

	for c := range relocateCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// Parse input string into TestValue slice
			values := parser.ParseTestValues(c.input)
			result := relocate(values, c.pkg)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestSourceLink verifies GitHub source link generation.
func TestSourceLink(t *testing.T) {
	t.Parallel()

	for c := range sourceLinkCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := sourceLink(c.baseURL, c.pos)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestQuote verifies string quoting.
func TestQuote(t *testing.T) {
	t.Parallel()

	for c := range quoteCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := quote(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestTitleize verifies title casing.
func TestTitleize(t *testing.T) {
	t.Parallel()

	for c := range titleizeCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := Titleize(c.input)

			if result != c.expected {
				t.Errorf("Expected: %q, Got: %q", c.expected, result)
			}
		})
	}
}

// TestGodocbadge verifies pkg.go.dev badge URL generation.
func TestGodocbadge(t *testing.T) {
	t.Parallel()

	for c := range godocbadgeCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result, err := godocbadge(c.input)

			if c.expectError {
				if err == nil {
					t.Error("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != c.expected {
					t.Errorf("Expected: %q, Got: %q", c.expected, result)
				}
			}
		})
	}
}

func TestDebug(t *testing.T) {
	t.Parallel()

	result := printDebug(1)
	const expected = "1"
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestData(t *testing.T) {
	t.Parallel()

	result := printDate()
	if result == "" {
		t.Errorf("Expected a non-empty string, Got: %q", result)
	}
}

/* Test case iterators */

type printImportsCase struct {
	name     string
	input    model.ImportMap
	expected string
}

func printImportsCases() iter.Seq[printImportsCase] {
	return slices.Values([]printImportsCase{
		{
			name:     "empty imports",
			input:    model.ImportMap{},
			expected: "",
		},
		{
			name: "single import without alias",
			input: model.ImportMap{
				"testing": "testing",
			},
			expected: "\t\"testing\"",
		},
		{
			name: "single import with alias",
			input: model.ImportMap{
				"req": testRequirePackage,
			},
			expected: fmt.Sprintf("\treq\t%q", testRequirePackage),
		},
		{
			name: "multiple imports sorted alphabetically",
			input: model.ImportMap{
				"testing":    "testing",
				"assertions": testPackage,
				"fmt":        "fmt",
			},
			// Sorted by the full formatted string, not package name
			// No alias for assertions since it matches path.Base
			expected: fmt.Sprintf("\t\"fmt\"\n\t%q\n\t\"testing\"", testPackage),
		},
	})
}

type commentCase struct {
	name     string
	input    string
	expected string
}

func commentCases() iter.Seq[commentCase] {
	return slices.Values([]commentCase{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single line",
			input:    "This is a comment",
			expected: "// This is a comment",
		},
		{
			name:     "multiline",
			input:    "Line 1\nLine 2\nLine 3",
			expected: "// Line 1\n// Line 2\n// Line 3",
		},
		{
			name:     "trailing empty lines trimmed",
			input:    "Line 1\nLine 2\n\n\n",
			expected: "// Line 1\n// Line 2",
		},
	})
}

type paramsCase struct {
	name     string
	input    model.Parameters
	expected string
}

func paramsCases() iter.Seq[paramsCase] {
	return slices.Values([]paramsCase{
		{
			name:     "empty parameters",
			input:    model.Parameters{},
			expected: "",
		},
		{
			name: "single parameter",
			input: model.Parameters{
				{Name: "t", GoType: "TestingT"},
			},
			expected: "t TestingT",
		},
		{
			name: "multiple parameters",
			input: model.Parameters{
				{Name: "expected", GoType: "any"},
				{Name: "actual", GoType: "any"},
			},
			expected: "expected any, actual any",
		},
		{
			name: "parameter with Func type and assertions selector",
			input: model.Parameters{
				{Name: "fn", GoType: "ComparisonAssertionFunc", Selector: "assertions"},
			},
			expected: "fn assertions.ComparisonAssertionFunc",
		},
		{
			name: "variadic parameter as last argument",
			input: model.Parameters{
				{Name: "t", GoType: "TestingT"},
				{Name: "expected", GoType: "any"},
				{Name: "msgAndArgs", GoType: "...any", IsVariadic: true},
			},
			expected: "t TestingT, expected any, msgAndArgs ...any",
		},
		{
			name: "variadic parameter as unique argument",
			input: model.Parameters{
				{Name: "msgAndArgs", GoType: "...any", IsVariadic: true},
			},
			expected: "msgAndArgs ...any",
		},
		{
			name: "special handling of xxxFunc type as unique argument",
			input: model.Parameters{
				{Name: "fn", GoType: "PanicFunc", Selector: "assertions"},
			},
			expected: "fn assertions.PanicFunc",
		},
		{
			name: "special handling of xxxFunc type with multiple arguments",
			input: model.Parameters{
				{Name: "t", GoType: "TestingT"},
				{Name: "fn", GoType: "PanicFunc", Selector: "assertions"},
			},
			expected: "t TestingT, fn assertions.PanicFunc",
		},
	})
}

type forwardCase struct {
	name     string
	input    model.Parameters
	expected string
}

func forwardCases() iter.Seq[forwardCase] {
	return slices.Values([]forwardCase{
		{
			name:     "empty parameters",
			input:    model.Parameters{},
			expected: "",
		},
		{
			name: "single parameter",
			input: model.Parameters{
				{Name: "expected"},
			},
			expected: "expected",
		},
		{
			name: "multiple parameters",
			input: model.Parameters{
				{Name: "expected"},
				{Name: "actual"},
				{Name: "msgAndArgs"},
			},
			expected: "expected, actual, msgAndArgs",
		},
		{
			name: "variadic parameter as last argument",
			input: model.Parameters{
				{Name: "expected"},
				{Name: "msgAndArgs", IsVariadic: true},
			},
			expected: "expected, msgAndArgs...",
		},
		{
			name: "variadic parameter as unique argument",
			input: model.Parameters{
				{Name: "msgAndArgs", GoType: "...any", IsVariadic: true},
			},
			expected: "msgAndArgs...",
		},
	})
}

type printReturnsCase struct {
	name     string
	input    model.Parameters
	expected string
}

func printReturnsCases() iter.Seq[printReturnsCase] {
	return slices.Values([]printReturnsCase{
		{
			name:     "no returns",
			input:    model.Parameters{},
			expected: "",
		},
		{
			name: "single unnamed return",
			input: model.Parameters{
				{GoType: "bool"},
			},
			expected: "bool",
		},
		{
			name: "single named return",
			input: model.Parameters{
				{Name: "success", GoType: "bool"},
			},
			expected: "(success bool)",
		},
		{
			name: "multiple unnamed returns",
			input: model.Parameters{
				{GoType: "string"},
				{GoType: "error"},
			},
			expected: "(string, error)",
		},
		{
			name: "multiple named returns",
			input: model.Parameters{
				{Name: "result", GoType: "string"},
				{Name: "err", GoType: "error"},
			},
			expected: "(result string, err error)",
		},
	})
}

type docStringForCase struct {
	name     string
	usage    string
	funcName string
	expected string
}

func docStringForCases() iter.Seq[docStringForCase] {
	return slices.Values([]docStringForCase{
		{
			name:     "format usage",
			usage:    "format",
			funcName: "assert.Equal",
			expected: "// Equalf is the same as [assert.Equal], but it accepts a format string to format arguments like [fmt.Printf].",
		},
		{
			name:     "forward usage",
			usage:    "forward",
			funcName: "assert.Equal",
			expected: "// Equal is the same as [assert.Equal], as a method rather than a package-level function.",
		},
		{
			name:     "unknown usage",
			usage:    "unknown",
			funcName: "assert.Equal",
			expected: "",
		},
	})
}

type docStringPackageCase struct {
	name     string
	pkg      string
	expected string
}

func docStringPackageCases() iter.Seq[docStringPackageCase] {
	return slices.Values([]docStringPackageCase{
		{
			name:     "assert package",
			pkg:      "assert",
			expected: `// Upon failure, the test [T] is marked as failed and continues execution.`,
		},
		{
			name:     "require package",
			pkg:      "require",
			expected: `// Upon failure, the test [T] is marked as failed and stops execution.`,
		},
		{
			name:     "unknown package",
			pkg:      "unknown",
			expected: "",
		},
	})
}

type concatStringsCase struct {
	name     string
	input    []string
	expected string
}

func concatStringsCases() iter.Seq[concatStringsCase] {
	return slices.Values([]concatStringsCase{
		{
			name:     "empty strings",
			input:    []string{},
			expected: "",
		},
		{
			name:     "single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "multiple strings",
			input:    []string{"hello", " ", "world"},
			expected: "hello world",
		},
	})
}

type pathPartsCase struct {
	name     string
	input    string
	expected string
}

func pathPartsCases() iter.Seq[pathPartsCase] {
	return slices.Values([]pathPartsCase{
		{
			name:     "simple path",
			input:    "foo/bar/baz",
			expected: `"foo","bar","baz"`,
		},
		{
			name:     "absolute path",
			input:    "/usr/local/bin",
			expected: `"usr","local","bin"`,
		},
		{
			name:     "single component",
			input:    "test",
			expected: `"test"`,
		},
	})
}

type relocateCase struct {
	name     string
	input    string
	pkg      string
	expected string
}

func relocateCases() iter.Seq[relocateCase] {
	return slices.Values([]relocateCase{
		{
			name:     "empty package does nothing",
			input:    "ErrTest{}, CollectT{}, True(c)",
			pkg:      "",
			expected: "ErrTest{}, CollectT{}, True(c)",
		},
		{
			name:     "relocates ErrTest",
			input:    "ErrTest{}",
			pkg:      "assert",
			expected: "assert.ErrTest{}",
		},
		{
			name:     "relocates CollectT",
			input:    "CollectT{}",
			pkg:      "assert",
			expected: "assert.CollectT{}",
		},
		{
			name:     "relocates True(c)",
			input:    "True(c)",
			pkg:      "assert",
			expected: "assert.True(c)",
		},
		{
			name:     "does not relocate with dot prefix",
			input:    "t.ErrTest{}",
			pkg:      "assert",
			expected: "t.ErrTest{}",
		},
		{
			name:     "relocates all three",
			input:    "ErrTest{}, CollectT{}, True(c)",
			pkg:      "assert",
			expected: "assert.ErrTest{}, assert.CollectT{}, assert.True(c)",
		},
	})
}

type sourceLinkCase struct {
	name     string
	baseURL  string
	pos      *token.Position
	expected string
}

func sourceLinkCases() iter.Seq[sourceLinkCase] {
	return slices.Values([]sourceLinkCase{
		{
			name:     "nil position",
			baseURL:  githubRepo,
			pos:      nil,
			expected: "",
		},
		{
			name:    "valid position",
			baseURL: githubRepo,
			pos: &token.Position{
				Filename: "/home/user/go/src/github.com/go-openapi/testify/internal/assertions/equal.go",
				Line:     42,
			},
			expected: githubRepo + "/blob/master/internal/assertions/equal.go#L42",
		},
	})
}

type quoteCase struct {
	name     string
	input    string
	expected string
}

func quoteCases() iter.Seq[quoteCase] {
	return slices.Values([]quoteCase{
		{
			name:     "simple string",
			input:    "hello",
			expected: `"hello"`,
		},
		{
			name:     "string with quotes",
			input:    `he said "hello"`,
			expected: `"he said \"hello\""`,
		},
		{
			name:     "empty string",
			input:    "",
			expected: `""`,
		},
	})
}

type titleizeCase struct {
	name     string
	input    string
	expected string
}

func titleizeCases() iter.Seq[titleizeCase] {
	return slices.Values([]titleizeCase{
		{
			name:     "lowercase",
			input:    "hello world",
			expected: "Hello World",
		},
		{
			name:     "mixed case keeps original case",
			input:    "hELLo WoRLD",
			expected: "HELLo WoRLD",
		},
		{
			name:     "upper-cased keeps original case",
			input:    "HTTP",
			expected: "HTTP",
		},
		{
			name:     "already titlecase",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	})
}

type godocbadgeCase struct {
	name        string
	input       string
	expected    string
	expectError bool
}

func godocbadgeCases() iter.Seq[godocbadgeCase] {
	return slices.Values([]godocbadgeCase{
		{
			name:        "valid URL",
			input:       "https://pkg.go.dev/github.com/go-openapi/testify/v2",
			expected:    "https://pkg.go.dev/badge/github.com/go-openapi/testify/v2",
			expectError: false,
		},
		{
			name:        "invalid URL",
			input:       "://invalid",
			expected:    "",
			expectError: true,
		},
	})
}

// TestSlugize verifies slugize converts names to markdown slugs.
func TestSlugize(t *testing.T) {
	t.Parallel()

	for tt := range slugizeCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := slugize(tt.input); got != tt.expected {
				t.Errorf("slugize(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

type slugizeCase struct {
	name     string
	input    string
	expected string
}

func slugizeCases() iter.Seq[slugizeCase] {
	return slices.Values([]slugizeCase{
		{
			name:     "simple name",
			input:    "Equal",
			expected: "equal",
		},
		{
			name:     "with dots",
			input:    "assert.Equal",
			expected: "assert-equal",
		},
		{
			name:     "with underscores",
			input:    "my_function",
			expected: "my-function",
		},
		{
			name:     "with spaces",
			input:    "My Function",
			expected: "my-function",
		},
		{
			name:     "with brackets",
			input:    "Greater[T cmp.Ordered]",
			expected: "greatert-cmp-ordered",
		},
		{
			name:     "with colons and tabs",
			input:    "section:\ttitle",
			expected: "section--title",
		},
		{
			name:     "with commas and tildes",
			input:    "A, B~C",
			expected: "a-bc",
		},
	})
}

// TestBlockquote verifies blockquote formatting.
func TestBlockquote(t *testing.T) {
	t.Parallel()

	for tt := range blockquoteCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := blockquote(tt.input); got != tt.expected {
				t.Errorf("blockquote(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

type blockquoteCase struct {
	name     string
	input    string
	expected string
}

func blockquoteCases() iter.Seq[blockquoteCase] {
	return slices.Values([]blockquoteCase{
		{
			name:     "single line",
			input:    "Hello world",
			expected: "> Hello world",
		},
		{
			name:     "line ending with period",
			input:    "Hello world.",
			expected: "> Hello world.\n>",
		},
		{
			name:     "multi-line",
			input:    "Line 1.\nLine 2",
			expected: "> Line 1.\n>\n> Line 2",
		},
	})
}

// TestHugoDelimiters verifies Hugo template delimiter functions.
func TestHugoDelimiters(t *testing.T) {
	t.Parallel()

	if got := hugoopen(); got != "{{" {
		t.Errorf("hugoopen() = %q, want %q", got, "{{")
	}
	if got := hugoclose(); got != "}}" {
		t.Errorf("hugoclose() = %q, want %q", got, "}}")
	}
}

// TestShouldLineFeed verifies conditional line feed.
func TestShouldLineFeed(t *testing.T) {
	t.Parallel()

	if got := shouldLineFeed(true); got != "" {
		t.Errorf("shouldLineFeed(true) = %q, want empty", got)
	}
	if got := shouldLineFeed(false); got != "\n" {
		t.Errorf(`shouldLineFeed(false) = %q, want "\n"`, got)
	}
}

// TestTestSetup verifies test parameterization for all variants.
func TestTestSetup(t *testing.T) {
	t.Parallel()

	for tt := range testSetupCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fn := model.Function{
				Name:    "Equal",
				UseMock: tt.useMock,
			}

			result := testSetup(fn, tt.variant, tt.receiver)

			if result.TestCall != tt.expectedCall {
				t.Errorf("TestCall = %q, want %q", result.TestCall, tt.expectedCall)
			}
			if result.TestMock != tt.expectedMock {
				t.Errorf("TestMock = %q, want %q", result.TestMock, tt.expectedMock)
			}
			if result.TestErrorPrefix != tt.expectedPrefix {
				t.Errorf("TestErrorPrefix = %q, want %q", result.TestErrorPrefix, tt.expectedPrefix)
			}
			if result.TestMsg != tt.expectedMsg {
				t.Errorf("TestMsg = %q, want %q", result.TestMsg, tt.expectedMsg)
			}
		})
	}
}

type testSetupCase struct {
	name           string
	variant        string
	receiver       string
	useMock        string
	expectedCall   string
	expectedMock   string
	expectedPrefix string
	expectedMsg    string
}

func testSetupCases() iter.Seq[testSetupCase] {
	return slices.Values([]testSetupCase{
		{
			name:           "assertions variant",
			variant:        "assertions",
			useMock:        "mockT",
			expectedCall:   "Equal(mock, ",
			expectedMock:   "mock := new(mockT)",
			expectedPrefix: "Equal",
			expectedMsg:    "",
		},
		{
			name:           "format variant",
			variant:        "format",
			useMock:        "mockT",
			expectedCall:   "Equalf(mock, ",
			expectedMock:   "mock := new(mockT)",
			expectedPrefix: "Equalf",
			expectedMsg:    "test message",
		},
		{
			name:           "forward variant",
			variant:        "forward",
			receiver:       "Assertions",
			useMock:        "mockT",
			expectedCall:   "a.Equal(",
			expectedMock:   "mock := new(mockT)\na := New(mock)",
			expectedPrefix: "Assertions.Equal",
			expectedMsg:    "",
		},
		{
			name:           "forward-format variant",
			variant:        "forward-format",
			receiver:       "Assertions",
			useMock:        "mockT",
			expectedCall:   "a.Equalf(",
			expectedMock:   "mock := new(mockT)\na := New(mock)",
			expectedPrefix: "Assertions.Equalf",
			expectedMsg:    "test message",
		},
	})
}
