package fifoutils

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queueutils/blockutils"
	"github.com/google/uuid"
)

// Fifo - Struct defining a FIFO queue
type Fifo struct {
	id        uuid.UUID
	tail      *blockutils.Block
	numBlocks int
}
