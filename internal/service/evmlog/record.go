package evmlog

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/TimeleapLabs/unchained/internal/service/correctness"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/TimeleapLabs/unchained/internal/utils/address"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
)

func (s *service) RecordSignature(
	ctx context.Context, signature bls12381.G1Affine, signer model.Signer, hash bls12381.G1Affine, info model.EventLog, debounce bool, historical bool,
) error {
	supportKey := SupportKey{
		Chain:   info.Chain,
		Address: info.Address,
		Event:   info.Event,
	}
	if supported := s.supportedEvents[supportKey]; !supported {
		utils.Logger.
			With("Chain", info.Chain).
			With("Address", info.Address).
			With("Event", info.Event).
			Debug("Event not supported")
		return consts.ErrEventNotSupported
	}

	// TODO: Standalone mode shouldn't call this or check consensus
	blockNumber, err := s.ethRPC.GetBlockNumber(ctx, info.Chain)
	if err != nil {
		s.ethRPC.RefreshRPC(info.Chain)
		utils.Logger.
			With("Network", info.Chain).
			With("Error", err).
			Error("Failed to get the latest block number")
		return err
	}

	if !historical {
		// TODO: this won't work for Arbitrum
		// TODO: we disallow syncing historical events here
		if blockNumber-info.Block > BlockOutOfRange {
			utils.Logger.
				With("Packet", info.Block).
				With("Current", blockNumber).
				Debug("Data too old")
			return consts.ErrDataTooOld
		}
	}

	s.signatureMutex.Lock()
	defer s.signatureMutex.Unlock()

	key := EventKey{
		Chain:    info.Chain,
		TxHash:   info.TxHash,
		LogIndex: info.LogIndex,
	}

	if !s.consensus.Contains(key) {
		s.consensus.Add(key, make(map[bls12381.G1Affine]big.Int))
	}

	votingPower, err := s.pos.GetVotingPowerOfEvm(ctx, signer.EvmAddress)
	if err != nil {
		publicKeyBytes, err := hex.DecodeString(signer.PublicKey)
		if err != nil {
			utils.Logger.Error("Can't decode public key: %v", err)
			return err
		}

		utils.Logger.
			With("Address", address.Calculate(publicKeyBytes)).
			With("Error", err).
			Error("Failed to get voting power")
		return err
	}

	reportedValues, _ := s.consensus.Get(key)
	voted := reportedValues[hash]
	totalVoted := new(big.Int).Add(votingPower, &voted)
	isMajority := true

	for _, reportCount := range reportedValues {
		if reportCount.Cmp(totalVoted) == 1 {
			isMajority = false
			break
		}
	}

	cached, _ := s.signatureCache.Get(hash)

	packed := correctness.Signature{
		Signature: signature,
		Signer:    signer,
	}

	for _, item := range cached {
		if item.Signer.PublicKey == signer.PublicKey {
			return consts.ErrDuplicateSignature
		}
	}

	reportedValues[hash] = *totalVoted
	cached = append(cached, packed)
	s.signatureCache.Add(hash, cached)

	saveArgs := SaveSignatureArgs{
		Hash:      hash,
		Info:      info,
		Consensus: isMajority,
		Voted:     totalVoted,
	}

	if debounce {
		s.DebouncedSaveSignatures(hash, saveArgs)
		return nil
	}

	err = s.SaveSignatures(ctx, saveArgs)
	if err != nil {
		return err
	}

	return nil
}
