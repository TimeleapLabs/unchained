package multisig

import (
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"golang.org/x/crypto/sha3"
)

func NewSignersFromStrings(signers []string) []party.ID {
	participantIds := []party.ID{}
	for _, signer := range signers {
		participantIds = append(participantIds, party.ID(signer))
	}

	return participantIds
}

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}
