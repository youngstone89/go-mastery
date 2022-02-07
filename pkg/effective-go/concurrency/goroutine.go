package concurrency

import (
	"fmt"
	"time"
)

func GoWithChannel() {
	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		doSomethingForAWhile()
		fmt.Println("finished goroutine")
		c <- 1 // Send a signal; value does not matter.
	}()
	fmt.Println("waiting for finish signal")
	<-c // Wait for sort to finish; discard sent value.
	fmt.Println("signal received")
}

func doSomethingForAWhile() {
	for i := 0; i < 10; i++ {
		fmt.Println("printing:", i)
		time.Sleep(1 * time.Second)
	}
}
