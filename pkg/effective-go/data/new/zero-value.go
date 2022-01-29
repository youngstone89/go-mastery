package new

import (
	"bytes"
	"sync"
)

type SyncedBuffer struct {
	lock   *sync.Mutex
	buffer *bytes.Buffer
}

func (sb *SyncedBuffer) Lock() *sync.Mutex {
	return sb.lock
}

func (sb *SyncedBuffer) Buffer() *bytes.Buffer {
	return sb.buffer
}
