package runtime

import (
	"context"
	"fmt"
	"net"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// RunUnixCall runs a function with the given name and parameters.
func RunUnixCall(_ context.Context, conn net.Conn, params *dto.RPCRequest) (*dto.RPCResponse, error) {
	fmt.Println(params.Sia().Bytes())
	_, err := conn.Write(params.Sia().Bytes())
	if err != nil {
		utils.Logger.With("err", err).Error("Error sending message")
		return nil, consts.ErrCantSendRPCRequest
	}

	// Wait for response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		utils.Logger.With("err", err).Error("Error receiving response")
		return nil, consts.ErrCantReceiveRPCResponse
	}

	return new(dto.RPCResponse).FromSiaBytes(buf[0:n]), nil
}
