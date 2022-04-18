package generate

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

//Function to generate a random self-signed cert and key for use in example operations.
func Generate(name string) error {
	key, err := generateKeyPair()
	if err != nil {
		return err
	}

	keyBlock, certBlock, err := generateCert(key)
	if err != nil {
		return err
	}

	err = writePemToFile(fmt.Sprintf("%s.key", name), keyBlock)
	if err != nil {
		return err
	}

	err = writePemToFile(fmt.Sprintf("%s.crt", name), certBlock)
	if err != nil {
		return err
	}
	return nil
}

func generateCert(key *ecdsa.PrivateKey) (*pem.Block, *pem.Block, error) {

	template := &x509.Certificate{
		Version:            1,
		SerialNumber:       big.NewInt(1),
		SignatureAlgorithm: x509.ECDSAWithSHA256,
		Subject: pkix.Name{
			CommonName:   "test.example.com",
			Organization: []string{"TestOrganization"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, key.Public(), key)
	if err != nil {
		return nil, nil, err
	}
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	keyBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyBytes,
	}

	return keyBlock, certBlock, nil
}

func generateKeyPair() (*ecdsa.PrivateKey, error) {
	keyPair, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("unable to generate key pair: %v", err)
	}
	return keyPair, nil
}

func writePemToFile(filename string, block *pem.Block) error {
	contents := &bytes.Buffer{}

	err := pem.Encode(contents, block)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, contents.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
