package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/KenshiTech/unchained/internal/address"
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/crypto/ethereum"
	"github.com/KenshiTech/unchained/internal/datasets"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

type MachineIdentity struct {
	Bls *bls.BlsSigner
	Eth *ethereum.EvmSigner
}

var Identity = &MachineIdentity{}

type Option func(identity *MachineIdentity) error

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

		address := ethCrypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		evmSignerInstance := &ethereum.EvmSigner{
			PublicKey:  publicKeyECDSA,
			PrivateKey: privateKey,
			Address:    address,
		}

		if privateKeyRegenerated || config.App.Secret.EvmAddress == "" {
			privateKeyBytes := ethCrypto.FromECDSA(evmSignerInstance.PrivateKey)

			config.App.Secret.EvmPrivateKey = hexutil.Encode(privateKeyBytes)[2:]
			config.App.Secret.EvmAddress = evmSignerInstance.Address
		}

		log.Logger.
			With("Address", evmSignerInstance.Address).
			Info("EVM identity initialized")

		return nil
	}
}

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
