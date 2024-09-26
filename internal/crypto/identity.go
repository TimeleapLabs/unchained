package crypto

import (
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// MachineIdentity holds machine identity and provide and manage keys.
type MachineIdentity struct {
	Bls *bls.Signer
	Eth *ethereum.Signer
}

// Identity is a global variable that holds machine identity.
var Identity = &MachineIdentity{}

// Option represents a function that can add new identity to machine identity.
type Option func(identity *MachineIdentity) error

// InitMachineIdentity loads all provided identities and save them to secret file.
func InitMachineIdentity(options ...Option) {
	for _, option := range options {
		err := option(Identity)
		if err != nil {
			panic(err)
		}
	}

	if config.App.System.AllowGenerateSecrets {
		err := config.App.Secret.Save()
		if err != nil {
			panic(err)
		}
	}
}

// ExportEvmSigner returns EVM signer from machine identity.
func (i *MachineIdentity) ExportEvmSigner() *model.Signer {
	publicKey := Identity.Bls.PublicKey.Bytes()
	shortPublicKey := Identity.Bls.ShortPublicKey.Bytes()

	return &model.Signer{
		Name:           config.App.System.Name,
		EvmAddress:     Identity.Eth.Address,
		PublicKey:      hex.EncodeToString(publicKey[:]),
		ShortPublicKey: hex.EncodeToString(shortPublicKey[:]),
	}
}

// WithEvmSigner initialize and will add Evm keys to machine identity.
func WithEvmSigner() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Eth = ethereum.NewIdentity()
		machineIdentity.Eth.WriteConfigs()

		utils.Logger.
			With("Address", machineIdentity.Eth.Address).
			Info("EVM identity initialized")

		return nil
	}
}

// WithBlsIdentity initialize and will add Bls keys to machine identity.
func WithBlsIdentity() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Bls = bls.NewIdentity()
		machineIdentity.Bls.WriteConfigs()

		utils.Logger.
			With("Address", machineIdentity.Bls.ShortPublicKey.String()).
			Info("Unchained identity initialized")

		return nil
	}
}
