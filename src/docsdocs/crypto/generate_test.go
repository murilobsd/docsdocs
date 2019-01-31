package crypto

import "testing"

func TestNewCrypto(t *testing.T) {
	size := 2048
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Errorf("failed new crypto")
	}
	privKey := docCrypto.GetPrivate()
	if bits := privKey.N.BitLen(); bits != size {
		t.Errorf("private key too short (%d vs %d)", bits, size)
	}
}
