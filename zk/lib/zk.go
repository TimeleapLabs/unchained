package lib

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_bls12381"
)

// Boneh-Lynn-Shacham (BLS) signature verification
// e(sig, g2) * e(hm, pk) == 1
// where:
//   - Sig (in G1) the signature
//   - G2 (in G2) the public generator of G2
//   - Hm (in G1) the hashed-to-curve message
//   - Pk (in G2) the public key of the signer
type Circuit struct {
	Sig sw_bls12381.G1Affine
	G2  sw_bls12381.G2Affine
	Hm  sw_bls12381.G1Affine
	Pk  sw_bls12381.G2Affine
}

func (circuit *Circuit) Define(api frontend.API) error {

	pairing, _ := sw_bls12381.NewPairing(api)

	err := sw_bls12381.Pairing.PairingCheck(*pairing,
		[]*sw_bls12381.G1Affine{&circuit.Sig, &circuit.Hm},
		[]*sw_bls12381.G2Affine{&circuit.G2, &circuit.Pk})

	if err != nil {
		return err
	}

	return nil
}
