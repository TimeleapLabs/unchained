package address

import (
	"fmt"

	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/btcsuite/btcutil/base58"
)

func Calculate(input []byte) string {
	hash := utils.Shake(input)
	address := base58.Encode(hash[:20])
	return address
}

func CalculateHex(input []byte) (string, [20]byte) {
	hash := utils.Shake(input)
	addressBytes := hash[:20]
	return fmt.Sprintf("0x%x", addressBytes), [20]byte(addressBytes)
}
