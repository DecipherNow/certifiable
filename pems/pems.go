package pems

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

// EncodeCertificate PEM encodes a DER certificate.
func EncodeCertificate(certificate []byte) (string, error) {

	buffer := bytes.NewBuffer(nil)

	err := pem.Encode(buffer, &pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	if err != nil {
		return "", errors.Wrap(err, "error encoding certificate")
	}

	return buffer.String(), nil
}

// EncodeKey PEM encodes an RSA key.
func EncodeKey(key *rsa.PrivateKey) (string, error) {

	buffer := bytes.NewBuffer(nil)

	err := pem.Encode(buffer, &pem.Block{Type: "RSA PRIVATE_KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return "", errors.Wrap(err, "error encoding key")
	}

	return buffer.String(), nil
}
