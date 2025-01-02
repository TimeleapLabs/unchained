package runtime

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// RunWebSocketCall runs a function with the given name and parameters.
func RunWebSocketCall(_ context.Context, conn *websocket.Conn, params *dto.RPCRequest) error {
	err := conn.WriteMessage(websocket.BinaryMessage, params.Sia().Bytes())
	if err != nil {
		utils.Logger.With("err", err).Error("Error sending message")
		return err
	}

	return nil
}
