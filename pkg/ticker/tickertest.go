package ticker

import (
	"log"
	"sync"
	"time"
)

func DoTickerTest(wg *sync.WaitGroup)  {
	//defer wg.Done()
	go func() {
		//for now := range time.Tick(time.Duration(1) * time.Second) {
		//	log.Println("Processing cache: ", now)
		//}
		for {
			select {
			case now :=<- time.Tick(time.Duration(1) * time.Second):
				log.Println("Processing cache: ", now)
			}
		}
	}()
}

