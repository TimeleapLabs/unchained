package crypto

import (
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto/bls"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/model"
)

// Signer represents a Signing method.
type Signer interface {
	Sign(data []byte) ([]byte, error)
	Verify(signature []byte, hashedMessage []byte, publicKey []byte) (bool, error)
	WriteConfigs()
}

// MachineIdentity holds machine identity and provide and manage keys.
type MachineIdentity struct {
	Bls   Signer
	Eth   Signer
	Frost *multisig.DistributedSigner
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
	blsPublicKey, err := hex.DecodeString(config.App.Secret.PublicKey)
	if err != nil {
		panic(err)
	}

	return &model.Signer{
		Name:           config.App.System.Name,
		EvmAddress:     config.App.Secret.EvmAddress,
		PublicKey:      [96]byte(blsPublicKey),
		ShortPublicKey: config.App.Secret.ShortPublicKey,
	}
}

// WithEvmSigner initialize and will add Evm identity to machine identity.
func WithEvmSigner() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Eth = ethereum.NewIdentity()
		machineIdentity.Eth.WriteConfigs()

		return nil
	}
}

// WithTssSigner initialize and will add Tss identity to machine identity.
// func WithTssSigner(signers []string, minThreshold int) func(machineIdentity *MachineIdentity) error {
//	return func(machineIdentity *MachineIdentity) error {
//		machineIdentity.Tss = tss.NewIdentity(
//			signers,
//			minThreshold,
//		)
//		//machineIdentity.Tss.WriteConfigs()
//
//		return nil
//	}
//}

// WithBlsIdentity initialize and will add Bls identity to machine identity.
func WithBlsIdentity() func(machineIdentity *MachineIdentity) error {
	return func(machineIdentity *MachineIdentity) error {
		machineIdentity.Bls = bls.NewIdentity()
		machineIdentity.Bls.WriteConfigs()

		return nil
	}
}
