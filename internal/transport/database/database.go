package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"gorm.io/gorm"
)

type Database interface {
	GetConnection() *gorm.DB
	Migrate()
	HealthCheck(ctx context.Context) bool
}

type MongoDatabase interface {
	GetConnection() *mongo.Client
	HealthCheck(ctx context.Context) bool
}
