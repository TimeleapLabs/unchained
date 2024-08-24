// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/TimeleapLabs/unchained/internal/model"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/ent/helpers"
	"github.com/TimeleapLabs/unchained/internal/ent/signer"
)

// EventLogCreate is the builder for creating a EventLog entity.
type EventLogCreate struct {
	config
	mutation *EventLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBlock sets the "block" field.
func (elc *EventLogCreate) SetBlock(u uint64) *EventLogCreate {
	elc.mutation.SetBlock(u)
	return elc
}

// SetSignersCount sets the "signersCount" field.
func (elc *EventLogCreate) SetSignersCount(u uint64) *EventLogCreate {
	elc.mutation.SetSignersCount(u)
	return elc
}

// SetSignature sets the "signature" field.
func (elc *EventLogCreate) SetSignature(b []byte) *EventLogCreate {
	elc.mutation.SetSignature(b)
	return elc
}

// SetAddress sets the "address" field.
func (elc *EventLogCreate) SetAddress(s string) *EventLogCreate {
	elc.mutation.SetAddress(s)
	return elc
}

// SetChain sets the "chain" field.
func (elc *EventLogCreate) SetChain(s string) *EventLogCreate {
	elc.mutation.SetChain(s)
	return elc
}

// SetIndex sets the "index" field.
func (elc *EventLogCreate) SetIndex(u uint64) *EventLogCreate {
	elc.mutation.SetIndex(u)
	return elc
}

// SetEvent sets the "event" field.
func (elc *EventLogCreate) SetEvent(s string) *EventLogCreate {
	elc.mutation.SetEvent(s)
	return elc
}

// SetTransaction sets the "transaction" field.
func (elc *EventLogCreate) SetTransaction(b []byte) *EventLogCreate {
	elc.mutation.SetTransaction(b)
	return elc
}

// SetArgs sets the "args" field.
func (elc *EventLogCreate) SetArgs(dla []model.EventLogArg) *EventLogCreate {
	elc.mutation.SetArgs(dla)
	return elc
}

// SetConsensus sets the "consensus" field.
func (elc *EventLogCreate) SetConsensus(b bool) *EventLogCreate {
	elc.mutation.SetConsensus(b)
	return elc
}

// SetNillableConsensus sets the "consensus" field if the given value is not nil.
func (elc *EventLogCreate) SetNillableConsensus(b *bool) *EventLogCreate {
	if b != nil {
		elc.SetConsensus(*b)
	}
	return elc
}

// SetVoted sets the "voted" field.
func (elc *EventLogCreate) SetVoted(hi *helpers.BigInt) *EventLogCreate {
	elc.mutation.SetVoted(hi)
	return elc
}

// AddSignerIDs adds the "signers" edge to the Signer entity by IDs.
func (elc *EventLogCreate) AddSignerIDs(ids ...int) *EventLogCreate {
	elc.mutation.AddSignerIDs(ids...)
	return elc
}

// AddSigners adds the "signers" edges to the Signer entity.
func (elc *EventLogCreate) AddSigners(s ...*Signer) *EventLogCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return elc.AddSignerIDs(ids...)
}

// Mutation returns the EventLogMutation object of the builder.
func (elc *EventLogCreate) Mutation() *EventLogMutation {
	return elc.mutation
}

