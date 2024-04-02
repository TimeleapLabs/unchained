package gql

import (
	"github.com/KenshiTech/unchained/log"
	"net/http"

	"github.com/KenshiTech/unchained/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func WithGraphQL() func() {
	return func() {
		log.Logger.Info("GraphQL is activated")

		srv := handler.NewDefaultServer(NewSchema(db.GetClient()))
		http.Handle("/gql", playground.Handler("Unchained Playground", "/gql/query"))
		http.Handle("/gql/query", srv)
	}
}
