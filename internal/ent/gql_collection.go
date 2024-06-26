// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/TimeleapLabs/unchained/internal/ent/assetprice"
	"github.com/TimeleapLabs/unchained/internal/ent/correctnessreport"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/ent/signer"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ap *AssetPriceQuery) CollectFields(ctx context.Context, satisfies ...string) (*AssetPriceQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ap, nil
	}
	if err := ap.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ap, nil
}

func (ap *AssetPriceQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(assetprice.Columns))
		selectedFields = []string{assetprice.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "signers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SignerClient{config: ap.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, signerImplementors)...); err != nil {
				return err
			}
			ap.WithNamedSigners(alias, func(wq *SignerQuery) {
				*wq = *query
			})
		case "block":
			if _, ok := fieldSeen[assetprice.FieldBlock]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldBlock)
				fieldSeen[assetprice.FieldBlock] = struct{}{}
			}
		case "signerscount":
			if _, ok := fieldSeen[assetprice.FieldSignersCount]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldSignersCount)
				fieldSeen[assetprice.FieldSignersCount] = struct{}{}
			}
		case "price":
			if _, ok := fieldSeen[assetprice.FieldPrice]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldPrice)
				fieldSeen[assetprice.FieldPrice] = struct{}{}
			}
		case "signature":
			if _, ok := fieldSeen[assetprice.FieldSignature]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldSignature)
				fieldSeen[assetprice.FieldSignature] = struct{}{}
			}
		case "asset":
			if _, ok := fieldSeen[assetprice.FieldAsset]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldAsset)
				fieldSeen[assetprice.FieldAsset] = struct{}{}
			}
		case "chain":
			if _, ok := fieldSeen[assetprice.FieldChain]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldChain)
				fieldSeen[assetprice.FieldChain] = struct{}{}
			}
		case "pair":
			if _, ok := fieldSeen[assetprice.FieldPair]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldPair)
				fieldSeen[assetprice.FieldPair] = struct{}{}
			}
		case "consensus":
			if _, ok := fieldSeen[assetprice.FieldConsensus]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldConsensus)
				fieldSeen[assetprice.FieldConsensus] = struct{}{}
			}
		case "voted":
			if _, ok := fieldSeen[assetprice.FieldVoted]; !ok {
				selectedFields = append(selectedFields, assetprice.FieldVoted)
				fieldSeen[assetprice.FieldVoted] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ap.Select(selectedFields...)
	}
	return nil
}

type assetpricePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AssetPricePaginateOption
}

func newAssetPricePaginateArgs(rv map[string]any) *assetpricePaginateArgs {
	args := &assetpricePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &AssetPriceOrder{Field: &AssetPriceOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAssetPriceOrder(order))
			}
		case *AssetPriceOrder:
			if v != nil {
				args.opts = append(args.opts, WithAssetPriceOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*AssetPriceWhereInput); ok {
		args.opts = append(args.opts, WithAssetPriceFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (cr *CorrectnessReportQuery) CollectFields(ctx context.Context, satisfies ...string) (*CorrectnessReportQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return cr, nil
	}
	if err := cr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return cr, nil
}

func (cr *CorrectnessReportQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(correctnessreport.Columns))
		selectedFields = []string{correctnessreport.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "signers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SignerClient{config: cr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, signerImplementors)...); err != nil {
				return err
			}
			cr.WithNamedSigners(alias, func(wq *SignerQuery) {
				*wq = *query
			})
		case "signerscount":
			if _, ok := fieldSeen[correctnessreport.FieldSignersCount]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldSignersCount)
				fieldSeen[correctnessreport.FieldSignersCount] = struct{}{}
			}
		case "timestamp":
			if _, ok := fieldSeen[correctnessreport.FieldTimestamp]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldTimestamp)
				fieldSeen[correctnessreport.FieldTimestamp] = struct{}{}
			}
		case "signature":
			if _, ok := fieldSeen[correctnessreport.FieldSignature]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldSignature)
				fieldSeen[correctnessreport.FieldSignature] = struct{}{}
			}
		case "hash":
			if _, ok := fieldSeen[correctnessreport.FieldHash]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldHash)
				fieldSeen[correctnessreport.FieldHash] = struct{}{}
			}
		case "topic":
			if _, ok := fieldSeen[correctnessreport.FieldTopic]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldTopic)
				fieldSeen[correctnessreport.FieldTopic] = struct{}{}
			}
		case "correct":
			if _, ok := fieldSeen[correctnessreport.FieldCorrect]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldCorrect)
				fieldSeen[correctnessreport.FieldCorrect] = struct{}{}
			}
		case "consensus":
			if _, ok := fieldSeen[correctnessreport.FieldConsensus]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldConsensus)
				fieldSeen[correctnessreport.FieldConsensus] = struct{}{}
			}
		case "voted":
			if _, ok := fieldSeen[correctnessreport.FieldVoted]; !ok {
				selectedFields = append(selectedFields, correctnessreport.FieldVoted)
				fieldSeen[correctnessreport.FieldVoted] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		cr.Select(selectedFields...)
	}
	return nil
}

type correctnessreportPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CorrectnessReportPaginateOption
}

