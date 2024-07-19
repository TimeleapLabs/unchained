package fees

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TxChecker contains the Ethereum client and transaction cache.
type TxChecker struct {
	client  *ethclient.Client
	txCache *TxCache
}

// NewTxChecker creates a new TxChecker.
func NewTxChecker(clientURL string) (*TxChecker, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	return &TxChecker{
		client:  client,
		txCache: NewTxCache(),
	}, nil
}

// CheckTransaction checks if a transaction meets the criteria.
func (tc *TxChecker) CheckTransaction(txHash common.Hash, toAddress common.Address, amount *big.Int) (bool, error) {
	if tc.txCache.IsExpired(txHash) {
		return false, fmt.Errorf("transaction is expired")
	}

	tx, isPending, err := tc.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return false, fmt.Errorf("could not retrieve transaction: %v", err)
	}

	if isPending {
		return false, fmt.Errorf("transaction is still pending")
	}

	receipt, err := tc.client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return false, fmt.Errorf("could not retrieve transaction receipt: %v", err)
	}

	if receipt.Status != 1 {
		return false, fmt.Errorf("transaction failed")
	}

	header, err := tc.client.HeaderByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return false, fmt.Errorf("could not retrieve block header: %v", err)
	}

	blockTime := time.Unix(int64(header.Time), 0)
	if time.Since(blockTime) > 5*time.Minute {
		tc.txCache.MarkExpired(txHash)
		return false, fmt.Errorf("transaction is older than 5 minutes")
	}

	if tx.To() == nil || *tx.To() != toAddress {
		return false, nil
	}

	if tx.Value().Cmp(amount) != 0 {
		return false, nil
	}

	return true, nil
}
