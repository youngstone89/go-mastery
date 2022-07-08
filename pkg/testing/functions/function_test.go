package functions

import (
	"errors"
	"fmt"
	"testing"
)

type jsonObj = map[string]interface{}

type Funtional interface {
	Fn | Fn2
}
type Fn func()
type Fn2 func() error

type MyConstraint interface {
	int | float64
}

func MyFunc[T MyConstraint](t T) {
	fmt.Println(t)
}

func TestWithMyConstraint(t *testing.T) {
	MyFunc[int](1)
	// MyFunc[string]("hello")
}

func TestTakeFunctionAsInterface(t *testing.T) {

	locFn := func() {
		fmt.Println("i am local")
	}
	TakeFnAsInf(Fn(locFn))

	locFn2 := func() error {
		fmt.Println("i am local2")
		return errors.New("damn")
	}
	TakeFnAsInf(Fn2(locFn2))

}

func TakeFnAsInf[F Funtional](fn F) {
	// fn(Fn)()
}

func f[T any](t T) {
	fmt.Println(t)
}

type s[T any] struct {
	t T
}

func TestAnyTypeConstractFunction(t *testing.T) {
	f[int](1)
	anyInstance := s[int]{t: 1}
	fmt.Println(anyInstance)
}
func TestJsonObject(t *testing.T) {
	obj := make(jsonObj)
	obj["key"] = "value"
	lclFn := func() error {
		obj["key"] = "changed"
		return nil
	}
	Do(lclFn)
	fmt.Println(obj)
}

type retryFunc func() error

func Do(fn retryFunc) error {
	err := fn()
	if err != nil {
		return err
	}
	return nil
}

type retriable func() error

type retrier interface {
	Retry(retriable)
}

func Test(t *testing.T) {

	a := func() error {
		fmt.Println("a")
		return nil
	}
	retriable(a)()

}
