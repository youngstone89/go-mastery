package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func DoWithBufferedChannel(bufsize int) {

	var sem = make(chan int, bufsize)
	var q = make(chan int)
	loadSize := 1000

	_process := func(r int) {
		fmt.Println("handling request:", r)
		time.Sleep(10 * time.Millisecond)
	}

	_handle := func(wg *sync.WaitGroup, r int) {
		defer wg.Done()
		sem <- 1    // Wait for active queue to drain.
		_process(r) // May take a long time.
		<-sem       // Done; enable next request to run.
		fmt.Println("completed:", r, "goroutine:", runtime.NumGoroutine())

	}

	_server := func(queue chan int, loadSize int) {
		var wg sync.WaitGroup
		wg.Add(loadSize)
		for i := 0; i < loadSize; i++ {
			req := <-queue
			go _handle(&wg, req) // Don't wait for handle to finish.
		}
		wg.Wait()

	}
	_genLoad := func(q chan<- int, size int) {
		for i := 0; i < size; i++ {
			q <- i
			fmt.Println("[genLoad] sent", i)
		}
	}
	go _genLoad(q, loadSize)
	_server(q, loadSize)

}
