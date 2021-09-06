package liskovsubstitution

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestReadOnlyCollection_Add(t *testing.T) {
	x := ReadOnlyCollection{}
	param := 1
	x.Add(param)
	g := x.Get(0)
	if g == nil {
		t.Errorf("g is nil:%v",g)
	}
	t.Logf("g is not nil:%v",g)
}
