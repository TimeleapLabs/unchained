package mongo

import (
	"context"
	"errors"

	"github.com/TimeleapLabs/unchained/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c AttestationRepo) Find(ctx context.Context, hash [32]byte) (model.Attestation, error) {
	decoded := model.AttestationDataFrame{}
	err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestation").
		FindOne(ctx, bson.M{"hash": hash}).
		Decode(&decoded)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Attestation{}, consts.ErrRecordNotfound
	}

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot decode attestation report from database")
		return model.Attestation{}, consts.ErrInternalError
	}

	return decoded.Data, nil
}

func (c AttestationRepo) Upsert(ctx context.Context, hash [32]byte, data model.Attestation) error {
	opt := options.Update().SetUpsert(true)

	_, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestation").
		UpdateOne(ctx, bson.M{
			"hash": hash,
		}, bson.M{
			"$setOnInsert": bson.M{
				"hash":           hash,
				"data.topic":     data.Topic,
				"data.timestamp": data.Timestamp,
				"data.meta":      data.Meta,
			},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot upsert attestation report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewAttestation(client database.MongoDatabase) repository.Attestation {
	return &AttestationRepo{
		client: client,
	}
}
