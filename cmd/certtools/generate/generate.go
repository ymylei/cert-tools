package generate

import (
	"crypto/x509"
	"fmt"

	"github.com/ymylei/cert-tools/pkg/tools"
)

//Function to generate a random self-signed cert and key for use in example operations.
func Generate(name string) error {
	key, err := tools.GenerateKeyPair()
	if err != nil {
		return err
	}

	keyBlock, certBlock, err := tools.GenerateCert(key)
	if err != nil {
		return err
	}

	err = tools.WritePemToFile(fmt.Sprintf("%s.key", name), keyBlock)
	if err != nil {
		return err
	}

	err = tools.WritePemToFile(fmt.Sprintf("%s.crt", name), certBlock)
	if err != nil {
		return err
	}
	return nil
}

//Function to generate a ECDSA Key Pair and output as PEM encoded files for other use
func GenerateKeyPair(name string) error {
	key, err := tools.GenerateKeyPair()
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

	privPem := tools.MarshallPem("EC PRIVATE KEY", privByes)
	pubPem := tools.MarshallPem("PUBLIC KEY", pubByes)

	err = tools.WritePemToFile(fmt.Sprintf("%s.key", name), privPem)
	if err != nil {
		return err
	}
	err = tools.WritePemToFile(fmt.Sprintf("%s.pub", name), pubPem)
	if err != nil {
		return err
	}
	return nil
}
