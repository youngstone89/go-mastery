package concurrency_test

import (
	"go-mastery/pkg/concurrency"
	"testing"
)

func TestFanIn(t *testing.T) {
	concurrency.DoFanIn()
}
