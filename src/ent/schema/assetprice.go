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
		field.Uint64("block"),
		field.Uint64("signersCount").Nillable().Optional(),
		field.String("price").
			GoType(&big.Int{}).
			ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.Bytes("signature").MaxLen(96),
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
