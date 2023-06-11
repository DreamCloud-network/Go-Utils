package datautils

import (
	"log"
	"testing"
)

type Teste struct {
	Name string
	Age  int
}

func TestEncodeDecode(t *testing.T) {
	// Create test struct
	var dataTeste = Teste{
		Name: "Teste",
		Age:  20,
	}

	// Serialize test struct
	dataTestBytes, err := Serialize(dataTeste)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println("Test data bytes: ", dataTestBytes)

	// Deserialize test struct
	var dataTestStruct Teste

	err = Deserialize(dataTestBytes, &dataTestStruct)
	if err != nil {
		t.Error(err)
		return
	}

	// Verify if values are equal
	if dataTestStruct.Name != dataTeste.Name {
		t.Error("Name not equal")
		return
	}

	if dataTestStruct.Age != dataTeste.Age {
		t.Error("Age not equal")
		return
	}
}
