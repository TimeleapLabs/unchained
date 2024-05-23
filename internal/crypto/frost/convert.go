package frost

import (
	"encoding/hex"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/bytemare/crypto"
	"github.com/bytemare/frost"
	"github.com/bytemare/frost/dkg"
)

// DecodeCommitment will decode commitment from bytes.
func (s *MessageSigner) DecodeCommitment(data []byte) (*frost.Commitment, error) {
	scalarLen := s.participant.Ciphersuite.Group.ScalarLength()
	commitment := &frost.Commitment{
		Identifier:   s.participant.Ciphersuite.Group.NewScalar(),
		HidingNonce:  s.participant.Ciphersuite.Group.NewElement(),
		BindingNonce: s.participant.Ciphersuite.Group.NewElement(),
	}

	err := commitment.Identifier.Decode(data[:scalarLen])
	if err != nil {
		utils.Logger.With("err", err, "value", hex.EncodeToString(data[:scalarLen])).Error("cant decode commitment identifier")
		return nil, err
	}

	err = commitment.HidingNonce.Decode(data[scalarLen : scalarLen*2])
	if err != nil {
		utils.Logger.With("err", err, "value", hex.EncodeToString(data[:scalarLen])).Error("cant decode commitment hiding_nonce")
		// return nil, err
	}

	err = commitment.BindingNonce.Decode(data[scalarLen*2 : scalarLen*3])
	if err != nil {
		utils.Logger.With("err", err, "value", hex.EncodeToString(data[:scalarLen])).Error("cant decode commitment binding_nonce")
		return nil, err
	}

	return commitment, nil
}

// DecodeSignature will decode signature from bytes.
func (s *MessageSigner) DecodeSignature(data []byte) (*frost.Signature, error) {
	signature := frost.Signature{
		R: s.participant.Ciphersuite.Group.NewElement(),
		Z: s.participant.Ciphersuite.Group.NewScalar(),
	}
	err := signature.Decode(s.participant.Ciphersuite.Group, data)
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode signature")
		return &signature, consts.ErrCantDecode
	}

	return &signature, nil
}

func EncodeRoundOneMessage(message *dkg.Round1Data) []byte {
	commitment := []byte{}
	for _, commit := range message.Commitment {
		commitment = append(commitment, commit.Encode()...)
	}

	out := []byte{}
	out = append(out, message.SenderIdentifier.Encode()...)
	out = append(out, message.ProofOfKnowledge.Encode()...)
	out = append(out, commitment...)

	return out
}

func EncodeRoundTwoMessages(message []*dkg.Round2Data) [][]byte {
	out := [][]byte{}

	for _, msg := range message {
		out = append(out, EncodeRoundTwoMessage(msg))
	}

	return out
}

func EncodeRoundTwoMessage(message *dkg.Round2Data) []byte {
	out := []byte{}

	out = append(out, message.SenderIdentifier.Encode()...)
	out = append(out, message.ReceiverIdentifier.Encode()...)
	out = append(out, message.SecretShare.Encode()...)

	return out
}

func (s *DistributedSigner) DecodeRoundTwoMessage(message []byte) (*dkg.Round2Data, error) {
	r2Data := &dkg.Round2Data{
		SenderIdentifier:   s.config.Ciphersuite.Group.NewScalar(),
		ReceiverIdentifier: s.config.Ciphersuite.Group.NewScalar(),
		SecretShare:        s.config.Ciphersuite.Group.NewScalar(),
	}

	err := r2Data.SenderIdentifier.Decode(message[:s.config.Ciphersuite.Group.ScalarLength()])
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode sender_identifier")
		return nil, err
	}

	err = r2Data.ReceiverIdentifier.Decode(message[s.config.Ciphersuite.Group.ScalarLength() : s.config.Ciphersuite.Group.ScalarLength()*2])
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode receiver_identifier")
		return nil, err
	}

	err = r2Data.SecretShare.Decode(message[s.config.Ciphersuite.Group.ScalarLength()*2:])
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode secret_share")
		return nil, err
	}

	return r2Data, nil
}

func (s *DistributedSigner) DecodeRoundOneMessage(message []byte) (*dkg.Round1Data, error) {
	senderBytes := message[:s.config.Ciphersuite.Group.ScalarLength()]
	proofBytes := message[s.config.Ciphersuite.Group.ScalarLength() : s.config.Ciphersuite.Group.ScalarLength()+64]
	commitmentsBytes := message[s.config.Ciphersuite.Group.ScalarLength()+64:]

	r1Data := &dkg.Round1Data{
		ProofOfKnowledge: frost.Signature{
			R: s.config.Ciphersuite.Group.NewElement(),
			Z: s.config.Ciphersuite.Group.NewScalar(),
		},
		SenderIdentifier: s.config.Ciphersuite.Group.NewScalar(),
		Commitment:       []*crypto.Element{},
	}

	for i := 0; i < len(commitmentsBytes); i += s.config.Ciphersuite.Group.ElementLength() {
		commitment := s.config.Ciphersuite.Group.NewElement()
		err := commitment.Decode(commitmentsBytes[i : i+s.config.Ciphersuite.Group.ElementLength()])
		if err != nil {
			utils.Logger.With("err", err).Error("cant decode commitment")
			return nil, err
		}

		r1Data.Commitment = append(r1Data.Commitment, commitment)
	}

	err := r1Data.ProofOfKnowledge.Decode(s.config.Ciphersuite.Group, proofBytes)
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode proof_of_knowledge")
		return nil, err
	}

	err = r1Data.SenderIdentifier.Decode(senderBytes)
	if err != nil {
		utils.Logger.With("err", err).Error("cant decode identifier")
		return nil, err
	}

	return r1Data, nil
}
