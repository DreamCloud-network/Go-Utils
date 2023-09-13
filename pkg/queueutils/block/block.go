package block

import (
	"sync"

	"github.com/google/uuid"
)

type Block struct {
	id uuid.UUID

	sync.Mutex

	PreviousBlock *Block
	NextBlock     *Block
	Data          interface{}
}

func New(data interface{}) *Block {
	return &Block{
		id:            uuid.New(),
		Mutex:         sync.Mutex{},
		PreviousBlock: nil,
		NextBlock:     nil,
		Data:          data,
	}
}

// GetID - Returns the ID of the block
func (block *Block) GetID() uuid.UUID {
	block.Lock()
	defer block.Unlock()

	return block.id
}

// Erase - Erases the block
func (block *Block) Erase() {
	block.Lock()
	defer block.Unlock()

	block.PreviousBlock = nil
	block.NextBlock = nil
	block.Data = nil
}
