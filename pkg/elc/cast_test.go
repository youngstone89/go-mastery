package elc_test

import (
	"testing"
)

type (
	Code      float64
	RetryFunc func() error
)

type ErrBody struct {
	ErrorCode float64 `json:"error_code"`
	Message   string  `json:"message"`
}

func TestCastNil(t *testing.T) {
	var errBody ErrBody

	code := Code(errBody.ErrorCode)
	checkNil(code)

}

func checkNil(code Code) {
	if code != 40403 {

	}

}
