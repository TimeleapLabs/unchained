//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithWhereInputs(true),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("gql/unchained.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}
	config := &gen.Config{Features: []gen.Feature{gen.FeatureUpsert}}
	if err := entc.Generate("./ent/schema", config, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
