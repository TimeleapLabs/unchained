package bls

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/consts"

	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/btcsuite/btcutil/base58"
	bls12381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"

	"github.com/TimeleapLabs/unchained/internal/utils/address"

	"github.com/TimeleapLabs/unchained/internal/config"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
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

// Sign will sign a data with bls secret and return signed data.
func (s *Signer) Sign(data []byte) ([]byte, error) {
	hashedMessage, err := Hash(data)
	if err != nil {
		panic(err)
	}

	signature := new(bls12381.G1Affine).ScalarMultiplication(&hashedMessage, s.SecretKey).Bytes()

	return signature[:], err
}

// WriteConfigs writes the secret key, public key and address to the global config object.
func (s *Signer) WriteConfigs() {
	pkBytes := s.PublicKey.Bytes()
	config.App.Secret.SecretKey = hex.EncodeToString(s.SecretKey.Bytes())
	config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])
	config.App.Secret.Address = address.Calculate(pkBytes[:])
}

// Verify verifies the signature of a message belongs to the public key.
func (s *Signer) Verify(signature []byte, hashedMessage []byte, publicKey []byte) (bool, error) {
	signatureBls, err := RecoverSignature([48]byte(signature))
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't recover bls signature")
		return false, consts.ErrInternalError
	}

	publicKeyBls, err := RecoverPublicKey([96]byte(publicKey))
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't recover pub-key")
		return false, consts.ErrInternalError
	}

	messageBls, err := Hash(hashedMessage)
	if err != nil {
		utils.Logger.With("Err", err).Error("Can't convert message to hash")
		return false, consts.ErrInternalError
	}

	pairingSigG2, err := bls12381.Pair(
		[]bls12381.G1Affine{signatureBls},
		[]bls12381.G2Affine{s.g2Aff})
	if err != nil {
		return false, err
	}

	pairingHmPk, pairingError := bls12381.Pair(
		[]bls12381.G1Affine{messageBls},
		[]bls12381.G2Affine{publicKeyBls})

	ok := pairingSigG2.Equal(&pairingHmPk)

	return ok, pairingError
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

	utils.Logger.
		With("Address", s.ShortPublicKey.String()).
		Info("Unchained identity initialized")

	return s
}
