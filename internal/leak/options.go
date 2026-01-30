//nolint:ireturn,inamedparam // temporary setup
package leak

import (
	"strings"
	"time"
)

// Option lets users specify custom verifications.
type Option interface {
	apply(*opts)
}

// We retry up to 20 times if we can't find the goroutine that
// we are looking for. In between each attempt, we will sleep for
// a short while to let any running goroutines complete.
const _defaultRetries = 20

const defaultMaxSleep = 100 * time.Millisecond

type opts struct {
	filters      []func(Stack) bool
	maxRetries   int
	maxSleep     time.Duration
	cleanup      func(int)
	runOnFailure bool
}

// implement apply so that opts struct itself can be used as
// an Option.
func (o *opts) apply(opts *opts) {
	opts.filters = o.filters
	opts.maxRetries = o.maxRetries
	opts.maxSleep = o.maxSleep
	opts.cleanup = o.cleanup
	opts.runOnFailure = o.runOnFailure
}

// optionFunc lets us easily write options without a custom type.
type optionFunc func(*opts)

func (f optionFunc) apply(opts *opts) { f(opts) }

// IgnoreTopFunction ignores any goroutines where the specified function
// is at the top of the  The function name should be fully qualified,
// e.g., go.uber.org/goleak.IgnoreTopFunction.
func IgnoreTopFunction(f string) Option {
	return addFilter(func(s Stack) bool {
		return s.FirstFunction() == f
	})
}

// IgnoreAnyFunction ignores goroutines where the specified function
// is present anywhere in the
//
// The function name must be fully qualified, e.g.,
//
//	go.uber.org/goleak.IgnoreAnyFunction
//
// For methods, the fully qualified form looks like:
//
//	go.uber.org/goleak.(*MyType).MyMethod
func IgnoreAnyFunction(f string) Option {
	return addFilter(func(s Stack) bool {
		return s.HasFunction(f)
	})
}

// IgnoreCreatedBy ignores any goroutines that where spawned from the
// specified function. The function name should be fully qualified, e.g.
// go.uber.org/goleak.IgnoreCreatedBy.
func IgnoreCreatedBy(f string) Option {
	return addFilter(func(s Stack) bool {
		return s.CreatedBy() == f
	})
}

// Cleanup sets up a cleanup function that will be executed at the
// end of the leak check.
// When passed to [VerifyTestMain], the exit code passed to cleanupFunc
// will be set to the exit code of TestMain.
// When passed to [VerifyNone], the exit code will be set to 0.
// This cannot be passed to [Find].
func Cleanup(cleanupFunc func(exitCode int)) Option {
	return optionFunc(func(opts *opts) {
		opts.cleanup = cleanupFunc
	})
}

// IgnoreCurrent records all current goroutines when the option is created, and ignores
// them in any future Find/Verify calls.
func IgnoreCurrent() Option {
	excludeIDSet := map[int]bool{}
	for _, s := range All() {
		excludeIDSet[s.ID()] = true
	}
	return addFilter(func(s Stack) bool {
		return excludeIDSet[s.ID()]
	})
}

// RunOnFailure makes goleak look for leaking goroutines upon test failures.
// By default goleak only looks for leaking goroutines when tests succeed.
func RunOnFailure() Option {
	return optionFunc(func(opts *opts) {
		opts.runOnFailure = true
	})
}

/*
func maxSleep(d time.Duration) Option {
	return optionFunc(func(opts *opts) {
		opts.maxSleep = d
	})
}
*/

func addFilter(f func(Stack) bool) Option {
	return optionFunc(func(opts *opts) {
		opts.filters = append(opts.filters, f)
	})
}

func buildOpts(options ...Option) *opts {
	opts := &opts{
		maxRetries: _defaultRetries,
		maxSleep:   defaultMaxSleep,
	}
	opts.filters = append(opts.filters,
		isTestStack,
		isSyscallStack,
		isStdLibStack,
		isTraceStack,
	)
	for _, option := range options {
		option.apply(opts)
	}
	return opts
}

func (o *opts) filter(s Stack) bool {
	for _, filter := range o.filters {
		if filter(s) {
			return true
		}
	}
	return false
}

func (o *opts) retry(i int) bool {
	if i >= o.maxRetries {
		return false
	}

	d := min(time.Duration(int(time.Microsecond)<<uint(i)), o.maxSleep) //nolint:gosec // conversion is fine
	time.Sleep(d)
	return true
}

// isTestStack is a default filter installed to automatically skip goroutines
// that the testing package runs while the user's tests are running.
func isTestStack(s Stack) bool {
	// Until go1.7, the main goroutine ran RunTests, which started
	// the test in a separate goroutine and waited for that test goroutine
	// to end by waiting on a channel.
	// Since go1.7, a separate goroutine is started to wait for signals.
	// T.Parallel is for parallel tests, which are blocked until all serial
	// tests have run with T.Parallel at the top of the
	// testing.runFuzzTests is for fuzz testing, it's blocked until the test
	// function with all seed corpus have run.
	// testing.runFuzzing is for fuzz testing, it's blocked until a failing
	// input is found.
	switch s.FirstFunction() {
	case "testing.RunTests", "testing.(*T).Run", "testing.(*T).Parallel", "testing.runFuzzing", "testing.runFuzzTests":
		// In pre1.7 and post-1.7, background goroutines started by the testing
		// package are blocked waiting on a channel.
		return strings.HasPrefix(s.State(), "chan receive")
	}
	return false
}

func isSyscallStack(s Stack) bool {
	// Typically runs in the background when code uses CGo:
	// https://github.com/golang/go/issues/16714
	return s.HasFunction("runtime.goexit") && strings.HasPrefix(s.State(), "syscall")
}

func isStdLibStack(s Stack) bool {
	// Importing os/signal starts a background goroutine.
	// The name of the function at the top has changed between versions.
	if f := s.FirstFunction(); f == "os/signal.signal_recv" || f == "os/signal.loop" {
		return true
	}

	// Using signal.Notify will start a runtime goroutine.
	return s.HasFunction("runtime.ensureSigM")
}

func isTraceStack(s Stack) bool {
	return s.HasFunction("runtime.ReadTrace")
}
