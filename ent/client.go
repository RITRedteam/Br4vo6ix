// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/BradHacker/Br4vo6ix/ent/migrate"

	"github.com/BradHacker/Br4vo6ix/ent/heartbeat"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/task"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Heartbeat is the client for interacting with the Heartbeat builders.
	Heartbeat *HeartbeatClient
	// Implant is the client for interacting with the Implant builders.
	Implant *ImplantClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Heartbeat = NewHeartbeatClient(c.config)
	c.Implant = NewImplantClient(c.config)
	c.Task = NewTaskClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Heartbeat: NewHeartbeatClient(cfg),
		Implant:   NewImplantClient(cfg),
		Task:      NewTaskClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:    cfg,
		Heartbeat: NewHeartbeatClient(cfg),
		Implant:   NewImplantClient(cfg),
		Task:      NewTaskClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Heartbeat.
//		Query().
//		Count(ctx)
//
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
	c.Heartbeat.Use(hooks...)
	c.Implant.Use(hooks...)
	c.Task.Use(hooks...)
}

// HeartbeatClient is a client for the Heartbeat schema.
type HeartbeatClient struct {
	config
}

// NewHeartbeatClient returns a client for the Heartbeat from the given config.
func NewHeartbeatClient(c config) *HeartbeatClient {
	return &HeartbeatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `heartbeat.Hooks(f(g(h())))`.
func (c *HeartbeatClient) Use(hooks ...Hook) {
	c.hooks.Heartbeat = append(c.hooks.Heartbeat, hooks...)
}

// Create returns a create builder for Heartbeat.
func (c *HeartbeatClient) Create() *HeartbeatCreate {
	mutation := newHeartbeatMutation(c.config, OpCreate)
	return &HeartbeatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Heartbeat entities.
func (c *HeartbeatClient) CreateBulk(builders ...*HeartbeatCreate) *HeartbeatCreateBulk {
	return &HeartbeatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Heartbeat.
func (c *HeartbeatClient) Update() *HeartbeatUpdate {
	mutation := newHeartbeatMutation(c.config, OpUpdate)
	return &HeartbeatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HeartbeatClient) UpdateOne(h *Heartbeat) *HeartbeatUpdateOne {
	mutation := newHeartbeatMutation(c.config, OpUpdateOne, withHeartbeat(h))
	return &HeartbeatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HeartbeatClient) UpdateOneID(id int) *HeartbeatUpdateOne {
	mutation := newHeartbeatMutation(c.config, OpUpdateOne, withHeartbeatID(id))
	return &HeartbeatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Heartbeat.
func (c *HeartbeatClient) Delete() *HeartbeatDelete {
	mutation := newHeartbeatMutation(c.config, OpDelete)
	return &HeartbeatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *HeartbeatClient) DeleteOne(h *Heartbeat) *HeartbeatDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *HeartbeatClient) DeleteOneID(id int) *HeartbeatDeleteOne {
	builder := c.Delete().Where(heartbeat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HeartbeatDeleteOne{builder}
}

// Query returns a query builder for Heartbeat.
func (c *HeartbeatClient) Query() *HeartbeatQuery {
	return &HeartbeatQuery{
		config: c.config,
	}
}

// Get returns a Heartbeat entity by its id.
func (c *HeartbeatClient) Get(ctx context.Context, id int) (*Heartbeat, error) {
	return c.Query().Where(heartbeat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HeartbeatClient) GetX(ctx context.Context, id int) *Heartbeat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImplant queries the implant edge of a Heartbeat.
func (c *HeartbeatClient) QueryImplant(h *Heartbeat) *ImplantQuery {
	query := &ImplantQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(heartbeat.Table, heartbeat.FieldID, id),
			sqlgraph.To(implant.Table, implant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, heartbeat.ImplantTable, heartbeat.ImplantColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HeartbeatClient) Hooks() []Hook {
	return c.hooks.Heartbeat
}

// ImplantClient is a client for the Implant schema.
type ImplantClient struct {
	config
}

// NewImplantClient returns a client for the Implant from the given config.
func NewImplantClient(c config) *ImplantClient {
	return &ImplantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `implant.Hooks(f(g(h())))`.
func (c *ImplantClient) Use(hooks ...Hook) {
	c.hooks.Implant = append(c.hooks.Implant, hooks...)
}

// Create returns a create builder for Implant.
func (c *ImplantClient) Create() *ImplantCreate {
	mutation := newImplantMutation(c.config, OpCreate)
	return &ImplantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Implant entities.
func (c *ImplantClient) CreateBulk(builders ...*ImplantCreate) *ImplantCreateBulk {
	return &ImplantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Implant.
func (c *ImplantClient) Update() *ImplantUpdate {
	mutation := newImplantMutation(c.config, OpUpdate)
	return &ImplantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImplantClient) UpdateOne(i *Implant) *ImplantUpdateOne {
	mutation := newImplantMutation(c.config, OpUpdateOne, withImplant(i))
	return &ImplantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImplantClient) UpdateOneID(id int) *ImplantUpdateOne {
	mutation := newImplantMutation(c.config, OpUpdateOne, withImplantID(id))
	return &ImplantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Implant.
func (c *ImplantClient) Delete() *ImplantDelete {
	mutation := newImplantMutation(c.config, OpDelete)
	return &ImplantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ImplantClient) DeleteOne(i *Implant) *ImplantDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ImplantClient) DeleteOneID(id int) *ImplantDeleteOne {
	builder := c.Delete().Where(implant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImplantDeleteOne{builder}
}

// Query returns a query builder for Implant.
func (c *ImplantClient) Query() *ImplantQuery {
	return &ImplantQuery{
		config: c.config,
	}
}

// Get returns a Implant entity by its id.
func (c *ImplantClient) Get(ctx context.Context, id int) (*Implant, error) {
	return c.Query().Where(implant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImplantClient) GetX(ctx context.Context, id int) *Implant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHeartbeats queries the heartbeats edge of a Implant.
func (c *ImplantClient) QueryHeartbeats(i *Implant) *HeartbeatQuery {
	query := &HeartbeatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(implant.Table, implant.FieldID, id),
			sqlgraph.To(heartbeat.Table, heartbeat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, implant.HeartbeatsTable, implant.HeartbeatsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTasks queries the tasks edge of a Implant.
func (c *ImplantClient) QueryTasks(i *Implant) *TaskQuery {
	query := &TaskQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(implant.Table, implant.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, implant.TasksTable, implant.TasksColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImplantClient) Hooks() []Hook {
	return c.hooks.Implant
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `task.Hooks(f(g(h())))`.
func (c *TaskClient) Use(hooks ...Hook) {
	c.hooks.Task = append(c.hooks.Task, hooks...)
}

// Create returns a create builder for Task.
func (c *TaskClient) Create() *TaskCreate {
	mutation := newTaskMutation(c.config, OpCreate)
	return &TaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Task entities.
func (c *TaskClient) CreateBulk(builders ...*TaskCreate) *TaskCreateBulk {
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	mutation := newTaskMutation(c.config, OpUpdate)
	return &TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTask(t))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id int) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTaskID(id))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	mutation := newTaskMutation(c.config, OpDelete)
	return &TaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TaskClient) DeleteOneID(id int) *TaskDeleteOne {
	builder := c.Delete().Where(task.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskDeleteOne{builder}
}

// Query returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{
		config: c.config,
	}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id int) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id int) *Task {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryImplant queries the implant edge of a Task.
func (c *TaskClient) QueryImplant(t *Task) *ImplantQuery {
	query := &ImplantQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(implant.Table, implant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, task.ImplantTable, task.ImplantColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TaskClient) Hooks() []Hook {
	return c.hooks.Task
}
