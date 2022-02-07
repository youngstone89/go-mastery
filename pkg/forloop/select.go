package forloop

import (
	"fmt"
	"time"
)

func DoSelectInLoop() {
	shutdown := make(chan struct{})
	go func() {
		for {
			select {
			case <-shutdown:
				return
			default:

				fmt.Println("before continue")
				continue
				fmt.Println("after continue")

			}
		}
	}()

	time.Sleep(5 * time.Second)
	shutdown <- struct{}{}
}
