package make

import "fmt"

func DoSlicing() {
	s := []int{2, 4, 5, 7, 11, 13}
	s = s[:0]
	printSlice(s)
	s = s[:6]
	printSlice(s)
	s = s[:5]
	printSlice(s)
	s = s[:6]
	printSlice(s)
	s = s[5:6]
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:1]
	printSlice(s)

}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
