package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"docsdocs/log"
	"encoding/gob"
	"errors"
)

var (
	errCreatingKey = errors.New("creating key fail")
)

func init() {
	gob.Register(rsa.PublicKey{})
}

// RSAKeySize given bit size using the random
const RSAKeySize int = 2048

// Encrypter ...
type Encrypter interface {
	// Encode(*rsa.PublicKey) ([]byte, error)
	// Decode([]byte) (*rsa.PublicKey, error)
	// WritePrivate() error
	// WritePublic() error
	GetPrivate() *rsa.PrivateKey
	GetPublic() rsa.PublicKey
	// ReadPrivate() (*rsa.PrivateKey, error)
	// ReadPublic() (*rsa.PublicKey, error)
}

// DocsCrypto This is struct to manager crypto files
type DocsCrypto struct {
	privateKey *rsa.PrivateKey
	log.Logger
}

// CreateKey generate RSA keypair
func CreateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, RSAKeySize)
}

// NewCrypto ...
func NewCrypto() (Encrypter, error) {
	priKey, err := CreateKey()
	if err != nil {
		return nil, errCreatingKey
	}
	docCrypto := &DocsCrypto{
		privateKey: priKey,
		Logger:     log.NewDocsLogger(),
	}
	return docCrypto, nil
}

// GetPrivate return privatekey
func (d *DocsCrypto) GetPrivate() *rsa.PrivateKey {
	d.Debug("getting private key")
	return d.privateKey
}

// GetPublic return public key
func (d *DocsCrypto) GetPublic() rsa.PublicKey {
	d.Debug("gettting public key")
	return d.privateKey.PublicKey
}
