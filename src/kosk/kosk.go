package kosk

// TODO: Move to "crypto"

import (
	"crypto/rand"

	"github.com/KenshiTech/unchained/src/crypto/bls"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

const (
	LenOfChallenge = 128
	LenOfSignature = 48
	LenOfPublicKey = 96
)

type Challenge struct {
	Passed    bool
	Random    [LenOfChallenge]byte
	Signature [LenOfSignature]byte
}

func (c *Challenge) Sia() *sia.Sia {
	return new(sia.Sia).
		AddBool(c.Passed).
		AddByteArray8(c.Random[:]).
		AddByteArray8(c.Signature[:])
}

func (c *Challenge) DeSia(sia *sia.Sia) *Challenge {
	c.Passed = sia.ReadBool()
	copy(c.Random[:], sia.ReadByteArray8())
	copy(c.Signature[:], sia.ReadByteArray8())

	return c
}

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
