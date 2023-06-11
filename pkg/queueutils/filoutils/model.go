package filoutils

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queueutils/blockutils"
	"github.com/google/uuid"
)

type Filo struct {
	id        uuid.UUID
	head      *blockutils.Block
	tail      *blockutils.Block
	numBlocks int
}
