package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DataSet holds the schema definition for the DataSet entity.
type DataSet struct {
	ent.Schema
}

// Fields of the DataSet.
func (DataSet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the DataSet.
func (DataSet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assetPrice", AssetPrice.Type).Ref("dataSet"),
	}
}

// Edges of the DataSet.
func (DataSet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
