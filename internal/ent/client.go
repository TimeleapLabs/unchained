// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/TimeleapLabs/unchained/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/TimeleapLabs/unchained/internal/ent/assetprice"
	"github.com/TimeleapLabs/unchained/internal/ent/correctnessreport"
	"github.com/TimeleapLabs/unchained/internal/ent/eventlog"
	"github.com/TimeleapLabs/unchained/internal/ent/signer"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AssetPrice is the client for interacting with the AssetPrice builders.
	AssetPrice *AssetPriceClient
	// CorrectnessReport is the client for interacting with the CorrectnessReport builders.
	CorrectnessReport *CorrectnessReportClient
	// EventLog is the client for interacting with the EventLog builders.
	EventLog *EventLogClient
	// Signer is the client for interacting with the Signer builders.
	Signer *SignerClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AssetPrice = NewAssetPriceClient(c.config)
	c.CorrectnessReport = NewCorrectnessReportClient(c.config)
	c.EventLog = NewEventLogClient(c.config)
	c.Signer = NewSignerClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		AssetPrice:        NewAssetPriceClient(cfg),
		CorrectnessReport: NewCorrectnessReportClient(cfg),
		EventLog:          NewEventLogClient(cfg),
		Signer:            NewSignerClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		AssetPrice:        NewAssetPriceClient(cfg),
		CorrectnessReport: NewCorrectnessReportClient(cfg),
		EventLog:          NewEventLogClient(cfg),
		Signer:            NewSignerClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AssetPrice.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AssetPrice.Use(hooks...)
	c.CorrectnessReport.Use(hooks...)
	c.EventLog.Use(hooks...)
	c.Signer.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AssetPrice.Intercept(interceptors...)
	c.CorrectnessReport.Intercept(interceptors...)
	c.EventLog.Intercept(interceptors...)
	c.Signer.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AssetPriceMutation:
		return c.AssetPrice.mutate(ctx, m)
	case *CorrectnessReportMutation:
		return c.CorrectnessReport.mutate(ctx, m)
	case *EventLogMutation:
		return c.EventLog.mutate(ctx, m)
	case *SignerMutation:
		return c.Signer.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AssetPriceClient is a client for the AssetPrice schema.
type AssetPriceClient struct {
	config
}

