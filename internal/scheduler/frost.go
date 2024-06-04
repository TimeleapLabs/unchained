package scheduler

import (
	"github.com/TimeleapLabs/unchained/internal/service/frost"
)

// FrostSync is a scheduler for syncing signer of Frost and keep task's dependencies.
type FrostSync struct {
	frostService frost.Service
}

// Run will trigger by the scheduler and process the Frost sync.
func (e *FrostSync) Run() {
	err := e.frostService.SyncFrost()
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
