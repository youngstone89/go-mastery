package pipeline

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Processor struct {
	storages  []*Storage
	rateLimit int
	sem       *semaphore.Weighted
	ctx       context.Context
}

func NewProcessor(storages []*Storage, rateLimit int) *Processor {
	p := &Processor{
		storages:  storages,
		sem:       semaphore.NewWeighted(int64(rateLimit)),
		ctx:       context.TODO(),
		rateLimit: rateLimit,
	}
	return p
}

func (p *Processor) ProcessRecord(n int, done chan int) error {

	p.SaveRecord(n, done)

	return nil
}

func (p *Processor) SaveRecord(n int, done chan int) {

	if err := p.sem.Acquire(p.ctx, 1); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v \n", err)
	} else {
		// Non-blocking write to storage...
		go func() {
			defer p.sem.Release(1)

			// Write to each configured storage object in parallel
			var wg sync.WaitGroup
			for _, s := range p.storages {
				wg.Add(1)
				go func(wg *sync.WaitGroup, s *Storage) {
					defer wg.Done()
					_, err := s.Write(n, done)
					if err != nil {
						fmt.Println(err.Error())
					}
				}(&wg, s)
			}
			wg.Wait()
		}()
	}
}
