package ioutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestOsFileWrite(t *testing.T) {
	os.MkdirAll(fmt.Sprintf("%s/%s", "test", "test"), 0775)
	ioutil.WriteFile(fmt.Sprintf("%s/%s/%s", "test", "test", "test"), []byte{}, 0775)
}
