package infembedding

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// only interfaces can be embedded within interface
type ReadWriter interface {
	Reader
	Writer
}
