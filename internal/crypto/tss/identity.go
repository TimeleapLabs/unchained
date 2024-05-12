package tss

import (
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

// DistributedSigner represents a TSS identity.
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

// Verify function will check message signature to be valid
func (s *DistributedSigner) Verify(_ []byte, _ []byte, _ []byte) (bool, error) {
	// TODO implement me
	panic("implement me")
}

// WriteConfigs will write identity keys to a persistent storage
func (s *DistributedSigner) WriteConfigs() {
	// nothing to do!
}

// Update function will update the identity key about other parties
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
