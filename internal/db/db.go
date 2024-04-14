package db

import (
	"context"
	"database/sql"

	"github.com/KenshiTech/unchained/internal/log"

	"github.com/KenshiTech/unchained/internal/config"
	"github.com/KenshiTech/unchained/internal/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	// these imports are required for ent to work with postgres.
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

var dbClient *ent.Client

// Open new connection.
func Open(databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Start() {
	if config.App.Postgres.URL == "" {
		return
	}

	var err error

	log.Logger.Info("Connecting to DB")

	dbClient, err = ent.Open("postgres", config.App.Postgres.URL)

	if err != nil {
		log.Logger.Error("failed opening connection to postgres: %v", err)
	}

	if err = dbClient.Schema.Create(context.Background()); err != nil {
		log.Logger.Error("failed creating schema resources: %v", err)
	}
}

func GetClient() *ent.Client {
	return dbClient
}
