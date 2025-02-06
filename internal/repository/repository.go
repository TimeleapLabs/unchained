package repository

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/model"
)

// Proof interface represents the methods that can be used to interact with the Proof table in the database.
type Proof interface {
	CreateProof(ctx context.Context, hash [32]byte, signatures []model.Signature) error
	Find(ctx context.Context, hash [32]byte) (model.Proof, error)
	GetSignerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error)
}

// Attestation interface represents the methods that can be used to interact with the Attestation table in the database.
type Attestation interface {
	Find(ctx context.Context, hash [32]byte) (model.Attestation, error)
	Upsert(ctx context.Context, hash [32]byte, data model.Attestation) error
}
