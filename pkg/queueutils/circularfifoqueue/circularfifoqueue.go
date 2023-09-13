package circularfifoqueue

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queueutils/blockutils"
	"github.com/google/uuid"
)

type CircularFifoQueue struct {
	id       uuid.UUID
	dataInit *blockutils.Block
	dataEnd  *blockutils.Block

	numBlocks int
	numData   int

	numMaxData int
	useMaxData bool
}

func NewCircularFifoQueue() *CircularFifoQueue {
	newCircularQueue := &CircularFifoQueue{
		id:        uuid.New(),
		dataInit:  nil,
		dataEnd:   nil,
		numBlocks: 0,
		numData:   0,
	}

	return newCircularQueue
}

func (circularQueue *CircularFifoQueue) GetID() uuid.UUID {
	return circularQueue.id
}

func (circularQueue *CircularFifoQueue) GetNumBlocks() int {
	return circularQueue.numBlocks
}

func (circularQueue *CircularFifoQueue) GetNumData() int {
	return circularQueue.numData
}

func (circularQueue *CircularFifoQueue) GetMaxData() int {
	return circularQueue.numMaxData
}

func (circularQueue *CircularFifoQueue) SetMaxData(numMaxData int) {
	circularQueue.numMaxData = numMaxData
	circularQueue.useMaxData = true
}

func (circularQueue *CircularFifoQueue) UnsetMaxData() {
	circularQueue.numMaxData = 0
	circularQueue.useMaxData = false
}

// Adds a new block with no data in the end of the queue
func (circularQueue *CircularFifoQueue) addNewBlock() {
	newBlock := blockutils.New(nil)

	if circularQueue.dataEnd == nil {
		newBlock.PreviousBlock = newBlock
		newBlock.NextBlock = newBlock

		circularQueue.dataInit = newBlock
	} else {
		newBlock.PreviousBlock = circularQueue.dataEnd
		newBlock.NextBlock = circularQueue.dataEnd.NextBlock

		circularQueue.dataEnd.NextBlock = newBlock
		newBlock.NextBlock.PreviousBlock = newBlock
	}

	circularQueue.dataEnd = newBlock
	circularQueue.numBlocks++
}

// Adds a new data with no data in the end of the queue. If the queue is full, the first data is removed
func (circularQueue *CircularFifoQueue) Push(data interface{}) interface{} {
	if data == nil {
		return nil
	}

	var popedData interface{}

	if circularQueue.useMaxData && circularQueue.numData == circularQueue.numMaxData {
		popedData = circularQueue.Pop()
	}

	if circularQueue.numBlocks == 0 {
		circularQueue.addNewBlock()
	} else if circularQueue.numData == circularQueue.numBlocks {
		circularQueue.addNewBlock()
	} else {
		circularQueue.dataEnd = circularQueue.dataEnd.NextBlock
	}

	circularQueue.dataEnd.Data = data
	circularQueue.numData++

	return popedData
}

// Pops the first data from the queue
func (circularQueue *CircularFifoQueue) Pop() interface{} {
	if circularQueue.dataInit == nil {
		return nil
	}

	if circularQueue.numData == 0 {
		return nil
	}

	data := circularQueue.dataInit.Data
	circularQueue.dataInit.Data = nil
	circularQueue.numData--

	circularQueue.dataInit = circularQueue.dataInit.NextBlock

	return data
}
