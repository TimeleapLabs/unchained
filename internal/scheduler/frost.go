package scheduler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"

	"github.com/TimeleapLabs/unchained/internal/service/frost"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// FrostSync is a scheduler for syncing signer of Frost and keep task's dependencies.
type FrostSync struct {
	frostService frost.Service
}

type FrostReadiness struct {
}

// Run will trigger by the scheduler and process the Frost sync.
func (e *FrostSync) Run() {
	utils.Logger.Info("Start synchronizing frost signers")
	ctx := context.TODO()
	err := e.frostService.SyncSigners(ctx)
	if err != nil {
		panic(err)
	}
}

// RunHeartbeatSender will trigger by the scheduler and send heartbeat to show workers readiness.
func (e *FrostReadiness) Run() {
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
