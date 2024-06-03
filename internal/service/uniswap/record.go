package uniswap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/TimeleapLabs/unchained/internal/utils/address"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/puzpuzpuz/xsync/v3"
)

// TODO: This needs to work with different datasets
// TODO: Can we turn this into a library func?
func (s *service) RecordSignature(
	ctx context.Context, signature []byte, signer model.Signer, hash bls12381.G1Affine, info model.PriceInfo, debounce bool, historical bool,
) error {
	if supported := s.SupportedTokens[info.Asset.Token]; !supported {
		utils.Logger.
			With("Name", info.Asset.Token.Name).
			With("Chain", info.Asset.Token.Chain).
			With("Pair", info.Asset.Token.Pair).
			Debug("Token not supported")
		return consts.ErrTokenNotSupported
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := s.ethRPC.GetBlockNumber(ctx, info.Asset.Token.Chain)
	if err != nil {
		s.ethRPC.RefreshRPC(info.Asset.Token.Chain)
		utils.Logger.
			With("Network", info.Asset.Token.Chain).
			With("Error", err).
			Error("Failed to get the latest block number")
		return err
	}

	if !historical && blockNumber-info.Asset.Block > MaxBlockNumberDelta {
		utils.Logger.
			With("Packet", info.Asset.Block).
			With("Current", blockNumber).
			Debug("Data too old")
		return consts.ErrDataTooOld
	}

	if !s.consensus.Contains(info.Asset) {
		s.consensus.Add(info.Asset, *xsync.NewMapOf[bls12381.G1Affine, big.Int]())
	}

	reportedValues, _ := s.consensus.Get(info.Asset)
	isMajority := true
	voted, ok := reportedValues.Load(hash)
	if !ok {
		voted = *big.NewInt(0)
	}

	votingPower, err := s.pos.GetVotingPowerOfEvm(ctx, signer.EvmAddress)
	if err != nil {
		utils.Logger.
			With("Address", address.Calculate(signer.PublicKey[:])).
			With("Error", err).
			Error("Failed to get voting power")
		return err
	}

	totalVoted := new(big.Int).Add(votingPower, &voted)

	reportedValues.Range(func(_ bls12381.G1Affine, value big.Int) bool {
		if value.Cmp(totalVoted) == 1 {
			isMajority = false
		}
		return isMajority
	})

	err = s.checkAndCacheSignature(&reportedValues, signature, signer, hash, totalVoted)
	if err != nil {
		return err
	}

	saveArgs := SaveSignatureArgs{
		Hash:      hash,
		Info:      info,
		Voted:     totalVoted,
		Consensus: isMajority,
	}

	if !debounce {
		err = s.SaveSignatures(ctx, saveArgs)
		if err != nil {
			return err
		}
		return nil
	}

	reportLog := utils.Logger.
		With("Block", info.Asset.Block).
		With("Price", info.Price.String()).
		With("Token", info.Asset.Token.Name)

	reportedValues.Range(func(hash bls12381.G1Affine, value big.Int) bool {
		reportLog = reportLog.With(
			fmt.Sprintf("%x", hash.Bytes())[:8],
			value.String(),
		)
		return true
	})

	reportedValues.Range(func(hash bls12381.G1Affine, value big.Int) bool {
		reportLog = reportLog.With(
			fmt.Sprintf("%x", hash.Bytes())[:8],
			value.String(),
		)
		return true
	})

	reportLog.
		With("Majority", fmt.Sprintf("%x", hash.Bytes())[:8]).
		Debug("Values")

	DebouncedSaveSignatures(info.Asset, saveArgs)

	return nil
}
