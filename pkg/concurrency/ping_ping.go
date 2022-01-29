package concurrency

import (
	"log"
	"time"
)

func player(table chan int) {
	for {
		ball := <-table
		log.Printf("hit the ball:%d", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
func DoPingPong() {
	var Ball int
	table := make(chan int)
	for i := 0; i < 100; i++ {
		go player(table)
	}
	// go player(table)
	// go player(table)
	// go player(table)
	// go player(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}
