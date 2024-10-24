package mongo

import (
	"context"
	"time"

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

func (c AttestationRepo) Find(ctx context.Context, hash []byte) (model.Attestation, error) {
	decoded := model.AttestationDataFrame{}
	err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestation").
		FindOne(ctx, bson.M{"hash": hash}).
		Decode(&decoded)

	if err == mongo.ErrNoDocuments {
		return model.Attestation{}, consts.ErrRecordNotfound
	}

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot decode attestation report from database")
		return model.Attestation{}, consts.ErrInternalError
	}

	return decoded.Data, nil
}

func (c AttestationRepo) Upsert(ctx context.Context, data model.Attestation) error {
	opt := options.Update().SetUpsert(true)

	_, err := c.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("attestation").
		UpdateOne(ctx, bson.M{
			"hash": data.Bls().Bytes(),
		}, bson.M{
			"$set": bson.M{
				"data.meta":          data.Meta,
				"data.signers_count": data.SignersCount,
				"data.signature":     data.Signature,
				"data.timestamp":     data.Timestamp,
				"data.consensus":     data.Consensus, // TODO: Remove this field?
				"data.voted":         data.Voted,     // TODO: Improve or remove this field
			},
			"$setOnInsert": bson.M{
				"hash":       data.Bls().Bytes(),
				"timestamp":  time.Now(),
				"data.hash":  data.Hash,
				"data.topic": data.Topic,
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
