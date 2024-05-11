package postgres

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/assetprice"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type AssetPriceRepo struct {
	client database.Database
}

func (a AssetPriceRepo) Upsert(ctx context.Context, data model.AssetPrice) error {
	err := a.client.
		GetConnection().
		AssetPrice.
		Create().
		SetPair(data.Pair).
		SetAsset(data.Name).
		SetChain(data.Chain).
		SetBlock(data.Block).
		SetPrice(&helpers.BigInt{Int: data.Price}).
		SetSignersCount(data.SignersCount).
		SetSignature(data.Signature).
		SetConsensus(data.Consensus).
		SetVoted(&helpers.BigInt{Int: data.Voted}).
		AddSignerIDs(data.SignerIDs...).
		OnConflictColumns("block", "chain", "asset", "pair").
		UpdateNewValues().
		Update(func(u *ent.AssetPriceUpsert) {
			u.Add("voted", 1)
		}).
		Exec(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert asset price record in database")
		return consts.ErrInternalError
	}

	return nil
}

func (a AssetPriceRepo) Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]*ent.AssetPrice, error) {
	currentRecords, err := a.client.
		GetConnection().
		AssetPrice.
		Query().
		Where(
			assetprice.Block(block),
			assetprice.Chain(chain),
			assetprice.Asset(name),
			assetprice.Pair(pair),
		).
		WithSigners().
		All(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func NewAssetPrice(client database.Database) repository.AssetPrice {
	return &AssetPriceRepo{
		client: client,
	}
}
