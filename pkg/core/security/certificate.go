package security

import (
	"crypto"
	"crypto/rand"
)

type Certificate [2048]byte

func GenerateCertificate() (Certificate, error) {
	var cert Certificate
	_, err := rand.Read(cert[:])
	if err != nil {
		return Certificate{}, err
	}
	return cert, nil
}

func (cert *Certificate) ID() string {
	return string(crypto.SHA256.New().Sum(cert[:]))
}
