package bls

import (
	"crypto/rand"
	"math/big"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	bls12_381_fr "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
)

var (
	g1Aff bls12381.G1Affine
)

func init() {
	_, _, g1Aff, _ = bls12381.Generators()
}

func GetPublicKey(sk *big.Int) *bls12381.G1Affine {
	return new(bls12381.G1Affine).ScalarMultiplication(&g1Aff, sk)
}

// generate BLS private and public key pair
func GenerateKeyPair() (*big.Int, *bls12381.G1Affine, error) {
	// generate a random point in G2
	g2Order := bls12_381_fr.Modulus()
	sk, err := rand.Int(rand.Reader, g2Order)

	if err != nil {
		return nil, nil, err
	}

	pk := GetPublicKey(sk)

	return sk, pk, nil
}

func Verify(
	signature bls12381.G2Affine,
	hashedMessage bls12381.G2Affine,
	publicKey bls12381.G1Affine) (bool, error) {

	pairingSigG1, _ := bls12381.Pair(
		[]bls12381.G1Affine{g1Aff},
		[]bls12381.G2Affine{signature})

	pairingHmPk, pairingError := bls12381.Pair(
		[]bls12381.G1Affine{publicKey},
		[]bls12381.G2Affine{hashedMessage})

	ok := pairingSigG1.Equal(&pairingHmPk)

	return ok, pairingError
}

func FastVerify(
	signature bls12381.G2Affine,
	invertedHashedMessage bls12381.G2Affine,
	publicKey bls12381.G1Affine) (bool, error) {

	ok, pairingError := bls12381.PairingCheck(
		[]bls12381.G1Affine{g1Aff, publicKey},
		[]bls12381.G2Affine{signature, invertedHashedMessage},
	)

	return ok, pairingError
}

func Hash(message []byte) (bls12381.G2Affine, error) {
	dst := []byte("UNCHAINED")
	return bls12381.HashToG2(message, dst)
}

func Sign(secretKey big.Int, message []byte) (bls12381.G2Affine, bls12381.G2Affine) {
	hashedMessage, _ := Hash(message)
	signature := new(bls12381.G2Affine).ScalarMultiplication(&hashedMessage, &secretKey)

	return *signature, hashedMessage
}

func RecoverSignature(bytes [96]byte) (bls12381.G2Affine, error) {
	signature := new(bls12381.G2Affine)
	_, err := signature.SetBytes(bytes[:])
	return *signature, err
}

func RecoverPublicKey(bytes [48]byte) (bls12381.G1Affine, error) {
	pk := new(bls12381.G1Affine)
	_, err := pk.SetBytes(bytes[:])
	return *pk, err
}

func AggregateSignatures(signatures []bls12381.G2Affine) (bls12381.G2Affine, error) {

	aggregated := new(bls12381.G2Jac).FromAffine(&signatures[0])

	for _, sig := range signatures[1:] {
		sigJac := new(bls12381.G2Jac).FromAffine(&sig)
		aggregated.AddAssign(sigJac)
	}

	aggregatedAffine := new(bls12381.G2Affine).FromJacobian(aggregated)

	return *aggregatedAffine, nil
}
