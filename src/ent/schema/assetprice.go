package schema

import (
	"math/big"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// DataSet holds the schema definition for the DataSet entity.
type AssetPrice struct {
	ent.Schema
}

// Fields of the DataSet.
func (AssetPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("block").Unique(),
		field.Uint64("signersCount").Nillable().Optional(),
		field.String("price").
			GoType(&big.Int{}).
			ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.Bytes("signature").MaxLen(96),
	}
}

// Edges of the DataSet.
func (AssetPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dataSet", DataSet.Type).Required(),
		edge.To("signers", Signer.Type).Required(),
	}
}

// Edges of the DataSet.
func (AssetPrice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("block").
			Unique(),
	}
}
