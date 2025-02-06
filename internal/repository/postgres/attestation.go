package postgres

import (
	"context"
	"encoding/hex"

	"gorm.io/gorm/clause"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/repository"
	"github.com/TimeleapLabs/timeleap/internal/transport/database"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

type AttestationRepo struct {
	client database.Database
}

func (c AttestationRepo) Find(ctx context.Context, hash [32]byte) (model.Attestation, error) {
	currentRecord := model.AttestationDataFrame{}

	tx := c.client.
		GetConnection().
		WithContext(ctx).
		Where("hash", hash).
		First(&currentRecord)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cannot fetch attestation reports from database")
		return currentRecord.Data, consts.ErrInternalError
	}

	return currentRecord.Data, nil
}

func (c AttestationRepo) Upsert(ctx context.Context, hash [32]byte, data model.Attestation) error {
	tx := c.client.
		GetConnection().
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "topic"}, {Name: "hash"}},
			UpdateAll: true,
		}).
		Create(&model.AttestationDataFrame{
			Hash: hex.EncodeToString(hash[:]),
			Data: data,
		})

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cannot upsert attestation report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewAttestation(client database.Database) repository.Attestation {
	return &AttestationRepo{
		client: client,
	}
}
