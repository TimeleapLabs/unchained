package ethereum

import (
	"crypto/ecdsa"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type EvmSigner struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

var EvmSignerInstance *EvmSigner

func saveConfig() error {
	privateKeyBytes := crypto.FromECDSA(EvmSignerInstance.PrivateKey)

	config.App.Secret.EvmPrivateKey = hexutil.Encode(privateKeyBytes)[2:]
	config.App.Secret.EvmAddress = EvmSignerInstance.Address
	err := config.App.Secret.Save()

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Can't save EVM identity to the config")
	}

	return err
}

func InitClientIdentity() error {
	var privateKey *ecdsa.PrivateKey
	var err error
	var privateKeyRegenerated bool

	if config.App.Secret.EvmPrivateKey != "" {
		privateKey, err = crypto.HexToECDSA(config.App.Secret.EvmPrivateKey)

		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Can't decode EVM private key")

			return err
		}

		privateKeyRegenerated = true
	} else {
		privateKey, err = crypto.GenerateKey()

		if err != nil {
			log.Logger.
				With("Error", err).
				Error("Can't generate EVM private key")

			return err
		}
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Logger.Error("Can't assert type: publicKey is not of type *ecdsa.PublicKey")
		return err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	EvmSignerInstance = &EvmSigner{
		PublicKey:  publicKeyECDSA,
		PrivateKey: privateKey,
		Address:    address,
	}

	if privateKeyRegenerated || config.App.Secret.EvmAddress == "" {
		err := saveConfig()

		if err != nil {
			return err
		}
	}

	log.Logger.
		With("Address", EvmSignerInstance.Address).
		Info("EVM identity initialized")

	return nil
}
