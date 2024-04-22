package uniswap

import (
	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/model"
)

func IsNewSigner(signature model.Signature, records []*ent.AssetPrice) bool {
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}
