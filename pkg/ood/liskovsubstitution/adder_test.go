package liskovsubstitution

import "testing"

func TestAdder(t *testing.T) {

	age := Age{}

	s := PrintSum(1,1, &age)

	if s == "" {
		t.Errorf("error")
	}

	//double := Double{}
	//
	//s := PrintSum(1,1, &double)
	//
	//if s == "" {
	//	t.Errorf("error")
	//}


}
