package postgres

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"

	// these imports are required for ent to work with postgres.
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

type connection struct {
	db *ent.Client
}

func (c *connection) HealthCheck(_ context.Context) bool {
	return true
}

func (c *connection) GetConnection() *ent.Client {
	if c.db != nil {
		return c.db
	}

	if config.App.Postgres.URL == "" {
		return nil
	}

	var err error

	utils.Logger.Info("Connecting to DB")

	c.db, err = ent.Open("postgres", config.App.Postgres.URL)

	if err != nil {
		utils.Logger.With("err", err).Error("failed opening connection to postgres")
	}

	if err = c.db.Schema.Create(context.Background()); err != nil {
		utils.Logger.With("err", err).Error("failed creating schema resources")
	}

	return c.db
}

func New() database.Database {
	return &connection{}
}
