package gcloudutils

import (
	"log"
	"testing"
)

func TestSymetricCryptographicChannel_Encrypt(t *testing.T) {
	testKey := "projects/greenman-387219/locations/us-central1/keyRings/greenman-key-ring/cryptoKeys/main-key"
	simetricalCryptographicChannel := NewSymetricCryptographicChannel(testKey)

	plaintext := []byte("Hello World!")

	ciphertext, err := simetricalCryptographicChannel.Encrypt(plaintext)
	if err != nil {
		t.Errorf("SymetricCryptographicChannel.Encrypt() error = %v", err)
		return
	}

	log.Println("Encrypted: " + string(ciphertext))

	decrypted, err := simetricalCryptographicChannel.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("SymetricCryptographicChannel.Decrypt() error = %v", err)
		return
	}

	log.Println("Decrypted: " + string(decrypted))

	if string(plaintext) != string(decrypted) {
		t.Errorf("SymetricCryptographicChannel.Decrypt() error = %v", err)
		return
	}

}
