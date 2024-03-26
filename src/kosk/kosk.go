package kosk

// TODO: Move to "crypto"

import (
	"crypto/rand"

	"github.com/KenshiTech/unchained/crypto/bls"
)

const (
	LenOfChallenge = 128
	LenOfSignature = 48
	LenOfPublicKey = 96
)

func NewChallenge() [LenOfChallenge]byte {
	challenge := make([]byte, LenOfChallenge)
	_, err := rand.Read(challenge)
	if err != nil {
		panic(err)
	}

	return [LenOfChallenge]byte(challenge)
}

// TODO: We should use small signatures

func VerifyChallenge(challenge [LenOfChallenge]byte,
	publicKeyBytes [LenOfPublicKey]byte,
	signatureBytes [LenOfSignature]byte) (bool, error) {
	signature, err := bls.RecoverSignature(signatureBytes)
	if err != nil {
		return false, err
	}

	hash, err := bls.Hash(challenge[:])
	if err != nil {
		return false, err
	}

	publicKey, err := bls.RecoverPublicKey(publicKeyBytes)
	if err != nil {
		return false, err
	}

	return bls.Verify(signature, hash, publicKey)
}
