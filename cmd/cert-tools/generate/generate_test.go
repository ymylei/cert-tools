package generate

import (
	"encoding/pem"
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
