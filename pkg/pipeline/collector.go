package pipeline

import (
	"fmt"
	"sync"
)

type (
	jsonObj = map[string]interface{}
	jsonArr = []interface{}
)

type Collector struct {
	consumer   *Consumer
	processors []*Processor
}

func (c *Collector) Start(wg *sync.WaitGroup, queueDepth int, done chan int) {
	defer wg.Done()

	stream := make(chan interface{}, queueDepth)
	errc := make(chan error)

	go c.consumer.Read(stream, errc)

	for {
		select {
		case n := <-stream:
			// fmt.Printf("New Stream:%v \n", n)
			wg.Add(1)
			go c.ProcessRecord(n.(int), done)
		case err := <-errc:
			fmt.Printf("%s", err)
			return
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

func NewCollector(consumer *Consumer, processors []*Processor) *Collector {
	c := &Collector{
		consumer:   consumer,
		processors: processors,
	}
	return c
}

func (c *Collector) ProcessRecord(n int, done chan int) {

	var wg sync.WaitGroup

	for _, p := range c.processors {
		wg.Add(1)
		go func(wg *sync.WaitGroup, p *Processor) {
			defer wg.Done()
			err := p.ProcessRecord(n, done)
			if err != nil {
				fmt.Printf("Error processing record: %v", err)
			}
		}(&wg, p)
	}
	wg.Wait()
}
