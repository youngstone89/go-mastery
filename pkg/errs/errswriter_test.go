package errs

import (
	"fmt"
	"testing"
)

func TestErrWriter(t *testing.T) {

	fw := FileWriter{}
	ew := &errWriter{w: fw}
	strVal := "a"
	ew.write([]byte(strVal))
	strVal = "b"
	ew.write([]byte(strVal))

	if ew.err != nil {
		fmt.Println(ew.err)
	}

}

func TestFileWriterWithErr(t *testing.T) {

	fw := FileWriterWithErr{}
	ew := &errWriter{w: fw}
	strVal := "a"
	ew.write([]byte(strVal))
	strVal = "b"
	ew.write([]byte(strVal))

	if ew.err != nil {
		fmt.Println(ew.err)
	}

}

type FileWriterWithErr struct {
}

func (fw FileWriterWithErr) Write(p []byte) (n int, err error) {
	err = fmt.Errorf("error in writing file")
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
