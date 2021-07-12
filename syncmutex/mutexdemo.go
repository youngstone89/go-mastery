package main

import (
	"sync"
)

var(
	Count int = 0
	Lock *sync.Mutex = &sync.Mutex{}
)

func main() {
	var wg *sync.WaitGroup
	Add1Mutex(wg)
	Sub1Mutex(wg)
	println(Count)
}
func Add1(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		Count += 1
	}
}

func Add1Mutex(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		Lock.Lock()
		Count += 1
		Lock.Unlock()
	}
}

func Add1Sem(wg *sync.WaitGroup, sem chan struct{}){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		sem <- struct{}{}
		Count += 1
		<-sem
	}
}

func Sub1(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		Count -= 1
	}
}

func Sub1Mutex(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 0; i < 10000; i++{
		Lock.Lock()
		Count -= 1
		Lock.Unlock()
	}
}