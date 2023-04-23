// Go program to delete the element from the given slice
package main

import "fmt"

func main() {

	// Declaring a slice
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original Slice:", numbers)

	var index int = 3

	// Get the element at the provided index in the slice
	elem := numbers[index]

	// Using append function to combine two slices
	// first slice is the slice of all the elements before the given index
	// second slice is the slice of all the elements after the given index
	// append function appends the second slice to the end of the first slice
	// returning a slice, so we store it in the form of a slice
	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Printf("The element %d was deleted.\n", elem)
	fmt.Println("Slice after deleting elements:", numbers)
}
