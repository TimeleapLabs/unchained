package queue

import (
	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/transport/server/packet"
	"github.com/TimeleapLabs/timeleap/internal/utils"
	"github.com/gorilla/websocket"
)

type WebSocketWriter struct {
	Conn   *websocket.Conn
	queue  chan []byte
	closed chan struct{}
}

// NewWebSocketWriter creates a new writer with a dedicated write queue.
func NewWebSocketWriter(conn *websocket.Conn, bufferSize int) *WebSocketWriter {
	writer := &WebSocketWriter{
		Conn:   conn,
		queue:  make(chan []byte, bufferSize),
		closed: make(chan struct{}),
	}

	go writer.run()
	return writer
}

// run processes the write queue.
func (w *WebSocketWriter) run() {
	for {
		select {
		case message := <-w.queue:
			err := w.Conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				utils.Logger.With("Error", err).Error("Cannot send packet")
				return // Exit goroutine on write error.
			}
		case <-w.closed:
			return
		}
	}
}

func (w *WebSocketWriter) SendRaw(payload []byte) {
	select {
	case w.queue <- payload:
		// Message enqueued successfully.
	default:
		// TODO!: Implement a proper queue overflow strategy.
		// We should never reach this point.
		// And if we do, we should NOT drop packets.
		utils.Logger.Error("Write queue is full, dropping packet")
	}
}

// SendRawSigned enqueues a signed message for writing.
func (w *WebSocketWriter) SendRawSigned(payload []byte) {
	signed := packet.New(payload).Sia().Bytes()
	w.SendRaw(signed)
}

// Send enqueues a message for writing.
func (w *WebSocketWriter) Send(opCode consts.OpCode, payload []byte) {
	message := append([]byte{byte(opCode)}, payload...)
	w.SendRaw(message)
}

// Send enqueues a message for writing.
func (w *WebSocketWriter) SendSigned(opCode consts.OpCode, payload []byte) {
	message := append([]byte{byte(opCode)}, payload...)
	signed := packet.New(message).Sia().Bytes()
	w.SendRaw(signed)
}

// SendMessage sends a message to the client.
func (w *WebSocketWriter) SendMessage(opCode consts.OpCode, message string) {
	w.Send(opCode, []byte(message))
}

// SendMessage sends a message to the client.
func (w *WebSocketWriter) SendMessageSigned(opCode consts.OpCode, message string) {
	w.SendSigned(opCode, []byte(message))
}

// SendError sends an error message to the client.
func (w *WebSocketWriter) SendError(opCode consts.OpCode, err error) {
	w.SendMessage(opCode, err.Error())
}

func (w *WebSocketWriter) SendErrorSigned(opCode consts.OpCode, err error) {
	w.SendMessageSigned(opCode, err.Error())
}

// Close shuts down the writer.
func (w *WebSocketWriter) Close() {
	close(w.closed)
	close(w.queue)
}
