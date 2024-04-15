// Code generated by ent, DO NOT EDIT.

package assetprice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/KenshiTech/unchained/internal/ent/helpers"
	"github.com/KenshiTech/unchained/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldID, id))
}

// Block applies equality check predicate on the "block" field. It's identical to BlockEQ.
func Block(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldBlock, v))
}

// SignersCount applies equality check predicate on the "signersCount" field. It's identical to SignersCountEQ.
func SignersCount(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldSignersCount, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldPrice, v))
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldSignature, v))
}

// Asset applies equality check predicate on the "asset" field. It's identical to AssetEQ.
func Asset(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldAsset, v))
}

// Chain applies equality check predicate on the "chain" field. It's identical to ChainEQ.
func Chain(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldChain, v))
}

// Pair applies equality check predicate on the "pair" field. It's identical to PairEQ.
func Pair(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldPair, v))
}

// Consensus applies equality check predicate on the "consensus" field. It's identical to ConsensusEQ.
func Consensus(v bool) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldConsensus, v))
}

// Voted applies equality check predicate on the "voted" field. It's identical to VotedEQ.
func Voted(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldVoted, v))
}

// BlockEQ applies the EQ predicate on the "block" field.
func BlockEQ(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldBlock, v))
}

// BlockNEQ applies the NEQ predicate on the "block" field.
func BlockNEQ(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldBlock, v))
}

// BlockIn applies the In predicate on the "block" field.
func BlockIn(vs ...uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldBlock, vs...))
}

// BlockNotIn applies the NotIn predicate on the "block" field.
func BlockNotIn(vs ...uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldBlock, vs...))
}

// BlockGT applies the GT predicate on the "block" field.
func BlockGT(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldBlock, v))
}

// BlockGTE applies the GTE predicate on the "block" field.
func BlockGTE(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldBlock, v))
}

// BlockLT applies the LT predicate on the "block" field.
func BlockLT(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldBlock, v))
}

// BlockLTE applies the LTE predicate on the "block" field.
func BlockLTE(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldBlock, v))
}

// SignersCountEQ applies the EQ predicate on the "signersCount" field.
func SignersCountEQ(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldSignersCount, v))
}

// SignersCountNEQ applies the NEQ predicate on the "signersCount" field.
func SignersCountNEQ(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldSignersCount, v))
}

// SignersCountIn applies the In predicate on the "signersCount" field.
func SignersCountIn(vs ...uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldSignersCount, vs...))
}

// SignersCountNotIn applies the NotIn predicate on the "signersCount" field.
func SignersCountNotIn(vs ...uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldSignersCount, vs...))
}

// SignersCountGT applies the GT predicate on the "signersCount" field.
func SignersCountGT(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldSignersCount, v))
}

// SignersCountGTE applies the GTE predicate on the "signersCount" field.
func SignersCountGTE(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldSignersCount, v))
}

// SignersCountLT applies the LT predicate on the "signersCount" field.
func SignersCountLT(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldSignersCount, v))
}

// SignersCountLTE applies the LTE predicate on the "signersCount" field.
func SignersCountLTE(v uint64) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldSignersCount, v))
}

// SignersCountIsNil applies the IsNil predicate on the "signersCount" field.
func SignersCountIsNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIsNull(FieldSignersCount))
}

// SignersCountNotNil applies the NotNil predicate on the "signersCount" field.
func SignersCountNotNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotNull(FieldSignersCount))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...*helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...*helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldPrice, v))
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldSignature, v))
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldSignature, v))
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...[]byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldSignature, vs...))
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...[]byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldSignature, vs...))
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldSignature, v))
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldSignature, v))
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldSignature, v))
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v []byte) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldSignature, v))
}

// AssetEQ applies the EQ predicate on the "asset" field.
func AssetEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldAsset, v))
}

// AssetNEQ applies the NEQ predicate on the "asset" field.
func AssetNEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldAsset, v))
}

// AssetIn applies the In predicate on the "asset" field.
func AssetIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldAsset, vs...))
}

// AssetNotIn applies the NotIn predicate on the "asset" field.
func AssetNotIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldAsset, vs...))
}

// AssetGT applies the GT predicate on the "asset" field.
func AssetGT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldAsset, v))
}

// AssetGTE applies the GTE predicate on the "asset" field.
func AssetGTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldAsset, v))
}

// AssetLT applies the LT predicate on the "asset" field.
func AssetLT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldAsset, v))
}

// AssetLTE applies the LTE predicate on the "asset" field.
func AssetLTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldAsset, v))
}

// AssetContains applies the Contains predicate on the "asset" field.
func AssetContains(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContains(FieldAsset, v))
}

// AssetHasPrefix applies the HasPrefix predicate on the "asset" field.
func AssetHasPrefix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasPrefix(FieldAsset, v))
}

// AssetHasSuffix applies the HasSuffix predicate on the "asset" field.
func AssetHasSuffix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasSuffix(FieldAsset, v))
}

// AssetIsNil applies the IsNil predicate on the "asset" field.
func AssetIsNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIsNull(FieldAsset))
}

// AssetNotNil applies the NotNil predicate on the "asset" field.
func AssetNotNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotNull(FieldAsset))
}

// AssetEqualFold applies the EqualFold predicate on the "asset" field.
func AssetEqualFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEqualFold(FieldAsset, v))
}

// AssetContainsFold applies the ContainsFold predicate on the "asset" field.
func AssetContainsFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContainsFold(FieldAsset, v))
}

// ChainEQ applies the EQ predicate on the "chain" field.
func ChainEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldChain, v))
}

