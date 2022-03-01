package structembedding_test

import (
	structembedding "go-mastery/pkg/embedding/struct"
	"testing"
)

func TestReaderWriterRead(t *testing.T) {

	rWriter := structembedding.ReadWriter{}

	data := `hi`
	n, err := rWriter.Read([]byte(data))
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}
