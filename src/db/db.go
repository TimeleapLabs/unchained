package db

import (
	"context"
	"log"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/ent"

	_ "github.com/lib/pq"
)

var dbClient *ent.Client

func Start() {

	if !config.Config.InConfig("database.url") {
		return
	}

	var err error

	dbUrl := config.Config.GetString("database.url")
	dbClient, err = ent.Open("postgres", dbUrl)

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
