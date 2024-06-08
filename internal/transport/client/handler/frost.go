package handler

import (
	"context"
)

func (h *postgresConsumer) ConfirmFrostHandshake(_ context.Context, _ []byte) {}
func (h *schnorrConsumer) ConfirmFrostHandshake(_ context.Context, _ []byte)  {}

func (w worker) ConfirmFrostHandshake(ctx context.Context, message []byte) {
	err := w.frostService.ConfirmHandshake(ctx, message)
	if err != nil {
		return
	}
}

func (h *postgresConsumer) StoreOnlineFrostParty(_ context.Context, _ []byte) {}

func (h *schnorrConsumer) StoreOnlineFrostParty(_ context.Context, evmAddressBytes []byte) {
	h.signerRepository.SetSignerIsAlive(string(evmAddressBytes))
}

func (w worker) StoreOnlineFrostParty(_ context.Context, _ []byte) {}
