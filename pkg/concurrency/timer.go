package concurrency

import (
	"log"
	"time"
)

func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func DoTimer() {
	for i := 0; i < 2; i++ {
		c := timer(1 * time.Second)
		v := <-c
		log.Printf("read: %d", v)
	}
}
