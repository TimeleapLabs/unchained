package ethereum

import (
	"crypto/ecdsa"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Signer represents an Ethereum identity.
type Signer struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

func (s *Signer) Verify(_ []byte, _ []byte, _ []byte) (bool, error) {
	// TODO implement me
	panic("implement me")
}

// WriteConfigs writes the secret key, public key and address to the global config object.
func (s *Signer) WriteConfigs() {
	privateKeyBytes := crypto.FromECDSA(s.PrivateKey)
	config.App.Secret.EvmPrivateKey = hexutil.Encode(privateKeyBytes)[2:]
	config.App.Secret.EvmAddress = s.Address
}

func (s *Signer) Sign(data []byte) ([]byte, error) {
	signature, err := crypto.Sign(data, s.PrivateKey)
	if err != nil {
		return nil, err
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}

// NewIdentity creates a new Ethereum identity.
func NewIdentity() *Signer {
	var privateKey *ecdsa.PrivateKey
	var err error

	if config.App.Secret.EvmPrivateKey == "" && !config.App.System.AllowGenerateSecrets {
		panic("EVM private key is not provided and secrets generation is not allowed")
	}

	if config.App.Secret.EvmPrivateKey != "" {
		privateKey, err = crypto.HexToECDSA(config.App.Secret.EvmPrivateKey)
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Can't decode EVM private key")

			panic(err)
		}
	} else {
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Can't generate EVM private key")

			panic(err)
		}
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("Can't assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	s := &Signer{
		PublicKey:  publicKeyECDSA,
		PrivateKey: privateKey,
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA).Hex(),
	}

	utils.Logger.
		With("Address", s.Address).
		Info("EVM identity initialized")

	return s
}
