package mongo

import (
	"context"

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

func (s proofRepo) CreateProof(ctx context.Context, signature [48]byte, signers []model.Signer) error {
	proof := model.NewProof(signers, signature[:])

	_, err := s.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("signer").
		InsertOne(ctx, proof)
	if err != nil {
		utils.Logger.With("err", err).Error("Cant create signers in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s proofRepo) GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	opt := options.Find().SetProjection(bson.M{"_id": 1})
	cursor, err := s.client.
		GetConnection().
		Database(config.App.Mongo.Database).
		Collection("signer").
		Find(ctx, bson.M{
			"data.key": bson.M{"$in": keys},
		}, opt)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	ids := []int{}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			utils.Logger.With("err", err).Error("Cant close cursor")
		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var result model.Signer
		err := cursor.Decode(&result)
		if err != nil {
			utils.Logger.With("err", err).Error("Cant decode signer record")
			return nil, err
		}

		ids = append(ids, int(result.ID))
	}
	if err := cursor.Err(); err != nil {
		utils.Logger.With("err", err).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	return ids, nil
}

func NewProof(client database.MongoDatabase) repository.Proof {
	return &proofRepo{
		client: client,
	}
}
