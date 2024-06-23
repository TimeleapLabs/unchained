package postgres

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type signerRepo struct {
	client database.Database
}

func (s signerRepo) CreateSigners(ctx context.Context, signers []model.Signer) error {
	err := s.client.
		GetConnection().
		WithContext(ctx).
		Table("signer").
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "shortkey"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"name":  "name",
				"evm":   "evm",
				"key":   "key",
				"point": gorm.Expr("signer.point + 1"),
			}),
		}).
		CreateInBatches(&signers, 100)

	// Update(func(su *ent.SignerUpsert) {
	//	su.AddPoints(1)
	// }).

	if err != nil {
		utils.Logger.With("err", err).Error("Cant create signers in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s signerRepo) GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	ids := []int{}

	err := s.client.
		GetConnection().
		WithContext(ctx).
		Table("signer").
		Select("id").
		Where("data.key in ?", keys).
		Find(&ids)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	return ids, nil
}

func NewSigner(client database.Database) repository.Signer {
	return &signerRepo{
		client: client,
	}
}
