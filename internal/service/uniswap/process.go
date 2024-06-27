package uniswap

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/uniswap/types"

	"github.com/TimeleapLabs/unchained/internal/config"
)

func (s *service) ProcessBlocks(ctx context.Context, chain string) error {
	currBlockNumber, err := s.ethRPC.GetBlockNumber(ctx, chain)
	if err != nil {
		s.ethRPC.RefreshRPC(chain)
		return err
	}

	for _, token := range types.NewTokensFromCfg(config.App.Plugins.Uniswap.Tokens) {
		if token.Chain != chain {
			continue
		}

		// TODO: this can be cached
		key := types.NewTokenKey(token.GetCrossTokenKeys(s.crossTokens), token)
		tokenLastBlock, exists := s.LastBlock.Load(*key)

		if !exists {
			s.LastBlock.Store(*key, currBlockNumber-1)
		} else if tokenLastBlock == currBlockNumber {
			return nil
		}

		err = s.SyncBlocks(ctx, token, *key, currBlockNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
