package runtime

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/queue"
)

// RunWebSocketCall runs a function with the given name and parameters.
func RunWebSocketCall(_ context.Context, wsQueue *queue.WebSocketWriter, params *dto.RPCRequest) error {
	wsQueue.SendRaw(params.Sia().Bytes()) // TODO: How to handle write errors?
	return nil
}
