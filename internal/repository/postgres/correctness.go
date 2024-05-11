package postgres

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/correctnessreport"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type CorrectnessRepo struct {
	client database.Database
}

func (c CorrectnessRepo) Find(ctx context.Context, hash []byte, topic []byte, timestamp uint64) ([]*ent.CorrectnessReport, error) {
	currentRecords, err := c.client.
		GetConnection().
		CorrectnessReport.
		Query().
		Where(correctnessreport.And(
			correctnessreport.Hash(hash),
			correctnessreport.Topic(topic),
			correctnessreport.Timestamp(timestamp),
		)).
		All(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch correctness reports from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (c CorrectnessRepo) Upsert(ctx context.Context, data model.Correctness) error {
	err := c.client.
		GetConnection().
		CorrectnessReport.
		Create().
		SetCorrect(data.Correct).
		SetSignersCount(data.SignersCount).
		SetSignature(data.Signature).
		SetHash(data.Hash).
		SetTimestamp(data.Timestamp).
		SetTopic(data.Topic[:]).
		SetConsensus(data.Consensus).
		SetVoted(&helpers.BigInt{Int: data.Voted}).
		AddSignerIDs(data.SignerIDs...).
		OnConflictColumns("topic", "hash").
		UpdateNewValues().
		Update(func(u *ent.CorrectnessReportUpsert) {
			u.Add("voted", 1)
		}).
		Exec(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert correctness report in database")
		return consts.ErrInternalError
	}

	return nil
}

func NewCorrectness(client database.Database) repository.CorrectnessReport {
	return &CorrectnessRepo{
		client: client,
	}
}
