package workerpools_test

import (
	"context"
	"fmt"
	"go-mastery/pkg/workerpools"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
)

func Test_Shutdown(t *testing.T) {
	consumer := workerpools.Consumer{
		IngestChan: make(chan int, 1),
		JobsChan:   make(chan int, workerpools.WorkerPoolSize),
	}

	// Simulate external lib sending us 10 events per second
	producer := workerpools.Producer{CallbackFunc: consumer.CallbackFunc}
	go producer.Start()

	// Set up cancellation context and waitgroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Start consumer with cancellation context passed
	go consumer.StartConsumer(ctx)

	// Start workers and Add [workerPoolSize] to WaitGroup
	wg.Add(workerpools.WorkerPoolSize)
	for i := 0; i < workerpools.WorkerPoolSize; i++ {
		go consumer.WorkerFunc(wg, i)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	// Handle shutdown
	fmt.Println("*********************************\nShutdown signal received\n*********************************")
	cancelFunc() // Signal cancellation to context.Context
	wg.Wait()    // Block here until are workers are done

	fmt.Println("All workers done, shutting down!")

}
