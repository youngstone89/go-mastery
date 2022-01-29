package algorithms

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	intData := []int{3, 1, 2, 5, 6, 4}
	minResult := FindMinInt(intData)
	fmt.Println("Minimum value in array: ", minResult)
}
