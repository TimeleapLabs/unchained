package database

import (
	"context"

	"gorm.io/gorm"
)

type Database interface {
	GetConnection() *gorm.DB
	HealthCheck(ctx context.Context) bool
}
