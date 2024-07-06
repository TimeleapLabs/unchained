package mongo

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type connection struct {
	db *mongo.Client
}

func (c *connection) HealthCheck(ctx context.Context) bool {
	if c.db == nil {
		return false
	}

	err := c.db.Ping(ctx, nil)
	if err != nil {
		utils.Logger.With("Error", err).Error("HealthCheck")
		return false
	}

	return true
}

func (c *connection) GetConnection() *mongo.Client {
	if c.db != nil {
		return c.db
	}

	utils.Logger.Info("Connecting to MongoDB")

	var err error
	nrMon := nrmongo.NewCommandMonitor(nil)
	c.db, err = mongo.Connect(context.Background(), options.Client().SetMonitor(nrMon).ApplyURI(config.App.Mongo.URL))
	if err != nil {
		panic(err)
	}

	return c.db
}

func New() database.MongoDatabase {
	return &connection{}
}
