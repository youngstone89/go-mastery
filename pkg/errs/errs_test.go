package errs

import (
	"errors"
	"fmt"
	"testing"
)

//Compare wrapper error instance to wrapped error type
func TestW(t *testing.T) {
	err := ErrNotFound
	werr := fmt.Errorf("access denied: %w", err)
	if errors.Is(werr, ErrNotFound) {
		fmt.Println("matched")
	}

}
func TestErrsNotFound(t *testing.T) {

	err := errors.New("not found")
	if err == ErrNotFound {
		fmt.Println("equal")
	} else {
		fmt.Println("Not Equal")
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("equal")
	} else {
		fmt.Println("Not Equal")
	}

}

func TestErrsNotFoundCompare(t *testing.T) {
	err := ErrNotFound
	if err == ErrNotFound {
		fmt.Println("equal")
	} else {
		fmt.Println("Not Equal")
	}

}

func TestCompareNotFoundErrorInstance(t *testing.T) {

	err := NewNotFoundError()

	if _, ok := err.(*NotFoundError); ok {
		fmt.Println("identical type")
	} else {
		fmt.Println("not identical type")
	}
	nfe := &NotFoundError{}
	fmt.Println(nfe)
	if errors.As(err, &nfe) {
		fmt.Println("identical type")
	} else {
		fmt.Println("not identical type")
	}

}

func TestQueryError(t *testing.T) {

	err := NewQueryError()
	if e, ok := err.(*QueryError); ok && e.Err == ErrNotFound {
		fmt.Println("wrapped error is ErrNotFound")
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("wrapped error is ErrNotFound")
	}

}
