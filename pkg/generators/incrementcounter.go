package generators

import (
	"fmt"
	"sync"
)

func incrementCounter() func() int {
	initializedNumber := 0
	return func() int {
		initializedNumber++
		return initializedNumber
	}
}

func DoIncrementCounter() {
	n1 := incrementCounter()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			n1()
		}(&wg)
	}
	wg.Wait()
	fmt.Println("n1 increment counter #1: ", n1())
	fmt.Println("n1 increment counter #2: ", n1())
	n2 := incrementCounter()
	fmt.Println("n2 increment counter #1: ", n2())
	fmt.Println("n1 increment counter #3: ", n1())
}
