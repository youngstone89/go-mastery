package new

import (
	"fmt"
	"os"
)

func NewFileWithLiteral() *os.File {
	return &os.File{}
}

func NewFileWithNew() *os.File {
	return new(os.File)
}
func NewArrays() [1]string {
	return [1]string{"hi"}
}

const (
	Enone = iota
	Eio
	Einval
)

func NewArraysWithCompositLiteral() {
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Printf("%T %v %v %v", a, a[0], a[1], a[2])
}

func NewArraysWithCompositLiteralOutOfBounds() {
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	fmt.Printf("%T %v %v %v", a, a[0], a[1], a[2], a[3])
}
func NewSliceWithCompositLiteral() {
	s := []string{0: "no error", 1: "Eio", 2: "invalid argument"}
	fmt.Printf("%T %v %v %v", s, s[0], s[1], s[2])
}

func NewArraysPointer() *[1]string {
	return &[1]string{"hi"}
}

func NewSliceWithLiteral() []string {
	return []string{"a"}
}

func NewMapWithLiteral() map[int]string {
	return map[int]string{0: "no error", 1: "Eio", 2: "invalid argument"}
}
