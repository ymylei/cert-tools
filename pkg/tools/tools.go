package tools

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

//GenerateCert creates a testing self-signed cert for playing with operations
func GenerateCert(key *ecdsa.PrivateKey) (*pem.Block, *pem.Block, error) {

	serial, err := rand.Int(rand.Reader, big.NewInt(10000000000))
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		Version:            1,
		SerialNumber:       serial,
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
	certBlock := MarshallPem("CERTIFICATE", cert)

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	keyBlock := MarshallPem("EC PRIVATE KEY", keyBytes)

	return keyBlock, certBlock, nil
}

//GenerateKeyPair creates a ECDSA key pair for use in operations
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	keyPair, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("unable to generate key pair: %v", err)
	}
	return keyPair, nil
}

//WritePemToFile lets you output the PEM as a file for use elsewhere with testing/otherwise
func WritePemToFile(filename string, block *pem.Block) error {
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

//MarshallPem simple wrapper for outputting PEM block
func MarshallPem(t string, d []byte) *pem.Block {
	return &pem.Block{
		Type:  t,
		Bytes: d,
	}
}
