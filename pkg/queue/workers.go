package queue

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/sync/semaphore"
)

var counter int
var counterMap map[int]int = make(map[int]int)
var mu sync.Mutex
var sem *semaphore.Weighted = semaphore.NewWeighted(int64(5))

const port = "5000"

type job struct {
	name    string
	payload string
}

func runJob(id int, individualJob job) {
	mu.Lock()
	defer mu.Unlock()
	counterMap[id]++
	log.Printf("Worker %d: Completed: %s with payload %s Counter:%v", id, individualJob.name, individualJob.payload, counterMap[id])
}

func DoJob(queueSize int, workers int) {
	// ctx, cancel := context.WithCancel(context.Background())
	jobQueue := make(chan job, queueSize)
	for i := 1; i <= workers; i++ {
		success := sem.TryAcquire(1)
		if !success {
			fmt.Println("no more goroutines allowed")
			break
		}
		go func(i int) {
			for j := range jobQueue {
				runJob(i, j)
			}
		}(i)

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		submittedJob := job{r.FormValue("name"), r.FormValue("payload")}
		jobQueue <- submittedJob
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":"+port, nil)

}
