package circularfifoqueue

import (
	"log"
	"testing"
)

func TestCircularFifoQueue(t *testing.T) {
	t.Log("Testing CircularFifoQueue")

	newTestCircularFifoQueue := NewCircularFifoQueue()

	log.Println("Testing pushing and poping one element...")
	log.Println("Pushing: ")
	num := 1
	newTestCircularFifoQueue.Push(num)

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	log.Println("Poping: ")
	newTestCircularFifoQueue.Pop()

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	if newTestCircularFifoQueue.GetNumBlocks() != 1 || newTestCircularFifoQueue.GetNumData() != 0 {
		t.Error("Error in pushing and poping one element")
		return
	}

	log.Println("Testing pushing and poping 10 elements element...")
	log.Println("Pushing: ")
	for i := 0; i < 10; i++ {
		newTestCircularFifoQueue.Push(i)
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	log.Println("Poping: ")
	for i := 0; i < 10; i++ {
		newTestCircularFifoQueue.Pop()
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	if newTestCircularFifoQueue.GetNumBlocks() != 10 || newTestCircularFifoQueue.GetNumData() != 0 {
		t.Error("Error in pushing and poping one element")
		return
	}

	log.Println("Testing pushing and poping more 20 elements element...")
	log.Println("Pushing: ")
	for i := 0; i < 20; i++ {
		newTestCircularFifoQueue.Push(i)
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	log.Println("Poping: ")
	for i := 0; i < 20; i++ {
		newTestCircularFifoQueue.Pop()
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	if newTestCircularFifoQueue.GetNumBlocks() != 20 || newTestCircularFifoQueue.GetNumData() != 0 {
		t.Error("Error in pushing and poping one element")
		return
	}

	log.Println("Testing setting max data 10 and pushing and poping more 20 elements element...")
	newTestCircularFifoQueue.SetMaxData(10)

	log.Println("Pushing: ")
	for i := 0; i < 20; i++ {
		popedData := newTestCircularFifoQueue.Push(i)
		if popedData != nil {
			log.Println("Poped data: ", popedData)
		}
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	log.Println("Poping: ")
	for i := 0; i < 20; i++ {
		newTestCircularFifoQueue.Pop()
	}

	log.Println("NumBlocks: ", newTestCircularFifoQueue.GetNumBlocks())
	log.Println("NumData: ", newTestCircularFifoQueue.GetNumData())

	if newTestCircularFifoQueue.GetNumBlocks() != 20 || newTestCircularFifoQueue.GetNumData() != 0 {
		t.Error("Error in pushing and poping one element")
		return
	}
}
