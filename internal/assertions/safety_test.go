// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"sync"
	"testing"
)

func TestNoGoRoutineLeak_Success(t *testing.T) {
	mockT := new(mockT)

	result := NoGoRoutineLeak(mockT, func() {
		// Clean function — no goroutines spawned.
	})

	if !result {
		t.Error("expected NoGoRoutineLeak to return true for clean function")
	}
	if mockT.failed {
		t.Error("expected no failure for clean function")
	}
}

func TestNoGoRoutineLeak_Failure(t *testing.T) {
	blocker := make(chan struct{})
	var wg sync.WaitGroup

	t.Cleanup(func() {
		close(blocker)
		wg.Wait()
	})

	mockT := new(mockT)

	wg.Add(1)
	result := NoGoRoutineLeak(mockT, func() {
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks until cleanup
		}()
	})

	if result {
		t.Error("expected NoGoRoutineLeak to return false for leaking function")
	}
	if !mockT.failed {
		t.Error("expected failure to be reported for leaking function")
	}
}
