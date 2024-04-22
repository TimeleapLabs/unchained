package uniswap

import (
	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/model"
)

func (s *service) ProcessBlocks(chain string) error {
	currBlockNumber, err := s.ethRPC.GetBlockNumber(chain)
	if err != nil {
		s.ethRPC.RefreshRPC(chain)
		return err
	}

	for _, token := range model.NewTokensFromCfg(config.App.Plugins.Uniswap.Tokens) {
		if token.Chain != chain {
			continue
		}

		// TODO: this can be cached
		key := s.TokenKey(token)
		tokenLastBlock, exists := s.LastBlock.Load(*key)

		if !exists {
			s.LastBlock.Store(*key, currBlockNumber-1)
		} else if tokenLastBlock == currBlockNumber {
			return nil
		}

		err = s.SyncBlocks(token, *key, currBlockNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
