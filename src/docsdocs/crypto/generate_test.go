package crypto

import (
	"testing"
)

func TestNewCrypto(t *testing.T) {
	_, err := NewCrypto()
	if err != nil {
		t.Errorf("failed new crypto")
	}

}

func TestGetPrivate(t *testing.T) {
	size := 2048
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Errorf("failed new crypto")
	}
	privKey := docCrypto.GetPrivate()
	if docCrypto.GetPrivate() == nil {
		t.Errorf("private key not exist")
	}
	if bits := privKey.N.BitLen(); bits != size {
		t.Errorf("private key too short (%d vs %d)", bits, size)
	}

}
func TestGetPublic(t *testing.T) {
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Errorf("failed new crypto")
	}

	if docCrypto.GetPublic().E == 0 {
		t.Errorf("public key not exist")
	}
}
