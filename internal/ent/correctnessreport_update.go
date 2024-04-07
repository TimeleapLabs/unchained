// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KenshiTech/unchained/internal/ent/correctnessreport"
	"github.com/KenshiTech/unchained/internal/ent/predicate"
	"github.com/KenshiTech/unchained/internal/ent/signer"
)

// CorrectnessReportUpdate is the builder for updating CorrectnessReport entities.
type CorrectnessReportUpdate struct {
	config
	hooks    []Hook
	mutation *CorrectnessReportMutation
}

// Where appends a list predicates to the CorrectnessReportUpdate builder.
func (cru *CorrectnessReportUpdate) Where(ps ...predicate.CorrectnessReport) *CorrectnessReportUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetSignersCount sets the "signersCount" field.
func (cru *CorrectnessReportUpdate) SetSignersCount(u uint64) *CorrectnessReportUpdate {
	cru.mutation.ResetSignersCount()
	cru.mutation.SetSignersCount(u)
	return cru
}

// SetNillableSignersCount sets the "signersCount" field if the given value is not nil.
func (cru *CorrectnessReportUpdate) SetNillableSignersCount(u *uint64) *CorrectnessReportUpdate {
	if u != nil {
		cru.SetSignersCount(*u)
	}
	return cru
}

// AddSignersCount adds u to the "signersCount" field.
func (cru *CorrectnessReportUpdate) AddSignersCount(u int64) *CorrectnessReportUpdate {
	cru.mutation.AddSignersCount(u)
	return cru
}

// SetTimestamp sets the "timestamp" field.
func (cru *CorrectnessReportUpdate) SetTimestamp(u uint64) *CorrectnessReportUpdate {
	cru.mutation.ResetTimestamp()
	cru.mutation.SetTimestamp(u)
	return cru
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (cru *CorrectnessReportUpdate) SetNillableTimestamp(u *uint64) *CorrectnessReportUpdate {
	if u != nil {
		cru.SetTimestamp(*u)
	}
	return cru
}

// AddTimestamp adds u to the "timestamp" field.
func (cru *CorrectnessReportUpdate) AddTimestamp(u int64) *CorrectnessReportUpdate {
	cru.mutation.AddTimestamp(u)
	return cru
}

// SetSignature sets the "signature" field.
func (cru *CorrectnessReportUpdate) SetSignature(b []byte) *CorrectnessReportUpdate {
	cru.mutation.SetSignature(b)
	return cru
}

// SetHash sets the "hash" field.
func (cru *CorrectnessReportUpdate) SetHash(b []byte) *CorrectnessReportUpdate {
	cru.mutation.SetHash(b)
	return cru
}

// SetTopic sets the "topic" field.
func (cru *CorrectnessReportUpdate) SetTopic(b []byte) *CorrectnessReportUpdate {
	cru.mutation.SetTopic(b)
	return cru
}

// SetCorrect sets the "correct" field.
func (cru *CorrectnessReportUpdate) SetCorrect(b bool) *CorrectnessReportUpdate {
	cru.mutation.SetCorrect(b)
	return cru
}

// SetNillableCorrect sets the "correct" field if the given value is not nil.
func (cru *CorrectnessReportUpdate) SetNillableCorrect(b *bool) *CorrectnessReportUpdate {
	if b != nil {
		cru.SetCorrect(*b)
	}
	return cru
}

// AddSignerIDs adds the "signers" edge to the Signer entity by IDs.
func (cru *CorrectnessReportUpdate) AddSignerIDs(ids ...int) *CorrectnessReportUpdate {
	cru.mutation.AddSignerIDs(ids...)
	return cru
}

// AddSigners adds the "signers" edges to the Signer entity.
func (cru *CorrectnessReportUpdate) AddSigners(s ...*Signer) *CorrectnessReportUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cru.AddSignerIDs(ids...)
}

// Mutation returns the CorrectnessReportMutation object of the builder.
func (cru *CorrectnessReportUpdate) Mutation() *CorrectnessReportMutation {
	return cru.mutation
}

// ClearSigners clears all "signers" edges to the Signer entity.
func (cru *CorrectnessReportUpdate) ClearSigners() *CorrectnessReportUpdate {
	cru.mutation.ClearSigners()
	return cru
}

// RemoveSignerIDs removes the "signers" edge to Signer entities by IDs.
func (cru *CorrectnessReportUpdate) RemoveSignerIDs(ids ...int) *CorrectnessReportUpdate {
	cru.mutation.RemoveSignerIDs(ids...)
	return cru
}

