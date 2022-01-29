package iterators

import "fmt"

var sumLoops int

func simpleLoop(n int) int {
	for i := 0; i < n; i++ {
		sumLoops += i
	}
	return sumLoops
}

func DoForLoop(n int) {

	fmt.Println(simpleLoop(n))

}
