package errs_test

import (
	"errors"
	"fmt"
	"testing"
)

type ErrorInternal struct {
	function string
}

func (e *ErrorInternal) Error() string {
	return fmt.Sprintf("[%s] error internal", e.function)
}

func getError(level int) error {
	level1Err := fmt.Errorf("level 1 error: %w", &ErrorInternal{function: "getData"})
	if level == 1 {
		return level1Err
	}
	if level == 2 {
		return fmt.Errorf("level 2 error: %w", level1Err)
	}

	return &ErrorInternal{function: "getData"}
}

func TestErrorAsWrap(t *testing.T) {

	err := getError(1)
	var internalErr *ErrorInternal

	if errors.As(err, &internalErr) {
		fmt.Printf("is error internal: %v\n", err)
	}
	fmt.Printf("unwrapped error: %v\n", errors.Unwrap(err))

	fmt.Printf("---\n")

	err = getError(2)
	if errors.As(err, &internalErr) {
		fmt.Printf("is error internal: %v\n", err)
	}
	unwrapped := errors.Unwrap(err)
	fmt.Printf("unwrapped error: %v\n", unwrapped)
	fmt.Printf("unwrapped unwrapped error: %v\n", errors.Unwrap(unwrapped))
}

func TestErrorsIsWrap(t *testing.T) {
	var ErrorInternal = errors.New("internal error")
	getError := func(level int) error {
		level1Err := fmt.Errorf("[getData] level 1 error: %w", ErrorInternal)
		if level == 1 {
			return level1Err
		}
		if level == 2 {
			return fmt.Errorf("[getData] level 2 error: %w", level1Err)
		}

		return ErrorInternal
	}

	err := getError(1)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
}
