package jwk

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

func ConvertPubToJWK(key *ecdsa.PrivateKey) error {
	k, err := generateJwks(key)
	if err != nil {
		return err
	}
	err = marshallJwk(k)
	if err != nil {
		return err
	}
	return nil
}

func generateJwks(key *ecdsa.PrivateKey) (*jwk.Key, error) {
	jwkKey, err := jwk.FromRaw(key.PublicKey)
	if err != nil {
		return nil, err
	}

	if _, ok := jwkKey.(jwk.ECDSAPublicKey); !ok {
		return nil, fmt.Errorf("key not expected format of ECDSAPublicKey")
	}
	return &jwkKey, nil
}

func marshallJwk(key *jwk.Key) error {
	k := *key
	err := k.Set(jwk.KeyIDKey, "test")
	if err != nil {
		return err
	}
	jsonKey, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", jsonKey)
	return nil
}
