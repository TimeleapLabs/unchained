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

type CorrectnessRepo struct {
	client database.Database
}

func (c CorrectnessRepo) Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]model.Correctness, error) {
	currentRecords := []model.CorrectnessDataFrame{}
	tx := c.client.
		GetConnection().
		WithContext(ctx).
		Where("hash", hash).
		Where("topic", topic).
		Where("timestamp", timestamp).
		Find(&currentRecords)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant fetch correctness reports from database")
		return nil, consts.ErrInternalError
	}

	results := []model.Correctness{}
	for _, record := range currentRecords {
		results = append(results, record.Data)
	}

	return results, nil
}

func (c CorrectnessRepo) Upsert(ctx context.Context, data model.Correctness) error {
	tx := c.client.
		GetConnection().
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "topic"}, {Name: "hash"}},
			UpdateAll: true,
		}).
		Create(&model.CorrectnessDataFrame{
			Hash:      hex.EncodeToString(data.Bls().Marshal()),
			Timestamp: time.Now(),
			Data:      data,
		})

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant upsert correctness report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewCorrectness(client database.Database) repository.CorrectnessReport {
	return &CorrectnessRepo{
		client: client,
	}
}
