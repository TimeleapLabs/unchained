package repository

import (
	"context"

	"github.com/KenshiTech/unchained/internal/model"

	"github.com/KenshiTech/unchained/internal/ent"
)

type EventLog interface {
	Find(ctx context.Context, block uint64, hash []byte, index uint64) ([]*ent.EventLog, error)
	Upsert(ctx context.Context, data model.EventLog) error
}

type Signer interface {
	CreateSigners(ctx context.Context, signers []model.Signer) error
	GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error)
}

type AssetPrice interface {
	Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]*ent.AssetPrice, error)
	Upsert(ctx context.Context, data model.AssetPrice) error
}

type CorrectnessReport interface {
	Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]*ent.CorrectnessReport, error)
	Upsert(ctx context.Context, data model.Correctness) error
}

type MessagingRepository interface {
	Send(ctx context.Context, channel string, payload []byte) error
	Listen(ctx context.Context, channel string) (chan []byte, error)
}
