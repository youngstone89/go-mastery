package new

import "os"

func NewFileWithLiteral() *os.File {
	return &os.File{}
}

func NewFileWithNew() *os.File {
	return new(os.File)
}
