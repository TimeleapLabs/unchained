package ed25519

import (
	"crypto/rand"

	"github.com/TimeleapLabs/unchained/internal/utils/address"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/btcsuite/btcutil/base58"

	"crypto/ed25519"
)

// Signer represents a Ed25519 identity.
type Signer struct {
	Name      string
	SecretKey ed25519.PrivateKey
	PublicKey ed25519.PublicKey
}

// WriteConfigs writes the secret key, public key and address to the global config object.
func (s *Signer) WriteConfigs() {
	config.App.Secret.SecretKey = base58.Encode(s.SecretKey)
	config.App.Secret.PublicKey = base58.Encode(s.PublicKey)
	config.App.Secret.Address = address.Calculate(s.PublicKey)
}

// NewIdentity creates a new Ed25519 identity.
func NewIdentity() *Signer {
	s := &Signer{}

	if config.App.Secret.SecretKey == "" && !config.App.System.AllowGenerateSecrets {
		panic("Ed25519 secret key is not provided and secrets generation is not allowed")
	}

	if config.App.Secret.SecretKey != "" {
		decoded := base58.Decode(config.App.Secret.SecretKey)

		s.SecretKey = ed25519.PrivateKey(decoded)
		pubKey, ok := s.SecretKey.Public().(ed25519.PublicKey)

		if !ok {
			panic("failed to assert type ed25519.PublicKey")
		}

		s.PublicKey = pubKey
	} else {
		pk, sk, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}

		s.SecretKey = sk
		s.PublicKey = pk
	}

	return s
}

// Sign signs a message with the identities secret key.
func (s *Signer) Sign(message []byte) ([]byte, error) {
	return s.SecretKey.Sign(rand.Reader, message, &ed25519.Options{})
}

// Verify verifies the signature of a message belongs to the public key.
func (s *Signer) Verify(signature []byte, message []byte, publicKey ed25519.PublicKey) bool {
	return ed25519.Verify(publicKey, message, signature)
}
