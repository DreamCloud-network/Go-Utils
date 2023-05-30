package fileutils

import (
	"log"
	"testing"
)

func TestReadJson(t *testing.T) {
	mapConfig, err := ReadJsonFile("config.json")
	if err != nil {
		t.Errorf("Error reading json file: %v", err)
	}

	log.Print(mapConfig)
}
