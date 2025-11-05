package blockbucket

import (
	"log"
	"testing"

	"github.com/DreamCloud-network/Go-Utils/pkg/queueutils/block"
)

func TestBlockBucket(t *testing.T) {

	log.Println("Create a new Block")
	newBlock := GetBlock("Test Data 1")

	if newBlock == nil {
		t.Errorf("NewBlock() returned nil")
		return
	}

	log.Println("Block ID: ", newBlock.GetID())
	if newBlock.Data == nil {
		t.Errorf("NewBlock() returned a block with nil data")
		return
	} else {
		log.Print("Block Data: ", newBlock.Data.(string))
	}

	log.Println("Return the block")
	ReturnBlock(newBlock)

	log.Println("Get the block to use again")
	newBlock = GetBlock("Test Data 2")

	if newBlock == nil {
		t.Errorf("NewBlock() returned a block with nil data")
		return
	}
	log.Println("Block ID: ", newBlock.GetID())
	log.Print("Block Data: ", newBlock.Data.(string))

	log.Println("Testing a lot of blocks")

	vetBlocks := make([]*block.Block, 0)
	for cont := 0; cont < 10; cont++ {
		newBlock := GetBlock(cont)
		log.Println("Block ID: ", newBlock.GetID())
		log.Print("Block Data: ", newBlock.Data.(int))
		vetBlocks = append(vetBlocks, newBlock)
	}

	for _, dataBlock := range vetBlocks {
		ReturnBlock(dataBlock)
	}
}
