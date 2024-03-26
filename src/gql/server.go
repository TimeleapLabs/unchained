package gql

import (
	"net/http"

	"github.com/KenshiTech/unchained/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func InstallHandlers() error {
	srv := handler.NewDefaultServer(NewSchema(db.GetClient()))
	http.Handle("/gql", playground.Handler("Unchained Playground", "/gql/query"))
	http.Handle("/gql/query", srv)
	return nil
}
