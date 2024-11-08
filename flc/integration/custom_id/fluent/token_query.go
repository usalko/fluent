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
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/account"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/token"
	"github.com/usalko/fluent/flc/integration/custom_id/sid"
	"github.com/usalko/fluent/schema/field"
)

// TokenQuery is the builder for querying Token entities.
type TokenQuery struct {
	config
	ctx         *QueryContext
	order       []token.OrderOption
	inters      []Interceptor
	predicates  []predicate.Token
	withAccount *AccountQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TokenQuery builder.
func (tq *TokenQuery) Where(ps ...predicate.Token) *TokenQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TokenQuery) Limit(limit int) *TokenQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TokenQuery) Offset(offset int) *TokenQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TokenQuery) Unique(unique bool) *TokenQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TokenQuery) Order(o ...token.OrderOption) *TokenQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryAccount chains the current query on the "account" edge.
func (tq *TokenQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, token.AccountTable, token.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Token entity from the query.
// Returns a *NotFoundError when no Token was found.
func (tq *TokenQuery) First(ctx context.Context) (*Token, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, fluent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{token.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TokenQuery) FirstX(ctx context.Context) *Token {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Token ID from the query.
// Returns a *NotFoundError when no Token ID was found.
func (tq *TokenQuery) FirstID(ctx context.Context) (id sid.ID, err error) {
	var ids []sid.ID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, fluent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{token.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TokenQuery) FirstIDX(ctx context.Context) sid.ID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Token entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Token entity is found.
// Returns a *NotFoundError when no Token entities are found.
func (tq *TokenQuery) Only(ctx context.Context) (*Token, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, fluent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{token.Label}
	default:
		return nil, &NotSingularError{token.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TokenQuery) OnlyX(ctx context.Context) *Token {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Token ID in the query.
// Returns a *NotSingularError when more than one Token ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TokenQuery) OnlyID(ctx context.Context) (id sid.ID, err error) {
	var ids []sid.ID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, fluent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{token.Label}
	default:
		err = &NotSingularError{token.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TokenQuery) OnlyIDX(ctx context.Context) sid.ID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tokens.
func (tq *TokenQuery) All(ctx context.Context) ([]*Token, error) {
	ctx = setContextOp(ctx, tq.ctx, fluent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Token, *TokenQuery]()
	return withInterceptors[[]*Token](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TokenQuery) AllX(ctx context.Context) []*Token {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Token IDs.
func (tq *TokenQuery) IDs(ctx context.Context) (ids []sid.ID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, fluent.OpQueryIDs)
	if err = tq.Select(token.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TokenQuery) IDsX(ctx context.Context) []sid.ID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, fluent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TokenQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TokenQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, fluent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TokenQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TokenQuery) Clone() *TokenQuery {
	if tq == nil {
		return nil
	}
	return &TokenQuery{
		config:      tq.config,
		ctx:         tq.ctx.Clone(),
		order:       append([]token.OrderOption{}, tq.order...),
		inters:      append([]Interceptor{}, tq.inters...),
		predicates:  append([]predicate.Token{}, tq.predicates...),
		withAccount: tq.withAccount.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TokenQuery) WithAccount(opts ...func(*AccountQuery)) *TokenQuery {
	query := (&AccountClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withAccount = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Body string `json:"body,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Token.Query().
//		GroupBy(token.FieldBody).
//		Aggregate(fluent.Count()).
//		Scan(ctx, &v)
func (tq *TokenQuery) GroupBy(field string, fields ...string) *TokenGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TokenGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = token.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Body string `json:"body,omitempty"`
//	}
//
//	client.Token.Query().
//		Select(token.FieldBody).
//		Scan(ctx, &v)
func (tq *TokenQuery) Select(fields ...string) *TokenSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TokenSelect{TokenQuery: tq}
	sbuild.label = token.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TokenSelect configured with the given aggregations.
func (tq *TokenQuery) Aggregate(fns ...AggregateFunc) *TokenSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !token.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Token, error) {
	var (
		nodes       = []*Token{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withAccount != nil,
		}
	)
	if tq.withAccount != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, token.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Token).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Token{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withAccount; query != nil {
		if err := tq.loadAccount(ctx, query, nodes, nil,
			func(n *Token, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TokenQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*Token, init func(*Token), assign func(*Token, *Account)) error {
	ids := make([]sid.ID, 0, len(nodes))
	nodeids := make(map[sid.ID][]*Token)
	for i := range nodes {
		if nodes[i].account_token == nil {
			continue
		}
		fk := *nodes[i].account_token
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_token" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeOther))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for i := range fields {
			if fields[i] != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(token.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = token.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TokenGroupBy is the group-by builder for Token entities.
type TokenGroupBy struct {
	selector
	build *TokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TokenGroupBy) Aggregate(fns ...AggregateFunc) *TokenGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, fluent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TokenQuery, *TokenGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TokenGroupBy) sqlScan(ctx context.Context, root *TokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TokenSelect is the builder for selecting fields of Token entities.
type TokenSelect struct {
	*TokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TokenSelect) Aggregate(fns ...AggregateFunc) *TokenSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, fluent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TokenQuery, *TokenSelect](ctx, ts.TokenQuery, ts, ts.inters, v)
}

func (ts *TokenSelect) sqlScan(ctx context.Context, root *TokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}