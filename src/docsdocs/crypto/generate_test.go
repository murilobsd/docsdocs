package crypto

import (
	"bytes"
	"testing"
)

func TestCreateKey(t *testing.T) {
	size := 2048
	key, err := CreateKey()
	if err != nil {
		t.Error(err)
	}

	if bits := key.N.BitLen(); bits != size {
		t.Errorf("private key too short (%d vs %d)", bits, size)
	}
}

func TestNewCrypto(t *testing.T) {
	_, err := NewCrypto()
	if err != nil {
		t.Error(err)
	}

}

func TestGetPrivate(t *testing.T) {
	size := 2048
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Error(err)
	}
	privKey := docCrypto.GetPrivate()
	if docCrypto.GetPrivate() == nil {
		t.Error("failed new crypto")
	}
	if bits := privKey.N.BitLen(); bits != size {
		t.Errorf("private key too short (%d vs %d)", bits, size)
	}

}
func TestGetPublic(t *testing.T) {
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Error(err)
	}

	if docCrypto.GetPublic().E == 0 {
		t.Errorf("public key not exist")
	}
}

func TestExportPrivate(t *testing.T) {
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Error(err)
	}
	buf := &bytes.Buffer{}
	err = docCrypto.ExportPrivate(buf)
	if err != nil {
		t.Error(err)
	}
}

func TestExportPublic(t *testing.T) {
	docCrypto, err := NewCrypto()
	if err != nil {
		t.Error(err)
	}
	buf := &bytes.Buffer{}
	err = docCrypto.ExportPublic(buf)
	if err != nil {
		t.Error(err)
	}
}
