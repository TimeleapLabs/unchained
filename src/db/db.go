package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/KenshiTech/unchained/src/config"
	"github.com/KenshiTech/unchained/src/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
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

	dbURL := config.App.Postgres.URL
	dbClient, err = ent.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err = dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// TODO: Properly close the conn
}

func GetClient() *ent.Client {
	return dbClient
}
