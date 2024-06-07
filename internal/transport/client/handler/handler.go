package handler

import (
	"context"
)

type Handler interface {
	Challenge(message []byte) []byte
	CorrectnessReport(ctx context.Context, message []byte)
	EventLog(ctx context.Context, message []byte)
	PriceReport(ctx context.Context, message []byte)

	ConfirmFrostHandshake(ctx context.Context, message []byte)
	StoreOnlineFrostParty(ctx context.Context, message []byte)
}
