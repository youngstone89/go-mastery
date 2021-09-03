package liskovsubstitution

import "testing"

func TestCar3(t *testing.T) {

	car3 := &Car3{}
	Go3(car3)

	buggy3 := &Buggy3{}
	Go3(buggy3)
}
