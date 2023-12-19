package lib

import (
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	bls12_381_ecc "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/emulated/sw_bls12381"

	"github.com/urfave/cli/v2"
)

func CompileCommand(cCtx *cli.Context) error {
	// compiles our circuit into a R1CS
	var circuit Circuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	// groth16 zkSNARK: Setup
	pk, vk, _ := groth16.Setup(ccs)

	// Generate Solidity file
	f, _ := os.Create("contract_g16.sol")
	exportErr := vk.ExportSolidity(f)

	if exportErr != nil {
		fmt.Println(exportErr)
		return exportErr
	}

	f, _ = os.Create("ccs.bin")
	_, exportErr = ccs.WriteTo(f)

	if exportErr != nil {
		fmt.Println(exportErr)
		return exportErr
	}

	f, _ = os.Create("pk.bin")
	_, exportErr = pk.WriteTo(f)

	if exportErr != nil {
		fmt.Println(exportErr)
		return exportErr
	}

	f, _ = os.Create("vk.bin")
	_, exportErr = vk.WriteTo(f)

	if exportErr != nil {
		fmt.Println(exportErr)
		return exportErr
	}

	return nil
}

func TestCommand(cCtx *cli.Context) error {
	secretKey, publicKey, _ := GenerateKeyPair()
	message := []byte("Test message")
	dst := []byte("UNCHAINED")

	hashedMessage, _ := bls12_381_ecc.HashToG1(message, dst)
	signature := new(bls12_381_ecc.G1Affine).ScalarMultiplication(&hashedMessage, secretKey)

	ok, pairingError := Verify(*signature, g2Gen, hashedMessage, *publicKey)

	if ok {
		fmt.Println("✅ First test passes")
	} else {
		fmt.Println("❌ First test fails")
		return pairingError
	}

	// Alternate verify why no work?

	var invertedHm bls12_381_ecc.G1Affine
	invertedHm.Neg(&hashedMessage) // Inverting hm for the pairing

	ok, pairingError = FastVerify(*signature, g2Gen, invertedHm, *publicKey)

	// This one doesn't pass!
	if ok {
		fmt.Println("✅ Second test passes")
	} else {
		fmt.Println("❌ Second test fails")
		return pairingError
	}

	// compiles our circuit into a R1CS
	ccs := groth16.NewCS(ecc.BN254)
	f, openErr := os.Open("ccs.bin")

	if openErr != nil {
		fmt.Println("❌ Cannot open the circuit")
		return openErr
	}

	_, readErr := ccs.ReadFrom(f)

	if readErr != nil {
		fmt.Println("❌ Cannot read the circuit")
		return readErr
	}

	// groth16 zkSNARK: Setup
	pk := groth16.NewProvingKey(ecc.BN254)
	vk := groth16.NewVerifyingKey(ecc.BN254)

	f, openErr = os.Open("pk.bin")

	if openErr != nil {
		fmt.Println("❌ Cannot open the proving key")
		return openErr
	}

	_, readErr = pk.ReadFrom(f)

	if readErr != nil {
		fmt.Println("❌ Cannot read the proving key")
		return readErr
	}

	f, openErr = os.Open("vk.bin")

	if openErr != nil {
		fmt.Println("❌ Cannot open the verifying key")
		return openErr
	}

	_, readErr = vk.ReadFrom(f)

	if readErr != nil {
		fmt.Println("Cannot read the verifying key")
		return readErr
	}

	// witness definition
	assignment := Circuit{
		Sig: sw_bls12381.NewG1Affine(*signature),
		G2:  sw_bls12381.NewG2Affine(g2Gen),
		Hm:  sw_bls12381.NewG1Affine(invertedHm),
		Pk:  sw_bls12381.NewG2Affine(*publicKey)}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	proof, _ := groth16.Prove(ccs, pk, witness)
	err := groth16.Verify(proof, vk, publicWitness)

	if err != nil {
		return err
	}

	return nil
}
