package postgres

import (
	"context"

	"github.com/TimeleapLabs/timeleap/internal/consts"
	"github.com/TimeleapLabs/timeleap/internal/model"
	"github.com/TimeleapLabs/timeleap/internal/repository"
	"github.com/TimeleapLabs/timeleap/internal/transport/database"
	"github.com/TimeleapLabs/timeleap/internal/utils"
)

type proofRepo struct {
	client database.Database
}

func (s proofRepo) CreateProof(ctx context.Context, hash [32]byte, signatures []model.Signature) error {
	proof := model.NewProof(signatures, hash[:])

	tx := s.client.
		GetConnection().
		WithContext(ctx).
		Create(&proof)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cannot create signers in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s proofRepo) Find(ctx context.Context, hash [32]byte) (model.Proof, error) {
	proof := model.Proof{}

	tx := s.client.
		GetConnection().
		WithContext(ctx).
		Where("hash = ?", hash).
		First(&proof)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cannot fetch signer record from database")
		return model.Proof{}, consts.ErrInternalError
	}

	return proof, nil
}

func (s proofRepo) GetSignerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	ids := []int{}

	tx := s.client.
		GetConnection().
		WithContext(ctx).
		Table("proofs").
		Select("id").
		Where("data.key in ?", keys).
		Find(&ids)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cannot fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	return ids, nil
}

func NewProof(client database.Database) repository.Proof {
	return &proofRepo{
		client: client,
	}
}
