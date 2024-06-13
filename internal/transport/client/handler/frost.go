package handler

import (
	"context"
	"encoding/json"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

func (h *consumer) ConfirmFrostHandshake(_ context.Context, _ []byte) {}

func (w *worker) ConfirmFrostHandshake(ctx context.Context, message []byte) {
	isReady, err := w.frostService.ConfirmHandshake(ctx, message)
	if err != nil {
		return
	}

	if isReady {
		conn.SendMessage(consts.OpcodeFrostSignerIsReady, crypto.Identity.ExportEvmSigner().EvmAddress)
	}
}

func (h *consumer) InitFrostIdentity(ctx context.Context, message []byte) {}
func (w *worker) InitFrostIdentity(ctx context.Context, message []byte) {
	utils.Logger.Info("Start init frost identity")
	onlineSigners := []string{}

	err := json.Unmarshal(message, &onlineSigners)
	if err != nil {
		return
	}

	handshakeChannel, err := w.frostService.SyncSigners(ctx, onlineSigners)
	if err != nil {
		return
	}

	w.frostService.CoordinateHandshake(handshakeChannel)
}
