package gcloudutils

import (
	"context"
	"errors"
	"hash/crc32"
	"log"

	gcloudkms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func NewSymetricCryptographicChannel(cryptographyKey string) *SymetricCryptographicChannel {
	return &SymetricCryptographicChannel{
		cryptographyKey: cryptographyKey,
	}
}

// encryptSymmetric encrypts the input plaintext with the specified symmetric
// Cloud KMS key.
func (ch *SymetricCryptographicChannel) Encrypt(plaintext []byte) ([]byte, error) {

	// Create the client.
	ctx := context.Background()
	client, err := gcloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		err := errors.New("failed to create kms client")
		log.Println("kms.encryptSymmetric - Error creating kms client: " + err.Error())
		return nil, err
	}
	defer client.Close()

	// Optional but recommended: Compute plaintext's CRC32C.
	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}
	plaintextCRC32C := crc32c(plaintext)

	// Build the request.
	req := &kmspb.EncryptRequest{
		Name:            ch.cryptographyKey,
		Plaintext:       plaintext,
		PlaintextCrc32C: wrapperspb.Int64(int64(plaintextCRC32C)),
	}

	// Call the API.
	result, err := client.Encrypt(ctx, req)
	if err != nil {
		err := errors.New("failed to encrypt")
		log.Println("kms.encryptSymmetric - Error encrypting: " + err.Error())
		return nil, err
	}

	// Optional, but recommended: perform integrity verification on result.
	// For more details on ensuring E2E in-transit integrity to and from Cloud KMS visit:
	// https://cloud.google.com/kms/docs/data-integrity-guidelines
	if !result.VerifiedPlaintextCrc32C {
		err := errors.New("encrypt: request corrupted in-transit")
		log.Println("kms.encryptSymmetric - Error encrypting: " + err.Error())
		return nil, err
	}
	if int64(crc32c(result.Ciphertext)) != result.CiphertextCrc32C.Value {
		err := errors.New("encrypt: response corrupted in-transit")
		log.Println("kms.encryptSymmetric - Error encrypting: " + err.Error())
		return nil, err
	}

	return result.Ciphertext, nil
}

// decryptSymmetric will decrypt the input ciphertext bytes using the specified symmetric key.
func (ch *SymetricCryptographicChannel) Decrypt(ciphertext []byte) ([]byte, error) {
	// Create the client.
	ctx := context.Background()
	client, err := gcloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		err := errors.New("failed to create kms client")
		log.Println("kms.decryptSymmetric - Error creating kms client: " + err.Error())
		return nil, err
	}
	defer client.Close()

	// Optional, but recommended: Compute ciphertext's CRC32C.
	crc32c := func(data []byte) uint32 {
		t := crc32.MakeTable(crc32.Castagnoli)
		return crc32.Checksum(data, t)
	}
	ciphertextCRC32C := crc32c(ciphertext)

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:             ch.cryptographyKey,
		Ciphertext:       ciphertext,
		CiphertextCrc32C: wrapperspb.Int64(int64(ciphertextCRC32C)),
	}

	// Call the API.
	result, err := client.Decrypt(ctx, req)
	if err != nil {
		err := errors.New("failed to decrypt")
		log.Println("kms.decryptSymmetric - Error decrypting: " + err.Error())
		return nil, err
	}

	// Optional, but recommended: perform integrity verification on result.
	// For more details on ensuring E2E in-transit integrity to and from Cloud KMS visit:
	// https://cloud.google.com/kms/docs/data-integrity-guidelines
	if int64(crc32c(result.Plaintext)) != result.PlaintextCrc32C.Value {
		err := errors.New("decrypt: request corrupted in-transit")
		log.Println("kms.decryptSymmetric - Error decrypting: " + err.Error())
		return nil, err
	}

	return result.Plaintext, nil
}
