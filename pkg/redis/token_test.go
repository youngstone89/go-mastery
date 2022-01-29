package redis

import (
	"fmt"
	"testing"
)

func Test_decodeToken(t *testing.T) {
	token := &Token{}
	err := decodeToken(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token.AccessToken)
	fmt.Println(token.ExpriresIn)
	fmt.Println(token.Scope)
	fmt.Println(token.TokenType)
}
