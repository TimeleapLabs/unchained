package scheduler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/evmlog"
)

// EvmLog is a scheduler for EvmLog and keep task's dependencies.
type EvmLog struct {
	chain         string
	evmLogService evmlog.Service
}

// Run will trigger by the scheduler and process the EvmLog blocks.
func (e *EvmLog) Run() {
	err := e.evmLogService.ProcessBlocks(context.TODO(), e.chain)
	if err != nil {
		panic(err)
	}
}

// NewEvmLog will create a new EvmLog task.
func NewEvmLog(
	chanName string,
	evmLogService evmlog.Service,
) *EvmLog {
	e := EvmLog{
		chain:         chanName,
		evmLogService: evmLogService,
	}

	return &e
}
