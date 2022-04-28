package useinterfaceeffectively_test

import "testing"

func TestTypeAssertion(t *testing.T) {
	var x interface{}
	x = "a"
	v, sucs := x.(string)
	if sucs {
		t.Log(v, &v)
	}
	y := x
	v, sucs = y.(string)
	v = "b"
	if sucs {
		t.Log(v, &v)
	}

	t.Log(&x, &y)
	t.Log(x, y)

}
func TestTypeAssertionThatFails(t *testing.T) {
	var x interface{}
	x = "a"
	v, sucs := x.(int)
	if sucs {
		t.Log(v)
	} else {
		t.Fail()
	}

}
