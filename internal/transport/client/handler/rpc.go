package handler

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/service/ai"
	"github.com/ethereum/go-ethereum/common"
)

var TIMELEAP_RPC = "https://devnet.timeleap.swiss/rpc"
var COLLECTOR_ADDRESS = common.HexToAddress("0xA2dEc4f8089f89F426e6beB76B555f3Cf9E7f499")

func (h *consumer) RpcRequest(ctx context.Context, message []byte) {
	return
}

func (w worker) RpcRequest(ctx context.Context, message []byte) {
	utils.Logger.Info("RPC Request")
	packet := new(dto.RpcRequest).FromSiaBytes(message)

	// check fees
	checker, err := ai.NewTxChecker(TIMELEAP_RPC)
	if err != nil {
		return
	}

	// 0.1 TLP
	fee, _ := new(big.Int).SetString("100000000000000000", 10)

	ok, err := checker.CheckTransaction(common.HexToHash(packet.TxHash), COLLECTOR_ADDRESS, fee)
	if err != nil || !ok {
		return
	}

	response, err := w.rpc.RunFunction(ctx, packet.Method, packet.Params)
	if err != nil {
		return
	}

	conn.Send(consts.OpCodeRpcResponse, response)
}

func (w worker) RpcResponse(ctx context.Context, message []byte) {}

func (h *consumer) RpcResponse(ctx context.Context, message []byte) {}