// Save creates the EventLog in the database.
func (elc *EventLogCreate) Save(ctx context.Context) (*EventLog, error) {
	elc.defaults()
	return withHooks(ctx, elc.sqlSave, elc.mutation, elc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (elc *EventLogCreate) SaveX(ctx context.Context) *EventLog {
	v, err := elc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elc *EventLogCreate) Exec(ctx context.Context) error {
	_, err := elc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elc *EventLogCreate) ExecX(ctx context.Context) {
	if err := elc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (elc *EventLogCreate) defaults() {
	if _, ok := elc.mutation.Consensus(); !ok {
		v := eventlog.DefaultConsensus
		elc.mutation.SetConsensus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (elc *EventLogCreate) check() error {
	if _, ok := elc.mutation.Block(); !ok {
		return &ValidationError{Name: "block", err: errors.New(`ent: missing required field "EventLog.block"`)}
	}
	if _, ok := elc.mutation.SignersCount(); !ok {
		return &ValidationError{Name: "signersCount", err: errors.New(`ent: missing required field "EventLog.signersCount"`)}
	}
	if _, ok := elc.mutation.Signature(); !ok {
		return &ValidationError{Name: "signature", err: errors.New(`ent: missing required field "EventLog.signature"`)}
	}
	if v, ok := elc.mutation.Signature(); ok {
		if err := eventlog.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "EventLog.signature": %w`, err)}
		}
	}
	if _, ok := elc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "EventLog.address"`)}
	}
	if _, ok := elc.mutation.Chain(); !ok {
		return &ValidationError{Name: "chain", err: errors.New(`ent: missing required field "EventLog.chain"`)}
	}
	if _, ok := elc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "EventLog.index"`)}
	}
	if _, ok := elc.mutation.Event(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required field "EventLog.event"`)}
	}
	if _, ok := elc.mutation.Transaction(); !ok {
		return &ValidationError{Name: "transaction", err: errors.New(`ent: missing required field "EventLog.transaction"`)}
	}
	if v, ok := elc.mutation.Transaction(); ok {
		if err := eventlog.TransactionValidator(v); err != nil {
			return &ValidationError{Name: "transaction", err: fmt.Errorf(`ent: validator failed for field "EventLog.transaction": %w`, err)}
		}
	}
	if _, ok := elc.mutation.Args(); !ok {
		return &ValidationError{Name: "args", err: errors.New(`ent: missing required field "EventLog.args"`)}
	}
	if _, ok := elc.mutation.Consensus(); !ok {
		return &ValidationError{Name: "consensus", err: errors.New(`ent: missing required field "EventLog.consensus"`)}
	}
	if _, ok := elc.mutation.Voted(); !ok {
		return &ValidationError{Name: "voted", err: errors.New(`ent: missing required field "EventLog.voted"`)}
	}
	if len(elc.mutation.SignersIDs()) == 0 {
		return &ValidationError{Name: "signers", err: errors.New(`ent: missing required edge "EventLog.signers"`)}
	}
	return nil
}

