package timer

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func DoTest2(LogFlushTimeoutSec int) {

	records := &Records{
		Records: make(jsonObj),
		val:     0,
	}
	records.Records["records"] = make(jsonArr, 0)
	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan int, 1)
	count := 8000

	var obj jsonObj
	err := json.Unmarshal([]byte(event), &obj)
	_obj, err := json.Marshal(obj)
	fmt.Printf("Object size in bytes: %v \n", len(_obj))
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Go Routine processing cache
	go func(done chan int, records *Records, wg *sync.WaitGroup, counter int, obj jsonObj) {
		for {
			select {
			case <-done:
				fmt.Println("Received Done")
				records.processCache2(done, wg, counter)
				wg.Done()
				return
			}
		}
	}(done, records, &wg, count, obj)

	// Go routine adding val
	for i := 0; i < count; i++ {
		wg.Add(1)
		go records.generateLoad2(done, &wg, obj, count)
	}
	wg.Wait()
	fmt.Println(records.val)
	fmt.Println(records.TotalSize)
}

func (r *Records) generateLoad2(done chan int, wg *sync.WaitGroup, obj jsonObj, counter int) {
	defer wg.Done()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Records["records"] = append(r.Records["records"].(jsonArr), obj)
	_records, err := json.Marshal(r.Records["records"].(jsonArr))
	if err != nil {
		panic(err)
	}
	size := len([]byte(_records))
	// fmt.Printf("increased size in bytes in total: %v \n", size)
	if size > 1*1e+2 {
		go r.processCache2(done, wg, counter)
	}
	// fmt.Printf("Goroutine: %v, Increments: %v \n", goid(), len(r.Records["records"].(jsonArr)))

}

func (r *Records) processCache2(done chan int, wg *sync.WaitGroup, counter int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	count := len(r.Records["records"].(jsonArr))
	if r.Records != nil && len(r.Records) > 0 {
		if count > 0 {
			r.val += count
			fmt.Printf("Goroutine: %v , Purging : %v, Total: %v \n", goid(), len(r.Records["records"].(jsonArr)), r.val)
			_records, err := json.Marshal(r.Records["records"].(jsonArr))
			if err != nil {
				panic(err)
			}
			byteSize := len([]byte(_records))
			fmt.Printf("Purged size in bytes in total: %v \n", byteSize)
			r.TotalSize += byteSize
			if r.val == counter {
				done <- 1
			}
			r.Records["records"] = make(jsonArr, 0)
		}
	} else {
		fmt.Println("records empty")
	}

}
