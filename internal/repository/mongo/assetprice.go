package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TimeleapLabs/unchained/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type AssetPriceRepo struct {
	client database.MongoDatabase
}

func (a AssetPriceRepo) Upsert(ctx context.Context, data model.AssetPrice) error {
	opt := options.Update().SetUpsert(true)

	_, err := a.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("assetprice").
		UpdateOne(ctx, bson.M{
			"block": data.Block,
			"chain": data.Chain,
			"asset": data.Name,
			"pair":  data.Pair,
		}, bson.M{
			"$set": bson.M{
				"name":          data.Name,
				"price":         data.Price,
				"signers_count": data.SignersCount,
				"signature":     data.Signature,
				"consensus":     data.Consensus,
				"voted":         data.Voted,
				"signer_ids":    data.SignerIDs,
			},
			"$setOnInsert": bson.M{
				"pair":  data.Pair,
				"chain": data.Chain,
				"block": data.Block,
				"asset": data.Name,
			},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert asset price record in database")
		return consts.ErrInternalError
	}

	return nil
}

func (a AssetPriceRepo) Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]*ent.AssetPrice, error) {
	currentRecords := []*ent.AssetPrice{}
	cursor, err := a.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("assetprice").
		Find(ctx, map[string]interface{}{
			"block": block,
			"chain": chain,
			"asset": name,
			"pair":  pair,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			utils.Logger.With("err", err).Error("Cant close cursor")
		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var result ent.AssetPrice
		err := cursor.Decode(&result)
		if err != nil {
			utils.Logger.With("err", err).Error("Cant decode signer record")
			return nil, err
		}

		currentRecords = append(currentRecords, &result)
	}
	if err := cursor.Err(); err != nil {
		utils.Logger.With("err", err).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func NewAssetPrice(client database.MongoDatabase) repository.AssetPrice {
	return &AssetPriceRepo{
		client: client,
	}
}
