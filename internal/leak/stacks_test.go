package leak

import (
	"sync"
	"testing"

	"github.com/go-openapi/testify/v2/internal/spew"
)

func TestStackCurrent(_ *testing.T) {
	var wg sync.WaitGroup

	var s Stack
	wg.Add(1)
	go func() {
		defer wg.Done()
		s = Current()
	}()

	wg.Wait()

	// spew.Config = spew.ConfigState{DisableMethods: true}
	spew.Dump(s)
}
