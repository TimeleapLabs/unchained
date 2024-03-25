package bls

import (
	"encoding/hex"
	"math/big"

	"github.com/KenshiTech/unchained/address"
	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/log"

	"github.com/btcsuite/btcutil/base58"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

var ClientSecretKey *big.Int
var ClientPublicKey *bls12381.G2Affine
var ClientShortPublicKey *bls12381.G1Affine
var ClientSigner Signer

func saveConfig() {
	pkBytes := ClientPublicKey.Bytes()

	config.Secrets.Set("secretKey", hex.EncodeToString(ClientSecretKey.Bytes()))
	config.Secrets.Set("publicKey", hex.EncodeToString(pkBytes[:]))

	err := config.Secrets.WriteConfig()

	if err != nil {
		panic(err)
	}
}

func InitClientIdentity() {
	var err error

	if config.Secrets.IsSet("secretKey") {
		secretKeyFromConfig := config.Secrets.GetString("secretKey")
		decoded, err := hex.DecodeString(secretKeyFromConfig)

		if err != nil {
			// TODO: Backwards compatibility with base58 encoded secret keys
			// Remove this after a few releases
			decoded = base58.Decode(secretKeyFromConfig)
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

	ClientSigner = Signer{
		Name:           config.Config.GetString("name"),
		EvmWallet:      config.Secrets.GetString("evmwallet"),
		PublicKey:      ClientPublicKey.Bytes(),
		ShortPublicKey: ClientShortPublicKey.Bytes(),
	}

	log.Logger.
		With("Address", addrStr).
		Info("Unchained")

	// TODO: Avoid recalculating this
	config.Secrets.Set("publicKey", hex.EncodeToString(pkBytes[:]))

	if !config.Secrets.IsSet("address") {
		config.Secrets.Set("address", addrStr)
		err := config.Secrets.WriteConfig()

		if err != nil {
			panic(err)
		}
	}
}
