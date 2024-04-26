package postgres

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/ent"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type EventLogRepo struct {
	client database.Database
}

func (r EventLogRepo) Find(ctx context.Context, block uint64, hash []byte, index uint64) ([]*ent.EventLog, error) {
	currentRecords, err := r.client.
		GetConnection().
		EventLog.
		Query().
		Where(
			eventlog.Block(block),
			eventlog.TransactionEQ(hash),
			eventlog.IndexEQ(index),
		).
		WithSigners().
		All(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch event log records from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (r EventLogRepo) Upsert(ctx context.Context, data model.EventLog) error {
	err := r.client.
		GetConnection().
		EventLog.
		Create().
		SetBlock(data.Block).
		SetChain(data.Chain).
		SetAddress(data.Address).
		SetEvent(data.Event).
		SetIndex(data.LogIndex).
		SetTransaction(data.TxHash[:]).
		SetSignersCount(data.SignersCount).
		SetSignature(data.Signature).
		SetArgs(data.Args).
		SetConsensus(data.Consensus).
		SetVoted(data.Voted).
		AddSignerIDs(data.SignerIDs...).
		OnConflictColumns("block", "transaction", "index").
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert event log record to database")
		return consts.ErrInternalError
	}

	return nil
}

func NewEventLog(client database.Database) repository.EventLog {
	return &EventLogRepo{
		client: client,
	}
}
