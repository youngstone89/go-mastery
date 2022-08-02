package errs

import (
	"io"
	"io/ioutil"
)

type errWriter struct {
	w   io.Writer
	err error
}

func (e *errWriter) write(p []byte) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.Write(p)
}

type FileWriter struct {
}

func (fw FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile("filename", p, 0755)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
