package pipeline

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Records struct {
	Records []int
	lock    sync.Mutex
}

type Storage struct {
	rateLimit *rate.Limiter
	timeout   int
	records   *Records
	readCount int
}

func NewStorage(rateLimit int, timeout int, readCount int, done chan int) *Storage {
	s := &Storage{
		timeout:   timeout,
		rateLimit: rate.NewLimiter(rate.Every(1*time.Second), rateLimit), // X requests / sec
		records: &Records{
			Records: make([]int, 0),
		},
		readCount: readCount,
	}

	go func() {
		for now := range time.Tick(time.Duration(timeout) * time.Second) {
			fmt.Printf("Processing cache: %v \n", now)
			s.processCache(done)
		}
	}()
	return s
}
func (s *Storage) processCache(done chan int) {

	s.records.lock.Lock()
	defer s.records.lock.Unlock()

	records := s.records.Records

	if records != nil && len(records) > 0 {
		fmt.Printf("Writing %d cached records to storage \n", len(records))
		s.records.Clear()
		go s.write(records, done)
	} else {
		fmt.Printf("Records cache is empty \n")
	}

}
func (s *Records) Clear() {
	s.Records = make([]int, 0)

}

func (s *Storage) write(records []int, done chan<- int) {
	ctx := context.Background()
	s.rateLimit.Wait(ctx)
	// var str string
	// for _, v := range records {
	// 	str += strconv.Itoa(v) + ", "
	// }
	fmt.Printf("succesfully written : %v \n", len(records))

	// pointer := records[len(records)-1] + 1
	// fmt.Printf("pointer : %v \n", pointer)
	// if s.readCount == pointer {
	// 	fmt.Printf("sending done : %v \n", pointer)
	// 	done <- 1
	// }

}

func (s *Storage) Write(n int, done chan<- int) (int, error) {
	s.records.lock.Lock()
	defer s.records.lock.Unlock()

	records := s.records.Records
	if len(records) >= 10000 {
		// fmt.Printf("Clearing cache : %v \n", records)
		s.records.Clear()
		go s.write(records, done)
	}
	// fmt.Printf("append number: %v \n", n)
	s.records.Records = append(s.records.Records, n)

	return 1, nil
}
