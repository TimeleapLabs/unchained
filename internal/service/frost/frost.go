package frost

import (
	"context"
	"errors"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto/multisig"
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
	SignData(ctx context.Context, data []byte) ([]byte, error)
}

type service struct {
	pos pos.Service

	frost *multisig.DistributedSigner

	currentSigners []common.Address
}

func (s *service) SignData(ctx context.Context, data []byte) ([]byte, error) {
	signer, msgCh, err := s.frost.NewSigner(data)
	if err != nil {

	}
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
		pos: pos,
	}
}