func newCorrectnessReportPaginateArgs(rv map[string]any) *correctnessreportPaginateArgs {
	args := &correctnessreportPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &CorrectnessReportOrder{Field: &CorrectnessReportOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCorrectnessReportOrder(order))
			}
		case *CorrectnessReportOrder:
			if v != nil {
				args.opts = append(args.opts, WithCorrectnessReportOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CorrectnessReportWhereInput); ok {
		args.opts = append(args.opts, WithCorrectnessReportFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (el *EventLogQuery) CollectFields(ctx context.Context, satisfies ...string) (*EventLogQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return el, nil
	}
	if err := el.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return el, nil
}

func (el *EventLogQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(eventlog.Columns))
		selectedFields = []string{eventlog.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "signers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SignerClient{config: el.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, signerImplementors)...); err != nil {
				return err
			}
			el.WithNamedSigners(alias, func(wq *SignerQuery) {
				*wq = *query
			})
		case "block":
			if _, ok := fieldSeen[eventlog.FieldBlock]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldBlock)
				fieldSeen[eventlog.FieldBlock] = struct{}{}
			}
		case "signerscount":
			if _, ok := fieldSeen[eventlog.FieldSignersCount]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldSignersCount)
				fieldSeen[eventlog.FieldSignersCount] = struct{}{}
			}
		case "signature":
			if _, ok := fieldSeen[eventlog.FieldSignature]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldSignature)
				fieldSeen[eventlog.FieldSignature] = struct{}{}
			}
		case "address":
			if _, ok := fieldSeen[eventlog.FieldAddress]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldAddress)
				fieldSeen[eventlog.FieldAddress] = struct{}{}
			}
		case "chain":
			if _, ok := fieldSeen[eventlog.FieldChain]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldChain)
				fieldSeen[eventlog.FieldChain] = struct{}{}
			}
		case "index":
			if _, ok := fieldSeen[eventlog.FieldIndex]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldIndex)
				fieldSeen[eventlog.FieldIndex] = struct{}{}
			}
		case "event":
			if _, ok := fieldSeen[eventlog.FieldEvent]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldEvent)
				fieldSeen[eventlog.FieldEvent] = struct{}{}
			}
		case "transaction":
			if _, ok := fieldSeen[eventlog.FieldTransaction]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldTransaction)
				fieldSeen[eventlog.FieldTransaction] = struct{}{}
			}
		case "args":
			if _, ok := fieldSeen[eventlog.FieldArgs]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldArgs)
				fieldSeen[eventlog.FieldArgs] = struct{}{}
			}
		case "consensus":
			if _, ok := fieldSeen[eventlog.FieldConsensus]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldConsensus)
				fieldSeen[eventlog.FieldConsensus] = struct{}{}
			}
		case "voted":
			if _, ok := fieldSeen[eventlog.FieldVoted]; !ok {
				selectedFields = append(selectedFields, eventlog.FieldVoted)
				fieldSeen[eventlog.FieldVoted] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		el.Select(selectedFields...)
	}
	return nil
}

type eventlogPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []EventLogPaginateOption
}

func newEventLogPaginateArgs(rv map[string]any) *eventlogPaginateArgs {
	args := &eventlogPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &EventLogOrder{Field: &EventLogOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithEventLogOrder(order))
			}
		case *EventLogOrder:
			if v != nil {
				args.opts = append(args.opts, WithEventLogOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*EventLogWhereInput); ok {
		args.opts = append(args.opts, WithEventLogFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SignerQuery) CollectFields(ctx context.Context, satisfies ...string) (*SignerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SignerQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(signer.Columns))
		selectedFields = []string{signer.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "assetprice":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AssetPriceClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, assetpriceImplementors)...); err != nil {
				return err
			}
			s.WithNamedAssetPrice(alias, func(wq *AssetPriceQuery) {
				*wq = *query
			})
		case "eventlogs":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EventLogClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, eventlogImplementors)...); err != nil {
				return err
			}
			s.WithNamedEventLogs(alias, func(wq *EventLogQuery) {
				*wq = *query
			})
		case "correctnessreport":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CorrectnessReportClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, correctnessreportImplementors)...); err != nil {
				return err
			}
			s.WithNamedCorrectnessReport(alias, func(wq *CorrectnessReportQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[signer.FieldName]; !ok {
				selectedFields = append(selectedFields, signer.FieldName)
				fieldSeen[signer.FieldName] = struct{}{}
			}
		case "evm":
			if _, ok := fieldSeen[signer.FieldEvm]; !ok {
				selectedFields = append(selectedFields, signer.FieldEvm)
				fieldSeen[signer.FieldEvm] = struct{}{}
			}
		case "key":
			if _, ok := fieldSeen[signer.FieldKey]; !ok {
				selectedFields = append(selectedFields, signer.FieldKey)
				fieldSeen[signer.FieldKey] = struct{}{}
			}
		case "shortkey":
			if _, ok := fieldSeen[signer.FieldShortkey]; !ok {
				selectedFields = append(selectedFields, signer.FieldShortkey)
				fieldSeen[signer.FieldShortkey] = struct{}{}
			}
		case "points":
			if _, ok := fieldSeen[signer.FieldPoints]; !ok {
				selectedFields = append(selectedFields, signer.FieldPoints)
				fieldSeen[signer.FieldPoints] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type signerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SignerPaginateOption
}

func newSignerPaginateArgs(rv map[string]any) *signerPaginateArgs {
	args := &signerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &SignerOrder{Field: &SignerOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithSignerOrder(order))
			}
		case *SignerOrder:
			if v != nil {
				args.opts = append(args.opts, WithSignerOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*SignerWhereInput); ok {
		args.opts = append(args.opts, WithSignerFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
