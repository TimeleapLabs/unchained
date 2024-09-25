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

type EventLogRepo struct {
	client database.MongoDatabase
}

func (r EventLogRepo) Find(ctx context.Context, block uint64, hash []byte, index uint64) ([]model.EventLog, error) {
	cursor, err := r.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("eventlog").
		Find(ctx, bson.M{
			"block":       block,
			"transaction": hash,
			"index":       index,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch event log records from database")
		return nil, consts.ErrInternalError
	}

	currentRecords, err := CursorToList[model.EventLog](ctx, cursor)
	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch event log records from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (r EventLogRepo) Upsert(ctx context.Context, data model.EventLog) error {
	opt := options.Update().SetUpsert(true)

	_, err := r.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("eventlog").
		UpdateOne(ctx, bson.M{
			"block":       data.Block,
			"transaction": data.TxHash[:],
			"index":       data.LogIndex,
		}, bson.M{
			"$set": bson.M{
				"data.chain":         data.Chain,
				"data.address":       data.Address,
				"data.event":         data.Event,
				"data.signers_count": data.SignersCount,
				"data.signature":     data.Signature,
				"data.args":          data.Args,
				"data.consensus":     data.Consensus,
				"data.voted":         data.Voted,
			},
			"$setOnInsert": bson.M{
				"hash":      data.Bls().Bytes(),
				"timestamp": time.Now(),
				"data": bson.M{
					"block":         data.Block,
					"chain":         data.Chain,
					"address":       data.Address,
					"event":         data.Event,
					"index":         data.LogIndex,
					"transaction":   data.TxHash[:],
					"signers_count": data.SignersCount,
					"signature":     data.Signature,
					"args":          data.Args,
					"consensus":     data.Consensus,
					"voted":         data.Voted,
				},
			},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert event log record to database")
		return consts.ErrInternalError
	}

	return nil
}

func NewEventLog(client database.MongoDatabase) repository.EventLog {
	return &EventLogRepo{
		client: client,
	}
}
