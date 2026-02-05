// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package leak

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
	"time"
)

// TestLeaked_NoLeak verifies that after a clean function completes,
// no goroutines with our label remain in the goroutine profile.
func TestLeaked_NoLeak(t *testing.T) {
	tested := func() {
		// Does nothing — should leave no goroutines behind.
	}

	leaked := Leaked(t.Context(), tested)
	if leaked != "" {
		t.Errorf("expected no leaked goroutines, but profile contains our label:\n%s", leaked)
	}
}

// TestLeaked_WithLeak verifies that a function that leaks a goroutine
// shows up in the goroutine profile with our label.
func TestLeaked_WithLeak(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup

	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	wg.Add(1)
	tested := func() {
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks forever until cleanup
		}()
	}

	leaked := Leaked(t.Context(), tested)
	if leaked == "" {
		t.Error("expected leaked goroutine with our label, but found none")

		return
	}

	t.Logf("detected leaked goroutine profile:\n%s", leaked)
}

// TestLeaked_TransitiveLeak verifies that grandchild goroutines
// inherit the label from the parent.
func TestLeaked_TransitiveLeak(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup
	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	wg.Add(2)
	tested := func() {
		defer wg.Done()
		go func() {
			defer wg.Done()
			// Child spawns a grandchild — label should propagate.
			go func() {
				<-blocker
			}()
		}()
	}

	// Give the child goroutine time to spawn the grandchild.
	time.Sleep(50 * time.Millisecond)

	leaked := Leaked(t.Context(), tested)
	if leaked == "" {
		t.Error("expected grandchild goroutine to inherit label, but found none")
	} else {
		t.Logf("detected transitive leak:\n%s", leaked)
	}
}

func TestLeaked_Panic(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup
	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	wg.Add(1)
	tested := func() {
		defer wg.Done()
		go func() {
			<-blocker
		}()
		panic("yay")
	}

	time.Sleep(50 * time.Millisecond)

	leaked := Leaked(t.Context(), tested)
	if leaked == "" {
		t.Error("expected panicked leaked goroutine to be detected")
	} else {
		t.Logf("detected leak while panicking:\n%s", leaked)
	}
}

func TestLeaked_Goexit(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup
	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	wg.Add(1)
	tested := func() {
		defer wg.Done()
		go func() {
			<-blocker
		}()
		runtime.Goexit()
	}

	time.Sleep(50 * time.Millisecond)

	leaked := Leaked(t.Context(), tested)
	if leaked == "" {
		t.Error("expected exited leaked goroutine to be detected")
	} else {
		t.Logf("detected leak while Goexit:\n%s", leaked)
	}
}

// TestLeaked_PreExistingGoroutineNotLabeled verifies that a goroutine
// started BEFORE pprof.Do is not attributed to our label.
func TestLeaked_PreExistingGoroutineNotLabeled(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup
	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	// Start a goroutine before the instrumented call.
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-blocker
	}()

	tested := func() {
		// Does nothing — the pre-existing goroutine is not ours.
	}

	leaked := Leaked(t.Context(), tested)
	if leaked != "" {
		t.Errorf("pre-existing goroutine should NOT carry our label, but profile contains:\n%s", leaked)
	}
}

// TestLeaked_ParallelIsolation verifies that two concurrent tests
// with different labels don't interfere.
func TestLeaked_ParallelIsolation(t *testing.T) {
	const workers = 5
	var wg sync.WaitGroup
	results := make([]string, workers)

	for i := range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()

			blocker := make(chan struct{})
			tested := func() {
				go func() {
					<-blocker
				}()
			}

			results[i] = Leaked(t.Context(), tested)
			close(blocker) // clean up
		}()
	}

	wg.Wait()

	for i, r := range results {
		if r == "" {
			t.Errorf("worker %d: expected to detect its own leaked goroutine", i)
		}
	}
}

func TestLeaked_MultipleLeaks(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup

	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	wg.Add(2)
	tested := func() {
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks forever until cleanup
		}()

		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks forever until cleanup
		}()
	}

	leaked := Leaked(t.Context(), tested)
	if leaked == "" {
		t.Error("expected leaked goroutine with our label, but found none")

		return
	}

	t.Logf("detected multiple leaked goroutine profiles:\n%s", leaked)
}

// TestLeaked_DumpFullProfile dumps the raw debug=1 profile
// for manual inspection of the label format.
func TestLeaked_DumpFullProfile(t *testing.T) {
	blocker := make(chan struct{})
	t.Cleanup(func() { close(blocker) })

	id := uniqueLabel()
	labels := pprof.Labels(labelKey, id)

	pprof.Do(context.Background(), labels, func(_ context.Context) {
		go func() {
			<-blocker
		}()
	})

	// Small delay to let the goroutine settle into its blocked state.
	time.Sleep(10 * time.Millisecond)

	var buf bytes.Buffer
	profile := pprof.Lookup("goroutine")
	if err := profile.WriteTo(&buf, 1); err != nil {
		t.Fatalf("failed to write goroutine profile: %v", err)
	}

	t.Logf("=== full goroutine profile (debug=1) ===\n%s", buf.String())
	t.Logf("=== searching for label key=%q id=%q ===", labelKey, id)

	needle := fmt.Appendf(nil, "%q:%q", labelKey, id)
	if bytes.Contains(buf.Bytes(), needle) {
		t.Logf("FOUND label in profile with needle: %s", needle)
	} else {
		t.Errorf("label NOT found in profile. Tried: %q", needle)
	}
}
