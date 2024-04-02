package bls

import (
	"encoding/hex"
	"math/big"

	"github.com/KenshiTech/unchained/src/config"

	"github.com/KenshiTech/unchained/src/address"
	"github.com/KenshiTech/unchained/src/datasets"
	"github.com/KenshiTech/unchained/src/log"

	"github.com/btcsuite/btcutil/base58"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

var ClientSecretKey *big.Int
var ClientPublicKey *bls12381.G2Affine
var ClientShortPublicKey *bls12381.G1Affine
var ClientSigner datasets.Signer

func saveConfig() {
	pkBytes := ClientPublicKey.Bytes()

	config.App.Secret.SecretKey = hex.EncodeToString(ClientSecretKey.Bytes())
	config.App.Secret.PublicKey = hex.EncodeToString(pkBytes[:])
	err := config.App.Secret.Save()

	if err != nil {
		panic(err)
	}
}

func InitClientIdentity() {
	var err error

	if config.App.Secret.SecretKey != "" {
		decoded, err := hex.DecodeString(config.App.Secret.SecretKey)

		if err != nil {
			// TODO: Backwards compatibility with base58 encoded secret keys
			// Remove this after a few releases
			decoded = base58.Decode(config.App.Secret.SecretKey)
		}

		ClientSecretKey = new(big.Int)
		ClientSecretKey.SetBytes(decoded)
		ClientPublicKey = GetPublicKey(ClientSecretKey)

		if err != nil {
			saveConfig()
		}
	} else {
		ClientSecretKey, ClientPublicKey, err = GenerateKeyPair()

		if err != nil {
			panic(err)
		}

		saveConfig()
	}

	ClientShortPublicKey = GetShortPublicKey(ClientSecretKey)

	pkBytes := ClientPublicKey.Bytes()
	addrStr := address.Calculate(pkBytes[:])

	ClientSigner = datasets.Signer{
		Name:           config.App.System.Name,
		EvmWallet:      config.App.Secret.EvmWallet,
		PublicKey:      ClientPublicKey.Bytes(),
		ShortPublicKey: ClientShortPublicKey.Bytes(),
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
