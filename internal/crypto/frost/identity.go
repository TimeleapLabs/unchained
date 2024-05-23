package frost

import (
	"errors"

	"github.com/bytemare/crypto"
	"github.com/bytemare/frost"
	"github.com/bytemare/frost/dkg"
)

// DistributedSigner represents a Frost identity.
type DistributedSigner struct {
	ID                  *crypto.Scalar
	currentParticipant  *dkg.Participant
	finalParticipant    *frost.Participant
	accumulatedMessages []*dkg.Round1Data
	ackMessages         []*dkg.Round2Data
	config              *frost.Configuration
	signerCount         int
	minSigningCount     int
}

// Update function will update the identity key about other parties.
func (s *DistributedSigner) Update(msg []byte) ([][]byte, error) {
	round1Data, err := s.DecodeRoundOneMessage(msg)
	if err != nil {
		return nil, err
	}

	s.accumulatedMessages = append(s.accumulatedMessages, round1Data)

	if len(s.accumulatedMessages) != s.signerCount {
		return nil, nil
	}

	round2Data, err := s.currentParticipant.Continue(s.accumulatedMessages)
	if err != nil {
		return nil, err
	}

	if len(round2Data) != len(s.accumulatedMessages)-1 {
		return nil, errors.New("number of accept messages is not correct")
	}

	return EncodeRoundTwoMessages(round2Data), nil
}

// Finalize function will confirm the identity key about other parties updates.
func (s *DistributedSigner) Finalize(msgByte []byte) error {
	msg, err := s.DecodeRoundTwoMessage(msgByte)
	if err != nil {
		return err
	}

	if msg.ReceiverIdentifier.Equal(s.currentParticipant.Identifier) == 0 {
		return nil
	}

	s.ackMessages = append(s.ackMessages, msg)

	if len(s.ackMessages) < len(s.accumulatedMessages)-1 {
		return nil
	}

	participantsSecretKey, _, groupPublicKeyGeneratedInDKG, err := s.currentParticipant.Finalize(
		s.accumulatedMessages,
		s.ackMessages,
	)
	if err != nil {
		return err
	}

	s.config.GroupPublicKey = groupPublicKeyGeneratedInDKG

	s.finalParticipant = s.config.Participant(s.ID, participantsSecretKey)

	return nil
}

// NewIdentity creates a new Frost identity.
func NewIdentity(id int, signerCount int, minSigningCount int) ([]byte, *DistributedSigner) {
	signer := DistributedSigner{
		accumulatedMessages: make([]*dkg.Round1Data, 0, signerCount),
		config:              frost.Ristretto255.Configuration(),
		signerCount:         signerCount,
		minSigningCount:     minSigningCount,
	}

	signer.ID = signer.config.IDFromInt(id)
	signer.currentParticipant = dkg.NewParticipant(
		signer.config.Ciphersuite,
		signer.ID,
		signerCount,
		minSigningCount,
	)

	round1Data := signer.currentParticipant.Init()
	if round1Data.SenderIdentifier.Equal(signer.ID) != 1 {
		panic("this is just a test, and it failed")
	}

	return EncodeRoundOneMessage(round1Data), &signer
}
