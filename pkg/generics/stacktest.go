package generics

import "fmt"

type Stack struct {
	vals []interface{}
}
func (s *Stack) Push(val interface{}) {
	s.vals = append(s.vals,val)
}
func (s *Stack) Pop() (interface{},bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals) - 1]
	s.vals = s.vals[:len(s.vals) - 1]
	return top, true
}

func DoStackTest() {
	var s Stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	v, ok := s.Pop()
	fmt.Println(v,ok)
}