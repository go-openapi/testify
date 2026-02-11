// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

// stretchrAssertPath and related constants define the import paths for
// stretchr/testify and their go-openapi/testify/v2 replacements.
const (
	stretchrAssertPath  = "github.com/stretchr/testify/assert"
	stretchrRequirePath = "github.com/stretchr/testify/require"
	stretchrRootPath    = "github.com/stretchr/testify"
	stretchrMockPath    = "github.com/stretchr/testify/mock"
	stretchrSuitePath   = "github.com/stretchr/testify/suite"
	stretchrHTTPPath    = "github.com/stretchr/testify/http"

	goopenapiAssertPath  = "github.com/go-openapi/testify/v2/assert"
	goopenapiRequirePath = "github.com/go-openapi/testify/v2/require"
	goopenapiRootPath    = "github.com/go-openapi/testify/v2"
	goopenapiYAMLEnable  = "github.com/go-openapi/testify/v2/enable/yaml"
)

// importRewrites maps stretchr import paths to go-openapi replacements.
var importRewrites = map[string]string{ //nolint:gochecknoglobals // lookup table
	stretchrAssertPath:  goopenapiAssertPath,
	stretchrRequirePath: goopenapiRequirePath,
	stretchrRootPath:    goopenapiRootPath,
}

// incompatibleImports are stretchr import paths with no go-openapi equivalent.
var incompatibleImports = map[string]string{ //nolint:gochecknoglobals // lookup table
	stretchrMockPath: "mock package is not available in go-openapi/testify/v2.\n" +
		"  → Use github.com/vektra/mockery or hand-written mocks.\n" +
		"  → See: https://go-openapi.github.io/testify/usage/migration/index.html#5-remove-suitemock-usage",
	stretchrSuitePath: "suite package is not available in go-openapi/testify/v2.\n" +
		"  → Use standard Go subtests (t.Run) and TestMain.\n" +
		"  → See: https://go-openapi.github.io/testify/usage/migration/index.html#5-remove-suitemock-usage",
	stretchrHTTPPath: "http package is not available in go-openapi/testify/v2.\n" +
		"  → Use net/http/httptest from the standard library.\n" +
		"  → See: https://go-openapi.github.io/testify/usage/migration/index.html#6-remove-use-of-the-testifyhttp-package",
}

// migrationRenames maps stretchr function names to their go-openapi equivalents
// where the names differ.
var migrationRenames = map[string]string{ //nolint:gochecknoglobals // lookup table
	"EventuallyWithT":  "EventuallyWith",
	"EventuallyWithTf": "EventuallyWithf",
	"NoDirExists":      "DirNotExists",
	"NoDirExistsf":     "DirNotExistsf",
	"NoFileExists":     "FileNotExists",
	"NoFileExistsf":    "FileNotExistsf",
}

// yamlFunctions lists assertion function names that use YAML, triggering
// injection of the enable/yaml import.
var yamlFunctions = map[string]bool{ //nolint:gochecknoglobals // lookup table
	"YAMLEq":       true,
	"YAMLEqf":      true,
	"YAMLEqT":      true,
	"YAMLEqTf":     true,
	"YAMLEqBytes":  true,
	"YAMLEqBytesf": true,
}

// constraintKind classifies the type constraint needed for a generic upgrade.
type constraintKind int

const (
	constraintComparable     constraintKind = iota
	constraintDeepComparable                // comparable AND == has same semantics as reflect.DeepEqual
	constraintOrdered
	constraintText
	constraintSignedNumeric
	constraintMeasurable
	constraintPointer
	constraintBoolean
	constraintRegExp
)

// containerKind classifies the container type for Contains upgrades.
type containerKind int

const (
	containerString containerKind = iota
	containerSlice
	containerMap
)

// upgradeRule defines how to upgrade a reflection-based assertion to a generic variant.
type upgradeRule struct {
	// target is the generic function name to upgrade to.
	target string
	// argConstraints defines the constraint required for each argument
	// (excluding t and msgAndArgs). For single-constraint functions, only
	// one entry is needed.
	argConstraints []constraintKind
	// sameType requires that the relevant arguments have identical types.
	sameType bool
	// containerUpgrade means the function dispatches to different generic
	// variants depending on the container type (string, slice, map).
	containerUpgrade bool
	// manualReview flags the upgrade as requiring manual review due to
	// signature changes (e.g., IsType → IsOfTypeT changes arg count).
	manualReview bool
}

