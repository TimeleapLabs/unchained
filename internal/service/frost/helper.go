package frost

import (
	"github.com/ethereum/go-ethereum/common"
)

func FilterOnlineSigners(signers []common.Address, onlines []string) []common.Address {
	for i := 0; i < len(signers); i++ {
		isExist := false
		for _, online := range onlines {
			if online == signers[i].String() {
				isExist = true
				break
			}
		}

		if !isExist {
			signers = append(signers[:i], signers[i+1:]...)
		}
	}

	return signers
}
