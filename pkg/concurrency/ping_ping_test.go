package concurrency_test

import (
	"go-mastery/pkg/concurrency"
	"testing"
)

func TestDoPingPong(t *testing.T) {
	concurrency.DoPingPong()
}
