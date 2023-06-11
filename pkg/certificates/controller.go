package certificates

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"os"
)

// SavePEMFile - Saves a private key to a PEM file.
func SavePEMFile(filePath string, privateKey *rsa.PrivateKey) error {

	pemPrivateFile, err := os.Create(filePath)
	if err != nil {
		log.Println("certificates.SavePEMFile - Error creating PEM file: " + filePath + ":" + err.Error())
		return err
	}

	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		log.Println("certificates.SavePEMFile - Error encoding PEM file: " + filePath + ":" + err.Error())
		return err
	}
	pemPrivateFile.Close()

	return nil
}

// LoadRSAPrivateKeyFromFile - Loads a private key from a PEM file.
func LoadRSAPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {

	pemPrivateFile, err := os.Open(filePath)
	if err != nil {
		log.Println("certificates.LoadRSAPrivateKeyFromFile - Error opening PEM file: " + filePath + ":" + err.Error())
		return nil, err
	}
	defer pemPrivateFile.Close()

	pemFileInfo, _ := pemPrivateFile.Stat()
	var size = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := pemBytes[:size]
	_, err = pemPrivateFile.Read(buffer)
	if err != nil {
		log.Println("certificates.LoadRSAPrivateKeyFromFile - Error reading PEM file: " + filePath + ":" + err.Error())
		return nil, err
	}

	data, _ := pem.Decode(buffer)
	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		log.Println("certificates.LoadRSAPrivateKeyFromFile - Error parsing PEM file: " + filePath + ":" + err.Error())
		return nil, err
	}

	return privateKey, nil
}

// LoadCertificateFromFile - Loads a certificate from a file.
func LoadCertificateFromFile(filePath string) ([]byte, error) {

	certificateFile, err := os.Open(filePath)
	if err != nil {
		log.Println("certificates.LoadCertificateFromFile - Error opening file: " + filePath + ":" + err.Error())
		return nil, err
	}
	defer certificateFile.Close()

	certificateFileInfo, _ := certificateFile.Stat()
	var size = certificateFileInfo.Size()
	certificateBytes := make([]byte, size)

	buffer := certificateBytes[:size]
	_, err = certificateFile.Read(buffer)
	if err != nil {
		log.Println("certificates.LoadCertificateFromFile - Error reading file: " + filePath + ":" + err.Error())
		return nil, err
	}

	data, _ := pem.Decode(buffer)
	//cert, err := x509.ParseCertificate(data.Bytes)

	return data.Bytes, nil
}

// SaveCertificateToFile - Saves a certificate to a file.
func SaveCertificateToFile(filePath string, certificate []byte) error {

	// pem encode
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate,
	})

	certificateFile, err := os.Create(filePath)
	if err != nil {
		log.Println("certificates.SaveCertificateFile - Error creating certificate file: " + filePath + ":" + err.Error())
		return err
	}

	var certificateBlock = &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate,
	}

	err = pem.Encode(certificateFile, certificateBlock)
	if err != nil {
		log.Println("certificates.SaveCertificateFile - Error encoding PEM file: " + filePath + ":" + err.Error())
		return err
	}
	certificateFile.Close()

	return nil
}

// GetPublicKey - Returns the public key from a private key.
func GetPublicKey(priv any) any {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

// GenereteSerialNumber - Generates a random serial number for a certificate.
func GenereteSerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
		return nil
	}

	return serialNumber
}
