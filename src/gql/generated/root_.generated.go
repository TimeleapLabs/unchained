// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/KenshiTech/unchained/ent"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	AssetPrice() AssetPriceResolver
	CorrectnessReport() CorrectnessReportResolver
	EventLog() EventLogResolver
	EventLogArg() EventLogArgResolver
	Query() QueryResolver
	Signer() SignerResolver
	AssetPriceWhereInput() AssetPriceWhereInputResolver
	CorrectnessReportWhereInput() CorrectnessReportWhereInputResolver
	SignerWhereInput() SignerWhereInputResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	AssetPrice struct {
		Asset        func(childComplexity int) int
		Block        func(childComplexity int) int
		Chain        func(childComplexity int) int
		ID           func(childComplexity int) int
		Pair         func(childComplexity int) int
		Price        func(childComplexity int) int
		Signature    func(childComplexity int) int
		Signers      func(childComplexity int) int
		SignersCount func(childComplexity int) int
	}

	AssetPriceConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	AssetPriceEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	CorrectnessReport struct {
		Correct      func(childComplexity int) int
		Hash         func(childComplexity int) int
		ID           func(childComplexity int) int
		Signature    func(childComplexity int) int
		Signers      func(childComplexity int) int
		SignersCount func(childComplexity int) int
		Timestamp    func(childComplexity int) int
		Topic        func(childComplexity int) int
	}

	CorrectnessReportConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	CorrectnessReportEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	EventLog struct {
		Address      func(childComplexity int) int
		Args         func(childComplexity int) int
		Block        func(childComplexity int) int
		Chain        func(childComplexity int) int
		Event        func(childComplexity int) int
		ID           func(childComplexity int) int
		Index        func(childComplexity int) int
		Signature    func(childComplexity int) int
		Signers      func(childComplexity int) int
		SignersCount func(childComplexity int) int
		Transaction  func(childComplexity int) int
	}

	EventLogArg struct {
		Name  func(childComplexity int) int
		Value func(childComplexity int) int
	}

	EventLogConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	EventLogEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		AssetPrices        func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AssetPriceOrder, where *ent.AssetPriceWhereInput) int
		CorrectnessReports func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.CorrectnessReportOrder, where *ent.CorrectnessReportWhereInput) int
		EventLogs          func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.EventLogOrder, where *ent.EventLogWhereInput) int
		Node               func(childComplexity int, id int) int
		Nodes              func(childComplexity int, ids []int) int
		Signers            func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.SignerOrder, where *ent.SignerWhereInput) int
	}

	Signer struct {
		AssetPrice        func(childComplexity int) int
		CorrectnessReport func(childComplexity int) int
		EventLogs         func(childComplexity int) int
		Evm               func(childComplexity int) int
		ID                func(childComplexity int) int
		Key               func(childComplexity int) int
		Name              func(childComplexity int) int
		Points            func(childComplexity int) int
		Shortkey          func(childComplexity int) int
	}

	SignerConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	SignerEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "AssetPrice.asset":
		if e.complexity.AssetPrice.Asset == nil {
			break
		}

		return e.complexity.AssetPrice.Asset(childComplexity), true

	case "AssetPrice.block":
		if e.complexity.AssetPrice.Block == nil {
			break
		}

		return e.complexity.AssetPrice.Block(childComplexity), true

	case "AssetPrice.chain":
		if e.complexity.AssetPrice.Chain == nil {
			break
		}

		return e.complexity.AssetPrice.Chain(childComplexity), true

	case "AssetPrice.id":
		if e.complexity.AssetPrice.ID == nil {
			break
		}

		return e.complexity.AssetPrice.ID(childComplexity), true

	case "AssetPrice.pair":
		if e.complexity.AssetPrice.Pair == nil {
			break
		}

		return e.complexity.AssetPrice.Pair(childComplexity), true

	case "AssetPrice.price":
		if e.complexity.AssetPrice.Price == nil {
			break
		}

		return e.complexity.AssetPrice.Price(childComplexity), true

	case "AssetPrice.signature":
		if e.complexity.AssetPrice.Signature == nil {
			break
		}

		return e.complexity.AssetPrice.Signature(childComplexity), true

	case "AssetPrice.signers":
		if e.complexity.AssetPrice.Signers == nil {
			break
		}

		return e.complexity.AssetPrice.Signers(childComplexity), true

	case "AssetPrice.signerscount":
		if e.complexity.AssetPrice.SignersCount == nil {
			break
		}

		return e.complexity.AssetPrice.SignersCount(childComplexity), true

	case "AssetPriceConnection.edges":
		if e.complexity.AssetPriceConnection.Edges == nil {
			break
		}

		return e.complexity.AssetPriceConnection.Edges(childComplexity), true

	case "AssetPriceConnection.pageInfo":
		if e.complexity.AssetPriceConnection.PageInfo == nil {
			break
		}

		return e.complexity.AssetPriceConnection.PageInfo(childComplexity), true

	case "AssetPriceConnection.totalCount":
		if e.complexity.AssetPriceConnection.TotalCount == nil {
			break
		}

		return e.complexity.AssetPriceConnection.TotalCount(childComplexity), true

	case "AssetPriceEdge.cursor":
		if e.complexity.AssetPriceEdge.Cursor == nil {
			break
		}

		return e.complexity.AssetPriceEdge.Cursor(childComplexity), true

	case "AssetPriceEdge.node":
		if e.complexity.AssetPriceEdge.Node == nil {
			break
		}

		return e.complexity.AssetPriceEdge.Node(childComplexity), true

	case "CorrectnessReport.correct":
		if e.complexity.CorrectnessReport.Correct == nil {
			break
		}

		return e.complexity.CorrectnessReport.Correct(childComplexity), true

	case "CorrectnessReport.hash":
		if e.complexity.CorrectnessReport.Hash == nil {
			break
		}

		return e.complexity.CorrectnessReport.Hash(childComplexity), true

	case "CorrectnessReport.id":
		if e.complexity.CorrectnessReport.ID == nil {
			break
		}

		return e.complexity.CorrectnessReport.ID(childComplexity), true

	case "CorrectnessReport.signature":
		if e.complexity.CorrectnessReport.Signature == nil {
			break
		}

		return e.complexity.CorrectnessReport.Signature(childComplexity), true

	case "CorrectnessReport.signers":
		if e.complexity.CorrectnessReport.Signers == nil {
			break
		}

		return e.complexity.CorrectnessReport.Signers(childComplexity), true

	case "CorrectnessReport.signerscount":
		if e.complexity.CorrectnessReport.SignersCount == nil {
			break
		}

		return e.complexity.CorrectnessReport.SignersCount(childComplexity), true

	case "CorrectnessReport.timestamp":
		if e.complexity.CorrectnessReport.Timestamp == nil {
			break
		}

		return e.complexity.CorrectnessReport.Timestamp(childComplexity), true

	case "CorrectnessReport.topic":
		if e.complexity.CorrectnessReport.Topic == nil {
			break
		}

		return e.complexity.CorrectnessReport.Topic(childComplexity), true

	case "CorrectnessReportConnection.edges":
		if e.complexity.CorrectnessReportConnection.Edges == nil {
			break
		}

		return e.complexity.CorrectnessReportConnection.Edges(childComplexity), true

	case "CorrectnessReportConnection.pageInfo":
		if e.complexity.CorrectnessReportConnection.PageInfo == nil {
			break
		}

		return e.complexity.CorrectnessReportConnection.PageInfo(childComplexity), true

	case "CorrectnessReportConnection.totalCount":
		if e.complexity.CorrectnessReportConnection.TotalCount == nil {
			break
		}

		return e.complexity.CorrectnessReportConnection.TotalCount(childComplexity), true

	case "CorrectnessReportEdge.cursor":
		if e.complexity.CorrectnessReportEdge.Cursor == nil {
			break
		}

		return e.complexity.CorrectnessReportEdge.Cursor(childComplexity), true

	case "CorrectnessReportEdge.node":
		if e.complexity.CorrectnessReportEdge.Node == nil {
			break
		}

		return e.complexity.CorrectnessReportEdge.Node(childComplexity), true

	case "EventLog.address":
		if e.complexity.EventLog.Address == nil {
			break
		}

		return e.complexity.EventLog.Address(childComplexity), true

	case "EventLog.args":
		if e.complexity.EventLog.Args == nil {
			break
		}

		return e.complexity.EventLog.Args(childComplexity), true

	case "EventLog.block":
		if e.complexity.EventLog.Block == nil {
			break
		}

		return e.complexity.EventLog.Block(childComplexity), true

	case "EventLog.chain":
		if e.complexity.EventLog.Chain == nil {
			break
		}

		return e.complexity.EventLog.Chain(childComplexity), true

	case "EventLog.event":
		if e.complexity.EventLog.Event == nil {
			break
		}

		return e.complexity.EventLog.Event(childComplexity), true

	case "EventLog.id":
		if e.complexity.EventLog.ID == nil {
			break
		}

		return e.complexity.EventLog.ID(childComplexity), true

	case "EventLog.index":
		if e.complexity.EventLog.Index == nil {
			break
		}

		return e.complexity.EventLog.Index(childComplexity), true

	case "EventLog.signature":
		if e.complexity.EventLog.Signature == nil {
			break
		}

		return e.complexity.EventLog.Signature(childComplexity), true

	case "EventLog.signers":
		if e.complexity.EventLog.Signers == nil {
			break
		}

		return e.complexity.EventLog.Signers(childComplexity), true

	case "EventLog.signerscount":
		if e.complexity.EventLog.SignersCount == nil {
			break
		}

		return e.complexity.EventLog.SignersCount(childComplexity), true

	case "EventLog.transaction":
		if e.complexity.EventLog.Transaction == nil {
			break
		}

		return e.complexity.EventLog.Transaction(childComplexity), true

	case "EventLogArg.name":
		if e.complexity.EventLogArg.Name == nil {
			break
		}

		return e.complexity.EventLogArg.Name(childComplexity), true

	case "EventLogArg.value":
		if e.complexity.EventLogArg.Value == nil {
			break
		}

		return e.complexity.EventLogArg.Value(childComplexity), true

	case "EventLogConnection.edges":
		if e.complexity.EventLogConnection.Edges == nil {
			break
		}

		return e.complexity.EventLogConnection.Edges(childComplexity), true

	case "EventLogConnection.pageInfo":
		if e.complexity.EventLogConnection.PageInfo == nil {
			break
		}

		return e.complexity.EventLogConnection.PageInfo(childComplexity), true

	case "EventLogConnection.totalCount":
		if e.complexity.EventLogConnection.TotalCount == nil {
			break
		}

		return e.complexity.EventLogConnection.TotalCount(childComplexity), true

	case "EventLogEdge.cursor":
		if e.complexity.EventLogEdge.Cursor == nil {
			break
		}

		return e.complexity.EventLogEdge.Cursor(childComplexity), true

	case "EventLogEdge.node":
		if e.complexity.EventLogEdge.Node == nil {
			break
		}

		return e.complexity.EventLogEdge.Node(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.assetPrices":
		if e.complexity.Query.AssetPrices == nil {
			break
		}

		args, err := ec.field_Query_assetPrices_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.AssetPrices(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["orderBy"].(*ent.AssetPriceOrder), args["where"].(*ent.AssetPriceWhereInput)), true

	case "Query.correctnessReports":
		if e.complexity.Query.CorrectnessReports == nil {
			break
		}

		args, err := ec.field_Query_correctnessReports_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.CorrectnessReports(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["orderBy"].(*ent.CorrectnessReportOrder), args["where"].(*ent.CorrectnessReportWhereInput)), true

	case "Query.eventLogs":
		if e.complexity.Query.EventLogs == nil {
			break
		}

		args, err := ec.field_Query_eventLogs_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.EventLogs(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["orderBy"].(*ent.EventLogOrder), args["where"].(*ent.EventLogWhereInput)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	case "Query.signers":
		if e.complexity.Query.Signers == nil {
			break
		}

		args, err := ec.field_Query_signers_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Signers(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["orderBy"].(*ent.SignerOrder), args["where"].(*ent.SignerWhereInput)), true

	case "Signer.assetprice":
		if e.complexity.Signer.AssetPrice == nil {
			break
		}

		return e.complexity.Signer.AssetPrice(childComplexity), true

	case "Signer.correctnessreport":
		if e.complexity.Signer.CorrectnessReport == nil {
			break
		}

		return e.complexity.Signer.CorrectnessReport(childComplexity), true

	case "Signer.eventlogs":
		if e.complexity.Signer.EventLogs == nil {
			break
		}

		return e.complexity.Signer.EventLogs(childComplexity), true

	case "Signer.evm":
		if e.complexity.Signer.Evm == nil {
			break
		}

		return e.complexity.Signer.Evm(childComplexity), true

	case "Signer.id":
		if e.complexity.Signer.ID == nil {
			break
		}

		return e.complexity.Signer.ID(childComplexity), true

	case "Signer.key":
		if e.complexity.Signer.Key == nil {
			break
		}

		return e.complexity.Signer.Key(childComplexity), true

	case "Signer.name":
		if e.complexity.Signer.Name == nil {
			break
		}

		return e.complexity.Signer.Name(childComplexity), true

	case "Signer.points":
		if e.complexity.Signer.Points == nil {
			break
		}

		return e.complexity.Signer.Points(childComplexity), true

	case "Signer.shortkey":
		if e.complexity.Signer.Shortkey == nil {
			break
		}

		return e.complexity.Signer.Shortkey(childComplexity), true

	case "SignerConnection.edges":
		if e.complexity.SignerConnection.Edges == nil {
			break
		}

		return e.complexity.SignerConnection.Edges(childComplexity), true

	case "SignerConnection.pageInfo":
		if e.complexity.SignerConnection.PageInfo == nil {
			break
		}

		return e.complexity.SignerConnection.PageInfo(childComplexity), true

	case "SignerConnection.totalCount":
		if e.complexity.SignerConnection.TotalCount == nil {
			break
		}

		return e.complexity.SignerConnection.TotalCount(childComplexity), true

	case "SignerEdge.cursor":
		if e.complexity.SignerEdge.Cursor == nil {
			break
		}

		return e.complexity.SignerEdge.Cursor(childComplexity), true

	case "SignerEdge.node":
		if e.complexity.SignerEdge.Node == nil {
			break
		}

		return e.complexity.SignerEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAssetPriceOrder,
		ec.unmarshalInputAssetPriceWhereInput,
		ec.unmarshalInputCorrectnessReportOrder,
		ec.unmarshalInputCorrectnessReportWhereInput,
		ec.unmarshalInputEventLogOrder,
		ec.unmarshalInputEventLogWhereInput,
		ec.unmarshalInputSignerOrder,
		ec.unmarshalInputSignerWhereInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../args.graphql", Input: `scalar BigInt
scalar Bytes

type EventLogArg {
  name: String!
  value: String!
}
`, BuiltIn: false},
	{Name: "../extensions.graphql", Input: `extend input SignerWhereInput {
  key: String
}

extend input CorrectnessReportWhereInput {
  topic: String
  hash: String
}
`, BuiltIn: false},
	{Name: "../unchained.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
type AssetPrice implements Node {
  id: ID!
  block: Uint!
  signerscount: Uint @goField(name: "SignersCount", forceResolver: false)
  price: Uint!
  signature: Bytes!
  asset: String
  chain: String
  pair: String
  signers: [Signer!]!
}
"""
A connection to a list of items.
"""
type AssetPriceConnection {
  """
  A list of edges.
  """
  edges: [AssetPriceEdge]
  """
  Information to aid in pagination.
  """
  pageInfo: PageInfo!
  """
  Identifies the total count of items in the connection.
  """
  totalCount: Int!
}
"""
An edge in a connection.
"""
type AssetPriceEdge {
  """
  The item at the end of the edge.
  """
  node: AssetPrice
  """
  A cursor for use in pagination.
  """
  cursor: Cursor!
}
"""
Ordering options for AssetPrice connections
"""
input AssetPriceOrder {
  """
  The ordering direction.
  """
  direction: OrderDirection! = ASC
  """
  The field by which to order AssetPrices.
  """
  field: AssetPriceOrderField!
}
"""
Properties by which AssetPrice connections can be ordered.
"""
enum AssetPriceOrderField {
  BLOCK
}
"""
AssetPriceWhereInput is used for filtering AssetPrice objects.
Input was generated by ent.
"""
input AssetPriceWhereInput {
  not: AssetPriceWhereInput
  and: [AssetPriceWhereInput!]
  or: [AssetPriceWhereInput!]
  """
  id field predicates
  """
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """
  block field predicates
  """
  block: Uint
  blockNEQ: Uint
  blockIn: [Uint!]
  blockNotIn: [Uint!]
  blockGT: Uint
  blockGTE: Uint
  blockLT: Uint
  blockLTE: Uint
  """
  signersCount field predicates
  """
  signerscount: Uint
  signerscountNEQ: Uint
  signerscountIn: [Uint!]
  signerscountNotIn: [Uint!]
  signerscountGT: Uint
  signerscountGTE: Uint
  signerscountLT: Uint
  signerscountLTE: Uint
  signerscountIsNil: Boolean
  signerscountNotNil: Boolean
  """
  price field predicates
  """
  price: Uint
  priceNEQ: Uint
  priceIn: [Uint!]
  priceNotIn: [Uint!]
  priceGT: Uint
  priceGTE: Uint
  priceLT: Uint
  priceLTE: Uint
  """
  asset field predicates
  """
  asset: String
  assetNEQ: String
  assetIn: [String!]
  assetNotIn: [String!]
  assetGT: String
  assetGTE: String
  assetLT: String
  assetLTE: String
  assetContains: String
  assetHasPrefix: String
  assetHasSuffix: String
  assetIsNil: Boolean
  assetNotNil: Boolean
  assetEqualFold: String
  assetContainsFold: String
  """
  chain field predicates
  """
  chain: String
  chainNEQ: String
  chainIn: [String!]
  chainNotIn: [String!]
  chainGT: String
  chainGTE: String
  chainLT: String
  chainLTE: String
  chainContains: String
  chainHasPrefix: String
  chainHasSuffix: String
  chainIsNil: Boolean
  chainNotNil: Boolean
  chainEqualFold: String
  chainContainsFold: String
  """
  pair field predicates
  """
  pair: String
  pairNEQ: String
  pairIn: [String!]
  pairNotIn: [String!]
  pairGT: String
  pairGTE: String
  pairLT: String
  pairLTE: String
  pairContains: String
  pairHasPrefix: String
  pairHasSuffix: String
  pairIsNil: Boolean
  pairNotNil: Boolean
  pairEqualFold: String
  pairContainsFold: String
  """
  signers edge predicates
  """
  hasSigners: Boolean
  hasSignersWith: [SignerWhereInput!]
}
type CorrectnessReport implements Node {
  id: ID!
  signerscount: Uint! @goField(name: "SignersCount", forceResolver: false)
  timestamp: Uint!
  signature: Bytes!
  hash: Bytes!
  topic: Bytes!
  correct: Boolean!
  signers: [Signer!]!
}
"""
A connection to a list of items.
"""
type CorrectnessReportConnection {
  """
  A list of edges.
  """
  edges: [CorrectnessReportEdge]
  """
  Information to aid in pagination.
  """
  pageInfo: PageInfo!
  """
  Identifies the total count of items in the connection.
  """
  totalCount: Int!
}
"""
An edge in a connection.
"""
type CorrectnessReportEdge {
  """
  The item at the end of the edge.
  """
  node: CorrectnessReport
  """
  A cursor for use in pagination.
  """
  cursor: Cursor!
}
"""
Ordering options for CorrectnessReport connections
"""
input CorrectnessReportOrder {
  """
  The ordering direction.
  """
  direction: OrderDirection! = ASC
  """
  The field by which to order CorrectnessReports.
  """
  field: CorrectnessReportOrderField!
}
"""
Properties by which CorrectnessReport connections can be ordered.
"""
enum CorrectnessReportOrderField {
  TIMESTAMP
}
"""
CorrectnessReportWhereInput is used for filtering CorrectnessReport objects.
Input was generated by ent.
"""
input CorrectnessReportWhereInput {
  not: CorrectnessReportWhereInput
  and: [CorrectnessReportWhereInput!]
  or: [CorrectnessReportWhereInput!]
  """
  id field predicates
  """
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """
  signersCount field predicates
  """
  signerscount: Uint
  signerscountNEQ: Uint
  signerscountIn: [Uint!]
  signerscountNotIn: [Uint!]
  signerscountGT: Uint
  signerscountGTE: Uint
  signerscountLT: Uint
  signerscountLTE: Uint
  """
  timestamp field predicates
  """
  timestamp: Uint
  timestampNEQ: Uint
  timestampIn: [Uint!]
  timestampNotIn: [Uint!]
  timestampGT: Uint
  timestampGTE: Uint
  timestampLT: Uint
  timestampLTE: Uint
  """
  correct field predicates
  """
  correct: Boolean
  correctNEQ: Boolean
  """
  signers edge predicates
  """
  hasSigners: Boolean
  hasSignersWith: [SignerWhereInput!]
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
type EventLog implements Node {
  id: ID!
  block: Uint!
  signerscount: Uint! @goField(name: "SignersCount", forceResolver: false)
  signature: Bytes!
  address: String!
  chain: String!
  index: Uint!
  event: String!
  transaction: Bytes!
  args: [EventLogArg!]!
  signers: [Signer!]!
}
"""
A connection to a list of items.
"""
type EventLogConnection {
  """
  A list of edges.
  """
  edges: [EventLogEdge]
  """
  Information to aid in pagination.
  """
  pageInfo: PageInfo!
  """
  Identifies the total count of items in the connection.
  """
  totalCount: Int!
}
"""
An edge in a connection.
"""
type EventLogEdge {
  """
  The item at the end of the edge.
  """
  node: EventLog
  """
  A cursor for use in pagination.
  """
  cursor: Cursor!
}
"""
Ordering options for EventLog connections
"""
input EventLogOrder {
  """
  The ordering direction.
  """
  direction: OrderDirection! = ASC
  """
  The field by which to order EventLogs.
  """
  field: EventLogOrderField!
}
"""
Properties by which EventLog connections can be ordered.
"""
enum EventLogOrderField {
  BLOCK
}
"""
EventLogWhereInput is used for filtering EventLog objects.
Input was generated by ent.
"""
input EventLogWhereInput {
  not: EventLogWhereInput
  and: [EventLogWhereInput!]
  or: [EventLogWhereInput!]
  """
  id field predicates
  """
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """
  block field predicates
  """
  block: Uint
  blockNEQ: Uint
  blockIn: [Uint!]
  blockNotIn: [Uint!]
  blockGT: Uint
  blockGTE: Uint
  blockLT: Uint
  blockLTE: Uint
  """
  signersCount field predicates
  """
  signerscount: Uint
  signerscountNEQ: Uint
  signerscountIn: [Uint!]
  signerscountNotIn: [Uint!]
  signerscountGT: Uint
  signerscountGTE: Uint
  signerscountLT: Uint
  signerscountLTE: Uint
  """
  address field predicates
  """
  address: String
  addressNEQ: String
  addressIn: [String!]
  addressNotIn: [String!]
  addressGT: String
  addressGTE: String
  addressLT: String
  addressLTE: String
  addressContains: String
  addressHasPrefix: String
  addressHasSuffix: String
  addressEqualFold: String
  addressContainsFold: String
  """
  chain field predicates
  """
  chain: String
  chainNEQ: String
  chainIn: [String!]
  chainNotIn: [String!]
  chainGT: String
  chainGTE: String
  chainLT: String
  chainLTE: String
  chainContains: String
  chainHasPrefix: String
  chainHasSuffix: String
  chainEqualFold: String
  chainContainsFold: String
  """
  index field predicates
  """
  index: Uint
  indexNEQ: Uint
  indexIn: [Uint!]
  indexNotIn: [Uint!]
  indexGT: Uint
  indexGTE: Uint
  indexLT: Uint
  indexLTE: Uint
  """
  event field predicates
  """
  event: String
  eventNEQ: String
  eventIn: [String!]
  eventNotIn: [String!]
  eventGT: String
  eventGTE: String
  eventLT: String
  eventLTE: String
  eventContains: String
  eventHasPrefix: String
  eventHasSuffix: String
  eventEqualFold: String
  eventContainsFold: String
  """
  signers edge predicates
  """
  hasSigners: Boolean
  hasSignersWith: [SignerWhereInput!]
}
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "github.com/KenshiTech/unchained/ent.Noder") {
  """
  The id of the object.
  """
  id: ID!
}
"""
Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument.
"""
enum OrderDirection {
  """
  Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument.
  """
  ASC
  """
  Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument.
  """
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """
  When paginating forwards, are there more items?
  """
  hasNextPage: Boolean!
  """
  When paginating backwards, are there more items?
  """
  hasPreviousPage: Boolean!
  """
  When paginating backwards, the cursor to continue.
  """
  startCursor: Cursor
  """
  When paginating forwards, the cursor to continue.
  """
  endCursor: Cursor
}
type Query {
  """
  Fetches an object given its ID.
  """
  node(
    """
    ID of the object.
    """
    id: ID!
  ): Node
  """
  Lookup nodes by a list of IDs.
  """
  nodes(
    """
    The list of node IDs.
    """
    ids: [ID!]!
  ): [Node]!
  assetPrices(
    """
    Returns the elements in the list that come after the specified cursor.
    """
    after: Cursor

    """
    Returns the first _n_ elements from the list.
    """
    first: Int

    """
    Returns the elements in the list that come before the specified cursor.
    """
    before: Cursor

    """
    Returns the last _n_ elements from the list.
    """
    last: Int

    """
    Ordering options for AssetPrices returned from the connection.
    """
    orderBy: AssetPriceOrder

    """
    Filtering options for AssetPrices returned from the connection.
    """
    where: AssetPriceWhereInput
  ): AssetPriceConnection!
  correctnessReports(
    """
    Returns the elements in the list that come after the specified cursor.
    """
    after: Cursor

    """
    Returns the first _n_ elements from the list.
    """
    first: Int

    """
    Returns the elements in the list that come before the specified cursor.
    """
    before: Cursor

    """
    Returns the last _n_ elements from the list.
    """
    last: Int

    """
    Ordering options for CorrectnessReports returned from the connection.
    """
    orderBy: CorrectnessReportOrder

    """
    Filtering options for CorrectnessReports returned from the connection.
    """
    where: CorrectnessReportWhereInput
  ): CorrectnessReportConnection!
  eventLogs(
    """
    Returns the elements in the list that come after the specified cursor.
    """
    after: Cursor

    """
    Returns the first _n_ elements from the list.
    """
    first: Int

    """
    Returns the elements in the list that come before the specified cursor.
    """
    before: Cursor

    """
    Returns the last _n_ elements from the list.
    """
    last: Int

    """
    Ordering options for EventLogs returned from the connection.
    """
    orderBy: EventLogOrder

    """
    Filtering options for EventLogs returned from the connection.
    """
    where: EventLogWhereInput
  ): EventLogConnection!
  signers(
    """
    Returns the elements in the list that come after the specified cursor.
    """
    after: Cursor

    """
    Returns the first _n_ elements from the list.
    """
    first: Int

    """
    Returns the elements in the list that come before the specified cursor.
    """
    before: Cursor

    """
    Returns the last _n_ elements from the list.
    """
    last: Int

    """
    Ordering options for Signers returned from the connection.
    """
    orderBy: SignerOrder

    """
    Filtering options for Signers returned from the connection.
    """
    where: SignerWhereInput
  ): SignerConnection!
}
type Signer implements Node {
  id: ID!
  name: String!
  evm: String
  key: Bytes!
  shortkey: Bytes!
  points: Int!
  assetprice: [AssetPrice!] @goField(name: "AssetPrice", forceResolver: false)
  eventlogs: [EventLog!] @goField(name: "EventLogs", forceResolver: false)
  correctnessreport: [CorrectnessReport!] @goField(name: "CorrectnessReport", forceResolver: false)
}
"""
A connection to a list of items.
"""
type SignerConnection {
  """
  A list of edges.
  """
  edges: [SignerEdge]
  """
  Information to aid in pagination.
  """
  pageInfo: PageInfo!
  """
  Identifies the total count of items in the connection.
  """
  totalCount: Int!
}
"""
An edge in a connection.
"""
type SignerEdge {
  """
  The item at the end of the edge.
  """
  node: Signer
  """
  A cursor for use in pagination.
  """
  cursor: Cursor!
}
"""
Ordering options for Signer connections
"""
input SignerOrder {
  """
  The ordering direction.
  """
  direction: OrderDirection! = ASC
  """
  The field by which to order Signers.
  """
  field: SignerOrderField!
}
"""
Properties by which Signer connections can be ordered.
"""
enum SignerOrderField {
  POINTS
}
"""
SignerWhereInput is used for filtering Signer objects.
Input was generated by ent.
"""
input SignerWhereInput {
  not: SignerWhereInput
  and: [SignerWhereInput!]
  or: [SignerWhereInput!]
  """
  id field predicates
  """
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """
  name field predicates
  """
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """
  evm field predicates
  """
  evm: String
  evmNEQ: String
  evmIn: [String!]
  evmNotIn: [String!]
  evmGT: String
  evmGTE: String
  evmLT: String
  evmLTE: String
  evmContains: String
  evmHasPrefix: String
  evmHasSuffix: String
  evmIsNil: Boolean
  evmNotNil: Boolean
  evmEqualFold: String
  evmContainsFold: String
  """
  points field predicates
  """
  points: Int
  pointsNEQ: Int
  pointsIn: [Int!]
  pointsNotIn: [Int!]
  pointsGT: Int
  pointsGTE: Int
  pointsLT: Int
  pointsLTE: Int
  """
  assetPrice edge predicates
  """
  hasAssetPrice: Boolean
  hasAssetPriceWith: [AssetPriceWhereInput!]
  """
  eventLogs edge predicates
  """
  hasEventLogs: Boolean
  hasEventLogsWith: [EventLogWhereInput!]
  """
  correctnessReport edge predicates
  """
  hasCorrectnessReport: Boolean
  hasCorrectnessReportWith: [CorrectnessReportWhereInput!]
}
"""
The builtin Uint type
"""
scalar Uint
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
