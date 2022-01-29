package algorithms

func FindMinInt(a []int) int {
	var minInt int = a[0]
	for _, i := range a {
		if minInt > i {
			minInt = i
		}
	}
	return minInt

}
