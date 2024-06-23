package postgres

import (
	"context"
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
	err := a.client.
		GetConnection().
		WithContext(ctx).
		Table("asset_price").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "block"}, {Name: "chain"}, {Name: "asset"}, {Name: "pair"}},
			UpdateAll: true,
		}).
		Create(&model.DataFrame{
			Hash:      data.Hash(),
			Timestamp: time.Now(),
			Data:      data,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert asset price record in database")
		return consts.ErrInternalError
	}

	return nil
}

func (a AssetPriceRepo) Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]model.AssetPrice, error) {
	currentRecords := []model.AssetPrice{}
	err := a.client.
		GetConnection().
		WithContext(ctx).
		Table("asset_price").
		Where("block", block).
		Where("chain", chain).
		Where("name", name).
		Where("pair", pair).
		Preload("Signer").
		Find(&currentRecords)

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
