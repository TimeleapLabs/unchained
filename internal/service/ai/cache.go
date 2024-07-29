package ai

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// TxCache is a simple in-memory cache for transaction hashes.
type TxCache struct {
	mu    sync.Mutex
	cache map[common.Hash]struct{}
}

// NewTxCache creates a new TxCache.
func NewTxCache() *TxCache {
	return &TxCache{
		cache: make(map[common.Hash]struct{}),
	}
}

// MarkExpired marks a transaction hash as expired.
func (tc *TxCache) MarkExpired(txHash common.Hash) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.cache[txHash] = struct{}{}
}

// IsExpired checks if a transaction hash is marked as expired.
func (tc *TxCache) IsExpired(txHash common.Hash) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	_, expired := tc.cache[txHash]
	return expired
}
