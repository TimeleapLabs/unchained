package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

// Send sends a packet to the client.
func Send(conn *websocket.Conn, opCode consts.OpCode, payload []byte) {
	err := conn.WriteMessage(
		websocket.BinaryMessage,
		append(
			[]byte{byte(opCode)},
			payload...),
	)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't send packet")
	}
}

// SendMessage sends a string packet to the client.
func SendMessage(conn *websocket.Conn, opCode consts.OpCode, message string) {
	Send(conn, opCode, []byte(message))
}

// BroadcastListener listens for messages on the channel and sends them to the client.
func BroadcastListener(ctx context.Context, conn *websocket.Conn, ch chan []byte) {
	for {
		select {
		case <-ctx.Done():
			utils.Logger.Info("Closing connection")
			close(ch)
			return
		case message := <-ch:
			err := conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				utils.Logger.Error(err.Error())
			}
		}
	}
}

// SendError sends an error message to the client.
func SendError(conn *websocket.Conn, opCode consts.OpCode, err error) {
	SendMessage(conn, opCode, err.Error())
}

// Close gracefully closes the connection.
func Close(conn *websocket.Conn) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		utils.Logger.With("Error", err).Error("Connection closed")
	}

	err = conn.Close()
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't close connection")
	}
}
