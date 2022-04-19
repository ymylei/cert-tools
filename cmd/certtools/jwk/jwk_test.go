package jwk

import (
	"testing"

	"github.com/ymylei/cert-tools/pkg/tools"
)

func TestConvertPubToJWK(t *testing.T) {
	testK, err := tools.GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	err = ConvertPubToJWK(testK)
	if err != nil {
		t.Fatal(err)
	}
}
