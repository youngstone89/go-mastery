package concurrency_test

import (
	"go-mastery/pkg/concurrency"
	"testing"
)

func TestDoServeAndLogWithWorkers(t *testing.T) {
	concurrency.DoServeAndLogWithWorkers()

}
