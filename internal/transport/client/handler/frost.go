package handler

import (
	"context"
)

func (h *consumer) ConfirmFrostHandshake(_ context.Context, _ []byte) {}

func (w worker) ConfirmFrostHandshake(ctx context.Context, message []byte) {
	err := w.frostService.ConfirmHandshake(ctx, message)
	if err != nil {
		return
	}
}

func (h *consumer) StoreOnlineFrostParty(_ context.Context, _ []byte) {

}

func (w worker) StoreOnlineFrostParty(_ context.Context, _ []byte) {}
