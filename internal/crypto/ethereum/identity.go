package ethereum

import (
	"crypto/ecdsa"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

// Signer represents an Ethereum identity.
type Signer struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

// WriteConfigs writes the secret key, public key and address to the global config object.
func (s *Signer) WriteConfigs() {
	privateKeyBytes := ethCrypto.FromECDSA(s.PrivateKey)
	config.App.Secret.EvmPrivateKey = hexutil.Encode(privateKeyBytes)[2:]
	config.App.Secret.EvmAddress = s.Address
}

// NewIdentity creates a new Ethereum identity.
func NewIdentity() *Signer {
	var privateKey *ecdsa.PrivateKey
	var err error

	if config.App.Secret.EvmPrivateKey == "" && !config.App.System.AllowGenerateSecrets {
		panic("EVM private key is not provided and secrets generation is not allowed")
	}

	if config.App.Secret.EvmPrivateKey != "" {
		privateKey, err = ethCrypto.HexToECDSA(config.App.Secret.EvmPrivateKey)
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Cannot decode EVM private key")

			panic(err)
		}
	} else {
		privateKey, err = ethCrypto.GenerateKey()
		if err != nil {
			utils.Logger.
				With("Error", err).
				Error("Cannot generate EVM private key")

			panic(err)
		}
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	s := &Signer{
		PublicKey:  publicKeyECDSA,
		PrivateKey: privateKey,
		Address:    ethCrypto.PubkeyToAddress(*publicKeyECDSA).Hex(),
	}

	return s
}
