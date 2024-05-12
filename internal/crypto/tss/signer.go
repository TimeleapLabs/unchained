package tss

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/tss"
	"math/big"
)

// MessageSigner represents state of signing of a message.
type MessageSigner struct {
	party *signing.LocalParty
}

// NewSigning will sign a data with tss secret and return signed data.
func (s *DistributedSigner) NewSigning(signers []*DistributedSigner, data []byte, out chan tss.Message, end chan<- common.SignatureData) (*MessageSigner, error) {
	if !s.localPartySaved.ValidateWithProof() {
		return &MessageSigner{}, consts.ErrSignerIsNotReady
	}

	dataBig := new(big.Int).SetBytes(data)

	partyIDs := make(tss.UnSortedPartyIDs, len(signers))
	for i, signer := range signers {
		partyIDs[i] = signer.PartyID
	}

	sortedPartyIDs := tss.SortPartyIDs(partyIDs)

	ctx := tss.NewPeerContext(sortedPartyIDs)

	params := tss.NewParameters(tss.S256(), ctx, s.PartyID, s.numberOfPeers, s.minThreshold)
	signer := &MessageSigner{
		party: signing.NewLocalParty(dataBig, params, *s.localPartySaved, out, end).(*signing.LocalParty),
	}

	go func(signer *signing.LocalParty) {
		err := signer.Start()
		if err != nil {
			panic(err)
		}

		utils.Logger.With("ID", signer.PartyID().Moniker).Info("New Tss signer started")
	}(signer.party)

	return signer, nil
}

// Acknowledge function will confirm an acceptance from other signers
func (s *MessageSigner) Acknowledge(msg tss.Message) error {
	if s.party.PartyID() == msg.GetFrom() {
		return nil
	}

	bz, _, err := msg.WireBytes()
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}

	pMsg, err := tss.ParseWireMessage(bz, msg.GetFrom(), msg.IsBroadcast())
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}

	isOK, tssErr := s.party.Update(pMsg)
	if tssErr != nil {
		utils.Logger.Error(tssErr.Error())
		return err
	}

	if !isOK {
		return consts.ErrInternalError
	}

	return nil
}
