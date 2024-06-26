package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"math/big"

	"entgo.io/contrib/entgql"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/transport/server/gql/generated"
	"github.com/TimeleapLabs/unchained/internal/transport/server/gql/types"
)

// Price is the resolver for the price field.
func (r *assetPriceResolver) Price(ctx context.Context, obj *ent.AssetPrice) (uint64, error) {
	return obj.Price.Uint64(), nil
}

// Signature is the resolver for the signature field.
func (r *assetPriceResolver) Signature(ctx context.Context, obj *ent.AssetPrice) (types.Bytes, error) {
	return obj.Signature, nil
}

// Voted is the resolver for the voted field.
func (r *assetPriceResolver) Voted(ctx context.Context, obj *ent.AssetPrice) (uint64, error) {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// Signature is the resolver for the signature field.
func (r *correctnessReportResolver) Signature(ctx context.Context, obj *ent.CorrectnessReport) (types.Bytes, error) {
	return obj.Signature, nil
}

// Hash is the resolver for the hash field.
func (r *correctnessReportResolver) Hash(ctx context.Context, obj *ent.CorrectnessReport) (types.Bytes, error) {
	return obj.Hash, nil
}

// Topic is the resolver for the topic field.
func (r *correctnessReportResolver) Topic(ctx context.Context, obj *ent.CorrectnessReport) (types.Bytes, error) {
	return obj.Topic, nil
}

// Voted is the resolver for the voted field.
func (r *correctnessReportResolver) Voted(ctx context.Context, obj *ent.CorrectnessReport) (uint64, error) {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// Signature is the resolver for the signature field.
func (r *eventLogResolver) Signature(ctx context.Context, obj *ent.EventLog) (types.Bytes, error) {
	return obj.Signature, nil
}

// Transaction is the resolver for the transaction field.
func (r *eventLogResolver) Transaction(ctx context.Context, obj *ent.EventLog) (types.Bytes, error) {
	return obj.Transaction, nil
}

// Voted is the resolver for the voted field.
func (r *eventLogResolver) Voted(ctx context.Context, obj *ent.EventLog) (uint64, error) {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// AssetPrices is the resolver for the assetPrices field.
func (r *queryResolver) AssetPrices(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AssetPriceOrder, where *ent.AssetPriceWhereInput) (*ent.AssetPriceConnection, error) {
	return r.client.AssetPrice.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithAssetPriceOrder(orderBy),
			ent.WithAssetPriceFilter(where.Filter),
		)
}

// CorrectnessReports is the resolver for the correctnessReports field.
func (r *queryResolver) CorrectnessReports(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.CorrectnessReportOrder, where *ent.CorrectnessReportWhereInput) (*ent.CorrectnessReportConnection, error) {
	return r.client.CorrectnessReport.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCorrectnessReportOrder(orderBy),
			ent.WithCorrectnessReportFilter(where.Filter),
		)
}

// EventLogs is the resolver for the eventLogs field.
func (r *queryResolver) EventLogs(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.EventLogOrder, where *ent.EventLogWhereInput) (*ent.EventLogConnection, error) {
	return r.client.EventLog.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithEventLogOrder(orderBy),
			ent.WithEventLogFilter(where.Filter),
		)
}

// Signers is the resolver for the signers field.
func (r *queryResolver) Signers(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SignerOrder, where *ent.SignerWhereInput) (*ent.SignerConnection, error) {
	return r.client.Signer.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithSignerOrder(orderBy),
			ent.WithSignerFilter(where.Filter),
		)
}

// Key is the resolver for the key field.
func (r *signerResolver) Key(ctx context.Context, obj *ent.Signer) (types.Bytes, error) {
	return obj.Key, nil
}

// Shortkey is the resolver for the shortkey field.
func (r *signerResolver) Shortkey(ctx context.Context, obj *ent.Signer) (types.Bytes, error) {
	return obj.Shortkey, nil
}

