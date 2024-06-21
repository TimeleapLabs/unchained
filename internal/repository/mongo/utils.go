package mongo

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func CursorToList[T any](ctx context.Context, cursor *mongo.Cursor) ([]T, error) {
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			utils.Logger.With("err", err).Error("Cant close cursor")
		}
	}(cursor, ctx)

	result := []T{}
	for cursor.Next(ctx) {
		var item T
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	if err := cursor.Err(); err != nil {
		utils.Logger.With("err", err).Error("Cant fetch asset price records from database")
		return nil, consts.ErrInternalError
	}

	return result, nil
}
