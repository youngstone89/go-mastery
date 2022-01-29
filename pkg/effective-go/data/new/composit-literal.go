package new

import (
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

func NewArraysPointer() *[1]string {
	return &[1]string{"hi"}
}

// func NewFileWithNew() *os.File {
// 	return new(os.File)
// }
// func NewFileWithNew() *os.File {
// 	return new(os.File)
// }
