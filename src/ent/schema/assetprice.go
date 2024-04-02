package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/KenshiTech/unchained/src/ent/helpers"
)

// DataSet holds the schema definition for the DataSet entity.
type AssetPrice struct {
	ent.Schema
}

// Fields of the DataSet.
func (AssetPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("block").
			Annotations(
				entgql.Type("Uint"),
				entgql.OrderField("BLOCK"),
			),
		field.Uint64("signersCount").Nillable().Optional().
			Annotations(entgql.Type("Uint")),
		field.Uint("price").
			GoType(new(helpers.BigInt)).
			SchemaType(map[string]string{
				// Uint256
				dialect.SQLite:   "numeric(78, 0)",
				dialect.Postgres: "numeric(78, 0)",
			}).
			Annotations(entgql.Type("Uint")),
		field.Bytes("signature").
			MaxLen(SignatureMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.String("asset").Optional(),
		field.String("chain").Optional(),
		field.String("pair").Optional(),
	}
}

// Edges of the DataSet.
func (AssetPrice) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Make these required on next migrate
		edge.To("signers", Signer.Type).Required(),
	}
}

// Edges of the DataSet.
func (AssetPrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("block", "chain", "asset", "pair").Unique(),
	}
}

func (AssetPrice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
