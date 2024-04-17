package gql

import (
	"github.com/KenshiTech/unchained/internal/utils"
	"net/http"

	"github.com/KenshiTech/unchained/internal/transport/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func WithGraphQL(db database.Database) func() {
	return func() {
		utils.Logger.Info("GraphQL service is activated")

		srv := handler.NewDefaultServer(NewSchema(db.GetConnection()))
		http.Handle("/gql", playground.Handler("Unchained Playground", "/gql/query"))
		http.Handle("/gql/query", srv)
	}
}