// NewAssetPriceClient returns a client for the AssetPrice from the given config.
func NewAssetPriceClient(c config) *AssetPriceClient {
	return &AssetPriceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `assetprice.Hooks(f(g(h())))`.
func (c *AssetPriceClient) Use(hooks ...Hook) {
	c.hooks.AssetPrice = append(c.hooks.AssetPrice, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `assetprice.Intercept(f(g(h())))`.
func (c *AssetPriceClient) Intercept(interceptors ...Interceptor) {
	c.inters.AssetPrice = append(c.inters.AssetPrice, interceptors...)
}

// Create returns a builder for creating a AssetPrice entity.
func (c *AssetPriceClient) Create() *AssetPriceCreate {
	mutation := newAssetPriceMutation(c.config, OpCreate)
	return &AssetPriceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AssetPrice entities.
func (c *AssetPriceClient) CreateBulk(builders ...*AssetPriceCreate) *AssetPriceCreateBulk {
	return &AssetPriceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AssetPriceClient) MapCreateBulk(slice any, setFunc func(*AssetPriceCreate, int)) *AssetPriceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AssetPriceCreateBulk{err: fmt.Errorf("calling to AssetPriceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AssetPriceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AssetPriceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AssetPrice.
func (c *AssetPriceClient) Update() *AssetPriceUpdate {
	mutation := newAssetPriceMutation(c.config, OpUpdate)
	return &AssetPriceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AssetPriceClient) UpdateOne(ap *AssetPrice) *AssetPriceUpdateOne {
	mutation := newAssetPriceMutation(c.config, OpUpdateOne, withAssetPrice(ap))
	return &AssetPriceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AssetPriceClient) UpdateOneID(id int) *AssetPriceUpdateOne {
	mutation := newAssetPriceMutation(c.config, OpUpdateOne, withAssetPriceID(id))
	return &AssetPriceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AssetPrice.
func (c *AssetPriceClient) Delete() *AssetPriceDelete {
	mutation := newAssetPriceMutation(c.config, OpDelete)
	return &AssetPriceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AssetPriceClient) DeleteOne(ap *AssetPrice) *AssetPriceDeleteOne {
	return c.DeleteOneID(ap.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AssetPriceClient) DeleteOneID(id int) *AssetPriceDeleteOne {
	builder := c.Delete().Where(assetprice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AssetPriceDeleteOne{builder}
}

// Query returns a query builder for AssetPrice.
func (c *AssetPriceClient) Query() *AssetPriceQuery {
	return &AssetPriceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAssetPrice},
		inters: c.Interceptors(),
	}
}

// Get returns a AssetPrice entity by its id.
func (c *AssetPriceClient) Get(ctx context.Context, id int) (*AssetPrice, error) {
	return c.Query().Where(assetprice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AssetPriceClient) GetX(ctx context.Context, id int) *AssetPrice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySigners queries the signers edge of a AssetPrice.
func (c *AssetPriceClient) QuerySigners(ap *AssetPrice) *SignerQuery {
	query := (&SignerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ap.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(assetprice.Table, assetprice.FieldID, id),
			sqlgraph.To(signer.Table, signer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, assetprice.SignersTable, assetprice.SignersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ap.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AssetPriceClient) Hooks() []Hook {
	return c.hooks.AssetPrice
}

// Interceptors returns the client interceptors.
func (c *AssetPriceClient) Interceptors() []Interceptor {
	return c.inters.AssetPrice
}

func (c *AssetPriceClient) mutate(ctx context.Context, m *AssetPriceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AssetPriceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AssetPriceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AssetPriceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AssetPriceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AssetPrice mutation op: %q", m.Op())
	}
}

// CorrectnessReportClient is a client for the CorrectnessReport schema.
type CorrectnessReportClient struct {
	config
}

// NewCorrectnessReportClient returns a client for the CorrectnessReport from the given config.
func NewCorrectnessReportClient(c config) *CorrectnessReportClient {
	return &CorrectnessReportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `correctnessreport.Hooks(f(g(h())))`.
func (c *CorrectnessReportClient) Use(hooks ...Hook) {
	c.hooks.CorrectnessReport = append(c.hooks.CorrectnessReport, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `correctnessreport.Intercept(f(g(h())))`.
func (c *CorrectnessReportClient) Intercept(interceptors ...Interceptor) {
	c.inters.CorrectnessReport = append(c.inters.CorrectnessReport, interceptors...)
}

// Create returns a builder for creating a CorrectnessReport entity.
func (c *CorrectnessReportClient) Create() *CorrectnessReportCreate {
	mutation := newCorrectnessReportMutation(c.config, OpCreate)
	return &CorrectnessReportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CorrectnessReport entities.
func (c *CorrectnessReportClient) CreateBulk(builders ...*CorrectnessReportCreate) *CorrectnessReportCreateBulk {
	return &CorrectnessReportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CorrectnessReportClient) MapCreateBulk(slice any, setFunc func(*CorrectnessReportCreate, int)) *CorrectnessReportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CorrectnessReportCreateBulk{err: fmt.Errorf("calling to CorrectnessReportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CorrectnessReportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CorrectnessReportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CorrectnessReport.
func (c *CorrectnessReportClient) Update() *CorrectnessReportUpdate {
	mutation := newCorrectnessReportMutation(c.config, OpUpdate)
	return &CorrectnessReportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CorrectnessReportClient) UpdateOne(cr *CorrectnessReport) *CorrectnessReportUpdateOne {
	mutation := newCorrectnessReportMutation(c.config, OpUpdateOne, withCorrectnessReport(cr))
	return &CorrectnessReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CorrectnessReportClient) UpdateOneID(id int) *CorrectnessReportUpdateOne {
	mutation := newCorrectnessReportMutation(c.config, OpUpdateOne, withCorrectnessReportID(id))
	return &CorrectnessReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CorrectnessReport.
func (c *CorrectnessReportClient) Delete() *CorrectnessReportDelete {
	mutation := newCorrectnessReportMutation(c.config, OpDelete)
	return &CorrectnessReportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CorrectnessReportClient) DeleteOne(cr *CorrectnessReport) *CorrectnessReportDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CorrectnessReportClient) DeleteOneID(id int) *CorrectnessReportDeleteOne {
	builder := c.Delete().Where(correctnessreport.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CorrectnessReportDeleteOne{builder}
}

// Query returns a query builder for CorrectnessReport.
func (c *CorrectnessReportClient) Query() *CorrectnessReportQuery {
	return &CorrectnessReportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCorrectnessReport},
		inters: c.Interceptors(),
	}
}

// Get returns a CorrectnessReport entity by its id.
func (c *CorrectnessReportClient) Get(ctx context.Context, id int) (*CorrectnessReport, error) {
	return c.Query().Where(correctnessreport.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CorrectnessReportClient) GetX(ctx context.Context, id int) *CorrectnessReport {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySigners queries the signers edge of a CorrectnessReport.
func (c *CorrectnessReportClient) QuerySigners(cr *CorrectnessReport) *SignerQuery {
	query := (&SignerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(correctnessreport.Table, correctnessreport.FieldID, id),
			sqlgraph.To(signer.Table, signer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, correctnessreport.SignersTable, correctnessreport.SignersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CorrectnessReportClient) Hooks() []Hook {
	return c.hooks.CorrectnessReport
}

// Interceptors returns the client interceptors.
func (c *CorrectnessReportClient) Interceptors() []Interceptor {
	return c.inters.CorrectnessReport
}

func (c *CorrectnessReportClient) mutate(ctx context.Context, m *CorrectnessReportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CorrectnessReportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CorrectnessReportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CorrectnessReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CorrectnessReportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CorrectnessReport mutation op: %q", m.Op())
	}
}

// EventLogClient is a client for the EventLog schema.
type EventLogClient struct {
	config
}

// NewEventLogClient returns a client for the EventLog from the given config.
func NewEventLogClient(c config) *EventLogClient {
	return &EventLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `eventlog.Hooks(f(g(h())))`.
func (c *EventLogClient) Use(hooks ...Hook) {
	c.hooks.EventLog = append(c.hooks.EventLog, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `eventlog.Intercept(f(g(h())))`.
func (c *EventLogClient) Intercept(interceptors ...Interceptor) {
	c.inters.EventLog = append(c.inters.EventLog, interceptors...)
}

// Create returns a builder for creating a EventLog entity.
func (c *EventLogClient) Create() *EventLogCreate {
	mutation := newEventLogMutation(c.config, OpCreate)
	return &EventLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EventLog entities.
func (c *EventLogClient) CreateBulk(builders ...*EventLogCreate) *EventLogCreateBulk {
	return &EventLogCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EventLogClient) MapCreateBulk(slice any, setFunc func(*EventLogCreate, int)) *EventLogCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EventLogCreateBulk{err: fmt.Errorf("calling to EventLogClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EventLogCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EventLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EventLog.
func (c *EventLogClient) Update() *EventLogUpdate {
	mutation := newEventLogMutation(c.config, OpUpdate)
	return &EventLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventLogClient) UpdateOne(el *EventLog) *EventLogUpdateOne {
	mutation := newEventLogMutation(c.config, OpUpdateOne, withEventLog(el))
	return &EventLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventLogClient) UpdateOneID(id int) *EventLogUpdateOne {
	mutation := newEventLogMutation(c.config, OpUpdateOne, withEventLogID(id))
	return &EventLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EventLog.
func (c *EventLogClient) Delete() *EventLogDelete {
	mutation := newEventLogMutation(c.config, OpDelete)
	return &EventLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventLogClient) DeleteOne(el *EventLog) *EventLogDeleteOne {
	return c.DeleteOneID(el.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EventLogClient) DeleteOneID(id int) *EventLogDeleteOne {
	builder := c.Delete().Where(eventlog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventLogDeleteOne{builder}
}

// Query returns a query builder for EventLog.
func (c *EventLogClient) Query() *EventLogQuery {
	return &EventLogQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEventLog},
		inters: c.Interceptors(),
	}
}

// Get returns a EventLog entity by its id.
func (c *EventLogClient) Get(ctx context.Context, id int) (*EventLog, error) {
	return c.Query().Where(eventlog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventLogClient) GetX(ctx context.Context, id int) *EventLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySigners queries the signers edge of a EventLog.
func (c *EventLogClient) QuerySigners(el *EventLog) *SignerQuery {
	query := (&SignerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := el.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eventlog.Table, eventlog.FieldID, id),
			sqlgraph.To(signer.Table, signer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, eventlog.SignersTable, eventlog.SignersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(el.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventLogClient) Hooks() []Hook {
	return c.hooks.EventLog
}

// Interceptors returns the client interceptors.
func (c *EventLogClient) Interceptors() []Interceptor {
	return c.inters.EventLog
}

func (c *EventLogClient) mutate(ctx context.Context, m *EventLogMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EventLogCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EventLogUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EventLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EventLogDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EventLog mutation op: %q", m.Op())
	}
}

// SignerClient is a client for the Signer schema.
type SignerClient struct {
	config
}

// NewSignerClient returns a client for the Signer from the given config.
func NewSignerClient(c config) *SignerClient {
	return &SignerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `signer.Hooks(f(g(h())))`.
func (c *SignerClient) Use(hooks ...Hook) {
	c.hooks.Signer = append(c.hooks.Signer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `signer.Intercept(f(g(h())))`.
func (c *SignerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Signer = append(c.inters.Signer, interceptors...)
}

// Create returns a builder for creating a Signer entity.
func (c *SignerClient) Create() *SignerCreate {
	mutation := newSignerMutation(c.config, OpCreate)
	return &SignerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Signer entities.
func (c *SignerClient) CreateBulk(builders ...*SignerCreate) *SignerCreateBulk {
	return &SignerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SignerClient) MapCreateBulk(slice any, setFunc func(*SignerCreate, int)) *SignerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SignerCreateBulk{err: fmt.Errorf("calling to SignerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SignerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SignerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Signer.
func (c *SignerClient) Update() *SignerUpdate {
	mutation := newSignerMutation(c.config, OpUpdate)
	return &SignerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SignerClient) UpdateOne(s *Signer) *SignerUpdateOne {
	mutation := newSignerMutation(c.config, OpUpdateOne, withSigner(s))
	return &SignerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SignerClient) UpdateOneID(id int) *SignerUpdateOne {
	mutation := newSignerMutation(c.config, OpUpdateOne, withSignerID(id))
	return &SignerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Signer.
func (c *SignerClient) Delete() *SignerDelete {
	mutation := newSignerMutation(c.config, OpDelete)
	return &SignerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SignerClient) DeleteOne(s *Signer) *SignerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SignerClient) DeleteOneID(id int) *SignerDeleteOne {
	builder := c.Delete().Where(signer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SignerDeleteOne{builder}
}

// Query returns a query builder for Signer.
func (c *SignerClient) Query() *SignerQuery {
	return &SignerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSigner},
		inters: c.Interceptors(),
	}
}

// Get returns a Signer entity by its id.
func (c *SignerClient) Get(ctx context.Context, id int) (*Signer, error) {
	return c.Query().Where(signer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SignerClient) GetX(ctx context.Context, id int) *Signer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAssetPrice queries the assetPrice edge of a Signer.
func (c *SignerClient) QueryAssetPrice(s *Signer) *AssetPriceQuery {
	query := (&AssetPriceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(signer.Table, signer.FieldID, id),
			sqlgraph.To(assetprice.Table, assetprice.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, signer.AssetPriceTable, signer.AssetPricePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEventLogs queries the eventLogs edge of a Signer.
func (c *SignerClient) QueryEventLogs(s *Signer) *EventLogQuery {
	query := (&EventLogClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(signer.Table, signer.FieldID, id),
			sqlgraph.To(eventlog.Table, eventlog.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, signer.EventLogsTable, signer.EventLogsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCorrectnessReport queries the correctnessReport edge of a Signer.
func (c *SignerClient) QueryCorrectnessReport(s *Signer) *CorrectnessReportQuery {
	query := (&CorrectnessReportClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(signer.Table, signer.FieldID, id),
			sqlgraph.To(correctnessreport.Table, correctnessreport.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, signer.CorrectnessReportTable, signer.CorrectnessReportPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SignerClient) Hooks() []Hook {
	return c.hooks.Signer
}

// Interceptors returns the client interceptors.
func (c *SignerClient) Interceptors() []Interceptor {
	return c.inters.Signer
}

func (c *SignerClient) mutate(ctx context.Context, m *SignerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SignerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SignerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SignerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SignerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Signer mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AssetPrice, CorrectnessReport, EventLog, Signer []ent.Hook
	}
	inters struct {
		AssetPrice, CorrectnessReport, EventLog, Signer []ent.Interceptor
	}
)