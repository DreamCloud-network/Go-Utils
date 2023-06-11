package certificates

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"os"
	"testing"
	"time"
)

func TestSaveAndLoadPEMFile(t *testing.T) {
	testFIle := "./test.pem"

	// Clear the test file if exists
	err := os.Remove(testFIle)
	if err != nil {
		t.Errorf("Error removing test file: " + err.Error())
		return
	}

	// create our private and public key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Errorf("Error generating private key: " + err.Error())
		return
	}

	err = SavePEMFile(testFIle, privateKey)
	if err != nil {
		t.Errorf("Error saving pem file: %v", err)
		return
	}

	_, err = LoadRSAPrivateKeyFromFile(testFIle)
	if err != nil {
		t.Errorf("Error loading pem file: %v", err)
		return
	}
}

func TestSaveCertificateFile(t *testing.T) {
	testFile := "./test.crt"

	// Clear the test file if exists
	err := os.Remove(testFile)
	if err != nil {
		t.Errorf("Error removing test file: " + err.Error())
		return
	}

	// set up our CA certificate
	certificate := &x509.Certificate{
		SerialNumber: GenereteSerialNumber(),
		Subject: pkix.Name{
			Organization:  []string{"Organization"},
			Country:       []string{"Country"},
			Province:      []string{"Province"},
			Locality:      []string{"Locality"},
			StreetAddress: []string{"StreetAddress"},
			PostalCode:    []string{"PostalCode"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// create our private and public key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Errorf("Error generating private Key: %v", err)
		return
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificate, GetPublicKey(privateKey), privateKey)
	if err != nil {
		t.Errorf("error generating certificate: %v", err)
		return
	}

	err = SaveCertificateToFile(testFile, certificateBytes)
	if err != nil {
		t.Errorf("Error saving certificate file: %v", err)
		return
	}
}
