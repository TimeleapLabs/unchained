package bls

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/btcsuite/btcutil/base58"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	bls12381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

// Signer represents a BLS identity.
type Signer struct {
	Name           string
	SecretKey      *big.Int
	PublicKey      *bls12381.G2Affine
	ShortPublicKey *bls12381.G1Affine

	g2Aff bls12381.G2Affine
	g1Aff bls12381.G1Affine
}

// WriteConfigs writes the secret key, public key and address to the global config object.
func (s *Signer) WriteConfigs() {
	pkBytes := s.PublicKey.Bytes()
	config.App.Secret.SecretKey = hex.EncodeToString(s.SecretKey.Bytes())
	config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])
	config.App.Secret.Address = address.Calculate(pkBytes[:])
}

// NewIdentity creates a new BLS identity.
func NewIdentity() *Signer {
	s := &Signer{
		SecretKey: new(big.Int),
	}

	if config.App.Secret.SecretKey == "" && !config.App.System.AllowGenerateSecrets {
		panic("BLS secret key is not provided and secrets generation is not allowed")
	}

	_, _, s.g1Aff, s.g2Aff = bls12381.Generators()

	if config.App.Secret.SecretKey != "" {
		decoded, err := hex.DecodeString(config.App.Secret.SecretKey)
		if err != nil {
			// TODO: Backwards compatibility with base58 encoded secret keys
			// Remove this after a few releases
			decoded = base58.Decode(config.App.Secret.SecretKey)
		}

		s.SecretKey.SetBytes(decoded)
		s.PublicKey = new(bls12381.G2Affine).ScalarMultiplication(&s.g2Aff, s.SecretKey)
	} else {
		// generate a random point in G2
		g2Order := bls12381_fr.Modulus()
		sk, err := rand.Int(rand.Reader, g2Order)
		if err != nil {
			panic(err)
		}

		pk := new(bls12381.G2Affine).ScalarMultiplication(&s.g2Aff, sk)

		s.SecretKey = sk
		s.PublicKey = pk
	}

	s.ShortPublicKey = new(bls12381.G1Affine).ScalarMultiplication(&s.g1Aff, s.SecretKey)

	return s
}

// Verify verifies the signature of a message belongs to the public key.
func (s *Signer) Verify(
	signature bls12381.G1Affine, hashedMessage bls12381.G1Affine, publicKey bls12381.G2Affine,
) (bool, error) {
	pairingSigG2, err := bls12381.Pair(
		[]bls12381.G1Affine{signature},
		[]bls12381.G2Affine{s.g2Aff})
	if err != nil {
		return false, err
	}

	pairingHmPk, pairingError := bls12381.Pair(
		[]bls12381.G1Affine{hashedMessage},
		[]bls12381.G2Affine{publicKey})

	ok := pairingSigG2.Equal(&pairingHmPk)

	return ok, pairingError
}
