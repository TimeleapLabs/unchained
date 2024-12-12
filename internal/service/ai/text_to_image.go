package ai

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
)

func TextToImage(prompt string, negativePrompt string, model string, loraWeights string, steps uint8) []byte {
	conn, _, err := websocket.DefaultDialer.Dial(
		"ws://localhost:8765", nil,
	)
	if err != nil {
		panic(err)
	}

	closed := false
	defer CloseSocket(conn, &closed)

	incoming := Read(conn, &closed)

	requestUUID, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	uuidBytes, err := requestUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}

	payload := sia.New().
		AddUInt16(0).
		AddByteArrayN(uuidBytes).
		AddString8(model).
		AddString8(loraWeights).
		AddUInt8(steps).
		AddString16(prompt).
		AddString16(negativePrompt).
		Bytes()

	err = conn.WriteMessage(websocket.BinaryMessage, payload)
	if err != nil {
		panic(err)
	}

	// Start a goroutine to print dots every second
	stopDots := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopDots:
				return
			case <-time.After(1 * time.Second):
				fmt.Print(".") //nolint:forbidigo // This is a CLI tool
			}
		}
	}()

	data := <-incoming

	// Stop the dot-printing goroutine
	close(stopDots)

	// process data
	s := sia.NewFromBytes(data)
	uuidBytesFromResponse := s.ReadByteArrayN(16)
	responseUUID, err := uuid.FromBytes(uuidBytesFromResponse)
	if err != nil {
		panic(err)
	}

	if requestUUID != responseUUID {
		panic("UUID mismatch")
	}

	return s.ReadByteArray32()
}
