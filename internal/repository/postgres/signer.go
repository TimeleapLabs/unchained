package postgres

import (
	"context"

	"github.com/KenshiTech/unchained/internal/model"
	"github.com/KenshiTech/unchained/internal/utils"

	"github.com/KenshiTech/unchained/internal/consts"

	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/ent/signer"
	"github.com/KenshiTech/unchained/internal/repository"
	"github.com/KenshiTech/unchained/internal/transport/database"
)

type signerRepo struct {
	client database.Database
}

func (s signerRepo) CreateSigners(ctx context.Context, signers []model.Signer) error {
	err := s.client.
		GetConnection().
		Signer.MapCreateBulk(signers, func(sc *ent.SignerCreate, i int) {
		signer := signers[i]
		sc.SetName(signer.Name).
			SetEvm(signer.EvmAddress).
			SetKey(signer.PublicKey[:]).
			SetShortkey(signer.ShortPublicKey[:]).
			SetPoints(0)
	}).
		OnConflictColumns("shortkey").
		UpdateName().
		UpdateEvm().
		UpdateKey().
		Update(func(su *ent.SignerUpsert) {
			su.AddPoints(1)
		}).
		Exec(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant create signers in database")
		return consts.ErrInternalError
	}

	return nil
}

func (s signerRepo) GetSingerIDsByKeys(ctx context.Context, keys [][]byte) ([]int, error) {
	signerIDs, err := s.client.
		GetConnection().
		Signer.
		Query().
		Where(signer.KeyIn(keys...)).
		IDs(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch signer IDs from database")
		return []int{}, consts.ErrInternalError
	}

	return signerIDs, nil
}

func NewSigner(client database.Database) repository.Signer {
	return &signerRepo{
		client: client,
	}
}
