package pipeline

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Run() {

	go func() {
		for {
			fmt.Printf("number of go routines:%v \n", runtime.NumGoroutine())
			time.Sleep(time.Duration(1) * time.Second)
		}

	}()

	start := time.Now()

	done := make(chan int, 1)

	readCount := 100000
	c := NewConsumer(readCount)

	storageRateLimit := 10
	timeout := 10
	s := NewStorage(storageRateLimit, timeout, readCount, done)

	var storages []*Storage
	storages = append(storages, s)

	processorRateLimit := 10000
	pcs := NewProcessor(storages, processorRateLimit)
	var processors []*Processor
	processors = append(processors, pcs)

	collector := NewCollector(c, processors)

	var wg sync.WaitGroup

	queueDepth := 100000

	wg.Add(1)
	go collector.Start(&wg, queueDepth, done)

	wg.Wait()

	fmt.Printf("Read Count: %v, Process Goroutines Limit: %v, Write Rate Limit: %v, Purge Timeout:%v, Execution Time : %v", readCount, processorRateLimit, storageRateLimit, timeout, time.Since(start))

}
