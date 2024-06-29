package postgres

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type proofRepo struct {
	client database.Database
}

func (s proofRepo) CreateProof(ctx context.Context, signature [48]byte, signers []model.Signer) error {
	proof := model.NewProof(signers, signature[:])

	tx := s.client.
		GetConnection().
		WithContext(ctx).
		Create(&proof)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant create signers in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s proofRepo) GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	ids := []int{}

	tx := s.client.
		GetConnection().
		WithContext(ctx).
		Table("proofs").
		Select("id").
		Where("data.key in ?", keys).
		Find(&ids)

	if tx.Error != nil {
		utils.Logger.With("err", tx.Error).Error("Cant fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	return ids, nil
}

func NewProof(client database.Database) repository.Proof {
	return &proofRepo{
		client: client,
	}
}
