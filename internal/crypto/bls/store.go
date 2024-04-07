package bls

import (
	"encoding/hex"

	"github.com/KenshiTech/unchained/internal/config"

	"github.com/KenshiTech/unchained/internal/address"
	"github.com/KenshiTech/unchained/internal/datasets"
	"github.com/KenshiTech/unchained/internal/log"
)

var ClientSigner datasets.Signer
var MachineIdentity *Identity

func InitClientIdentity() {
	MachineIdentity = NewIdentity()
	MachineIdentity.Save()

	pkBytes := MachineIdentity.PublicKey.Bytes()
	addrStr := address.Calculate(pkBytes[:])

	ClientSigner = datasets.Signer{
		Name:           config.App.System.Name,
		EvmWallet:      config.App.Secret.EvmWallet,
		PublicKey:      MachineIdentity.PublicKey.Bytes(),
		ShortPublicKey: MachineIdentity.ShortPublicKey.Bytes(),
	}

	log.Logger.
		With("Address", addrStr).
		Info("Unchained")

	// TODO: Avoid recalculating this
	config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])

	if config.App.Secret.Address != "" {
		config.App.Secret.Address = addrStr
		err := config.App.Secret.Save()

		if err != nil {
			panic(err)
		}
	}
}
