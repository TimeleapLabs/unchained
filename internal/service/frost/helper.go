package frost

import (
	"github.com/ethereum/go-ethereum/common"
)

func FilterOnlineSigners(signers []common.Address, onlines []string) []common.Address {
	lookup := make(map[string]bool)
	for _, item := range onlines {
		lookup[item] = true
	}

	var filtered []common.Address
	for _, item := range signers {
		if _, exists := lookup[item.String()]; exists {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
