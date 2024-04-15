package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/KenshiTech/unchained/internal/datasets"
	"github.com/KenshiTech/unchained/internal/ent/helpers"
)

// DataSet holds the schema definition for the DataSet entity.
type EventLog struct {
	ent.Schema
}

const (
	TransactionMaxLen = 32
)

// Fields of the DataSet.
func (EventLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("block").
			Annotations(
				entgql.Type("Uint"),
				entgql.OrderField("BLOCK"),
			),
		field.Uint64("signersCount").
			Annotations(entgql.Type("Uint")),
		field.Bytes("signature").
			MaxLen(SignatureMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.String("address"),
		field.String("chain"),
		field.Uint64("index").
			Annotations(entgql.Type("Uint")),
		field.String("event"),
		field.Bytes("transaction").
			MaxLen(TransactionMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.JSON("args", []datasets.EventLogArg{}),
		field.Bool("consensus").Default(false).
			Annotations(entgql.Type("Boolean")),
		field.Uint("voted").
			GoType(new(helpers.BigInt)).
			SchemaType(map[string]string{
				// Uint256
				dialect.SQLite:   "numeric(78, 0)",
				dialect.Postgres: "numeric(78, 0)",
			}).
			Annotations(entgql.Type("Uint")),
	}
}

// Edges of the DataSet.
func (EventLog) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Make these required on next migrate
		edge.To("signers", Signer.Type).Required(),
	}
}

// Edges of the DataSet.
func (EventLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("block", "transaction", "index").Unique(),
		index.Fields("block", "address", "event", "consensus"),
	}
}

func (EventLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
