package tss

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/tss"
)

// Signer represents a TSS identity.
type DistributedSigner struct {
	ctx             *tss.PeerContext
	partyIDs        tss.UnSortedPartyIDs
	PartyID         *tss.PartyID
	keyParty        *keygen.LocalParty
	localPartySaved *keygen.LocalPartySaveData
	result          chan keygen.LocalPartySaveData
	minThreshold    int
	numberOfPeers   int
}

type MessageSigner struct {
	party *signing.LocalParty
}

func (s *DistributedSigner) Verify(_ []byte, _ []byte, _ []byte) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func (s *DistributedSigner) WriteConfigs() {
	// TODO implement me
	panic("implement me")
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

func (s *MessageSigner) AckSignature(msg tss.Message) error {
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

func (s *DistributedSigner) Update(msg tss.Message) error {
	if s.keyParty.PartyID() == msg.GetFrom() {
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

	isOK, tssErr := s.keyParty.Update(pMsg)
	if tssErr != nil {
		utils.Logger.Error(tssErr.Error())
		return err
	}

	if !isOK {
		return consts.ErrInternalError
	}

	return nil
}

// NewIdentity creates a new BLS identity.
func NewIdentity(signerID int, signers []string, outCh chan tss.Message, done chan struct{}, minThreshold int) *DistributedSigner {
	partyIDs := make(tss.UnSortedPartyIDs, len(signers))
	for i, signer := range signers {
		party := tss.NewPartyID(signer, signer, big.NewInt(int64(i+1)))
		partyIDs[i] = party
	}

	signer := &DistributedSigner{
		minThreshold:  minThreshold,
		numberOfPeers: len(signers),
		//rawPartyID:    signers[signerID],
		partyIDs: partyIDs,
		PartyID:  partyIDs[signerID],
		result:   make(chan keygen.LocalPartySaveData, len(signers)),
	}

	sortedPartyIDs := tss.SortPartyIDs(partyIDs)

	signer.ctx = tss.NewPeerContext(sortedPartyIDs)

	var isOk bool
	signer.keyParty, isOk = keygen.NewLocalParty(
		tss.NewParameters(tss.S256(), signer.ctx, sortedPartyIDs[signerID], len(signers), minThreshold),
		outCh,
		signer.result,
	).(*keygen.LocalParty)

	if !isOk {
		panic(consts.ErrInternalError)
	}

	go func(signer *DistributedSigner) {
		err := signer.keyParty.Start()
		if err != nil {
			panic(err)
		}
		utils.Logger.With("ID", signer.PartyID.Moniker).Info("New Tss party started")

		for save := range signer.result {
			localPartySaved := save
			signer.localPartySaved = &localPartySaved
			done <- struct{}{}
		}
	}(signer)

	return signer
}
