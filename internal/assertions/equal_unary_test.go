package assertions

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"testing"
)

func TestEqualUnaryErrorMessages(t *testing.T) {
	// error messages validation
	for tc := range equalEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mock := new(captureT)

			res := Empty(mock, tc.value)
			mock.checkResultAndErrMsg(t, res, tc.expectedResult, tc.expectedErrMsg)
		})
	}
}

// Unary assertion tests (Nil, NotNil, Empty, NotEmpty).
func TestEqualUnaryAssertions(t *testing.T) {
	t.Parallel()

	for tc := range unifiedUnaryCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Run("with Nil", testUnaryAssertion(tc, nilKind, Nil))
			t.Run("with NotNil", testUnaryAssertion(tc, notNilKind, NotNil))
			t.Run("with Empty", testUnaryAssertion(tc, emptyKind, Empty))
			t.Run("with NotEmpty", testUnaryAssertion(tc, notEmptyKind, NotEmpty))
		})
	}
}

type unaryTestCase struct {
	name     string
	object   any
	category objectCategory
}

func unifiedUnaryCases() iter.Seq[unaryTestCase] {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	x := 1
	xP := &x
	z := 0
	zP := &z
	var arr [1]int

	type TString string
	type TStruct struct {
		x int
	}
	type FStruct struct {
		x func()
	}

	return slices.Values([]unaryTestCase{
		// Nil category
		{"nil/nil-ptr", (*int)(nil), nilCategory},
		{"nil/nil-slice", []int(nil), nilCategory},
		{"nil/nil-interface", (any)(nil), nilCategory},
		{"nil/nil-struct-ptr", (*struct{})(nil), nilCategory},

		// Empty non-nil category
		{"empty/slice", []int{}, emptyNonNil},
		{"empty/string", "", emptyNonNil},
		{"empty/zero-int", 0, emptyNonNil},
		{"empty/zero-bool", false, emptyNonNil},
		{"empty/channel", make(chan struct{}), emptyNonNil},
		{"empty/zero-struct", TStruct{}, emptyNonNil},
		{"empty/aliased-string", TString(""), emptyNonNil},
		{"empty/zero-array", [1]int{}, emptyNonNil},
		{"empty/zero-ptr", zP, emptyNonNil},
		{"empty/zero-struct-ptr", &TStruct{}, emptyNonNil},
		{"empty/zero-array-ptr", &arr, emptyNonNil},
		{"empty/rune", '\u0000', emptyNonNil},
		{"empty/complex", 0i, emptyNonNil},
		{"empty/error", errors.New(""), emptyNonNil},
		{"empty/struct-with-func", FStruct{x: nil}, emptyNonNil},

		// Non-empty comparable category
		{"non-empty/int", 42, nonEmptyComparable},
		{"non-empty/rune", 'A', nonEmptyComparable},
		{"non-empty/string", "hello", nonEmptyComparable},
		{"non-empty/bool", true, nonEmptyComparable},
		{"non-empty/slice", []int{1}, nonEmptyComparable},
		{"non-empty/channel", chWithValue, nonEmptyComparable},
		{"non-empty/struct", TStruct{x: 1}, nonEmptyComparable},
		{"non-empty/aliased-string", TString("abc"), nonEmptyComparable},
		{"non-empty/ptr", xP, nonEmptyComparable},
		{"non-empty/array", [1]int{42}, nonEmptyComparable},

		// Non-empty non-comparable category
		{"non-empty/error", errors.New("something"), nonEmptyNonComparable},
		{"non-empty/slice-error", []error{errors.New("")}, nonEmptyNonComparable},
		{"non-empty/slice-nil-error", []error{nil}, nonEmptyNonComparable},
		{"non-empty/slice-zero", []int{0}, nonEmptyNonComparable},
		{"non-empty/slice-nil", []*int{nil}, nonEmptyNonComparable},
		{"non-empty/struct-with-func", FStruct{x: func() {}}, nonEmptyNonComparable},
	})
}

type unaryAssertionKind int

const (
	nilKind unaryAssertionKind = iota
	notNilKind
	emptyKind
	notEmptyKind
)

type objectCategory int

const (
	nilCategory objectCategory = iota
	emptyNonNil
	nonEmptyComparable
	nonEmptyNonComparable
)

