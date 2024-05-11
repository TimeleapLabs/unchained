package tss

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/tss"
	"math/big"
)

// Signer represents a TSS identity.
type Signer struct {
	ctx             *tss.PeerContext
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

func (s *Signer) Verify(signature []byte, hashedMessage []byte, publicKey []byte) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Signer) WriteConfigs() {
	//TODO implement me
	panic("implement me")
}

// NewSigning will sign a data with tss secret and return signed data.
func (s *Signer) NewSigning(data []byte, out chan tss.Message, end chan<- common.SignatureData) (*MessageSigner, error) {
	if !s.localPartySaved.ValidateWithProof() {
		return &MessageSigner{}, consts.ErrSignerIsNotReady
	}

	dataBig := new(big.Int).SetBytes(data)

	params := tss.NewParameters(tss.S256(), s.ctx, s.PartyID, s.numberOfPeers, s.minThreshold)
	signer := &MessageSigner{
		party: signing.NewLocalParty(dataBig, params, *s.localPartySaved, out, end).(*signing.LocalParty),
	}

	go func(signer *MessageSigner) {
		err := signer.party.Start()
		if err != nil {
			panic(err)
		}
	}(signer)

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

	isOK, err := s.party.Update(pMsg)
	if err != nil && !isOK {
		utils.Logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Signer) Update(msg tss.Message) error {
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

	isOK, err := s.keyParty.Update(pMsg)
	if err != nil && !isOK {
		utils.Logger.Error(err.Error())
		return err
	}

	return nil
}

// NewIdentity creates a new BLS identity.
func NewIdentity(signerID int, signers []string, outCh chan tss.Message, done chan struct{}, minThreshold int) *Signer {
	partyIDs := make(tss.UnSortedPartyIDs, len(signers))
	for i, signer := range signers {
		party := tss.NewPartyID(signer, signer, big.NewInt(int64(i+1)))
		partyIDs[i] = party
	}

	signer := &Signer{
		minThreshold:  minThreshold,
		numberOfPeers: len(signers),
		//rawPartyID:    signers[signerID],
		PartyID: partyIDs[signerID],
		result:  make(chan keygen.LocalPartySaveData, len(signers)),
	}

	sortedPartyIDs := tss.SortPartyIDs(partyIDs)

	signer.ctx = tss.NewPeerContext(sortedPartyIDs)

	signer.keyParty = keygen.NewLocalParty(
		tss.NewParameters(tss.S256(), signer.ctx, sortedPartyIDs[signerID], len(signers), minThreshold),
		outCh,
		signer.result,
	).(*keygen.LocalParty)

	go func(signer *Signer) {
		err := signer.keyParty.Start()
		if err != nil {
			panic(err)
		}
		utils.Logger.With("ID", signer.PartyID.Moniker).Info("New Tss party started")

		for {
			select {
			case save := <-signer.result:
				signer.localPartySaved = &save
				done <- struct{}{}
				break
			}
		}

	}(signer)

	return signer
}
