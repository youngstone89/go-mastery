package mutex

import (
	"fmt"
	"sync"
)

func DoLockInForLoop() {

	var count = 0
	var mu sync.Mutex

	go func() {
		for {
			mu.Lock()
			count++
			fmt.Println("increase:", count)
			mu.Unlock()
			if count == 10 {
				return
			}
		}
	}()

	for {
		mu.Lock()
		fmt.Println("check", count)
		if count == 10 {
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}
