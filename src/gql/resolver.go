package gql

import (
	"github.com/KenshiTech/unchained/src/ent"
	"github.com/KenshiTech/unchained/src/gql/generated"

	"github.com/99designs/gqlgen/graphql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
