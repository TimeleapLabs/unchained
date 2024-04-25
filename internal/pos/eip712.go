package pos

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *Repository) Slash(address [20]byte, to common.Address, amount *big.Int, nftIDs []*big.Int) error {
	evmAddress, err := s.posContract.EvmAddressOf(nil, address)

	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get EVM address of the staker")
		return err
	}

	transfer := contracts.UnchainedStakingEIP712Transfer{
		From:   evmAddress,
		To:     to,
		Amount: amount,
		NftIds: nftIDs,
	}

	signature, err := s.eip712Signer.SignTransferRequest(crypto.Identity.Eth, &transfer)

	if err != nil {
		utils.Logger.
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
		utils.Logger.
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
		utils.Logger.
			With("Error", err).
			Error("Failed to wait for transaction to be mined")
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		utils.Logger.
			With("Error", err).
			Error("Transaction failed")
		return err
	}

	utils.Logger.
		With("Address", evmAddress.Hex()).
		With("To", to.Hex()).
		With("Amount", amount.String()).
		With("NftIds", nftIDs).
		Info("Slashed")

	return nil
}
