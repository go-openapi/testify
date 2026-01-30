//nolint:inamedparam // temporary setup
package leak

import (
	"errors"
	"fmt"
)

// TestingT is the minimal subset of testing.TB that we use.
type TestingT interface {
	Error(...any)
}

// filterStacks will filter any stacks excluded by the given opts.
// filterStacks modifies the passed in stacks slice.
func filterStacks(stacks []Stack, skipID int, opts *opts) []Stack {
	filtered := stacks[:0]
	for _, stack := range stacks {
		// Always skip the running goroutine.
		if stack.ID() == skipID {
			continue
		}
		// Run any default or user-specified filters.
		if opts.filter(stack) {
			continue
		}
		filtered = append(filtered, stack)
	}
	return filtered
}

// Find looks for extra goroutines, and returns a descriptive error if
// any are found.
func Find(options ...Option) error {
	cur := Current().ID()

	opts := buildOpts(options...)
	if opts.cleanup != nil {
		return errors.New("Cleanup can only be passed to VerifyNone or VerifyTestMain")
	}
	if opts.runOnFailure {
		return errors.New("RunOnFailure can only be passed to VerifyTestMain")
	}
	var stacks []Stack
	retry := true
	for i := 0; retry; i++ {
		stacks = filterStacks(All(), cur, opts)

		if len(stacks) == 0 {
			return nil
		}
		retry = opts.retry(i)
	}

	return fmt.Errorf("found unexpected goroutines:\n%s", stacks)
}

/*
type testHelper interface {
	Helper()
}

// VerifyNone marks the given TestingT as failed if any extra goroutines are
// found by Find. This is a helper method to make it easier to integrate in
// tests by doing:
//
//	defer VerifyNone(t)
//
// VerifyNone is currently incompatible with t.Parallel because it cannot
// associate specific goroutines with specific tests. Thus, non-leaking
// goroutines from other tests running in parallel could fail this check.
// If you need to run tests in parallel, use [VerifyTestMain] instead,
// which will verify that no leaking goroutines exist after ALL tests finish.
func VerifyNone(t TestingT, options ...Option) {
	opts := buildOpts(options...)
	var cleanup func(int)
	cleanup, opts.cleanup = opts.cleanup, nil

	if h, ok := t.(testHelper); ok {
		// Mark this function as a test helper, if available.
		h.Helper()
	}

	if err := Find(opts); err != nil {
		t.Error(err)
	}

	if cleanup != nil {
		cleanup(0)
	}
}
*/
