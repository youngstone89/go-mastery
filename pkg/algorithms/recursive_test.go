package algorithms_test

import "testing"

func TestRecursive(t *testing.T) {
	day := 1
	choco := []int{1, 1}
	output := Joy(choco, day)
	t.Log(output)

}

func Joy(choco []int, day int) int {
	n := len(choco)
	if n == 1 {
		return day * choco[0]
	}
	left := day*choco[0] + Joy(choco[1:], day+1)
	right := day*choco[n-1] + Joy(choco[:n-1], day+1)
	return Max(left, right)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
