package memorydb

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("data not found")

type MemoryDB struct {
	memoryDb map[string][]byte
	mx       *sync.Mutex
}
