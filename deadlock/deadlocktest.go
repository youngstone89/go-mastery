package deadlock

import (
	"fmt"
)

type DeadLock struct {
	
}

func NewDeadLock() *DeadLock {
	return &DeadLock{}
}

func (d *DeadLock) Gotestdeadlock() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		//chan2 is read
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	//v := 2
	//ch2 <- v
	////ch1 read
	//v2 := <-ch1
	//fmt.Println(v, v2)

	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
	
}
