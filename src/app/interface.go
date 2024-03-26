package app

import "context"

type App interface {
	// Run is the entry point for the application. the Thread will be locked
	Run(ctx context.Context) error
}
