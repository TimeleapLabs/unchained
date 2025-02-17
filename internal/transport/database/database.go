package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabase interface {
	GetConnection() *mongo.Client
	HealthCheck(ctx context.Context) bool
}
