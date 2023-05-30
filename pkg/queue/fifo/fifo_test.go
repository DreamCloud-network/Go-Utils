package fifo

import (
	"log"
	"math/rand"
	"testing"

	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
)

func TestFifo(t *testing.T) {
	fifo := New()

	if fifo == nil {
		t.Errorf("New() returned nil")
	}

	log.Println("Fifo ID: ", fifo.GetID())

	log.Println("Remove a block from an empty queue")
	// Remove a block from an empty queue
	var newBlock *block.Block

	newBlock = fifo.Pop()

	if newBlock != nil {
		t.Errorf("Pop() returned a block from an empty queue")
	}

	if newBlock == nil {
		log.Println("Pop() returned nil")
	}

	log.Println("Add and remove a block")
	// Add and remove a block
	fifo.Push(block.New("Test Data 1"))

	log.Println("Fifo size: ", fifo.GetNumBlocks())

	newBlock = fifo.Pop()

	log.Println("Block ID: ", newBlock.GetID())
	log.Println("Block Data: ", newBlock.Data.(string))

	log.Println("Fifo size: ", fifo.GetNumBlocks())

	log.Println("Add and remove multiple blocks")
	// Add and remove multiple blocks
	for i := 0; i < rand.Intn(10000); i++ {
		fifo.Push(block.New(i))
	}

	log.Println("Filo size: ", fifo.GetNumBlocks())

	for fifo.GetNumBlocks() > 0 {
		fifo.Pop()
	}

	log.Println("Fifo size: ", fifo.GetNumBlocks())
}
