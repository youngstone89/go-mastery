package make

import "fmt"

func AllocateSliceStructureWithNew() {
	var p *[]int = new([]int) // rarely useful
	s := *p
	fmt.Printf("%T %v %v\n", p, p, *p)
	if s == nil {
		fmt.Printf("%T %v %v\n", s, s, s)
	}

}

func AllocateSliceStructureWithNewComplex() {
	var p *[]int = new([]int)
	*p = make([]int, 100, 100)
	s := *p
	for i := 0; i < 100; i++ {
		s[i] = i
	}
	fmt.Printf("%T %v %v\n", p, p, *p)
	if s == nil {
		fmt.Printf("%T %v %v\n", s, s, s)
	}

}

func AllocateSliceStructureIdiomatic() {
	p := make([]int, 100)
	for i := 0; i < 100; i++ {
		p[i] = i
	}
	fmt.Printf("%T %v %v\n", p, p, p)
	if p == nil {
		fmt.Printf("%T %v %v\n", p, p, p)
	}

}

func AllocateSliceStructureWithMake() {
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
	fmt.Printf("%T %v %v\n", v, v, &v)
}
