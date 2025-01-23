package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type proofRepo struct {
	client database.MongoDatabase
}

func (s proofRepo) CreateProof(ctx context.Context, hash [32]byte, signers []model.Signature) error {
	_, err := s.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("signatures").
		UpdateOne(ctx, bson.M{
			"hash": hash[:],
		}, bson.M{
			"$addToSet": bson.M{
				"signatures": bson.M{
					"$each": signers,
				},
			},
		}, options.Update().SetUpsert(true))

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot create signatures in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s proofRepo) Find(ctx context.Context, hash [32]byte) (model.Proof, error) {
	var result model.Proof
	err := s.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("signatures").
		FindOne(ctx, bson.M{
			"hash": hash[:],
		}).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Proof{}, consts.ErrRecordNotfound
	}

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot fetch signature record from database")
		return model.Proof{}, consts.ErrInternalError
	}

	return result, nil
}

func (s proofRepo) GetSignerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	opt := options.Find().SetProjection(bson.M{"_id": 1})
	cursor, err := s.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("signatures").
		Find(ctx, bson.M{
			"data.key": bson.M{"$in": keys},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cannot fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	ids := []int{}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			utils.Logger.With("err", err).Error("Cannot close cursor")
		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var result model.Signer
		err := cursor.Decode(&result)
		if err != nil {
			utils.Logger.With("err", err).Error("Cannot decode signer record")
			return nil, err
		}

		ids = append(ids, int(result.ID))
	}
	if err := cursor.Err(); err != nil {
		utils.Logger.With("err", err).Error("Cannot fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	return ids, nil
}

func NewProof(client database.MongoDatabase) repository.Proof {
	return &proofRepo{
		client: client,
	}
}
