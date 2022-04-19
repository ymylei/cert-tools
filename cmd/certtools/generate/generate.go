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

//Function to generate a ECDSA Key Pair and output as PEM encoded files for other use
func GenerateKeyPair(name string) error {
	key, err := generateKeyPair()
	if err != nil {
		return err
	}

	privByes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}
	pubByes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return err
	}

	privPem := marshallPem("EC PRIVATE KEY", privByes)
	pubPem := marshallPem("PUBLIC KEY", pubByes)

	err = writePemToFile(fmt.Sprintf("%s.key", name), privPem)
	if err != nil {
		return err
	}
	err = writePemToFile(fmt.Sprintf("%s.pub", name), pubPem)
	if err != nil {
		return err
	}
	return nil
}

func generateCert(key *ecdsa.PrivateKey) (*pem.Block, *pem.Block, error) {

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
	certBlock := marshallPem("CERTIFICATE", cert)

	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	keyBlock := marshallPem("EC PRIVATE KEY", keyBytes)

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

func marshallPem(t string, d []byte) *pem.Block {
	return &pem.Block{
		Type:  t,
		Bytes: d,
	}
}
