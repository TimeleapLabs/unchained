package utils

import (
	"crypto/rand"
)

const (
	LenOfChallenge = 128
)

// NewChallenge generate a new challenge.
func NewChallenge() [LenOfChallenge]byte {
	challenge := make([]byte, LenOfChallenge)
	_, err := rand.Read(challenge)
	if err != nil {
		panic(err)
	}

	return [LenOfChallenge]byte(challenge)
}
