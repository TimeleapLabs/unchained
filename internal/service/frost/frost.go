package frost

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type Service interface {
	SyncSigners(ctx context.Context) error
	ConfirmHandshake(ctx context.Context, message []byte) error
}

type service struct {
	pos pos.Service

	currentSigners []common.Address
}

func (s *service) ConfirmHandshake(_ context.Context, message []byte) error {
	isReady, err := crypto.Identity.Frost.ConfirmFromBytes(message)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't confirm handshake message")
		return err
	}

	if isReady {
		conn.SendMessage(consts.OpcodeFrostSignerIsReady, crypto.Identity.ExportEvmSigner().EvmAddress)
	}

	return nil
}

func New(pos pos.Service) Service {
	return &service{
		pos: pos,
	}
}
