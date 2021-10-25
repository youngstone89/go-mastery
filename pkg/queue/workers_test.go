package queue

import "testing"

func TesMain(m *testing.M) {
	m.Run()
}

func TestDoJob(t *testing.T) {
	DoJob(10000, 10)
}
