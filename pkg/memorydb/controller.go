package memorydb

import (
	"fmt"
	"log"
	"sync"

	"github.com/GreenMan-Network/Go-Utils/pkg/datautils"
)

// This a memony key value store for test purpose only

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		memoryDb: make(map[string][]byte),
		mx:       &sync.Mutex{},
	}
}

// Save a new data in the DB or update if the key already exist
func (db *MemoryDB) Push(key string, value any) error {
	db.mx.Lock()
	defer db.mx.Unlock()

	dataBytes, err := datautils.Serialize(value)
	if err != nil {
		return fmt.Errorf("memorydb.Save - Error serializing data: %w", err)
	}

	db.memoryDb[key] = dataBytes

	return nil
}

// Read a data from the DB
func (db *MemoryDB) Read(key string, target any) error {
	db.mx.Lock()
	defer db.mx.Unlock()

	value, ok := db.memoryDb[key]

	if !ok {
		err := ErrNotFound
		log.Println("memorydb.Read - Error reading data: ", err)
		return err
	}

	err := datautils.Deserialize(value, target)
	if err != nil {
		log.Println("memorydb.Read - Error deserializing data: ", err)
		return err
	}

	return nil
}

// Read a data from the DB and delete it
func (db *MemoryDB) Pull(key string, target any) error {
	err := db.Read(key, target)
	if err != nil {
		return fmt.Errorf("memorydb.Pull - %w", err)
	}

	db.mx.Lock()
	defer db.mx.Unlock()
	delete(db.memoryDb, key)

	return nil
}