// genericUpgrades maps reflection-based assertion names to their generic upgrade rules.
var genericUpgrades = map[string]upgradeRule{ //nolint:gochecknoglobals // lookup table
	// Equality — must be deep-comparable (== same as reflect.DeepEqual)
	// to preserve semantics. Pointer types are excluded because EqualT
	// uses == (address comparison) while Equal uses reflect.DeepEqual.
	"Equal": {
		target:         "EqualT",
		argConstraints: []constraintKind{constraintDeepComparable},
		sameType:       true,
	},
	"NotEqual": {
		target:         "NotEqualT",
		argConstraints: []constraintKind{constraintDeepComparable},
		sameType:       true,
	},

	// Comparison / ordering
	"Greater": {
		target:         "GreaterT",
		argConstraints: []constraintKind{constraintOrdered},
		sameType:       true,
	},
	"GreaterOrEqual": {
		target:         "GreaterOrEqualT",
		argConstraints: []constraintKind{constraintOrdered},
		sameType:       true,
	},
	"Less": {
		target:         "LessT",
		argConstraints: []constraintKind{constraintOrdered},
		sameType:       true,
	},
	"LessOrEqual": {
		target:         "LessOrEqualT",
		argConstraints: []constraintKind{constraintOrdered},
		sameType:       true,
	},

	// Numeric
	"Positive": {
		target:         "PositiveT",
		argConstraints: []constraintKind{constraintSignedNumeric},
	},
	"Negative": {
		target:         "NegativeT",
		argConstraints: []constraintKind{constraintSignedNumeric},
	},
	"InDelta": {
		target:         "InDeltaT",
		argConstraints: []constraintKind{constraintMeasurable},
		sameType:       true,
	},
	"InEpsilon": {
		target:         "InEpsilonT",
		argConstraints: []constraintKind{constraintMeasurable},
		sameType:       true,
	},

	// Container (dispatches to different targets)
	"Contains": {
		containerUpgrade: true,
	},
	"NotContains": {
		containerUpgrade: true,
	},

	// Collection
	"ElementsMatch": {
		target:         "ElementsMatchT",
		argConstraints: []constraintKind{constraintComparable},
		sameType:       true,
	},
	"Subset": {
		target:         "SliceSubsetT",
		argConstraints: []constraintKind{constraintComparable},
		sameType:       true,
	},

	// Ordering (slice)
	"IsIncreasing": {
		target:         "IsIncreasingT",
		argConstraints: []constraintKind{constraintOrdered},
	},
	"IsDecreasing": {
		target:         "IsDecreasingT",
		argConstraints: []constraintKind{constraintOrdered},
	},
	"IsNonIncreasing": {
		target:         "IsNonIncreasingT",
		argConstraints: []constraintKind{constraintOrdered},
	},
	"IsNonDecreasing": {
		target:         "IsNonDecreasingT",
		argConstraints: []constraintKind{constraintOrdered},
	},

	// String / regex
	"Regexp": {
		target:         "RegexpT",
		argConstraints: []constraintKind{constraintRegExp, constraintText},
	},
	"NotRegexp": {
		target:         "NotRegexpT",
		argConstraints: []constraintKind{constraintRegExp, constraintText},
	},

	// JSON / YAML
	"JSONEq": {
		target:         "JSONEqT",
		argConstraints: []constraintKind{constraintText},
		sameType:       true,
	},
	"YAMLEq": {
		target:         "YAMLEqT",
		argConstraints: []constraintKind{constraintText},
		sameType:       true,
	},

	// Pointer
	"Same": {
		target:         "SameT",
		argConstraints: []constraintKind{constraintPointer},
		sameType:       true,
	},
	"NotSame": {
		target:         "NotSameT",
		argConstraints: []constraintKind{constraintPointer},
		sameType:       true,
	},

	// Type (manual review — arg count changes)
	"IsType": {
		target:         "IsOfTypeT",
		manualReview:   true,
		argConstraints: []constraintKind{},
	},

	// Boolean
	"True": {
		target:         "TrueT",
		argConstraints: []constraintKind{constraintBoolean},
	},
	"False": {
		target:         "FalseT",
		argConstraints: []constraintKind{constraintBoolean},
	},
}

// containerUpgradeTargets maps container kind to (funcName, notFuncName) pairs.
var containerUpgradeTargets = map[containerKind][2]string{ //nolint:gochecknoglobals // lookup table
	containerString: {"StringContainsT", "StringNotContainsT"},
	containerSlice:  {"SliceContainsT", "SliceNotContainsT"},
	containerMap:    {"MapContainsT", "MapNotContainsT"},
}

// skipReason describes why a generic upgrade was skipped for an assertion call.
type skipReason string

const (
	skipPointerSemantics           skipReason = "pointer type (== compares addresses, not values)"
	skipInterfaceField             skipReason = "struct contains pointer/interface fields"
	skipUnresolvableType           skipReason = "type not statically resolvable"
	skipInterfaceType              skipReason = "argument is interface{}/any"
	skipTypeMismatch               skipReason = "arguments have different types"
	skipNotOrdered                 skipReason = "type does not satisfy Ordered constraint"
	skipNotText                    skipReason = "type does not satisfy Text constraint"
	skipNotSignedNumeric           skipReason = "type does not satisfy SignedNumeric constraint"
	skipNotMeasurable              skipReason = "type does not satisfy Measurable constraint"
	skipNotComparable              skipReason = "type does not satisfy comparable constraint"
	skipNotBoolean                 skipReason = "type does not satisfy Boolean constraint"
	skipNotPointer                 skipReason = "type is not a pointer"
	skipNotSlice                   skipReason = "argument is not a slice type"
	skipNotRegExp                  skipReason = "type does not satisfy RegExp constraint"
	skipSliceElemNotDeepComparable skipReason = "slice element not deeply comparable"
	skipSliceElemNotOrdered        skipReason = "slice element not ordered"
	skipContainerTypeUnknown       skipReason = "container is not string, slice, or map"
)
