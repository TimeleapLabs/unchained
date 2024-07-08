package handler

import (
	"context"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/fees"
	"github.com/TimeleapLabs/unchained/internal/rpc"
	"github.com/TimeleapLabs/unchained/internal/service/ai"
	"github.com/ethereum/go-ethereum/common"
)

var TIMELEAP_RPC = "https://devnet.timeleap.swiss/rpc"
var COLLECTOR_ADDRESS = common.HexToAddress("0xA2dEc4f8089f89F426e6beB76B555f3Cf9E7f499")

func (h *consumer) RpcRequest(ctx context.Context, message []byte) []byte {
	return nil
}

func (w worker) RpcRequest(ctx context.Context, message []byte) []byte {
	packet := new(rpc.TextToImageRpcRequest).FromSiaBytes(message)

	// check fees
	checker, err := fees.NewTxChecker(TIMELEAP_RPC)
	if err != nil {
		return nil
	}

	// 0.1 TLP
	fee, _ := new(big.Int).SetString("100000000000000000", 10)

	ok, err := checker.CheckTransaction(common.HexToHash(packet.TxHash), COLLECTOR_ADDRESS, fee)
	if err != nil || !ok {
		return nil
	}

	// process request
	return ai.TextToImage(
		packet.Prompt,
		packet.NegativePrompt,
		packet.Model,
		packet.LoraWeights,
		packet.Steps,
	)
}

func (w worker) RpcResponse(ctx context.Context, message []byte) {}

func (h *consumer) RpcResponse(ctx context.Context, message []byte) {}
