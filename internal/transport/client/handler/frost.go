package handler

import (
	"context"
)

func (h *consumer) AskIfFrostSigner(ctx context.Context, message []byte) {
	// packet := new([]model.Signer).FromBytes(message)

}

func (w worker) AskIfFrostSigner(ctx context.Context, message []byte) {
	err := w.frostService.InitSigner()
	if err != nil {
		return
	}

	// packet := new(model.Signers).FromBytes(message)
	// TODO implement me
	panic("implement me")
}
