// Package security
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

type PairToken [128]byte

func GeneratePairToken() (PairToken, error) {
	var token PairToken
	_, err := rand.Read(token[:])
	if err != nil {
		return PairToken{}, err
	}
	return token, nil
}