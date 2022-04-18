package generate

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestGeneratePair(t *testing.T) {
	_, err := generateKeyPair()
	if err != nil {
		t.Fatal("key generate failed")
	}
}

func TestGenerateCert(t *testing.T) {
	key, _ := generateKeyPair()

	_, _, err := generateCert(key)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestWritePemToFile(t *testing.T) {
	key, _ := generateKeyPair()

	_, cert, err := generateCert(key)
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
