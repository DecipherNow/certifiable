package authorities

import (
	"crypto/rsa"

	"github.com/deciphernow/certifiable/pems"
)

// Leaf represents a leaf node in a certificate chain.
type Leaf struct {
	certificate []byte
	key         *rsa.PrivateKey
}

// EncodeCertificate returns the PEM encoded certificate for the leaf.
func (l *Leaf) EncodeCertificate() (string, error) {
	return pems.EncodeCertificate(l.certificate)
}

// EncodeKey returns the PEM encoded key for the leaf.
func (l *Leaf) EncodeKey() (string, error) {
	return pems.EncodeKey(l.key)
}
