package recover_test

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	IamGoingToPanic()
}

func IamGoingToPanic() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("panic was %q", r)
		}
	}()
	panic("OMG")
}
