package security

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
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
	return hex.EncodeToString(crypto.SHA256.New().Sum(cert[:]))
}
