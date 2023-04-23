package elc_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRLock(t *testing.T) {
	a := 0

	lock := sync.RWMutex{}

	for i := 1; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			fmt.Printf("Lock: from go routine %d: a = %d\n", i, a)
			time.Sleep(time.Second)
			lock.Unlock()
		}(i)
	}

	b := 0

	for i := 11; i < 20; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("RLock: from go routine %d: b = %d\n", i, b)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)
	}

	<-time.After(time.Second * 10)
}
