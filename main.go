package main

import (
	"fmt"
	b "go-mastery/basic"
	g "go-mastery/basic/gross"

)

func main() {

	b.Basic = 10000
	fmt.Println(b.Basic)

	fmt.Println(g.GrossSalary())
}
