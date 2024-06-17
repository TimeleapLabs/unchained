package frost

import (
	"context"
	"crypto/sha1" // #nosec
	"errors"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"

	"github.com/ethereum/go-ethereum/common"

	"github.com/TimeleapLabs/unchained/internal/service/pos"
)

type Service interface {
	SendOnlineSigners(ctx context.Context) error
	SyncSigners(ctx context.Context, onlineSigners []string) (<-chan *protocol.Message, error)
	CoordinateHandshake(ch <-chan *protocol.Message)
	ConfirmHandshake(ctx context.Context, message []byte) (bool, error)
	ConfirmHandshakeRaw(ctx context.Context, message *protocol.Message) (bool, error)

	RequestToSign(data []byte)
	SignData(ctx context.Context, data []byte) ([]byte, <-chan *protocol.Message, error)
	CoordinateSignData(hashOfMessage []byte, msgCh <-chan *protocol.Message)
	ConfirmSignedData(data []byte) ([]byte, error)
}

type service struct {
	pos pos.Service

	frost *multisig.DistributedSigner

	currentSigners  []common.Address
	signingMessages *utils.TTLMap
}

func (s *service) RequestToSign(data []byte) {
	conn.Send(consts.OpCodeFrostRequestSign, data)
}

func (s *service) SignData(ctx context.Context, data []byte) ([]byte, <-chan *protocol.Message, error) {
	signer, msgCh, err := s.frost.NewSigner(data)
	if err != nil {
		return nil, nil, err
	}

	hasher := sha1.New() // #nosec
	hasher.Write(data)
	hashOfMessage := hasher.Sum(nil)
	s.signingMessages.Set(ctx, string(hashOfMessage), signer)

	return hashOfMessage, msgCh, nil
}

func (s *service) CoordinateSignData(hashOfMessage []byte, msgCh <-chan *protocol.Message) {
	for msg := range msgCh {
		msgBytes, err := msg.MarshalBinary()
		if err != nil {
			utils.Logger.With("Error", err).Error("Can't marshal message")
			return
		}

		conn.Send(consts.OpCodeFrostConfirmMessageSign, append(hashOfMessage, msgBytes...))
	}
}

func (s *service) ConfirmSignedData(data []byte) ([]byte, error) {
	hashOfMessage := data[:20]
	signatureBytes := data[20:]
	signer, isExist := s.signingMessages.Get(string(hashOfMessage))
	if !isExist {
		utils.Logger.With("Error", "").Error("Can't get signer")
		return nil, consts.ErrSignerIsNotReady
	}

	signature, err := signer.(*multisig.MessageSigner).ConfirmFromBytes(signatureBytes)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't confirm signed message")
		return nil, err
	}

	return signature, nil
}

func (s *service) ConfirmHandshakeRaw(_ context.Context, message *protocol.Message) (bool, error) {
	isReady, err := s.frost.Confirm(message)
	if err != nil && !errors.Is(err, consts.ErrInvalidSignature) {
		utils.Logger.With("Error", err).Error("Can't confirm handshake message")
		return false, err
	}

	return isReady, nil
}

func (s *service) ConfirmHandshake(_ context.Context, message []byte) (bool, error) {
	isReady, err := s.frost.ConfirmFromBytes(message)
	if err != nil && errors.Is(err, consts.ErrInvalidSignature) {
		utils.Logger.With("Error", err).Error("Can't confirm handshake message")
		return false, err
	}

	return isReady, nil
}

func New(pos pos.Service) Service {
	return &service{
		pos:             pos,
		signingMessages: utils.NewDataStore(),
	}
}
