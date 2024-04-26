package repository

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/model"
)

// EvenetLog interface represents the methods that can be used to interact with the EventLog table in the database.
type EventLog interface {
	Find(ctx context.Context, block uint64, hash []byte, index uint64) ([]*ent.EventLog, error)
	Upsert(ctx context.Context, data model.EventLog) error
}

// Signer interface represents the methods that can be used to interact with the Signer table in the database.
type Signer interface {
	CreateSigners(ctx context.Context, signers []model.Signer) error
	GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error)
}

// AssetPrice interface represents the methods that can be used to interact with the AssetPrice table in the database.
type AssetPrice interface {
	Find(ctx context.Context, block uint64, chain string, name string, pair string) ([]*ent.AssetPrice, error)
	Upsert(ctx context.Context, data model.AssetPrice) error
}

// CorrectnessReport interface represents the methods that can be used to interact with the CorrectnessReport table in the database.
type CorrectnessReport interface {
	Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]*ent.CorrectnessReport, error)
	Upsert(ctx context.Context, data model.Correctness) error
}

// MessagingRepository interface represents the methods that can be used to interact with the messaging service.
type MessagingRepository interface {
	Send(ctx context.Context, channel string, payload []byte) error
	Listen(ctx context.Context, channel string) (chan []byte, error)
}
