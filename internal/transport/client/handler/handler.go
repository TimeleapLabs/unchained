package handler

import (
	"context"
)

type Handler interface {
	Challenge(message []byte) []byte
	CorrectnessReport(ctx context.Context, message []byte)
	RPCRequest(ctx context.Context, message []byte)
}
