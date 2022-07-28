package concurrency

import (
	"fmt"
	"time"
)

func Produce(ch chan<- int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++

		time.Sleep(d)
	}
}

func Consume(ch <-chan int, id int) {
	for i := range ch {
		fmt.Printf("worker %v got %v \n", id, i)
	}
}
