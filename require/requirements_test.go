package require

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

// AssertionTesterInterface defines an interface to be used for testing assertion methods.
type AssertionTesterInterface interface {
	TestMethod()
}

// AssertionTesterConformingObject is an object that conforms to the AssertionTesterInterface interface.
type AssertionTesterConformingObject struct {
}

func (a *AssertionTesterConformingObject) TestMethod() {
}

// AssertionTesterNonConformingObject is an object that does not conform to the AssertionTesterInterface interface.
type AssertionTesterNonConformingObject struct {
}

type MockT struct {
	Failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (MockT) Helper() {}

func (t *MockT) FailNow() {
	t.Failed = true
}

func (t *MockT) Errorf(format string, args ...any) {
	_, _ = format, args
}

func TestImplements(t *testing.T) {
	t.Parallel()

	Implements(t, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject))

	mockT := new(MockT)
	Implements(mockT, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject))
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestIsType(t *testing.T) {
	t.Parallel()

	IsType(t, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject))

	mockT := new(MockT)
	IsType(mockT, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject))
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()

	Equal(t, 1, 1)

	mockT := new(MockT)
	Equal(mockT, 1, 2)
	if !mockT.Failed {
		t.Error("Check should fail")
	}

}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	NotEqual(t, 1, 2)
	mockT := new(MockT)
	NotEqual(mockT, 2, 2)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestExactly(t *testing.T) {
	t.Parallel()

	a := float32(1)
	b := float32(1)
	c := float64(1)

	Exactly(t, a, b)

	mockT := new(MockT)
	Exactly(mockT, a, c)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNotNil(t *testing.T) {
	t.Parallel()

	NotNil(t, new(AssertionTesterConformingObject))

	mockT := new(MockT)
	NotNil(mockT, nil)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNil(t *testing.T) {
	t.Parallel()

	Nil(t, nil)

	mockT := new(MockT)
	Nil(mockT, new(AssertionTesterConformingObject))
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestTrue(t *testing.T) {
	t.Parallel()

	True(t, true)

	mockT := new(MockT)
	True(mockT, false)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestFalse(t *testing.T) {
	t.Parallel()

	False(t, false)

	mockT := new(MockT)
	False(mockT, true)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestContains(t *testing.T) {
	t.Parallel()

	Contains(t, "Hello World", "Hello")

	mockT := new(MockT)
	Contains(mockT, "Hello World", "Salut")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNotContains(t *testing.T) {
	t.Parallel()

	NotContains(t, "Hello World", "Hello!")

	mockT := new(MockT)
	NotContains(mockT, "Hello World", "Hello")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestPanics(t *testing.T) {
	t.Parallel()

	Panics(t, func() {
		panic("Panic!")
	})

	mockT := new(MockT)
	Panics(mockT, func() {})
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNotPanics(t *testing.T) {
	t.Parallel()

	NotPanics(t, func() {})

	mockT := new(MockT)
	NotPanics(mockT, func() {
		panic("Panic!")
	})
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNoError(t *testing.T) {
	t.Parallel()

	NoError(t, nil)

	mockT := new(MockT)
	NoError(mockT, someError())
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	Error(t, someError())

	mockT := new(MockT)
	Error(mockT, nil)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestErrorContains(t *testing.T) {
	t.Parallel()

	ErrorContains(t, fmt.Errorf("some error: another error: %w", errSentinel), "some error")

	mockT := new(MockT)
	ErrorContains(mockT, someError(), "different error")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestEqualError(t *testing.T) {
	t.Parallel()

	EqualError(t, someError(), "some error: test error")

	mockT := new(MockT)
	EqualError(mockT, someError(), "Not some error")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	Empty(t, "")

	mockT := new(MockT)
	Empty(mockT, "x")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNotEmpty(t *testing.T) {
	t.Parallel()

	NotEmpty(t, "x")

	mockT := new(MockT)
	NotEmpty(mockT, "")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestWithinDuration(t *testing.T) {
	t.Parallel()

	a := time.Now()
	b := a.Add(10 * time.Second)

	WithinDuration(t, a, b, 15*time.Second)

	mockT := new(MockT)
	WithinDuration(mockT, a, b, 5*time.Second)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestInDelta(t *testing.T) {
	t.Parallel()

	InDelta(t, 1.001, 1, 0.01)

	mockT := new(MockT)
	InDelta(mockT, 1, 2, 0.5)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestZero(t *testing.T) {
	t.Parallel()

	Zero(t, "")

	mockT := new(MockT)
	Zero(mockT, "x")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestNotZero(t *testing.T) {
	t.Parallel()

	NotZero(t, "x")

	mockT := new(MockT)
	NotZero(mockT, "")
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_EqualSONString(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, simpleJSONObject, simpleJSONObject)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEq_EquivalentButNotEqual(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, simpleJSONObject, simpleJSONObjectReversed)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEq_HashOfArraysAndHashes(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, nestedJSONObject, nestedJSONObjectShuffled)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEq_Array(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, simpleJSONNested, simpleJSONNestedReversed)
	if mockT.Failed {
		t.Error("Check should pass")
	}
}

func TestJSONEq_HashAndArrayNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, simpleJSONNested, simpleJSONNestedNotEq)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_HashesNotEquivalent(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, fooBarObject, simpleJSONObjectReversed)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_ActualIsNotJSON(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, fooBarObject, notJSONString)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_ExpectedIsNotJSON(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, notJSONString, simpleJSONObjectReversed)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_ExpectedAndActualNotJSON(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, notJSONString, notJSONString)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestJSONEq_ArraysOfDifferentOrder(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)
	JSONEq(mockT, simpleJSONArray, simpleJSONArrayReversed)
	if !mockT.Failed {
		t.Error("Check should fail")
	}
}

func TestComparisonAssertionFunc(t *testing.T) {
	t.Parallel()

	type iface interface {
		Name() string
	}

	tests := []struct {
		name      string
		expect    any
		got       any
		assertion ComparisonAssertionFunc
	}{
		{"implements", (*iface)(nil), t, Implements},
		{"isType", (*testing.T)(nil), t, IsType},
		{"equal", t, t, Equal},
		{"equalValues", t, t, EqualValues},
		{"exactly", t, t, Exactly},
		{"notEqual", t, nil, NotEqual},
		{"NotEqualValues", t, nil, NotEqualValues},
		{"notContains", []int{1, 2, 3}, 4, NotContains},
		{"subset", []int{1, 2, 3, 4}, []int{2, 3}, Subset},
		{"notSubset", []int{1, 2, 3, 4}, []int{0, 3}, NotSubset},
		{"elementsMatch", []byte("abc"), []byte("bac"), ElementsMatch},
		{"regexp", "^t.*y$", "testify", Regexp},
		{"notRegexp", "^t.*y$", "Testify", NotRegexp},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.expect, tt.got)
		})
	}
}

func TestValueAssertionFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		value     any
		assertion ValueAssertionFunc
	}{
		{"notNil", true, NotNil},
		{"nil", nil, Nil},
		{"empty", []int{}, Empty},
		{"notEmpty", []int{1}, NotEmpty},
		{"zero", false, Zero},
		{"notZero", 42, NotZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.value)
		})
	}
}

func TestBoolAssertionFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		value     bool
		assertion BoolAssertionFunc
	}{
		{"true", true, True},
		{"false", false, False},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.value)
		})
	}
}

func TestErrorAssertionFunc(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		err       error
		assertion ErrorAssertionFunc
	}{
		{"noError", nil, NoError},
		{"error", fmt.Errorf("whoops: %w", errSentinel), Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, tt.err)
		})
	}
}

func TestEventuallyWithTFalse(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)

	condition := func(collect *assert.CollectT) {
		True(collect, false)
	}

	EventuallyWithT(mockT, condition, 100*time.Millisecond, 20*time.Millisecond)
	True(t, mockT.Failed, "Check should fail")
}

func TestEventuallyWithTTrue(t *testing.T) {
	t.Parallel()

	mockT := new(MockT)

	counter := 0
	condition := func(collect *assert.CollectT) {
		defer func() {
			counter++
		}()
		True(collect, counter == 1)
	}

	EventuallyWithT(mockT, condition, 100*time.Millisecond, 20*time.Millisecond)
	False(t, mockT.Failed, "Check should pass")
	Equal(t, 2, counter, "Condition is expected to be called 2 times")
}

var errSentinel = errors.New("test error")

func someError() error {
	return fmt.Errorf("some error: %w", errSentinel)
}