// ChainNEQ applies the NEQ predicate on the "chain" field.
func ChainNEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldChain, v))
}

// ChainIn applies the In predicate on the "chain" field.
func ChainIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldChain, vs...))
}

// ChainNotIn applies the NotIn predicate on the "chain" field.
func ChainNotIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldChain, vs...))
}

// ChainGT applies the GT predicate on the "chain" field.
func ChainGT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldChain, v))
}

// ChainGTE applies the GTE predicate on the "chain" field.
func ChainGTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldChain, v))
}

// ChainLT applies the LT predicate on the "chain" field.
func ChainLT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldChain, v))
}

// ChainLTE applies the LTE predicate on the "chain" field.
func ChainLTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldChain, v))
}

// ChainContains applies the Contains predicate on the "chain" field.
func ChainContains(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContains(FieldChain, v))
}

// ChainHasPrefix applies the HasPrefix predicate on the "chain" field.
func ChainHasPrefix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasPrefix(FieldChain, v))
}

// ChainHasSuffix applies the HasSuffix predicate on the "chain" field.
func ChainHasSuffix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasSuffix(FieldChain, v))
}

// ChainIsNil applies the IsNil predicate on the "chain" field.
func ChainIsNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIsNull(FieldChain))
}

// ChainNotNil applies the NotNil predicate on the "chain" field.
func ChainNotNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotNull(FieldChain))
}

// ChainEqualFold applies the EqualFold predicate on the "chain" field.
func ChainEqualFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEqualFold(FieldChain, v))
}

// ChainContainsFold applies the ContainsFold predicate on the "chain" field.
func ChainContainsFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContainsFold(FieldChain, v))
}

// PairEQ applies the EQ predicate on the "pair" field.
func PairEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldPair, v))
}

// PairNEQ applies the NEQ predicate on the "pair" field.
func PairNEQ(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldPair, v))
}

// PairIn applies the In predicate on the "pair" field.
func PairIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldPair, vs...))
}

// PairNotIn applies the NotIn predicate on the "pair" field.
func PairNotIn(vs ...string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldPair, vs...))
}

// PairGT applies the GT predicate on the "pair" field.
func PairGT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldPair, v))
}

// PairGTE applies the GTE predicate on the "pair" field.
func PairGTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldPair, v))
}

// PairLT applies the LT predicate on the "pair" field.
func PairLT(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldPair, v))
}

// PairLTE applies the LTE predicate on the "pair" field.
func PairLTE(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldPair, v))
}

// PairContains applies the Contains predicate on the "pair" field.
func PairContains(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContains(FieldPair, v))
}

// PairHasPrefix applies the HasPrefix predicate on the "pair" field.
func PairHasPrefix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasPrefix(FieldPair, v))
}

// PairHasSuffix applies the HasSuffix predicate on the "pair" field.
func PairHasSuffix(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldHasSuffix(FieldPair, v))
}

// PairIsNil applies the IsNil predicate on the "pair" field.
func PairIsNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIsNull(FieldPair))
}

// PairNotNil applies the NotNil predicate on the "pair" field.
func PairNotNil() predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotNull(FieldPair))
}

// PairEqualFold applies the EqualFold predicate on the "pair" field.
func PairEqualFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEqualFold(FieldPair, v))
}

// PairContainsFold applies the ContainsFold predicate on the "pair" field.
func PairContainsFold(v string) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldContainsFold(FieldPair, v))
}

// ConsensusEQ applies the EQ predicate on the "consensus" field.
func ConsensusEQ(v bool) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldConsensus, v))
}

// ConsensusNEQ applies the NEQ predicate on the "consensus" field.
func ConsensusNEQ(v bool) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldConsensus, v))
}

// VotedEQ applies the EQ predicate on the "voted" field.
func VotedEQ(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldEQ(FieldVoted, v))
}

// VotedNEQ applies the NEQ predicate on the "voted" field.
func VotedNEQ(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNEQ(FieldVoted, v))
}

// VotedIn applies the In predicate on the "voted" field.
func VotedIn(vs ...*helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldIn(FieldVoted, vs...))
}

// VotedNotIn applies the NotIn predicate on the "voted" field.
func VotedNotIn(vs ...*helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldNotIn(FieldVoted, vs...))
}

// VotedGT applies the GT predicate on the "voted" field.
func VotedGT(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGT(FieldVoted, v))
}

// VotedGTE applies the GTE predicate on the "voted" field.
func VotedGTE(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldGTE(FieldVoted, v))
}

// VotedLT applies the LT predicate on the "voted" field.
func VotedLT(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLT(FieldVoted, v))
}

// VotedLTE applies the LTE predicate on the "voted" field.
func VotedLTE(v *helpers.BigInt) predicate.AssetPrice {
	return predicate.AssetPrice(sql.FieldLTE(FieldVoted, v))
}

// HasSigners applies the HasEdge predicate on the "signers" edge.
func HasSigners() predicate.AssetPrice {
	return predicate.AssetPrice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SignersTable, SignersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSignersWith applies the HasEdge predicate on the "signers" edge with a given conditions (other predicates).
func HasSignersWith(preds ...predicate.Signer) predicate.AssetPrice {
	return predicate.AssetPrice(func(s *sql.Selector) {
		step := newSignersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AssetPrice) predicate.AssetPrice {
	return predicate.AssetPrice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AssetPrice) predicate.AssetPrice {
	return predicate.AssetPrice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AssetPrice) predicate.AssetPrice {
	return predicate.AssetPrice(sql.NotPredicates(p))
}