// expectedStatusForUnaryAssertion returns the expected semantics for a given assertion (Nil, Empty, ...)
// and a given category of input.
func expectedStatusForUnaryAssertion(kind unaryAssertionKind, category objectCategory) bool {
	switch kind {
	case nilKind:
		return category == nilCategory
	case notNilKind:
		return category != nilCategory
	case emptyKind:
		return category == nilCategory || category == emptyNonNil
	case notEmptyKind:
		return category == nonEmptyComparable || category == nonEmptyNonComparable
	default:
		panic(fmt.Errorf("test case configuration error: invalid unaryAssertionKind: %d", kind))
	}
}

func testUnaryAssertion(tc unaryTestCase, kind unaryAssertionKind, unaryAssertion func(T, any, ...any) bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := unaryAssertion(mock, tc.object)
		shouldPass := expectedStatusForUnaryAssertion(kind, tc.category)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

type equalEmptyCase struct {
	name           string
	value          any
	expectedResult bool
	expectedErrMsg string
}

func equalEmptyCases() iter.Seq[equalEmptyCase] {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	// var tiP *time.Time
	// var tiNP time.Time
	// var s *string
	// var f *os.File
	// sP := &s
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	return slices.Values([]equalEmptyCase{
		{
			name:           "Non Empty string is not empty",
			value:          "something",
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was something\n",
		},
		{
			name:           "Non nil object is not empty",
			value:          errors.New("something"),
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was something\n",
		},
		{
			name:           "Non empty string array is not empty",
			value:          []string{"something"},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was [something]\n",
		},
		{
			name:           "Non-zero int value is not empty",
			value:          1,
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was 1\n",
		},
		{
			name:           "True value is not empty",
			value:          true,
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was true\n",
		},
		{
			name:           "Channel with values is not empty",
			value:          chWithValue,
			expectedResult: false,
			expectedErrMsg: fmt.Sprintf("Should be empty, but was %v\n", chWithValue),
		},
		{
			name:           "struct with initialized values is empty",
			value:          TStruct{x: 1},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was {1}\n",
		},
		{
			name:           "non-empty aliased string is empty",
			value:          TString("abc"),
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was abc\n",
		},
		{
			name:           "ptr to non-nil value is not empty",
			value:          xP,
			expectedResult: false,
			expectedErrMsg: fmt.Sprintf("Should be empty, but was %p\n", xP),
		},
		{
			name:           "array is not state",
			value:          [1]int{42},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was [42]\n",
		},

		// Here are some edge cases
		{
			name:           "string with only spaces is not empty",
			value:          "   ",
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was    \n", // Proposal for enhancement: FIX THIS strange error message
		},
		{
			name:           "string with a line feed is not empty",
			value:          "\n",
			expectedResult: false,
			// Proposal for enhancement: This is the exact same error message as for an empty string
			expectedErrMsg: "Should be empty, but was \n", // Proposal for enhancement: FIX THIS strange error message
		},
		{
			name:           "string with only tabulation and lines feed is not empty",
			value:          "\n\t\n",
			expectedResult: false,
			// Proposal for enhancement: The line feeds and tab are not helping to spot what is expected
			expectedErrMsg: "" + // this syntax is used to show how errors are reported.
				"Should be empty, but was \n" +
				"\t\n",
		},
		{
			name:           "string with trailing lines feed is not empty",
			value:          "foo\n\n",
			expectedResult: false,
			// Proposal for enhancement: it's not clear if one or two lines feed are expected
			expectedErrMsg: "Should be empty, but was foo\n\n",
		},
		{
			name:           "string with leading and trailing tabulation and lines feed is not empty",
			value:          "\n\nfoo\t\n\t\n",
			expectedResult: false,
			// Proposal for enhancement: The line feeds and tab are not helping to figure what is expected
			expectedErrMsg: "" +
				"Should be empty, but was \n" +
				"\n" +
				"foo\t\n" +
				"\t\n",
		},
		{
			name:           "non-printable character is not empty",
			value:          "\u00a0", // NO-BREAK SPACE UNICODE CHARACTER
			expectedResult: false,
			// Proposal for enhancement: here you cannot figure out what is expected
			expectedErrMsg: "Should be empty, but was \u00a0\n",
		},
		// Here we are testing there is no error message on success
		{
			name:           "Empty string is empty",
			value:          "",
			expectedResult: true,
			expectedErrMsg: "",
		},
	})
}
