package blockbucket

import (
	"sync"

	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/block"
)

type BlockBucket struct {
	sync.Mutex

	pos *block.Block
}

var bucket *BlockBucket

// New - Creates a new block bucket
func initBucket() {
	bucket = &BlockBucket{
		Mutex: sync.Mutex{},

		pos: nil,
	}
}

// GetBlock - Creates a new block or get one from the bucket
func GetBlock(data interface{}) *block.Block {
	if bucket == nil {
		initBucket()
	}
	bucket.Lock()
	defer bucket.Unlock()

	if bucket.pos == nil {
		return block.New(data)
	} else {
		newBlock := bucket.pos

		bucket.pos = bucket.pos.PreviousBlock

		newBlock.Erase()
		newBlock.Data = data

		return newBlock
	}
}

// ReturnBlock - Erases and returns a block to the bucket
func ReturnBlock(block *block.Block) {
	if block == nil {
		return
	}

	if bucket == nil {
		initBucket()
	}
	bucket.Lock()
	defer bucket.Unlock()

	block.Erase()

	if bucket.pos != nil {
		block.PreviousBlock = bucket.pos
		bucket.pos.NextBlock = block
	}

	bucket.pos = block
}
