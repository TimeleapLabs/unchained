package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DataSet holds the schema definition for the DataSet entity.
type CorrectnessReport struct {
	ent.Schema
}

const (
	SignatureMaxLen = 48
	HashMaxLen      = 64
)

// Fields of the DataSet.
func (CorrectnessReport) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("signersCount").
			Annotations(entgql.Type("Uint")),
		field.Uint64("timestamp").
			Annotations(entgql.OrderField("TIMESTAMP")).
			Annotations(entgql.Type("Uint")),
		field.Bytes("signature").
			MaxLen(SignatureMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.Bytes("hash").
			MaxLen(HashMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.Bytes("topic").
			MaxLen(HashMaxLen).
			Annotations(entgql.Type("Bytes")),
		field.Bool("correct"),
	}
}

// Edges of the DataSet.
func (CorrectnessReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("signers", Signer.Type).Required(),
	}
}

// Edges of the DataSet.
func (CorrectnessReport) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic", "hash").Unique(),
		index.Fields("topic", "timestamp", "hash"),
	}
}

func (CorrectnessReport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
