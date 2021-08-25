package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x,y = y, x+y
			fmt.Printf("Channel C Value: %v , X Value: %v, Y Value: %v \n",c,x,y)
		//case v := <- c:
		//	//x, y = y, x+y
		//	fmt.Printf("Channel C Value: %v \n",v)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go feedChannelWithI(c, quit)
	fibonacci(c, quit)
}

func feedChannelWithI(c chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Looping Index: %d \n", i)
		c <- i
	}
	fmt.Printf("For loop ends \n")
	quit <- 0
}

