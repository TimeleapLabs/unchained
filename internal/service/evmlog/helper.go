package evmlog

import (
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/model"
)

func isNewSigner(signature model.Signature, records []*ent.EventLog) bool {
	for _, record := range records {
		for _, signer := range record.Edges.Signers {
			if signature.Signer.PublicKey == [96]byte(signer.Key) {
				return false
			}
		}
	}

	return true
}

func sortEventArgs(lhs model.EventLogArg, rhs model.EventLogArg) int {
	if lhs.Name < rhs.Name {
		return -1
	}
	return 1
}
