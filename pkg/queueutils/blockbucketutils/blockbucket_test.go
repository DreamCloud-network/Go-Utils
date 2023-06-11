package blockbucketutils

import (
	"log"
	"testing"
)

func TestBlockBucket(t *testing.T) {
	Init()

	log.Println("Create a new Block")
	newBlock := NewBlock("Test Data 1")

	if newBlock == nil {
		t.Errorf("NewBlock() returned nil")
		return
	}

	log.Println("Block ID: ", newBlock.GetID())
	if newBlock.Data == nil {
		t.Errorf("NewBlock() returned a block with nil data")
	} else {
		log.Print("Block Data: ", newBlock.Data.(string))
	}

	log.Println("Return the block")
	ReturnBlock(newBlock)
	log.Println("Bucket size: ", GetNumBlocks())

	log.Println("Get the block to use again")
	newBlock = NewBlock("Test Data 2")

	if newBlock == nil {
		t.Errorf("NewBlock() returned a block with nil data")
	}
	log.Println("Block ID: ", newBlock.GetID())
}
