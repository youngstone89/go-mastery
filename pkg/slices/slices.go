package slices

import "fmt"

func DoWithPointerIntegerSlice(elements []int) {
	if len(elements) <= 0 {
		fmt.Printf("no elements: %d ", len(elements))
		return
	}
	var pointerArray []int
	var pointerStore []*int
	for _, v := range elements {
		fmt.Println(v)
		c := v
		p := &c
		pointerArray = append(pointerArray, c)

		pointerStore = append(pointerStore, p)
	}
	for i := 0; i < len(pointerStore); i++ {
		*pointerStore[len(pointerStore)-1-i] = i + 1
	}
	for _, v := range pointerArray {
		fmt.Printf("%T,%v, %v", v, v, v)
		fmt.Println()
	}

}

func DoWithPointerSlice(element ...int) {
	var slice []int = nil
	slice = append(slice, element...)

	for _, v := range slice {
		fmt.Println(v)
	}

	var pointerSlice *[]int = &[]int{}
	*pointerSlice = append(*pointerSlice, element...)

	for _, v := range *pointerSlice {
		fmt.Println(v)
	}

}
