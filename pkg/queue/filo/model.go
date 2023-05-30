package filo

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
	"github.com/google/uuid"
)

type Filo struct {
	id        uuid.UUID
	head      *block.Block
	tail      *block.Block
	numBlocks int
}
