package bls

import (
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

func InitClientIdentity() {
	var err error
	var pkBytes [96]byte

	if config.Secrets.IsSet("secretKey") {

		decoded := base58.Decode(config.Secrets.GetString("secretKey"))

		ClientSecretKey = new(big.Int)
		ClientSecretKey.SetBytes(decoded)

		ClientPublicKey = GetPublicKey(ClientSecretKey)
		pkBytes = ClientPublicKey.Bytes()

	} else {
		ClientSecretKey, ClientPublicKey, err = GenerateKeyPair()

		if err != nil {
			panic(err)
		}

		pkBytes = ClientPublicKey.Bytes()

		config.Secrets.Set("secretKey", base58.Encode(ClientSecretKey.Bytes()))
		config.Secrets.Set("publicKey", base58.Encode(pkBytes[:]))

		err = config.Secrets.WriteConfig()

		if err != nil {
			panic(err)
		}
	}

	ClientShortPublicKey = GetShortPublicKey(ClientSecretKey)
	addrStr := address.Calculate(pkBytes[:])

	ClientSigner = Signer{
		Name:           config.Config.GetString("name"),
		EvmWallet:      config.Config.GetString("evmWallet"),
		PublicKey:      ClientPublicKey.Bytes(),
		ShortPublicKey: ClientShortPublicKey.Bytes(),
	}

	log.Logger.
		With("Address", addrStr).
		Info("Unchained")

	// TODO: Avoid recalculating this
	config.Secrets.Set("publicKey", base58.Encode(pkBytes[:]))

	if !config.Secrets.IsSet("address") {
		config.Secrets.Set("address", addrStr)
		err := config.Secrets.WriteConfig()

		if err != nil {
			panic(err)
		}
	}
}
