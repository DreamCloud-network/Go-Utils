package memorydb

import (
	"fmt"
	"sync"

	"github.com/GreenMan-Network/Go-Utils/pkg/datautils"
	"github.com/google/uuid"
)

// This a memony key value store for test purpose only

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		memoryDb: make(map[string][]byte),
		mx:       &sync.Mutex{},
	}
}

// Save a new data in the DB and return the key
func (db *MemoryDB) Push(value any) (string, error) {
	db.mx.Lock()
	db.mx.Unlock()

	key := uuid.New().String()

	err := db.Save(key, value)
	if err != nil {
		return "", fmt.Errorf("memorydb.SaveNewdata - %w", err)
	}

	return key, nil
}

// Save a new data in the DB or update if the key already exist
func (db *MemoryDB) Save(key string, value any) error {
	db.mx.Lock()
	db.mx.Unlock()

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
	db.mx.Unlock()

	value, ok := db.memoryDb[key]

	if !ok {
		return fmt.Errorf("memorydb.Read - Error reading data: %w", ErrNotFound)
	}

	err := datautils.Deserialize(value, target)
	if err != nil {
		return fmt.Errorf("memorydb.Read - Error deserializing data: %w", err)
	}

	return nil
}

// Read a data from the DB and delete it
func (db *MemoryDB) Pull(key string, target any) error {
	db.mx.Lock()
	db.mx.Unlock()

	err := db.Read(key, target)
	if err != nil {
		return fmt.Errorf("memorydb.Pull - %w", err)
	}

	delete(db.memoryDb, key)

	return nil
}
