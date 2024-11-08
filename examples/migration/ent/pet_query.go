// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"fmt"
	"math"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/examples/migration/fluent/pet"
	"github.com/usalko/fluent/examples/migration/fluent/predicate"
	"github.com/usalko/fluent/examples/migration/fluent/user"
	"github.com/usalko/fluent/schema/field"
	"github.com/google/uuid"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	ctx            *QueryContext
	order          []pet.OrderOption
	inters         []Interceptor
	predicates     []predicate.Pet
	withBestFriend *PetQuery
	withOwner      *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PetQuery builder.
func (pq *PetQuery) Where(ps ...predicate.Pet) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PetQuery) Unique(unique bool) *PetQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PetQuery) Order(o ...pet.OrderOption) *PetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryBestFriend chains the current query on the "best_friend" edge.
func (pq *PetQuery) QueryBestFriend() *PetQuery {
	query := (&PetClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, pet.BestFriendTable, pet.BestFriendColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (pq *PetQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, pet.OwnerTable, pet.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pet entity from the query.
// Returns a *NotFoundError when no Pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pet ID from the query.
// Returns a *NotFoundError when no Pet ID was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pet entity is found.
// Returns a *NotFoundError when no Pet entities are found.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pet.Label}
	default:
		return nil, &NotSingularError{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pet ID in the query.
// Returns a *NotSingularError when more than one Pet ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PetQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = &NotSingularError{pet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pet, *PetQuery]()
	return withInterceptors[[]*Pet](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pet IDs.
func (pq *PetQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryIDs)
	if err = pq.Select(pet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PetQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, fluent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	if pq == nil {
		return nil
	}
	return &PetQuery{
		config:         pq.config,
		ctx:            pq.ctx.Clone(),
		order:          append([]pet.OrderOption{}, pq.order...),
		inters:         append([]Interceptor{}, pq.inters...),
		predicates:     append([]predicate.Pet{}, pq.predicates...),
		withBestFriend: pq.withBestFriend.Clone(),
		withOwner:      pq.withOwner.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithBestFriend tells the query-builder to eager-load the nodes that are connected to
// the "best_friend" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithBestFriend(opts ...func(*PetQuery)) *PetQuery {
	query := (&PetClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withBestFriend = query
	return pq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithOwner(opts ...func(*UserQuery)) *PetQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pet.Query().
//		GroupBy(pet.FieldName).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (pq *PetQuery) GroupBy(field string, fields ...string) *PetGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PetGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Pet.Query().
//		Select(pet.FieldName).
//		Scan(ctx, &v)
func (pq *PetQuery) Select(fields ...string) *PetSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PetSelect{PetQuery: pq}
	sbuild.label = pet.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PetSelect configured with the given aggregations.
func (pq *PetQuery) Aggregate(fns ...AggregateFunc) *PetSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pet, error) {
	var (
		nodes       = []*Pet{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withBestFriend != nil,
			pq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pet{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withBestFriend; query != nil {
		if err := pq.loadBestFriend(ctx, query, nodes, nil,
			func(n *Pet, e *Pet) { n.Edges.BestFriend = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withOwner; query != nil {
		if err := pq.loadOwner(ctx, query, nodes, nil,
			func(n *Pet, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PetQuery) loadBestFriend(ctx context.Context, query *PetQuery, nodes []*Pet, init func(*Pet), assign func(*Pet, *Pet)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Pet)
	for i := range nodes {
		fk := nodes[i].BestFriendID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(pet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "best_friend_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PetQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Pet, init func(*Pet), assign func(*Pet, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Pet)
	for i := range nodes {
		fk := nodes[i].OwnerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pet.FieldID)
		for i := range fields {
			if fields[i] != pet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withBestFriend != nil {
			_spec.Node.AddColumnOnce(pet.FieldBestFriendID)
		}
		if pq.withOwner != nil {
			_spec.Node.AddColumnOnce(pet.FieldOwnerID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pet.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PetGroupBy is the group-by builder for Pet entities.
type PetGroupBy struct {
	selector
	build *PetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PetGroupBy) Aggregate(fns ...AggregateFunc) *PetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, fluent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetQuery, *PetGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PetGroupBy) sqlScan(ctx context.Context, root *PetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PetSelect is the builder for selecting fields of Pet entities.
type PetSelect struct {
	*PetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PetSelect) Aggregate(fns ...AggregateFunc) *PetSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, fluent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetQuery, *PetSelect](ctx, ps.PetQuery, ps, ps.inters, v)
}

func (ps *PetSelect) sqlScan(ctx context.Context, root *PetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}