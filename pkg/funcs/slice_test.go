package funcs_test

import (
	"fmt"
	"testing"
)

func TestDoAppendToIntSlice(t *testing.T) {
	x := []int{1, 2, 3}

	doAppend(x)
	fmt.Println("outside: ", x) // outside:  [1 2 3]
	x = doAppendAndReturn(x)

	fmt.Println("outside: ", x) // outside:  [1 2 3]
}
func doAppend(sl []int) {
	sl = append(sl, 100)
	fmt.Println("inside: ", sl) // inside:  [1 2 3 100]
}

func doAppendAndReturn(sl []int) []int {
	sl = append(sl, 100)
	fmt.Println("inside: ", sl) // inside:  [1 2 3 100]
	return sl
}

func TestDoAppendToMapSlice(t *testing.T) {
	x := []map[string]any{{"original": nil}}
	modifyMap(x)
	fmt.Println("outside: ", x)

}

func modifyMap(m []map[string]any) {
	newObj := map[string]any{"hi": "there"}
	m[0] = newObj
	fmt.Println("inside: ", m)

}
