package repository

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/model"
)

// Proof interface represents the methods that can be used to interact with the Proof table in the database.
type Proof interface {
	CreateProof(ctx context.Context, signature [48]byte, signers []model.Signer) error
	GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error)
}

// Signer interface represents the methods that can be used to interact with the Signer table in the database.
type Signer interface {
	CreateSigners(ctx context.Context, signers []model.Signer) error
	GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error)
}

// Attestation interface represents the methods that can be used to interact with the Attestation table in the database.
type Attestation interface {
	Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]model.Attestation, error)
	Upsert(ctx context.Context, data model.Attestation) error
}

// MessagingRepository interface represents the methods that can be used to interact with the messaging service.
type MessagingRepository interface {
	Send(ctx context.Context, channel string, payload []byte) error
	Listen(ctx context.Context, channel string) (chan []byte, error)
}
