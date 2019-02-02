package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"docsdocs/log"
	"encoding/gob"
	"encoding/pem"
	"errors"
	"io"
)

var (
	errCreatingKey        = errors.New("creating key fail")
	errExportPrivateKey   = errors.New("failed to encode private key")
	errSerializePublicKey = errors.New("failed to serializing public key")
	errExportPublicKey    = errors.New("failed to encode public key")
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
	ExportPrivate(io.Writer) error
	ExportPublic(io.Writer) error
	GetPrivate() *rsa.PrivateKey
	GetPublic() *rsa.PublicKey
	// ReadPrivateFile(r io.Reader) (*rsa.PrivateKey, error)
	// ReadPublicFile(r io.Reader) (*rsa.PublicKey, error)
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
	privKey, err := CreateKey()
	if err != nil {
		return nil, errCreatingKey
	}
	docCrypto := &DocsCrypto{
		privateKey: privKey,
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
func (d *DocsCrypto) GetPublic() *rsa.PublicKey {
	d.Debug("gettting public key")
	return d.privateKey.Public().(*rsa.PublicKey)
}

// ExportPrivate convert a private to pem and storage to disk
func (d *DocsCrypto) ExportPrivate(out io.Writer) error {
	d.Info("Exporting private key")
	if err := pem.Encode(out, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(d.GetPrivate()),
	}); err != nil {
		d.Error(err)
		return errExportPrivateKey
	}
	return nil
}

// ExportPublic convert a public key to pem and storage to disk
func (d *DocsCrypto) ExportPublic(out io.Writer) error {
	d.Info("Exporting public key")
	pubKey, err := x509.MarshalPKIXPublicKey(d.GetPublic())
	d.Debug("Serializing private key")
	if err != nil {
		d.Error(err)
		return errSerializePublicKey
	}
	if err = pem.Encode(out, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKey,
	}); err != nil {
		d.Error(err)
		return errExportPublicKey
	}
	return nil
}
