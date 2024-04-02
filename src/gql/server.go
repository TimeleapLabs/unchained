package gql

import (
	"net/http"

	"github.com/KenshiTech/unchained/src/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func InstallHandlers() {
	srv := handler.NewDefaultServer(NewSchema(db.GetClient()))
	http.Handle("/gql", playground.Handler("Unchained Playground", "/gql/query"))
	http.Handle("/gql/query", srv)
}
