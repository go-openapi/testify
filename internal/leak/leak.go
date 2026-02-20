// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package leak

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

const labelKey = "testify-leak-check"

func init() { //nolint:gochecknoinits // this init check is justify by the use of an internal volatile API.
	// check that the profile API behaves as expected or panic.
	//
	// The exact format of the labels reported in the profile stack is not documented and not guaranteed
	// to remain stable across go versions. We panic here to detect as early as possible any go API change
	// so we can quickly adapt to a new format.
	//
	// Even though we don't parse the complete stack, our detection method remains sentitive to labels formatting,
	// e.g. "labels: {key:value}".
	var wg sync.WaitGroup
	blocker := make(chan struct{})
	id := uniqueLabel()
	labels := pprof.Labels(labelKey, id)
	pprof.Do(context.Background(), labels, func(_ context.Context) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks forever until cleanup
		}()
	})
	needle := buildNeedle(id)
	profile := captureProfile()
	match := extractLabeledBlocks(profile, needle)
	close(blocker)
	wg.Wait()
	if match == "" {
		panic("unrecognized goroutine profile format: go API has changed unexpectedly")
	}
}

// Leaked instruments the tested function with a [pprof] label.
//
// It waits briefly for goroutines to settle, then checks the goroutine profile
// for any goroutines still carrying our label.
//
// Returns the matching portion of the profile text if leaks are found,
// or the empty string if clean.
func Leaked(ctx context.Context, tested func()) string {
	id := uniqueLabel()
	labels := pprof.Labels(labelKey, id)

	didPanic, panicVal, panicStack := runWithPanicGuard(ctx, labels, tested)

	needle := buildNeedle(id)
	profile := captureProfile()
	match := extractLabeledBlocks(profile, needle)
	if match == "" {
		return "" // early exit: clean state
	}

	const (
		maxAttempts = 20
		maxWait     = 100 * time.Millisecond
		waitFactor  = 2
	)
	wait := time.Microsecond
	for range maxAttempts {
		time.Sleep(wait) // brief retry: goroutines may be mid-exit.
		profile := captureProfile()
		match := extractLabeledBlocks(profile, needle)
		if match == "" {
			return "" // clean
		}

		// retry — goroutine might still be exiting
		// wait exponential backoff, capped to maxWait
		wait = min(wait*waitFactor, maxWait)
	}

	return formatOutput(match, didPanic, panicVal, panicStack)
}

func uniqueLabel() string {
	var buf [16]byte

	_, _ = rand.Read(buf[:])

	return hex.EncodeToString(buf[:])
}

func runWithPanicGuard(ctx context.Context, labels pprof.LabelSet, tested func()) (panicked bool, panicVal any, stack []byte) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer func() {
			if r := recover(); r != nil {
				stack = debug.Stack()
				panicked = true
				panicVal = r
			}
		}()
		pprof.Do(ctx, labels, func(_ context.Context) {
			tested()
		})
	}()

	select {
	case <-ctx.Done():
	case <-done:
	}

	return panicked, panicVal, stack
}

// buildNeedle returns the byte sequence to search for in the debug=1 profile.
func buildNeedle(id string) []byte {
	// The Go runtime formats labels as: {"key":"value"}
	// with %q-quoted strings and no space after the colon.
	return fmt.Appendf(nil, "%q:%q", labelKey, id)
}

// captureProfile writes the goroutine profile with debug=1 to a buffer.
func captureProfile() []byte {
	var buf bytes.Buffer
	profile := pprof.Lookup("goroutine")
	if profile == nil {
		return nil
	}
	_ = profile.WriteTo(&buf, 1)

	return buf.Bytes()
}

// extractLabeledBlocks scans the profile text for blocks containing the needle.
// Returns all matching blocks concatenated, or empty string if none found.
func extractLabeledBlocks(profile []byte, needle []byte) string {
	if !bytes.Contains(profile, needle) {
		return ""
	}

	// Extract the blocks containing our label for diagnostic output.
	// In debug=1, each block starts with a count line like "1 @ 0x..."
	// and ends at the next count line or EOF.
	var result bytes.Buffer
	for block := range bytes.SplitSeq(profile, []byte("\n\n")) {
		if !bytes.Contains(block, needle) {
			continue
		}

		if result.Len() > 0 {
			result.WriteString("\n---\n")
		}

		result.Write(block)
	}

	return result.String()
}

func formatOutput(match string, didPanic bool, panicVal any, panicStack []byte) string {
	// Sample raw match:
	//     1 @ 0x47890e 0x410193 0x40fd12 0x530449 0x47ff01
	// # labels: {"testify-leak-check":"37bf0473408cc5be3a122dd960862a74"}
	// #	0x530448	github.com/go-openapi/testify/v2/internal/leak.TestLeaked_MultipleLeaks.func2.2+0x48	/home/fred/src/github.com/go-openapi/testify/internal/leak/leak_test.go:169
	//  ---
	// 1 @ 0x47890e 0x410193 0x40fd12 0x530529 0x47ff01
	// # labels: {"testify-leak-check":"37bf0473408cc5be3a122dd960862a74"}
	// #	0x530528	github.com/go-openapi/testify/v2/internal/leak.TestLeaked_MultipleLeaks.func2.1+0x48	/home/fred/src/github.com/go-openapi/testify/internal/leak/leak_test.go:164

	var (
		w        strings.Builder
		routines int
	)
	scanner := bufio.NewScanner(strings.NewReader(match))

	pick := false
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), []byte(labelKey)) {
			pick = true
			continue
		}
		if !pick {
			continue
		}
		_, _ = w.Write(scanner.Bytes())
		_ = w.WriteByte('\n')
		routines++
		pick = false
	}

	if didPanic {
		fmt.Fprintf(&w, "panicked with %v\n", panicVal)
		_, _ = w.Write(panicStack)
		_ = w.WriteByte('\n')
	}

	if routines == 1 {
		return w.String()
	}

	return fmt.Sprintf("%d goroutines:\n%v", routines, w.String())
}
