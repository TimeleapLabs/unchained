package handler

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/transport/server/pubsub"

	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/gorilla/websocket"
)

func BroadcastListener(ctx context.Context, conn *websocket.Conn, topic string, ch chan []byte) {
	for {
		select {
		case <-ctx.Done():
			utils.Logger.Info("Closing connection")
			pubsub.Unsubscribe(topic, ch)
			return
		case message := <-ch:
			err := conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				utils.Logger.Error(err.Error())
			}
		}
	}
}

func Close(conn *websocket.Conn) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		utils.Logger.With("Error", err).Error("Connection closed")
	}

	err = conn.Close()
	if err != nil {
		utils.Logger.With("Error", err).Error("Cannot close connection")
	}
}
