package concurrency

import (
	"log"
	"net"
	"time"
)

func DoServeAndLogWithWorkers() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan string)
	go workerpool(ch, 4)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}
func workerpool(ch chan string, n int) {
	wch := make(chan int)
	results := make(chan int)
	for i := 0; i < n; i++ {
		go logWorker(wch, results)
	}
	go parse(results)
	for {
		addr := <-ch
		l := len(addr)
		wch <- l
	}
}

func parse(results chan int) {
	for {
		log.Printf("parsing results:%d", <-results)
	}
}
func logWorker(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}
