package syscall_test

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	var stopChan = make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		time.Sleep(time.Duration(10) * time.Second)
		stopChan <- os.Interrupt
	}()
	fmt.Println("waiting for signal")
	<-stopChan // wait for SIGINT
	fmt.Println("interrupted")
	// syscall.Do()
}
