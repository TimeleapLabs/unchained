package net

import (
	"github.com/KenshiTech/unchained/src/constants/opcodes"
	"github.com/KenshiTech/unchained/src/log"
	"github.com/KenshiTech/unchained/src/net/consumer"
	"github.com/gorilla/websocket"
)

func Send(conn *websocket.Conn, messageType int, opCode opcodes.OpCode, payload []byte) {
	err := conn.WriteMessage(
		messageType,
		append(
			[]byte{byte(opCode)},
			payload...),
	)
	if err != nil {
		log.Logger.With("Error", err).Error("Can't send packet")
	}
}

func SendMessage(conn *websocket.Conn, messageType int, opCode opcodes.OpCode, message string) {
	Send(conn, messageType, opCode, []byte(message))
}

func BroadcastPayload(opCode opcodes.OpCode, message []byte) {
	consumer.Broadcast(
		append(
			[]byte{byte(opCode)},
			message...,
		),
	)
}

func SendError(conn *websocket.Conn, messageType int, opCode opcodes.OpCode, err error) {
	SendMessage(conn, messageType, opCode, err.Error())
}

func Close(conn *websocket.Conn) {
	err := conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Logger.With("Error", err).Error("Connection closed")
	}

	err = conn.Close()
	if err != nil {
		log.Logger.With("Error", err).Error("Can't close connection")
	}
}
