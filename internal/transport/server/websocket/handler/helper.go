package handler

import (
	"github.com/KenshiTech/unchained/internal/constants/opcodes"
	"github.com/KenshiTech/unchained/internal/log"
	"github.com/KenshiTech/unchained/internal/transport/server/websocket/store"
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
	store.Consumers.Range(func(consumer *websocket.Conn, _ bool) bool {
		mu, ok := store.BroadcastMutex.Load(consumer)
		if ok {
			mu.Lock()
			defer mu.Unlock()
			err := consumer.WriteMessage(websocket.BinaryMessage, append(
				[]byte{byte(opCode)},
				message...,
			))
			if err != nil {
				log.Logger.Error(err.Error())
			}
		}
		return true
	})
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
