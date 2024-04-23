package bls

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/btcsuite/btcutil/base58"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	bls12381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

type Signer struct {
	Name           string
	EvmAddress     string
	SecretKey      *big.Int
	PublicKey      *bls12381.G2Affine
	ShortPublicKey *bls12381.G1Affine

	g2Aff bls12381.G2Affine
	g1Aff bls12381.G1Affine
}

func NewIdentity() *Signer {
	b := &Signer{
		SecretKey: new(big.Int),
	}

	_, _, b.g1Aff, b.g2Aff = bls12381.Generators()

	if config.App.Secret.SecretKey != "" {
		decoded, err := hex.DecodeString(config.App.Secret.SecretKey)

		if err != nil {
			// TODO: Backwards compatibility with base58 encoded secret keys
			// Remove this after a few releases
			decoded = base58.Decode(config.App.Secret.SecretKey)
		}

		b.SecretKey.SetBytes(decoded)
		b.PublicKey = b.getPublicKey(b.SecretKey)
	} else {
		b.generateKeyPair()
	}

	b.ShortPublicKey = b.getShortPublicKey(b.SecretKey)

	return b
}

func (b *Signer) getPublicKey(sk *big.Int) *bls12381.G2Affine {
	return new(bls12381.G2Affine).ScalarMultiplication(&b.g2Aff, sk)
}

func (b *Signer) getShortPublicKey(sk *big.Int) *bls12381.G1Affine {
	return new(bls12381.G1Affine).ScalarMultiplication(&b.g1Aff, sk)
}

func (b *Signer) generateKeyPair() {
	// generate a random point in G2
	g2Order := bls12381_fr.Modulus()
	sk, err := rand.Int(rand.Reader, g2Order)

	if err != nil {
		panic(err)
	}

	pk := b.getPublicKey(sk)

	b.SecretKey = sk
	b.PublicKey = pk
}

func (b *Signer) Verify(
	signature bls12381.G1Affine,
	hashedMessage bls12381.G1Affine,
	publicKey bls12381.G2Affine) (bool, error) {
	pairingSigG2, err := bls12381.Pair(
		[]bls12381.G1Affine{signature},
		[]bls12381.G2Affine{b.g2Aff})
	if err != nil {
		return false, err
	}

	pairingHmPk, pairingError := bls12381.Pair(
		[]bls12381.G1Affine{hashedMessage},
		[]bls12381.G2Affine{publicKey})

	ok := pairingSigG2.Equal(&pairingHmPk)

	return ok, pairingError
}
