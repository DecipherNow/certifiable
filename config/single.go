package config

import "time"

// Single represents the configuration for a single certificate.
type Single struct {

	// AlternativeName defines the subject alternative names for the certificate.
	AlternativeNames []string `json:"alternativeNames" mapstructure:"alternativeNames" yaml:"alternativeNames"`

	// CommonName defines the subject common name for the certificate.
	CommonName string `json:"commonName" mapstructure:"commonName" yaml:"commonName"`

	// Expires defines the duration for which the certificate is valid.
	Expires time.Duration `json:"expires" mapstructure:"expires" yaml:"expires"`
}
