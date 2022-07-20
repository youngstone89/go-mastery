package funcs_test

import (
	"fmt"
	"os"
	"testing"
)

func TestVariadicFunction(t *testing.T) {
	args := []int{1, 2, 3, 4, 5}
	fmt.Println(args)
	fmt.Println(sum(args))
	fmt.Println(args)
	fmt.Fprintln(os.Stdout, args)
	fmt.Println(variadicSum(args...))

	args2 := []string{"1", "2"}
	fmt.Println(args2)

	fmt.Println(variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func variadicSum(numbers ...int) (sum int) {

	for _, n := range numbers {
		sum += n
	}
	return sum
}
func sum(numbers []int) int {

	sum := 0

	for _, n := range numbers {

		sum += n

	}

	return sum

}
