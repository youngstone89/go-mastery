package gross

import b "go-mastery/basic"

func GrossSalary() int  {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}