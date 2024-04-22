package utils

import "golang.org/x/crypto/sha3"

func Shake(input []byte) []byte {
	shake := sha3.NewShake256()
	shake.Write(input)
	hash := shake.Sum(nil)
	return hash
}
