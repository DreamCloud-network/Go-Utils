package fifo

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
	"github.com/google/uuid"
)

// New - Creates a new FIFO queue
func New() *Fifo {
	return &Fifo{
		id:        uuid.New(),
		tail:      nil,
		numBlocks: 0,
	}
}

// Push - Pushes a new block to the end of the queue
func (fifo *Fifo) GetID() uuid.UUID {
	return fifo.id
}

// Push - Pushes a new block to the end of the queue
func (fifo *Fifo) Push(newBlock *block.Block) {
	if newBlock == nil {
		return
	}

	if fifo.tail == nil {
		fifo.tail = newBlock
	} else {
		fifo.tail.NextBlock = newBlock
		newBlock.PreviousBlock = fifo.tail

		fifo.tail = newBlock
	}
	fifo.numBlocks++
}

// Pop - Pops the first block from the queue
func (fifo *Fifo) Pop() *block.Block {
	if fifo.tail == nil {
		return nil
	}

	block := fifo.tail

	if fifo.numBlocks == 1 {
		fifo.tail = nil
	} else {
		fifo.tail = block.PreviousBlock
	}

	fifo.numBlocks--

	block.PreviousBlock = nil
	block.NextBlock = nil

	return block
}

// GetNumBlocks - Returns the number of blocks in the queue
func (fifo *Fifo) GetNumBlocks() int {
	return fifo.numBlocks
}
