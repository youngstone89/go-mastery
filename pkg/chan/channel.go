package main

import "fmt"

var i int

func doSomething() {
	i++
	fmt.Printf("I am doing something... %d", i)
}

func DoChannelTest() {
	ch := make(chan int)
	go sendingGoRoutine(ch)
	receivingGoRoutine(ch)

	//x := <- ch
	//fmt.Println(x)

}

func receivingGoRoutine(ch chan int) {
	v := <-ch
	fmt.Println("Received value", v)
}

func sendingGoRoutine(ch chan int) {
	ch <- 45
}
