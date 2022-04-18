package generate

import (
	"encoding/pem"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestGeneratePair(t *testing.T) {
	_, err := generateKeyPair()
	if err != nil {
		t.Fatal("key generate failed")
	}
}

func TestGenerateCert(t *testing.T) {
	key, cert, err := GenerateCert()
	if err != nil {
		t.Fatalf("%v", err)
	}
	log.Info().Str("cert", string(pem.EncodeToMemory(cert))).Str("key", string(pem.EncodeToMemory(key))).Msg("output")
}

func TestWritePemToFile(t *testing.T) {
	_, cert, err := GenerateCert()
	if err != nil {
		t.Fatal(err)
	}

	filename := fmt.Sprintf("test_%d.pem", rand.Int())

	err = writePemToFile(filename, cert)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(filename)
	if err != nil {
		t.Error(err)
	}
}
