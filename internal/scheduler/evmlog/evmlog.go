package evmlog

import (
	"github.com/KenshiTech/unchained/internal/service/evmlog"
)

type EvmLog struct {
	chain         string
	evmLogService *evmlog.Service
}

func (e *EvmLog) Run() {
	err := e.evmLogService.ProcessBlocks(e.chain)
	if err != nil {
		panic(err)
	}
}

func New(
	chanName string,
	evmLogService *evmlog.Service,
) *EvmLog {
	e := EvmLog{
		chain:         chanName,
		evmLogService: evmLogService,
	}

	return &e
}
