package useinterfaceeffectively_test

import (
	"fmt"
	"testing"
)

// Compile-time checks for ensuring a type implements a particular
// interface.
var (
	// Works but allocates a dummy Foo instance on the heap.
	_ fmt.Stringer = &Foo{}

	// Preferred way that does not allocate anything on the heap.
	_ fmt.Stringer = (*Foo)(nil)
)

type Foo struct{}

func (*Foo) String() string { return "Foo" }

func TestFooCompileTimeCheck(t *testing.T) {

}
