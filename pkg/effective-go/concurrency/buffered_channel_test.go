package concurrency_test

import (
	"go-mastery/pkg/effective-go/concurrency"
	"testing"
)

func TestDoWithBufferedChannel(t *testing.T) {
	concurrency.DoWithBufferedChannel(0)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.DoWithBufferedChannel(1)
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrency.DoWithBufferedChannel(1000)
	}
}
