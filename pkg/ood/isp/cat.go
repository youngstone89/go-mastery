package isp

import "fmt"

type Cat struct{}

// Implicitly implement the fmt.Stringer interface
func (c Cat) String() string {
	return "Meow!"
}

func Meow()  {

	kitty := Cat{}
	fmt.Printf("Kitty %s", kitty)
	
}