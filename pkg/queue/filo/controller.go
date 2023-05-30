package filo

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
	"github.com/google/uuid"
)

func New() *Filo {
	return &Filo{
		id:        uuid.New(),
		head:      nil,
		tail:      nil,
		numBlocks: 0,
	}
}

// GetID - Returns the ID of the queue
func (filo *Filo) GetID() uuid.UUID {
	return filo.id
}

// Push - Pushes a new block to the end of the queue
func (filo *Filo) Push(newBlock *block.Block) {
	if newBlock == nil {
		return
	}

	if filo.head == nil {
		filo.head = newBlock
		filo.tail = newBlock
	} else {
		filo.tail.NextBlock = newBlock
		newBlock.PreviousBlock = filo.tail

		filo.tail = newBlock
	}
	filo.numBlocks++
}

// Pop - Pops the first block from the queue
func (filo *Filo) Pop() *block.Block {
	if filo.head == nil {
		return nil
	}

	block := filo.head

	if filo.numBlocks == 1 {
		filo.head = nil
		filo.tail = nil
	} else {
		filo.head = filo.head.NextBlock
	}

	filo.numBlocks--

	block.PreviousBlock = nil
	block.NextBlock = nil

	return block
}

// GetNumBlocks - Returns the number of blocks in the queue
func (filo *Filo) GetNumBlocks() int {
	return filo.numBlocks
}
