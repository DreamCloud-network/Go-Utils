package blockutils

import (
	"log"
	"testing"
)

func TestBlock(t *testing.T) {
	block := New("teste")

	log.Println("Block ID: ", block.GetID())
	log.Println("Block Test: ", block.Data.(string))
}
