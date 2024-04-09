package pos

import (
	"context"
	"math/big"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ethereum/contracts"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/pos/eip712"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *Repository) Slash(address [20]byte, to common.Address, amount *big.Int, nftIds []*big.Int) error {
	evmAddress, err := s.posContract.EvmAddressOf(nil, address)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to get EVM address of the staker")
		return err
	}

	transfer := contracts.UnchainedStakingEIP712Transfer{
		From:   evmAddress,
		To:     to,
		Amount: amount,
		NftIds: nftIds,
	}

	signature, err := eip712.SignTransferRequest(s.evmSigner, &transfer)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to sign transfer request")
		return err
	}

	tx, err := s.posContract.Transfer(
		nil,
		[]contracts.UnchainedStakingEIP712Transfer{transfer},
		[]contracts.UnchainedStakingSignature{*signature},
	)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to transfer")
		return err
	}

	receipt, err := bind.WaitMined(
		context.Background(),
		s.ethRPC.Clients[config.App.ProofOfStake.Chain],
		tx,
	)

	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Failed to wait for transaction to be mined")
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Logger.
			With("Error", err).
			Error("Transaction failed")
		return err
	}

	log.Logger.
		With("Address", evmAddress.Hex()).
		With("To", to.Hex()).
		With("Amount", amount.String()).
		With("NftIds", nftIds).
		Info("Slashed")

	return nil
}
