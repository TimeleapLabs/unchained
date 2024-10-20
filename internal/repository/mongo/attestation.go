package mongo

import (
	"context"
	"time"

	"github.com/TimeleapLabs/unchained/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type AttestationRepo struct {
	client database.MongoDatabase
}

func (c AttestationRepo) Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]model.Attestation, error) {
	cursor, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestationreport").
		Find(ctx, bson.M{
			"hash":      hash,
			"topic":     topic,
			"timestamp": timestamp,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch attestation reports from database")
		return nil, consts.ErrInternalError
	}

	currentRecords, err := CursorToList[model.Attestation](ctx, cursor)
	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch attestation reports from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (c AttestationRepo) Upsert(ctx context.Context, data model.Attestation) error {
	opt := options.Update().SetUpsert(true)

	_, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestationreport").
		UpdateOne(ctx, bson.M{
			"hash":  data.Hash,
			"topic": data.Topic,
		}, bson.M{
			"$set": bson.M{
				"data.correct":       data.Correct,
				"data.signers_count": data.SignersCount,
				"data.signature":     data.Signature,
				"data.timestamp":     data.Timestamp,
				"data.consensus":     data.Consensus,
				"data.voted":         data.Voted,
			},
			"$setOnInsert": bson.M{
				"hash":       data.Bls().Bytes(),
				"timestamp":  time.Now(),
				"data.hash":  data.Hash,
				"data.topic": data.Topic,
			},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert attestation report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewAttestation(client database.MongoDatabase) repository.Attestation {
	return &AttestationRepo{
		client: client,
	}
}
