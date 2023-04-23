package elc_test

import (
	"errors"
	"fmt"
	"testing"
)

type CustomError struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (m *CustomError) Error() string {
	return fmt.Sprintf("{\"error_code\":%d,\"message\":\"%s\"}", m.ErrorCode, m.Message)
}

func TestErrorString(t *testing.T) {

	code := 401
	httpError := errors.New("")
	httpError = nil
	err := &CustomError{ErrorCode: 90401, Message: fmt.Sprintf("[code:%d][body:%s][error:%v]", code, "Unauthorized", httpError)}
	t.Log(err.Error())

}
