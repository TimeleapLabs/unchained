package crypto

import (
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/ed25519"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
)

// MachineIdentity holds machine identity and provide and manage keys.
type MachineIdentity struct {
	Ed25519 *ed25519.Signer
	Eth     *ethereum.Signer
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
	return &model.Signer{
		Name:       config.App.System.Name,
		EvmAddress: Identity.Eth.Address,
		PublicKey:  i.Ed25519.PublicKey,
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

// WithEd25519Identity initialize and will add Ed25519 keys to machine identity.
func WithEd25519Identity() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Ed25519 = ed25519.NewIdentity()
		machineIdentity.Ed25519.WriteConfigs()

		utils.Logger.
			With("Address", address.Calculate(machineIdentity.Ed25519.PublicKey)).
			Info("Unchained identity initialized")

		return nil
	}
}