func (elc *EventLogCreate) sqlSave(ctx context.Context) (*EventLog, error) {
	if err := elc.check(); err != nil {
		return nil, err
	}
	_node, _spec := elc.createSpec()
	if err := sqlgraph.CreateNode(ctx, elc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	elc.mutation.id = &_node.ID
	elc.mutation.done = true
	return _node, nil
}

func (elc *EventLogCreate) createSpec() (*EventLog, *sqlgraph.CreateSpec) {
	var (
		_node = &EventLog{config: elc.config}
		_spec = sqlgraph.NewCreateSpec(eventlog.Table, sqlgraph.NewFieldSpec(eventlog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = elc.conflict
	if value, ok := elc.mutation.Block(); ok {
		_spec.SetField(eventlog.FieldBlock, field.TypeUint64, value)
		_node.Block = value
	}
	if value, ok := elc.mutation.SignersCount(); ok {
		_spec.SetField(eventlog.FieldSignersCount, field.TypeUint64, value)
		_node.SignersCount = value
	}
	if value, ok := elc.mutation.Signature(); ok {
		_spec.SetField(eventlog.FieldSignature, field.TypeBytes, value)
		_node.Signature = value
	}
	if value, ok := elc.mutation.Address(); ok {
		_spec.SetField(eventlog.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := elc.mutation.Chain(); ok {
		_spec.SetField(eventlog.FieldChain, field.TypeString, value)
		_node.Chain = value
	}
	if value, ok := elc.mutation.Index(); ok {
		_spec.SetField(eventlog.FieldIndex, field.TypeUint64, value)
		_node.Index = value
	}
	if value, ok := elc.mutation.Event(); ok {
		_spec.SetField(eventlog.FieldEvent, field.TypeString, value)
		_node.Event = value
	}
	if value, ok := elc.mutation.Transaction(); ok {
		_spec.SetField(eventlog.FieldTransaction, field.TypeBytes, value)
		_node.Transaction = value
	}
	if value, ok := elc.mutation.Args(); ok {
		_spec.SetField(eventlog.FieldArgs, field.TypeJSON, value)
		_node.Args = value
	}
	if value, ok := elc.mutation.Consensus(); ok {
		_spec.SetField(eventlog.FieldConsensus, field.TypeBool, value)
		_node.Consensus = value
	}
	if value, ok := elc.mutation.Voted(); ok {
		_spec.SetField(eventlog.FieldVoted, field.TypeUint, value)
		_node.Voted = value
	}
	if nodes := elc.mutation.SignersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   eventlog.SignersTable,
			Columns: eventlog.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EventLog.Create().
//		SetBlock(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventLogUpsert) {
//			SetBlock(v+v).
//		}).
//		Exec(ctx)
func (elc *EventLogCreate) OnConflict(opts ...sql.ConflictOption) *EventLogUpsertOne {
	elc.conflict = opts
	return &EventLogUpsertOne{
		create: elc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EventLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (elc *EventLogCreate) OnConflictColumns(columns ...string) *EventLogUpsertOne {
	elc.conflict = append(elc.conflict, sql.ConflictColumns(columns...))
	return &EventLogUpsertOne{
		create: elc,
	}
}

type (
	// EventLogUpsertOne is the builder for "upsert"-ing
	//  one EventLog node.
	EventLogUpsertOne struct {
		create *EventLogCreate
	}

	// EventLogUpsert is the "OnConflict" setter.
	EventLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetBlock sets the "block" field.
func (u *EventLogUpsert) SetBlock(v uint64) *EventLogUpsert {
	u.Set(eventlog.FieldBlock, v)
	return u
}

// UpdateBlock sets the "block" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateBlock() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldBlock)
	return u
}

// AddBlock adds v to the "block" field.
func (u *EventLogUpsert) AddBlock(v uint64) *EventLogUpsert {
	u.Add(eventlog.FieldBlock, v)
	return u
}

// SetSignersCount sets the "signersCount" field.
func (u *EventLogUpsert) SetSignersCount(v uint64) *EventLogUpsert {
	u.Set(eventlog.FieldSignersCount, v)
	return u
}

// UpdateSignersCount sets the "signersCount" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateSignersCount() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldSignersCount)
	return u
}

// AddSignersCount adds v to the "signersCount" field.
func (u *EventLogUpsert) AddSignersCount(v uint64) *EventLogUpsert {
	u.Add(eventlog.FieldSignersCount, v)
	return u
}

// SetSignature sets the "signature" field.
func (u *EventLogUpsert) SetSignature(v []byte) *EventLogUpsert {
	u.Set(eventlog.FieldSignature, v)
	return u
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateSignature() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldSignature)
	return u
}

// SetAddress sets the "address" field.
func (u *EventLogUpsert) SetAddress(v string) *EventLogUpsert {
	u.Set(eventlog.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateAddress() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldAddress)
	return u
}

// SetChain sets the "chain" field.
func (u *EventLogUpsert) SetChain(v string) *EventLogUpsert {
	u.Set(eventlog.FieldChain, v)
	return u
}

// UpdateChain sets the "chain" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateChain() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldChain)
	return u
}

// SetIndex sets the "index" field.
func (u *EventLogUpsert) SetIndex(v uint64) *EventLogUpsert {
	u.Set(eventlog.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateIndex() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *EventLogUpsert) AddIndex(v uint64) *EventLogUpsert {
	u.Add(eventlog.FieldIndex, v)
	return u
}

// SetEvent sets the "event" field.
func (u *EventLogUpsert) SetEvent(v string) *EventLogUpsert {
	u.Set(eventlog.FieldEvent, v)
	return u
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateEvent() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldEvent)
	return u
}

// SetTransaction sets the "transaction" field.
func (u *EventLogUpsert) SetTransaction(v []byte) *EventLogUpsert {
	u.Set(eventlog.FieldTransaction, v)
	return u
}

// UpdateTransaction sets the "transaction" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateTransaction() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldTransaction)
	return u
}

// SetArgs sets the "args" field.
func (u *EventLogUpsert) SetArgs(v []model.EventLogArg) *EventLogUpsert {
	u.Set(eventlog.FieldArgs, v)
	return u
}

// UpdateArgs sets the "args" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateArgs() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldArgs)
	return u
}

// SetConsensus sets the "consensus" field.
func (u *EventLogUpsert) SetConsensus(v bool) *EventLogUpsert {
	u.Set(eventlog.FieldConsensus, v)
	return u
}

// UpdateConsensus sets the "consensus" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateConsensus() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldConsensus)
	return u
}

// SetVoted sets the "voted" field.
func (u *EventLogUpsert) SetVoted(v *helpers.BigInt) *EventLogUpsert {
	u.Set(eventlog.FieldVoted, v)
	return u
}

// UpdateVoted sets the "voted" field to the value that was provided on create.
func (u *EventLogUpsert) UpdateVoted() *EventLogUpsert {
	u.SetExcluded(eventlog.FieldVoted)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.EventLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EventLogUpsertOne) UpdateNewValues() *EventLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EventLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EventLogUpsertOne) Ignore() *EventLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventLogUpsertOne) DoNothing() *EventLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventLogCreate.OnConflict
// documentation for more info.
func (u *EventLogUpsertOne) Update(set func(*EventLogUpsert)) *EventLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlock sets the "block" field.
func (u *EventLogUpsertOne) SetBlock(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetBlock(v)
	})
}

