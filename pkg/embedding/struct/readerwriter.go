package structembedding

import "bufio"

type ReadWriter struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func (rw *ReadWriter) Read(p []byte) (n int, err error) {
	return rw.reader.Read(p)
}

func (rw *ReadWriter) Write(p []byte) (n int, err error) {
	return rw.writer.Write(p)
}
