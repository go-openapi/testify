//go:build integrationtest

package examplespkg

import "fmt"

func ExampleGreet() {
	fmt.Println(Greet("World"))
	// Output: Hello, World!
}

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	// Output: 3
}

func ExampleAdd_negative() {
	fmt.Println(Add(-1, -2))
	// Output: -3
}