// Price is the resolver for the price field.
func (r *assetPriceWhereInputResolver) Price(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.Price = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// PriceNeq is the resolver for the priceNEQ field.
func (r *assetPriceWhereInputResolver) PriceNeq(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.PriceNEQ = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// PriceIn is the resolver for the priceIn field.
func (r *assetPriceWhereInputResolver) PriceIn(ctx context.Context, obj *ent.AssetPriceWhereInput, data []uint64) error {
	for _, value := range data {
		obj.PriceIn = append(obj.PriceIn, &helpers.BigInt{Int: *big.NewInt(int64(value))})
	}
	return nil
}

// PriceNotIn is the resolver for the priceNotIn field.
func (r *assetPriceWhereInputResolver) PriceNotIn(ctx context.Context, obj *ent.AssetPriceWhereInput, data []uint64) error {
	for _, value := range data {
		obj.PriceNotIn = append(obj.PriceNotIn, &helpers.BigInt{Int: *big.NewInt(int64(value))})
	}
	return nil
}

// PriceGt is the resolver for the priceGT field.
func (r *assetPriceWhereInputResolver) PriceGt(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.PriceGT = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// PriceGte is the resolver for the priceGTE field.
func (r *assetPriceWhereInputResolver) PriceGte(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.PriceGTE = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// PriceLt is the resolver for the priceLT field.
func (r *assetPriceWhereInputResolver) PriceLt(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.PriceLT = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// PriceLte is the resolver for the priceLTE field.
func (r *assetPriceWhereInputResolver) PriceLte(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	obj.PriceLTE = &helpers.BigInt{Int: *big.NewInt(int64(*data))}
	return nil
}

// Voted is the resolver for the voted field.
func (r *assetPriceWhereInputResolver) Voted(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// VotedNeq is the resolver for the votedNEQ field.
func (r *assetPriceWhereInputResolver) VotedNeq(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedNeq - votedNEQ"))
}

// VotedIn is the resolver for the votedIn field.
func (r *assetPriceWhereInputResolver) VotedIn(ctx context.Context, obj *ent.AssetPriceWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedIn - votedIn"))
}

// VotedNotIn is the resolver for the votedNotIn field.
func (r *assetPriceWhereInputResolver) VotedNotIn(ctx context.Context, obj *ent.AssetPriceWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedNotIn - votedNotIn"))
}

// VotedGt is the resolver for the votedGT field.
func (r *assetPriceWhereInputResolver) VotedGt(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGt - votedGT"))
}

// VotedGte is the resolver for the votedGTE field.
func (r *assetPriceWhereInputResolver) VotedGte(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGte - votedGTE"))
}

// VotedLt is the resolver for the votedLT field.
func (r *assetPriceWhereInputResolver) VotedLt(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLt - votedLT"))
}

// VotedLte is the resolver for the votedLTE field.
func (r *assetPriceWhereInputResolver) VotedLte(ctx context.Context, obj *ent.AssetPriceWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLte - votedLTE"))
}

// Voted is the resolver for the voted field.
func (r *correctnessReportWhereInputResolver) Voted(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// VotedNeq is the resolver for the votedNEQ field.
func (r *correctnessReportWhereInputResolver) VotedNeq(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedNeq - votedNEQ"))
}

// VotedIn is the resolver for the votedIn field.
func (r *correctnessReportWhereInputResolver) VotedIn(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedIn - votedIn"))
}

// VotedNotIn is the resolver for the votedNotIn field.
func (r *correctnessReportWhereInputResolver) VotedNotIn(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedNotIn - votedNotIn"))
}

// VotedGt is the resolver for the votedGT field.
func (r *correctnessReportWhereInputResolver) VotedGt(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGt - votedGT"))
}

// VotedGte is the resolver for the votedGTE field.
func (r *correctnessReportWhereInputResolver) VotedGte(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGte - votedGTE"))
}

// VotedLt is the resolver for the votedLT field.
func (r *correctnessReportWhereInputResolver) VotedLt(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLt - votedLT"))
}

// VotedLte is the resolver for the votedLTE field.
func (r *correctnessReportWhereInputResolver) VotedLte(ctx context.Context, obj *ent.CorrectnessReportWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLte - votedLTE"))
}

// Voted is the resolver for the voted field.
func (r *eventLogWhereInputResolver) Voted(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: Voted - voted"))
}

// VotedNeq is the resolver for the votedNEQ field.
func (r *eventLogWhereInputResolver) VotedNeq(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedNeq - votedNEQ"))
}

// VotedIn is the resolver for the votedIn field.
func (r *eventLogWhereInputResolver) VotedIn(ctx context.Context, obj *ent.EventLogWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedIn - votedIn"))
}

// VotedNotIn is the resolver for the votedNotIn field.
func (r *eventLogWhereInputResolver) VotedNotIn(ctx context.Context, obj *ent.EventLogWhereInput, data []uint64) error {
	panic(fmt.Errorf("not implemented: VotedNotIn - votedNotIn"))
}

// VotedGt is the resolver for the votedGT field.
func (r *eventLogWhereInputResolver) VotedGt(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGt - votedGT"))
}

// VotedGte is the resolver for the votedGTE field.
func (r *eventLogWhereInputResolver) VotedGte(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedGte - votedGTE"))
}

// VotedLt is the resolver for the votedLT field.
func (r *eventLogWhereInputResolver) VotedLt(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLt - votedLT"))
}

// VotedLte is the resolver for the votedLTE field.
func (r *eventLogWhereInputResolver) VotedLte(ctx context.Context, obj *ent.EventLogWhereInput, data *uint64) error {
	panic(fmt.Errorf("not implemented: VotedLte - votedLTE"))
}

// AssetPrice returns generated.AssetPriceResolver implementation.
func (r *Resolver) AssetPrice() generated.AssetPriceResolver { return &assetPriceResolver{r} }

// CorrectnessReport returns generated.CorrectnessReportResolver implementation.
func (r *Resolver) CorrectnessReport() generated.CorrectnessReportResolver {
	return &correctnessReportResolver{r}
}

// EventLog returns generated.EventLogResolver implementation.
func (r *Resolver) EventLog() generated.EventLogResolver { return &eventLogResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Signer returns generated.SignerResolver implementation.
func (r *Resolver) Signer() generated.SignerResolver { return &signerResolver{r} }

// AssetPriceWhereInput returns generated.AssetPriceWhereInputResolver implementation.
func (r *Resolver) AssetPriceWhereInput() generated.AssetPriceWhereInputResolver {
	return &assetPriceWhereInputResolver{r}
}

// CorrectnessReportWhereInput returns generated.CorrectnessReportWhereInputResolver implementation.
func (r *Resolver) CorrectnessReportWhereInput() generated.CorrectnessReportWhereInputResolver {
	return &correctnessReportWhereInputResolver{r}
}

// EventLogWhereInput returns generated.EventLogWhereInputResolver implementation.
func (r *Resolver) EventLogWhereInput() generated.EventLogWhereInputResolver {
	return &eventLogWhereInputResolver{r}
}

// SignerWhereInput returns generated.SignerWhereInputResolver implementation.
func (r *Resolver) SignerWhereInput() generated.SignerWhereInputResolver {
	return &signerWhereInputResolver{r}
}

type assetPriceResolver struct{ *Resolver }
type correctnessReportResolver struct{ *Resolver }
type eventLogResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type signerResolver struct{ *Resolver }
type assetPriceWhereInputResolver struct{ *Resolver }
type correctnessReportWhereInputResolver struct{ *Resolver }
type eventLogWhereInputResolver struct{ *Resolver }
type signerWhereInputResolver struct{ *Resolver }
