package security

import "crypto/rand"

type PairToken [128]byte

func GeneratePairToken() (PairToken, error) {
	var token PairToken
	_, err := rand.Read(token[:])
	if err != nil {
		return PairToken{}, err
	}
	return token, nil
}
