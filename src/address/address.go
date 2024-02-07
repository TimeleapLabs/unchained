package address

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

var chars = "0123456789ABCDEFGHJKMNPQTSTUVXYZ"

func ToBase32(input []byte) string {
	var output []byte
	var temp int
	var bits int

	for _, b := range input {
		temp = (temp << 8) | int(b)
		bits += 8

		for bits >= 5 {
			bits -= 5
			index := (temp >> bits) & 0x1F
			output = append(output, chars[index])
		}
	}

	if bits > 0 {
		lastChunk := (temp << (5 - bits)) & 0x1F
		output = append(output, chars[lastChunk])
	}

	return string(output)
}

func Shake(input []byte) []byte {
	shake := sha3.NewShake256()
	shake.Write(input)
	hash := shake.Sum(nil)
	return hash
}

func Calculate(input []byte) string {
	hash := Shake(input)
	address := ToBase32(hash[:20])
	checksum := Shake(input[:20])
	checkchars := []byte{chars[checksum[0]%32], chars[checksum[1]%32]}

	return fmt.Sprintf("%s%s", address, checkchars)
}
