package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TimeleapLabs/unchained/internal/ent"
)

type Database interface {
	GetConnection() *ent.Client
	HealthCheck(ctx context.Context) bool
}

type MongoDatabase interface {
	GetConnection() *mongo.Client
	HealthCheck(ctx context.Context) bool
}
