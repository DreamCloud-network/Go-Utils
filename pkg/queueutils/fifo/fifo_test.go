package fifo

import (
	"log"
	"testing"
)

func TestFifo(t *testing.T) {
	testFifo := NewFifo(0)

	if testFifo == nil {
		t.Errorf("New() returned nil")
		return
	}

	log.Println("Fifo ID: ", testFifo.GetID())

	log.Println("Remove a block from an empty queue")
	// Remove a block from an empty queue

	testData := testFifo.Pop()

	if testData != nil {
		t.Errorf("Pop() returned a block from an empty queue")
		return
	}

	if testData == nil {
		log.Println("Pop() returned nil")
	}

	log.Println("Add and remove a block")
	// Add and remove a block
	testFifo.Push("Test Data 1")

	log.Println("Fifo size: ", testFifo.GetSize())

	testDataInt := testFifo.Pop()
	if testDataInt == nil {
		t.Errorf("Pop() returned nil")
		return
	}

	log.Println("Block Data: ", testDataInt.(string))

	log.Println("Fifo size: ", testFifo.GetSize())

	log.Println("Add and remove multiple blocks")
	// Add and remove multiple blocks
	for i := 0; i < 1000; i++ {
		testFifo.Push(i)
	}

	log.Println("Filo size: ", testFifo.GetSize())

	for testFifo.GetSize() > 0 {
		testDataInt = testFifo.Pop()
		log.Println("Data: ", testDataInt.(int))
	}

	log.Println("Fifo size: ", testFifo.GetSize())

	log.Println("Test Set size")
	testFifo.SetMaxSize(10)
	// Add and remove multiple blocks
	for i := 0; i < 20; i++ {
		testFifo.Push(i)
	}
	log.Println("Filo size: ", testFifo.GetSize())
	for testFifo.GetSize() > 0 {
		testDataInt = testFifo.Pop()
		log.Println("Data: ", testDataInt.(int))
	}
	log.Println("Fifo size: ", testFifo.GetSize())

	log.Println("Test reset size to a value smaller than the actual size.")
	// Add and remove multiple blocks
	for i := 0; i < 20; i++ {
		testFifo.Push(i)
	}

	log.Println("Filo size: ", testFifo.GetSize())
	testFifo.SetMaxSize(5)
	log.Println("Filo size: ", testFifo.GetSize())

	for testFifo.GetSize() > 0 {
		testDataInt = testFifo.Pop()
		log.Println("Data: ", testDataInt.(int))
	}
	log.Println("Filo size: ", testFifo.GetSize())
}
