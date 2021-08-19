package gross

import b "go-mastery/pkg/basic"

func GrossSalary() int  {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}