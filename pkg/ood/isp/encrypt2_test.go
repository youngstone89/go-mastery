package isp

import (
	"golang.org/x/net/context"
	"testing"
)

func TestEncryptV2(t *testing.T) {
	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// store the key
	ctx = context.WithValue(ctx, "encryption-key", "-secret-")

	// call the function
	data, _ := EncryptV2(ctx, ctx, []byte("my data"))
	t.Logf("data :%v",string(data))

}
