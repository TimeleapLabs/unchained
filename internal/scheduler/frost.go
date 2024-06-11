package scheduler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"

	"github.com/TimeleapLabs/unchained/internal/service/frost"
)

// FrostSync is a scheduler for syncing signer of Frost and keep task's dependencies.
type FrostSync struct {
	frostService frost.Service
}

type FrostReadiness struct {
}

// Run will trigger by the scheduler and process the Frost sync.
func (e *FrostSync) Run() {
	ctx := context.TODO()
	err := e.frostService.SendOnlineSigners(ctx)
	if err != nil {
		panic(err)
	}
}

// Run will trigger by the scheduler and send heartbeat to show workers readiness.
func (e *FrostReadiness) Run() {
	if conn.IsClosed {
		return
	}

	conn.SendMessage(consts.OpCodeFrostSignerHeartBeat, crypto.Identity.ExportEvmSigner().EvmAddress)
}

// NewFrostSync will create a new FrostSync task.
func NewFrostSync(frostService frost.Service) *FrostSync {
	e := FrostSync{
		frostService: frostService,
	}

	return &e
}

// NewFrostReadiness will create a new FrostReadiness task.
func NewFrostReadiness() *FrostReadiness {
	e := FrostReadiness{}

	return &e
}
