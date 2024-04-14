package bls

import (
	"math/big"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

func FastVerify(
	signature bls12381.G1Affine,
	g2Gen bls12381.G2Affine,
	invertedHashedMessage bls12381.G1Affine,
	publicKey bls12381.G2Affine) (bool, error) {
	ok, pairingError := bls12381.PairingCheck(
		[]bls12381.G1Affine{signature, invertedHashedMessage},
		[]bls12381.G2Affine{g2Gen, publicKey})

	return ok, pairingError
}

func Hash(message []byte) (bls12381.G1Affine, error) {
	dst := []byte("UNCHAINED")
	return bls12381.HashToG1(message, dst)
}

func Sign(secretKey big.Int, message []byte) (bls12381.G1Affine, bls12381.G1Affine) {
	hashedMessage, err := Hash(message)
	if err != nil {
		panic(err)
	}
	signature := new(bls12381.G1Affine).ScalarMultiplication(&hashedMessage, &secretKey)

	return *signature, hashedMessage
}

func RecoverSignature(bytes [48]byte) (bls12381.G1Affine, error) {
	signature := new(bls12381.G1Affine)
	_, err := signature.SetBytes(bytes[:])
	return *signature, err
}

func RecoverPublicKey(bytes [96]byte) (bls12381.G2Affine, error) {
	pk := new(bls12381.G2Affine)
	_, err := pk.SetBytes(bytes[:])
	return *pk, err
}

func AggregateSignatures(signatures []bls12381.G1Affine) (bls12381.G1Affine, error) {
	aggregated := new(bls12381.G1Jac).FromAffine(&signatures[0])

	for _, sig := range signatures[1:] {
		sigJac := new(bls12381.G1Jac).FromAffine(&sig)
		aggregated.AddAssign(sigJac)
	}

	aggregatedAffine := new(bls12381.G1Affine).FromJacobian(aggregated)

	return *aggregatedAffine, nil
}
