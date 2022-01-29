package waitgroups

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func DoBasic() {
	var wg sync.WaitGroup
	var urls = []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
	for i := range urls {
		// WaitGroup Counter++ when new goroutine is called
		wg.Add(1)
		go retrieve(urls[i], &wg)
	}
	// Wait for the collection of goroutines to finish
	wg.Wait()
}

func retrieve(url string, wg *sync.WaitGroup) {
	// WaitGroup Counter-- when goroutine is finished
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	// print the status code from the response
	fmt.Println(url, res.StatusCode, end)

}
