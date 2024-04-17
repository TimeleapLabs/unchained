package database

import (
	"context"

	"github.com/KenshiTech/unchained/internal/ent"
)

type Database interface {
	GetConnection() *ent.Client
	HealthCheck(ctx context.Context) bool
}
