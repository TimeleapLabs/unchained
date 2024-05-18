package bls

import (
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

func Hash(message []byte) (bls12381.G1Affine, error) {
	dst := []byte("UNCHAINED")
	return bls12381.HashToG1(message, dst)
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

func AggregateSignatures(signatures [][]byte) ([]byte, error) {
	signaturesBls := []bls12381.G1Affine{}

	for _, signature := range signatures {
		signatureBls, err := RecoverSignature([48]byte(signature))
		if err != nil {
			return nil, err
		}

		signaturesBls = append(signaturesBls, signatureBls)
	}

	aggregated := new(bls12381.G1Jac).FromAffine(&signaturesBls[0])

	for _, sig := range signaturesBls[1:] {
		sigJac := new(bls12381.G1Jac).FromAffine(&sig)
		aggregated.AddAssign(sigJac)
	}

	aggregatedAffine := new(bls12381.G1Affine).FromJacobian(aggregated)
	aggregatedAffineByte := aggregatedAffine.Bytes()

	return aggregatedAffineByte[:], nil
}
