package memorydb

import (
	"log"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int
}

func TestMemoryDB(t *testing.T) {

	testData := TestStruct{
		Name: "Test",
		Age:  10,
	}

	db := NewMemoryDB()

	id, err := db.Push(testData)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println("Data pushed with id:", id)

	var testReadedData TestStruct
	err = db.Read(id, &testReadedData)
	if err != nil {
		t.Error(err)
		return
	}

	if (testReadedData.Name != testData.Name) || (testReadedData.Age != testData.Age) {
		t.Error("Data not equal")
		return
	}

	var testPullData TestStruct
	err = db.Pull(id, &testPullData)
	if err != nil {
		t.Error(err)
		return
	}

	if (testPullData.Name != testData.Name) || (testPullData.Age != testData.Age) {
		t.Error("Data not equal")
		return
	}
}
