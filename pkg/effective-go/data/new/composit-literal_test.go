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

	arr := n.NewArrays()
	// value copy doesn't change value of arrays
	arr2 := arr
	arr2[0] = "there"
	fmt.Printf("%T %v  %v \n", arr, arr, arr[0])

	arrPointer := n.NewArraysPointer()
	// pointer copy changes value of original arrays
	arrPointer2 := arrPointer
	arrPointer[0] = "there"
	fmt.Printf("%T %v  %v \n", arrPointer2, arrPointer2, arrPointer2[0])
}
