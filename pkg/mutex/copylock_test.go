package mutex_test

import (
	"fmt"
	"sync"
	"testing"
)

type Counter struct {
	count int
	mu    *sync.Mutex
}

func NewCounter(mu *sync.Mutex) *Counter {
	return &Counter{mu: mu}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *Counter) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}
func (c *Counter) Print() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("\n counter:%d", c.count)
}

type CounterClient struct {
	*Counter
}

func NewCounterClient() *CounterClient {
	mu := sync.Mutex{}
	counter := NewCounter(&mu)
	fmt.Printf("\n counter mutex pointer address:%p", counter.mu)
	return &CounterClient{Counter: counter}
}

func TestCounter(t *testing.T) {
	counterClient := NewCounterClient()
	fmt.Printf("\n counter client mutex pointer address:%p", counterClient.mu)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			counterClient.Inc()
			counterClient.Print()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			counterClient.Dec()
			counterClient.Print()

		}
		wg.Done()
	}()
	wg.Wait()
	counterClient.Print()

}
