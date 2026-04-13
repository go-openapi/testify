// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"net"
	"os"
	"runtime"
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

func TestNoFileDescriptorLeak_Success(t *testing.T) {
	mockT := new(mockT)

	result := NoFileDescriptorLeak(mockT, func() {
		// Clean function — no file descriptors opened.
	})

	if !result {
		t.Error("expected NoFileDescriptorLeak to return true for clean function")
	}
	if mockT.failed {
		t.Error("expected no failure for clean function")
	}
}

func TestNoFileDescriptorLeak_Failure(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("file descriptor leak detection requires Linux")
	}

	mockT := new(mockT)

	var leakedFile *os.File

	result := NoFileDescriptorLeak(mockT, func() {
		f, err := os.CreateTemp(t.TempDir(), "fdleak-test-*")
		if err != nil {
			t.Fatalf("CreateTemp: %v", err)
		}

		leakedFile = f // intentionally not closed
	})

	t.Cleanup(func() {
		if leakedFile != nil {
			leakedFile.Close()
			os.Remove(leakedFile.Name())
		}
	})

	if result {
		t.Error("expected NoFileDescriptorLeak to return false for leaking function")
	}
	if !mockT.failed {
		t.Error("expected failure to be reported for leaking function")
	}
}

func TestNoFileDescriptorLeak_SocketFiltered(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("file descriptor leak detection requires Linux")
	}

	mockT := new(mockT)

	var leakedListener net.Listener

	result := NoFileDescriptorLeak(mockT, func() {
		var lc net.ListenConfig
		ln, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatalf("net.Listen: %v", err)
		}

		leakedListener = ln // intentionally not closed — socket FD should be filtered
	})

	t.Cleanup(func() {
		if leakedListener != nil {
			leakedListener.Close()
		}
	})

	if !result {
		t.Error("expected socket FD to be filtered, but assertion reported a leak")
	}
	if mockT.failed {
		t.Error("expected no failure when socket FD is filtered")
	}
}
