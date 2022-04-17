package generate

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/rs/zerolog/log"
)

//Function to generate a random cert and key for use in example operations.
func GenerateCert() {

}

func generateKeyPair() *rsa.PrivateKey {
	keyPair, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to generate key")
	}
	return keyPair
}
