package clientidentity

import (
	"math/big"

	"github.com/KenshiTech/unchained/crypto/bls"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

//go:generate mockgen -destination=../../test/testutils/mocks/mock_client_identity.go -package=mocks -source=interface.go

// ClientIdentity
// TODO add proper comments
type ClientIdentity interface {
	GetSecretKey() *big.Int
	GetPublicKey() *bls12381.G2Affine
	GetShortPublicKey() *bls12381.G1Affine
	GetSigner() *bls.Signer
}

// Default is singleton pattern for the client identity
var Default ClientIdentity

func GetSecretKey() *big.Int {
	return Default.GetSecretKey()
}

func GetPublicKey() *bls12381.G2Affine {
	return Default.GetPublicKey()
}

func GetShortPublicKey() *bls12381.G1Affine {
	return Default.GetShortPublicKey()
}

func GetSigner() *bls.Signer {
	return Default.GetSigner()
}
