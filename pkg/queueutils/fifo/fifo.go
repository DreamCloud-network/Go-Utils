package fifo

import (
	"sync"

	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/block"
	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/blockbucket"
	"github.com/google/uuid"
)

type Fifo struct {
	id uuid.UUID

	sync.Mutex

	head *block.Block

	size    int
	maxSize int
}

// If sizeMax is 0, the queue will have no size limit
func NewFifo(sizeMax int) *Fifo {
	newFifo := &Fifo{
		id: uuid.New(),

		Mutex: sync.Mutex{},

		head: nil,

		size:    0,
		maxSize: sizeMax,
	}

	return newFifo
}

func (fifo *Fifo) GetID() uuid.UUID {
	fifo.Lock()
	defer fifo.Unlock()

	return fifo.id
}

func (fifo *Fifo) GetSize() int {
	fifo.Lock()
	defer fifo.Unlock()

	return fifo.size
}

func (fifo *Fifo) GetMaxSize() int {
	fifo.Lock()
	defer fifo.Unlock()

	return fifo.maxSize
}

// Set a new max size for the queue.
// If the queue size is greater than the new max size, all the necessary data will be discarted from the queue until the max size is reached.
func (fifo *Fifo) SetMaxSize(sizeMax int) {
	fifo.Lock()
	defer fifo.Unlock()

	fifo.maxSize = sizeMax

	if fifo.size > 0 {
		for fifo.size > fifo.maxSize {
			fifo.internalPop()
		}
	}
}

// Add data to the last position.
// If the queue is full, the first element will be removed and returned.
func (fifo *Fifo) Push(data interface{}) interface{} {
	fifo.Lock()
	defer fifo.Unlock()

	if data == nil {
		return nil
	}

	var popedData interface{}

	if fifo.size == fifo.maxSize {
		popedData = fifo.internalPop()
	}

	newBlock := blockbucket.GetBlock(data)

	// Add the block in the end of the queue
	if fifo.head == nil {
		fifo.head = newBlock

		fifo.head.PreviousBlock = fifo.head
		fifo.head.NextBlock = fifo.head
	} else {
		newBlock.PreviousBlock = fifo.head.PreviousBlock
		newBlock.NextBlock = fifo.head

		fifo.head.PreviousBlock.NextBlock = newBlock
		fifo.head.PreviousBlock = newBlock
	}

	fifo.size++

	return popedData
}

func (fifo *Fifo) internalPop() interface{} {
	if fifo.head == nil {
		return nil
	}

	popedData := fifo.head.Data

	if fifo.size == 1 {
		blockbucket.ReturnBlock(fifo.head)
		fifo.head = nil
	} else {
		fifo.head.PreviousBlock.NextBlock = fifo.head.NextBlock
		fifo.head.NextBlock.PreviousBlock = fifo.head.PreviousBlock

		removedBLock := fifo.head
		fifo.head = fifo.head.NextBlock

		blockbucket.ReturnBlock(removedBLock)
	}

	fifo.size--

	return popedData
}

func (fifo *Fifo) Pop() interface{} {
	fifo.Lock()
	defer fifo.Unlock()

	return fifo.internalPop()
}
