// Package testsrc is a package for testing spew with cgo.
//
// NOTE: Due to the following build constraints, this file will only be compiled
// when both cgo is supported and "-tags testcgo" is added to the go test
// command line.  This code should really only be in the dumpcgo_test.go file,
// but unfortunately Go will not allow cgo in test files, so this is a
// workaround to allow cgo types to be tested.  This configuration is used
// because spew itself does not require cgo to run even though it does handle
// certain cgo types specially.  Rather than forcing all clients to require cgo
// and an external C compiler just to run the tests, this scheme makes them
// optional.
package testsrc
