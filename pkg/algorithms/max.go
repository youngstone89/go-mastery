package algorithms

func FindMaxInt(b []int) int {
	var max int = b[0]
	for _, i := range b {
		if max < i {
			max = i
		}
	}
	return max
}
