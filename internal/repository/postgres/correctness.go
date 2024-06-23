package postgres

import (
	"context"
	"time"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"gorm.io/gorm/clause"
)

type CorrectnessRepo struct {
	client database.Database
}

func (c CorrectnessRepo) Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]model.Correctness, error) {
	currentRecords := []model.Correctness{}
	err := c.client.
		GetConnection().
		WithContext(ctx).
		Table("correctness").
		Where("hash", hash).
		Where("topic", topic).
		Where("timestamp", timestamp).
		Find(&currentRecords)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch correctness reports from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (c CorrectnessRepo) Upsert(ctx context.Context, data model.Correctness) error {
	dataBls := data.Bls()
	dataBlsHash := (&dataBls).Marshal()

	err := c.client.
		GetConnection().
		WithContext(ctx).
		Table("correctness").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "topic"}, {Name: "hash"}},
			UpdateAll: true,
		}).
		Create(&model.DataFrame{
			Hash:      dataBlsHash,
			Timestamp: time.Now(),
			Data:      data,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert correctness report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewCorrectness(client database.Database) repository.CorrectnessReport {
	return &CorrectnessRepo{
		client: client,
	}
}
