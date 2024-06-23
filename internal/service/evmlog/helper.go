package evmlog

import (
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/service/correctness"
)

func isNewSigner(signature correctness.Signature, records []model.EventLog) bool {
	for _, record := range records {
		for _, signer := range record.Signers {
			if signature.Signer.PublicKey == signer.PublicKey {
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
