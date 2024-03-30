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
type Signer struct {
	ent.Schema
}

const (
	KeyMaxLen = 96
)

// Fields of the DataSet.
func (Signer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("evm").
			Nillable().
			Optional(),
		field.Bytes("key").
			MaxLen(KeyMaxLen).
			Unique().
			NotEmpty().
			Annotations(entgql.Type("Bytes")),
		field.Bytes("shortkey").
			MaxLen(KeyMaxLen).
			Unique().
			NotEmpty().
			Annotations(entgql.Type("Bytes")),
		field.Int64("points").
			Annotations(entgql.OrderField("POINTS")),
	}
}

// Edges of the DataSet.
func (Signer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assetPrice", AssetPrice.Type).Ref("signers"),
		edge.From("eventLogs", EventLog.Type).Ref("signers"),
		edge.From("correctnessReport", CorrectnessReport.Type).Ref("signers"),
	}
}

// Indexes of the DataSet.
func (Signer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").
			Unique(),
		index.Fields("shortkey").
			Unique(),
	}
}

func (Signer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
