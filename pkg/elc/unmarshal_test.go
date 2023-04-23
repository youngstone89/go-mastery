package elc_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type ErrorBody struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func TestUnmarshal(t *testing.T) {

	var errBody ErrBody

	err := &mixHTTPError{
		ErrorCode: 90404,
		Message:   "error in response [code:401][body:Unauthorized][error:<nil>]",
	}
	merr := json.Unmarshal([]byte(err.Error()), &errBody)
	if merr != nil {
		t.Errorf(merr.Error())
	}
	t.Log(errBody)
	t.Log(errBody.ErrorCode)

}

type mixHTTPError struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (m mixHTTPError) Error() string {
	return fmt.Sprintf("{\"error_code\":%d,\"message\":\"%s\"}", m.ErrorCode, strings.ReplaceAll(m.Message, "\"", ""))
}

func TestReplaceUnmarshall(t *testing.T) {

	mixHTTPErr := &mixHTTPError{
		ErrorCode: 90401,
		Message:   `Get "https://log.api.nuance.com/consumers/records": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`,
	}
	var errBody mixHTTPError
	err := json.Unmarshal([]byte(mixHTTPErr.Error()), &errBody)
	if err != nil {
		t.Errorf(err.Error())
	}

}
