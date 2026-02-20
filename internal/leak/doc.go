// Package leak provide goroutine leak detection
// using [pprof] labels instead of stack-trace heuristics.
//
// Instead of parsing [runtime.Stack] and filtering known
// system goroutines, we label the tested function's goroutine using
// [pprof.Do].
//
// Child goroutines inherit the label automatically. After the
// function returns, any goroutines still carrying our label are leaks.
//
// This approach makes leak detection immune to goroutines running concurrently
// to the leak check (e.g. parallel tests) and immune to background goroutines
// started by other programs (e.g. database connections, http or grpc pools...).
package leak
