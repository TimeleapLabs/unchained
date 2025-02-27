package handler

import (
	"context"
)

// Handler is an interface that represent the handlers of client nodes.
type Handler interface {
	RPCRequest(ctx context.Context, message []byte)
	RPCResponse(ctx context.Context, message []byte)
}
