package filo

import (
	"sync"

	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/block"
	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/blockbucket"
	"github.com/google/uuid"
)

type Filo struct {
	id uuid.UUID

	sync.Mutex

	top *block.Block

	size int

	maxSize int
}

// New - Creates a new Filo queue
// If sizeMax is 0, the queue will have no size limit
func NewFilo(sizeMax int) *Filo {
	return &Filo{
		id:      uuid.New(),
		Mutex:   sync.Mutex{},
		top:     nil,
		size:    0,
		maxSize: sizeMax,
	}
}

// GetID - Returns the ID of the queue
func (filo *Filo) GetID() uuid.UUID {
	filo.Lock()
	defer filo.Unlock()

	return filo.id
}

func (filo *Filo) GetSize() int {
	filo.Lock()
	defer filo.Unlock()

	return filo.size
}

func (filo *Filo) GetMaxSize() int {
	filo.Lock()
	defer filo.Unlock()

	return filo.maxSize
}

// Set a new max size for the queue.
// If the queue size is greater than the new max size, all the necessary data will be discarted from the queue until the max size is reached.
func (filo *Filo) SetMaxSize(sizeMax int) {
	filo.Lock()
	defer filo.Unlock()

	filo.maxSize = sizeMax

	if filo.size > 0 {
		for filo.size > filo.maxSize {
			filo.internalPop()
		}
	}
}

// Push - Pushes a new block to the end of the queue
func (filo *Filo) Push(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	filo.Lock()
	defer filo.Unlock()

	var popedData interface{}

	if filo.size == filo.maxSize {
		popedData = filo.internalPop()
	}

	newBlock := blockbucket.GetBlock(data)

	if filo.top != nil {
		newBlock.PreviousBlock = filo.top
		filo.top.NextBlock = newBlock
	}

	filo.top = newBlock
	filo.size++

	return popedData
}

// Pop - Pops the first block from the queue
func (filo *Filo) internalPop() interface{} {
	if filo.top == nil {
		return nil
	}

	removedBlock := filo.top

	if filo.size == 1 {
		filo.top = nil
	} else {
		filo.top = filo.top.PreviousBlock
		filo.top.NextBlock = nil
	}

	data := removedBlock.Data
	blockbucket.ReturnBlock(removedBlock)

	filo.size--

	return data
}

func (filo *Filo) Pop() interface{} {
	filo.Lock()
	defer filo.Unlock()

	return filo.internalPop()
}
