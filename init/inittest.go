package main

import "fmt"

var v = varaibleInit()


func varaibleInit() int {
	fmt.Println("variable")
	return 0
}

func main() {
fmt.Println("main")

}

func init() {
	fmt.Println("init")
}
