package frost

import (
	"errors"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bytemare/frost"
)

// MessageSigner represents state of signing of a message.
type MessageSigner struct {
	partySize   int
	data        []byte
	commitments frost.CommitmentList
	participant *frost.Participant
	commitment  *frost.Commitment
}

// Confirm function will set other parties confirms.
func (s *MessageSigner) Confirm(commitment *frost.Commitment) (*frost.SignatureShare, error) {
	s.commitments = append(s.commitments, commitment)

	if len(s.commitments.Participants()) < s.partySize {
		return nil, nil
	}

	signatureShare, err := s.participant.Sign(s.data, s.commitments)
	if err != nil {
		utils.Logger.With("err", err).Error("cant sign message")
		return nil, consts.ErrCantSign
	}

	if !s.participant.VerifySignatureShare(
		s.commitment,
		s.participant.PublicKey,
		signatureShare.SignatureShare,
		s.commitments,
		s.data,
	) {
		return nil, consts.ErrCantVerify
	}

	return signatureShare, nil
}

func (s *MessageSigner) AggregateSignatures(signatureShares []*frost.SignatureShare) ([]byte, error) {
	signature := s.participant.Aggregate(s.commitments, s.data, signatureShares)

	return signature.Encode(), nil
}

// NewSigner create a new signing state for a message.
func (s *DistributedSigner) NewSigner(data []byte) (*MessageSigner, *frost.Commitment, error) {
	if s.finalParticipant == nil {
		return nil, nil, consts.ErrSignerIsNotReady
	}

	commitment := s.finalParticipant.Commit()
	if commitment.Identifier.Equal(s.finalParticipant.KeyShare.Identifier) != 1 {
		return nil, nil, errors.New("identifier is not correct")
	}

	signer := &MessageSigner{
		partySize:   s.minSigningCount,
		data:        data,
		commitments: make(frost.CommitmentList, 0, s.signerCount),
		participant: s.finalParticipant,
		commitment:  commitment,
	}

	return signer, commitment, nil
}
