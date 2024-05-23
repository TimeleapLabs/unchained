package handler

import (
	"context"
)

func (h *consumer) InitFrostSigner(ctx context.Context, message []byte) {
	// packet := new([]model.Signer).FromBytes(message)

}

func (w worker) InitFrostSigner(ctx context.Context, message []byte) {
	// packet := new(model.Signers).FromBytes(message)
	// TODO implement me
	panic("implement me")
}
