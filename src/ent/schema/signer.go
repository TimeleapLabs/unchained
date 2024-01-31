package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DataSet holds the schema definition for the DataSet entity.
type Signer struct {
	ent.Schema
}

// Fields of the DataSet.
func (Signer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Bytes("key").MaxLen(48).Unique().NotEmpty(),
		field.Int64("points"),
	}
}

// Edges of the DataSet.
func (Signer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assetPrice", AssetPrice.Type).Ref("signers"),
	}
}

// Edges of the DataSet.
func (Signer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").
			Unique(),
	}
}
