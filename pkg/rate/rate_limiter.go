package rate

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func DoRateLimit() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("requests", req, time.Now())
	}
}

func DoBurstyLimit() {

	burstyLimiter := make(chan time.Time, 5)

	for i := 0; i < 5; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func DoRateLimitWithGOPackage() {

	// target_per_second := 10
	target_per_second := 1
	// target_per_second := 100
	sec := 1.0
	interval_sec := sec / float64(target_per_second)
	interval_ms := interval_sec * 1000
	rateLimiter := rate.NewLimiter(rate.Every(time.Duration(interval_ms)*time.Millisecond), 1)

	ctx := context.Background()

	_process := func(i int, wg *sync.WaitGroup) {
		defer wg.Done()
		start := time.Now()
		rateLimiter.Wait(ctx)
		fmt.Println(i, "waited", time.Since(start).Seconds())
	}
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go _process(i, &wg)
	}
	// for i := 0; i < 250; i++ {
	// 	wg.Add(1)
	// 	go _process(i, &wg)
	// }
	// for i := 250; i < 500; i++ {
	// 	wg.Add(1)
	// 	go _process(i, &wg)
	// }
	// for i := 500; i < 750; i++ {
	// 	wg.Add(1)
	// 	go _process(i, &wg)
	// }
	// for i := 750; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go _process(i, &wg)
	// }
	wg.Wait()
	fmt.Println("took", time.Since(start).Seconds())
}
