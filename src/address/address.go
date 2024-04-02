package address

import (
	"fmt"

	"github.com/KenshiTech/unchained/src/crypto/shake"
)

var chars = "0123456789ABCDEFGHJKMNPQRSTUVXYZ"

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

func Calculate(input []byte) string {
	hash := shake.Shake(input)
	address := ToBase32(hash[:20])
	checksum := shake.Shake([]byte(address))
	checkchars := []byte{chars[checksum[0]%32], chars[checksum[1]%32]}

	return fmt.Sprintf("%s%s", address, checkchars)
}

func CalculateHex(input []byte) (string, [20]byte) {
	hash := shake.Shake(input)
	addressBytes := hash[:20]
	return fmt.Sprintf("0x%x", addressBytes), [20]byte(addressBytes)
}
