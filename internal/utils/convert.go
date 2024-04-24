package utils

import (
	"math/big"

	"golang.org/x/crypto/sha3"
)

func BigIntToFloat(power *big.Int) *big.Float {
	decimalPlaces := big.NewInt(1e18)
	powerFloat := new(big.Float).SetInt(power)
	powerFloat.Quo(powerFloat, new(big.Float).SetInt(decimalPlaces))
	return powerFloat
}

func Shake(input []byte) []byte {
	shake := sha3.NewShake256()
	shake.Write(input)
	hash := shake.Sum(nil)
	return hash
}
