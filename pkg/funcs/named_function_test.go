package funcs_test

import (
	"fmt"
	"testing"
)

func TestNamedFunction(t *testing.T) {
	dvide := func(num, div int) (res, rem int) {

		res = num / div

		rem = num % div

		return res, rem

	}
	res, rem := dvide(10, 2)
	fmt.Printf("res: %v, rem: %v \n", res, rem)
}
