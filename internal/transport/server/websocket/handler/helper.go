package handler

import (
	"github.com/KenshiTech/unchained/internal/consts"
	"github.com/KenshiTech/unchained/internal/utils"
	"github.com/gorilla/websocket"
)

func Send(conn *websocket.Conn, messageType int, opCode consts.OpCode, payload []byte) {
	err := conn.WriteMessage(
		messageType,
		append(
			[]byte{byte(opCode)},
			payload...),
	)
	if err != nil {
		utils.Logger.With("Error", err).Error("Can't send packet")
	}
}

func SendMessage(conn *websocket.Conn, messageType int, opCode consts.OpCode, message string) {
	Send(conn, messageType, opCode, []byte(message))
}

func BroadcastListener(conn *websocket.Conn, ch <-chan []byte) {
	for message := range ch {
		err := conn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			utils.Logger.Error(err.Error())
		}
	}
}

func SendError(conn *websocket.Conn, messageType int, opCode consts.OpCode, err error) {
	SendMessage(conn, messageType, opCode, err.Error())
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
		utils.Logger.With("Error", err).Error("Can't close connection")
	}
}
