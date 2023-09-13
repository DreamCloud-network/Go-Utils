package filo

import (
	"log"
	"testing"
)

func TestFilo(t *testing.T) {
	testFilo := NewFilo(0)

	if testFilo == nil {
		t.Errorf("New() returned nil")
		return
	}

	log.Println("Filo ID: ", testFilo.GetID())

	log.Println("Remove a block from an empty queue")
	// Remove a block from an empty queue
	newdataInt := testFilo.Pop()

	if newdataInt != nil {
		t.Errorf("Pop() returned a block from an empty queue")
		return
	}

	if newdataInt == nil {
		log.Println("Pop() returned nil")
	}

	log.Println("Add and remove a block")
	// Add and remove a block
	testFilo.Push("Test Data 1")

	log.Println("Filo size: ", testFilo.GetSize())

	newdataInt = testFilo.Pop()
	if newdataInt == nil {
		t.Errorf("Pop() returned nil")
		return
	}

	log.Println("Block Data: ", newdataInt.(string))

	log.Println("Filo size: ", testFilo.GetSize())

	log.Println("Add and remove multiple blocks")
	// Add and remove multiple blocks
	for i := 0; i < 1000; i++ {
		testFilo.Push(i)
	}

	log.Println("Filo size: ", testFilo.GetSize())

	for testFilo.GetSize() > 0 {
		newdataInt = testFilo.Pop()
		log.Println("Block Data: ", newdataInt.(int))
	}

	log.Println("Filo size: ", testFilo.GetSize())
}
