package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DataSet holds the schema definition for the DataSet entity.
type EventLog struct {
	ent.Schema
}

// Fields of the DataSet.
func (EventLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("block"),
		field.Uint64("signersCount"),
		field.Bytes("signature").MaxLen(96),
		field.String("address"),
		field.String("chain"),
		field.String("index"),
		field.String("event"),
		field.String("transaction"),
	}
}

// Edges of the DataSet.
func (EventLog) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Make these required on next migrate
		edge.To("signers", Signer.Type).Required(),
		edge.To("args", EventLogArg.Type),
	}
}

// Edges of the DataSet.
func (EventLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("transaction", "index").Unique(),
		index.Fields("block", "address", "event"),
	}
}
