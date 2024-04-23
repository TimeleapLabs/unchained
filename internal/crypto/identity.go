package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/address"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/datasets"
	"github.com/TimeleapLabs/unchained/internal/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

// MachineIdentity holds machine identity and provide and manage keys.
type MachineIdentity struct {
	Bls *bls.Signer
	Eth *ethereum.EvmSigner
}

var Identity = &MachineIdentity{}

type Option func(identity *MachineIdentity) error

// InitMachineIdentity loads all provided identities and save them to secret file.
func InitMachineIdentity(options ...Option) {
	for _, option := range options {
		err := option(Identity)
		if err != nil {
			panic(err)
		}
	}

	err := config.App.Secret.Save()
	if err != nil {
		panic(err)
	}
}

func (i *MachineIdentity) ExportBlsSigner() *datasets.Signer {
	return &datasets.Signer{
		Name:           config.App.System.Name,
		EvmAddress:     Identity.Eth.Address,
		PublicKey:      Identity.Bls.PublicKey.Bytes(),
		ShortPublicKey: Identity.Bls.ShortPublicKey.Bytes(),
	}
}

// WithEvmSigner initialize and will add Evm keys to machine identity.
func WithEvmSigner() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		var privateKey *ecdsa.PrivateKey
		var err error
		var privateKeyRegenerated bool

		if config.App.Secret.EvmPrivateKey != "" {
			privateKey, err = ethCrypto.HexToECDSA(config.App.Secret.EvmPrivateKey)

			if err != nil {
				log.Logger.
					With("Error", err).
					Error("Can't decode EVM private key")

				return err
			}

			privateKeyRegenerated = true
		} else {
			privateKey, err = ethCrypto.GenerateKey()

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

		ethAddress := ethCrypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		machineIdentity.Eth = &ethereum.EvmSigner{
			PublicKey:  publicKeyECDSA,
			PrivateKey: privateKey,
			Address:    ethAddress,
		}

		if privateKeyRegenerated || config.App.Secret.EvmAddress == "" {
			privateKeyBytes := ethCrypto.FromECDSA(machineIdentity.Eth.PrivateKey)

			config.App.Secret.EvmPrivateKey = hexutil.Encode(privateKeyBytes)[2:]
			config.App.Secret.EvmAddress = machineIdentity.Eth.Address
		}

		log.Logger.
			With("Address", machineIdentity.Eth.Address).
			Info("EVM identity initialized")

		return nil
	}
}

// WithBlsIdentity initialize and will add Bls keys to machine identity.
func WithBlsIdentity() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Bls = bls.NewIdentity()
		pkBytes := machineIdentity.Bls.PublicKey.Bytes()

		config.App.Secret.SecretKey = hex.EncodeToString(machineIdentity.Bls.SecretKey.Bytes())
		config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])

		addrStr := address.Calculate(pkBytes[:])

		log.Logger.
			With("Address", addrStr).
			Info("Unchained identity initialized")

		// TODO: Avoid recalculating this
		config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])

		if config.App.Secret.Address != "" {
			config.App.Secret.Address = addrStr
		}

		return nil
	}
}
