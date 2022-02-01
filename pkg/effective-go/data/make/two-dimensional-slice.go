package make

import "fmt"

func DoTwoDimensionalSlice() {
	// Allocate the top-level slice.
	YSize := 2
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	XSize := 2
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}
	fmt.Println(picture)
}
func DoTwoDimensionalSliceOneAllocation() {
	YSize := 2
	// Allocate the top-level slice, the same as before.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	XSize := 2
	pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
	for i, _ := range pixels {
		pixels[i] = uint8(i)
	}

	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}
	fmt.Println(picture)
	fmt.Println(pixels)
}
