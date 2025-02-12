package handler

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/service/rpc/dto"
	"github.com/TimeleapLabs/timeleap/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var TimeleapRPC = "https://devnet.timeleap.swiss/rpc"
var CollectorAddress = common.HexToAddress("0xA2dEc4f8089f89F426e6beB76B555f3Cf9E7f499")

func (w *worker) Message(_ context.Context, _ []byte) {}

// RPCRequest is a method that handles RPC request packets and call the corresponding function.
func (w worker) RPCRequest(ctx context.Context, message []byte) {
	packet := new(dto.RPCRequest).FromSiaBytes(message)

	utils.Logger.
		With("ID", packet.ID).
		With("Plugin", packet.Plugin).
		With("Function", packet.Method).
		Info("RPC Request")

	// check fees
	// checker, err := ai.NewTxChecker(TimeleapRPC)
	// if err != nil {
	//	 return
	// }

	// 0.1 TLP
	// fee, _ := new(big.Int).SetString("100000000000000000", 10)

	// ok, err := checker.CheckTransaction(common.HexToHash(packet.TxHash), CollectorAddress, fee)
	// if err != nil || !ok {
	//	 return
	// }

	err := w.rpc.RunFunction(ctx, packet.Plugin, packet)
	if err != nil {
		return
	}
}

// RPCResponse is not defined for worker nodes.
func (w worker) RPCResponse(_ context.Context, _ []byte) {}