// AddBlock adds v to the "block" field.
func (u *EventLogUpsertOne) AddBlock(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.AddBlock(v)
	})
}

// UpdateBlock sets the "block" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateBlock() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateBlock()
	})
}

// SetSignersCount sets the "signersCount" field.
func (u *EventLogUpsertOne) SetSignersCount(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetSignersCount(v)
	})
}

// AddSignersCount adds v to the "signersCount" field.
func (u *EventLogUpsertOne) AddSignersCount(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.AddSignersCount(v)
	})
}

// UpdateSignersCount sets the "signersCount" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateSignersCount() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateSignersCount()
	})
}

// SetSignature sets the "signature" field.
func (u *EventLogUpsertOne) SetSignature(v []byte) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateSignature() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateSignature()
	})
}

// SetAddress sets the "address" field.
func (u *EventLogUpsertOne) SetAddress(v string) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateAddress() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateAddress()
	})
}

// SetChain sets the "chain" field.
func (u *EventLogUpsertOne) SetChain(v string) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetChain(v)
	})
}

// UpdateChain sets the "chain" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateChain() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateChain()
	})
}

// SetIndex sets the "index" field.
func (u *EventLogUpsertOne) SetIndex(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *EventLogUpsertOne) AddIndex(v uint64) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateIndex() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateIndex()
	})
}

// SetEvent sets the "event" field.
func (u *EventLogUpsertOne) SetEvent(v string) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetEvent(v)
	})
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateEvent() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateEvent()
	})
}

// SetTransaction sets the "transaction" field.
func (u *EventLogUpsertOne) SetTransaction(v []byte) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetTransaction(v)
	})
}

// UpdateTransaction sets the "transaction" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateTransaction() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateTransaction()
	})
}

// SetArgs sets the "args" field.
func (u *EventLogUpsertOne) SetArgs(v []model.EventLogArg) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetArgs(v)
	})
}

// UpdateArgs sets the "args" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateArgs() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateArgs()
	})
}

// SetConsensus sets the "consensus" field.
func (u *EventLogUpsertOne) SetConsensus(v bool) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetConsensus(v)
	})
}

// UpdateConsensus sets the "consensus" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateConsensus() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateConsensus()
	})
}

// SetVoted sets the "voted" field.
func (u *EventLogUpsertOne) SetVoted(v *helpers.BigInt) *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.SetVoted(v)
	})
}

// UpdateVoted sets the "voted" field to the value that was provided on create.
func (u *EventLogUpsertOne) UpdateVoted() *EventLogUpsertOne {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateVoted()
	})
}

