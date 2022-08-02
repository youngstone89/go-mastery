package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestReturnGoroutine(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		fmt.Println("outer go routing...")
		go func() {
			fmt.Println("inner go routing...")
			go func() {
				fmt.Println("inner most go routing...")
				ch <- struct{}{}
			}()
		}()
	}()
	start := time.Now()
	<-ch
	fmt.Println(time.Since(start))
}

func TestGoInLoop(t *testing.T) {

	inputChan := make(chan int)
	go func() {
		inputChan <- 1
	}()
	go func() {
		inputChan <- 2
	}()

	for {
		select {
		case i := <-inputChan:
			go func() {
				defer func() {
					fmt.Println("out of go func")
				}()
				if i%2 == 0 {
					fmt.Println("returning")
					return
				}
				fmt.Println(i)
			}()
		default:
			time.Sleep(1 * time.Second)
		}
	}

}
