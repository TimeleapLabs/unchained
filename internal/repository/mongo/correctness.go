package mongo

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type CorrectnessRepo struct {
	client database.MongoDatabase
}

func (c CorrectnessRepo) Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]model.Correctness, error) {
	cursor, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("correctnessreport").
		Find(ctx, bson.M{
			"hash":      hash,
			"topic":     topic,
			"timestamp": timestamp,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch correctness reports from database")
		return nil, consts.ErrInternalError
	}

	currentRecords, err := CursorToList[model.Correctness](ctx, cursor)
	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch correctness reports from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (c CorrectnessRepo) Upsert(ctx context.Context, data model.Correctness) error {
	opt := options.Update().SetUpsert(true)

	_, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("correctnessreport").
		UpdateOne(ctx, bson.M{
			"hash":  data.Hash,
			"topic": data.Topic[:],
		}, bson.M{
			"$set": bson.M{
				"correct":       data.Correct,
				"signers_count": data.SignersCount,
				"signature":     data.Signature,
				"timestamp":     data.Timestamp,
				"consensus":     data.Consensus,
				"voted":         data.Voted,
				"signer_ids":    data.SignerIDs,
			},
			"$setOnInsert": bson.M{
				"correct":       data.Correct,
				"signers_count": data.SignersCount,
				"signature":     data.Signature,
				"hash":          data.Hash,
				"timestamp":     data.Timestamp,
				"topic":         data.Topic[:],
				"consensus":     data.Consensus,
				"voted":         data.Voted,
				"signer_ids":    data.SignerIDs,
			},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert correctness report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewCorrectness(client database.MongoDatabase) repository.CorrectnessReport {
	return &CorrectnessRepo{
		client: client,
	}
}
