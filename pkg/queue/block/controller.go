package block

import "github.com/google/uuid"

func New(data interface{}) *Block {
	return &Block{
		id:            uuid.New(),
		PreviousBlock: nil,
		NextBlock:     nil,
		Data:          data,
	}
}

// GetID - Returns the ID of the block
func (block *Block) GetID() uuid.UUID {
	return block.id
}

// Erase - Erases the block
func (block *Block) Erase() {
	block.PreviousBlock = nil
	block.NextBlock = nil
	block.Data = nil
}
