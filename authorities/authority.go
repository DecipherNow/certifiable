package authorities

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/deciphernow/certifiable/pems"
	"github.com/pkg/errors"
)

// Authority represetns a certificate authority.
type Authority struct {
	certificate []byte
	key         *rsa.PrivateKey
	template    *x509.Certificate
}

// Root generates a root certificate authority.
func Root(commonName string) (*Authority, error) {
	return authority(commonName, nil, nil)
}

// EncodeCertificate returns the PEM encoded certificate for the authority.
func (a *Authority) EncodeCertificate() (string, error) {
	return pems.EncodeCertificate(a.certificate)
}

// Branch creates a branch (i.e., intermediate certificate authority).
func (a *Authority) Branch(commonName string) (*Authority, error) {
	return authority(commonName, a.template, a.key)
}

// Leaf creates a leaf certificate.
func (a *Authority) Leaf(commonName string, alternativeNames []string, expires time.Duration) (*Leaf, error) {

	template := leafTemplate(commonName, alternativeNames, expires)

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.Wrapf(err, "error generating private key for [%s]", commonName)
	}

	certificate, err := x509.CreateCertificate(rand.Reader, template, a.template, &key.PublicKey, a.key)
	if err != nil {
		return nil, errors.Wrapf(err, "error signing certificate for [%s]", commonName)
	}

	return &Leaf{certificate: certificate, key: key}, nil
}

// authority creates a new authority.
func authority(commonName string, parentTemplate *x509.Certificate, parentKey *rsa.PrivateKey) (*Authority, error) {

	template := authorityTemplate(commonName)

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, errors.Wrapf(err, "error generating private key for [%s]", commonName)
	}

	if parentTemplate == nil {
		parentTemplate = template
	}

	if parentKey == nil {
		parentKey = key
	}

	certificate, err := x509.CreateCertificate(rand.Reader, template, parentTemplate, &key.PublicKey, parentKey)
	if err != nil {
		return nil, errors.Wrapf(err, "error signing certificate for [%s]", commonName)
	}

	return &Authority{certificate: certificate, key: key, template: template}, nil
}

// authorityTemplate returns the template for a root or branch certificate.
func authorityTemplate(commonName string) *x509.Certificate {
	return &x509.Certificate{
		BasicConstraintsValid: true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		NotAfter:              time.Now().AddDate(10, 0, 0),
		NotBefore:             time.Now(),
		SerialNumber:          big.NewInt(time.Now().Unix()),
		Subject:               subjectTemplate(commonName),
	}
}

// leafTemplate returns the template for a leaf certificate.
func leafTemplate(commonName string, alternativeNames []string, expires time.Duration) *x509.Certificate {
	return &x509.Certificate{
		BasicConstraintsValid: true,
		DNSNames:              alternativeNames,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		NotAfter:              time.Now().Add(expires),
		NotBefore:             time.Now(),
		SerialNumber:          big.NewInt(time.Now().Unix()),
		Subject:               subjectTemplate(commonName),
	}
}

// subjectTemplate returns the template for a subject.
func subjectTemplate(commonName string) pkix.Name {
	return pkix.Name{
		CommonName:         commonName,
		Country:            []string{"US"},
		Locality:           []string{"Alexandria"},
		Organization:       []string{"Decipher Technology Studios"},
		OrganizationalUnit: []string{"Engineering"},
		Province:           []string{"Virginia"},
		PostalCode:         []string{"22314"},
		StreetAddress:      []string{"110 S. Union St, Floor 2"},
	}
}
