package lib

import (
	"crypto/rand"
	"math/big"

	bls12_381_ecc "github.com/consensys/gnark-crypto/ecc/bls12-381"
	bls12_381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

var (
	g2Gen bls12_381_ecc.G2Affine
)

func init() {
	_, _, _, g2Gen = bls12_381_ecc.Generators()
}

// generate BLS private and public key pair
func GenerateKeyPair() (*big.Int, *bls12_381_ecc.G2Affine, error) {
	// generate a random point in G2
	g2Order := bls12_381_fr.Modulus()
	sk, err := rand.Int(rand.Reader, g2Order)
	if err != nil {
		return nil, nil, err
	}

	pk := new(bls12_381_ecc.G2Affine).ScalarMultiplication(&g2Gen, sk)

	return sk, pk, nil
}

func Verify(
	signature bls12_381_ecc.G1Affine,
	g2Gen bls12_381_ecc.G2Affine,
	hashedMessage bls12_381_ecc.G1Affine,
	publicKey bls12_381_ecc.G2Affine) (bool, error) {

	pairingSigG2, _ := bls12_381_ecc.Pair(
		[]bls12_381_ecc.G1Affine{signature},
		[]bls12_381_ecc.G2Affine{g2Gen})

	pairingHmPk, pairingError := bls12_381_ecc.Pair(
		[]bls12_381_ecc.G1Affine{hashedMessage},
		[]bls12_381_ecc.G2Affine{publicKey})

	ok := pairingSigG2.Equal(&pairingHmPk)

	return ok, pairingError
}

func FastVerify(
	signature bls12_381_ecc.G1Affine,
	g2Gen bls12_381_ecc.G2Affine,
	invertedHashedMessage bls12_381_ecc.G1Affine,
	publicKey bls12_381_ecc.G2Affine) (bool, error) {

	ok, pairingError := bls12_381_ecc.PairingCheck(
		[]bls12_381_ecc.G1Affine{signature, invertedHashedMessage},
		[]bls12_381_ecc.G2Affine{g2Gen, publicKey})

	return ok, pairingError
}
