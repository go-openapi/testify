//go:build integrationtest

package examplespkg

import "fmt"

type helper struct {
	value string
}

func ExampleFormatter() {
	f := Formatter{Prefix: ">>"}
	h := helper{value: "test"}
	fmt.Printf("%s %s", f.Prefix, h.value)
	// Output: >> test
}
