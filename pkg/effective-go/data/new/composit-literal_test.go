package new_test

import (
	"fmt"
	n "go-mastery/pkg/effective-go/data/new"
	"testing"
)

func TestCompositLiteral(t *testing.T) {
	lfile := n.NewFileWithLiteral()
	fmt.Printf("%T %v \n", lfile, lfile)
	nFile := n.NewFileWithNew()
	fmt.Printf("%T %v \n", nFile, lfile)
}
