package concurrency_test

import (
	"go-mastery/pkg/concurrency"
	"testing"
	"time"
)

func TestFanOut(t *testing.T) {
	ch := make(chan int)
	go concurrency.Produce(ch, 1*time.Second)
	go concurrency.Consume(ch, 1)
	go concurrency.Consume(ch, 2)

	select {}
}
