package new_test

import (
	"fmt"
	n "go-mastery/pkg/effective-go/data/new"
	"testing"
)

func TestZeroValue(t *testing.T) {
	p := new(n.SyncedBuffer)
	var v n.SyncedBuffer

	fmt.Printf("%T %v %v\n", p, p.Buffer(), p.Lock())
	fmt.Printf("%T %v %v\n", v, v.Buffer(), v.Lock())

}
