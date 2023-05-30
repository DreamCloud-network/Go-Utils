package fifo

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
	"github.com/google/uuid"
)

// Fifo - Struct defining a FIFO queue
type Fifo struct {
	id        uuid.UUID
	tail      *block.Block
	numBlocks int
}
