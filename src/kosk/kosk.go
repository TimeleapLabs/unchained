package kosk

// TODO: Move to "crypto"

import (
	"crypto/rand"

	"github.com/KenshiTech/unchained/bls"
)

type Challenge struct {
	Passed    bool
	Random    [128]byte
	Signature [48]byte
}

func NewChallenge() [128]byte {
	challenge := make([]byte, 128)
	rand.Read(challenge)
	return [128]byte(challenge)
}

// TODO: We should use small signatures

func VerifyChallenge(challenge [128]byte,
	publicKeyBytes [96]byte,
	signatureBytes [48]byte) (bool, error) {

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
