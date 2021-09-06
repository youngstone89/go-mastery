package concurrency

import "testing"

func TestGetCounter(t *testing.T) {
	counter := getCounter()
	if counter != 5000 {
		t.Error("unexpected counter:", counter)
	}
}
