package scheduler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/service/frost"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// FrostSync is a scheduler for syncing signer of Frost and keep task's dependencies.
type FrostSync struct {
	frostService frost.Service
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

// NewFrostSync will create a new FrostSync task.
func NewFrostSync(frostService frost.Service) *FrostSync {
	e := FrostSync{
		frostService: frostService,
	}

	return &e
}
