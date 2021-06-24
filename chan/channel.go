package main

import "fmt"

var i int

func doSomething() {
	i++
	fmt.Printf("I am doing something... %d",i)
}

func main() {
	ch := make(chan int)
	go sendingGoRoutine(ch)
	go receivingGoRoutine(ch)

}

func receivingGoRoutine(ch chan int) {
	v :=<- ch
	fmt.Println("Received value",v)
}

func sendingGoRoutine(ch chan int) {
	ch <- 45
}



