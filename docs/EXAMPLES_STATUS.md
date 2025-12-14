# Test Examples Status

This document tracks which assertion functions have test examples added.

## ✅ Completed (with values)

### boolean.go
- [x] True - `success: 1 == 1`, `failure: 1 == 0`
- [x] False - `success: 1 == 0`, `failure: 1 == 1`

### collection.go
- [x] Len - `success: []string{"A","B"}, 2`, `failure: []string{"A","B"}, 1`
- [x] Contains - `success: []string{"A","B"}, "A"`, `failure: []string{"A","B"}, "C"`
- [x] NotContains - `success: []string{"A","B"}, "C"`, `failure: []string{"A","B"}, "B"`

### equal.go
- [x] Equal - `success: 123, 123`, `failure: 123, 456`
- [x] EqualValues - `success: uint32(123), int32(123)`, `failure: uint32(123), int32(456)`

### error.go
- [x] NoError - `success: nil`, `failure: ErrTest`
- [x] Error - `success: ErrTest`, `failure: nil`
- [x] EqualError - `success: ErrTest, "assert.ErrTest general error for testing"`, `failure: ErrTest, "wrong error message"`
- [x] ErrorContains - `success: ErrTest, "general error"`, `failure: ErrTest, "not in message"`

### yaml.go
- [x] YAMLEq - `panic:` (with message about yaml feature)

## 📝 Completed (with TODO placeholders)

### equal.go
- [x] Same - TODO: add pointer test values
- [x] NotSame - TODO: add different pointer test values
- [x] EqualExportedValues - TODO: add struct test values

## ⏳ Remaining Functions (need examples)

### equal.go
- [ ] NotEqual
- [ ] NotEqualValues
- [ ] Exactly

### collection.go
- [ ] Empty
- [ ] NotEmpty
- [ ] Subset
- [ ] NotSubset
- [ ] ElementsMatch
- [ ] NotElementsMatch

### compare.go
- [ ] Greater
- [ ] GreaterOrEqual
- [ ] Less
- [ ] LessOrEqual
- [ ] Positive
- [ ] Negative

### error.go
- [ ] ErrorIs
- [ ] ErrorAs
- [ ] NotErrorIs

### object.go
- [ ] Nil
- [ ] NotNil
- [ ] IsType
- [ ] NotNil
- [ ] Zero
- [ ] NotZero

### panic.go
- [ ] Panics
- [ ] PanicsWithValue
- [ ] PanicsWithError
- [ ] NotPanics

### string.go
- [ ] Regexp
- [ ] NotRegexp

### number.go
- [ ] InDelta
- [ ] InDeltaSlice
- [ ] InDeltaMapValues
- [ ] InEpsilon
- [ ] InEpsilonSlice

### condition.go
- [ ] Condition

### http.go
- [ ] HTTPSuccess
- [ ] HTTPRedirect
- [ ] HTTPError
- [ ] HTTPStatusCode
- [ ] HTTPBodyContains
- [ ] HTTPBodyNotContains

### json.go
- [ ] JSONEq

### file.go
- [ ] FileExists
- [ ] NoFileExists
- [ ] DirExists
- [ ] NoDirExists

### time.go
- [ ] WithinDuration
- [ ] WithinRange

### type.go
- [ ] Implements

### testing.go
- [ ] Never
- [ ] Eventually
- [ ] EventuallyWithT

### assertion.go
- [ ] Fail
- [ ] FailNow

## Template for Adding Examples

```go
// FunctionName does something.
//
// Usage:
//
//	assert.FunctionName(t, args...)
//
// Examples:
//
//	success: <values that should succeed>
//	failure: <values that should fail>
//	panic: <values that cause panic>
//	       <panic assertion message>
func FunctionName(...) bool {
```

## Notes

- Use actual Go values that can be directly inserted into test code
- For complex types, use TODO comments
- success/failure are required, panic is optional
- All values will be injected directly into generated test code
