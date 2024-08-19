package runtime

import (
	"context"
	"io"
	"net"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type UnixPayload struct {
	Size   uint32
	Params []byte
}

func NewUnixPayload(params *dto.RPCRequest) *UnixPayload {
	payload := params.Sia().Bytes()
	return &UnixPayload{
		Size:   uint32(len(payload)),
		Params: payload,
	}
}

func (p *UnixPayload) Sia() sia.Sia {
	return sia.New().AddUInt32(p.Size).EmbedBytes(p.Params)
}

// RunUnixCall runs a function with the given name and parameters.
func RunUnixCall(_ context.Context, conn net.Conn, params *dto.RPCRequest) (*dto.RPCResponse, error) {
	payload := NewUnixPayload(params)
	_, err := conn.Write(payload.Sia().Bytes())
	if err != nil {
		utils.Logger.With("err", err).Error("Error sending message")
		return nil, consts.ErrCantSendRPCRequest
	}

	// Wait for response
	var response []byte
	buf := make([]byte, 1024)
	payloadSize := int(4)

	for payloadSize > len(response) {
		n, err := conn.Read(buf)
		if err == io.EOF {
			utils.Logger.Error("Connection closed")
			break // End of file or connection closed
		} else if err != nil {
			utils.Logger.With("err", err).Error("Error receiving response")
			return nil, consts.ErrCantReceiveRPCResponse
		}

		response = append(response, buf[:n]...)
		if payloadSize == 4 && len(response) >= 4 {
			payloadSize += int(sia.NewFromBytes(response).ReadUInt32())
		}

		if len(response) >= payloadSize {
			break
		}
	}

	utils.Logger.With("length", len(response)).Info("Received response")

	return new(dto.RPCResponse).FromSiaBytes(response[4:]), nil
}
