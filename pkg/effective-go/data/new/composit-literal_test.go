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

	slice := n.NewSliceWithLiteral()
	slice[0] = "b"
	slice2 := slice
	fmt.Printf("%T %v  %v \n", slice, slice, slice[0])
	fmt.Printf("%T %v  %v \n", slice2, slice2, slice2[0])

	m := n.NewMapWithLiteral()
	fmt.Printf("%T %v  %v \n", m, m, m[0])

}

func TestArraysWithCompositLiteral(t *testing.T) {
	n.NewArraysWithCompositLiteral()
}

func TestArraysWithCompositLiteralOutOfBounds(t *testing.T) {
	n.NewArraysWithCompositLiteralOutOfBounds()
}

func TestSliceWithCompositLiteral(t *testing.T) {
	n.NewSliceWithCompositLiteral()
}

func TestNewMapWithCompositLiteral(t *testing.T) {
	m := n.NewMapWithCompositLiteral()
	v := m["UCT"]
	//zero value for the type of entries in the map is returend
	// when the entry type is integer, the returning zero value is 0
	fmt.Printf("value: %v \n", v)
}

func TestCommaOkayIdom(t *testing.T) {
	rs := n.CommaOkIdom("X")
	fmt.Printf("value: %v \n", rs)
	rs = n.CommaOkIdom("PST")
	fmt.Printf("value: %v \n", rs)
}

func TestCommaOkayIdomClosure(t *testing.T) {
	tz := n.CommaOkIdomClosure()
	tz()["UTC"] = 1
	tz()
}
