package algorithms

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	intData := []int{3, 1, 2, 5, 6, 4}
	maxResult := FindMaxInt(intData)
	fmt.Println("Maximum value in array: ", maxResult)
}
