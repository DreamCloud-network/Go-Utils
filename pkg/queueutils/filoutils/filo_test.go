package filoutils

import (
	"log"
	"math/rand"
	"testing"

	"github.com/GreenMan-Network/Go-Utils/pkg/queueutils/blockutils"
)

func TestFilo(t *testing.T) {
	filo := New()

	if filo == nil {
		t.Errorf("New() returned nil")
	}

	log.Println("Filo ID: ", filo.GetID())

	log.Println("Remove a block from an empty queue")
	// Remove a block from an empty queue
	var newBlock *blockutils.Block

	newBlock = filo.Pop()

	if newBlock != nil {
		t.Errorf("Pop() returned a block from an empty queue")
	}

	if newBlock == nil {
		log.Println("Pop() returned nil")
	}

	log.Println("Add and remove a block")
	// Add and remove a block
	filo.Push(blockutils.New("Test Data 1"))

	log.Println("Filo size: ", filo.GetNumBlocks())

	newBlock = filo.Pop()

	log.Println("Block ID: ", newBlock.GetID())
	log.Println("Block Data: ", newBlock.Data.(string))

	log.Println("Filo size: ", filo.GetNumBlocks())

	log.Println("Add and remove multiple blocks")
	// Add and remove multiple blocks
	for i := 0; i < rand.Intn(10000); i++ {
		filo.Push(blockutils.New(i))
	}

	log.Println("Filo size: ", filo.GetNumBlocks())

	for filo.GetNumBlocks() > 0 {
		filo.Pop()
	}

	log.Println("Filo size: ", filo.GetNumBlocks())
}
