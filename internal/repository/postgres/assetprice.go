package postgres

import (
	"context"
	"encoding/hex"
	"time"

	"gorm.io/gorm/clause"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type AssetPriceRepo struct {
	client database.Database
}

func (a AssetPriceRepo) Upsert(ctx context.Context, data model.AssetPrice) error {
	dataHash := data.Bls().Bytes()

	tx := a.client.
		GetConnection().
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "block"}, {Name: "chain"}, {Name: "name"}, {Name: "pair"}},
			UpdateAll: true,
		}).
		Create(&model.AssetPriceDataFrame{
			Hash:      hex.EncodeToString(dataHash[:]),
			Timestamp: time.Now(),
			Data:      data,
		})

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant upsert asset price record in database")
		return consts.ErrInternalError
	}

	return nil
}

func (a AssetPriceRepo) Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]model.AssetPrice, error) {
	currentRecords := []model.AssetPriceDataFrame{}
	tx := a.client.
		GetConnection().
		WithContext(ctx).
		Where(model.AssetPriceDataFrame{Data: model.AssetPrice{
			Pair:  pair,
			Name:  name,
			Chain: chain,
			Block: block,
		}}).
		Find(&currentRecords)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	results := []model.AssetPrice{}
	for _, record := range currentRecords {
		results = append(results, record.Data)
	}

	return results, nil
}

func NewAssetPrice(client database.Database) repository.AssetPrice {
	return &AssetPriceRepo{
		client: client,
	}
}
