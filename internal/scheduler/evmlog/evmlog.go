package evmlog

import (
	"context"
	"github.com/KenshiTech/unchained/internal/service/evmlog"
)

type EvmLog struct {
	chain         string
	evmLogService evmlog.Service
}

func (e *EvmLog) Run() {
	err := e.evmLogService.ProcessBlocks(context.TODO(), e.chain)
	if err != nil {
		panic(err)
	}
}

func New(
	chanName string,
	evmLogService evmlog.Service,
) *EvmLog {
	e := EvmLog{
		chain:         chanName,
		evmLogService: evmLogService,
	}

	return &e
}
