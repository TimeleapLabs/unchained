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
func (s *MessageSigner) Confirm(commitmentBytes []byte) ([]byte, error) {
	commitment, err := s.DecodeCommitment(commitmentBytes)
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode commitment")
		return nil, consts.ErrCantDecode
	}

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

	return signatureShare.Encode(), nil
}

// Aggregate function will generate final signature by combining parties signature shares.
func (s *MessageSigner) Aggregate(signatureSharesBytes [][]byte) ([]byte, error) {
	signatureShares := []*frost.SignatureShare{}

	for _, signatureShareBytes := range signatureSharesBytes {
		signatureShare, err := frost.Ristretto255.Configuration().DecodeSignatureShare(signatureShareBytes)
		if err != nil {
			utils.Logger.With("err", err).Error("cant decode signature share")
			return nil, consts.ErrCantDecode
		}

		signatureShares = append(signatureShares, signatureShare)
	}

	signature := s.participant.Aggregate(s.commitments, s.data, signatureShares)

	return signature.Encode(), nil
}

// Verify function will verify the signature.
func (s *MessageSigner) Verify(signatureBytes []byte) (bool, error) {
	signature, err := s.DecodeSignature(signatureBytes)
	if err != nil {
		return false, err
	}

	return frost.Verify(
		frost.Ristretto255.Configuration().Ciphersuite,
		s.data,
		signature,
		s.participant.GroupPublicKey,
	), nil
}

// NewSigner create a new signing state for a message.
func (s *DistributedSigner) NewSigner(data []byte) (*MessageSigner, []byte, error) {
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

	return signer, commitment.Encode(), nil
}
