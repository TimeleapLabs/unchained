// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KenshiTech/unchained/internal/ent/correctnessreport"
	"github.com/KenshiTech/unchained/internal/ent/predicate"
	"github.com/KenshiTech/unchained/internal/ent/signer"
)

// CorrectnessReportQuery is the builder for querying CorrectnessReport entities.
type CorrectnessReportQuery struct {
	config
	ctx              *QueryContext
	order            []correctnessreport.OrderOption
	inters           []Interceptor
	predicates       []predicate.CorrectnessReport
	withSigners      *SignerQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*CorrectnessReport) error
	withNamedSigners map[string]*SignerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CorrectnessReportQuery builder.
func (crq *CorrectnessReportQuery) Where(ps ...predicate.CorrectnessReport) *CorrectnessReportQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *CorrectnessReportQuery) Limit(limit int) *CorrectnessReportQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *CorrectnessReportQuery) Offset(offset int) *CorrectnessReportQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *CorrectnessReportQuery) Unique(unique bool) *CorrectnessReportQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *CorrectnessReportQuery) Order(o ...correctnessreport.OrderOption) *CorrectnessReportQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QuerySigners chains the current query on the "signers" edge.
func (crq *CorrectnessReportQuery) QuerySigners() *SignerQuery {
	query := (&SignerClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(correctnessreport.Table, correctnessreport.FieldID, selector),
			sqlgraph.To(signer.Table, signer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, correctnessreport.SignersTable, correctnessreport.SignersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CorrectnessReport entity from the query.
// Returns a *NotFoundError when no CorrectnessReport was found.
func (crq *CorrectnessReportQuery) First(ctx context.Context) (*CorrectnessReport, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{correctnessreport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *CorrectnessReportQuery) FirstX(ctx context.Context) *CorrectnessReport {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CorrectnessReport ID from the query.
// Returns a *NotFoundError when no CorrectnessReport ID was found.
func (crq *CorrectnessReportQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{correctnessreport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *CorrectnessReportQuery) FirstIDX(ctx context.Context) int {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CorrectnessReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CorrectnessReport entity is found.
// Returns a *NotFoundError when no CorrectnessReport entities are found.
func (crq *CorrectnessReportQuery) Only(ctx context.Context) (*CorrectnessReport, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{correctnessreport.Label}
	default:
		return nil, &NotSingularError{correctnessreport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *CorrectnessReportQuery) OnlyX(ctx context.Context) *CorrectnessReport {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CorrectnessReport ID in the query.
// Returns a *NotSingularError when more than one CorrectnessReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *CorrectnessReportQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{correctnessreport.Label}
	default:
		err = &NotSingularError{correctnessreport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *CorrectnessReportQuery) OnlyIDX(ctx context.Context) int {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CorrectnessReports.
func (crq *CorrectnessReportQuery) All(ctx context.Context) ([]*CorrectnessReport, error) {
	ctx = setContextOp(ctx, crq.ctx, "All")
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CorrectnessReport, *CorrectnessReportQuery]()
	return withInterceptors[[]*CorrectnessReport](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *CorrectnessReportQuery) AllX(ctx context.Context) []*CorrectnessReport {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CorrectnessReport IDs.
func (crq *CorrectnessReportQuery) IDs(ctx context.Context) (ids []int, err error) {
	if crq.ctx.Unique == nil && crq.path != nil {
		crq.Unique(true)
	}
	ctx = setContextOp(ctx, crq.ctx, "IDs")
	if err = crq.Select(correctnessreport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *CorrectnessReportQuery) IDsX(ctx context.Context) []int {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *CorrectnessReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, "Count")
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*CorrectnessReportQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *CorrectnessReportQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *CorrectnessReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, "Exist")
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *CorrectnessReportQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CorrectnessReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *CorrectnessReportQuery) Clone() *CorrectnessReportQuery {
	if crq == nil {
		return nil
	}
	return &CorrectnessReportQuery{
		config:      crq.config,
		ctx:         crq.ctx.Clone(),
		order:       append([]correctnessreport.OrderOption{}, crq.order...),
		inters:      append([]Interceptor{}, crq.inters...),
		predicates:  append([]predicate.CorrectnessReport{}, crq.predicates...),
		withSigners: crq.withSigners.Clone(),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// WithSigners tells the query-builder to eager-load the nodes that are connected to
// the "signers" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *CorrectnessReportQuery) WithSigners(opts ...func(*SignerQuery)) *CorrectnessReportQuery {
	query := (&SignerClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withSigners = query
	return crq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SignersCount uint64 `json:"signersCount,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CorrectnessReport.Query().
//		GroupBy(correctnessreport.FieldSignersCount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *CorrectnessReportQuery) GroupBy(field string, fields ...string) *CorrectnessReportGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CorrectnessReportGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = correctnessreport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SignersCount uint64 `json:"signersCount,omitempty"`
//	}
//
//	client.CorrectnessReport.Query().
//		Select(correctnessreport.FieldSignersCount).
//		Scan(ctx, &v)
func (crq *CorrectnessReportQuery) Select(fields ...string) *CorrectnessReportSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &CorrectnessReportSelect{CorrectnessReportQuery: crq}
	sbuild.label = correctnessreport.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CorrectnessReportSelect configured with the given aggregations.
func (crq *CorrectnessReportQuery) Aggregate(fns ...AggregateFunc) *CorrectnessReportSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *CorrectnessReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !correctnessreport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *CorrectnessReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CorrectnessReport, error) {
	var (
		nodes       = []*CorrectnessReport{}
		_spec       = crq.querySpec()
		loadedTypes = [1]bool{
			crq.withSigners != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CorrectnessReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CorrectnessReport{config: crq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(crq.modifiers) > 0 {
		_spec.Modifiers = crq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := crq.withSigners; query != nil {
		if err := crq.loadSigners(ctx, query, nodes,
			func(n *CorrectnessReport) { n.Edges.Signers = []*Signer{} },
			func(n *CorrectnessReport, e *Signer) { n.Edges.Signers = append(n.Edges.Signers, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range crq.withNamedSigners {
		if err := crq.loadSigners(ctx, query, nodes,
			func(n *CorrectnessReport) { n.appendNamedSigners(name) },
			func(n *CorrectnessReport, e *Signer) { n.appendNamedSigners(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range crq.loadTotal {
		if err := crq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (crq *CorrectnessReportQuery) loadSigners(ctx context.Context, query *SignerQuery, nodes []*CorrectnessReport, init func(*CorrectnessReport), assign func(*CorrectnessReport, *Signer)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*CorrectnessReport)
	nids := make(map[int]map[*CorrectnessReport]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(correctnessreport.SignersTable)
		s.Join(joinT).On(s.C(signer.FieldID), joinT.C(correctnessreport.SignersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(correctnessreport.SignersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(correctnessreport.SignersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*CorrectnessReport]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Signer](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "signers" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (crq *CorrectnessReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	if len(crq.modifiers) > 0 {
		_spec.Modifiers = crq.modifiers
	}
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *CorrectnessReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(correctnessreport.Table, correctnessreport.Columns, sqlgraph.NewFieldSpec(correctnessreport.FieldID, field.TypeInt))
	_spec.From = crq.sql
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crq.path != nil {
		_spec.Unique = true
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, correctnessreport.FieldID)
		for i := range fields {
			if fields[i] != correctnessreport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *CorrectnessReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(correctnessreport.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = correctnessreport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSigners tells the query-builder to eager-load the nodes that are connected to the "signers"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (crq *CorrectnessReportQuery) WithNamedSigners(name string, opts ...func(*SignerQuery)) *CorrectnessReportQuery {
	query := (&SignerClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if crq.withNamedSigners == nil {
		crq.withNamedSigners = make(map[string]*SignerQuery)
	}
	crq.withNamedSigners[name] = query
	return crq
}

// CorrectnessReportGroupBy is the group-by builder for CorrectnessReport entities.
type CorrectnessReportGroupBy struct {
	selector
	build *CorrectnessReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *CorrectnessReportGroupBy) Aggregate(fns ...AggregateFunc) *CorrectnessReportGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *CorrectnessReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, "GroupBy")
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CorrectnessReportQuery, *CorrectnessReportGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *CorrectnessReportGroupBy) sqlScan(ctx context.Context, root *CorrectnessReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CorrectnessReportSelect is the builder for selecting fields of CorrectnessReport entities.
type CorrectnessReportSelect struct {
	*CorrectnessReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *CorrectnessReportSelect) Aggregate(fns ...AggregateFunc) *CorrectnessReportSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *CorrectnessReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, "Select")
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CorrectnessReportQuery, *CorrectnessReportSelect](ctx, crs.CorrectnessReportQuery, crs, crs.inters, v)
}

func (crs *CorrectnessReportSelect) sqlScan(ctx context.Context, root *CorrectnessReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
