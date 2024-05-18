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
}

func (s *MessageSigner) Confirm(commitment *frost.Commitment) error {
	s.commitments = append(s.commitments, commitment)

	if len(s.commitments.Participants()) < s.partySize {
		return nil
	}

	signatureShare, err := s.participant.Sign(s.data, s.commitments)
	if err != nil {
		utils.Logger.With("err", err).Error("cant sign message")
		return consts.ErrCantSign
	}

	if !s.participant.VerifySignatureShare(
		commitment,
		s.participant.GroupPublicKey,
		signatureShare.SignatureShare,
		s.commitments,
		s.data,
	) {
		return consts.ErrCantVerify
	}

	return nil
}

func (s *DistributedSigner) NewSigner(data []byte) (*MessageSigner, *frost.Commitment, error) {
	if s.finalParticipant == nil {
		return nil, nil, consts.ErrSignerIsNotReady
	}

	signer := &MessageSigner{
		partySize:   s.signerCount,
		data:        data,
		commitments: make(frost.CommitmentList, 0, s.signerCount),
		participant: s.finalParticipant,
	}

	commitment := s.finalParticipant.Commit()
	if commitment.Identifier.Equal(s.finalParticipant.KeyShare.Identifier) != 1 {
		return nil, nil, errors.New("identifier is not correct")
	}

	return signer, commitment, nil
}