// RemoveSigners removes "signers" edges to Signer entities.
func (cru *CorrectnessReportUpdate) RemoveSigners(s ...*Signer) *CorrectnessReportUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cru.RemoveSignerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CorrectnessReportUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CorrectnessReportUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CorrectnessReportUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CorrectnessReportUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CorrectnessReportUpdate) check() error {
	if v, ok := cru.mutation.Signature(); ok {
		if err := correctnessreport.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.signature": %w`, err)}
		}
	}
	if v, ok := cru.mutation.Hash(); ok {
		if err := correctnessreport.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.hash": %w`, err)}
		}
	}
	if v, ok := cru.mutation.Topic(); ok {
		if err := correctnessreport.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.topic": %w`, err)}
		}
	}
	return nil
}

func (cru *CorrectnessReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(correctnessreport.Table, correctnessreport.Columns, sqlgraph.NewFieldSpec(correctnessreport.FieldID, field.TypeInt))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.SignersCount(); ok {
		_spec.SetField(correctnessreport.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := cru.mutation.AddedSignersCount(); ok {
		_spec.AddField(correctnessreport.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := cru.mutation.Timestamp(); ok {
		_spec.SetField(correctnessreport.FieldTimestamp, field.TypeUint64, value)
	}
	if value, ok := cru.mutation.AddedTimestamp(); ok {
		_spec.AddField(correctnessreport.FieldTimestamp, field.TypeUint64, value)
	}
	if value, ok := cru.mutation.Signature(); ok {
		_spec.SetField(correctnessreport.FieldSignature, field.TypeBytes, value)
	}
	if value, ok := cru.mutation.Hash(); ok {
		_spec.SetField(correctnessreport.FieldHash, field.TypeBytes, value)
	}
	if value, ok := cru.mutation.Topic(); ok {
		_spec.SetField(correctnessreport.FieldTopic, field.TypeBytes, value)
	}
	if value, ok := cru.mutation.Correct(); ok {
		_spec.SetField(correctnessreport.FieldCorrect, field.TypeBool, value)
	}
	if cru.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.RemovedSignersIDs(); len(nodes) > 0 && !cru.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.SignersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{correctnessreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CorrectnessReportUpdateOne is the builder for updating a single CorrectnessReport entity.
type CorrectnessReportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CorrectnessReportMutation
}

// SetSignersCount sets the "signersCount" field.
func (cruo *CorrectnessReportUpdateOne) SetSignersCount(u uint64) *CorrectnessReportUpdateOne {
	cruo.mutation.ResetSignersCount()
	cruo.mutation.SetSignersCount(u)
	return cruo
}

// SetNillableSignersCount sets the "signersCount" field if the given value is not nil.
func (cruo *CorrectnessReportUpdateOne) SetNillableSignersCount(u *uint64) *CorrectnessReportUpdateOne {
	if u != nil {
		cruo.SetSignersCount(*u)
	}
	return cruo
}

// AddSignersCount adds u to the "signersCount" field.
func (cruo *CorrectnessReportUpdateOne) AddSignersCount(u int64) *CorrectnessReportUpdateOne {
	cruo.mutation.AddSignersCount(u)
	return cruo
}

// SetTimestamp sets the "timestamp" field.
func (cruo *CorrectnessReportUpdateOne) SetTimestamp(u uint64) *CorrectnessReportUpdateOne {
	cruo.mutation.ResetTimestamp()
	cruo.mutation.SetTimestamp(u)
	return cruo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (cruo *CorrectnessReportUpdateOne) SetNillableTimestamp(u *uint64) *CorrectnessReportUpdateOne {
	if u != nil {
		cruo.SetTimestamp(*u)
	}
	return cruo
}

// AddTimestamp adds u to the "timestamp" field.
func (cruo *CorrectnessReportUpdateOne) AddTimestamp(u int64) *CorrectnessReportUpdateOne {
	cruo.mutation.AddTimestamp(u)
	return cruo
}

// SetSignature sets the "signature" field.
func (cruo *CorrectnessReportUpdateOne) SetSignature(b []byte) *CorrectnessReportUpdateOne {
	cruo.mutation.SetSignature(b)
	return cruo
}

// SetHash sets the "hash" field.
func (cruo *CorrectnessReportUpdateOne) SetHash(b []byte) *CorrectnessReportUpdateOne {
	cruo.mutation.SetHash(b)
	return cruo
}

// SetTopic sets the "topic" field.
func (cruo *CorrectnessReportUpdateOne) SetTopic(b []byte) *CorrectnessReportUpdateOne {
	cruo.mutation.SetTopic(b)
	return cruo
}

// SetCorrect sets the "correct" field.
func (cruo *CorrectnessReportUpdateOne) SetCorrect(b bool) *CorrectnessReportUpdateOne {
	cruo.mutation.SetCorrect(b)
	return cruo
}

// SetNillableCorrect sets the "correct" field if the given value is not nil.
func (cruo *CorrectnessReportUpdateOne) SetNillableCorrect(b *bool) *CorrectnessReportUpdateOne {
	if b != nil {
		cruo.SetCorrect(*b)
	}
	return cruo
}

// AddSignerIDs adds the "signers" edge to the Signer entity by IDs.
func (cruo *CorrectnessReportUpdateOne) AddSignerIDs(ids ...int) *CorrectnessReportUpdateOne {
	cruo.mutation.AddSignerIDs(ids...)
	return cruo
}

// AddSigners adds the "signers" edges to the Signer entity.
func (cruo *CorrectnessReportUpdateOne) AddSigners(s ...*Signer) *CorrectnessReportUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cruo.AddSignerIDs(ids...)
}

// Mutation returns the CorrectnessReportMutation object of the builder.
func (cruo *CorrectnessReportUpdateOne) Mutation() *CorrectnessReportMutation {
	return cruo.mutation
}

// ClearSigners clears all "signers" edges to the Signer entity.
func (cruo *CorrectnessReportUpdateOne) ClearSigners() *CorrectnessReportUpdateOne {
	cruo.mutation.ClearSigners()
	return cruo
}

// RemoveSignerIDs removes the "signers" edge to Signer entities by IDs.
func (cruo *CorrectnessReportUpdateOne) RemoveSignerIDs(ids ...int) *CorrectnessReportUpdateOne {
	cruo.mutation.RemoveSignerIDs(ids...)
	return cruo
}

// RemoveSigners removes "signers" edges to Signer entities.
func (cruo *CorrectnessReportUpdateOne) RemoveSigners(s ...*Signer) *CorrectnessReportUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cruo.RemoveSignerIDs(ids...)
}

// Where appends a list predicates to the CorrectnessReportUpdate builder.
func (cruo *CorrectnessReportUpdateOne) Where(ps ...predicate.CorrectnessReport) *CorrectnessReportUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CorrectnessReportUpdateOne) Select(field string, fields ...string) *CorrectnessReportUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CorrectnessReport entity.
func (cruo *CorrectnessReportUpdateOne) Save(ctx context.Context) (*CorrectnessReport, error) {
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CorrectnessReportUpdateOne) SaveX(ctx context.Context) *CorrectnessReport {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CorrectnessReportUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CorrectnessReportUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CorrectnessReportUpdateOne) check() error {
	if v, ok := cruo.mutation.Signature(); ok {
		if err := correctnessreport.SignatureValidator(v); err != nil {
			return &ValidationError{Name: "signature", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.signature": %w`, err)}
		}
	}
	if v, ok := cruo.mutation.Hash(); ok {
		if err := correctnessreport.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.hash": %w`, err)}
		}
	}
	if v, ok := cruo.mutation.Topic(); ok {
		if err := correctnessreport.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf(`ent: validator failed for field "CorrectnessReport.topic": %w`, err)}
		}
	}
	return nil
}

func (cruo *CorrectnessReportUpdateOne) sqlSave(ctx context.Context) (_node *CorrectnessReport, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(correctnessreport.Table, correctnessreport.Columns, sqlgraph.NewFieldSpec(correctnessreport.FieldID, field.TypeInt))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CorrectnessReport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, correctnessreport.FieldID)
		for _, f := range fields {
			if !correctnessreport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != correctnessreport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.SignersCount(); ok {
		_spec.SetField(correctnessreport.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := cruo.mutation.AddedSignersCount(); ok {
		_spec.AddField(correctnessreport.FieldSignersCount, field.TypeUint64, value)
	}
	if value, ok := cruo.mutation.Timestamp(); ok {
		_spec.SetField(correctnessreport.FieldTimestamp, field.TypeUint64, value)
	}
	if value, ok := cruo.mutation.AddedTimestamp(); ok {
		_spec.AddField(correctnessreport.FieldTimestamp, field.TypeUint64, value)
	}
	if value, ok := cruo.mutation.Signature(); ok {
		_spec.SetField(correctnessreport.FieldSignature, field.TypeBytes, value)
	}
	if value, ok := cruo.mutation.Hash(); ok {
		_spec.SetField(correctnessreport.FieldHash, field.TypeBytes, value)
	}
	if value, ok := cruo.mutation.Topic(); ok {
		_spec.SetField(correctnessreport.FieldTopic, field.TypeBytes, value)
	}
	if value, ok := cruo.mutation.Correct(); ok {
		_spec.SetField(correctnessreport.FieldCorrect, field.TypeBool, value)
	}
	if cruo.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.RemovedSignersIDs(); len(nodes) > 0 && !cruo.mutation.SignersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.SignersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   correctnessreport.SignersTable,
			Columns: correctnessreport.SignersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CorrectnessReport{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{correctnessreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}