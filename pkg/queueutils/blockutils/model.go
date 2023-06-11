package blockutils

import "github.com/google/uuid"

// Block - Struct to store a data block
type Block struct {
	id            uuid.UUID
	PreviousBlock *Block
	NextBlock     *Block
	Data          interface{}
}