// Exec executes the query.
func (u *EventLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventLogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventLogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventLogCreateBulk is the builder for creating many EventLog entities in bulk.
type EventLogCreateBulk struct {
	config
	err      error
	builders []*EventLogCreate
	conflict []sql.ConflictOption
}

// Save creates the EventLog entities in the database.
func (elcb *EventLogCreateBulk) Save(ctx context.Context) ([]*EventLog, error) {
	if elcb.err != nil {
		return nil, elcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(elcb.builders))
	nodes := make([]*EventLog, len(elcb.builders))
	mutators := make([]Mutator, len(elcb.builders))
	for i := range elcb.builders {
		func(i int, root context.Context) {
			builder := elcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, elcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = elcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, elcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, elcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (elcb *EventLogCreateBulk) SaveX(ctx context.Context) []*EventLog {
	v, err := elcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elcb *EventLogCreateBulk) Exec(ctx context.Context) error {
	_, err := elcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elcb *EventLogCreateBulk) ExecX(ctx context.Context) {
	if err := elcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EventLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventLogUpsert) {
//			SetBlock(v+v).
//		}).
//		Exec(ctx)
func (elcb *EventLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventLogUpsertBulk {
	elcb.conflict = opts
	return &EventLogUpsertBulk{
		create: elcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EventLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (elcb *EventLogCreateBulk) OnConflictColumns(columns ...string) *EventLogUpsertBulk {
	elcb.conflict = append(elcb.conflict, sql.ConflictColumns(columns...))
	return &EventLogUpsertBulk{
		create: elcb,
	}
}

// EventLogUpsertBulk is the builder for "upsert"-ing
// a bulk of EventLog nodes.
type EventLogUpsertBulk struct {
	create *EventLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EventLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EventLogUpsertBulk) UpdateNewValues() *EventLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EventLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EventLogUpsertBulk) Ignore() *EventLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventLogUpsertBulk) DoNothing() *EventLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventLogCreateBulk.OnConflict
// documentation for more info.
func (u *EventLogUpsertBulk) Update(set func(*EventLogUpsert)) *EventLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlock sets the "block" field.
func (u *EventLogUpsertBulk) SetBlock(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetBlock(v)
	})
}

// AddBlock adds v to the "block" field.
func (u *EventLogUpsertBulk) AddBlock(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.AddBlock(v)
	})
}

// UpdateBlock sets the "block" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateBlock() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateBlock()
	})
}

// SetSignersCount sets the "signersCount" field.
func (u *EventLogUpsertBulk) SetSignersCount(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetSignersCount(v)
	})
}

// AddSignersCount adds v to the "signersCount" field.
func (u *EventLogUpsertBulk) AddSignersCount(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.AddSignersCount(v)
	})
}

// UpdateSignersCount sets the "signersCount" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateSignersCount() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateSignersCount()
	})
}

// SetSignature sets the "signature" field.
func (u *EventLogUpsertBulk) SetSignature(v []byte) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateSignature() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateSignature()
	})
}

// SetAddress sets the "address" field.
func (u *EventLogUpsertBulk) SetAddress(v string) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateAddress() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateAddress()
	})
}

// SetChain sets the "chain" field.
func (u *EventLogUpsertBulk) SetChain(v string) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetChain(v)
	})
}

// UpdateChain sets the "chain" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateChain() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateChain()
	})
}

// SetIndex sets the "index" field.
func (u *EventLogUpsertBulk) SetIndex(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *EventLogUpsertBulk) AddIndex(v uint64) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateIndex() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateIndex()
	})
}

// SetEvent sets the "event" field.
func (u *EventLogUpsertBulk) SetEvent(v string) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetEvent(v)
	})
}

// UpdateEvent sets the "event" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateEvent() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateEvent()
	})
}

// SetTransaction sets the "transaction" field.
func (u *EventLogUpsertBulk) SetTransaction(v []byte) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetTransaction(v)
	})
}

// UpdateTransaction sets the "transaction" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateTransaction() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateTransaction()
	})
}

// SetArgs sets the "args" field.
func (u *EventLogUpsertBulk) SetArgs(v []model.EventLogArg) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetArgs(v)
	})
}

// UpdateArgs sets the "args" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateArgs() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateArgs()
	})
}

// SetConsensus sets the "consensus" field.
func (u *EventLogUpsertBulk) SetConsensus(v bool) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetConsensus(v)
	})
}

// UpdateConsensus sets the "consensus" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateConsensus() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateConsensus()
	})
}

// SetVoted sets the "voted" field.
func (u *EventLogUpsertBulk) SetVoted(v *helpers.BigInt) *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.SetVoted(v)
	})
}

// UpdateVoted sets the "voted" field to the value that was provided on create.
func (u *EventLogUpsertBulk) UpdateVoted() *EventLogUpsertBulk {
	return u.Update(func(s *EventLogUpsert) {
		s.UpdateVoted()
	})
}

// Exec executes the query.
func (u *EventLogUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}