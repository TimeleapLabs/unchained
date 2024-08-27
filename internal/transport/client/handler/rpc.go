package handler

import (
	"context"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/service/rpc/dto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"github.com/TimeleapLabs/unchained/internal/service/ai"
	"github.com/ethereum/go-ethereum/common"
)

var TimeleapRPC = "https://devnet.timeleap.swiss/rpc"
var CollectorAddress = common.HexToAddress("0xA2dEc4f8089f89F426e6beB76B555f3Cf9E7f499")

func (h *consumer) RPCRequest(_ context.Context, _ []byte) {}

func (w worker) RPCRequest(ctx context.Context, message []byte) {
	utils.Logger.Info("RPC Request")
	packet := new(dto.RPCRequest).FromSiaBytes(message)

	// check fees
	checker, err := ai.NewTxChecker(TimeleapRPC)
	if err != nil {
		return
	}

	// 0.1 TLP
	fee, _ := new(big.Int).SetString("100000000000000000", 10)

	ok, err := checker.CheckTransaction(common.HexToHash(packet.TxHash), CollectorAddress, fee)
	if err != nil || !ok {
		return
	}

	response, err := w.rpc.RunFunction(ctx, packet.Method, packet)
	if err != nil {
		return
	}

	conn.Send(consts.OpCodeRPCResponse, response)
}

func (w worker) RPCResponse(_ context.Context, _ []byte) {}

func (h *consumer) RPCResponse(_ context.Context, _ []byte) {}
