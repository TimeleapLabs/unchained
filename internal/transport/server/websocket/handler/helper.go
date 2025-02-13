package handler

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/transport/server/pubsub"

	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/gorilla/websocket"
)

func BroadcastManager(connCtx context.Context, subCtx context.Context, topic string, sub pubsub.Subscriber) {
	for {
		select {
		case <-connCtx.Done():
			pubsub.Unsubscribe(topic, sub.Writer)
			return
		case <-subCtx.Done():
			return
		case message := <-sub.Channel:
			sub.Writer.SendRaw(message)
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
