package uniswap

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
)

func IsNewSigner(signature correctness.Signature, records []model.AssetPrice) bool {
	for _, record := range records {
		for _, signer := range record.Signers {
			if signature.Signer.PublicKey == signer.PublicKey {
				return false
			}
		}
	}

	return true
}
