package generate

import "testing"

func TestGeneratePair(t *testing.T) {
	result := generateKeyPair()
	if result.Validate() != nil {
		t.Fatal("key generate failed")
	}
}
